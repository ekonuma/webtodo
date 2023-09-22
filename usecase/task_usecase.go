package usecase

import (
	"github.com/ekonuma/webtodo/model"
	"github.com/ekonuma/webtodo/repository"
)

type ITaskUsecase interface{
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	taskRepository repository.ITaskRepository
}

func NewTaskUsecase(taskRepository repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{taskRepository}
}

func (taskUsecase *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := taskUsecase.taskRepository.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

func (taskUsercase *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := taskUsercase.taskRepository.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (taskUsercase *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := taskUsercase.taskRepository.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := taskUsercase.taskRepository.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (taskUsercase *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	if err := taskUsercase.taskRepository.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	if err := taskUsercase.taskRepository.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (taskUsercase *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := taskUsercase.taskRepository.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}