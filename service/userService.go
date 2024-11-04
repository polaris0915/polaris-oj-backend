package service

import (
	"errors"

	"polaris-oj-backend/common"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/database/mysql"

	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"

	"github.com/gin-contrib/sessions"
	"gorm.io/gorm"
)

// 实例化 UserService提供给全局
var GUserService = NewUserService(mysql.DB)

// UserService 定义用户相关的服务
type UserService struct {
	db *gorm.DB
}

// NewUserService 初始化 UserService
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// 注册用户
func (s *UserService) RegisterUser(user *allModels.User) error {
	if err := validateRegistrationInput(user.UserAccount, user.UserPassword); err != nil {
		return err
	}
	var cnt int64
	if err := s.db.Model(user).Where("userAccount = ?", user.UserAccount).Count(&cnt).Error; err != nil || cnt > 0 {
		return errors.New("账号已存在")
	}
	user.Identity = utils.GetUUID()
	user.UserPassword = utils.GetMd5(user.UserPassword)
	user.UserRole = constant.DEFAULT_ROLE

	if err := s.db.Create(user).Error; err != nil {
		return errors.New("创建用户失败")
	}
	return nil
}

// 用户登录
func (s *UserService) LoginUser(user *allModels.User) (string, error) {
	if utils.IsAnyBlank(user.UserPassword, user.UserAccount) {
		return "", errors.New("账号或者密码不能为空")
	}
	password := utils.GetMd5(user.UserPassword)

	if err := s.db.Where("userAccount = ? AND userPassword = ?", user.UserAccount, password).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("账号或密码错误")
		}
		return "", errors.New("系统错误: 查询失败")
	}

	return utils.GetToken(user.Identity, user.UserAccount, user.UserRole)
}

// 获取当前登录用户信息
func (s *UserService) GetLoginUser(session sessions.Session, user *allModels.User) error {
	// 问题是如何获取用户信息呢？仅凭一个cookie
	var userInfo *utils.Claims
	var err error
	// 通过common.GetLoginUser解析用户信息，如果没有则是过期
	if userInfo, err = common.GetLoginUser(session); err != nil {
		return err
	}
	// 查询用户信息
	// TODO unfinished: 可以引入redis，提升性能
	identity := userInfo.Identity

	if err := mysql.DB.Model(&user).First(user, "identity = ?", identity).Error; err != nil {
		return errors.New("系统错误")
	}
	return nil
}

// 验证用户注册信息的输入
func validateRegistrationInput(userAccount, userPassword string) error {
	if utils.IsAnyBlank(userAccount, userPassword) {
		return errors.New("参数缺失")
	}
	if len(userAccount) < 4 {
		return errors.New("账号长度不足")
	}
	if len(userPassword) < 8 {
		return errors.New("密码长度不足")
	}
	return nil
}

// TODO: 逻辑有问题
func (s *UserService) UpdateMyUser(session sessions.Session, user *allModels.User) error {
	// 首先判断用户是否自己已经登录，如果没登录则返回
	userInfo := new(utils.Claims)
	var err error
	if userInfo, err = common.GetLoginUser(session); err != nil {
		return errors.New("登录信息过期，请重新登录")
	}

	dbUser := new(allModels.User)
	if err = s.db.First(dbUser, "identity = ?", userInfo.Identity).Error; err != nil {
		return errors.New(constant.SYSTEM_ERROR.Message)
	}

	if err = utils.CopyModels(dbUser, user); err != nil {
		return errors.New(constant.PARAMS_ERROR.Message)
	}
	if err = s.db.Save(dbUser).Error; err != nil {
		return errors.New(constant.PARAMS_ERROR.Message)
	}
	*user = *dbUser
	return nil

}
