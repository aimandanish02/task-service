//just fgor define contract for storage. so maybe later can swap from sqlite -> postgre, db -> redis or mock for tests
package repository

import "task-service/pkg/models"

type TaskRepository interface {
	Create(task models.Task) error
	GetByID(id string) (models.Task, error)
	Update(task models.Task) error
}


