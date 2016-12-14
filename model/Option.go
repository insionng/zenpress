package model

// Option 存储在系统默认及后台中用户设置的系统选项、插件及主题配置信息，
// 如用户设置的私人站名私人站描述是什么、用的什么主题、主题里用了什么功能、
// 是否开放注册、是否用了固定链接及其形式是什么等等；
type Option struct {
	OptionId    int64  `xorm:"not null pk autoincr BIGINT(20)"` //自增唯一ID
	BlogId      int64  //博客ID，用于多用户博客，默认0
	OptionName  string `xorm:"not null default '' unique VARCHAR(191)"` //键名
	OptionValue string `xorm:"not null LONGTEXT"`                       //键值
	Autoload    string `xorm:"not null default 'yes' VARCHAR(20)"`      //在系统载入时自动载入（yes/no）
}
