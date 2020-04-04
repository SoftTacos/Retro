package server

import (
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
	endpoint "github.com/softtacos/retroBot/endpoint"
	models "github.com/softtacos/retroBot/models"
	"github.com/softtacos/retroBot/transport"
)

func StartHttpRetroServer(service models.RetroServices, port string, errChan chan error) {
	httpHandler := httpTransport.NewServer(
		endpoint.MakeRetroEndpoint(service),
		transport.ParseRetroRequest,
		transport.EncodeHttpResponse,
	)

	http.Handle("/", httpHandler)

	//logger.Log("msg", "HTTP", "addr", port)
	go func() {
		for {
			errChan <- http.ListenAndServe(":"+port, nil)
		}
	}()
}
