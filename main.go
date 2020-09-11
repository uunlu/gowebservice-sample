package main

import (
	"net/http"

	"github.com/uunlu/gowebservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3001", nil)
}
