package core

type ProjectContact struct {
	UID        string `json:"uid"`
	ProjectUID string `json:"project_uid"`
	ContactUID string `json:"contact_uid"`
	IsActive   bool   `json:"is_active"`
}

func (p *ProjectContact) TableName() string {
	return "project_contacts"
}
