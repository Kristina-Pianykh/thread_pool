package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func randomDuration(maxSec int) time.Duration {
	return time.Duration(rand.Float64() * float64(maxSec) * float64(time.Second))
}

func TestMultipleWorkersDoDifferentWork(t *testing.T) {
	totalTasks := 10
	results := make(chan string, totalTasks)
	expResults := []string{}
	tasks := []func(){}

	for i := range totalTasks {
		task := func() {
			time.Sleep(randomDuration(5))
			results <- fmt.Sprintf("task %d", i)
		}
		tasks = append(tasks, task)

		expResults = append(expResults, fmt.Sprintf("task %d", i))
	}

	pool := CreateThreadPool(5)
	for _, task := range tasks {
		pool.Execute(task)
	}
	pool.WaitTasks()
	require.Len(t, results, totalTasks)

	for range totalTasks {
		require.Contains(t, expResults, <-results)
	}
}
