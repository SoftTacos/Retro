package server

import (
	"fmt"
	"html"
	"net/http"

	models "github.com/softtacos/retroBot/models"
)

func StartHttpRetroServer(service models.RetroServices, port string, errChan chan error) {
	// httpHandler := httpTransport.NewServer(
	// 	endpoint.MakeRetroEndpoint(service),
	// 	transport.ParseRetroRequest,
	// 	transport.EncodeHttpResponse,
	// )
	// http.Handle("/", httpHandler)

	webpage := []byte("THIS IS WEBPAGE BLYAT")
	http.HandleFunc("/index.html", WebpageHandler(webpage))
	//logger.Log("msg", "HTTP", "addr", port)
	go func() {
		for {
			errChan <- http.ListenAndServe(":"+port, nil)
		}
	}()
}

func WebpageHandler(webpage []byte) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}
}
