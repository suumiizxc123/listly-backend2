package core

type ProjectRoom struct {
	UID        string `json:"uid"`
	ProjectUID string `json:"project_uid"`
	Room       string `json:"room"`
	FloorPlan  string `json:"floor_plan"`

	IsActive bool `json:"is_active"`
}

func (p *ProjectRoom) TableName() string {
	return "project_rooms"
}
