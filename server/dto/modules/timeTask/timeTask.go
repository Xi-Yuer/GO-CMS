package timeTaskResponsiesModules

type StartTimeTack struct {
	TimeTaskID string `form:"timeTaskID" binding:"required"`
}

type StopTimeTack struct {
	TimeTaskID string `form:"timeTaskID" binding:"required"`
}

type DeleteTimeTack struct {
	TimeTaskID string `form:"timeTaskID" binding:"required"`
}

type CreateTimeTack struct {
	TaskName string `form:"taskName" binding:"required"`
	Cron     string `form:"cron" binding:"required"`
}

type UpdateTimeTask struct {
	TaskName string `form:"taskName" binding:"required"`
	Cron     string `form:"cron" binding:"required"`
	Status   *bool  `form:"status" binding:"required"`
}

type TimeTaskResponse struct {
	TimeTaskID  string  `json:"timeTaskID"`
	TaskName    string  `json:"taskName"`
	Cron        string  `json:"cron"`
	Status      bool    `json:"status"`
	LastRunTime string  `json:"lastRunTime"`
	RunTimes    float64 `json:"runTimes"`
}
