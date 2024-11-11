package question_vo

import (
	judgeconfig "polaris-oj-backend/models/dto/judgeconfig"
	"polaris-oj-backend/models/vo/user_vo"

	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type QuestionVO struct {
	Identity    string                  `json:"identity"`    // id
	Title       string                  `json:"title"`       // 标题
	Content     string                  `json:"content"`     // 内容
	Tags        []string                `json:"tags"`        // 标签列表
	SubmitNum   int                     `json:"submitNum"`   // 题目提交数
	AcceptedNum int                     `json:"acceptedNum"` // 题目通过数
	JudgeConfig judgeconfig.JudgeConfig `json:"judgeConfig"` // 判题配置（json 对象）
	ThumbNum    int                     `json:"thumbNum"`    // 点赞数
	FavourNum   int                     `json:"favourNum"`   // 收藏数
	UserID      string                  `json:"userId"`      // 创建用户 id
	// TODO question: 这边时间的名字没有对应上，不知道会不会出问题
	CreatedAt time.Time      `json:"createTime"` // 创建时间
	UpdatedAt time.Time      `json:"updateTime"` // 更新时间
	UserVO    user_vo.UserVO `json:"userVO"`     // 创建题目人的信息
}

func (u *QuestionVO) GetValidator() *validator.Validate {
	return validator.New()
}

func (u *QuestionVO) GetResponseVo(question *allModels.Question) error {
	// ============能直接拷贝的字段=======================
	if err := copier.Copy(u, question); err != nil {
		return err
	}
	// ============自定义字段转换为json字符串格式===============
	// 需要将问题表中的json字符串转换回QuestionVO的字段
	utils.JsonToModel(question.JudgeConfig, u.JudgeConfig)
	utils.JsonToModel(question.Tags, u.Tags)

	// 将question中的User脱敏
	var userVo = new(user_vo.UserVO)
	if err := userVo.GetResponseVo(question.User); err != nil {
		return err
	}
	u.UserVO = *userVo

	// 校验
	if err := u.GetValidator().Struct(u); err != nil {
		return err
	}
	return nil
}
