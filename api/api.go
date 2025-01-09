package api

import (
	"fmt"
	"main/count"
	threadpool "main/threadPool"
	"time"
)

var pool = threadpool.NewThreadPool(1, 1)

func Execute() string {
	uniqueID := fmt.Sprintf("%d", count.IncrementCount())
	mp.set(uniqueID, Queued)
	err := pool.Execute(threadpool.Task{
		UniqueID: uniqueID,
		Func: func(uniqueID string) {
			mp.set(uniqueID, Processing)
			time.Sleep(10 * time.Second)
			mp.set(uniqueID, Completed)
		},
	})
	if err != nil {
		return err.Error()
	}
	return uniqueID
}

func Poll(uniqueID string) State {
	return mp.get(uniqueID)
}
