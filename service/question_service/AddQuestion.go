package questionservice

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"polaris-oj-backend/common"

	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"
)

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
