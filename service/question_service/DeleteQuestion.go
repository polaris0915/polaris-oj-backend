package questionservice

import (
	"errors"
	"time"

	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"
)

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
