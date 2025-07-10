package cron

import (
	"forum_backend/app/forum/cmd/api/internal/cron/task"
	"forum_backend/app/forum/cmd/api/internal/svc"
	"time"
)

func StartCronTasks(svcCtx *svc.ServiceContext) {
	go startDataDeleteCron(svcCtx)
}

func startDataDeleteCron(svcCtx *svc.ServiceContext) {
	ticker := time.NewTicker(time.Hour * 24)
	defer ticker.Stop()

	cleanUpTask := task.NewDataDeleteTask(svcCtx)
	for {
		select {
		case <-ticker.C:
			err := cleanUpTask.Run()
			if err != nil {
				cleanUpTask.Logger.Errorf("clean up task failed: %v", err)
				time.Sleep(time.Second * 10)
				continue
			}
		}
	}
}
