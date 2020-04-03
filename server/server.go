package server

import (
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
	endpoint "github.com/softtacos/retroBot/endpoint"
)

func NewHttpMessageServer(service models.MessagingService, port string, errChan chan error) {
	httpSendEmailHandler := httpTransport.NewServer(
		endpoint.MakeSendEmailEndpoint(service),
		transport.ParseHttpSendEmailRequest,
		transport.EncodeHttpResponse,
	)

	http.Handle("/", httpSendEmailHandler)

	//logger.Log("msg", "HTTP", "addr", port)
	go func() {
		for {
			errChan <- http.ListenAndServe(port, nil)
		}
	}()
}
