package models

type Job struct {
	JobID           int64        `json:"job_id,omitempty" url:"job_id,omitempty"`
	CreatorUserName string       `json:"creator_user_name,omitempty" url:"creator_user_name,omitempty"`
	Settings        *JobSettings `json:"settings,omitempty" url:"settings,omitempty"`
	CreatedTime     int64        `json:"created_time,omitempty" url:"created_time,omitempty"`
}
