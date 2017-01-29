package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		result := ExecuteQuery(r.URL.Query()["query"][0], schema)
		json.NewEncoder(w).Encode(result)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
