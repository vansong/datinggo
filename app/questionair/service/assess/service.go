package assess

import (
	"goinspect.cn/app/conf"
	papaerdao "goinspect.cn/app/questionair/dao/paper"
)
type Service struct {
	c *conf.Config

	// dao
	paperDao *papaerdao.Dao
}

func New(c *conf.Config) (s *Service) {
	s = &Service {
		c:        c,
		paperDao: papaerdao.New(c),
	}
	return
}