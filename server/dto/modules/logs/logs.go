package logsResponsiesModules

type CreateLogRecordRequest struct {
	ID              int64
	UserID          string
	UserName        string
	UserIP          string
	RequestPath     string
	RequestMethod   string
	RequestStatus   int
	RequestDuration string
	CreateTime      int64
}

type GetLogRecordResponse struct {
	ID              string `json:"id"`
	UserID          string `json:"userID"`
	UserName        string `json:"userName"`
	UserIP          string `json:"userIP"`
	RequestPath     string `json:"requestPath"`
	RequestMethod   string `json:"requestMethod"`
	RequestStatus   string `json:"requestStatus"`
	RequestDuration string `json:"requestDuration"`
	CreateTime      string `json:"createTime"`
}
