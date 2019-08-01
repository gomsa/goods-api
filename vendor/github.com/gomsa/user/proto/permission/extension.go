package permission

import (
	"time"

	"github.com/jinzhu/gorm"
)

// TimeLayout 转换字符
const TimeLayout = "2006-01-02 15:04:05"

var (
	local = time.FixedZone("CST", 8*3600)
)

// BeforeCreate 插入前数据处理
func (p *Permission) BeforeCreate(scope *gorm.Scope) (err error) {
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
