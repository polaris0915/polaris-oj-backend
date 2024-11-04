package service

import (
	"errors"
	"fmt"
	"polaris-oj-backend/common"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/database/mysql"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"gorm.io/gorm"
)

// 实例化 QuestionService提供给全局
var GQuestionService = NewQuestionService(mysql.DB)

// QuestionService 定义用户相关的服务
type QuestionService struct {
	db *gorm.DB
}

// NewQuestion 初始化 QuestionService
func NewQuestionService(db *gorm.DB) *QuestionService {
	return &QuestionService{db: db}
}

// TODO unfinished: 通过id获取问题详情
func (s *QuestionService) GetQuestionById(session sessions.Session, question *allModels.Question) error {
	// TODO unfinished: 需要编写具体逻辑
	// 获取当前登录用户
	var userInfo *utils.Claims
	var err error
	// 如果登录信息无效也就不能获取题目信息了
	if userInfo, err = common.GetLoginUser(session); err != nil {
		return err
	}

	// 查询题目信息，如果没有问题，信息也就在question中了
	if err = s.db.Preload("User").First(question, "identity = ?", question.Identity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	fmt.Printf("question: %+v\n", question)

	// 只有创建题目的作者以及管理员才可以查询到问题的详情
	if userInfo.UserRole != constant.ADMIN_ROLE && userInfo.Identity != question.User.Identity {
		return errors.New("权限不足")
	}
	return nil
}

/*
	TODO Unfinished: 整个问题添加以及修改的逻辑是需要完善的
	因为如果是用户本身去修改题目，肯定是要规范用户修改题目的规则的
	因此这里第一个构想是，用户提交修改请求之后，通知管理员进行审核
	管理员审核无误之后，再将数据进行修改
	创建题目也是同样的逻辑
*/
// 目前只是开发测试基本的功能是否能够接通数据库
// 更新问题
func (s *QuestionService) UpdateQuestion(session sessions.Session, question *allModels.Question, user *allModels.User) error {
	// 获取题目创建人的用户信息
	var userInfo *utils.Claims
	var err error
	if userInfo, err = common.GetLoginUser(session); err != nil {
		return err
	}
	// TODO: 在调用这个接口的时候就应该是已经通过中间件鉴权了，这边再次验证一下
	// 但是开发阶段可以先关闭
	// if userInfo.UserRole != constant.ADMIN_ROLE {
	// 	return errors.New("权限不足")
	// }

	// 查找修改题目的用户的信息
	if err = s.db.First(user, "identity = ?", userInfo.Identity).Error; err != nil {
		return err
	}
	// 更新问题业务
	dbQuestion := new(allModels.Question)
	// 根据问题的identity查找出问题
	if err := s.db.Model(dbQuestion).First(&dbQuestion, "identity = ?", question.Identity).Error; err != nil {
		return errors.New("没有该数据")
	}
	// TODO: 逻辑有问题
	// 更新dbQuestion即数据库中该问题需要更新的字段
	if err := utils.CopyModels(dbQuestion, question); err != nil {
		return err
	}

	return s.db.Save(dbQuestion).Error
}

// 添加问题
func (s *QuestionService) AddQuestion(session sessions.Session, question *allModels.Question, user *allModels.User) error {
	// 首先经过中间件在controller层排除未登录的用户到达次接口
	// 因此后面鉴权中间件加入之后就不再需要去校验用户是否登录，而是直接去取用户信息
	// 但是用户信息拿到之后如果实效过期也是得返回登录错误
	var userInfo *utils.Claims
	var err error
	if userInfo, err = common.GetLoginUser(session); err != nil {
		return err
	}
	if err = s.db.First(user, "identity = ?", userInfo.Identity).Error; err != nil {
		return err
	}

	// TODO unfinished: 需要改善更严格的问题内容校验
	// 例如题目标题不能重复等等
	if utils.IsAnyBlank(question.Title, question.Content, question.Answer) {
		return errors.New("参数不正确")
	}

	// 设置Identity
	question.Identity = utils.GetUUID()
	// 设置创建人Identity
	question.UserID = userInfo.Identity

	if err = s.db.Save(question).Error; err != nil {
		return err
	}

	return nil
}

// 删除问题
func (s *QuestionService) DeleteQuestion(question *allModels.Question) error {
	if utils.IsAnyBlank(question.Identity) {
		return errors.New("参数不正确")
	}

	// 已经删除就不要再删除了
	// 可以不要，万一gorm的First查询不自带校验过期时间就需要
	if question.DeletedAt.Valid {
		// fmt.Println(question.DeletedAt.Time.Before(time.Now()))
		if question.DeletedAt.Time.Before(time.Now()) {
			return nil
		}
	}

	res := s.db.Delete(question, "identity = ?", question.Identity)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected == 0 {
		return errors.New("已经删除")
	}
	return nil
}
