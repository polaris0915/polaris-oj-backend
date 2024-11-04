package vo

import (
	"github.com/gin-gonic/gin"
)

// ------------------BaseRequest Model------------------
type BaseRequest struct {
}

// ------------------BaseResponse Model------------------
// 最基础的响应
type BaseResponse struct {
	Err     error
	Code    int         `json:"code"`    // 响应码
	Data    interface{} `json:"data"`    // 接口，表示具体信息
	Message string      `json:"message"` // 请求结果[发生错误则是错误信息，如果没有错误则是SUCCESS.Code的值]
}

// 嵌套是为了自定义新的响应，方便以后拓展
type SubResponse struct {
	BaseResponse
	C      *gin.Context
	Status int
}

func NewSubResponse(c *gin.Context, status int, code int, data interface{}, message string) *SubResponse {
	return &SubResponse{
		BaseResponse: BaseResponse{
			Code:    code,
			Data:    data,
			Message: message,
		},
		C:      c,
		Status: status,
	}
}

func (br *SubResponse) Response() {
	br.C.JSON(br.Status, gin.H{
		"code":    br.Code,
		"data":    br.Data,
		"message": br.Message,
	})
}
