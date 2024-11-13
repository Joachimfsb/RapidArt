package models

import "time"

type Report struct {
	ReportId         int       `json:"report_id"`
	UserId           int       `json:"user_id,omitempty"`
	PostId           int       `json:"post_id,omitempty"`
	Message          string    `json:"message,omitempty"`
	CreationDateTime time.Time `json:"creation_date_time,omitempty"`
}
