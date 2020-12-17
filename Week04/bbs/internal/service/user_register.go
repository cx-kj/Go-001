package service

import "bbs/internal/biz"

type UserRegister struct {
	 useCase biz.UserRegisterCase
}

func NewUserRegister(useCase biz.UserRegisterCase) *UserRegister {
	return &UserRegister{
		useCase,
	}
}

func (ur *UserRegister) Register(register UserRegister) error {
	bizUr := biz.UserRegister{}
	_, err := ur.useCase.Register(bizUr)
	return err
}