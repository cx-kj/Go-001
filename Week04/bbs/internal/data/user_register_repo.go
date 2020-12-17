package data

import "bbs/internal/biz"

type DB interface {
	Find(field string, value string) (uid int, err error)
	Register(map[string]string) (uid int, err error)
}

type UserRegisterRepo struct {
	db DB
}

func (r *UserRegisterRepo) GetByUserName(username string) (uid int, err error) {
	return r.db.Find("username", username)
}

func (r *UserRegisterRepo) Register(register biz.UserRegister) (uid int, err error) {
	m := map[string]string{
		"username": register.Username,
		"password": register.Password,
	}
	return r.db.Register(m)
}