package models

import (
	"errors"
	"fmt"
	"github.com/lunny/xorm"
	_ "github.com/mattn/go-sqlite3"
	"toropress/helper"
	//_ "github.com/lib/pq" //当某时间字段表现为0001-01-01 07:36:42+07:36:42形式的时候 会读不出数据
	//_ "github.com/bylevel/pq"
	"os"
	"time"
)

var (
	Engine *xorm.Engine
)

const (
	DbName         = "./data/sqlite.db"
	DbUser         = "root"
	mysqlDriver    = "mymysql"
	mysqlDrvformat = "%v/%v/"
	pgDriver       = "postgres"
	pgDrvFormat    = "user=%v dbname=%v sslmode=disable"
	sqlite3Driver  = "sqlite3"
	dbtype         = "sqlite"
)

type User struct {
	Id            int64
	Email         string `xorm:"index"`
	Password      string
	Nickname      string `xorm:"index"`
	Realname      string
	Avatar        string
	Avatar_min    string
	Avatar_max    string
	Birth         time.Time
	Province      string
	City          string
	Company       string
	Address       string
	Postcode      string
	Mobile        string
	Website       string
	Sex           int64
	Qq            string
	Msn           string
	Weibo         string
	Ctype         int64
	Role          int64
	Created       time.Time `xorm:"index"`
	Hotness       float64   `xorm:"index"`
	Hotup         int64     `xorm:"index"`
	Hotdown       int64     `xorm:"index"`
	Views         int64     `xorm:"index"`
	LastLoginTime time.Time
	LastLoginIp   string
	LoginCount    int64
}

//category,Pid:root
type Category struct {
	Id             int64
	Pid            int64 `xorm:"index"`
	Uid            int64 `xorm:"index"`
	Ctype          int64
	Title          string
	Content        string
	Attachment     string
	Created        time.Time `xorm:"index"`
	Hotness        float64   `xorm:"index"`
	Hotup          int64     `xorm:"index"`
	Hotdown        int64     `xorm:"index"`
	Views          int64     `xorm:"index"`
	Author         string
	NodeTime       time.Time
	NodeCount      int64
	NodeLastUserId int64
}

//node,Pid:category
type Node struct {
	Id              int64
	Pid             int64 `xorm:"index"`
	Uid             int64 `xorm:"index"`
	Order           int64
	Ctype           int64
	Title           string
	Content         string    `xorm:"text"`
	Attachment      string    `xorm:"text"`
	Created         time.Time `xorm:"index"`
	Updated         time.Time `xorm:"index"`
	Hotness         float64   `xorm:"index"`
	Hotup           int64     `xorm:"index"`
	Hotdown         int64     `xorm:"index"`
	Hotscore        int64     `xorm:"index"`
	Views           int64     `xorm:"index"`
	Author          string    //节点的创建者
	TopicTime       time.Time
	TopicCount      int64
	TopicLastUserId int64
}

//由于cid nid uid都可以是topic的上级所以默认不设置pid字段,这里默认Pid是nid
type Topic struct {
	Id                int64
	Cid               int64 `xorm:"index"`
	Nid               int64 `xorm:"index"`
	Uid               int64 `xorm:"index"`
	Order             int64
	Ctype             int64
	Title             string
	Content           string    `xorm:"text"`
	Attachment        string    `xorm:"text"`
	Created           time.Time `xorm:"index"`
	Updated           time.Time `xorm:"index"`
	Hotness           float64   `xorm:"index"`
	Hotup             int64     `xorm:"index"`
	Hotdown           int64     `xorm:"index"`
	Hotscore          int64     `xorm:"index"`
	Views             int64     `xorm:"index"`
	Author            string
	Category          string
	Node              string
	ReplyTime         time.Time
	ReplyCount        int64
	ReplyLastUserId   int64
	ReplyLastUsername string
	ReplyLastNickname string
}

//reply,Pid:topic
type Reply struct {
	Id         int64
	Uid        int64 `xorm:"index"`
	Pid        int64 `xorm:"index"` //Topic id
	Ctype      int64
	Content    string
	Attachment string
	Created    time.Time `xorm:"index"`
	Hotness    float64   `xorm:"index"`
	Hotup      int64     `xorm:"index"`
	Hotdown    int64     `xorm:"index"`
	Views      int64     `xorm:"index"`
	Author     string
	Email      string
	Website    string
}

