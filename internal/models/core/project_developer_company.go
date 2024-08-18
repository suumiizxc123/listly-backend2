package core

type ProjectDeveloperCompany struct {
	UID          string `json:"uid"`
	ProjectUID   string `json:"project_uid"`
	DeveloperUID string `json:"developer_uid"`
	IsActive     bool   `json:"is_active"`
}

func (p *ProjectDeveloperCompany) TableName() string {
	return "project_developer_companies"
}
