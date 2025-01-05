package main

import (
	"net/http"

	"github.com/louislef299/go-scripts/projects/bare_webserver/pkg/router"
)

func main() {
	r := router.NewRouter()
	r.GET("/", func(r *http.Request) (statusCode int, data map[string]interface{}) {
		return 200, map[string]interface{}{
			"name": "elonmux",
		}
	})

	http.ListenAndServe(":8080", r)
}