type File struct {
	Id              int64
	Cid             int64 `xorm:"index"`
	Nid             int64 `xorm:"index"`
	Uid             int64 `xorm:"index"`
	Pid             int64 `xorm:"index"`
	Ctype           int64
	Filename        string
	Content         string
	Hash            string
	Location        string
	Url             string
	Size            int64
	Created         time.Time `xorm:"index"`
	Updated         time.Time `xorm:"index"`
	Hotness         float64   `xorm:"index"`
	Hotup           int64     `xorm:"index"`
	Hotdown         int64     `xorm:"index"`
	Views           int64     `xorm:"index"`
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

// k/v infomation
type Kvs struct {
	Id int64
	K  string
	V  string
}

func init() {
	_, err := SetEngine()
	if err != nil {
		fmt.Println(err)
	}
}

func XConDb() (*xorm.Engine, error) {
	switch {
	case dbtype == "sqlite":
		return xorm.NewEngine("sqlite3", DbName)

	case dbtype == "mysql":
		return xorm.NewEngine("mysql", "user=mysql password=jn!@#9^&* dbname=mysql")

	case dbtype == "pgsql":
		return xorm.NewEngine("postgres", "user=postgres password=jn!@#$%^&* dbname=pgsql sslmode=disable")
	}
	return nil, errors.New("尚未设定数据库连接")
}

func SetEngine() (*xorm.Engine, error) {

	var err error
	Engine, err = XConDb()
	//Engine.Mapper = xorm.SameMapper{}
	//Engine.SetMaxConns(5)
	//Engine.ShowSQL = true

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	Engine.SetDefaultCacher(cacher)

	return Engine, err
}

func CreateDb() {

	if err := Engine.Sync(new(User), new(Category), new(Node), new(Topic), new(Reply), new(Kvs), new(File)); err != nil {
		fmt.Println("Database struct sync failed")
		fmt.Println("Engine.Sync ERRORS:", err)
	} else {
		fmt.Println("Database struct sync successfully")
	}

	//用户等级划分：正数是普通用户，负数是管理员各种等级划分，为0则尚未注册
	if GetUserByRole(-1000).Role != -1000 {
		AddUser("root@localhost", "root", "系统默认管理员", helper.Encrypt_password("rootpass", nil), -1000)
		fmt.Println("Default User:root,Password:rootpass")

		if GetAllTopic(0, 0, "id") == nil {
			//分類默認數據
			AddCategory("Category！", "This is Category！")

			AddNode("Node！", "This is Node!", 1, 1)
			SetTopic(0, 1, 1, 1, 0, "Topic Title", `<p>This is Topic!</p>`, "root", "")

		}
	}

	if GetKV("author") != "Insion" {
		SetKV("author", "Insion")
		SetKV("title", "Toropress")
		SetKV("title_en", "Toropress")
		SetKV("keywords", "Toropress,")
		SetKV("description", "Toropress,")

		SetKV("company", "Toropress")
		SetKV("copyright", "2013 Copyright Toropress .All Right Reserved")
		SetKV("site_email", "info@verywave.com")

		SetKV("tweibo", "http://t.qq.com/yours")
		SetKV("sweibo", "http://weibo.com/yours")
	}

}

func Counts() (categorys int, nodes int, topics int, menbers int) {
	var categoryz []*Category
	if e := Engine.Find(&categoryz); e != nil {
		categorys = 0
		fmt.Println(e)
	} else {
		categorys = len(categoryz)
	}

	var nodez []*Node
	if e := Engine.Find(&nodez); e != nil {
		nodes = 0
		fmt.Println(e)
	} else {
		nodes = len(nodez)
	}

	var topicz []*Topic
	if e := Engine.Find(&topicz); e != nil {
		topics = 0
		fmt.Println(e)
	} else {
		topics = len(topicz)
	}

	var menberz []*User
	if e := Engine.Find(&menberz); e != nil {
		menbers = 0
		fmt.Println(e)
	} else {
		menbers = len(menberz)
	}

	return categorys, nodes, topics, menbers
}

func TopicCount() (today int, this_week int, this_month int) {
	var topict, topicw, topicm []*Topic
	k := time.Now()

	//一天之前
	d, _ := time.ParseDuration("-24h")
	t := k.Add(d)
	e := Engine.Where("created>?", t).Find(&topict)
	if e != nil {
		today = 0
		fmt.Println(e)
	} else {
		today = len(topict)
	}

	//一周之前
	w := k.Add(d * 7)
	e = Engine.Where("created>?", w).Find(&topicw)
	if e != nil {
		this_week = 0
		fmt.Println(e)
	} else {
		this_week = len(topicw)
	}

	//一月之前
	m := k.Add(d * 30)
	e = Engine.Where("created>?", m).Find(&topicm)
	if e != nil {
		this_month = 0
		fmt.Println(e)
	} else {
		this_month = len(topicm)
	}

	return today, this_week, this_month
}

func PutTopic(tid int64, tp *Topic) (int64, error) {
	//覆盖式更新话题
	tp.Updated = time.Now()
	row, err := Engine.Update(tp, &Topic{Id: tid}) //该方法目前返回的row为执行SQL所影响的行数

	return row, err

}

func PutNode(nid int64, nd *Node) (int64, error) {
	//覆盖式更新节点
	nd.Updated = time.Now()
	row, err := Engine.Update(nd, &Node{Id: nid})

	return row, err

}

func SetTopic(id int64, cid int64, nid int64, uid int64, ctype int64, title string, content string, author string, attachment string) error {
	var tp Topic
	if has, err := Engine.Id(id).Get(&tp); !has {
		_, err = Engine.Insert(&Topic{Id: id, Cid: cid, Nid: nid, Uid: uid, Ctype: ctype, Title: title, Content: content, Author: author, Attachment: attachment})
		return err
	} else {
		_, err := Engine.Id(id).Update(&Topic{Cid: cid, Nid: nid, Uid: uid, Ctype: ctype, Title: title, Content: content, Author: author, Attachment: attachment})
		return err
	}
	return nil
}

func AddFile(ctype int64, location string, url string) error {
	_, err := Engine.Insert(&File{Ctype: ctype, Location: location, Url: url})
	return err
}

func DelFile(id int64) error {
	f := GetFile(id)

	if helper.Exist("." + f.Location) {
		if err := os.Remove("." + f.Location); err != nil {
			return err
			fmt.Println(err)
		}
	}

	//不管实际路径中是否存在文件均删除该数据库记录，以免数据库记录陷入死循环无法删掉
	_, err := Engine.Id(id).Delete(new(File))
	fmt.Println(err)
	return err
}

func GetFile(id int64) *File {
	f := new(File)
	Engine.Id(id).Get(f)
	return f
}

func GetAllFile() []*File {
	f := make([]*File, 0)
	Engine.Desc("id").Find(&f)
	return f
}

func GetAllFileByCtype(ctype int64) []*File {
	f := make([]*File, 0)
	Engine.Where("ctype=?", ctype).Desc("id").Find(&f)
	return f
}

func SaveFile(f *File) error {
	_, e := Engine.Insert(&f)
	return e
}

func SaveUser(usr *User) error {
	_, e := Engine.Insert(&usr)
	return e
}

func SaveCategory(cat *Category) error {
	_, err := Engine.Insert(&cat)
	return err
}

func SaveNode(nd *Node) error {

	_, err := Engine.Insert(nd)
	return err
}

func SaveTopic(tp *Topic) error {
	_, err := Engine.Insert(&tp)
	return err
}

func SetFile(id int64, pid int64, ctype int64, filename string, content string, hash string, location string, url string, size int64) error {
	var f File
	if has, err := Engine.Id(id).Get(&f); !has {
		_, err = Engine.Insert(&File{Id: id, Pid: pid, Ctype: ctype, Filename: filename, Content: content, Hash: hash, Location: location, Url: url, Size: size})
		return err
	} else {
		_, err := Engine.Id(id).Update(&File{Pid: pid, Ctype: ctype, Filename: filename, Content: content, Hash: hash, Location: location, Url: url, Size: size})

		return err
	}
	return nil
}

func AddKV(k string, v string) error {
	_, err := Engine.Insert(&Kvs{K: k, V: v})
	return err
}

func SetKV(k string, v string) error {
	var kvs Kvs
	if has, err := Engine.Where("k=?", k).Get(&kvs); !has {
		_, err = Engine.Insert(&Kvs{K: k, V: v})
		return err
	} else {
		type Kvs struct {
			K string
			V string
		}

		_, err := Engine.Where("k=?", k).Update(&Kvs{K: k, V: v})

		return err
	}
	return nil
}

func GetKV(k string) (v string) {
	var kvs Kvs
	Engine.Where("k=?", k).Get(&kvs)
	return kvs.V
}

func AddUser(email string, nickname string, realname string, password string, role int64) error {
	_, err := Engine.Insert(&User{Email: email, Nickname: nickname, Realname: realname, Password: password, Role: role, Created: time.Now()})

	return err
}

func UpdateUser(uid int64, ur User) error {
	_, err := Engine.Id(uid).Update(&ur)
	return err
}

func GetUser(id int64) *User {
	usr := new(User)
	Engine.Id(id).Get(usr)
	return usr
}

func DelUser(uid int64) error {
	usr := GetUser(uid)
	_, err := Engine.Delete(&usr)

	return err
}

func GetUserByRole(role int) *User {
	usr := new(User)
	Engine.Where("role=?", role).Get(usr)
	return usr
}

func GetAllUserByRole(role int) *[]User {
	users := new([]User)
	Engine.Where("role=?", role).Desc("id").Find(users)
	return users
}

func GetUserByNickname(nickname string) *User {
	user := new(User)
	Engine.Where("nickname=?", nickname).Get(user)
	return user
}

func AddCategory(title string, content string) error {
	_, err := Engine.Insert(&Category{Title: title, Content: content, Created: time.Now()})

	return err
}

func AddNode(title string, content string, cid int64, uid int64) error {
	if _, err := Engine.Insert(&Node{Pid: cid, Uid: uid, Title: title, Content: content, Created: time.Now()}); err != nil {
		return err
	}

	if _, err := Engine.Where("id=?", cid).Update(&Category{NodeTime: time.Now(), NodeCount: int64(len(GetAllNodeByCid(cid, 0, 0, 0, "id"))), NodeLastUserId: uid}); err != nil {
		return err
	}
	return nil
}

func SetNode(id int64, title string, content string, cid int64, uid int64) error {
	var nd Node
	if has, err := Engine.Id(id).Get(&nd); !has {
		_, err = Engine.Insert(&Node{Id: id, Pid: cid, Uid: uid, Title: title, Content: content})
		return err
	} else {
		_, err := Engine.Id(id).Update(&Node{Pid: cid, Uid: uid, Title: title, Content: content})
		return err
	}
	return nil
}

func AddTopic(title string, content string, cid int64, nid int64, uid int64) error {
	if _, err := Engine.Insert(&Topic{Cid: cid, Nid: nid, Title: title, Content: content, Created: time.Now()}); err != nil {
		return err
	}
	if _, err := Engine.Id(nid).Update(&Node{TopicTime: time.Now(), TopicCount: int64(len(GetAllTopicByNid(nid, 0, 0, 0, "id"))), TopicLastUserId: uid}); err != nil {
		return err
	}
	return nil
}

func AddReply(tid int64, uid int64, content string, author string, email string, website string) error {
	if _, err := Engine.Insert(&Reply{Pid: tid, Uid: uid, Content: content, Created: time.Now(), Author: author, Email: email, Website: website}); err != nil {
		return err
	}

	if _, err := Engine.Id(tid).Update(&Topic{ReplyTime: time.Now(), ReplyCount: int64(len(GetReplyByPid(tid, 0, 0, "id"))), ReplyLastUserId: uid}); err != nil {
		return err
	}
	return nil
}

func DelNodePlus(nid int64) error {
	node := GetNode(nid)
	_, err := Engine.Delete(&node)

	for i, v := range GetAllTopicByNid(nid, 0, 0, 0, "id") {
		if i > 0 {
			DelTopic(v.Id)
			for ii, vv := range GetReplyByPid(v.Id, 0, 0, "id") {
				if ii > 0 {
					DelReply(vv.Id)
				}
			}
		}
	}

	return err
}

func DelCategory(id int64) error {
	_, err := Engine.Id(id).Delete(&Category{})

	return err
}

func DelTopic(id int64) error {
	topic := GetTopic(id)
	if helper.Exist("." + topic.Attachment) {
		if err := os.Remove("." + topic.Attachment); err != nil {
			//return err
			//可以输出错误，但不要反回错误，以免陷入死循环无法删掉
			fmt.Println("DEL TOPIC", id, err)
		}
	}

	//不管实际路径中是否存在文件均删除该数据库记录，以免数据库记录陷入死循环无法删掉
	_, err := Engine.Id(id).Delete(new(Topic))

	return err
}

func DelNode(nid int64) error {
	_, err := Engine.Id(nid).Delete(new(Node))

	return err
}

func DelReply(tid int64) error {
	_, err := Engine.Id(tid).Delete(new(Reply))

	return err
}

func GetAllCategory() []Category {
	cats := make([]Category, 0)
	Engine.Find(&cats)
	return cats
}

func GetAllNode() []Node {
	nds := make([]Node, 0)
	Engine.Desc("created").Find(&nds)
	return nds
}

func GetAllTopic(offset int, limit int, path string) []Topic {
	tps := make([]Topic, 0)
	Engine.Limit(limit, offset).Desc(path, "created").Find(&tps)
	return tps
}

func GetAllNodeByCid(cid int64, offset int, limit int, ctype int64, path string) []*Node {
	//排序首先是热值优先，然后是时间优先。
	alln := make([]*Node, 0)
	switch {
	case path == "asc":
		if ctype != 0 {
			Engine.Where("pid=? and ctype=?", cid, ctype).Limit(limit, offset).Find(&alln)
		} else {
			if cid == 0 {
				Engine.Limit(limit, offset).Find(&alln)
			} else {
				Engine.Where("pid=?", cid).Limit(limit, offset).Find(&alln)
			}
		}
	case path == "views" || path == "topic_count":
		if ctype != 0 {
			Engine.Where("pid=? and ctype=?", cid, ctype).Desc(path).Limit(limit, offset).Find(&alln)
		} else {
			if cid == 0 {
				Engine.Desc(path).Limit(limit, offset).Find(&alln)
			} else {
				Engine.Where("pid=?", cid).Desc(path).Limit(limit, offset).Find(&alln)
			}
		}
	default:
		if ctype != 0 {
			Engine.Where("pid=? and ctype=?", cid, ctype).Limit(limit, offset).Desc(path, "views, topic_count, created").Find(&alln)

		} else {
			if cid == 0 {
				Engine.Limit(limit, offset).Desc(path, "views,topic_count,created").Find(&alln)
			} else {
				Engine.Where("pid=?", cid).Limit(limit, offset).Desc(path, "views,topic_count,created").Find(&alln)
			}
		}

	}
	return alln
}

func GetAllTopicByCid(cid int64, offset int, limit int, ctype int64, path string) []*Topic {
	//排序首先是热值优先，然后是时间优先。
	allt := make([]*Topic, 0)
	switch {
	case path == "asc":
		if ctype != 0 {
			Engine.Where("cid=? and ctype=?", cid, ctype).Limit(limit, offset).Find(&allt)
		} else {
			Engine.Where("cid=?", cid).Limit(limit, offset).Find(&allt)
		}
	case path == "views" || path == "reply_count":
		if ctype != 0 {
			Engine.Where("cid=? and ctype=?", cid, ctype).Desc(path).Limit(limit, offset).Find(&allt)

		} else {
			if cid == 0 {
				Engine.Desc(path).Limit(limit, offset).Find(&allt)
			} else {
				Engine.Where("cid=?", cid).Desc(path).Limit(limit, offset).Find(&allt)
			}
		}
	default:
		if ctype != 0 {
			Engine.Where("cid=? and ctype=?", cid, ctype).Limit(limit, offset).Desc(path, "views, reply_count, created").Find(&allt)
		} else {
			if cid == 0 {
				Engine.Limit(limit, offset).Desc(path, "views, reply_count, created").Find(&allt)
			} else {
				Engine.Where("cid=?", cid).Limit(limit, offset).Desc(path, "views, reply_count, created").Find(&allt)
			}
		}
	}
	return allt
}

func GetAllTopicByCidNid(cid int64, nid int64, offset int, limit int, ctype int64, path string) []*Topic {
	allt := make([]*Topic, 0)
	switch {
	case path == "asc":
		if ctype != 0 {
			Engine.Where("cid=? and nid=? and ctype=?", cid, nid, ctype).Limit(limit, offset).Find(&allt)
		} else {
			Engine.Where("cid=? and nid=?", cid, nid).Limit(limit, offset).Find(&allt)
		}
	default:
		if ctype != 0 {
			Engine.Where("cid=? and nid=? and ctype=?", cid, nid, ctype).Limit(limit, offset).Desc(path, "views,reply_count,created").Find(&allt)
		} else {
			Engine.Where("cid=? and nid=?", cid, nid).Limit(limit, offset).Desc("views,reply_count,created").Find(&allt)
		}
	}
	return allt
}

func GetAllTopicByNid(nodeid int64, offset int, limit int, ctype int64, path string) []*Topic {
	//排序首先是热值优先，然后是时间优先。
	allt := make([]*Topic, 0)
	switch {
	case path == "asc":
		if nodeid == 0 {
			//q.Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("reply_count").OrderByDesc("created").FindAll(&allt)
			return nil
		} else {
			if ctype != 0 {
				Engine.Where("nid=? and ctype=?", nodeid, ctype).Limit(limit, offset).Find(&allt)
			} else {
				Engine.Where("nid=?", nodeid).Limit(limit, offset).Find(&allt)
			}
		}
	default:
		if nodeid == 0 {
			return nil
		} else {
			if ctype != 0 {
				Engine.Where("nid=? and ctype=?", nodeid, ctype).Limit(limit, offset).Desc(path, "views,reply_count,created").Find(&allt)
			} else {
				Engine.Where("nid=?", nodeid).Limit(limit, offset).Desc(path, "views,reply_count,created").Find(&allt)
			}
		}
	}
	return allt
}

func SearchTopic(content string, offset int, limit int, path string) *[]Topic {
	tps := new([]Topic)
	if content != "" {
		keyword := "%" + content + "%"

		Engine.Where("title like ? or content like ?", keyword, keyword).Limit(limit, offset).Desc(path, "views", "reply_count", "created").Find(tps)

		return tps
	}
	return nil
}

func GetCategory(id int64) (category Category) {
	Engine.Where("id=?", id).Get(&category)
	return category
}

func GetNode(id int64) *Node {
	nd := new(Node)
	Engine.Where("id=?", id).Get(nd)
	return nd
}

func GetTopic(id int64) *Topic {
	tp := new(Topic)
	Engine.Where("id=?", id).Get(tp)
	return tp
}

func UpdateCategory(cid int64, cg Category) error {
	_, err := Engine.Id(cid).Update(&cg)
	return err
}

func UpdateNode(nid int64, nd *Node) error {
	_, err := Engine.Where("id=?", int64(nid)).Update(nd)
	return err
}

func UpdateTopic(tid int64, tp *Topic) error {
	_, err := Engine.Where("id=?", int64(tid)).Update(tp)
	return err
}

func EditNode(nid int64, cid int64, uid int64, title string, content string) error {
	nd := GetNode(nid)
	nd.Pid = cid
	nd.Title = title
	nd.Content = content
	nd.Updated = time.Now()
	if err := UpdateNode(nid, nd); err != nil {
		return err
	}

	if _, err := Engine.Id(cid).Update(&Category{NodeTime: time.Now(), NodeCount: int64(len(GetAllNodeByCid(cid, 0, 0, 0, "id"))), NodeLastUserId: int64(uid)}); err != nil {
		return err
	}

	return nil
}

func EditTopic(tid int64, nid int64, cid int64, uid int64, title string, content string) error {
	tpc := GetTopic(tid)
	tpc.Cid = int64(cid)
	tpc.Nid = int64(nid)
	tpc.Title = title
	tpc.Content = content
	tpc.Updated = time.Now()

	if err := UpdateTopic(tid, tpc); err != nil {
		return err
	}

	if _, err := Engine.Where("id=?", nid).Update(&Node{TopicTime: tpc.Created, TopicCount: int64(len(GetAllTopicByNid(nid, 0, 0, 0, "id"))), TopicLastUserId: int64(uid)}); err != nil {
		return err
	}

	return nil
}

func GetAllReply() (allr []*Reply) {
	allr = make([]*Reply, 0)
	Engine.Desc("id").Find(&allr)
	return allr
}

func GetReply(id int64) (reply Reply) {
	Engine.Where("id=?", id).Get(&reply)
	return reply
}

func GetReplyByPid(tid int64, offset int, limit int, path string) []*Reply {
	allr := make([]*Reply, 0)
	if tid == 0 {
		Engine.Limit(limit, offset).Desc(path).Find(&allr)
	} else {
		//最热回复
		//q.Where("pid=?", tid).Offset(offset).Limit(limit).OrderByDesc("hotness").FindAll(&allr)
		Engine.Where("pid=?", tid).Limit(limit, offset).Desc(path).Find(&allr)
	}
	return allr
}
