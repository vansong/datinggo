package assess

import (
	"goinspect.cn/app/questionair/model/paper"
)

func (s *Service) GetPaperPaginated(page int, pageSize int) (res []*paper.Paper, err error) {
	offset := (page -1 ) * pageSize
	res, err = s.paperDao.GetPapers(offset, pageSize)
	return
}

func (s *Service) GetPaper(id int) (res *paper.Paper, err error) {
	res, err = s.paperDao.GetPaper(id)
	return
}