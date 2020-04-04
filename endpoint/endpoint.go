package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	models "github.com/softtacos/retroBot/models"
)

func MakeRetroEndpoint(next models.RetroServices) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//parse the request

		//call the service

		//return
		return &models.Response{
			Status: -1,
			Body:   "",
		}, nil //if err !=, then it won't return the response, just the error. We want to return everything.
	}
}
