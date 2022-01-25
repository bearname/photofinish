package tasks

import (
	"github.com/pkg/errors"
	"photofinish/pkg/common/util"
	"photofinish/pkg/domain/tasks"
)

type ServiceImpl struct {
	repo tasks.Repository
}

func NewService(eventRepo tasks.Repository) *ServiceImpl {
	s := new(ServiceImpl)
	s.repo = eventRepo
	return s
}

func (s *ServiceImpl) GetStatistics(taskId string) (*tasks.TaskStats, error) {
	if !util.IsUUID(taskId) {
		return nil, errors.New("invalid task id")
	}

	return s.repo.GetStatsByTask(taskId)
}
