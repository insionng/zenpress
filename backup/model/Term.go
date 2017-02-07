package model

// terms、term_relationships、term_taxonomy
// 系统里，我们常见的分类有文章的分类、链接的分类，
// 实际上还有TAG，它也是一种特殊的分类方式，
// 我们甚至还可以创建自己的分类方法。
// 将所有的分类及分类方法、对应结构都记录在这三个表中。
// terms记录了每个分类的名字以及基本信息，
// 如本站分为“系统开发”、“CEO插件”等，
// 这里的分类指广义上的分类，所以每个TAG也是一个“分类”。
// term_taxonomy记录了每个分类所归属的分类方法，
// 如“系统开发”、“CEO插件”是文章分类（category），
// 放置友情链接的“我的朋友”、“我的同事”分类属于友情链接分类（link_category）
// term_relationships记录了每个文章（或链接）所对应的分类方法。
// 庆幸的是，关于term的使用，系统中相关函数的使用方法还是比较清晰明了，
// 我们就没必要纠结于它的构造了

// Term 存储菜单分类、标签分类名称及URL信息；
type Term struct {
	TermId    int64  `xorm:"not null pk autoincr BIGINT(20)"`        //分类ID
	Name      string `xorm:"not null default '' index VARCHAR(200)"` //分类名
	Slug      string `xorm:"not null default '' index VARCHAR(200)"` //缩略名
	TermGroup int64  `xorm:"not null default 0 BIGINT(10)"`          //未知
}
