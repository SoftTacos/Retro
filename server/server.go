package server

import (
	"net/http"

	// httpTransport "github.com/go-kit/kit/transport/http"
	// endpoint "github.com/softtacos/retroBot/endpoint"

	httprouter "github.com/julienschmidt/httprouter"
	models "github.com/softtacos/retroBot/models"
	"github.com/softtacos/retroBot/util"
	// transport"github.com/softtacos/retroBot/transport"
)

// func StartHttpRetroServer() {
// 	// httpHandler := httpTransport.NewServer(
// 	// 	endpoint.MakeRetroEndpoint(service),
// 	// 	transport.ParseRetroRequest,
// 	// 	transport.EncodeHttpResponse,
// 	// )
// 	// http.Handle("/", httpHandler)
// 	webpage := "\nTHIS IS WEBPAGE BLYAT %q"
// 	http.HandleFunc("/index.html", WebpageHandler(webpage))
// 	//logger.Log("msg", "HTTP", "addr", port)
// 	go func() {
// 		for {
// 			errChan <- http.ListenAndServe(":"+port, nil)
// 		}
// 	}()
// }

func NewRetroServer(service models.RetroServices, port string, errChan chan error) *RetroServer {
	webpages := [][2]string{
		[2]string{"/index.html", string(util.LoadTextFile("html/index.html"))},
	}
	server := &RetroServer{
		router:  NewRouter(webpages),
		port:    port,
		errChan: errChan,
		service: service,
	}
	return server
}

type RetroServer struct {
	errChan chan error
	router  *httprouter.Router
	port    string
	service models.RetroServices
}

func (rs *RetroServer) Start() {
	go func() {
		for {
			rs.errChan <- http.ListenAndServe(rs.port, rs.router)
		}
	}()
}
