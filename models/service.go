package models

type RetroRequestHandleFunc func(*RetroRequest) *Response

type RetroServices interface {
	HandleRetroRequest(*RetroRequest) *Response
}
