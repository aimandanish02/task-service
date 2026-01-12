package service

import (
	"log"
	"sync"
	"time"

	"github.com/google/uuid"

	"task-service/internal/repository"
	"task-service/pkg/models"
)

type TaskService struct {
	repo  repository.TaskRepository
	queue chan models.Task
	wg    sync.WaitGroup
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo:  repo,
		queue: make(chan models.Task, 100),
	}
}

func (s *TaskService) CreateTask(title string) (models.Task, error) {
	task := models.Task{
		ID:        uuid.NewString(),
		Title:     title,
		Status:    models.StatusPending,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(task); err != nil {
		return task, err
	}

	s.queue <- task
	return task, nil
}

func (s *TaskService) StartWorkers(n int) {
	for i := 0; i < n; i++ {
		s.wg.Add(1)
		go s.worker(i)
	}
}


func (s *TaskService) worker(id int) {
	defer s.wg.Done()

	for task := range s.queue {
		log.Printf("worker %d processing task %s\n", id, task.ID)

		task.Status = models.StatusProcessing
		s.repo.Update(task)

		time.Sleep(2 * time.Second)

		task.Status = models.StatusCompleted
		s.repo.Update(task)

		log.Printf("worker %d completed task %s\n", id, task.ID)
	}
}


func (s *TaskService) GetTask(id string) (models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) Shutdown() {
	close(s.queue)
	s.wg.Wait()
}
