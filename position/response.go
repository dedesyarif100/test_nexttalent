package position

type PositionResponse struct {
	ID                   int64  `json:"id"`
	Name                 string `json:"name"`
	JobDescription       string `json:"job_description"`
	MinimumQualification string `json:"minimum_qualification"`
	MinimumExperience    string `json:"minimum_experience"`
	Skills               string `json:"skills"`
	Benefit              string `json:"benefit"`
}
