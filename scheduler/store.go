package scheduler

import (
	"fmt"
	"time"
)

type InMemoryJobStore struct {
	jobs map[string]*Job
}

type JobStore interface {
	Add(job *Job) error
	Get(id string) (*Job, error)
	Update(job *Job) error
	ListPending(now time.Time) ([]*Job, error)
}

func NewInMemoryJobStore() *InMemoryJobStore {
	return &InMemoryJobStore{
		jobs: make(map[string]*Job),
	}
}

func (s *InMemoryJobStore) Add(job *Job) error {
	if _, exists := s.jobs[job.ID]; exists {
		return fmt.Errorf("job already exists")
	}

	s.jobs[job.ID] = job
	return nil
}

func (s *InMemoryJobStore) Get(id string) (*Job, error) {
	job, exists := s.jobs[id]
	if !exists {
		return nil, fmt.Errorf("job not found")
	}

	return job, nil

}

func (s *InMemoryJobStore) Update(job *Job) error {
	if _, exists := s.jobs[job.ID]; !exists {
		return fmt.Errorf("job not found")
	}

	s.jobs[job.ID] = job
	return nil
}

func (s *InMemoryJobStore) ListPending(now time.Time) ([]*Job, error) {
	var result []*Job

	for _, job := range s.jobs {
		if job.Status == StatusPending && job.RunAt.Before(now) {
			result = append(result, job)
		}
	}

	return result, nil
}
