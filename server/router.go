package server

import (
	"fmt"
	"html"
	"net/http"

	httprouter "github.com/julienschmidt/httprouter"
	handlers "github.com/softtacos/retroBot/server/handlers"
)

func NewRouter(webpages [][2]string) *httprouter.Router {
	router := httprouter.New()
	for i := range webpages {
		router.GET(webpages[i][0], PageHandler(webpages[i][1]))
	}
	socketHandler := handlers.NewSocketManager()
	router.GET("/join/:meeting", socketHandler.AddConnection)
	return router
}

func PageHandler(webpage string) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, webpage, html.EscapeString(r.URL.Path))
	}
}
