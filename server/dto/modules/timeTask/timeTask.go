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
	TimeTaskName string `form:"timeTaskName" binding:"required"`
	Cron         string `form:"cron" binding:"required"`
}

type UpdateTimeTask struct {
	TimeTaskName string `form:"timeTaskName" binding:"required"`
	Cron         string `form:"cron" binding:"required"`
	Status       bool   `form:"status" binding:"required"`
}

type TimeTaskResponse struct {
	TimeTaskID   string  `json:"timeTaskID"`
	TimeTaskName string  `json:"taskName"`
	Cron         string  `json:"cron"`
	Status       bool    `json:"status"`
	LastRunTime  string  `json:"lastRunTime"`
	RunTimes     float64 `json:"runTimes"`
}
