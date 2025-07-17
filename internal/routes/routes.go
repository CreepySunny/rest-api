package routes

import (
	"net/http"
	CalcAddHandler "github.com/CreepySunny/rest-api/internal/handlers"
)

func RegisterRoutes(router *http.ServeMux) {
		router.HandleFunc("POST /add", CalcAddHandler.CalcAddHandler)
}
