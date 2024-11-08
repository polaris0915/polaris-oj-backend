package vo

import (
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ------------------BaseResponse Model------------------
// 最基础的响应
type BaseResponse[T any] struct {
	Code    int    `json:"code"`    // 响应码
	Data    *T     `json:"data"`    // 接口，表示具体信息
	Message string `json:"message"` // 请求结果[发生错误则是错误信息，如果没有错误则是SUCCESS.Code的值]
}

func (b *BaseResponse[T]) Response(c *gin.Context, status int) {
	// if b.Data == nil {
	// 	polaris_log.Logger.Errorf(b.Message)
	// }
	c.JSON(status, *b)
}

// /////////////////////////////////////////////////////////
// ~ 表示支持类型的衍生类型
// | 表示取并集
// 多行之间取交集
type mysqlModels interface {
	~*allModels.User | ~*allModels.Question | ~*allModels.QuestionSubmit | []*allModels.User | []*allModels.Question | []*allModels.QuestionSubmit
}

type ResponVo[T mysqlModels] interface {
	GetResponseVo(model T) error
	GetValidator() *validator.Validate
}

type PageVo struct {
	CountID          string           `json:"countId"`
	Current          int32            `json:"current"`
	MaxLimit         int32            `json:"MaxLimit"`
	OptimizeCountSql bool             `json:"optimizeCountSql"`
	Orders           []map[string]any `json:"oders"`
	Pages            int32            `json:"pages"`
	SearchCount      bool             `json:"searchCount"`
	Size             int32            `json:"size"`
	Total            int64            `json:"total"`
}

func (u *PageVo) GetPageVO(info map[string]any) {
	if CountID, ok := info["CountID"].(string); !ok {
		u.CountID = CountID
	}
	if Current, ok := info["Current"].(int32); !ok {
		u.Current = Current
	}
	if MaxLimit, ok := info["MaxLimit"].(int32); !ok {
		u.MaxLimit = MaxLimit
	}
	if OptimizeCountSql, ok := info["OptimizeCountSql"].(bool); !ok {
		u.OptimizeCountSql = OptimizeCountSql
	}
	if Pages, ok := info["Pages"].(int32); !ok {
		u.Pages = Pages
	}
	if SearchCount, ok := info["SearchCount"].(bool); !ok {
		u.SearchCount = SearchCount
	}
	if Size, ok := info["Size"].(int32); !ok {
		u.Size = Size
	}
	if Total, ok := info["Total"].(int64); !ok {
		u.Total = Total
	}
	if Orders, ok := info["Orders"].([]map[string]any); !ok {
		u.Orders = Orders
	}
}
