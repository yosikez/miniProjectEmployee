package input

type JsonDataOpportunity struct {
	Code            string     `json:"code" binding:"required"`
	ClientCode      string     `json:"client_code" binding:"required"`
	PicEmail        string     `json:"pic_email" binding:"required,email"`
	OpportunityName string     `json:"opportunity_name" binding:"required"`
	Description     string     `json:"description" binding:"required"`
	SalesEmail      string     `json:"sales_email" binding:"required,email"`
	Status          string     `json:"status" binding:"required"`
	LastModified    string     `json:"last_modified" binding:"required,datetime=2006-01-02 15:04:05"`
	Resources       []Resource `json:"resources" binding:"required,dive"`
}

type Resource struct {
	Qty             int64    `json:"qty" binding:"required"`
	Position        string   `json:"position" binding:"required"`
	Level           string   `json:"level" binding:"required"`
	Ctc             *float64 `json:"ctc"`
	ProjectDuration int64    `json:"project_duration" binding:"required"`
}