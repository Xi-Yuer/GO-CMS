package timeTask

import (
	"errors"
	"github.com/robfig/cron/v3"
)

var c *cron.Cron

var TimeTask = &timeTask{
	taskMap: make(map[string]struct {
		TaskID cron.EntryID
		Task   *cron.Cron
	}),
}

type timeTask struct {
	taskMap map[string]struct {
		TaskID cron.EntryID
		Task   *cron.Cron
	}
}

// AddTask 添加任务
func (t *timeTask) AddTask(TimeTaskID string, corn string, task func(), status bool) error {
	c = cron.New(cron.WithSeconds())
	id, err := c.AddFunc(corn, task)
	if status {
		go c.Start()
	}
	if err != nil {
		return err
	}
	t.taskMap[TimeTaskID] = struct {
		TaskID cron.EntryID
		Task   *cron.Cron
	}{
		TaskID: id,
		Task:   c,
	}
	return nil
}

// StartTask 启动任务
func (t *timeTask) StartTask(TimeTaskID string) error {
	if task, ok := t.taskMap[TimeTaskID]; ok {
		go task.Task.Start()
		return nil
	} else {
		return errors.New("任务不存在")
	}
}

// StopTask 停止任务
func (t *timeTask) StopTask(TimeTaskID string) error {
	if task, ok := t.taskMap[TimeTaskID]; ok {
		task.Task.Stop()
		return nil
	} else {
		return errors.New("任务不存在")
	}
}

// RemoveTask 移除任务
func (t *timeTask) RemoveTask(TimeTaskID string) error {
	if task, ok := t.taskMap[TimeTaskID]; ok {
		task.Task.Stop()
		c.Remove(task.TaskID)
		delete(t.taskMap, TimeTaskID)
	} else {
		return errors.New("任务不存在")
	}
	return nil
}

// ParseCron 校验 cron 表达式是否正确
func (t *timeTask) ParseCron(s string) error {
	c := cron.New(cron.WithSeconds())
	id, err := c.AddFunc(s, func() {})
	if err != nil {
		return errors.New("cron 表达式错误")
	} else {
		c.Remove(id)
	}
	return nil
}
