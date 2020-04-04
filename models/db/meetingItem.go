package models

type RetroMeetingItem struct {
	ID        uint
	MeetingID uint
	UserID    uint
	Body      string
	Sentiment uint16

	Meeting *RetroMeeting
}
