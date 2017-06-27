package model

import (
	"github.com/jinzhu/gorm"
)

// Option 基本配置信息表，通常通过get_option来操作，该表通常作为插件存储数据的一个地方。
type Option struct {
	OptionID    uint64 `gorm:"primary_key"`
	OptionName  string
	OptionValue string //`gorm:"not null LONGTEXT"`
	Autoload    string
}

/*
这是一种向选项数据库表中添加有名称的选项/值对的安全方法。
如果所需选项已存在，add_option()不添加内容。
选项被保存后，可通过get_option()来访问选项，
通过update_option()来修改选项，还可以通过delete_option()删除该选项。
*/

// AddOption 新增选项
func AddOption(key, value string, autoload ...string) (db *gorm.DB) {
	var iAutoload = "yes"
	if len(autoload) > 0 {
		if !((autoload[0] != "yes") && (autoload[0] != "no")) {
			iAutoload = autoload[0]
		}
	}
	db = Database.Create(&Option{OptionName: key, OptionValue: value, Autoload: iAutoload})
	return
}

// GetOption 获得选项
func GetOption(key string) (db *gorm.DB, option Option) {
	db = Database.First(&option, "option_name = ?", key)
	return
}

// UpdateOption 更新选项
func UpdateOption(key, value string) (db *gorm.DB, option Option) {
	option = Option{OptionValue: value}
	db = Database.Model(&option).Update("option_name", key)
	return
}

// DeleteOption 删除选项
func DeleteOption(key string) (db *gorm.DB) {
	db = Database.Delete(Option{}, "option_name = ?", key)
	return
}
