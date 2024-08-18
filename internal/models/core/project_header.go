package core

type ProjectHeader struct {
	UID string `json:"uid"`

	ProjectUID string `json:"project_uid"`

	URL string `json:"url"`

	IsActive bool `json:"is_active"`
}

func (p *ProjectHeader) TableName() string {
	return "project_header"
}
