package models

type RetroServices interface {
	HandleRetroRequest(*RetroRequest) *Response
}
