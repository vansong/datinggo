package service

import (
	"fmt"
	"goinspect.cn/app/account"
	"goinspect.cn/app/account/dao"
	"goinspect.cn/app/conf"
)

type Service struct {
	c *conf.Config

	// dao
	userDao *dao.Dao
}

func New(c *conf.Config) (s *Service) {
	s = &Service {
		c: c,
		userDao: dao.New(c),
	}
	return
}

func (s *Service) GetUser(id int) *account.User {
	u, err := s.userDao.GetUser(id);
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return u
}