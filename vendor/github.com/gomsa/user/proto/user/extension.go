package user

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// TimeLayout 转换字符
const TimeLayout = "2006-01-02 15:04:05"

var (
	local = time.FixedZone("CST", 8*3600)
)

// BeforeCreate 插入前数据处理
func (user *User) BeforeCreate(scope *gorm.Scope) (err error) {
	uuid := uuid.NewV4()
	err = scope.SetColumn("Id", uuid.String())
	if err != nil {
		return err
	}
	err = scope.SetColumn("CreatedAt", time.Now().In(local).Format(TimeLayout))
	if err != nil {
		return err
	}
	err = scope.SetColumn("UpdatedAt", time.Now().In(local).Format(TimeLayout))
	if err != nil {
		return err
	}
	return nil
}
