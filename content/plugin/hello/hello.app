/*
pluginName = "Hello World Plugin"
pluginURI = "http://www.yougam.com"
description = "这是Zenpress系统的插件示例"
version = "1.0.1"
author = "Insion Ng"
authorURI = "http://www.at3.net"
license = "bsd"
*/

/*
index.php：控制板
edit.php：文章
upload.php：媒体
link-manager.php：链接
edit.php?post_type=page：页面
edit-comments.php：评论
theme.php：主题
plugin.php：插件
users.php：用户
tools.php：工具
options-general.php：设置
*/

helloAction = fn() {
	str = "<Hello Plugin>"
	println(str)
	return str
}

hook.AddActionHook("plugin", helloAction)
