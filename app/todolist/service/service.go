package service

import (
	"goinspect.cn/app/conf"
	"goinspect.cn/app/todolist"
	"goinspect.cn/app/todolist/dao"
)

type Service struct {
	c *conf.Config

	// dao
	taskDao *dao.Dao
}

func New(c *conf.Config) (s *Service) {
	s = &Service {
		c: c,
		taskDao: dao.New(c),
	}
	return
}

func (s *Service) GetUserTasks(uid int) (todo []*todolist.Task, done []*todolist.Task, err error) {
	res, err := s.taskDao.GetTasks(uid)
	if err != nil {
		return
	}
	todo = []*todolist.Task{}
	done = []*todolist.Task{}
	for _, t := range res {
		if t.Status == todolist.STATUS_PENDING {
			todo = append(todo, t)
		} else {
			done = append(done, t)
		}
	}

	return
}

func (s *Service) NewTask(uid int, title string) bool {
	return s.taskDao.NewTask(uid, title)
}

func (s *Service) SetTaskDone(id int) bool {
	return s.taskDao.UpdateTaskStatus(id, todolist.STATUS_DONE)
}

func (s *Service) RemoveTask(id int) bool {
	return s.taskDao.RemoveTask(id)
}