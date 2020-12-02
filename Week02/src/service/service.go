package service

import(
	"github.com/pkg/errors"
	"dao"
)

type Service struct{
	dao *dao.Dao
}

func NewService() *Service {
	return &Service{dao.NewDao()}
}

func (s *Service) GetUsernameByUserId(id int) (u dao.User, err error) {
	s = NewService()
	u, err = s.dao.FindUserById(id)
	return u, errors.Wrapf(err, "service -%d- not match", id)
}
