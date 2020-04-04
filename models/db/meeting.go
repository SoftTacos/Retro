package models

import "time"

type RetroMeeting struct {
	ID   uint
	Date *time.Time

	MeetingItems []*RetroMeetingItem
}
