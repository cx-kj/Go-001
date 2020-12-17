package biz

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

/*
type BizError struct {
	Code int // http code
	ErrCode int // 业务错误码
	ErrMessage string // 错误消息
}
*/

type UserRegister struct {
	Username string
	Password string
}

// 数据访问层接口
type UserRegisterRepo interface {
	GetByUserName(username string) (uid int, err error)
	Register(register UserRegister) (uid int, err error)
}

type UserRegisterCase struct {
	repo UserRegisterRepo
}

// 业务层实现
func NewUserRegisterCase(repo UserRegisterRepo) *UserRegisterCase {
	return &UserRegisterCase{
		repo: repo,
	}
}

func (urc *UserRegisterCase) Register(ur UserRegister) (uid int, err error) {
	uid, err = urc.repo.GetByUserName(ur.Username)
	if err != sql.ErrNoRows && err != nil {
		// 使用自定义 BusinessError 更好
		err = errors.Wrap(err, "数据库出错了")
		return
	}
	if err == nil {
		// 使用自定义 BusinessError 更好
		err = errors.Wrap(err, fmt.Sprintf("用户已经存在: %d", uid))
		return
	}
	uid, err = urc.Register(ur)
	if err != nil {
		// 使用自定义 BusinessError 更好
		err = errors.Wrap(err, "用户注册失败")
		return
	}
	return uid, nil
}