package core

type ProjectAgent struct {
	UID        string `json:"uid"`
	ProjectUID string `json:"project_uid"`
	AgencyUID  string `json:"agency_uid"`
	IsActive   bool   `json:"is_active"`
}

func (p *ProjectAgent) TableName() string {
	return "project_agents"
}
