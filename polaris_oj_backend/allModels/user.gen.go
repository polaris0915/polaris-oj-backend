// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package allModels

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "User"

// User 用户
type User struct {
	ID           int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:id" json:"id"`                                                                    // id
	Identity     string         `gorm:"column:identity;type:varchar(36);not null;uniqueIndex:unique_identity,priority:1;index:idx_identity,priority:1;comment:唯一ID" json:"identity"` // 唯一ID
	UserAccount  string         `gorm:"column:userAccount;type:varchar(256);not null;comment:账号" json:"userAccount"`                                                                 // 账号
	UserPassword string         `gorm:"column:userPassword;type:varchar(512);not null;comment:密码" json:"userPassword"`                                                               // 密码
	UnionID      string         `gorm:"column:unionId;type:varchar(256);index:idx_unionId,priority:1;comment:微信开放平台id" json:"unionId"`                                               // 微信开放平台id
	MpOpenID     string         `gorm:"column:mpOpenId;type:varchar(256);comment:公众号openId" json:"mpOpenId"`                                                                         // 公众号openId
	UserName     string         `gorm:"column:userName;type:varchar(256);comment:用户昵称" json:"userName"`                                                                              // 用户昵称
	UserAvatar   string         `gorm:"column:userAvatar;type:varchar(1024);comment:用户头像" json:"userAvatar"`                                                                         // 用户头像
	UserProfile  string         `gorm:"column:userProfile;type:varchar(512);comment:用户简介" json:"userProfile"`                                                                        // 用户简介
	UserRole     string         `gorm:"column:userRole;type:varchar(256);not null;default:user;comment:用户角色：user/admin/ban" json:"userRole"`                                         // 用户角色：user/admin/ban
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at"`                                                                     // 创建时间
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime;not null;comment:更新时间" json:"updated_at"`                                                                     // 更新时间
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deleted_at"`                                                                              // 删除时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
