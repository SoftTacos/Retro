package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/softtacos/retroBot/models"
)

func ParseHttpSendEmailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	sendEmailRequest := &models.RetroRequest{}
	uri := r.URL.RequestURI()
	defer r.Body.Close()
	sendEmailRequest.URI = strings.Split(uri, "/")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error decoding SendEmailRequest ", err)
		return "Error decoding SendEmailRequest ", err
	}
	sendEmailRequest.Body = string(bodyBytes)

	return sendEmailRequest, nil
}

func EncodeHttpResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
