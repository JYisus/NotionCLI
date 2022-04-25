package task

import (
	"context"
	"errors"
	"github.com/jyisus/notioncli/entity"
)

type Service struct {
	taskRepository  entity.TaskRepository
	databases       []entity.Database
	defaultDatabase entity.Database
}

func NewService(cfg entity.Config, taskRepository entity.TaskRepository) Service {
	var defaultDatabase entity.Database
	for _, db := range cfg.Databases {
		if db.Name == cfg.DefaultDatabase {
			defaultDatabase = db
			break
		}
	}
	return Service{
		taskRepository:  taskRepository,
		databases:       cfg.Databases,
		defaultDatabase: defaultDatabase,
	}
}

func (s Service) AddTask(databaseName, task string) error {
	database, err := s.getDatabaseByName(databaseName)
	if err != nil {
		return err
	}

	return s.taskRepository.AddTask(context.TODO(), database, task)
}

func (s Service) ListTasks(writer entity.Writer, databaseName string) error {
	database, err := s.getDatabaseByName(databaseName)
	if err != nil {
		return err
	}

	tasks, err := s.taskRepository.ListTasks(context.TODO(), database)
	if err != nil {
		return err
	}

	//for index, task := range tasks {
	//	fmt.Printf("%d. %s # %s\n", index+1, task.Text, task.Id)
	//}

	writer.PrintTasks(tasks)

	return nil
}

func (s Service) DeleteTask(taskId string) error {
	return s.taskRepository.DeleteTask(context.TODO(), taskId)
}

func (s Service) getDatabaseByName(databaseName string) (entity.Database, error) {
	if databaseName == "default" {
		return s.defaultDatabase, nil
	}
	for _, db := range s.databases {
		if db.Name == databaseName {
			return db, nil
		}
	}
	return entity.Database{}, errors.New("unable to find database in configuration file")
}
