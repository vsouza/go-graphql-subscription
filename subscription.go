package main

import (
	"fmt"
	"rand"

	"github.com/graphql-go/graphql"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data map[string]user

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"count": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var subscriptionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Subscription",
		Fields: graphql.Fields{
			"count": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return fmt.Print(rand.Intn(100)), nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Subscription: subscriptionType,
	},
)

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
