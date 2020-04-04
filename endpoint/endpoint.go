package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	models "github.com/softtacos/retroBot/models"
)

func MakeRetroEndpoint(next models.RetroServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//type assert the request
		req := request.(*models.RetroRequest)
		//call the service
		res := next.HandleRetroRequest(req)
		//return
		return res, nil //if err !=, then it won't return the response, just the error. We want to return everything.
	}
}
