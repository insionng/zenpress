package handler

import (
	"github.com/insionng/macross"
)

// SignoutHandler 用户离开时候执行相关记录动作
// 客户端需要执行删除TOKEN操作
func SignoutHandler(self *macross.Context) error {
	return self.Redirect("/", 302)
}
