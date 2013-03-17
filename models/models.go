package models

import (
	"../utils"
	"database/sql"
	"fmt"
	"github.com/coocood/qbs"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

const (
	dbName         = "./data/sqlite.db"
	dbUser         = "root"
	mysqlDriver    = "mymysql"
	mysqlDrvformat = "%v/%v/"
	pgDriver       = "postgres"
	pgDrvFormat    = "user=%v dbname=%v sslmode=disable"
	sqlite3Driver  = "sqlite3"
)

type User struct {
	Id              int64
	Email           string
	Password        string
	Nickname        string
	Realname        string
	Avatar          string
	Avatar_min      string
	Avatar_max      string
	Birth           time.Time
	Province        string
	City            string
	Address         string
	Postcode        string
	Mobile          string
	Website         string
	Sex             string
	Qq              string
	Msn             string
	Weibo           string
	Ctype           int64
	Role            int64
	Created         time.Time
	Hotness         float64
	Hotup           int64
	Hotdown         int64
	Views           int64
	Last_login_time time.Time
	Last_login_ip   string
	Login_count     int64
}

//category,Pid:root
type Category struct {
	Id         int64
	Pid        int64
	Uid        int64
	Ctype      int64
	Title      string
	Content    string
	Attachment string
	Created    time.Time
	Hotness    float64
	Hotup      int64
	Hotdown    int64
	Views      int64
}

//node,Pid:category
type Node struct {
	Id         int64
	Pid        int64
	Uid        int64
	Ctype      int64
	Title      string
	Content    string
	Attachment string
	Created    time.Time
	Hotness    float64
	Hotup      int64
	Hotdown    int64
	Views      int64
}

//topic,Pid:node
type Topic struct {
	Id                 int64
	Cid                int64
	Nid                int64
	Uid                int64
	Ctype              int64
	Title              string
	Content            string
	Attachment         string
	Created            time.Time
	Hotness            float64
	Hotup              int64
	Hotdown            int64
	Views              int64
	Reply_time         time.Time
	Reply_count        int64
	Reply_last_user_id int64
}

//reply,Pid:topic
type Reply struct {
	Id         int64
	Uid        int64
	Pid        int64 //Topic id
	Ctype      int64
	Content    string
	Attachment string
	Created    time.Time
	Hotness    float64
	Hotup      int64
	Hotdown    int64
	Views      int64
	Author     string
	Email      string
	Website    string
}

func ConnDb() (*qbs.Qbs, error) {
	db, err := sql.Open(sqlite3Driver, dbName)
	q := qbs.New(db, qbs.NewSqlite3())
	return q, err
}

func SetMg() (*qbs.Migration, error) {
	db, err := sql.Open(sqlite3Driver, dbName)
	mg := qbs.NewMigration(db, dbName, qbs.NewSqlite3())
	return mg, err
}

func Ct() {
	q, err := ConnDb()
	defer q.Db.Close()
	if err != nil {
		fmt.Println(err)
	}

	mg, _ := SetMg()
	defer mg.Db.Close()

	mg.CreateTableIfNotExists(new(User))
	mg.CreateTableIfNotExists(new(Category))
	mg.CreateTableIfNotExists(new(Node))
	mg.CreateTableIfNotExists(new(Topic))
	mg.CreateTableIfNotExists(new(Reply))

	if GetUserByNickname("root").Nickname == "" {
		AddUser("root@insion.co", "root", utils.Encrypt_password("rootpass", nil), 100)
		fmt.Println("Default User:root,Password:rootpass")
	}

	if GetTopic(1).Title == "" {
		AddCategory("Hello Category", "This is Category!")
		AddNode("Hello Node!", "This is Node!", 1)
		AddTopic("Hello World!", "This is Toropress!", 1, 1)
	}
}

func AddUser(email string, nickname string, password string, role int) error {
	q, _ := ConnDb()
	defer q.Db.Close()
	_, err := q.Save(&User{Email: email, Nickname: nickname, Password: password, Role: int64(role), Created: time.Now()})

	return err
}

func SaveUser(usr User) User {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.Save(&usr)
	return usr
}

func GetUser(id int) (user User) {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.Where("id=?", id).Find(&user)
	return user
}

func GetUserByNickname(nickname string) (user User) {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.Where("nickname=?", nickname).Find(&user)
	return user
}

func AddCategory(title string, content string) error {
	q, _ := ConnDb()
	defer q.Db.Close()
	_, err := q.Save(&Category{Title: title, Content: content, Created: time.Now()})

	return err
}

func SaveCategory(cat Category) Category {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.Save(&cat)
	return cat
}

func AddNode(title string, content string, categoryid int) error {
	q, _ := ConnDb()
	defer q.Db.Close()
	_, err := q.Save(&Node{Pid: int64(categoryid), Title: title, Content: content, Created: time.Now()})

	return err
}

func SaveNode(nd Node) Node {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.Save(&nd)
	return nd
}

func DelNode(nid int) error {
	q, _ := ConnDb()
	defer q.Db.Close()
	node := GetNode(nid)
	_, err := q.Delete(&node)

	for i, v := range GetAllTopicByNodeid(nid, 0, 0, "id") {
		if i > 0 {
			DelTopic(int(v.Id))
			for ii, vv := range GetReplyByPid(int(v.Id), 0, 0, "id") {
				if ii > 0 {
					DelReply(int(vv.Id))
				}
			}
		}
	}

	return err
}

func DelReply(tid int) error {
	q, _ := ConnDb()
	defer q.Db.Close()
	reply := GetReply(tid)
	_, err := q.Delete(&reply)

	return err
}

func AddTopic(title string, content string, cid int, nodeid int) error {
	q, _ := ConnDb()
	defer q.Db.Close()
	_, err := q.Save(&Topic{Cid: int64(cid), Nid: int64(nodeid), Title: title, Content: content, Created: time.Now()})

	return err
}

func EditTopic(tp Topic) Topic {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.Save(&tp)
	return tp
}

func DelTopic(id int) error {
	q, _ := ConnDb()
	defer q.Db.Close()
	topic := GetTopic(id)
	_, err := q.Delete(&topic)

	return err
}

func AddReply(pid int, uid int, content string, author string, email string, website string) error {
	q, _ := ConnDb()
	defer q.Db.Close()
	_, err := q.Save(&Reply{Pid: int64(pid), Uid: int64(uid), Content: content, Created: time.Now(), Author: author, Email: email, Website: website})

	return err
}

func GetAllCategory() (allc []*Category) {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.FindAll(&allc)
	return allc
}

func GetCategory(id int) (category Category) {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.Where("id=?", id).Find(&category)
	return category
}

func GetAllNode() (alln []*Node) {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.FindAll(&alln)
	return alln
}

func GetAllNodeByCategoryId(id int, offset int, limit int, path string) (alln []*Node) {
	q, _ := ConnDb()
	defer q.Db.Close()
	if id == 0 {
		q.Offset(offset).Limit(limit).OrderByDesc(path).FindAll(&alln)
	} else {
		//最热节点
		//q.Where("pid=?", id).Offset(offset).Limit(limit).OrderByDesc("hotness").FindAll(&alln)
		q.WhereEqual("pid", id).Offset(offset).Limit(limit).OrderByDesc(path).FindAll(&alln)
	}
	return alln
}

func GetNode(id int) (node Node) {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.Where("id=?", id).Find(&node)
	return node
}

func SearchAllTopicByTitle(title string, offset int, limit int, path string) (allt []*Topic) {
	q, _ := ConnDb()
	defer q.Db.Close()
	keyword := "%" + title + "%"
	q.Where("title like ?", keyword).Offset(offset).Limit(limit).OrderByDesc(path).FindAll(&allt)
	return allt
}

func SearchAllTopicByContent(content string, offset int, limit int, path string) (allt []*Topic) {
	q, _ := ConnDb()
	defer q.Db.Close()
	keyword := "%" + content + "%"
	q.Where("content like ?", keyword).Offset(offset).Limit(limit).OrderByDesc(path).FindAll(&allt)
	return allt
}

func GetAllTopic(offset int, limit int, path string) (allt []*Topic) {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.Offset(offset).Limit(limit).OrderByDesc(path).FindAll(&allt)
	return allt
}

func GetAllTopicByNodeid(nodeid int, offset int, limit int, path string) (allt []*Topic) {
	q, _ := ConnDb()
	defer q.Db.Close()
	if nodeid == 0 {
		q.Offset(offset).Limit(limit).OrderByDesc(path).FindAll(&allt)
	} else {
		q.WhereEqual("nid", nodeid).Offset(offset).Limit(limit).OrderByDesc(path).FindAll(&allt)
	}
	return allt
}

func GetTopic(id int) (topic Topic) {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.Where("id=?", id).Find(&topic)
	return topic
}

func SaveTopic(tp Topic) Topic {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.Save(&tp)
	return tp
}

func GetAllReply() (allr []*Reply) {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.FindAll(&allr)
	return allr
}

func GetReply(id int) (reply Reply) {
	q, _ := ConnDb()
	defer q.Db.Close()
	q.Where("id=?", id).Find(&reply)
	return reply
}

func GetReplyByPid(pid int, offset int, limit int, path string) (allr []*Reply) {
	q, _ := ConnDb()
	defer q.Db.Close()
	if pid == 0 {
		q.Offset(offset).Limit(limit).OrderByDesc(path).FindAll(&allr)
	} else {
		//最热回复
		//q.Where("pid=?", pid).Offset(offset).Limit(limit).OrderByDesc("hotness").FindAll(&allr)
		q.WhereEqual("pid", pid).Offset(offset).Limit(limit).OrderByDesc(path).FindAll(&allr)
	}
	return allr
}

/*
func main() {

	ct()

	for i := 0; i < 100; i++ {
		AddCategory("我系标题", "我系内容啊~")
	}
	for i := 0; i < 100; i++ {
		AddUser("insion@lihuaer.com", "insion", "huhjj897857hggfgjhghsjg")
	}
	for i := 0; i < 100; i++ {
		AddNode("node title", "node content")
	}
	for i := 0; i < 100; i++ {
		AddTopic("topic title", "topic content")
	}
	for i := 0; i < 100; i++ {
		AddReply(int64(i), "a reply's content")
	}
	cc := GetAllCategory()
	for _, info := range cc {
		fmt.Println(info.Content)
	}

	c := GetCategory(1)
	fmt.Println(c.Content)

	n := GetNode(1)
	fmt.Println(n.Title)

	t := GetTopic(1)
	fmt.Println(t.Content)

	r := GetReply(1)
	fmt.Println(r.Content)

	for _, info := range GetAllCategory() {
		fmt.Println(info.Title)
	}
	for _, info := range GetAllNode() {
		fmt.Println(info.Content)
	}
	for _, info := range GetAllTopic() {
		fmt.Println(info.Title)
	}
	for _, info := range GetAllReply() {
		fmt.Println(info.Content)
	}
}
*/
