package transport

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	models "github.com/softtacos/retroBot/models"
)

func ParseRetroRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	request := &models.RetroRequest{}
	uri := r.URL.RequestURI()
	defer r.Body.Close()
	request.URI = strings.Split(uri, "/")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error decoding Request ", err)
		return "Error decoding Request ", err
	}
	request.Body = bodyBytes

	return request, nil
}

func EncodeHttpResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
