package main

import (
	"distributed-job-scheduler/scheduler"
	"fmt"
	"time"
)

func main() {

	store := scheduler.NewInMemoryJobStore()

	job := &scheduler.Job{
		ID:        "job1",
		TaskType:  "print_hello",
		Payload:   []byte(`{"msg":"hello"}`),
		RunAt:     time.Now(),
		Status:    scheduler.StatusPending,
		CreatedAt: time.Now(),
	}

	err := store.Add(job)
	if err != nil {
		panic(err)
	}

	fmt.Println(" Job added")

	// Get
	fetched, err := store.Get("job1")
	if err != nil {
		panic(err)
	}

	fmt.Println("📦 Fetched Job:", fetched.ID, fetched.Status)

	// Update
	fetched.Status = scheduler.StatusRunning
	err = store.Update(fetched)
	if err != nil {
		panic(err)
	}

	fmt.Println("🔄 Job updated to:", fetched.Status)

	// List Pending
	jobs, _ := store.ListPending(time.Now())
	fmt.Println("📋 Pending Jobs Count:", len(jobs))
}
