package core

type ProjectLocationNearby struct {
	UID        string  `json:"uid"`
	ProjectUID string  `json:"project_uid"`
	Distance   float32 `json:"distance"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	IsActive   bool    `json:"is_active"`
}
