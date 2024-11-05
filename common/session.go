package common

import (
	"errors"
	"polaris-oj-backend/models/enums/userrole_enum"
	"polaris-oj-backend/utils"

	"github.com/gin-contrib/sessions"
)

func SetCookies(session sessions.Session, key, value string) error {
	session.Set(key, value) // 设置 session 中的值
	err := session.Save()   // 保存 session
	if err != nil {
		return errors.New("系统错误: cookie无法获取")
	}
	return nil
}

// 如果用户session中的token不能断言回字符串返回boolean
// 如果token检验失败返回error
// 没问题返回用户信息的*utils.Claims
func ParseUserInfoByToken(token interface{}) interface{} {
	tokenString, ok := token.(string)
	if !ok {
		return false
	}
	userInfo, err := utils.ValidateToken(tokenString)
	if err != nil {
		return err
	}
	return userInfo
}

func GetLoginUser(session sessions.Session) (*utils.Claims, error) {
	stringToken := session.Get(userrole_enum.USER_LOGIN_STATE) // 获取用户
	userInfo, ok := ParseUserInfoByToken(stringToken).(*utils.Claims)
	if !ok {
		return nil, errors.New("未登录或用户信息过期，请重新登录")
	}
	return userInfo, nil
}
