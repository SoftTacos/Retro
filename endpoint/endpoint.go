package endpoint

import (
	"context"
)

func MakeSendEmailEndpoint(svc services.MessagingService) endpoint.Endpoint {

	return func(ctx context.Context, sendEmailRequest interface{}) (interface{}, error) {
		sendEmailRequestAssert := sendEmailRequest.(*models.SendEmailRequest)
		status, msg, err := svc.SendEmail(sendEmailRequestAssert)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		return &models.SendEmailResponse{
			Status:  status,
			Message: msg,
			Err:     errString,
		}, nil //if err !=, then it won't return the response, just the error. We want to return everything.
	}
}
