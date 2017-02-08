package model

import (
	"time"
)

type Comments struct {
	CommentId          int64     `xorm:"not null pk autoincr BIGINT(20)"`
	CommentPostId      int64     `xorm:"not null default 0 index BIGINT(20)"`
	CommentAuthor      string    `xorm:"not null TINYTEXT"`
	CommentAuthorEmail string    `xorm:"not null default '' index VARCHAR(100)"`
	CommentAuthorUrl   string    `xorm:"not null default '' VARCHAR(200)"`
	CommentAuthorIp    string    `xorm:"not null default '' VARCHAR(100)"`
	CommentDate        time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	CommentDateGmt     time.Time `xorm:"not null default '0000-00-00 00:00:00' index index(comment_approved_date_gmt) DATETIME"`
	CommentContent     string    `xorm:"not null TEXT"`
	CommentKarma       int       `xorm:"not null default 0 INT(11)"`
	CommentApproved    string    `xorm:"not null default '1' index(comment_approved_date_gmt) VARCHAR(20)"`
	CommentAgent       string    `xorm:"not null default '' VARCHAR(255)"`
	CommentType        string    `xorm:"not null default '' VARCHAR(20)"`
	CommentParent      int64     `xorm:"not null default 0 index BIGINT(20)"`
	UserId             int64     `xorm:"not null default 0 BIGINT(20)"`
}

// GetCommentsesCount Commentss' Count
func GetCommentsesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Comments{})
	return total, err
}

// GetCommentsCountViaCommentId Get Comments via CommentId
func GetCommentsCountViaCommentId(iCommentId int64) int64 {
	n, _ := Engine.Where("comment_ID = ?", iCommentId).Count(&Comments{CommentId: iCommentId})
	return n
}

// GetCommentsCountViaCommentPostId Get Comments via CommentPostId
func GetCommentsCountViaCommentPostId(iCommentPostId int64) int64 {
	n, _ := Engine.Where("comment_post_ID = ?", iCommentPostId).Count(&Comments{CommentPostId: iCommentPostId})
	return n
}

// GetCommentsCountViaCommentAuthor Get Comments via CommentAuthor
func GetCommentsCountViaCommentAuthor(iCommentAuthor string) int64 {
	n, _ := Engine.Where("comment_author = ?", iCommentAuthor).Count(&Comments{CommentAuthor: iCommentAuthor})
	return n
}

// GetCommentsCountViaCommentAuthorEmail Get Comments via CommentAuthorEmail
func GetCommentsCountViaCommentAuthorEmail(iCommentAuthorEmail string) int64 {
	n, _ := Engine.Where("comment_author_email = ?", iCommentAuthorEmail).Count(&Comments{CommentAuthorEmail: iCommentAuthorEmail})
	return n
}

// GetCommentsCountViaCommentAuthorUrl Get Comments via CommentAuthorUrl
func GetCommentsCountViaCommentAuthorUrl(iCommentAuthorUrl string) int64 {
	n, _ := Engine.Where("comment_author_url = ?", iCommentAuthorUrl).Count(&Comments{CommentAuthorUrl: iCommentAuthorUrl})
	return n
}

// GetCommentsCountViaCommentAuthorIp Get Comments via CommentAuthorIp
func GetCommentsCountViaCommentAuthorIp(iCommentAuthorIp string) int64 {
	n, _ := Engine.Where("comment_author_IP = ?", iCommentAuthorIp).Count(&Comments{CommentAuthorIp: iCommentAuthorIp})
	return n
}

// GetCommentsCountViaCommentDate Get Comments via CommentDate
func GetCommentsCountViaCommentDate(iCommentDate time.Time) int64 {
	n, _ := Engine.Where("comment_date = ?", iCommentDate).Count(&Comments{CommentDate: iCommentDate})
	return n
}

// GetCommentsCountViaCommentDateGmt Get Comments via CommentDateGmt
func GetCommentsCountViaCommentDateGmt(iCommentDateGmt time.Time) int64 {
	n, _ := Engine.Where("comment_date_gmt = ?", iCommentDateGmt).Count(&Comments{CommentDateGmt: iCommentDateGmt})
	return n
}

// GetCommentsCountViaCommentContent Get Comments via CommentContent
func GetCommentsCountViaCommentContent(iCommentContent string) int64 {
	n, _ := Engine.Where("comment_content = ?", iCommentContent).Count(&Comments{CommentContent: iCommentContent})
	return n
}

// GetCommentsCountViaCommentKarma Get Comments via CommentKarma
func GetCommentsCountViaCommentKarma(iCommentKarma int) int64 {
	n, _ := Engine.Where("comment_karma = ?", iCommentKarma).Count(&Comments{CommentKarma: iCommentKarma})
	return n
}

// GetCommentsCountViaCommentApproved Get Comments via CommentApproved
func GetCommentsCountViaCommentApproved(iCommentApproved string) int64 {
	n, _ := Engine.Where("comment_approved = ?", iCommentApproved).Count(&Comments{CommentApproved: iCommentApproved})
	return n
}

// GetCommentsCountViaCommentAgent Get Comments via CommentAgent
func GetCommentsCountViaCommentAgent(iCommentAgent string) int64 {
	n, _ := Engine.Where("comment_agent = ?", iCommentAgent).Count(&Comments{CommentAgent: iCommentAgent})
	return n
}

// GetCommentsCountViaCommentType Get Comments via CommentType
func GetCommentsCountViaCommentType(iCommentType string) int64 {
	n, _ := Engine.Where("comment_type = ?", iCommentType).Count(&Comments{CommentType: iCommentType})
	return n
}

// GetCommentsCountViaCommentParent Get Comments via CommentParent
func GetCommentsCountViaCommentParent(iCommentParent int64) int64 {
	n, _ := Engine.Where("comment_parent = ?", iCommentParent).Count(&Comments{CommentParent: iCommentParent})
	return n
}

// GetCommentsCountViaUserId Get Comments via UserId
func GetCommentsCountViaUserId(iUserId int64) int64 {
	n, _ := Engine.Where("user_id = ?", iUserId).Count(&Comments{UserId: iUserId})
	return n
}

// GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthor Get Commentses via CommentIdAndCommentPostIdAndCommentAuthor
func GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthor(offset int, limit int, CommentId_ int64, CommentPostId_ int64, CommentAuthor_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and comment_author = ?", CommentId_, CommentPostId_, CommentAuthor_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorEmail Get Commentses via CommentIdAndCommentPostIdAndCommentAuthorEmail
func GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorEmail(offset int, limit int, CommentId_ int64, CommentPostId_ int64, CommentAuthorEmail_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and comment_author_email = ?", CommentId_, CommentPostId_, CommentAuthorEmail_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorUrl Get Commentses via CommentIdAndCommentPostIdAndCommentAuthorUrl
func GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorUrl(offset int, limit int, CommentId_ int64, CommentPostId_ int64, CommentAuthorUrl_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and comment_author_url = ?", CommentId_, CommentPostId_, CommentAuthorUrl_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorIp Get Commentses via CommentIdAndCommentPostIdAndCommentAuthorIp
func GetCommentsesByCommentIdAndCommentPostIdAndCommentAuthorIp(offset int, limit int, CommentId_ int64, CommentPostId_ int64, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and comment_author_ip = ?", CommentId_, CommentPostId_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostIdAndCommentDate Get Commentses via CommentIdAndCommentPostIdAndCommentDate
func GetCommentsesByCommentIdAndCommentPostIdAndCommentDate(offset int, limit int, CommentId_ int64, CommentPostId_ int64, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and comment_date = ?", CommentId_, CommentPostId_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostIdAndCommentDateGmt Get Commentses via CommentIdAndCommentPostIdAndCommentDateGmt
func GetCommentsesByCommentIdAndCommentPostIdAndCommentDateGmt(offset int, limit int, CommentId_ int64, CommentPostId_ int64, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and comment_date_gmt = ?", CommentId_, CommentPostId_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostIdAndCommentContent Get Commentses via CommentIdAndCommentPostIdAndCommentContent
func GetCommentsesByCommentIdAndCommentPostIdAndCommentContent(offset int, limit int, CommentId_ int64, CommentPostId_ int64, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and comment_content = ?", CommentId_, CommentPostId_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostIdAndCommentKarma Get Commentses via CommentIdAndCommentPostIdAndCommentKarma
func GetCommentsesByCommentIdAndCommentPostIdAndCommentKarma(offset int, limit int, CommentId_ int64, CommentPostId_ int64, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and comment_karma = ?", CommentId_, CommentPostId_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostIdAndCommentApproved Get Commentses via CommentIdAndCommentPostIdAndCommentApproved
func GetCommentsesByCommentIdAndCommentPostIdAndCommentApproved(offset int, limit int, CommentId_ int64, CommentPostId_ int64, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and comment_approved = ?", CommentId_, CommentPostId_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostIdAndCommentAgent Get Commentses via CommentIdAndCommentPostIdAndCommentAgent
func GetCommentsesByCommentIdAndCommentPostIdAndCommentAgent(offset int, limit int, CommentId_ int64, CommentPostId_ int64, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and comment_agent = ?", CommentId_, CommentPostId_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostIdAndCommentType Get Commentses via CommentIdAndCommentPostIdAndCommentType
func GetCommentsesByCommentIdAndCommentPostIdAndCommentType(offset int, limit int, CommentId_ int64, CommentPostId_ int64, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and comment_type = ?", CommentId_, CommentPostId_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostIdAndCommentParent Get Commentses via CommentIdAndCommentPostIdAndCommentParent
func GetCommentsesByCommentIdAndCommentPostIdAndCommentParent(offset int, limit int, CommentId_ int64, CommentPostId_ int64, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and comment_parent = ?", CommentId_, CommentPostId_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostIdAndUserId Get Commentses via CommentIdAndCommentPostIdAndUserId
func GetCommentsesByCommentIdAndCommentPostIdAndUserId(offset int, limit int, CommentId_ int64, CommentPostId_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ? and user_id = ?", CommentId_, CommentPostId_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorEmail Get Commentses via CommentIdAndCommentAuthorAndCommentAuthorEmail
func GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorEmail(offset int, limit int, CommentId_ int64, CommentAuthor_ string, CommentAuthorEmail_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ? and comment_author_email = ?", CommentId_, CommentAuthor_, CommentAuthorEmail_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorUrl Get Commentses via CommentIdAndCommentAuthorAndCommentAuthorUrl
func GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorUrl(offset int, limit int, CommentId_ int64, CommentAuthor_ string, CommentAuthorUrl_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ? and comment_author_url = ?", CommentId_, CommentAuthor_, CommentAuthorUrl_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorIp Get Commentses via CommentIdAndCommentAuthorAndCommentAuthorIp
func GetCommentsesByCommentIdAndCommentAuthorAndCommentAuthorIp(offset int, limit int, CommentId_ int64, CommentAuthor_ string, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ? and comment_author_ip = ?", CommentId_, CommentAuthor_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorAndCommentDate Get Commentses via CommentIdAndCommentAuthorAndCommentDate
func GetCommentsesByCommentIdAndCommentAuthorAndCommentDate(offset int, limit int, CommentId_ int64, CommentAuthor_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ? and comment_date = ?", CommentId_, CommentAuthor_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorAndCommentDateGmt Get Commentses via CommentIdAndCommentAuthorAndCommentDateGmt
func GetCommentsesByCommentIdAndCommentAuthorAndCommentDateGmt(offset int, limit int, CommentId_ int64, CommentAuthor_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ? and comment_date_gmt = ?", CommentId_, CommentAuthor_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorAndCommentContent Get Commentses via CommentIdAndCommentAuthorAndCommentContent
func GetCommentsesByCommentIdAndCommentAuthorAndCommentContent(offset int, limit int, CommentId_ int64, CommentAuthor_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ? and comment_content = ?", CommentId_, CommentAuthor_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorAndCommentKarma Get Commentses via CommentIdAndCommentAuthorAndCommentKarma
func GetCommentsesByCommentIdAndCommentAuthorAndCommentKarma(offset int, limit int, CommentId_ int64, CommentAuthor_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ? and comment_karma = ?", CommentId_, CommentAuthor_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorAndCommentApproved Get Commentses via CommentIdAndCommentAuthorAndCommentApproved
func GetCommentsesByCommentIdAndCommentAuthorAndCommentApproved(offset int, limit int, CommentId_ int64, CommentAuthor_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ? and comment_approved = ?", CommentId_, CommentAuthor_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorAndCommentAgent Get Commentses via CommentIdAndCommentAuthorAndCommentAgent
func GetCommentsesByCommentIdAndCommentAuthorAndCommentAgent(offset int, limit int, CommentId_ int64, CommentAuthor_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ? and comment_agent = ?", CommentId_, CommentAuthor_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorAndCommentType Get Commentses via CommentIdAndCommentAuthorAndCommentType
func GetCommentsesByCommentIdAndCommentAuthorAndCommentType(offset int, limit int, CommentId_ int64, CommentAuthor_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ? and comment_type = ?", CommentId_, CommentAuthor_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorAndCommentParent Get Commentses via CommentIdAndCommentAuthorAndCommentParent
func GetCommentsesByCommentIdAndCommentAuthorAndCommentParent(offset int, limit int, CommentId_ int64, CommentAuthor_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ? and comment_parent = ?", CommentId_, CommentAuthor_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorAndUserId Get Commentses via CommentIdAndCommentAuthorAndUserId
func GetCommentsesByCommentIdAndCommentAuthorAndUserId(offset int, limit int, CommentId_ int64, CommentAuthor_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ? and user_id = ?", CommentId_, CommentAuthor_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAuthorUrl Get Commentses via CommentIdAndCommentAuthorEmailAndCommentAuthorUrl
func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAuthorUrl(offset int, limit int, CommentId_ int64, CommentAuthorEmail_ string, CommentAuthorUrl_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_email = ? and comment_author_url = ?", CommentId_, CommentAuthorEmail_, CommentAuthorUrl_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAuthorIp Get Commentses via CommentIdAndCommentAuthorEmailAndCommentAuthorIp
func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAuthorIp(offset int, limit int, CommentId_ int64, CommentAuthorEmail_ string, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_email = ? and comment_author_ip = ?", CommentId_, CommentAuthorEmail_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentDate Get Commentses via CommentIdAndCommentAuthorEmailAndCommentDate
func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentDate(offset int, limit int, CommentId_ int64, CommentAuthorEmail_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_email = ? and comment_date = ?", CommentId_, CommentAuthorEmail_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentDateGmt Get Commentses via CommentIdAndCommentAuthorEmailAndCommentDateGmt
func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentDateGmt(offset int, limit int, CommentId_ int64, CommentAuthorEmail_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_email = ? and comment_date_gmt = ?", CommentId_, CommentAuthorEmail_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentContent Get Commentses via CommentIdAndCommentAuthorEmailAndCommentContent
func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentContent(offset int, limit int, CommentId_ int64, CommentAuthorEmail_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_email = ? and comment_content = ?", CommentId_, CommentAuthorEmail_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentKarma Get Commentses via CommentIdAndCommentAuthorEmailAndCommentKarma
func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentKarma(offset int, limit int, CommentId_ int64, CommentAuthorEmail_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_email = ? and comment_karma = ?", CommentId_, CommentAuthorEmail_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentApproved Get Commentses via CommentIdAndCommentAuthorEmailAndCommentApproved
func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentApproved(offset int, limit int, CommentId_ int64, CommentAuthorEmail_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_email = ? and comment_approved = ?", CommentId_, CommentAuthorEmail_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAgent Get Commentses via CommentIdAndCommentAuthorEmailAndCommentAgent
func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentAgent(offset int, limit int, CommentId_ int64, CommentAuthorEmail_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_email = ? and comment_agent = ?", CommentId_, CommentAuthorEmail_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentType Get Commentses via CommentIdAndCommentAuthorEmailAndCommentType
func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentType(offset int, limit int, CommentId_ int64, CommentAuthorEmail_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_email = ? and comment_type = ?", CommentId_, CommentAuthorEmail_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentParent Get Commentses via CommentIdAndCommentAuthorEmailAndCommentParent
func GetCommentsesByCommentIdAndCommentAuthorEmailAndCommentParent(offset int, limit int, CommentId_ int64, CommentAuthorEmail_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_email = ? and comment_parent = ?", CommentId_, CommentAuthorEmail_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorEmailAndUserId Get Commentses via CommentIdAndCommentAuthorEmailAndUserId
func GetCommentsesByCommentIdAndCommentAuthorEmailAndUserId(offset int, limit int, CommentId_ int64, CommentAuthorEmail_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_email = ? and user_id = ?", CommentId_, CommentAuthorEmail_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentAuthorIp Get Commentses via CommentIdAndCommentAuthorUrlAndCommentAuthorIp
func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentAuthorIp(offset int, limit int, CommentId_ int64, CommentAuthorUrl_ string, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_url = ? and comment_author_ip = ?", CommentId_, CommentAuthorUrl_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentDate Get Commentses via CommentIdAndCommentAuthorUrlAndCommentDate
func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentDate(offset int, limit int, CommentId_ int64, CommentAuthorUrl_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_url = ? and comment_date = ?", CommentId_, CommentAuthorUrl_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentDateGmt Get Commentses via CommentIdAndCommentAuthorUrlAndCommentDateGmt
func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentDateGmt(offset int, limit int, CommentId_ int64, CommentAuthorUrl_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_url = ? and comment_date_gmt = ?", CommentId_, CommentAuthorUrl_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentContent Get Commentses via CommentIdAndCommentAuthorUrlAndCommentContent
func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentContent(offset int, limit int, CommentId_ int64, CommentAuthorUrl_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_url = ? and comment_content = ?", CommentId_, CommentAuthorUrl_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentKarma Get Commentses via CommentIdAndCommentAuthorUrlAndCommentKarma
func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentKarma(offset int, limit int, CommentId_ int64, CommentAuthorUrl_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_url = ? and comment_karma = ?", CommentId_, CommentAuthorUrl_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentApproved Get Commentses via CommentIdAndCommentAuthorUrlAndCommentApproved
func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentApproved(offset int, limit int, CommentId_ int64, CommentAuthorUrl_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_url = ? and comment_approved = ?", CommentId_, CommentAuthorUrl_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentAgent Get Commentses via CommentIdAndCommentAuthorUrlAndCommentAgent
func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentAgent(offset int, limit int, CommentId_ int64, CommentAuthorUrl_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_url = ? and comment_agent = ?", CommentId_, CommentAuthorUrl_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentType Get Commentses via CommentIdAndCommentAuthorUrlAndCommentType
func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentType(offset int, limit int, CommentId_ int64, CommentAuthorUrl_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_url = ? and comment_type = ?", CommentId_, CommentAuthorUrl_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentParent Get Commentses via CommentIdAndCommentAuthorUrlAndCommentParent
func GetCommentsesByCommentIdAndCommentAuthorUrlAndCommentParent(offset int, limit int, CommentId_ int64, CommentAuthorUrl_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_url = ? and comment_parent = ?", CommentId_, CommentAuthorUrl_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorUrlAndUserId Get Commentses via CommentIdAndCommentAuthorUrlAndUserId
func GetCommentsesByCommentIdAndCommentAuthorUrlAndUserId(offset int, limit int, CommentId_ int64, CommentAuthorUrl_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_url = ? and user_id = ?", CommentId_, CommentAuthorUrl_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorIpAndCommentDate Get Commentses via CommentIdAndCommentAuthorIpAndCommentDate
func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentDate(offset int, limit int, CommentId_ int64, CommentAuthorIp_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_ip = ? and comment_date = ?", CommentId_, CommentAuthorIp_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorIpAndCommentDateGmt Get Commentses via CommentIdAndCommentAuthorIpAndCommentDateGmt
func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentDateGmt(offset int, limit int, CommentId_ int64, CommentAuthorIp_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_ip = ? and comment_date_gmt = ?", CommentId_, CommentAuthorIp_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorIpAndCommentContent Get Commentses via CommentIdAndCommentAuthorIpAndCommentContent
func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentContent(offset int, limit int, CommentId_ int64, CommentAuthorIp_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_ip = ? and comment_content = ?", CommentId_, CommentAuthorIp_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorIpAndCommentKarma Get Commentses via CommentIdAndCommentAuthorIpAndCommentKarma
func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentKarma(offset int, limit int, CommentId_ int64, CommentAuthorIp_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_ip = ? and comment_karma = ?", CommentId_, CommentAuthorIp_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorIpAndCommentApproved Get Commentses via CommentIdAndCommentAuthorIpAndCommentApproved
func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentApproved(offset int, limit int, CommentId_ int64, CommentAuthorIp_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_ip = ? and comment_approved = ?", CommentId_, CommentAuthorIp_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorIpAndCommentAgent Get Commentses via CommentIdAndCommentAuthorIpAndCommentAgent
func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentAgent(offset int, limit int, CommentId_ int64, CommentAuthorIp_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_ip = ? and comment_agent = ?", CommentId_, CommentAuthorIp_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorIpAndCommentType Get Commentses via CommentIdAndCommentAuthorIpAndCommentType
func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentType(offset int, limit int, CommentId_ int64, CommentAuthorIp_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_ip = ? and comment_type = ?", CommentId_, CommentAuthorIp_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorIpAndCommentParent Get Commentses via CommentIdAndCommentAuthorIpAndCommentParent
func GetCommentsesByCommentIdAndCommentAuthorIpAndCommentParent(offset int, limit int, CommentId_ int64, CommentAuthorIp_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_ip = ? and comment_parent = ?", CommentId_, CommentAuthorIp_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorIpAndUserId Get Commentses via CommentIdAndCommentAuthorIpAndUserId
func GetCommentsesByCommentIdAndCommentAuthorIpAndUserId(offset int, limit int, CommentId_ int64, CommentAuthorIp_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_ip = ? and user_id = ?", CommentId_, CommentAuthorIp_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateAndCommentDateGmt Get Commentses via CommentIdAndCommentDateAndCommentDateGmt
func GetCommentsesByCommentIdAndCommentDateAndCommentDateGmt(offset int, limit int, CommentId_ int64, CommentDate_ time.Time, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date = ? and comment_date_gmt = ?", CommentId_, CommentDate_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateAndCommentContent Get Commentses via CommentIdAndCommentDateAndCommentContent
func GetCommentsesByCommentIdAndCommentDateAndCommentContent(offset int, limit int, CommentId_ int64, CommentDate_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date = ? and comment_content = ?", CommentId_, CommentDate_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateAndCommentKarma Get Commentses via CommentIdAndCommentDateAndCommentKarma
func GetCommentsesByCommentIdAndCommentDateAndCommentKarma(offset int, limit int, CommentId_ int64, CommentDate_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date = ? and comment_karma = ?", CommentId_, CommentDate_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateAndCommentApproved Get Commentses via CommentIdAndCommentDateAndCommentApproved
func GetCommentsesByCommentIdAndCommentDateAndCommentApproved(offset int, limit int, CommentId_ int64, CommentDate_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date = ? and comment_approved = ?", CommentId_, CommentDate_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateAndCommentAgent Get Commentses via CommentIdAndCommentDateAndCommentAgent
func GetCommentsesByCommentIdAndCommentDateAndCommentAgent(offset int, limit int, CommentId_ int64, CommentDate_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date = ? and comment_agent = ?", CommentId_, CommentDate_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateAndCommentType Get Commentses via CommentIdAndCommentDateAndCommentType
func GetCommentsesByCommentIdAndCommentDateAndCommentType(offset int, limit int, CommentId_ int64, CommentDate_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date = ? and comment_type = ?", CommentId_, CommentDate_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateAndCommentParent Get Commentses via CommentIdAndCommentDateAndCommentParent
func GetCommentsesByCommentIdAndCommentDateAndCommentParent(offset int, limit int, CommentId_ int64, CommentDate_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date = ? and comment_parent = ?", CommentId_, CommentDate_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateAndUserId Get Commentses via CommentIdAndCommentDateAndUserId
func GetCommentsesByCommentIdAndCommentDateAndUserId(offset int, limit int, CommentId_ int64, CommentDate_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date = ? and user_id = ?", CommentId_, CommentDate_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateGmtAndCommentContent Get Commentses via CommentIdAndCommentDateGmtAndCommentContent
func GetCommentsesByCommentIdAndCommentDateGmtAndCommentContent(offset int, limit int, CommentId_ int64, CommentDateGmt_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date_gmt = ? and comment_content = ?", CommentId_, CommentDateGmt_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateGmtAndCommentKarma Get Commentses via CommentIdAndCommentDateGmtAndCommentKarma
func GetCommentsesByCommentIdAndCommentDateGmtAndCommentKarma(offset int, limit int, CommentId_ int64, CommentDateGmt_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date_gmt = ? and comment_karma = ?", CommentId_, CommentDateGmt_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateGmtAndCommentApproved Get Commentses via CommentIdAndCommentDateGmtAndCommentApproved
func GetCommentsesByCommentIdAndCommentDateGmtAndCommentApproved(offset int, limit int, CommentId_ int64, CommentDateGmt_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date_gmt = ? and comment_approved = ?", CommentId_, CommentDateGmt_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateGmtAndCommentAgent Get Commentses via CommentIdAndCommentDateGmtAndCommentAgent
func GetCommentsesByCommentIdAndCommentDateGmtAndCommentAgent(offset int, limit int, CommentId_ int64, CommentDateGmt_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date_gmt = ? and comment_agent = ?", CommentId_, CommentDateGmt_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateGmtAndCommentType Get Commentses via CommentIdAndCommentDateGmtAndCommentType
func GetCommentsesByCommentIdAndCommentDateGmtAndCommentType(offset int, limit int, CommentId_ int64, CommentDateGmt_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date_gmt = ? and comment_type = ?", CommentId_, CommentDateGmt_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateGmtAndCommentParent Get Commentses via CommentIdAndCommentDateGmtAndCommentParent
func GetCommentsesByCommentIdAndCommentDateGmtAndCommentParent(offset int, limit int, CommentId_ int64, CommentDateGmt_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date_gmt = ? and comment_parent = ?", CommentId_, CommentDateGmt_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateGmtAndUserId Get Commentses via CommentIdAndCommentDateGmtAndUserId
func GetCommentsesByCommentIdAndCommentDateGmtAndUserId(offset int, limit int, CommentId_ int64, CommentDateGmt_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date_gmt = ? and user_id = ?", CommentId_, CommentDateGmt_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentContentAndCommentKarma Get Commentses via CommentIdAndCommentContentAndCommentKarma
func GetCommentsesByCommentIdAndCommentContentAndCommentKarma(offset int, limit int, CommentId_ int64, CommentContent_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_content = ? and comment_karma = ?", CommentId_, CommentContent_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentContentAndCommentApproved Get Commentses via CommentIdAndCommentContentAndCommentApproved
func GetCommentsesByCommentIdAndCommentContentAndCommentApproved(offset int, limit int, CommentId_ int64, CommentContent_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_content = ? and comment_approved = ?", CommentId_, CommentContent_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentContentAndCommentAgent Get Commentses via CommentIdAndCommentContentAndCommentAgent
func GetCommentsesByCommentIdAndCommentContentAndCommentAgent(offset int, limit int, CommentId_ int64, CommentContent_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_content = ? and comment_agent = ?", CommentId_, CommentContent_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentContentAndCommentType Get Commentses via CommentIdAndCommentContentAndCommentType
func GetCommentsesByCommentIdAndCommentContentAndCommentType(offset int, limit int, CommentId_ int64, CommentContent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_content = ? and comment_type = ?", CommentId_, CommentContent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentContentAndCommentParent Get Commentses via CommentIdAndCommentContentAndCommentParent
func GetCommentsesByCommentIdAndCommentContentAndCommentParent(offset int, limit int, CommentId_ int64, CommentContent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_content = ? and comment_parent = ?", CommentId_, CommentContent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentContentAndUserId Get Commentses via CommentIdAndCommentContentAndUserId
func GetCommentsesByCommentIdAndCommentContentAndUserId(offset int, limit int, CommentId_ int64, CommentContent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_content = ? and user_id = ?", CommentId_, CommentContent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentKarmaAndCommentApproved Get Commentses via CommentIdAndCommentKarmaAndCommentApproved
func GetCommentsesByCommentIdAndCommentKarmaAndCommentApproved(offset int, limit int, CommentId_ int64, CommentKarma_ int, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_karma = ? and comment_approved = ?", CommentId_, CommentKarma_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentKarmaAndCommentAgent Get Commentses via CommentIdAndCommentKarmaAndCommentAgent
func GetCommentsesByCommentIdAndCommentKarmaAndCommentAgent(offset int, limit int, CommentId_ int64, CommentKarma_ int, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_karma = ? and comment_agent = ?", CommentId_, CommentKarma_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentKarmaAndCommentType Get Commentses via CommentIdAndCommentKarmaAndCommentType
func GetCommentsesByCommentIdAndCommentKarmaAndCommentType(offset int, limit int, CommentId_ int64, CommentKarma_ int, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_karma = ? and comment_type = ?", CommentId_, CommentKarma_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentKarmaAndCommentParent Get Commentses via CommentIdAndCommentKarmaAndCommentParent
func GetCommentsesByCommentIdAndCommentKarmaAndCommentParent(offset int, limit int, CommentId_ int64, CommentKarma_ int, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_karma = ? and comment_parent = ?", CommentId_, CommentKarma_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentKarmaAndUserId Get Commentses via CommentIdAndCommentKarmaAndUserId
func GetCommentsesByCommentIdAndCommentKarmaAndUserId(offset int, limit int, CommentId_ int64, CommentKarma_ int, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_karma = ? and user_id = ?", CommentId_, CommentKarma_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentApprovedAndCommentAgent Get Commentses via CommentIdAndCommentApprovedAndCommentAgent
func GetCommentsesByCommentIdAndCommentApprovedAndCommentAgent(offset int, limit int, CommentId_ int64, CommentApproved_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_approved = ? and comment_agent = ?", CommentId_, CommentApproved_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentApprovedAndCommentType Get Commentses via CommentIdAndCommentApprovedAndCommentType
func GetCommentsesByCommentIdAndCommentApprovedAndCommentType(offset int, limit int, CommentId_ int64, CommentApproved_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_approved = ? and comment_type = ?", CommentId_, CommentApproved_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentApprovedAndCommentParent Get Commentses via CommentIdAndCommentApprovedAndCommentParent
func GetCommentsesByCommentIdAndCommentApprovedAndCommentParent(offset int, limit int, CommentId_ int64, CommentApproved_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_approved = ? and comment_parent = ?", CommentId_, CommentApproved_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentApprovedAndUserId Get Commentses via CommentIdAndCommentApprovedAndUserId
func GetCommentsesByCommentIdAndCommentApprovedAndUserId(offset int, limit int, CommentId_ int64, CommentApproved_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_approved = ? and user_id = ?", CommentId_, CommentApproved_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAgentAndCommentType Get Commentses via CommentIdAndCommentAgentAndCommentType
func GetCommentsesByCommentIdAndCommentAgentAndCommentType(offset int, limit int, CommentId_ int64, CommentAgent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_agent = ? and comment_type = ?", CommentId_, CommentAgent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAgentAndCommentParent Get Commentses via CommentIdAndCommentAgentAndCommentParent
func GetCommentsesByCommentIdAndCommentAgentAndCommentParent(offset int, limit int, CommentId_ int64, CommentAgent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_agent = ? and comment_parent = ?", CommentId_, CommentAgent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAgentAndUserId Get Commentses via CommentIdAndCommentAgentAndUserId
func GetCommentsesByCommentIdAndCommentAgentAndUserId(offset int, limit int, CommentId_ int64, CommentAgent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_agent = ? and user_id = ?", CommentId_, CommentAgent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentTypeAndCommentParent Get Commentses via CommentIdAndCommentTypeAndCommentParent
func GetCommentsesByCommentIdAndCommentTypeAndCommentParent(offset int, limit int, CommentId_ int64, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_type = ? and comment_parent = ?", CommentId_, CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentTypeAndUserId Get Commentses via CommentIdAndCommentTypeAndUserId
func GetCommentsesByCommentIdAndCommentTypeAndUserId(offset int, limit int, CommentId_ int64, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_type = ? and user_id = ?", CommentId_, CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentParentAndUserId Get Commentses via CommentIdAndCommentParentAndUserId
func GetCommentsesByCommentIdAndCommentParentAndUserId(offset int, limit int, CommentId_ int64, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_parent = ? and user_id = ?", CommentId_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorEmail Get Commentses via CommentPostIdAndCommentAuthorAndCommentAuthorEmail
func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorEmail(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string, CommentAuthorEmail_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ? and comment_author_email = ?", CommentPostId_, CommentAuthor_, CommentAuthorEmail_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorUrl Get Commentses via CommentPostIdAndCommentAuthorAndCommentAuthorUrl
func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorUrl(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string, CommentAuthorUrl_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ? and comment_author_url = ?", CommentPostId_, CommentAuthor_, CommentAuthorUrl_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorIp Get Commentses via CommentPostIdAndCommentAuthorAndCommentAuthorIp
func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAuthorIp(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ? and comment_author_ip = ?", CommentPostId_, CommentAuthor_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorAndCommentDate Get Commentses via CommentPostIdAndCommentAuthorAndCommentDate
func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentDate(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ? and comment_date = ?", CommentPostId_, CommentAuthor_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorAndCommentDateGmt Get Commentses via CommentPostIdAndCommentAuthorAndCommentDateGmt
func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentDateGmt(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ? and comment_date_gmt = ?", CommentPostId_, CommentAuthor_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorAndCommentContent Get Commentses via CommentPostIdAndCommentAuthorAndCommentContent
func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentContent(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ? and comment_content = ?", CommentPostId_, CommentAuthor_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorAndCommentKarma Get Commentses via CommentPostIdAndCommentAuthorAndCommentKarma
func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentKarma(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ? and comment_karma = ?", CommentPostId_, CommentAuthor_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorAndCommentApproved Get Commentses via CommentPostIdAndCommentAuthorAndCommentApproved
func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentApproved(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ? and comment_approved = ?", CommentPostId_, CommentAuthor_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAgent Get Commentses via CommentPostIdAndCommentAuthorAndCommentAgent
func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentAgent(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ? and comment_agent = ?", CommentPostId_, CommentAuthor_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorAndCommentType Get Commentses via CommentPostIdAndCommentAuthorAndCommentType
func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentType(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ? and comment_type = ?", CommentPostId_, CommentAuthor_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorAndCommentParent Get Commentses via CommentPostIdAndCommentAuthorAndCommentParent
func GetCommentsesByCommentPostIdAndCommentAuthorAndCommentParent(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ? and comment_parent = ?", CommentPostId_, CommentAuthor_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorAndUserId Get Commentses via CommentPostIdAndCommentAuthorAndUserId
func GetCommentsesByCommentPostIdAndCommentAuthorAndUserId(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ? and user_id = ?", CommentPostId_, CommentAuthor_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAuthorUrl Get Commentses via CommentPostIdAndCommentAuthorEmailAndCommentAuthorUrl
func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAuthorUrl(offset int, limit int, CommentPostId_ int64, CommentAuthorEmail_ string, CommentAuthorUrl_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_email = ? and comment_author_url = ?", CommentPostId_, CommentAuthorEmail_, CommentAuthorUrl_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAuthorIp Get Commentses via CommentPostIdAndCommentAuthorEmailAndCommentAuthorIp
func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAuthorIp(offset int, limit int, CommentPostId_ int64, CommentAuthorEmail_ string, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_email = ? and comment_author_ip = ?", CommentPostId_, CommentAuthorEmail_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentDate Get Commentses via CommentPostIdAndCommentAuthorEmailAndCommentDate
func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentDate(offset int, limit int, CommentPostId_ int64, CommentAuthorEmail_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_email = ? and comment_date = ?", CommentPostId_, CommentAuthorEmail_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentDateGmt Get Commentses via CommentPostIdAndCommentAuthorEmailAndCommentDateGmt
func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentDateGmt(offset int, limit int, CommentPostId_ int64, CommentAuthorEmail_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_email = ? and comment_date_gmt = ?", CommentPostId_, CommentAuthorEmail_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentContent Get Commentses via CommentPostIdAndCommentAuthorEmailAndCommentContent
func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentContent(offset int, limit int, CommentPostId_ int64, CommentAuthorEmail_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_email = ? and comment_content = ?", CommentPostId_, CommentAuthorEmail_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentKarma Get Commentses via CommentPostIdAndCommentAuthorEmailAndCommentKarma
func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentKarma(offset int, limit int, CommentPostId_ int64, CommentAuthorEmail_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_email = ? and comment_karma = ?", CommentPostId_, CommentAuthorEmail_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentApproved Get Commentses via CommentPostIdAndCommentAuthorEmailAndCommentApproved
func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentApproved(offset int, limit int, CommentPostId_ int64, CommentAuthorEmail_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_email = ? and comment_approved = ?", CommentPostId_, CommentAuthorEmail_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAgent Get Commentses via CommentPostIdAndCommentAuthorEmailAndCommentAgent
func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentAgent(offset int, limit int, CommentPostId_ int64, CommentAuthorEmail_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_email = ? and comment_agent = ?", CommentPostId_, CommentAuthorEmail_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentType Get Commentses via CommentPostIdAndCommentAuthorEmailAndCommentType
func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentType(offset int, limit int, CommentPostId_ int64, CommentAuthorEmail_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_email = ? and comment_type = ?", CommentPostId_, CommentAuthorEmail_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentParent Get Commentses via CommentPostIdAndCommentAuthorEmailAndCommentParent
func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndCommentParent(offset int, limit int, CommentPostId_ int64, CommentAuthorEmail_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_email = ? and comment_parent = ?", CommentPostId_, CommentAuthorEmail_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorEmailAndUserId Get Commentses via CommentPostIdAndCommentAuthorEmailAndUserId
func GetCommentsesByCommentPostIdAndCommentAuthorEmailAndUserId(offset int, limit int, CommentPostId_ int64, CommentAuthorEmail_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_email = ? and user_id = ?", CommentPostId_, CommentAuthorEmail_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentAuthorIp Get Commentses via CommentPostIdAndCommentAuthorUrlAndCommentAuthorIp
func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentAuthorIp(offset int, limit int, CommentPostId_ int64, CommentAuthorUrl_ string, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_url = ? and comment_author_ip = ?", CommentPostId_, CommentAuthorUrl_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentDate Get Commentses via CommentPostIdAndCommentAuthorUrlAndCommentDate
func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentDate(offset int, limit int, CommentPostId_ int64, CommentAuthorUrl_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_url = ? and comment_date = ?", CommentPostId_, CommentAuthorUrl_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentDateGmt Get Commentses via CommentPostIdAndCommentAuthorUrlAndCommentDateGmt
func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentDateGmt(offset int, limit int, CommentPostId_ int64, CommentAuthorUrl_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_url = ? and comment_date_gmt = ?", CommentPostId_, CommentAuthorUrl_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentContent Get Commentses via CommentPostIdAndCommentAuthorUrlAndCommentContent
func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentContent(offset int, limit int, CommentPostId_ int64, CommentAuthorUrl_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_url = ? and comment_content = ?", CommentPostId_, CommentAuthorUrl_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentKarma Get Commentses via CommentPostIdAndCommentAuthorUrlAndCommentKarma
func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentKarma(offset int, limit int, CommentPostId_ int64, CommentAuthorUrl_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_url = ? and comment_karma = ?", CommentPostId_, CommentAuthorUrl_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentApproved Get Commentses via CommentPostIdAndCommentAuthorUrlAndCommentApproved
func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentApproved(offset int, limit int, CommentPostId_ int64, CommentAuthorUrl_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_url = ? and comment_approved = ?", CommentPostId_, CommentAuthorUrl_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentAgent Get Commentses via CommentPostIdAndCommentAuthorUrlAndCommentAgent
func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentAgent(offset int, limit int, CommentPostId_ int64, CommentAuthorUrl_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_url = ? and comment_agent = ?", CommentPostId_, CommentAuthorUrl_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentType Get Commentses via CommentPostIdAndCommentAuthorUrlAndCommentType
func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentType(offset int, limit int, CommentPostId_ int64, CommentAuthorUrl_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_url = ? and comment_type = ?", CommentPostId_, CommentAuthorUrl_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentParent Get Commentses via CommentPostIdAndCommentAuthorUrlAndCommentParent
func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndCommentParent(offset int, limit int, CommentPostId_ int64, CommentAuthorUrl_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_url = ? and comment_parent = ?", CommentPostId_, CommentAuthorUrl_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorUrlAndUserId Get Commentses via CommentPostIdAndCommentAuthorUrlAndUserId
func GetCommentsesByCommentPostIdAndCommentAuthorUrlAndUserId(offset int, limit int, CommentPostId_ int64, CommentAuthorUrl_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_url = ? and user_id = ?", CommentPostId_, CommentAuthorUrl_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentDate Get Commentses via CommentPostIdAndCommentAuthorIpAndCommentDate
func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentDate(offset int, limit int, CommentPostId_ int64, CommentAuthorIp_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_ip = ? and comment_date = ?", CommentPostId_, CommentAuthorIp_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentDateGmt Get Commentses via CommentPostIdAndCommentAuthorIpAndCommentDateGmt
func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentDateGmt(offset int, limit int, CommentPostId_ int64, CommentAuthorIp_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_ip = ? and comment_date_gmt = ?", CommentPostId_, CommentAuthorIp_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentContent Get Commentses via CommentPostIdAndCommentAuthorIpAndCommentContent
func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentContent(offset int, limit int, CommentPostId_ int64, CommentAuthorIp_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_ip = ? and comment_content = ?", CommentPostId_, CommentAuthorIp_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentKarma Get Commentses via CommentPostIdAndCommentAuthorIpAndCommentKarma
func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentKarma(offset int, limit int, CommentPostId_ int64, CommentAuthorIp_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_ip = ? and comment_karma = ?", CommentPostId_, CommentAuthorIp_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentApproved Get Commentses via CommentPostIdAndCommentAuthorIpAndCommentApproved
func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentApproved(offset int, limit int, CommentPostId_ int64, CommentAuthorIp_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_ip = ? and comment_approved = ?", CommentPostId_, CommentAuthorIp_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentAgent Get Commentses via CommentPostIdAndCommentAuthorIpAndCommentAgent
func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentAgent(offset int, limit int, CommentPostId_ int64, CommentAuthorIp_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_ip = ? and comment_agent = ?", CommentPostId_, CommentAuthorIp_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentType Get Commentses via CommentPostIdAndCommentAuthorIpAndCommentType
func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentType(offset int, limit int, CommentPostId_ int64, CommentAuthorIp_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_ip = ? and comment_type = ?", CommentPostId_, CommentAuthorIp_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentParent Get Commentses via CommentPostIdAndCommentAuthorIpAndCommentParent
func GetCommentsesByCommentPostIdAndCommentAuthorIpAndCommentParent(offset int, limit int, CommentPostId_ int64, CommentAuthorIp_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_ip = ? and comment_parent = ?", CommentPostId_, CommentAuthorIp_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorIpAndUserId Get Commentses via CommentPostIdAndCommentAuthorIpAndUserId
func GetCommentsesByCommentPostIdAndCommentAuthorIpAndUserId(offset int, limit int, CommentPostId_ int64, CommentAuthorIp_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_ip = ? and user_id = ?", CommentPostId_, CommentAuthorIp_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateAndCommentDateGmt Get Commentses via CommentPostIdAndCommentDateAndCommentDateGmt
func GetCommentsesByCommentPostIdAndCommentDateAndCommentDateGmt(offset int, limit int, CommentPostId_ int64, CommentDate_ time.Time, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date = ? and comment_date_gmt = ?", CommentPostId_, CommentDate_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateAndCommentContent Get Commentses via CommentPostIdAndCommentDateAndCommentContent
func GetCommentsesByCommentPostIdAndCommentDateAndCommentContent(offset int, limit int, CommentPostId_ int64, CommentDate_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date = ? and comment_content = ?", CommentPostId_, CommentDate_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateAndCommentKarma Get Commentses via CommentPostIdAndCommentDateAndCommentKarma
func GetCommentsesByCommentPostIdAndCommentDateAndCommentKarma(offset int, limit int, CommentPostId_ int64, CommentDate_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date = ? and comment_karma = ?", CommentPostId_, CommentDate_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateAndCommentApproved Get Commentses via CommentPostIdAndCommentDateAndCommentApproved
func GetCommentsesByCommentPostIdAndCommentDateAndCommentApproved(offset int, limit int, CommentPostId_ int64, CommentDate_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date = ? and comment_approved = ?", CommentPostId_, CommentDate_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateAndCommentAgent Get Commentses via CommentPostIdAndCommentDateAndCommentAgent
func GetCommentsesByCommentPostIdAndCommentDateAndCommentAgent(offset int, limit int, CommentPostId_ int64, CommentDate_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date = ? and comment_agent = ?", CommentPostId_, CommentDate_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateAndCommentType Get Commentses via CommentPostIdAndCommentDateAndCommentType
func GetCommentsesByCommentPostIdAndCommentDateAndCommentType(offset int, limit int, CommentPostId_ int64, CommentDate_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date = ? and comment_type = ?", CommentPostId_, CommentDate_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateAndCommentParent Get Commentses via CommentPostIdAndCommentDateAndCommentParent
func GetCommentsesByCommentPostIdAndCommentDateAndCommentParent(offset int, limit int, CommentPostId_ int64, CommentDate_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date = ? and comment_parent = ?", CommentPostId_, CommentDate_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateAndUserId Get Commentses via CommentPostIdAndCommentDateAndUserId
func GetCommentsesByCommentPostIdAndCommentDateAndUserId(offset int, limit int, CommentPostId_ int64, CommentDate_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date = ? and user_id = ?", CommentPostId_, CommentDate_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentContent Get Commentses via CommentPostIdAndCommentDateGmtAndCommentContent
func GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentContent(offset int, limit int, CommentPostId_ int64, CommentDateGmt_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date_gmt = ? and comment_content = ?", CommentPostId_, CommentDateGmt_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentKarma Get Commentses via CommentPostIdAndCommentDateGmtAndCommentKarma
func GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentKarma(offset int, limit int, CommentPostId_ int64, CommentDateGmt_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date_gmt = ? and comment_karma = ?", CommentPostId_, CommentDateGmt_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentApproved Get Commentses via CommentPostIdAndCommentDateGmtAndCommentApproved
func GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentApproved(offset int, limit int, CommentPostId_ int64, CommentDateGmt_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date_gmt = ? and comment_approved = ?", CommentPostId_, CommentDateGmt_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentAgent Get Commentses via CommentPostIdAndCommentDateGmtAndCommentAgent
func GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentAgent(offset int, limit int, CommentPostId_ int64, CommentDateGmt_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date_gmt = ? and comment_agent = ?", CommentPostId_, CommentDateGmt_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentType Get Commentses via CommentPostIdAndCommentDateGmtAndCommentType
func GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentType(offset int, limit int, CommentPostId_ int64, CommentDateGmt_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date_gmt = ? and comment_type = ?", CommentPostId_, CommentDateGmt_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentParent Get Commentses via CommentPostIdAndCommentDateGmtAndCommentParent
func GetCommentsesByCommentPostIdAndCommentDateGmtAndCommentParent(offset int, limit int, CommentPostId_ int64, CommentDateGmt_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date_gmt = ? and comment_parent = ?", CommentPostId_, CommentDateGmt_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateGmtAndUserId Get Commentses via CommentPostIdAndCommentDateGmtAndUserId
func GetCommentsesByCommentPostIdAndCommentDateGmtAndUserId(offset int, limit int, CommentPostId_ int64, CommentDateGmt_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date_gmt = ? and user_id = ?", CommentPostId_, CommentDateGmt_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentContentAndCommentKarma Get Commentses via CommentPostIdAndCommentContentAndCommentKarma
func GetCommentsesByCommentPostIdAndCommentContentAndCommentKarma(offset int, limit int, CommentPostId_ int64, CommentContent_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_content = ? and comment_karma = ?", CommentPostId_, CommentContent_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentContentAndCommentApproved Get Commentses via CommentPostIdAndCommentContentAndCommentApproved
func GetCommentsesByCommentPostIdAndCommentContentAndCommentApproved(offset int, limit int, CommentPostId_ int64, CommentContent_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_content = ? and comment_approved = ?", CommentPostId_, CommentContent_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentContentAndCommentAgent Get Commentses via CommentPostIdAndCommentContentAndCommentAgent
func GetCommentsesByCommentPostIdAndCommentContentAndCommentAgent(offset int, limit int, CommentPostId_ int64, CommentContent_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_content = ? and comment_agent = ?", CommentPostId_, CommentContent_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentContentAndCommentType Get Commentses via CommentPostIdAndCommentContentAndCommentType
func GetCommentsesByCommentPostIdAndCommentContentAndCommentType(offset int, limit int, CommentPostId_ int64, CommentContent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_content = ? and comment_type = ?", CommentPostId_, CommentContent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentContentAndCommentParent Get Commentses via CommentPostIdAndCommentContentAndCommentParent
func GetCommentsesByCommentPostIdAndCommentContentAndCommentParent(offset int, limit int, CommentPostId_ int64, CommentContent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_content = ? and comment_parent = ?", CommentPostId_, CommentContent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentContentAndUserId Get Commentses via CommentPostIdAndCommentContentAndUserId
func GetCommentsesByCommentPostIdAndCommentContentAndUserId(offset int, limit int, CommentPostId_ int64, CommentContent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_content = ? and user_id = ?", CommentPostId_, CommentContent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentKarmaAndCommentApproved Get Commentses via CommentPostIdAndCommentKarmaAndCommentApproved
func GetCommentsesByCommentPostIdAndCommentKarmaAndCommentApproved(offset int, limit int, CommentPostId_ int64, CommentKarma_ int, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_karma = ? and comment_approved = ?", CommentPostId_, CommentKarma_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentKarmaAndCommentAgent Get Commentses via CommentPostIdAndCommentKarmaAndCommentAgent
func GetCommentsesByCommentPostIdAndCommentKarmaAndCommentAgent(offset int, limit int, CommentPostId_ int64, CommentKarma_ int, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_karma = ? and comment_agent = ?", CommentPostId_, CommentKarma_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentKarmaAndCommentType Get Commentses via CommentPostIdAndCommentKarmaAndCommentType
func GetCommentsesByCommentPostIdAndCommentKarmaAndCommentType(offset int, limit int, CommentPostId_ int64, CommentKarma_ int, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_karma = ? and comment_type = ?", CommentPostId_, CommentKarma_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentKarmaAndCommentParent Get Commentses via CommentPostIdAndCommentKarmaAndCommentParent
func GetCommentsesByCommentPostIdAndCommentKarmaAndCommentParent(offset int, limit int, CommentPostId_ int64, CommentKarma_ int, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_karma = ? and comment_parent = ?", CommentPostId_, CommentKarma_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentKarmaAndUserId Get Commentses via CommentPostIdAndCommentKarmaAndUserId
func GetCommentsesByCommentPostIdAndCommentKarmaAndUserId(offset int, limit int, CommentPostId_ int64, CommentKarma_ int, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_karma = ? and user_id = ?", CommentPostId_, CommentKarma_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentApprovedAndCommentAgent Get Commentses via CommentPostIdAndCommentApprovedAndCommentAgent
func GetCommentsesByCommentPostIdAndCommentApprovedAndCommentAgent(offset int, limit int, CommentPostId_ int64, CommentApproved_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_approved = ? and comment_agent = ?", CommentPostId_, CommentApproved_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentApprovedAndCommentType Get Commentses via CommentPostIdAndCommentApprovedAndCommentType
func GetCommentsesByCommentPostIdAndCommentApprovedAndCommentType(offset int, limit int, CommentPostId_ int64, CommentApproved_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_approved = ? and comment_type = ?", CommentPostId_, CommentApproved_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentApprovedAndCommentParent Get Commentses via CommentPostIdAndCommentApprovedAndCommentParent
func GetCommentsesByCommentPostIdAndCommentApprovedAndCommentParent(offset int, limit int, CommentPostId_ int64, CommentApproved_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_approved = ? and comment_parent = ?", CommentPostId_, CommentApproved_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentApprovedAndUserId Get Commentses via CommentPostIdAndCommentApprovedAndUserId
func GetCommentsesByCommentPostIdAndCommentApprovedAndUserId(offset int, limit int, CommentPostId_ int64, CommentApproved_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_approved = ? and user_id = ?", CommentPostId_, CommentApproved_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAgentAndCommentType Get Commentses via CommentPostIdAndCommentAgentAndCommentType
func GetCommentsesByCommentPostIdAndCommentAgentAndCommentType(offset int, limit int, CommentPostId_ int64, CommentAgent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_agent = ? and comment_type = ?", CommentPostId_, CommentAgent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAgentAndCommentParent Get Commentses via CommentPostIdAndCommentAgentAndCommentParent
func GetCommentsesByCommentPostIdAndCommentAgentAndCommentParent(offset int, limit int, CommentPostId_ int64, CommentAgent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_agent = ? and comment_parent = ?", CommentPostId_, CommentAgent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAgentAndUserId Get Commentses via CommentPostIdAndCommentAgentAndUserId
func GetCommentsesByCommentPostIdAndCommentAgentAndUserId(offset int, limit int, CommentPostId_ int64, CommentAgent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_agent = ? and user_id = ?", CommentPostId_, CommentAgent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentTypeAndCommentParent Get Commentses via CommentPostIdAndCommentTypeAndCommentParent
func GetCommentsesByCommentPostIdAndCommentTypeAndCommentParent(offset int, limit int, CommentPostId_ int64, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_type = ? and comment_parent = ?", CommentPostId_, CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentTypeAndUserId Get Commentses via CommentPostIdAndCommentTypeAndUserId
func GetCommentsesByCommentPostIdAndCommentTypeAndUserId(offset int, limit int, CommentPostId_ int64, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_type = ? and user_id = ?", CommentPostId_, CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentParentAndUserId Get Commentses via CommentPostIdAndCommentParentAndUserId
func GetCommentsesByCommentPostIdAndCommentParentAndUserId(offset int, limit int, CommentPostId_ int64, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_parent = ? and user_id = ?", CommentPostId_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAuthorUrl Get Commentses via CommentAuthorAndCommentAuthorEmailAndCommentAuthorUrl
func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAuthorUrl(offset int, limit int, CommentAuthor_ string, CommentAuthorEmail_ string, CommentAuthorUrl_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_email = ? and comment_author_url = ?", CommentAuthor_, CommentAuthorEmail_, CommentAuthorUrl_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAuthorIp Get Commentses via CommentAuthorAndCommentAuthorEmailAndCommentAuthorIp
func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAuthorIp(offset int, limit int, CommentAuthor_ string, CommentAuthorEmail_ string, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_email = ? and comment_author_ip = ?", CommentAuthor_, CommentAuthorEmail_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentDate Get Commentses via CommentAuthorAndCommentAuthorEmailAndCommentDate
func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentDate(offset int, limit int, CommentAuthor_ string, CommentAuthorEmail_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_email = ? and comment_date = ?", CommentAuthor_, CommentAuthorEmail_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentDateGmt Get Commentses via CommentAuthorAndCommentAuthorEmailAndCommentDateGmt
func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentDateGmt(offset int, limit int, CommentAuthor_ string, CommentAuthorEmail_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_email = ? and comment_date_gmt = ?", CommentAuthor_, CommentAuthorEmail_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentContent Get Commentses via CommentAuthorAndCommentAuthorEmailAndCommentContent
func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentContent(offset int, limit int, CommentAuthor_ string, CommentAuthorEmail_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_email = ? and comment_content = ?", CommentAuthor_, CommentAuthorEmail_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentKarma Get Commentses via CommentAuthorAndCommentAuthorEmailAndCommentKarma
func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentKarma(offset int, limit int, CommentAuthor_ string, CommentAuthorEmail_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_email = ? and comment_karma = ?", CommentAuthor_, CommentAuthorEmail_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentApproved Get Commentses via CommentAuthorAndCommentAuthorEmailAndCommentApproved
func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentApproved(offset int, limit int, CommentAuthor_ string, CommentAuthorEmail_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_email = ? and comment_approved = ?", CommentAuthor_, CommentAuthorEmail_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAgent Get Commentses via CommentAuthorAndCommentAuthorEmailAndCommentAgent
func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentAgent(offset int, limit int, CommentAuthor_ string, CommentAuthorEmail_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_email = ? and comment_agent = ?", CommentAuthor_, CommentAuthorEmail_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentType Get Commentses via CommentAuthorAndCommentAuthorEmailAndCommentType
func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentType(offset int, limit int, CommentAuthor_ string, CommentAuthorEmail_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_email = ? and comment_type = ?", CommentAuthor_, CommentAuthorEmail_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentParent Get Commentses via CommentAuthorAndCommentAuthorEmailAndCommentParent
func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndCommentParent(offset int, limit int, CommentAuthor_ string, CommentAuthorEmail_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_email = ? and comment_parent = ?", CommentAuthor_, CommentAuthorEmail_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorEmailAndUserId Get Commentses via CommentAuthorAndCommentAuthorEmailAndUserId
func GetCommentsesByCommentAuthorAndCommentAuthorEmailAndUserId(offset int, limit int, CommentAuthor_ string, CommentAuthorEmail_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_email = ? and user_id = ?", CommentAuthor_, CommentAuthorEmail_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentAuthorIp Get Commentses via CommentAuthorAndCommentAuthorUrlAndCommentAuthorIp
func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentAuthorIp(offset int, limit int, CommentAuthor_ string, CommentAuthorUrl_ string, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_url = ? and comment_author_ip = ?", CommentAuthor_, CommentAuthorUrl_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentDate Get Commentses via CommentAuthorAndCommentAuthorUrlAndCommentDate
func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentDate(offset int, limit int, CommentAuthor_ string, CommentAuthorUrl_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_url = ? and comment_date = ?", CommentAuthor_, CommentAuthorUrl_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentDateGmt Get Commentses via CommentAuthorAndCommentAuthorUrlAndCommentDateGmt
func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentDateGmt(offset int, limit int, CommentAuthor_ string, CommentAuthorUrl_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_url = ? and comment_date_gmt = ?", CommentAuthor_, CommentAuthorUrl_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentContent Get Commentses via CommentAuthorAndCommentAuthorUrlAndCommentContent
func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentContent(offset int, limit int, CommentAuthor_ string, CommentAuthorUrl_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_url = ? and comment_content = ?", CommentAuthor_, CommentAuthorUrl_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentKarma Get Commentses via CommentAuthorAndCommentAuthorUrlAndCommentKarma
func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentKarma(offset int, limit int, CommentAuthor_ string, CommentAuthorUrl_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_url = ? and comment_karma = ?", CommentAuthor_, CommentAuthorUrl_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentApproved Get Commentses via CommentAuthorAndCommentAuthorUrlAndCommentApproved
func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentApproved(offset int, limit int, CommentAuthor_ string, CommentAuthorUrl_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_url = ? and comment_approved = ?", CommentAuthor_, CommentAuthorUrl_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentAgent Get Commentses via CommentAuthorAndCommentAuthorUrlAndCommentAgent
func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentAgent(offset int, limit int, CommentAuthor_ string, CommentAuthorUrl_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_url = ? and comment_agent = ?", CommentAuthor_, CommentAuthorUrl_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentType Get Commentses via CommentAuthorAndCommentAuthorUrlAndCommentType
func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentType(offset int, limit int, CommentAuthor_ string, CommentAuthorUrl_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_url = ? and comment_type = ?", CommentAuthor_, CommentAuthorUrl_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentParent Get Commentses via CommentAuthorAndCommentAuthorUrlAndCommentParent
func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndCommentParent(offset int, limit int, CommentAuthor_ string, CommentAuthorUrl_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_url = ? and comment_parent = ?", CommentAuthor_, CommentAuthorUrl_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorUrlAndUserId Get Commentses via CommentAuthorAndCommentAuthorUrlAndUserId
func GetCommentsesByCommentAuthorAndCommentAuthorUrlAndUserId(offset int, limit int, CommentAuthor_ string, CommentAuthorUrl_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_url = ? and user_id = ?", CommentAuthor_, CommentAuthorUrl_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentDate Get Commentses via CommentAuthorAndCommentAuthorIpAndCommentDate
func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentDate(offset int, limit int, CommentAuthor_ string, CommentAuthorIp_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_ip = ? and comment_date = ?", CommentAuthor_, CommentAuthorIp_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentDateGmt Get Commentses via CommentAuthorAndCommentAuthorIpAndCommentDateGmt
func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentDateGmt(offset int, limit int, CommentAuthor_ string, CommentAuthorIp_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_ip = ? and comment_date_gmt = ?", CommentAuthor_, CommentAuthorIp_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentContent Get Commentses via CommentAuthorAndCommentAuthorIpAndCommentContent
func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentContent(offset int, limit int, CommentAuthor_ string, CommentAuthorIp_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_ip = ? and comment_content = ?", CommentAuthor_, CommentAuthorIp_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentKarma Get Commentses via CommentAuthorAndCommentAuthorIpAndCommentKarma
func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentKarma(offset int, limit int, CommentAuthor_ string, CommentAuthorIp_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_ip = ? and comment_karma = ?", CommentAuthor_, CommentAuthorIp_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentApproved Get Commentses via CommentAuthorAndCommentAuthorIpAndCommentApproved
func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentApproved(offset int, limit int, CommentAuthor_ string, CommentAuthorIp_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_ip = ? and comment_approved = ?", CommentAuthor_, CommentAuthorIp_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentAgent Get Commentses via CommentAuthorAndCommentAuthorIpAndCommentAgent
func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentAgent(offset int, limit int, CommentAuthor_ string, CommentAuthorIp_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_ip = ? and comment_agent = ?", CommentAuthor_, CommentAuthorIp_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentType Get Commentses via CommentAuthorAndCommentAuthorIpAndCommentType
func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentType(offset int, limit int, CommentAuthor_ string, CommentAuthorIp_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_ip = ? and comment_type = ?", CommentAuthor_, CommentAuthorIp_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentParent Get Commentses via CommentAuthorAndCommentAuthorIpAndCommentParent
func GetCommentsesByCommentAuthorAndCommentAuthorIpAndCommentParent(offset int, limit int, CommentAuthor_ string, CommentAuthorIp_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_ip = ? and comment_parent = ?", CommentAuthor_, CommentAuthorIp_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorIpAndUserId Get Commentses via CommentAuthorAndCommentAuthorIpAndUserId
func GetCommentsesByCommentAuthorAndCommentAuthorIpAndUserId(offset int, limit int, CommentAuthor_ string, CommentAuthorIp_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_ip = ? and user_id = ?", CommentAuthor_, CommentAuthorIp_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateAndCommentDateGmt Get Commentses via CommentAuthorAndCommentDateAndCommentDateGmt
func GetCommentsesByCommentAuthorAndCommentDateAndCommentDateGmt(offset int, limit int, CommentAuthor_ string, CommentDate_ time.Time, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date = ? and comment_date_gmt = ?", CommentAuthor_, CommentDate_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateAndCommentContent Get Commentses via CommentAuthorAndCommentDateAndCommentContent
func GetCommentsesByCommentAuthorAndCommentDateAndCommentContent(offset int, limit int, CommentAuthor_ string, CommentDate_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date = ? and comment_content = ?", CommentAuthor_, CommentDate_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateAndCommentKarma Get Commentses via CommentAuthorAndCommentDateAndCommentKarma
func GetCommentsesByCommentAuthorAndCommentDateAndCommentKarma(offset int, limit int, CommentAuthor_ string, CommentDate_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date = ? and comment_karma = ?", CommentAuthor_, CommentDate_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateAndCommentApproved Get Commentses via CommentAuthorAndCommentDateAndCommentApproved
func GetCommentsesByCommentAuthorAndCommentDateAndCommentApproved(offset int, limit int, CommentAuthor_ string, CommentDate_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date = ? and comment_approved = ?", CommentAuthor_, CommentDate_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateAndCommentAgent Get Commentses via CommentAuthorAndCommentDateAndCommentAgent
func GetCommentsesByCommentAuthorAndCommentDateAndCommentAgent(offset int, limit int, CommentAuthor_ string, CommentDate_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date = ? and comment_agent = ?", CommentAuthor_, CommentDate_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateAndCommentType Get Commentses via CommentAuthorAndCommentDateAndCommentType
func GetCommentsesByCommentAuthorAndCommentDateAndCommentType(offset int, limit int, CommentAuthor_ string, CommentDate_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date = ? and comment_type = ?", CommentAuthor_, CommentDate_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateAndCommentParent Get Commentses via CommentAuthorAndCommentDateAndCommentParent
func GetCommentsesByCommentAuthorAndCommentDateAndCommentParent(offset int, limit int, CommentAuthor_ string, CommentDate_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date = ? and comment_parent = ?", CommentAuthor_, CommentDate_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateAndUserId Get Commentses via CommentAuthorAndCommentDateAndUserId
func GetCommentsesByCommentAuthorAndCommentDateAndUserId(offset int, limit int, CommentAuthor_ string, CommentDate_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date = ? and user_id = ?", CommentAuthor_, CommentDate_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentContent Get Commentses via CommentAuthorAndCommentDateGmtAndCommentContent
func GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentContent(offset int, limit int, CommentAuthor_ string, CommentDateGmt_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date_gmt = ? and comment_content = ?", CommentAuthor_, CommentDateGmt_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentKarma Get Commentses via CommentAuthorAndCommentDateGmtAndCommentKarma
func GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentKarma(offset int, limit int, CommentAuthor_ string, CommentDateGmt_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date_gmt = ? and comment_karma = ?", CommentAuthor_, CommentDateGmt_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentApproved Get Commentses via CommentAuthorAndCommentDateGmtAndCommentApproved
func GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentApproved(offset int, limit int, CommentAuthor_ string, CommentDateGmt_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date_gmt = ? and comment_approved = ?", CommentAuthor_, CommentDateGmt_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentAgent Get Commentses via CommentAuthorAndCommentDateGmtAndCommentAgent
func GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentAgent(offset int, limit int, CommentAuthor_ string, CommentDateGmt_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date_gmt = ? and comment_agent = ?", CommentAuthor_, CommentDateGmt_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentType Get Commentses via CommentAuthorAndCommentDateGmtAndCommentType
func GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentType(offset int, limit int, CommentAuthor_ string, CommentDateGmt_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date_gmt = ? and comment_type = ?", CommentAuthor_, CommentDateGmt_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentParent Get Commentses via CommentAuthorAndCommentDateGmtAndCommentParent
func GetCommentsesByCommentAuthorAndCommentDateGmtAndCommentParent(offset int, limit int, CommentAuthor_ string, CommentDateGmt_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date_gmt = ? and comment_parent = ?", CommentAuthor_, CommentDateGmt_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateGmtAndUserId Get Commentses via CommentAuthorAndCommentDateGmtAndUserId
func GetCommentsesByCommentAuthorAndCommentDateGmtAndUserId(offset int, limit int, CommentAuthor_ string, CommentDateGmt_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date_gmt = ? and user_id = ?", CommentAuthor_, CommentDateGmt_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentContentAndCommentKarma Get Commentses via CommentAuthorAndCommentContentAndCommentKarma
func GetCommentsesByCommentAuthorAndCommentContentAndCommentKarma(offset int, limit int, CommentAuthor_ string, CommentContent_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_content = ? and comment_karma = ?", CommentAuthor_, CommentContent_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentContentAndCommentApproved Get Commentses via CommentAuthorAndCommentContentAndCommentApproved
func GetCommentsesByCommentAuthorAndCommentContentAndCommentApproved(offset int, limit int, CommentAuthor_ string, CommentContent_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_content = ? and comment_approved = ?", CommentAuthor_, CommentContent_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentContentAndCommentAgent Get Commentses via CommentAuthorAndCommentContentAndCommentAgent
func GetCommentsesByCommentAuthorAndCommentContentAndCommentAgent(offset int, limit int, CommentAuthor_ string, CommentContent_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_content = ? and comment_agent = ?", CommentAuthor_, CommentContent_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentContentAndCommentType Get Commentses via CommentAuthorAndCommentContentAndCommentType
func GetCommentsesByCommentAuthorAndCommentContentAndCommentType(offset int, limit int, CommentAuthor_ string, CommentContent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_content = ? and comment_type = ?", CommentAuthor_, CommentContent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentContentAndCommentParent Get Commentses via CommentAuthorAndCommentContentAndCommentParent
func GetCommentsesByCommentAuthorAndCommentContentAndCommentParent(offset int, limit int, CommentAuthor_ string, CommentContent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_content = ? and comment_parent = ?", CommentAuthor_, CommentContent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentContentAndUserId Get Commentses via CommentAuthorAndCommentContentAndUserId
func GetCommentsesByCommentAuthorAndCommentContentAndUserId(offset int, limit int, CommentAuthor_ string, CommentContent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_content = ? and user_id = ?", CommentAuthor_, CommentContent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentKarmaAndCommentApproved Get Commentses via CommentAuthorAndCommentKarmaAndCommentApproved
func GetCommentsesByCommentAuthorAndCommentKarmaAndCommentApproved(offset int, limit int, CommentAuthor_ string, CommentKarma_ int, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_karma = ? and comment_approved = ?", CommentAuthor_, CommentKarma_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentKarmaAndCommentAgent Get Commentses via CommentAuthorAndCommentKarmaAndCommentAgent
func GetCommentsesByCommentAuthorAndCommentKarmaAndCommentAgent(offset int, limit int, CommentAuthor_ string, CommentKarma_ int, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_karma = ? and comment_agent = ?", CommentAuthor_, CommentKarma_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentKarmaAndCommentType Get Commentses via CommentAuthorAndCommentKarmaAndCommentType
func GetCommentsesByCommentAuthorAndCommentKarmaAndCommentType(offset int, limit int, CommentAuthor_ string, CommentKarma_ int, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_karma = ? and comment_type = ?", CommentAuthor_, CommentKarma_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentKarmaAndCommentParent Get Commentses via CommentAuthorAndCommentKarmaAndCommentParent
func GetCommentsesByCommentAuthorAndCommentKarmaAndCommentParent(offset int, limit int, CommentAuthor_ string, CommentKarma_ int, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_karma = ? and comment_parent = ?", CommentAuthor_, CommentKarma_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentKarmaAndUserId Get Commentses via CommentAuthorAndCommentKarmaAndUserId
func GetCommentsesByCommentAuthorAndCommentKarmaAndUserId(offset int, limit int, CommentAuthor_ string, CommentKarma_ int, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_karma = ? and user_id = ?", CommentAuthor_, CommentKarma_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentApprovedAndCommentAgent Get Commentses via CommentAuthorAndCommentApprovedAndCommentAgent
func GetCommentsesByCommentAuthorAndCommentApprovedAndCommentAgent(offset int, limit int, CommentAuthor_ string, CommentApproved_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_approved = ? and comment_agent = ?", CommentAuthor_, CommentApproved_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentApprovedAndCommentType Get Commentses via CommentAuthorAndCommentApprovedAndCommentType
func GetCommentsesByCommentAuthorAndCommentApprovedAndCommentType(offset int, limit int, CommentAuthor_ string, CommentApproved_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_approved = ? and comment_type = ?", CommentAuthor_, CommentApproved_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentApprovedAndCommentParent Get Commentses via CommentAuthorAndCommentApprovedAndCommentParent
func GetCommentsesByCommentAuthorAndCommentApprovedAndCommentParent(offset int, limit int, CommentAuthor_ string, CommentApproved_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_approved = ? and comment_parent = ?", CommentAuthor_, CommentApproved_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentApprovedAndUserId Get Commentses via CommentAuthorAndCommentApprovedAndUserId
func GetCommentsesByCommentAuthorAndCommentApprovedAndUserId(offset int, limit int, CommentAuthor_ string, CommentApproved_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_approved = ? and user_id = ?", CommentAuthor_, CommentApproved_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAgentAndCommentType Get Commentses via CommentAuthorAndCommentAgentAndCommentType
func GetCommentsesByCommentAuthorAndCommentAgentAndCommentType(offset int, limit int, CommentAuthor_ string, CommentAgent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_agent = ? and comment_type = ?", CommentAuthor_, CommentAgent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAgentAndCommentParent Get Commentses via CommentAuthorAndCommentAgentAndCommentParent
func GetCommentsesByCommentAuthorAndCommentAgentAndCommentParent(offset int, limit int, CommentAuthor_ string, CommentAgent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_agent = ? and comment_parent = ?", CommentAuthor_, CommentAgent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAgentAndUserId Get Commentses via CommentAuthorAndCommentAgentAndUserId
func GetCommentsesByCommentAuthorAndCommentAgentAndUserId(offset int, limit int, CommentAuthor_ string, CommentAgent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_agent = ? and user_id = ?", CommentAuthor_, CommentAgent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentTypeAndCommentParent Get Commentses via CommentAuthorAndCommentTypeAndCommentParent
func GetCommentsesByCommentAuthorAndCommentTypeAndCommentParent(offset int, limit int, CommentAuthor_ string, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_type = ? and comment_parent = ?", CommentAuthor_, CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentTypeAndUserId Get Commentses via CommentAuthorAndCommentTypeAndUserId
func GetCommentsesByCommentAuthorAndCommentTypeAndUserId(offset int, limit int, CommentAuthor_ string, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_type = ? and user_id = ?", CommentAuthor_, CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentParentAndUserId Get Commentses via CommentAuthorAndCommentParentAndUserId
func GetCommentsesByCommentAuthorAndCommentParentAndUserId(offset int, limit int, CommentAuthor_ string, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_parent = ? and user_id = ?", CommentAuthor_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentAuthorIp Get Commentses via CommentAuthorEmailAndCommentAuthorUrlAndCommentAuthorIp
func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentAuthorIp(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorUrl_ string, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_url = ? and comment_author_ip = ?", CommentAuthorEmail_, CommentAuthorUrl_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentDate Get Commentses via CommentAuthorEmailAndCommentAuthorUrlAndCommentDate
func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentDate(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorUrl_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_url = ? and comment_date = ?", CommentAuthorEmail_, CommentAuthorUrl_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentDateGmt Get Commentses via CommentAuthorEmailAndCommentAuthorUrlAndCommentDateGmt
func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentDateGmt(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorUrl_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_url = ? and comment_date_gmt = ?", CommentAuthorEmail_, CommentAuthorUrl_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentContent Get Commentses via CommentAuthorEmailAndCommentAuthorUrlAndCommentContent
func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentContent(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorUrl_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_url = ? and comment_content = ?", CommentAuthorEmail_, CommentAuthorUrl_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentKarma Get Commentses via CommentAuthorEmailAndCommentAuthorUrlAndCommentKarma
func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentKarma(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorUrl_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_url = ? and comment_karma = ?", CommentAuthorEmail_, CommentAuthorUrl_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentApproved Get Commentses via CommentAuthorEmailAndCommentAuthorUrlAndCommentApproved
func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentApproved(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorUrl_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_url = ? and comment_approved = ?", CommentAuthorEmail_, CommentAuthorUrl_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentAgent Get Commentses via CommentAuthorEmailAndCommentAuthorUrlAndCommentAgent
func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentAgent(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorUrl_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_url = ? and comment_agent = ?", CommentAuthorEmail_, CommentAuthorUrl_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentType Get Commentses via CommentAuthorEmailAndCommentAuthorUrlAndCommentType
func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentType(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorUrl_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_url = ? and comment_type = ?", CommentAuthorEmail_, CommentAuthorUrl_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentParent Get Commentses via CommentAuthorEmailAndCommentAuthorUrlAndCommentParent
func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndCommentParent(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorUrl_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_url = ? and comment_parent = ?", CommentAuthorEmail_, CommentAuthorUrl_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndUserId Get Commentses via CommentAuthorEmailAndCommentAuthorUrlAndUserId
func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrlAndUserId(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorUrl_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_url = ? and user_id = ?", CommentAuthorEmail_, CommentAuthorUrl_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentDate Get Commentses via CommentAuthorEmailAndCommentAuthorIpAndCommentDate
func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentDate(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorIp_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_ip = ? and comment_date = ?", CommentAuthorEmail_, CommentAuthorIp_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentDateGmt Get Commentses via CommentAuthorEmailAndCommentAuthorIpAndCommentDateGmt
func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentDateGmt(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorIp_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_ip = ? and comment_date_gmt = ?", CommentAuthorEmail_, CommentAuthorIp_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentContent Get Commentses via CommentAuthorEmailAndCommentAuthorIpAndCommentContent
func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentContent(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorIp_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_ip = ? and comment_content = ?", CommentAuthorEmail_, CommentAuthorIp_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentKarma Get Commentses via CommentAuthorEmailAndCommentAuthorIpAndCommentKarma
func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentKarma(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorIp_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_ip = ? and comment_karma = ?", CommentAuthorEmail_, CommentAuthorIp_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentApproved Get Commentses via CommentAuthorEmailAndCommentAuthorIpAndCommentApproved
func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentApproved(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorIp_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_ip = ? and comment_approved = ?", CommentAuthorEmail_, CommentAuthorIp_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentAgent Get Commentses via CommentAuthorEmailAndCommentAuthorIpAndCommentAgent
func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentAgent(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorIp_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_ip = ? and comment_agent = ?", CommentAuthorEmail_, CommentAuthorIp_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentType Get Commentses via CommentAuthorEmailAndCommentAuthorIpAndCommentType
func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentType(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorIp_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_ip = ? and comment_type = ?", CommentAuthorEmail_, CommentAuthorIp_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentParent Get Commentses via CommentAuthorEmailAndCommentAuthorIpAndCommentParent
func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndCommentParent(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorIp_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_ip = ? and comment_parent = ?", CommentAuthorEmail_, CommentAuthorIp_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndUserId Get Commentses via CommentAuthorEmailAndCommentAuthorIpAndUserId
func GetCommentsesByCommentAuthorEmailAndCommentAuthorIpAndUserId(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorIp_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_ip = ? and user_id = ?", CommentAuthorEmail_, CommentAuthorIp_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentDateGmt Get Commentses via CommentAuthorEmailAndCommentDateAndCommentDateGmt
func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentDateGmt(offset int, limit int, CommentAuthorEmail_ string, CommentDate_ time.Time, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date = ? and comment_date_gmt = ?", CommentAuthorEmail_, CommentDate_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentContent Get Commentses via CommentAuthorEmailAndCommentDateAndCommentContent
func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentContent(offset int, limit int, CommentAuthorEmail_ string, CommentDate_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date = ? and comment_content = ?", CommentAuthorEmail_, CommentDate_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentKarma Get Commentses via CommentAuthorEmailAndCommentDateAndCommentKarma
func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentKarma(offset int, limit int, CommentAuthorEmail_ string, CommentDate_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date = ? and comment_karma = ?", CommentAuthorEmail_, CommentDate_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentApproved Get Commentses via CommentAuthorEmailAndCommentDateAndCommentApproved
func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentApproved(offset int, limit int, CommentAuthorEmail_ string, CommentDate_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date = ? and comment_approved = ?", CommentAuthorEmail_, CommentDate_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentAgent Get Commentses via CommentAuthorEmailAndCommentDateAndCommentAgent
func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentAgent(offset int, limit int, CommentAuthorEmail_ string, CommentDate_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date = ? and comment_agent = ?", CommentAuthorEmail_, CommentDate_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentType Get Commentses via CommentAuthorEmailAndCommentDateAndCommentType
func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentType(offset int, limit int, CommentAuthorEmail_ string, CommentDate_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date = ? and comment_type = ?", CommentAuthorEmail_, CommentDate_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentParent Get Commentses via CommentAuthorEmailAndCommentDateAndCommentParent
func GetCommentsesByCommentAuthorEmailAndCommentDateAndCommentParent(offset int, limit int, CommentAuthorEmail_ string, CommentDate_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date = ? and comment_parent = ?", CommentAuthorEmail_, CommentDate_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateAndUserId Get Commentses via CommentAuthorEmailAndCommentDateAndUserId
func GetCommentsesByCommentAuthorEmailAndCommentDateAndUserId(offset int, limit int, CommentAuthorEmail_ string, CommentDate_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date = ? and user_id = ?", CommentAuthorEmail_, CommentDate_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentContent Get Commentses via CommentAuthorEmailAndCommentDateGmtAndCommentContent
func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentContent(offset int, limit int, CommentAuthorEmail_ string, CommentDateGmt_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date_gmt = ? and comment_content = ?", CommentAuthorEmail_, CommentDateGmt_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentKarma Get Commentses via CommentAuthorEmailAndCommentDateGmtAndCommentKarma
func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentKarma(offset int, limit int, CommentAuthorEmail_ string, CommentDateGmt_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date_gmt = ? and comment_karma = ?", CommentAuthorEmail_, CommentDateGmt_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentApproved Get Commentses via CommentAuthorEmailAndCommentDateGmtAndCommentApproved
func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentApproved(offset int, limit int, CommentAuthorEmail_ string, CommentDateGmt_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date_gmt = ? and comment_approved = ?", CommentAuthorEmail_, CommentDateGmt_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentAgent Get Commentses via CommentAuthorEmailAndCommentDateGmtAndCommentAgent
func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentAgent(offset int, limit int, CommentAuthorEmail_ string, CommentDateGmt_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date_gmt = ? and comment_agent = ?", CommentAuthorEmail_, CommentDateGmt_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentType Get Commentses via CommentAuthorEmailAndCommentDateGmtAndCommentType
func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentType(offset int, limit int, CommentAuthorEmail_ string, CommentDateGmt_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date_gmt = ? and comment_type = ?", CommentAuthorEmail_, CommentDateGmt_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentParent Get Commentses via CommentAuthorEmailAndCommentDateGmtAndCommentParent
func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndCommentParent(offset int, limit int, CommentAuthorEmail_ string, CommentDateGmt_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date_gmt = ? and comment_parent = ?", CommentAuthorEmail_, CommentDateGmt_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndUserId Get Commentses via CommentAuthorEmailAndCommentDateGmtAndUserId
func GetCommentsesByCommentAuthorEmailAndCommentDateGmtAndUserId(offset int, limit int, CommentAuthorEmail_ string, CommentDateGmt_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date_gmt = ? and user_id = ?", CommentAuthorEmail_, CommentDateGmt_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentKarma Get Commentses via CommentAuthorEmailAndCommentContentAndCommentKarma
func GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentKarma(offset int, limit int, CommentAuthorEmail_ string, CommentContent_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_content = ? and comment_karma = ?", CommentAuthorEmail_, CommentContent_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentApproved Get Commentses via CommentAuthorEmailAndCommentContentAndCommentApproved
func GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentApproved(offset int, limit int, CommentAuthorEmail_ string, CommentContent_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_content = ? and comment_approved = ?", CommentAuthorEmail_, CommentContent_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentAgent Get Commentses via CommentAuthorEmailAndCommentContentAndCommentAgent
func GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentAgent(offset int, limit int, CommentAuthorEmail_ string, CommentContent_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_content = ? and comment_agent = ?", CommentAuthorEmail_, CommentContent_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentType Get Commentses via CommentAuthorEmailAndCommentContentAndCommentType
func GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentType(offset int, limit int, CommentAuthorEmail_ string, CommentContent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_content = ? and comment_type = ?", CommentAuthorEmail_, CommentContent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentParent Get Commentses via CommentAuthorEmailAndCommentContentAndCommentParent
func GetCommentsesByCommentAuthorEmailAndCommentContentAndCommentParent(offset int, limit int, CommentAuthorEmail_ string, CommentContent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_content = ? and comment_parent = ?", CommentAuthorEmail_, CommentContent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentContentAndUserId Get Commentses via CommentAuthorEmailAndCommentContentAndUserId
func GetCommentsesByCommentAuthorEmailAndCommentContentAndUserId(offset int, limit int, CommentAuthorEmail_ string, CommentContent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_content = ? and user_id = ?", CommentAuthorEmail_, CommentContent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentApproved Get Commentses via CommentAuthorEmailAndCommentKarmaAndCommentApproved
func GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentApproved(offset int, limit int, CommentAuthorEmail_ string, CommentKarma_ int, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_karma = ? and comment_approved = ?", CommentAuthorEmail_, CommentKarma_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentAgent Get Commentses via CommentAuthorEmailAndCommentKarmaAndCommentAgent
func GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentAgent(offset int, limit int, CommentAuthorEmail_ string, CommentKarma_ int, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_karma = ? and comment_agent = ?", CommentAuthorEmail_, CommentKarma_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentType Get Commentses via CommentAuthorEmailAndCommentKarmaAndCommentType
func GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentType(offset int, limit int, CommentAuthorEmail_ string, CommentKarma_ int, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_karma = ? and comment_type = ?", CommentAuthorEmail_, CommentKarma_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentParent Get Commentses via CommentAuthorEmailAndCommentKarmaAndCommentParent
func GetCommentsesByCommentAuthorEmailAndCommentKarmaAndCommentParent(offset int, limit int, CommentAuthorEmail_ string, CommentKarma_ int, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_karma = ? and comment_parent = ?", CommentAuthorEmail_, CommentKarma_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentKarmaAndUserId Get Commentses via CommentAuthorEmailAndCommentKarmaAndUserId
func GetCommentsesByCommentAuthorEmailAndCommentKarmaAndUserId(offset int, limit int, CommentAuthorEmail_ string, CommentKarma_ int, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_karma = ? and user_id = ?", CommentAuthorEmail_, CommentKarma_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentAgent Get Commentses via CommentAuthorEmailAndCommentApprovedAndCommentAgent
func GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentAgent(offset int, limit int, CommentAuthorEmail_ string, CommentApproved_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_approved = ? and comment_agent = ?", CommentAuthorEmail_, CommentApproved_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentType Get Commentses via CommentAuthorEmailAndCommentApprovedAndCommentType
func GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentType(offset int, limit int, CommentAuthorEmail_ string, CommentApproved_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_approved = ? and comment_type = ?", CommentAuthorEmail_, CommentApproved_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentParent Get Commentses via CommentAuthorEmailAndCommentApprovedAndCommentParent
func GetCommentsesByCommentAuthorEmailAndCommentApprovedAndCommentParent(offset int, limit int, CommentAuthorEmail_ string, CommentApproved_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_approved = ? and comment_parent = ?", CommentAuthorEmail_, CommentApproved_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentApprovedAndUserId Get Commentses via CommentAuthorEmailAndCommentApprovedAndUserId
func GetCommentsesByCommentAuthorEmailAndCommentApprovedAndUserId(offset int, limit int, CommentAuthorEmail_ string, CommentApproved_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_approved = ? and user_id = ?", CommentAuthorEmail_, CommentApproved_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAgentAndCommentType Get Commentses via CommentAuthorEmailAndCommentAgentAndCommentType
func GetCommentsesByCommentAuthorEmailAndCommentAgentAndCommentType(offset int, limit int, CommentAuthorEmail_ string, CommentAgent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_agent = ? and comment_type = ?", CommentAuthorEmail_, CommentAgent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAgentAndCommentParent Get Commentses via CommentAuthorEmailAndCommentAgentAndCommentParent
func GetCommentsesByCommentAuthorEmailAndCommentAgentAndCommentParent(offset int, limit int, CommentAuthorEmail_ string, CommentAgent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_agent = ? and comment_parent = ?", CommentAuthorEmail_, CommentAgent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAgentAndUserId Get Commentses via CommentAuthorEmailAndCommentAgentAndUserId
func GetCommentsesByCommentAuthorEmailAndCommentAgentAndUserId(offset int, limit int, CommentAuthorEmail_ string, CommentAgent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_agent = ? and user_id = ?", CommentAuthorEmail_, CommentAgent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentTypeAndCommentParent Get Commentses via CommentAuthorEmailAndCommentTypeAndCommentParent
func GetCommentsesByCommentAuthorEmailAndCommentTypeAndCommentParent(offset int, limit int, CommentAuthorEmail_ string, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_type = ? and comment_parent = ?", CommentAuthorEmail_, CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentTypeAndUserId Get Commentses via CommentAuthorEmailAndCommentTypeAndUserId
func GetCommentsesByCommentAuthorEmailAndCommentTypeAndUserId(offset int, limit int, CommentAuthorEmail_ string, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_type = ? and user_id = ?", CommentAuthorEmail_, CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentParentAndUserId Get Commentses via CommentAuthorEmailAndCommentParentAndUserId
func GetCommentsesByCommentAuthorEmailAndCommentParentAndUserId(offset int, limit int, CommentAuthorEmail_ string, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_parent = ? and user_id = ?", CommentAuthorEmail_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentDate Get Commentses via CommentAuthorUrlAndCommentAuthorIpAndCommentDate
func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentDate(offset int, limit int, CommentAuthorUrl_ string, CommentAuthorIp_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_author_ip = ? and comment_date = ?", CommentAuthorUrl_, CommentAuthorIp_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentDateGmt Get Commentses via CommentAuthorUrlAndCommentAuthorIpAndCommentDateGmt
func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentDateGmt(offset int, limit int, CommentAuthorUrl_ string, CommentAuthorIp_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_author_ip = ? and comment_date_gmt = ?", CommentAuthorUrl_, CommentAuthorIp_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentContent Get Commentses via CommentAuthorUrlAndCommentAuthorIpAndCommentContent
func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentContent(offset int, limit int, CommentAuthorUrl_ string, CommentAuthorIp_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_author_ip = ? and comment_content = ?", CommentAuthorUrl_, CommentAuthorIp_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentKarma Get Commentses via CommentAuthorUrlAndCommentAuthorIpAndCommentKarma
func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentKarma(offset int, limit int, CommentAuthorUrl_ string, CommentAuthorIp_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_author_ip = ? and comment_karma = ?", CommentAuthorUrl_, CommentAuthorIp_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentApproved Get Commentses via CommentAuthorUrlAndCommentAuthorIpAndCommentApproved
func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentApproved(offset int, limit int, CommentAuthorUrl_ string, CommentAuthorIp_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_author_ip = ? and comment_approved = ?", CommentAuthorUrl_, CommentAuthorIp_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentAgent Get Commentses via CommentAuthorUrlAndCommentAuthorIpAndCommentAgent
func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentAgent(offset int, limit int, CommentAuthorUrl_ string, CommentAuthorIp_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_author_ip = ? and comment_agent = ?", CommentAuthorUrl_, CommentAuthorIp_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentType Get Commentses via CommentAuthorUrlAndCommentAuthorIpAndCommentType
func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentType(offset int, limit int, CommentAuthorUrl_ string, CommentAuthorIp_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_author_ip = ? and comment_type = ?", CommentAuthorUrl_, CommentAuthorIp_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentParent Get Commentses via CommentAuthorUrlAndCommentAuthorIpAndCommentParent
func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndCommentParent(offset int, limit int, CommentAuthorUrl_ string, CommentAuthorIp_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_author_ip = ? and comment_parent = ?", CommentAuthorUrl_, CommentAuthorIp_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndUserId Get Commentses via CommentAuthorUrlAndCommentAuthorIpAndUserId
func GetCommentsesByCommentAuthorUrlAndCommentAuthorIpAndUserId(offset int, limit int, CommentAuthorUrl_ string, CommentAuthorIp_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_author_ip = ? and user_id = ?", CommentAuthorUrl_, CommentAuthorIp_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentDateGmt Get Commentses via CommentAuthorUrlAndCommentDateAndCommentDateGmt
func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentDateGmt(offset int, limit int, CommentAuthorUrl_ string, CommentDate_ time.Time, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date = ? and comment_date_gmt = ?", CommentAuthorUrl_, CommentDate_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentContent Get Commentses via CommentAuthorUrlAndCommentDateAndCommentContent
func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentContent(offset int, limit int, CommentAuthorUrl_ string, CommentDate_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date = ? and comment_content = ?", CommentAuthorUrl_, CommentDate_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentKarma Get Commentses via CommentAuthorUrlAndCommentDateAndCommentKarma
func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentKarma(offset int, limit int, CommentAuthorUrl_ string, CommentDate_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date = ? and comment_karma = ?", CommentAuthorUrl_, CommentDate_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentApproved Get Commentses via CommentAuthorUrlAndCommentDateAndCommentApproved
func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentApproved(offset int, limit int, CommentAuthorUrl_ string, CommentDate_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date = ? and comment_approved = ?", CommentAuthorUrl_, CommentDate_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentAgent Get Commentses via CommentAuthorUrlAndCommentDateAndCommentAgent
func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentAgent(offset int, limit int, CommentAuthorUrl_ string, CommentDate_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date = ? and comment_agent = ?", CommentAuthorUrl_, CommentDate_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentType Get Commentses via CommentAuthorUrlAndCommentDateAndCommentType
func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentType(offset int, limit int, CommentAuthorUrl_ string, CommentDate_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date = ? and comment_type = ?", CommentAuthorUrl_, CommentDate_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentParent Get Commentses via CommentAuthorUrlAndCommentDateAndCommentParent
func GetCommentsesByCommentAuthorUrlAndCommentDateAndCommentParent(offset int, limit int, CommentAuthorUrl_ string, CommentDate_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date = ? and comment_parent = ?", CommentAuthorUrl_, CommentDate_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateAndUserId Get Commentses via CommentAuthorUrlAndCommentDateAndUserId
func GetCommentsesByCommentAuthorUrlAndCommentDateAndUserId(offset int, limit int, CommentAuthorUrl_ string, CommentDate_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date = ? and user_id = ?", CommentAuthorUrl_, CommentDate_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentContent Get Commentses via CommentAuthorUrlAndCommentDateGmtAndCommentContent
func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentContent(offset int, limit int, CommentAuthorUrl_ string, CommentDateGmt_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date_gmt = ? and comment_content = ?", CommentAuthorUrl_, CommentDateGmt_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentKarma Get Commentses via CommentAuthorUrlAndCommentDateGmtAndCommentKarma
func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentKarma(offset int, limit int, CommentAuthorUrl_ string, CommentDateGmt_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date_gmt = ? and comment_karma = ?", CommentAuthorUrl_, CommentDateGmt_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentApproved Get Commentses via CommentAuthorUrlAndCommentDateGmtAndCommentApproved
func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentApproved(offset int, limit int, CommentAuthorUrl_ string, CommentDateGmt_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date_gmt = ? and comment_approved = ?", CommentAuthorUrl_, CommentDateGmt_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentAgent Get Commentses via CommentAuthorUrlAndCommentDateGmtAndCommentAgent
func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentAgent(offset int, limit int, CommentAuthorUrl_ string, CommentDateGmt_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date_gmt = ? and comment_agent = ?", CommentAuthorUrl_, CommentDateGmt_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentType Get Commentses via CommentAuthorUrlAndCommentDateGmtAndCommentType
func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentType(offset int, limit int, CommentAuthorUrl_ string, CommentDateGmt_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date_gmt = ? and comment_type = ?", CommentAuthorUrl_, CommentDateGmt_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentParent Get Commentses via CommentAuthorUrlAndCommentDateGmtAndCommentParent
func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndCommentParent(offset int, limit int, CommentAuthorUrl_ string, CommentDateGmt_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date_gmt = ? and comment_parent = ?", CommentAuthorUrl_, CommentDateGmt_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndUserId Get Commentses via CommentAuthorUrlAndCommentDateGmtAndUserId
func GetCommentsesByCommentAuthorUrlAndCommentDateGmtAndUserId(offset int, limit int, CommentAuthorUrl_ string, CommentDateGmt_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date_gmt = ? and user_id = ?", CommentAuthorUrl_, CommentDateGmt_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentKarma Get Commentses via CommentAuthorUrlAndCommentContentAndCommentKarma
func GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentKarma(offset int, limit int, CommentAuthorUrl_ string, CommentContent_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_content = ? and comment_karma = ?", CommentAuthorUrl_, CommentContent_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentApproved Get Commentses via CommentAuthorUrlAndCommentContentAndCommentApproved
func GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentApproved(offset int, limit int, CommentAuthorUrl_ string, CommentContent_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_content = ? and comment_approved = ?", CommentAuthorUrl_, CommentContent_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentAgent Get Commentses via CommentAuthorUrlAndCommentContentAndCommentAgent
func GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentAgent(offset int, limit int, CommentAuthorUrl_ string, CommentContent_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_content = ? and comment_agent = ?", CommentAuthorUrl_, CommentContent_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentType Get Commentses via CommentAuthorUrlAndCommentContentAndCommentType
func GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentType(offset int, limit int, CommentAuthorUrl_ string, CommentContent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_content = ? and comment_type = ?", CommentAuthorUrl_, CommentContent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentParent Get Commentses via CommentAuthorUrlAndCommentContentAndCommentParent
func GetCommentsesByCommentAuthorUrlAndCommentContentAndCommentParent(offset int, limit int, CommentAuthorUrl_ string, CommentContent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_content = ? and comment_parent = ?", CommentAuthorUrl_, CommentContent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentContentAndUserId Get Commentses via CommentAuthorUrlAndCommentContentAndUserId
func GetCommentsesByCommentAuthorUrlAndCommentContentAndUserId(offset int, limit int, CommentAuthorUrl_ string, CommentContent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_content = ? and user_id = ?", CommentAuthorUrl_, CommentContent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentApproved Get Commentses via CommentAuthorUrlAndCommentKarmaAndCommentApproved
func GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentApproved(offset int, limit int, CommentAuthorUrl_ string, CommentKarma_ int, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_karma = ? and comment_approved = ?", CommentAuthorUrl_, CommentKarma_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentAgent Get Commentses via CommentAuthorUrlAndCommentKarmaAndCommentAgent
func GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentAgent(offset int, limit int, CommentAuthorUrl_ string, CommentKarma_ int, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_karma = ? and comment_agent = ?", CommentAuthorUrl_, CommentKarma_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentType Get Commentses via CommentAuthorUrlAndCommentKarmaAndCommentType
func GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentType(offset int, limit int, CommentAuthorUrl_ string, CommentKarma_ int, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_karma = ? and comment_type = ?", CommentAuthorUrl_, CommentKarma_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentParent Get Commentses via CommentAuthorUrlAndCommentKarmaAndCommentParent
func GetCommentsesByCommentAuthorUrlAndCommentKarmaAndCommentParent(offset int, limit int, CommentAuthorUrl_ string, CommentKarma_ int, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_karma = ? and comment_parent = ?", CommentAuthorUrl_, CommentKarma_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentKarmaAndUserId Get Commentses via CommentAuthorUrlAndCommentKarmaAndUserId
func GetCommentsesByCommentAuthorUrlAndCommentKarmaAndUserId(offset int, limit int, CommentAuthorUrl_ string, CommentKarma_ int, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_karma = ? and user_id = ?", CommentAuthorUrl_, CommentKarma_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentAgent Get Commentses via CommentAuthorUrlAndCommentApprovedAndCommentAgent
func GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentAgent(offset int, limit int, CommentAuthorUrl_ string, CommentApproved_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_approved = ? and comment_agent = ?", CommentAuthorUrl_, CommentApproved_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentType Get Commentses via CommentAuthorUrlAndCommentApprovedAndCommentType
func GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentType(offset int, limit int, CommentAuthorUrl_ string, CommentApproved_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_approved = ? and comment_type = ?", CommentAuthorUrl_, CommentApproved_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentParent Get Commentses via CommentAuthorUrlAndCommentApprovedAndCommentParent
func GetCommentsesByCommentAuthorUrlAndCommentApprovedAndCommentParent(offset int, limit int, CommentAuthorUrl_ string, CommentApproved_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_approved = ? and comment_parent = ?", CommentAuthorUrl_, CommentApproved_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentApprovedAndUserId Get Commentses via CommentAuthorUrlAndCommentApprovedAndUserId
func GetCommentsesByCommentAuthorUrlAndCommentApprovedAndUserId(offset int, limit int, CommentAuthorUrl_ string, CommentApproved_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_approved = ? and user_id = ?", CommentAuthorUrl_, CommentApproved_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAgentAndCommentType Get Commentses via CommentAuthorUrlAndCommentAgentAndCommentType
func GetCommentsesByCommentAuthorUrlAndCommentAgentAndCommentType(offset int, limit int, CommentAuthorUrl_ string, CommentAgent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_agent = ? and comment_type = ?", CommentAuthorUrl_, CommentAgent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAgentAndCommentParent Get Commentses via CommentAuthorUrlAndCommentAgentAndCommentParent
func GetCommentsesByCommentAuthorUrlAndCommentAgentAndCommentParent(offset int, limit int, CommentAuthorUrl_ string, CommentAgent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_agent = ? and comment_parent = ?", CommentAuthorUrl_, CommentAgent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAgentAndUserId Get Commentses via CommentAuthorUrlAndCommentAgentAndUserId
func GetCommentsesByCommentAuthorUrlAndCommentAgentAndUserId(offset int, limit int, CommentAuthorUrl_ string, CommentAgent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_agent = ? and user_id = ?", CommentAuthorUrl_, CommentAgent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentTypeAndCommentParent Get Commentses via CommentAuthorUrlAndCommentTypeAndCommentParent
func GetCommentsesByCommentAuthorUrlAndCommentTypeAndCommentParent(offset int, limit int, CommentAuthorUrl_ string, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_type = ? and comment_parent = ?", CommentAuthorUrl_, CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentTypeAndUserId Get Commentses via CommentAuthorUrlAndCommentTypeAndUserId
func GetCommentsesByCommentAuthorUrlAndCommentTypeAndUserId(offset int, limit int, CommentAuthorUrl_ string, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_type = ? and user_id = ?", CommentAuthorUrl_, CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentParentAndUserId Get Commentses via CommentAuthorUrlAndCommentParentAndUserId
func GetCommentsesByCommentAuthorUrlAndCommentParentAndUserId(offset int, limit int, CommentAuthorUrl_ string, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_parent = ? and user_id = ?", CommentAuthorUrl_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateAndCommentDateGmt Get Commentses via CommentAuthorIpAndCommentDateAndCommentDateGmt
func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentDateGmt(offset int, limit int, CommentAuthorIp_ string, CommentDate_ time.Time, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date = ? and comment_date_gmt = ?", CommentAuthorIp_, CommentDate_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateAndCommentContent Get Commentses via CommentAuthorIpAndCommentDateAndCommentContent
func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentContent(offset int, limit int, CommentAuthorIp_ string, CommentDate_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date = ? and comment_content = ?", CommentAuthorIp_, CommentDate_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateAndCommentKarma Get Commentses via CommentAuthorIpAndCommentDateAndCommentKarma
func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentKarma(offset int, limit int, CommentAuthorIp_ string, CommentDate_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date = ? and comment_karma = ?", CommentAuthorIp_, CommentDate_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateAndCommentApproved Get Commentses via CommentAuthorIpAndCommentDateAndCommentApproved
func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentApproved(offset int, limit int, CommentAuthorIp_ string, CommentDate_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date = ? and comment_approved = ?", CommentAuthorIp_, CommentDate_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateAndCommentAgent Get Commentses via CommentAuthorIpAndCommentDateAndCommentAgent
func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentAgent(offset int, limit int, CommentAuthorIp_ string, CommentDate_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date = ? and comment_agent = ?", CommentAuthorIp_, CommentDate_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateAndCommentType Get Commentses via CommentAuthorIpAndCommentDateAndCommentType
func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentType(offset int, limit int, CommentAuthorIp_ string, CommentDate_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date = ? and comment_type = ?", CommentAuthorIp_, CommentDate_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateAndCommentParent Get Commentses via CommentAuthorIpAndCommentDateAndCommentParent
func GetCommentsesByCommentAuthorIpAndCommentDateAndCommentParent(offset int, limit int, CommentAuthorIp_ string, CommentDate_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date = ? and comment_parent = ?", CommentAuthorIp_, CommentDate_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateAndUserId Get Commentses via CommentAuthorIpAndCommentDateAndUserId
func GetCommentsesByCommentAuthorIpAndCommentDateAndUserId(offset int, limit int, CommentAuthorIp_ string, CommentDate_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date = ? and user_id = ?", CommentAuthorIp_, CommentDate_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentContent Get Commentses via CommentAuthorIpAndCommentDateGmtAndCommentContent
func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentContent(offset int, limit int, CommentAuthorIp_ string, CommentDateGmt_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date_gmt = ? and comment_content = ?", CommentAuthorIp_, CommentDateGmt_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentKarma Get Commentses via CommentAuthorIpAndCommentDateGmtAndCommentKarma
func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentKarma(offset int, limit int, CommentAuthorIp_ string, CommentDateGmt_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date_gmt = ? and comment_karma = ?", CommentAuthorIp_, CommentDateGmt_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentApproved Get Commentses via CommentAuthorIpAndCommentDateGmtAndCommentApproved
func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentApproved(offset int, limit int, CommentAuthorIp_ string, CommentDateGmt_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date_gmt = ? and comment_approved = ?", CommentAuthorIp_, CommentDateGmt_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentAgent Get Commentses via CommentAuthorIpAndCommentDateGmtAndCommentAgent
func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentAgent(offset int, limit int, CommentAuthorIp_ string, CommentDateGmt_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date_gmt = ? and comment_agent = ?", CommentAuthorIp_, CommentDateGmt_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentType Get Commentses via CommentAuthorIpAndCommentDateGmtAndCommentType
func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentType(offset int, limit int, CommentAuthorIp_ string, CommentDateGmt_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date_gmt = ? and comment_type = ?", CommentAuthorIp_, CommentDateGmt_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentParent Get Commentses via CommentAuthorIpAndCommentDateGmtAndCommentParent
func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndCommentParent(offset int, limit int, CommentAuthorIp_ string, CommentDateGmt_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date_gmt = ? and comment_parent = ?", CommentAuthorIp_, CommentDateGmt_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateGmtAndUserId Get Commentses via CommentAuthorIpAndCommentDateGmtAndUserId
func GetCommentsesByCommentAuthorIpAndCommentDateGmtAndUserId(offset int, limit int, CommentAuthorIp_ string, CommentDateGmt_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date_gmt = ? and user_id = ?", CommentAuthorIp_, CommentDateGmt_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentContentAndCommentKarma Get Commentses via CommentAuthorIpAndCommentContentAndCommentKarma
func GetCommentsesByCommentAuthorIpAndCommentContentAndCommentKarma(offset int, limit int, CommentAuthorIp_ string, CommentContent_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_content = ? and comment_karma = ?", CommentAuthorIp_, CommentContent_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentContentAndCommentApproved Get Commentses via CommentAuthorIpAndCommentContentAndCommentApproved
func GetCommentsesByCommentAuthorIpAndCommentContentAndCommentApproved(offset int, limit int, CommentAuthorIp_ string, CommentContent_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_content = ? and comment_approved = ?", CommentAuthorIp_, CommentContent_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentContentAndCommentAgent Get Commentses via CommentAuthorIpAndCommentContentAndCommentAgent
func GetCommentsesByCommentAuthorIpAndCommentContentAndCommentAgent(offset int, limit int, CommentAuthorIp_ string, CommentContent_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_content = ? and comment_agent = ?", CommentAuthorIp_, CommentContent_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentContentAndCommentType Get Commentses via CommentAuthorIpAndCommentContentAndCommentType
func GetCommentsesByCommentAuthorIpAndCommentContentAndCommentType(offset int, limit int, CommentAuthorIp_ string, CommentContent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_content = ? and comment_type = ?", CommentAuthorIp_, CommentContent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentContentAndCommentParent Get Commentses via CommentAuthorIpAndCommentContentAndCommentParent
func GetCommentsesByCommentAuthorIpAndCommentContentAndCommentParent(offset int, limit int, CommentAuthorIp_ string, CommentContent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_content = ? and comment_parent = ?", CommentAuthorIp_, CommentContent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentContentAndUserId Get Commentses via CommentAuthorIpAndCommentContentAndUserId
func GetCommentsesByCommentAuthorIpAndCommentContentAndUserId(offset int, limit int, CommentAuthorIp_ string, CommentContent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_content = ? and user_id = ?", CommentAuthorIp_, CommentContent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentApproved Get Commentses via CommentAuthorIpAndCommentKarmaAndCommentApproved
func GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentApproved(offset int, limit int, CommentAuthorIp_ string, CommentKarma_ int, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_karma = ? and comment_approved = ?", CommentAuthorIp_, CommentKarma_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentAgent Get Commentses via CommentAuthorIpAndCommentKarmaAndCommentAgent
func GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentAgent(offset int, limit int, CommentAuthorIp_ string, CommentKarma_ int, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_karma = ? and comment_agent = ?", CommentAuthorIp_, CommentKarma_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentType Get Commentses via CommentAuthorIpAndCommentKarmaAndCommentType
func GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentType(offset int, limit int, CommentAuthorIp_ string, CommentKarma_ int, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_karma = ? and comment_type = ?", CommentAuthorIp_, CommentKarma_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentParent Get Commentses via CommentAuthorIpAndCommentKarmaAndCommentParent
func GetCommentsesByCommentAuthorIpAndCommentKarmaAndCommentParent(offset int, limit int, CommentAuthorIp_ string, CommentKarma_ int, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_karma = ? and comment_parent = ?", CommentAuthorIp_, CommentKarma_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentKarmaAndUserId Get Commentses via CommentAuthorIpAndCommentKarmaAndUserId
func GetCommentsesByCommentAuthorIpAndCommentKarmaAndUserId(offset int, limit int, CommentAuthorIp_ string, CommentKarma_ int, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_karma = ? and user_id = ?", CommentAuthorIp_, CommentKarma_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentAgent Get Commentses via CommentAuthorIpAndCommentApprovedAndCommentAgent
func GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentAgent(offset int, limit int, CommentAuthorIp_ string, CommentApproved_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_approved = ? and comment_agent = ?", CommentAuthorIp_, CommentApproved_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentType Get Commentses via CommentAuthorIpAndCommentApprovedAndCommentType
func GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentType(offset int, limit int, CommentAuthorIp_ string, CommentApproved_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_approved = ? and comment_type = ?", CommentAuthorIp_, CommentApproved_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentParent Get Commentses via CommentAuthorIpAndCommentApprovedAndCommentParent
func GetCommentsesByCommentAuthorIpAndCommentApprovedAndCommentParent(offset int, limit int, CommentAuthorIp_ string, CommentApproved_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_approved = ? and comment_parent = ?", CommentAuthorIp_, CommentApproved_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentApprovedAndUserId Get Commentses via CommentAuthorIpAndCommentApprovedAndUserId
func GetCommentsesByCommentAuthorIpAndCommentApprovedAndUserId(offset int, limit int, CommentAuthorIp_ string, CommentApproved_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_approved = ? and user_id = ?", CommentAuthorIp_, CommentApproved_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentAgentAndCommentType Get Commentses via CommentAuthorIpAndCommentAgentAndCommentType
func GetCommentsesByCommentAuthorIpAndCommentAgentAndCommentType(offset int, limit int, CommentAuthorIp_ string, CommentAgent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_agent = ? and comment_type = ?", CommentAuthorIp_, CommentAgent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentAgentAndCommentParent Get Commentses via CommentAuthorIpAndCommentAgentAndCommentParent
func GetCommentsesByCommentAuthorIpAndCommentAgentAndCommentParent(offset int, limit int, CommentAuthorIp_ string, CommentAgent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_agent = ? and comment_parent = ?", CommentAuthorIp_, CommentAgent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentAgentAndUserId Get Commentses via CommentAuthorIpAndCommentAgentAndUserId
func GetCommentsesByCommentAuthorIpAndCommentAgentAndUserId(offset int, limit int, CommentAuthorIp_ string, CommentAgent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_agent = ? and user_id = ?", CommentAuthorIp_, CommentAgent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentTypeAndCommentParent Get Commentses via CommentAuthorIpAndCommentTypeAndCommentParent
func GetCommentsesByCommentAuthorIpAndCommentTypeAndCommentParent(offset int, limit int, CommentAuthorIp_ string, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_type = ? and comment_parent = ?", CommentAuthorIp_, CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentTypeAndUserId Get Commentses via CommentAuthorIpAndCommentTypeAndUserId
func GetCommentsesByCommentAuthorIpAndCommentTypeAndUserId(offset int, limit int, CommentAuthorIp_ string, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_type = ? and user_id = ?", CommentAuthorIp_, CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentParentAndUserId Get Commentses via CommentAuthorIpAndCommentParentAndUserId
func GetCommentsesByCommentAuthorIpAndCommentParentAndUserId(offset int, limit int, CommentAuthorIp_ string, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_parent = ? and user_id = ?", CommentAuthorIp_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentDateGmtAndCommentContent Get Commentses via CommentDateAndCommentDateGmtAndCommentContent
func GetCommentsesByCommentDateAndCommentDateGmtAndCommentContent(offset int, limit int, CommentDate_ time.Time, CommentDateGmt_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_date_gmt = ? and comment_content = ?", CommentDate_, CommentDateGmt_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentDateGmtAndCommentKarma Get Commentses via CommentDateAndCommentDateGmtAndCommentKarma
func GetCommentsesByCommentDateAndCommentDateGmtAndCommentKarma(offset int, limit int, CommentDate_ time.Time, CommentDateGmt_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_date_gmt = ? and comment_karma = ?", CommentDate_, CommentDateGmt_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentDateGmtAndCommentApproved Get Commentses via CommentDateAndCommentDateGmtAndCommentApproved
func GetCommentsesByCommentDateAndCommentDateGmtAndCommentApproved(offset int, limit int, CommentDate_ time.Time, CommentDateGmt_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_date_gmt = ? and comment_approved = ?", CommentDate_, CommentDateGmt_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentDateGmtAndCommentAgent Get Commentses via CommentDateAndCommentDateGmtAndCommentAgent
func GetCommentsesByCommentDateAndCommentDateGmtAndCommentAgent(offset int, limit int, CommentDate_ time.Time, CommentDateGmt_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_date_gmt = ? and comment_agent = ?", CommentDate_, CommentDateGmt_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentDateGmtAndCommentType Get Commentses via CommentDateAndCommentDateGmtAndCommentType
func GetCommentsesByCommentDateAndCommentDateGmtAndCommentType(offset int, limit int, CommentDate_ time.Time, CommentDateGmt_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_date_gmt = ? and comment_type = ?", CommentDate_, CommentDateGmt_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentDateGmtAndCommentParent Get Commentses via CommentDateAndCommentDateGmtAndCommentParent
func GetCommentsesByCommentDateAndCommentDateGmtAndCommentParent(offset int, limit int, CommentDate_ time.Time, CommentDateGmt_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_date_gmt = ? and comment_parent = ?", CommentDate_, CommentDateGmt_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentDateGmtAndUserId Get Commentses via CommentDateAndCommentDateGmtAndUserId
func GetCommentsesByCommentDateAndCommentDateGmtAndUserId(offset int, limit int, CommentDate_ time.Time, CommentDateGmt_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_date_gmt = ? and user_id = ?", CommentDate_, CommentDateGmt_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentContentAndCommentKarma Get Commentses via CommentDateAndCommentContentAndCommentKarma
func GetCommentsesByCommentDateAndCommentContentAndCommentKarma(offset int, limit int, CommentDate_ time.Time, CommentContent_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_content = ? and comment_karma = ?", CommentDate_, CommentContent_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentContentAndCommentApproved Get Commentses via CommentDateAndCommentContentAndCommentApproved
func GetCommentsesByCommentDateAndCommentContentAndCommentApproved(offset int, limit int, CommentDate_ time.Time, CommentContent_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_content = ? and comment_approved = ?", CommentDate_, CommentContent_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentContentAndCommentAgent Get Commentses via CommentDateAndCommentContentAndCommentAgent
func GetCommentsesByCommentDateAndCommentContentAndCommentAgent(offset int, limit int, CommentDate_ time.Time, CommentContent_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_content = ? and comment_agent = ?", CommentDate_, CommentContent_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentContentAndCommentType Get Commentses via CommentDateAndCommentContentAndCommentType
func GetCommentsesByCommentDateAndCommentContentAndCommentType(offset int, limit int, CommentDate_ time.Time, CommentContent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_content = ? and comment_type = ?", CommentDate_, CommentContent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentContentAndCommentParent Get Commentses via CommentDateAndCommentContentAndCommentParent
func GetCommentsesByCommentDateAndCommentContentAndCommentParent(offset int, limit int, CommentDate_ time.Time, CommentContent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_content = ? and comment_parent = ?", CommentDate_, CommentContent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentContentAndUserId Get Commentses via CommentDateAndCommentContentAndUserId
func GetCommentsesByCommentDateAndCommentContentAndUserId(offset int, limit int, CommentDate_ time.Time, CommentContent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_content = ? and user_id = ?", CommentDate_, CommentContent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentKarmaAndCommentApproved Get Commentses via CommentDateAndCommentKarmaAndCommentApproved
func GetCommentsesByCommentDateAndCommentKarmaAndCommentApproved(offset int, limit int, CommentDate_ time.Time, CommentKarma_ int, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_karma = ? and comment_approved = ?", CommentDate_, CommentKarma_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentKarmaAndCommentAgent Get Commentses via CommentDateAndCommentKarmaAndCommentAgent
func GetCommentsesByCommentDateAndCommentKarmaAndCommentAgent(offset int, limit int, CommentDate_ time.Time, CommentKarma_ int, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_karma = ? and comment_agent = ?", CommentDate_, CommentKarma_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentKarmaAndCommentType Get Commentses via CommentDateAndCommentKarmaAndCommentType
func GetCommentsesByCommentDateAndCommentKarmaAndCommentType(offset int, limit int, CommentDate_ time.Time, CommentKarma_ int, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_karma = ? and comment_type = ?", CommentDate_, CommentKarma_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentKarmaAndCommentParent Get Commentses via CommentDateAndCommentKarmaAndCommentParent
func GetCommentsesByCommentDateAndCommentKarmaAndCommentParent(offset int, limit int, CommentDate_ time.Time, CommentKarma_ int, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_karma = ? and comment_parent = ?", CommentDate_, CommentKarma_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentKarmaAndUserId Get Commentses via CommentDateAndCommentKarmaAndUserId
func GetCommentsesByCommentDateAndCommentKarmaAndUserId(offset int, limit int, CommentDate_ time.Time, CommentKarma_ int, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_karma = ? and user_id = ?", CommentDate_, CommentKarma_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentApprovedAndCommentAgent Get Commentses via CommentDateAndCommentApprovedAndCommentAgent
func GetCommentsesByCommentDateAndCommentApprovedAndCommentAgent(offset int, limit int, CommentDate_ time.Time, CommentApproved_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_approved = ? and comment_agent = ?", CommentDate_, CommentApproved_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentApprovedAndCommentType Get Commentses via CommentDateAndCommentApprovedAndCommentType
func GetCommentsesByCommentDateAndCommentApprovedAndCommentType(offset int, limit int, CommentDate_ time.Time, CommentApproved_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_approved = ? and comment_type = ?", CommentDate_, CommentApproved_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentApprovedAndCommentParent Get Commentses via CommentDateAndCommentApprovedAndCommentParent
func GetCommentsesByCommentDateAndCommentApprovedAndCommentParent(offset int, limit int, CommentDate_ time.Time, CommentApproved_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_approved = ? and comment_parent = ?", CommentDate_, CommentApproved_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentApprovedAndUserId Get Commentses via CommentDateAndCommentApprovedAndUserId
func GetCommentsesByCommentDateAndCommentApprovedAndUserId(offset int, limit int, CommentDate_ time.Time, CommentApproved_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_approved = ? and user_id = ?", CommentDate_, CommentApproved_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentAgentAndCommentType Get Commentses via CommentDateAndCommentAgentAndCommentType
func GetCommentsesByCommentDateAndCommentAgentAndCommentType(offset int, limit int, CommentDate_ time.Time, CommentAgent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_agent = ? and comment_type = ?", CommentDate_, CommentAgent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentAgentAndCommentParent Get Commentses via CommentDateAndCommentAgentAndCommentParent
func GetCommentsesByCommentDateAndCommentAgentAndCommentParent(offset int, limit int, CommentDate_ time.Time, CommentAgent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_agent = ? and comment_parent = ?", CommentDate_, CommentAgent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentAgentAndUserId Get Commentses via CommentDateAndCommentAgentAndUserId
func GetCommentsesByCommentDateAndCommentAgentAndUserId(offset int, limit int, CommentDate_ time.Time, CommentAgent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_agent = ? and user_id = ?", CommentDate_, CommentAgent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentTypeAndCommentParent Get Commentses via CommentDateAndCommentTypeAndCommentParent
func GetCommentsesByCommentDateAndCommentTypeAndCommentParent(offset int, limit int, CommentDate_ time.Time, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_type = ? and comment_parent = ?", CommentDate_, CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentTypeAndUserId Get Commentses via CommentDateAndCommentTypeAndUserId
func GetCommentsesByCommentDateAndCommentTypeAndUserId(offset int, limit int, CommentDate_ time.Time, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_type = ? and user_id = ?", CommentDate_, CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentParentAndUserId Get Commentses via CommentDateAndCommentParentAndUserId
func GetCommentsesByCommentDateAndCommentParentAndUserId(offset int, limit int, CommentDate_ time.Time, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_parent = ? and user_id = ?", CommentDate_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentContentAndCommentKarma Get Commentses via CommentDateGmtAndCommentContentAndCommentKarma
func GetCommentsesByCommentDateGmtAndCommentContentAndCommentKarma(offset int, limit int, CommentDateGmt_ time.Time, CommentContent_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_content = ? and comment_karma = ?", CommentDateGmt_, CommentContent_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentContentAndCommentApproved Get Commentses via CommentDateGmtAndCommentContentAndCommentApproved
func GetCommentsesByCommentDateGmtAndCommentContentAndCommentApproved(offset int, limit int, CommentDateGmt_ time.Time, CommentContent_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_content = ? and comment_approved = ?", CommentDateGmt_, CommentContent_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentContentAndCommentAgent Get Commentses via CommentDateGmtAndCommentContentAndCommentAgent
func GetCommentsesByCommentDateGmtAndCommentContentAndCommentAgent(offset int, limit int, CommentDateGmt_ time.Time, CommentContent_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_content = ? and comment_agent = ?", CommentDateGmt_, CommentContent_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentContentAndCommentType Get Commentses via CommentDateGmtAndCommentContentAndCommentType
func GetCommentsesByCommentDateGmtAndCommentContentAndCommentType(offset int, limit int, CommentDateGmt_ time.Time, CommentContent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_content = ? and comment_type = ?", CommentDateGmt_, CommentContent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentContentAndCommentParent Get Commentses via CommentDateGmtAndCommentContentAndCommentParent
func GetCommentsesByCommentDateGmtAndCommentContentAndCommentParent(offset int, limit int, CommentDateGmt_ time.Time, CommentContent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_content = ? and comment_parent = ?", CommentDateGmt_, CommentContent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentContentAndUserId Get Commentses via CommentDateGmtAndCommentContentAndUserId
func GetCommentsesByCommentDateGmtAndCommentContentAndUserId(offset int, limit int, CommentDateGmt_ time.Time, CommentContent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_content = ? and user_id = ?", CommentDateGmt_, CommentContent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentApproved Get Commentses via CommentDateGmtAndCommentKarmaAndCommentApproved
func GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentApproved(offset int, limit int, CommentDateGmt_ time.Time, CommentKarma_ int, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_karma = ? and comment_approved = ?", CommentDateGmt_, CommentKarma_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentAgent Get Commentses via CommentDateGmtAndCommentKarmaAndCommentAgent
func GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentAgent(offset int, limit int, CommentDateGmt_ time.Time, CommentKarma_ int, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_karma = ? and comment_agent = ?", CommentDateGmt_, CommentKarma_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentType Get Commentses via CommentDateGmtAndCommentKarmaAndCommentType
func GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentType(offset int, limit int, CommentDateGmt_ time.Time, CommentKarma_ int, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_karma = ? and comment_type = ?", CommentDateGmt_, CommentKarma_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentParent Get Commentses via CommentDateGmtAndCommentKarmaAndCommentParent
func GetCommentsesByCommentDateGmtAndCommentKarmaAndCommentParent(offset int, limit int, CommentDateGmt_ time.Time, CommentKarma_ int, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_karma = ? and comment_parent = ?", CommentDateGmt_, CommentKarma_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentKarmaAndUserId Get Commentses via CommentDateGmtAndCommentKarmaAndUserId
func GetCommentsesByCommentDateGmtAndCommentKarmaAndUserId(offset int, limit int, CommentDateGmt_ time.Time, CommentKarma_ int, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_karma = ? and user_id = ?", CommentDateGmt_, CommentKarma_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentAgent Get Commentses via CommentDateGmtAndCommentApprovedAndCommentAgent
func GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentAgent(offset int, limit int, CommentDateGmt_ time.Time, CommentApproved_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_approved = ? and comment_agent = ?", CommentDateGmt_, CommentApproved_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentType Get Commentses via CommentDateGmtAndCommentApprovedAndCommentType
func GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentType(offset int, limit int, CommentDateGmt_ time.Time, CommentApproved_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_approved = ? and comment_type = ?", CommentDateGmt_, CommentApproved_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentParent Get Commentses via CommentDateGmtAndCommentApprovedAndCommentParent
func GetCommentsesByCommentDateGmtAndCommentApprovedAndCommentParent(offset int, limit int, CommentDateGmt_ time.Time, CommentApproved_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_approved = ? and comment_parent = ?", CommentDateGmt_, CommentApproved_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentApprovedAndUserId Get Commentses via CommentDateGmtAndCommentApprovedAndUserId
func GetCommentsesByCommentDateGmtAndCommentApprovedAndUserId(offset int, limit int, CommentDateGmt_ time.Time, CommentApproved_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_approved = ? and user_id = ?", CommentDateGmt_, CommentApproved_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentAgentAndCommentType Get Commentses via CommentDateGmtAndCommentAgentAndCommentType
func GetCommentsesByCommentDateGmtAndCommentAgentAndCommentType(offset int, limit int, CommentDateGmt_ time.Time, CommentAgent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_agent = ? and comment_type = ?", CommentDateGmt_, CommentAgent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentAgentAndCommentParent Get Commentses via CommentDateGmtAndCommentAgentAndCommentParent
func GetCommentsesByCommentDateGmtAndCommentAgentAndCommentParent(offset int, limit int, CommentDateGmt_ time.Time, CommentAgent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_agent = ? and comment_parent = ?", CommentDateGmt_, CommentAgent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentAgentAndUserId Get Commentses via CommentDateGmtAndCommentAgentAndUserId
func GetCommentsesByCommentDateGmtAndCommentAgentAndUserId(offset int, limit int, CommentDateGmt_ time.Time, CommentAgent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_agent = ? and user_id = ?", CommentDateGmt_, CommentAgent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentTypeAndCommentParent Get Commentses via CommentDateGmtAndCommentTypeAndCommentParent
func GetCommentsesByCommentDateGmtAndCommentTypeAndCommentParent(offset int, limit int, CommentDateGmt_ time.Time, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_type = ? and comment_parent = ?", CommentDateGmt_, CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentTypeAndUserId Get Commentses via CommentDateGmtAndCommentTypeAndUserId
func GetCommentsesByCommentDateGmtAndCommentTypeAndUserId(offset int, limit int, CommentDateGmt_ time.Time, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_type = ? and user_id = ?", CommentDateGmt_, CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentParentAndUserId Get Commentses via CommentDateGmtAndCommentParentAndUserId
func GetCommentsesByCommentDateGmtAndCommentParentAndUserId(offset int, limit int, CommentDateGmt_ time.Time, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_parent = ? and user_id = ?", CommentDateGmt_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentKarmaAndCommentApproved Get Commentses via CommentContentAndCommentKarmaAndCommentApproved
func GetCommentsesByCommentContentAndCommentKarmaAndCommentApproved(offset int, limit int, CommentContent_ string, CommentKarma_ int, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_karma = ? and comment_approved = ?", CommentContent_, CommentKarma_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentKarmaAndCommentAgent Get Commentses via CommentContentAndCommentKarmaAndCommentAgent
func GetCommentsesByCommentContentAndCommentKarmaAndCommentAgent(offset int, limit int, CommentContent_ string, CommentKarma_ int, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_karma = ? and comment_agent = ?", CommentContent_, CommentKarma_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentKarmaAndCommentType Get Commentses via CommentContentAndCommentKarmaAndCommentType
func GetCommentsesByCommentContentAndCommentKarmaAndCommentType(offset int, limit int, CommentContent_ string, CommentKarma_ int, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_karma = ? and comment_type = ?", CommentContent_, CommentKarma_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentKarmaAndCommentParent Get Commentses via CommentContentAndCommentKarmaAndCommentParent
func GetCommentsesByCommentContentAndCommentKarmaAndCommentParent(offset int, limit int, CommentContent_ string, CommentKarma_ int, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_karma = ? and comment_parent = ?", CommentContent_, CommentKarma_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentKarmaAndUserId Get Commentses via CommentContentAndCommentKarmaAndUserId
func GetCommentsesByCommentContentAndCommentKarmaAndUserId(offset int, limit int, CommentContent_ string, CommentKarma_ int, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_karma = ? and user_id = ?", CommentContent_, CommentKarma_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentApprovedAndCommentAgent Get Commentses via CommentContentAndCommentApprovedAndCommentAgent
func GetCommentsesByCommentContentAndCommentApprovedAndCommentAgent(offset int, limit int, CommentContent_ string, CommentApproved_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_approved = ? and comment_agent = ?", CommentContent_, CommentApproved_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentApprovedAndCommentType Get Commentses via CommentContentAndCommentApprovedAndCommentType
func GetCommentsesByCommentContentAndCommentApprovedAndCommentType(offset int, limit int, CommentContent_ string, CommentApproved_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_approved = ? and comment_type = ?", CommentContent_, CommentApproved_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentApprovedAndCommentParent Get Commentses via CommentContentAndCommentApprovedAndCommentParent
func GetCommentsesByCommentContentAndCommentApprovedAndCommentParent(offset int, limit int, CommentContent_ string, CommentApproved_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_approved = ? and comment_parent = ?", CommentContent_, CommentApproved_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentApprovedAndUserId Get Commentses via CommentContentAndCommentApprovedAndUserId
func GetCommentsesByCommentContentAndCommentApprovedAndUserId(offset int, limit int, CommentContent_ string, CommentApproved_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_approved = ? and user_id = ?", CommentContent_, CommentApproved_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentAgentAndCommentType Get Commentses via CommentContentAndCommentAgentAndCommentType
func GetCommentsesByCommentContentAndCommentAgentAndCommentType(offset int, limit int, CommentContent_ string, CommentAgent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_agent = ? and comment_type = ?", CommentContent_, CommentAgent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentAgentAndCommentParent Get Commentses via CommentContentAndCommentAgentAndCommentParent
func GetCommentsesByCommentContentAndCommentAgentAndCommentParent(offset int, limit int, CommentContent_ string, CommentAgent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_agent = ? and comment_parent = ?", CommentContent_, CommentAgent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentAgentAndUserId Get Commentses via CommentContentAndCommentAgentAndUserId
func GetCommentsesByCommentContentAndCommentAgentAndUserId(offset int, limit int, CommentContent_ string, CommentAgent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_agent = ? and user_id = ?", CommentContent_, CommentAgent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentTypeAndCommentParent Get Commentses via CommentContentAndCommentTypeAndCommentParent
func GetCommentsesByCommentContentAndCommentTypeAndCommentParent(offset int, limit int, CommentContent_ string, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_type = ? and comment_parent = ?", CommentContent_, CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentTypeAndUserId Get Commentses via CommentContentAndCommentTypeAndUserId
func GetCommentsesByCommentContentAndCommentTypeAndUserId(offset int, limit int, CommentContent_ string, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_type = ? and user_id = ?", CommentContent_, CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentParentAndUserId Get Commentses via CommentContentAndCommentParentAndUserId
func GetCommentsesByCommentContentAndCommentParentAndUserId(offset int, limit int, CommentContent_ string, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_parent = ? and user_id = ?", CommentContent_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentApprovedAndCommentAgent Get Commentses via CommentKarmaAndCommentApprovedAndCommentAgent
func GetCommentsesByCommentKarmaAndCommentApprovedAndCommentAgent(offset int, limit int, CommentKarma_ int, CommentApproved_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_approved = ? and comment_agent = ?", CommentKarma_, CommentApproved_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentApprovedAndCommentType Get Commentses via CommentKarmaAndCommentApprovedAndCommentType
func GetCommentsesByCommentKarmaAndCommentApprovedAndCommentType(offset int, limit int, CommentKarma_ int, CommentApproved_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_approved = ? and comment_type = ?", CommentKarma_, CommentApproved_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentApprovedAndCommentParent Get Commentses via CommentKarmaAndCommentApprovedAndCommentParent
func GetCommentsesByCommentKarmaAndCommentApprovedAndCommentParent(offset int, limit int, CommentKarma_ int, CommentApproved_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_approved = ? and comment_parent = ?", CommentKarma_, CommentApproved_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentApprovedAndUserId Get Commentses via CommentKarmaAndCommentApprovedAndUserId
func GetCommentsesByCommentKarmaAndCommentApprovedAndUserId(offset int, limit int, CommentKarma_ int, CommentApproved_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_approved = ? and user_id = ?", CommentKarma_, CommentApproved_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentAgentAndCommentType Get Commentses via CommentKarmaAndCommentAgentAndCommentType
func GetCommentsesByCommentKarmaAndCommentAgentAndCommentType(offset int, limit int, CommentKarma_ int, CommentAgent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_agent = ? and comment_type = ?", CommentKarma_, CommentAgent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentAgentAndCommentParent Get Commentses via CommentKarmaAndCommentAgentAndCommentParent
func GetCommentsesByCommentKarmaAndCommentAgentAndCommentParent(offset int, limit int, CommentKarma_ int, CommentAgent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_agent = ? and comment_parent = ?", CommentKarma_, CommentAgent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentAgentAndUserId Get Commentses via CommentKarmaAndCommentAgentAndUserId
func GetCommentsesByCommentKarmaAndCommentAgentAndUserId(offset int, limit int, CommentKarma_ int, CommentAgent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_agent = ? and user_id = ?", CommentKarma_, CommentAgent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentTypeAndCommentParent Get Commentses via CommentKarmaAndCommentTypeAndCommentParent
func GetCommentsesByCommentKarmaAndCommentTypeAndCommentParent(offset int, limit int, CommentKarma_ int, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_type = ? and comment_parent = ?", CommentKarma_, CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentTypeAndUserId Get Commentses via CommentKarmaAndCommentTypeAndUserId
func GetCommentsesByCommentKarmaAndCommentTypeAndUserId(offset int, limit int, CommentKarma_ int, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_type = ? and user_id = ?", CommentKarma_, CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentParentAndUserId Get Commentses via CommentKarmaAndCommentParentAndUserId
func GetCommentsesByCommentKarmaAndCommentParentAndUserId(offset int, limit int, CommentKarma_ int, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_parent = ? and user_id = ?", CommentKarma_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentApprovedAndCommentAgentAndCommentType Get Commentses via CommentApprovedAndCommentAgentAndCommentType
func GetCommentsesByCommentApprovedAndCommentAgentAndCommentType(offset int, limit int, CommentApproved_ string, CommentAgent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_approved = ? and comment_agent = ? and comment_type = ?", CommentApproved_, CommentAgent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentApprovedAndCommentAgentAndCommentParent Get Commentses via CommentApprovedAndCommentAgentAndCommentParent
func GetCommentsesByCommentApprovedAndCommentAgentAndCommentParent(offset int, limit int, CommentApproved_ string, CommentAgent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_approved = ? and comment_agent = ? and comment_parent = ?", CommentApproved_, CommentAgent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentApprovedAndCommentAgentAndUserId Get Commentses via CommentApprovedAndCommentAgentAndUserId
func GetCommentsesByCommentApprovedAndCommentAgentAndUserId(offset int, limit int, CommentApproved_ string, CommentAgent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_approved = ? and comment_agent = ? and user_id = ?", CommentApproved_, CommentAgent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentApprovedAndCommentTypeAndCommentParent Get Commentses via CommentApprovedAndCommentTypeAndCommentParent
func GetCommentsesByCommentApprovedAndCommentTypeAndCommentParent(offset int, limit int, CommentApproved_ string, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_approved = ? and comment_type = ? and comment_parent = ?", CommentApproved_, CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentApprovedAndCommentTypeAndUserId Get Commentses via CommentApprovedAndCommentTypeAndUserId
func GetCommentsesByCommentApprovedAndCommentTypeAndUserId(offset int, limit int, CommentApproved_ string, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_approved = ? and comment_type = ? and user_id = ?", CommentApproved_, CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentApprovedAndCommentParentAndUserId Get Commentses via CommentApprovedAndCommentParentAndUserId
func GetCommentsesByCommentApprovedAndCommentParentAndUserId(offset int, limit int, CommentApproved_ string, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_approved = ? and comment_parent = ? and user_id = ?", CommentApproved_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAgentAndCommentTypeAndCommentParent Get Commentses via CommentAgentAndCommentTypeAndCommentParent
func GetCommentsesByCommentAgentAndCommentTypeAndCommentParent(offset int, limit int, CommentAgent_ string, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_agent = ? and comment_type = ? and comment_parent = ?", CommentAgent_, CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAgentAndCommentTypeAndUserId Get Commentses via CommentAgentAndCommentTypeAndUserId
func GetCommentsesByCommentAgentAndCommentTypeAndUserId(offset int, limit int, CommentAgent_ string, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_agent = ? and comment_type = ? and user_id = ?", CommentAgent_, CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAgentAndCommentParentAndUserId Get Commentses via CommentAgentAndCommentParentAndUserId
func GetCommentsesByCommentAgentAndCommentParentAndUserId(offset int, limit int, CommentAgent_ string, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_agent = ? and comment_parent = ? and user_id = ?", CommentAgent_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentTypeAndCommentParentAndUserId Get Commentses via CommentTypeAndCommentParentAndUserId
func GetCommentsesByCommentTypeAndCommentParentAndUserId(offset int, limit int, CommentType_ string, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_type = ? and comment_parent = ? and user_id = ?", CommentType_, CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentPostId Get Commentses via CommentIdAndCommentPostId
func GetCommentsesByCommentIdAndCommentPostId(offset int, limit int, CommentId_ int64, CommentPostId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_post_id = ?", CommentId_, CommentPostId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthor Get Commentses via CommentIdAndCommentAuthor
func GetCommentsesByCommentIdAndCommentAuthor(offset int, limit int, CommentId_ int64, CommentAuthor_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author = ?", CommentId_, CommentAuthor_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorEmail Get Commentses via CommentIdAndCommentAuthorEmail
func GetCommentsesByCommentIdAndCommentAuthorEmail(offset int, limit int, CommentId_ int64, CommentAuthorEmail_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_email = ?", CommentId_, CommentAuthorEmail_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorUrl Get Commentses via CommentIdAndCommentAuthorUrl
func GetCommentsesByCommentIdAndCommentAuthorUrl(offset int, limit int, CommentId_ int64, CommentAuthorUrl_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_url = ?", CommentId_, CommentAuthorUrl_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAuthorIp Get Commentses via CommentIdAndCommentAuthorIp
func GetCommentsesByCommentIdAndCommentAuthorIp(offset int, limit int, CommentId_ int64, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_author_ip = ?", CommentId_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDate Get Commentses via CommentIdAndCommentDate
func GetCommentsesByCommentIdAndCommentDate(offset int, limit int, CommentId_ int64, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date = ?", CommentId_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentDateGmt Get Commentses via CommentIdAndCommentDateGmt
func GetCommentsesByCommentIdAndCommentDateGmt(offset int, limit int, CommentId_ int64, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_date_gmt = ?", CommentId_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentContent Get Commentses via CommentIdAndCommentContent
func GetCommentsesByCommentIdAndCommentContent(offset int, limit int, CommentId_ int64, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_content = ?", CommentId_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentKarma Get Commentses via CommentIdAndCommentKarma
func GetCommentsesByCommentIdAndCommentKarma(offset int, limit int, CommentId_ int64, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_karma = ?", CommentId_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentApproved Get Commentses via CommentIdAndCommentApproved
func GetCommentsesByCommentIdAndCommentApproved(offset int, limit int, CommentId_ int64, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_approved = ?", CommentId_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentAgent Get Commentses via CommentIdAndCommentAgent
func GetCommentsesByCommentIdAndCommentAgent(offset int, limit int, CommentId_ int64, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_agent = ?", CommentId_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentType Get Commentses via CommentIdAndCommentType
func GetCommentsesByCommentIdAndCommentType(offset int, limit int, CommentId_ int64, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_type = ?", CommentId_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndCommentParent Get Commentses via CommentIdAndCommentParent
func GetCommentsesByCommentIdAndCommentParent(offset int, limit int, CommentId_ int64, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and comment_parent = ?", CommentId_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentIdAndUserId Get Commentses via CommentIdAndUserId
func GetCommentsesByCommentIdAndUserId(offset int, limit int, CommentId_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_id = ? and user_id = ?", CommentId_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthor Get Commentses via CommentPostIdAndCommentAuthor
func GetCommentsesByCommentPostIdAndCommentAuthor(offset int, limit int, CommentPostId_ int64, CommentAuthor_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author = ?", CommentPostId_, CommentAuthor_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorEmail Get Commentses via CommentPostIdAndCommentAuthorEmail
func GetCommentsesByCommentPostIdAndCommentAuthorEmail(offset int, limit int, CommentPostId_ int64, CommentAuthorEmail_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_email = ?", CommentPostId_, CommentAuthorEmail_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorUrl Get Commentses via CommentPostIdAndCommentAuthorUrl
func GetCommentsesByCommentPostIdAndCommentAuthorUrl(offset int, limit int, CommentPostId_ int64, CommentAuthorUrl_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_url = ?", CommentPostId_, CommentAuthorUrl_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAuthorIp Get Commentses via CommentPostIdAndCommentAuthorIp
func GetCommentsesByCommentPostIdAndCommentAuthorIp(offset int, limit int, CommentPostId_ int64, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_author_ip = ?", CommentPostId_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDate Get Commentses via CommentPostIdAndCommentDate
func GetCommentsesByCommentPostIdAndCommentDate(offset int, limit int, CommentPostId_ int64, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date = ?", CommentPostId_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentDateGmt Get Commentses via CommentPostIdAndCommentDateGmt
func GetCommentsesByCommentPostIdAndCommentDateGmt(offset int, limit int, CommentPostId_ int64, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_date_gmt = ?", CommentPostId_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentContent Get Commentses via CommentPostIdAndCommentContent
func GetCommentsesByCommentPostIdAndCommentContent(offset int, limit int, CommentPostId_ int64, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_content = ?", CommentPostId_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentKarma Get Commentses via CommentPostIdAndCommentKarma
func GetCommentsesByCommentPostIdAndCommentKarma(offset int, limit int, CommentPostId_ int64, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_karma = ?", CommentPostId_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentApproved Get Commentses via CommentPostIdAndCommentApproved
func GetCommentsesByCommentPostIdAndCommentApproved(offset int, limit int, CommentPostId_ int64, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_approved = ?", CommentPostId_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentAgent Get Commentses via CommentPostIdAndCommentAgent
func GetCommentsesByCommentPostIdAndCommentAgent(offset int, limit int, CommentPostId_ int64, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_agent = ?", CommentPostId_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentType Get Commentses via CommentPostIdAndCommentType
func GetCommentsesByCommentPostIdAndCommentType(offset int, limit int, CommentPostId_ int64, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_type = ?", CommentPostId_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndCommentParent Get Commentses via CommentPostIdAndCommentParent
func GetCommentsesByCommentPostIdAndCommentParent(offset int, limit int, CommentPostId_ int64, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and comment_parent = ?", CommentPostId_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentPostIdAndUserId Get Commentses via CommentPostIdAndUserId
func GetCommentsesByCommentPostIdAndUserId(offset int, limit int, CommentPostId_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_id = ? and user_id = ?", CommentPostId_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorEmail Get Commentses via CommentAuthorAndCommentAuthorEmail
func GetCommentsesByCommentAuthorAndCommentAuthorEmail(offset int, limit int, CommentAuthor_ string, CommentAuthorEmail_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_email = ?", CommentAuthor_, CommentAuthorEmail_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorUrl Get Commentses via CommentAuthorAndCommentAuthorUrl
func GetCommentsesByCommentAuthorAndCommentAuthorUrl(offset int, limit int, CommentAuthor_ string, CommentAuthorUrl_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_url = ?", CommentAuthor_, CommentAuthorUrl_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAuthorIp Get Commentses via CommentAuthorAndCommentAuthorIp
func GetCommentsesByCommentAuthorAndCommentAuthorIp(offset int, limit int, CommentAuthor_ string, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_author_ip = ?", CommentAuthor_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDate Get Commentses via CommentAuthorAndCommentDate
func GetCommentsesByCommentAuthorAndCommentDate(offset int, limit int, CommentAuthor_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date = ?", CommentAuthor_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentDateGmt Get Commentses via CommentAuthorAndCommentDateGmt
func GetCommentsesByCommentAuthorAndCommentDateGmt(offset int, limit int, CommentAuthor_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_date_gmt = ?", CommentAuthor_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentContent Get Commentses via CommentAuthorAndCommentContent
func GetCommentsesByCommentAuthorAndCommentContent(offset int, limit int, CommentAuthor_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_content = ?", CommentAuthor_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentKarma Get Commentses via CommentAuthorAndCommentKarma
func GetCommentsesByCommentAuthorAndCommentKarma(offset int, limit int, CommentAuthor_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_karma = ?", CommentAuthor_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentApproved Get Commentses via CommentAuthorAndCommentApproved
func GetCommentsesByCommentAuthorAndCommentApproved(offset int, limit int, CommentAuthor_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_approved = ?", CommentAuthor_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentAgent Get Commentses via CommentAuthorAndCommentAgent
func GetCommentsesByCommentAuthorAndCommentAgent(offset int, limit int, CommentAuthor_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_agent = ?", CommentAuthor_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentType Get Commentses via CommentAuthorAndCommentType
func GetCommentsesByCommentAuthorAndCommentType(offset int, limit int, CommentAuthor_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_type = ?", CommentAuthor_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndCommentParent Get Commentses via CommentAuthorAndCommentParent
func GetCommentsesByCommentAuthorAndCommentParent(offset int, limit int, CommentAuthor_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and comment_parent = ?", CommentAuthor_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorAndUserId Get Commentses via CommentAuthorAndUserId
func GetCommentsesByCommentAuthorAndUserId(offset int, limit int, CommentAuthor_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ? and user_id = ?", CommentAuthor_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorUrl Get Commentses via CommentAuthorEmailAndCommentAuthorUrl
func GetCommentsesByCommentAuthorEmailAndCommentAuthorUrl(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorUrl_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_url = ?", CommentAuthorEmail_, CommentAuthorUrl_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAuthorIp Get Commentses via CommentAuthorEmailAndCommentAuthorIp
func GetCommentsesByCommentAuthorEmailAndCommentAuthorIp(offset int, limit int, CommentAuthorEmail_ string, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_author_ip = ?", CommentAuthorEmail_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDate Get Commentses via CommentAuthorEmailAndCommentDate
func GetCommentsesByCommentAuthorEmailAndCommentDate(offset int, limit int, CommentAuthorEmail_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date = ?", CommentAuthorEmail_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentDateGmt Get Commentses via CommentAuthorEmailAndCommentDateGmt
func GetCommentsesByCommentAuthorEmailAndCommentDateGmt(offset int, limit int, CommentAuthorEmail_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_date_gmt = ?", CommentAuthorEmail_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentContent Get Commentses via CommentAuthorEmailAndCommentContent
func GetCommentsesByCommentAuthorEmailAndCommentContent(offset int, limit int, CommentAuthorEmail_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_content = ?", CommentAuthorEmail_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentKarma Get Commentses via CommentAuthorEmailAndCommentKarma
func GetCommentsesByCommentAuthorEmailAndCommentKarma(offset int, limit int, CommentAuthorEmail_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_karma = ?", CommentAuthorEmail_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentApproved Get Commentses via CommentAuthorEmailAndCommentApproved
func GetCommentsesByCommentAuthorEmailAndCommentApproved(offset int, limit int, CommentAuthorEmail_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_approved = ?", CommentAuthorEmail_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentAgent Get Commentses via CommentAuthorEmailAndCommentAgent
func GetCommentsesByCommentAuthorEmailAndCommentAgent(offset int, limit int, CommentAuthorEmail_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_agent = ?", CommentAuthorEmail_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentType Get Commentses via CommentAuthorEmailAndCommentType
func GetCommentsesByCommentAuthorEmailAndCommentType(offset int, limit int, CommentAuthorEmail_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_type = ?", CommentAuthorEmail_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndCommentParent Get Commentses via CommentAuthorEmailAndCommentParent
func GetCommentsesByCommentAuthorEmailAndCommentParent(offset int, limit int, CommentAuthorEmail_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and comment_parent = ?", CommentAuthorEmail_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorEmailAndUserId Get Commentses via CommentAuthorEmailAndUserId
func GetCommentsesByCommentAuthorEmailAndUserId(offset int, limit int, CommentAuthorEmail_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ? and user_id = ?", CommentAuthorEmail_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAuthorIp Get Commentses via CommentAuthorUrlAndCommentAuthorIp
func GetCommentsesByCommentAuthorUrlAndCommentAuthorIp(offset int, limit int, CommentAuthorUrl_ string, CommentAuthorIp_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_author_ip = ?", CommentAuthorUrl_, CommentAuthorIp_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDate Get Commentses via CommentAuthorUrlAndCommentDate
func GetCommentsesByCommentAuthorUrlAndCommentDate(offset int, limit int, CommentAuthorUrl_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date = ?", CommentAuthorUrl_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentDateGmt Get Commentses via CommentAuthorUrlAndCommentDateGmt
func GetCommentsesByCommentAuthorUrlAndCommentDateGmt(offset int, limit int, CommentAuthorUrl_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_date_gmt = ?", CommentAuthorUrl_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentContent Get Commentses via CommentAuthorUrlAndCommentContent
func GetCommentsesByCommentAuthorUrlAndCommentContent(offset int, limit int, CommentAuthorUrl_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_content = ?", CommentAuthorUrl_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentKarma Get Commentses via CommentAuthorUrlAndCommentKarma
func GetCommentsesByCommentAuthorUrlAndCommentKarma(offset int, limit int, CommentAuthorUrl_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_karma = ?", CommentAuthorUrl_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentApproved Get Commentses via CommentAuthorUrlAndCommentApproved
func GetCommentsesByCommentAuthorUrlAndCommentApproved(offset int, limit int, CommentAuthorUrl_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_approved = ?", CommentAuthorUrl_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentAgent Get Commentses via CommentAuthorUrlAndCommentAgent
func GetCommentsesByCommentAuthorUrlAndCommentAgent(offset int, limit int, CommentAuthorUrl_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_agent = ?", CommentAuthorUrl_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentType Get Commentses via CommentAuthorUrlAndCommentType
func GetCommentsesByCommentAuthorUrlAndCommentType(offset int, limit int, CommentAuthorUrl_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_type = ?", CommentAuthorUrl_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndCommentParent Get Commentses via CommentAuthorUrlAndCommentParent
func GetCommentsesByCommentAuthorUrlAndCommentParent(offset int, limit int, CommentAuthorUrl_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and comment_parent = ?", CommentAuthorUrl_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorUrlAndUserId Get Commentses via CommentAuthorUrlAndUserId
func GetCommentsesByCommentAuthorUrlAndUserId(offset int, limit int, CommentAuthorUrl_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ? and user_id = ?", CommentAuthorUrl_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDate Get Commentses via CommentAuthorIpAndCommentDate
func GetCommentsesByCommentAuthorIpAndCommentDate(offset int, limit int, CommentAuthorIp_ string, CommentDate_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date = ?", CommentAuthorIp_, CommentDate_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentDateGmt Get Commentses via CommentAuthorIpAndCommentDateGmt
func GetCommentsesByCommentAuthorIpAndCommentDateGmt(offset int, limit int, CommentAuthorIp_ string, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_date_gmt = ?", CommentAuthorIp_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentContent Get Commentses via CommentAuthorIpAndCommentContent
func GetCommentsesByCommentAuthorIpAndCommentContent(offset int, limit int, CommentAuthorIp_ string, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_content = ?", CommentAuthorIp_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentKarma Get Commentses via CommentAuthorIpAndCommentKarma
func GetCommentsesByCommentAuthorIpAndCommentKarma(offset int, limit int, CommentAuthorIp_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_karma = ?", CommentAuthorIp_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentApproved Get Commentses via CommentAuthorIpAndCommentApproved
func GetCommentsesByCommentAuthorIpAndCommentApproved(offset int, limit int, CommentAuthorIp_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_approved = ?", CommentAuthorIp_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentAgent Get Commentses via CommentAuthorIpAndCommentAgent
func GetCommentsesByCommentAuthorIpAndCommentAgent(offset int, limit int, CommentAuthorIp_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_agent = ?", CommentAuthorIp_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentType Get Commentses via CommentAuthorIpAndCommentType
func GetCommentsesByCommentAuthorIpAndCommentType(offset int, limit int, CommentAuthorIp_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_type = ?", CommentAuthorIp_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndCommentParent Get Commentses via CommentAuthorIpAndCommentParent
func GetCommentsesByCommentAuthorIpAndCommentParent(offset int, limit int, CommentAuthorIp_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and comment_parent = ?", CommentAuthorIp_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAuthorIpAndUserId Get Commentses via CommentAuthorIpAndUserId
func GetCommentsesByCommentAuthorIpAndUserId(offset int, limit int, CommentAuthorIp_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_ip = ? and user_id = ?", CommentAuthorIp_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentDateGmt Get Commentses via CommentDateAndCommentDateGmt
func GetCommentsesByCommentDateAndCommentDateGmt(offset int, limit int, CommentDate_ time.Time, CommentDateGmt_ time.Time) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_date_gmt = ?", CommentDate_, CommentDateGmt_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentContent Get Commentses via CommentDateAndCommentContent
func GetCommentsesByCommentDateAndCommentContent(offset int, limit int, CommentDate_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_content = ?", CommentDate_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentKarma Get Commentses via CommentDateAndCommentKarma
func GetCommentsesByCommentDateAndCommentKarma(offset int, limit int, CommentDate_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_karma = ?", CommentDate_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentApproved Get Commentses via CommentDateAndCommentApproved
func GetCommentsesByCommentDateAndCommentApproved(offset int, limit int, CommentDate_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_approved = ?", CommentDate_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentAgent Get Commentses via CommentDateAndCommentAgent
func GetCommentsesByCommentDateAndCommentAgent(offset int, limit int, CommentDate_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_agent = ?", CommentDate_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentType Get Commentses via CommentDateAndCommentType
func GetCommentsesByCommentDateAndCommentType(offset int, limit int, CommentDate_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_type = ?", CommentDate_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndCommentParent Get Commentses via CommentDateAndCommentParent
func GetCommentsesByCommentDateAndCommentParent(offset int, limit int, CommentDate_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and comment_parent = ?", CommentDate_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateAndUserId Get Commentses via CommentDateAndUserId
func GetCommentsesByCommentDateAndUserId(offset int, limit int, CommentDate_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ? and user_id = ?", CommentDate_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentContent Get Commentses via CommentDateGmtAndCommentContent
func GetCommentsesByCommentDateGmtAndCommentContent(offset int, limit int, CommentDateGmt_ time.Time, CommentContent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_content = ?", CommentDateGmt_, CommentContent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentKarma Get Commentses via CommentDateGmtAndCommentKarma
func GetCommentsesByCommentDateGmtAndCommentKarma(offset int, limit int, CommentDateGmt_ time.Time, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_karma = ?", CommentDateGmt_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentApproved Get Commentses via CommentDateGmtAndCommentApproved
func GetCommentsesByCommentDateGmtAndCommentApproved(offset int, limit int, CommentDateGmt_ time.Time, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_approved = ?", CommentDateGmt_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentAgent Get Commentses via CommentDateGmtAndCommentAgent
func GetCommentsesByCommentDateGmtAndCommentAgent(offset int, limit int, CommentDateGmt_ time.Time, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_agent = ?", CommentDateGmt_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentType Get Commentses via CommentDateGmtAndCommentType
func GetCommentsesByCommentDateGmtAndCommentType(offset int, limit int, CommentDateGmt_ time.Time, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_type = ?", CommentDateGmt_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndCommentParent Get Commentses via CommentDateGmtAndCommentParent
func GetCommentsesByCommentDateGmtAndCommentParent(offset int, limit int, CommentDateGmt_ time.Time, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and comment_parent = ?", CommentDateGmt_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentDateGmtAndUserId Get Commentses via CommentDateGmtAndUserId
func GetCommentsesByCommentDateGmtAndUserId(offset int, limit int, CommentDateGmt_ time.Time, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ? and user_id = ?", CommentDateGmt_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentKarma Get Commentses via CommentContentAndCommentKarma
func GetCommentsesByCommentContentAndCommentKarma(offset int, limit int, CommentContent_ string, CommentKarma_ int) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_karma = ?", CommentContent_, CommentKarma_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentApproved Get Commentses via CommentContentAndCommentApproved
func GetCommentsesByCommentContentAndCommentApproved(offset int, limit int, CommentContent_ string, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_approved = ?", CommentContent_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentAgent Get Commentses via CommentContentAndCommentAgent
func GetCommentsesByCommentContentAndCommentAgent(offset int, limit int, CommentContent_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_agent = ?", CommentContent_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentType Get Commentses via CommentContentAndCommentType
func GetCommentsesByCommentContentAndCommentType(offset int, limit int, CommentContent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_type = ?", CommentContent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndCommentParent Get Commentses via CommentContentAndCommentParent
func GetCommentsesByCommentContentAndCommentParent(offset int, limit int, CommentContent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and comment_parent = ?", CommentContent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentContentAndUserId Get Commentses via CommentContentAndUserId
func GetCommentsesByCommentContentAndUserId(offset int, limit int, CommentContent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ? and user_id = ?", CommentContent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentApproved Get Commentses via CommentKarmaAndCommentApproved
func GetCommentsesByCommentKarmaAndCommentApproved(offset int, limit int, CommentKarma_ int, CommentApproved_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_approved = ?", CommentKarma_, CommentApproved_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentAgent Get Commentses via CommentKarmaAndCommentAgent
func GetCommentsesByCommentKarmaAndCommentAgent(offset int, limit int, CommentKarma_ int, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_agent = ?", CommentKarma_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentType Get Commentses via CommentKarmaAndCommentType
func GetCommentsesByCommentKarmaAndCommentType(offset int, limit int, CommentKarma_ int, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_type = ?", CommentKarma_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndCommentParent Get Commentses via CommentKarmaAndCommentParent
func GetCommentsesByCommentKarmaAndCommentParent(offset int, limit int, CommentKarma_ int, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and comment_parent = ?", CommentKarma_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentKarmaAndUserId Get Commentses via CommentKarmaAndUserId
func GetCommentsesByCommentKarmaAndUserId(offset int, limit int, CommentKarma_ int, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ? and user_id = ?", CommentKarma_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentApprovedAndCommentAgent Get Commentses via CommentApprovedAndCommentAgent
func GetCommentsesByCommentApprovedAndCommentAgent(offset int, limit int, CommentApproved_ string, CommentAgent_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_approved = ? and comment_agent = ?", CommentApproved_, CommentAgent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentApprovedAndCommentType Get Commentses via CommentApprovedAndCommentType
func GetCommentsesByCommentApprovedAndCommentType(offset int, limit int, CommentApproved_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_approved = ? and comment_type = ?", CommentApproved_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentApprovedAndCommentParent Get Commentses via CommentApprovedAndCommentParent
func GetCommentsesByCommentApprovedAndCommentParent(offset int, limit int, CommentApproved_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_approved = ? and comment_parent = ?", CommentApproved_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentApprovedAndUserId Get Commentses via CommentApprovedAndUserId
func GetCommentsesByCommentApprovedAndUserId(offset int, limit int, CommentApproved_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_approved = ? and user_id = ?", CommentApproved_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAgentAndCommentType Get Commentses via CommentAgentAndCommentType
func GetCommentsesByCommentAgentAndCommentType(offset int, limit int, CommentAgent_ string, CommentType_ string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_agent = ? and comment_type = ?", CommentAgent_, CommentType_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAgentAndCommentParent Get Commentses via CommentAgentAndCommentParent
func GetCommentsesByCommentAgentAndCommentParent(offset int, limit int, CommentAgent_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_agent = ? and comment_parent = ?", CommentAgent_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentAgentAndUserId Get Commentses via CommentAgentAndUserId
func GetCommentsesByCommentAgentAndUserId(offset int, limit int, CommentAgent_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_agent = ? and user_id = ?", CommentAgent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentTypeAndCommentParent Get Commentses via CommentTypeAndCommentParent
func GetCommentsesByCommentTypeAndCommentParent(offset int, limit int, CommentType_ string, CommentParent_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_type = ? and comment_parent = ?", CommentType_, CommentParent_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentTypeAndUserId Get Commentses via CommentTypeAndUserId
func GetCommentsesByCommentTypeAndUserId(offset int, limit int, CommentType_ string, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_type = ? and user_id = ?", CommentType_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentsesByCommentParentAndUserId Get Commentses via CommentParentAndUserId
func GetCommentsesByCommentParentAndUserId(offset int, limit int, CommentParent_ int64, UserId_ int64) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_parent = ? and user_id = ?", CommentParent_, UserId_).Limit(limit, offset).Find(_Comments)
	return _Comments, err
}

// GetCommentses Get Commentses via field
func GetCommentses(offset int, limit int, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentId Get Commentss via CommentId
func GetCommentsesViaCommentId(offset int, limit int, CommentId_ int64, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_ID = ?", CommentId_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentPostId Get Commentss via CommentPostId
func GetCommentsesViaCommentPostId(offset int, limit int, CommentPostId_ int64, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_post_ID = ?", CommentPostId_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentAuthor Get Commentss via CommentAuthor
func GetCommentsesViaCommentAuthor(offset int, limit int, CommentAuthor_ string, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author = ?", CommentAuthor_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentAuthorEmail Get Commentss via CommentAuthorEmail
func GetCommentsesViaCommentAuthorEmail(offset int, limit int, CommentAuthorEmail_ string, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_email = ?", CommentAuthorEmail_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentAuthorUrl Get Commentss via CommentAuthorUrl
func GetCommentsesViaCommentAuthorUrl(offset int, limit int, CommentAuthorUrl_ string, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_url = ?", CommentAuthorUrl_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentAuthorIp Get Commentss via CommentAuthorIp
func GetCommentsesViaCommentAuthorIp(offset int, limit int, CommentAuthorIp_ string, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_author_IP = ?", CommentAuthorIp_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentDate Get Commentss via CommentDate
func GetCommentsesViaCommentDate(offset int, limit int, CommentDate_ time.Time, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date = ?", CommentDate_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentDateGmt Get Commentss via CommentDateGmt
func GetCommentsesViaCommentDateGmt(offset int, limit int, CommentDateGmt_ time.Time, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_date_gmt = ?", CommentDateGmt_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentContent Get Commentss via CommentContent
func GetCommentsesViaCommentContent(offset int, limit int, CommentContent_ string, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_content = ?", CommentContent_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentKarma Get Commentss via CommentKarma
func GetCommentsesViaCommentKarma(offset int, limit int, CommentKarma_ int, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_karma = ?", CommentKarma_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentApproved Get Commentss via CommentApproved
func GetCommentsesViaCommentApproved(offset int, limit int, CommentApproved_ string, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_approved = ?", CommentApproved_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentAgent Get Commentss via CommentAgent
func GetCommentsesViaCommentAgent(offset int, limit int, CommentAgent_ string, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_agent = ?", CommentAgent_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentType Get Commentss via CommentType
func GetCommentsesViaCommentType(offset int, limit int, CommentType_ string, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_type = ?", CommentType_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaCommentParent Get Commentss via CommentParent
func GetCommentsesViaCommentParent(offset int, limit int, CommentParent_ int64, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("comment_parent = ?", CommentParent_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// GetCommentsesViaUserId Get Commentss via UserId
func GetCommentsesViaUserId(offset int, limit int, UserId_ int64, field string) (*[]*Comments, error) {
	var _Comments = new([]*Comments)
	err := Engine.Table("comments").Where("user_id = ?", UserId_).Limit(limit, offset).Desc(field).Find(_Comments)
	return _Comments, err
}

// HasCommentsViaCommentId Has Comments via CommentId
func HasCommentsViaCommentId(iCommentId int64) bool {
	if has, err := Engine.Where("comment_ID = ?", iCommentId).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentPostId Has Comments via CommentPostId
func HasCommentsViaCommentPostId(iCommentPostId int64) bool {
	if has, err := Engine.Where("comment_post_ID = ?", iCommentPostId).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentAuthor Has Comments via CommentAuthor
func HasCommentsViaCommentAuthor(iCommentAuthor string) bool {
	if has, err := Engine.Where("comment_author = ?", iCommentAuthor).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentAuthorEmail Has Comments via CommentAuthorEmail
func HasCommentsViaCommentAuthorEmail(iCommentAuthorEmail string) bool {
	if has, err := Engine.Where("comment_author_email = ?", iCommentAuthorEmail).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentAuthorUrl Has Comments via CommentAuthorUrl
func HasCommentsViaCommentAuthorUrl(iCommentAuthorUrl string) bool {
	if has, err := Engine.Where("comment_author_url = ?", iCommentAuthorUrl).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentAuthorIp Has Comments via CommentAuthorIp
func HasCommentsViaCommentAuthorIp(iCommentAuthorIp string) bool {
	if has, err := Engine.Where("comment_author_IP = ?", iCommentAuthorIp).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentDate Has Comments via CommentDate
func HasCommentsViaCommentDate(iCommentDate time.Time) bool {
	if has, err := Engine.Where("comment_date = ?", iCommentDate).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentDateGmt Has Comments via CommentDateGmt
func HasCommentsViaCommentDateGmt(iCommentDateGmt time.Time) bool {
	if has, err := Engine.Where("comment_date_gmt = ?", iCommentDateGmt).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentContent Has Comments via CommentContent
func HasCommentsViaCommentContent(iCommentContent string) bool {
	if has, err := Engine.Where("comment_content = ?", iCommentContent).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentKarma Has Comments via CommentKarma
func HasCommentsViaCommentKarma(iCommentKarma int) bool {
	if has, err := Engine.Where("comment_karma = ?", iCommentKarma).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentApproved Has Comments via CommentApproved
func HasCommentsViaCommentApproved(iCommentApproved string) bool {
	if has, err := Engine.Where("comment_approved = ?", iCommentApproved).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentAgent Has Comments via CommentAgent
func HasCommentsViaCommentAgent(iCommentAgent string) bool {
	if has, err := Engine.Where("comment_agent = ?", iCommentAgent).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentType Has Comments via CommentType
func HasCommentsViaCommentType(iCommentType string) bool {
	if has, err := Engine.Where("comment_type = ?", iCommentType).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaCommentParent Has Comments via CommentParent
func HasCommentsViaCommentParent(iCommentParent int64) bool {
	if has, err := Engine.Where("comment_parent = ?", iCommentParent).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasCommentsViaUserId Has Comments via UserId
func HasCommentsViaUserId(iUserId int64) bool {
	if has, err := Engine.Where("user_id = ?", iUserId).Get(new(Comments)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetCommentsViaCommentId Get Comments via CommentId
func GetCommentsViaCommentId(iCommentId int64) (*Comments, error) {
	var _Comments = &Comments{CommentId: iCommentId}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentPostId Get Comments via CommentPostId
func GetCommentsViaCommentPostId(iCommentPostId int64) (*Comments, error) {
	var _Comments = &Comments{CommentPostId: iCommentPostId}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentAuthor Get Comments via CommentAuthor
func GetCommentsViaCommentAuthor(iCommentAuthor string) (*Comments, error) {
	var _Comments = &Comments{CommentAuthor: iCommentAuthor}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentAuthorEmail Get Comments via CommentAuthorEmail
func GetCommentsViaCommentAuthorEmail(iCommentAuthorEmail string) (*Comments, error) {
	var _Comments = &Comments{CommentAuthorEmail: iCommentAuthorEmail}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentAuthorUrl Get Comments via CommentAuthorUrl
func GetCommentsViaCommentAuthorUrl(iCommentAuthorUrl string) (*Comments, error) {
	var _Comments = &Comments{CommentAuthorUrl: iCommentAuthorUrl}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentAuthorIp Get Comments via CommentAuthorIp
func GetCommentsViaCommentAuthorIp(iCommentAuthorIp string) (*Comments, error) {
	var _Comments = &Comments{CommentAuthorIp: iCommentAuthorIp}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentDate Get Comments via CommentDate
func GetCommentsViaCommentDate(iCommentDate time.Time) (*Comments, error) {
	var _Comments = &Comments{CommentDate: iCommentDate}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentDateGmt Get Comments via CommentDateGmt
func GetCommentsViaCommentDateGmt(iCommentDateGmt time.Time) (*Comments, error) {
	var _Comments = &Comments{CommentDateGmt: iCommentDateGmt}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentContent Get Comments via CommentContent
func GetCommentsViaCommentContent(iCommentContent string) (*Comments, error) {
	var _Comments = &Comments{CommentContent: iCommentContent}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentKarma Get Comments via CommentKarma
func GetCommentsViaCommentKarma(iCommentKarma int) (*Comments, error) {
	var _Comments = &Comments{CommentKarma: iCommentKarma}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentApproved Get Comments via CommentApproved
func GetCommentsViaCommentApproved(iCommentApproved string) (*Comments, error) {
	var _Comments = &Comments{CommentApproved: iCommentApproved}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentAgent Get Comments via CommentAgent
func GetCommentsViaCommentAgent(iCommentAgent string) (*Comments, error) {
	var _Comments = &Comments{CommentAgent: iCommentAgent}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentType Get Comments via CommentType
func GetCommentsViaCommentType(iCommentType string) (*Comments, error) {
	var _Comments = &Comments{CommentType: iCommentType}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaCommentParent Get Comments via CommentParent
func GetCommentsViaCommentParent(iCommentParent int64) (*Comments, error) {
	var _Comments = &Comments{CommentParent: iCommentParent}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// GetCommentsViaUserId Get Comments via UserId
func GetCommentsViaUserId(iUserId int64) (*Comments, error) {
	var _Comments = &Comments{UserId: iUserId}
	has, err := Engine.Get(_Comments)
	if has {
		return _Comments, err
	} else {
		return nil, err
	}
}

// SetCommentsViaCommentId Set Comments via CommentId
func SetCommentsViaCommentId(iCommentId int64, comments *Comments) (int64, error) {
	comments.CommentId = iCommentId
	return Engine.Insert(comments)
}

// SetCommentsViaCommentPostId Set Comments via CommentPostId
func SetCommentsViaCommentPostId(iCommentPostId int64, comments *Comments) (int64, error) {
	comments.CommentPostId = iCommentPostId
	return Engine.Insert(comments)
}

// SetCommentsViaCommentAuthor Set Comments via CommentAuthor
func SetCommentsViaCommentAuthor(iCommentAuthor string, comments *Comments) (int64, error) {
	comments.CommentAuthor = iCommentAuthor
	return Engine.Insert(comments)
}

// SetCommentsViaCommentAuthorEmail Set Comments via CommentAuthorEmail
func SetCommentsViaCommentAuthorEmail(iCommentAuthorEmail string, comments *Comments) (int64, error) {
	comments.CommentAuthorEmail = iCommentAuthorEmail
	return Engine.Insert(comments)
}

// SetCommentsViaCommentAuthorUrl Set Comments via CommentAuthorUrl
func SetCommentsViaCommentAuthorUrl(iCommentAuthorUrl string, comments *Comments) (int64, error) {
	comments.CommentAuthorUrl = iCommentAuthorUrl
	return Engine.Insert(comments)
}

// SetCommentsViaCommentAuthorIp Set Comments via CommentAuthorIp
func SetCommentsViaCommentAuthorIp(iCommentAuthorIp string, comments *Comments) (int64, error) {
	comments.CommentAuthorIp = iCommentAuthorIp
	return Engine.Insert(comments)
}

// SetCommentsViaCommentDate Set Comments via CommentDate
func SetCommentsViaCommentDate(iCommentDate time.Time, comments *Comments) (int64, error) {
	comments.CommentDate = iCommentDate
	return Engine.Insert(comments)
}

// SetCommentsViaCommentDateGmt Set Comments via CommentDateGmt
func SetCommentsViaCommentDateGmt(iCommentDateGmt time.Time, comments *Comments) (int64, error) {
	comments.CommentDateGmt = iCommentDateGmt
	return Engine.Insert(comments)
}

// SetCommentsViaCommentContent Set Comments via CommentContent
func SetCommentsViaCommentContent(iCommentContent string, comments *Comments) (int64, error) {
	comments.CommentContent = iCommentContent
	return Engine.Insert(comments)
}

// SetCommentsViaCommentKarma Set Comments via CommentKarma
func SetCommentsViaCommentKarma(iCommentKarma int, comments *Comments) (int64, error) {
	comments.CommentKarma = iCommentKarma
	return Engine.Insert(comments)
}

// SetCommentsViaCommentApproved Set Comments via CommentApproved
func SetCommentsViaCommentApproved(iCommentApproved string, comments *Comments) (int64, error) {
	comments.CommentApproved = iCommentApproved
	return Engine.Insert(comments)
}

// SetCommentsViaCommentAgent Set Comments via CommentAgent
func SetCommentsViaCommentAgent(iCommentAgent string, comments *Comments) (int64, error) {
	comments.CommentAgent = iCommentAgent
	return Engine.Insert(comments)
}

// SetCommentsViaCommentType Set Comments via CommentType
func SetCommentsViaCommentType(iCommentType string, comments *Comments) (int64, error) {
	comments.CommentType = iCommentType
	return Engine.Insert(comments)
}

// SetCommentsViaCommentParent Set Comments via CommentParent
func SetCommentsViaCommentParent(iCommentParent int64, comments *Comments) (int64, error) {
	comments.CommentParent = iCommentParent
	return Engine.Insert(comments)
}

// SetCommentsViaUserId Set Comments via UserId
func SetCommentsViaUserId(iUserId int64, comments *Comments) (int64, error) {
	comments.UserId = iUserId
	return Engine.Insert(comments)
}

// AddComments Add Comments via all columns
func AddComments(iCommentId int64, iCommentPostId int64, iCommentAuthor string, iCommentAuthorEmail string, iCommentAuthorUrl string, iCommentAuthorIp string, iCommentDate time.Time, iCommentDateGmt time.Time, iCommentContent string, iCommentKarma int, iCommentApproved string, iCommentAgent string, iCommentType string, iCommentParent int64, iUserId int64) error {
	_Comments := &Comments{CommentId: iCommentId, CommentPostId: iCommentPostId, CommentAuthor: iCommentAuthor, CommentAuthorEmail: iCommentAuthorEmail, CommentAuthorUrl: iCommentAuthorUrl, CommentAuthorIp: iCommentAuthorIp, CommentDate: iCommentDate, CommentDateGmt: iCommentDateGmt, CommentContent: iCommentContent, CommentKarma: iCommentKarma, CommentApproved: iCommentApproved, CommentAgent: iCommentAgent, CommentType: iCommentType, CommentParent: iCommentParent, UserId: iUserId}
	if _, err := Engine.Insert(_Comments); err != nil {
		return err
	} else {
		return nil
	}
}

// PostComments Post Comments via iComments
func PostComments(iComments *Comments) (int64, error) {
	_, err := Engine.Insert(iComments)
	return iComments.CommentId, err
}

// PutComments Put Comments
func PutComments(iComments *Comments) (int64, error) {
	_, err := Engine.Id(iComments.CommentId).Update(iComments)
	return iComments.CommentId, err
}

// PutCommentsViaCommentId Put Comments via CommentId
func PutCommentsViaCommentId(CommentId_ int64, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentId: CommentId_})
	return row, err
}

// PutCommentsViaCommentPostId Put Comments via CommentPostId
func PutCommentsViaCommentPostId(CommentPostId_ int64, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentPostId: CommentPostId_})
	return row, err
}

// PutCommentsViaCommentAuthor Put Comments via CommentAuthor
func PutCommentsViaCommentAuthor(CommentAuthor_ string, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentAuthor: CommentAuthor_})
	return row, err
}

// PutCommentsViaCommentAuthorEmail Put Comments via CommentAuthorEmail
func PutCommentsViaCommentAuthorEmail(CommentAuthorEmail_ string, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentAuthorEmail: CommentAuthorEmail_})
	return row, err
}

// PutCommentsViaCommentAuthorUrl Put Comments via CommentAuthorUrl
func PutCommentsViaCommentAuthorUrl(CommentAuthorUrl_ string, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentAuthorUrl: CommentAuthorUrl_})
	return row, err
}

// PutCommentsViaCommentAuthorIp Put Comments via CommentAuthorIp
func PutCommentsViaCommentAuthorIp(CommentAuthorIp_ string, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentAuthorIp: CommentAuthorIp_})
	return row, err
}

// PutCommentsViaCommentDate Put Comments via CommentDate
func PutCommentsViaCommentDate(CommentDate_ time.Time, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentDate: CommentDate_})
	return row, err
}

// PutCommentsViaCommentDateGmt Put Comments via CommentDateGmt
func PutCommentsViaCommentDateGmt(CommentDateGmt_ time.Time, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentDateGmt: CommentDateGmt_})
	return row, err
}

// PutCommentsViaCommentContent Put Comments via CommentContent
func PutCommentsViaCommentContent(CommentContent_ string, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentContent: CommentContent_})
	return row, err
}

// PutCommentsViaCommentKarma Put Comments via CommentKarma
func PutCommentsViaCommentKarma(CommentKarma_ int, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentKarma: CommentKarma_})
	return row, err
}

// PutCommentsViaCommentApproved Put Comments via CommentApproved
func PutCommentsViaCommentApproved(CommentApproved_ string, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentApproved: CommentApproved_})
	return row, err
}

// PutCommentsViaCommentAgent Put Comments via CommentAgent
func PutCommentsViaCommentAgent(CommentAgent_ string, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentAgent: CommentAgent_})
	return row, err
}

// PutCommentsViaCommentType Put Comments via CommentType
func PutCommentsViaCommentType(CommentType_ string, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentType: CommentType_})
	return row, err
}

// PutCommentsViaCommentParent Put Comments via CommentParent
func PutCommentsViaCommentParent(CommentParent_ int64, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{CommentParent: CommentParent_})
	return row, err
}

// PutCommentsViaUserId Put Comments via UserId
func PutCommentsViaUserId(UserId_ int64, iComments *Comments) (int64, error) {
	row, err := Engine.Update(iComments, &Comments{UserId: UserId_})
	return row, err
}

// UpdateCommentsViaCommentId via map[string]interface{}{}
func UpdateCommentsViaCommentId(iCommentId int64, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_ID = ?", iCommentId).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentPostId via map[string]interface{}{}
func UpdateCommentsViaCommentPostId(iCommentPostId int64, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_post_ID = ?", iCommentPostId).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentAuthor via map[string]interface{}{}
func UpdateCommentsViaCommentAuthor(iCommentAuthor string, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_author = ?", iCommentAuthor).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentAuthorEmail via map[string]interface{}{}
func UpdateCommentsViaCommentAuthorEmail(iCommentAuthorEmail string, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_author_email = ?", iCommentAuthorEmail).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentAuthorUrl via map[string]interface{}{}
func UpdateCommentsViaCommentAuthorUrl(iCommentAuthorUrl string, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_author_url = ?", iCommentAuthorUrl).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentAuthorIp via map[string]interface{}{}
func UpdateCommentsViaCommentAuthorIp(iCommentAuthorIp string, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_author_IP = ?", iCommentAuthorIp).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentDate via map[string]interface{}{}
func UpdateCommentsViaCommentDate(iCommentDate time.Time, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_date = ?", iCommentDate).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentDateGmt via map[string]interface{}{}
func UpdateCommentsViaCommentDateGmt(iCommentDateGmt time.Time, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_date_gmt = ?", iCommentDateGmt).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentContent via map[string]interface{}{}
func UpdateCommentsViaCommentContent(iCommentContent string, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_content = ?", iCommentContent).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentKarma via map[string]interface{}{}
func UpdateCommentsViaCommentKarma(iCommentKarma int, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_karma = ?", iCommentKarma).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentApproved via map[string]interface{}{}
func UpdateCommentsViaCommentApproved(iCommentApproved string, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_approved = ?", iCommentApproved).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentAgent via map[string]interface{}{}
func UpdateCommentsViaCommentAgent(iCommentAgent string, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_agent = ?", iCommentAgent).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentType via map[string]interface{}{}
func UpdateCommentsViaCommentType(iCommentType string, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_type = ?", iCommentType).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaCommentParent via map[string]interface{}{}
func UpdateCommentsViaCommentParent(iCommentParent int64, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("comment_parent = ?", iCommentParent).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateCommentsViaUserId via map[string]interface{}{}
func UpdateCommentsViaUserId(iUserId int64, iCommentsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Comments)).Where("user_id = ?", iUserId).Update(iCommentsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteComments Delete Comments
func DeleteComments(iCommentId int64) (int64, error) {
	row, err := Engine.Id(iCommentId).Delete(new(Comments))
	return row, err
}

// DeleteCommentsViaCommentId Delete Comments via CommentId
func DeleteCommentsViaCommentId(iCommentId int64) (err error) {
	var has bool
	var _Comments = &Comments{CommentId: iCommentId}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_ID = ?", iCommentId).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentPostId Delete Comments via CommentPostId
func DeleteCommentsViaCommentPostId(iCommentPostId int64) (err error) {
	var has bool
	var _Comments = &Comments{CommentPostId: iCommentPostId}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_post_ID = ?", iCommentPostId).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentAuthor Delete Comments via CommentAuthor
func DeleteCommentsViaCommentAuthor(iCommentAuthor string) (err error) {
	var has bool
	var _Comments = &Comments{CommentAuthor: iCommentAuthor}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_author = ?", iCommentAuthor).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentAuthorEmail Delete Comments via CommentAuthorEmail
func DeleteCommentsViaCommentAuthorEmail(iCommentAuthorEmail string) (err error) {
	var has bool
	var _Comments = &Comments{CommentAuthorEmail: iCommentAuthorEmail}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_author_email = ?", iCommentAuthorEmail).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentAuthorUrl Delete Comments via CommentAuthorUrl
func DeleteCommentsViaCommentAuthorUrl(iCommentAuthorUrl string) (err error) {
	var has bool
	var _Comments = &Comments{CommentAuthorUrl: iCommentAuthorUrl}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_author_url = ?", iCommentAuthorUrl).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentAuthorIp Delete Comments via CommentAuthorIp
func DeleteCommentsViaCommentAuthorIp(iCommentAuthorIp string) (err error) {
	var has bool
	var _Comments = &Comments{CommentAuthorIp: iCommentAuthorIp}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_author_IP = ?", iCommentAuthorIp).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentDate Delete Comments via CommentDate
func DeleteCommentsViaCommentDate(iCommentDate time.Time) (err error) {
	var has bool
	var _Comments = &Comments{CommentDate: iCommentDate}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_date = ?", iCommentDate).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentDateGmt Delete Comments via CommentDateGmt
func DeleteCommentsViaCommentDateGmt(iCommentDateGmt time.Time) (err error) {
	var has bool
	var _Comments = &Comments{CommentDateGmt: iCommentDateGmt}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_date_gmt = ?", iCommentDateGmt).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentContent Delete Comments via CommentContent
func DeleteCommentsViaCommentContent(iCommentContent string) (err error) {
	var has bool
	var _Comments = &Comments{CommentContent: iCommentContent}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_content = ?", iCommentContent).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentKarma Delete Comments via CommentKarma
func DeleteCommentsViaCommentKarma(iCommentKarma int) (err error) {
	var has bool
	var _Comments = &Comments{CommentKarma: iCommentKarma}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_karma = ?", iCommentKarma).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentApproved Delete Comments via CommentApproved
func DeleteCommentsViaCommentApproved(iCommentApproved string) (err error) {
	var has bool
	var _Comments = &Comments{CommentApproved: iCommentApproved}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_approved = ?", iCommentApproved).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentAgent Delete Comments via CommentAgent
func DeleteCommentsViaCommentAgent(iCommentAgent string) (err error) {
	var has bool
	var _Comments = &Comments{CommentAgent: iCommentAgent}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_agent = ?", iCommentAgent).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentType Delete Comments via CommentType
func DeleteCommentsViaCommentType(iCommentType string) (err error) {
	var has bool
	var _Comments = &Comments{CommentType: iCommentType}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_type = ?", iCommentType).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaCommentParent Delete Comments via CommentParent
func DeleteCommentsViaCommentParent(iCommentParent int64) (err error) {
	var has bool
	var _Comments = &Comments{CommentParent: iCommentParent}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("comment_parent = ?", iCommentParent).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteCommentsViaUserId Delete Comments via UserId
func DeleteCommentsViaUserId(iUserId int64) (err error) {
	var has bool
	var _Comments = &Comments{UserId: iUserId}
	if has, err = Engine.Get(_Comments); (has == true) && (err == nil) {
		if row, err := Engine.Where("user_id = ?", iUserId).Delete(new(Comments)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
