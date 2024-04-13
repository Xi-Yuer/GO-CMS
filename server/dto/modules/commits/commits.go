package commitsResponsiesModules

type CommitResponse struct {
	CommitID string `json:"commitID"`
	Email    string `json:"email"`
	Author   string `json:"author"`
	Message  string `json:"message"`
	Date     string `json:"date"`
}
