package core

type ProjectFaq struct {
	UID        string `json:"uid"`
	ProjectUID string `json:"project_uid"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	IsActive   bool   `json:"is_active"`
}

func (p *ProjectFaq) TableName() string {
	return "project_faqs"
}
