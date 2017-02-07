package model

import (
	"time"
)

type Links struct {
	LinkId          int64     `xorm:"not null pk autoincr BIGINT(20)"`
	LinkUrl         string    `xorm:"not null default '' VARCHAR(255)"`
	LinkName        string    `xorm:"not null default '' VARCHAR(255)"`
	LinkImage       string    `xorm:"not null default '' VARCHAR(255)"`
	LinkTarget      string    `xorm:"not null default '' VARCHAR(25)"`
	LinkDescription string    `xorm:"not null default '' VARCHAR(255)"`
	LinkVisible     string    `xorm:"not null default 'Y' index VARCHAR(20)"`
	LinkOwner       int64     `xorm:"not null default 1 BIGINT(20)"`
	LinkRating      int       `xorm:"not null default 0 INT(11)"`
	LinkUpdated     time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	LinkRel         string    `xorm:"not null default '' VARCHAR(255)"`
	LinkNotes       string    `xorm:"not null MEDIUMTEXT"`
	LinkRss         string    `xorm:"not null default '' VARCHAR(255)"`
}

// GetLinksesCount Linkss' Count
func GetLinksesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Links{})
	return total, err
}

// GetLinksCountByLinkId Get Links via LinkId
func GetLinksCountByLinkId(iLinkId int64) int64 {
	n, _ := Engine.Where("link_id = ?", iLinkId).Count(&Links{LinkId: iLinkId})
	return n
}

// GetLinksCountByLinkUrl Get Links via LinkUrl
func GetLinksCountByLinkUrl(iLinkUrl string) int64 {
	n, _ := Engine.Where("link_url = ?", iLinkUrl).Count(&Links{LinkUrl: iLinkUrl})
	return n
}

// GetLinksCountByLinkName Get Links via LinkName
func GetLinksCountByLinkName(iLinkName string) int64 {
	n, _ := Engine.Where("link_name = ?", iLinkName).Count(&Links{LinkName: iLinkName})
	return n
}

// GetLinksCountByLinkImage Get Links via LinkImage
func GetLinksCountByLinkImage(iLinkImage string) int64 {
	n, _ := Engine.Where("link_image = ?", iLinkImage).Count(&Links{LinkImage: iLinkImage})
	return n
}

// GetLinksCountByLinkTarget Get Links via LinkTarget
func GetLinksCountByLinkTarget(iLinkTarget string) int64 {
	n, _ := Engine.Where("link_target = ?", iLinkTarget).Count(&Links{LinkTarget: iLinkTarget})
	return n
}

// GetLinksCountByLinkDescription Get Links via LinkDescription
func GetLinksCountByLinkDescription(iLinkDescription string) int64 {
	n, _ := Engine.Where("link_description = ?", iLinkDescription).Count(&Links{LinkDescription: iLinkDescription})
	return n
}

// GetLinksCountByLinkVisible Get Links via LinkVisible
func GetLinksCountByLinkVisible(iLinkVisible string) int64 {
	n, _ := Engine.Where("link_visible = ?", iLinkVisible).Count(&Links{LinkVisible: iLinkVisible})
	return n
}

// GetLinksCountByLinkOwner Get Links via LinkOwner
func GetLinksCountByLinkOwner(iLinkOwner int64) int64 {
	n, _ := Engine.Where("link_owner = ?", iLinkOwner).Count(&Links{LinkOwner: iLinkOwner})
	return n
}

// GetLinksCountByLinkRating Get Links via LinkRating
func GetLinksCountByLinkRating(iLinkRating int) int64 {
	n, _ := Engine.Where("link_rating = ?", iLinkRating).Count(&Links{LinkRating: iLinkRating})
	return n
}

// GetLinksCountByLinkUpdated Get Links via LinkUpdated
func GetLinksCountByLinkUpdated(iLinkUpdated time.Time) int64 {
	n, _ := Engine.Where("link_updated = ?", iLinkUpdated).Count(&Links{LinkUpdated: iLinkUpdated})
	return n
}

// GetLinksCountByLinkRel Get Links via LinkRel
func GetLinksCountByLinkRel(iLinkRel string) int64 {
	n, _ := Engine.Where("link_rel = ?", iLinkRel).Count(&Links{LinkRel: iLinkRel})
	return n
}

// GetLinksCountByLinkNotes Get Links via LinkNotes
func GetLinksCountByLinkNotes(iLinkNotes string) int64 {
	n, _ := Engine.Where("link_notes = ?", iLinkNotes).Count(&Links{LinkNotes: iLinkNotes})
	return n
}

// GetLinksCountByLinkRss Get Links via LinkRss
func GetLinksCountByLinkRss(iLinkRss string) int64 {
	n, _ := Engine.Where("link_rss = ?", iLinkRss).Count(&Links{LinkRss: iLinkRss})
	return n
}

// GetLinksesByLinkIdAndLinkUrlAndLinkName Get Linkses via LinkIdAndLinkUrlAndLinkName
func GetLinksesByLinkIdAndLinkUrlAndLinkName(offset int, limit int, LinkId_ int64, LinkUrl_ string, LinkName_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_url = ? and link_name = ?", LinkId_, LinkUrl_, LinkName_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUrlAndLinkImage Get Linkses via LinkIdAndLinkUrlAndLinkImage
func GetLinksesByLinkIdAndLinkUrlAndLinkImage(offset int, limit int, LinkId_ int64, LinkUrl_ string, LinkImage_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_url = ? and link_image = ?", LinkId_, LinkUrl_, LinkImage_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUrlAndLinkTarget Get Linkses via LinkIdAndLinkUrlAndLinkTarget
func GetLinksesByLinkIdAndLinkUrlAndLinkTarget(offset int, limit int, LinkId_ int64, LinkUrl_ string, LinkTarget_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_url = ? and link_target = ?", LinkId_, LinkUrl_, LinkTarget_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUrlAndLinkDescription Get Linkses via LinkIdAndLinkUrlAndLinkDescription
func GetLinksesByLinkIdAndLinkUrlAndLinkDescription(offset int, limit int, LinkId_ int64, LinkUrl_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_url = ? and link_description = ?", LinkId_, LinkUrl_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUrlAndLinkVisible Get Linkses via LinkIdAndLinkUrlAndLinkVisible
func GetLinksesByLinkIdAndLinkUrlAndLinkVisible(offset int, limit int, LinkId_ int64, LinkUrl_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_url = ? and link_visible = ?", LinkId_, LinkUrl_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUrlAndLinkOwner Get Linkses via LinkIdAndLinkUrlAndLinkOwner
func GetLinksesByLinkIdAndLinkUrlAndLinkOwner(offset int, limit int, LinkId_ int64, LinkUrl_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_url = ? and link_owner = ?", LinkId_, LinkUrl_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUrlAndLinkRating Get Linkses via LinkIdAndLinkUrlAndLinkRating
func GetLinksesByLinkIdAndLinkUrlAndLinkRating(offset int, limit int, LinkId_ int64, LinkUrl_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_url = ? and link_rating = ?", LinkId_, LinkUrl_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUrlAndLinkUpdated Get Linkses via LinkIdAndLinkUrlAndLinkUpdated
func GetLinksesByLinkIdAndLinkUrlAndLinkUpdated(offset int, limit int, LinkId_ int64, LinkUrl_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_url = ? and link_updated = ?", LinkId_, LinkUrl_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUrlAndLinkRel Get Linkses via LinkIdAndLinkUrlAndLinkRel
func GetLinksesByLinkIdAndLinkUrlAndLinkRel(offset int, limit int, LinkId_ int64, LinkUrl_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_url = ? and link_rel = ?", LinkId_, LinkUrl_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUrlAndLinkNotes Get Linkses via LinkIdAndLinkUrlAndLinkNotes
func GetLinksesByLinkIdAndLinkUrlAndLinkNotes(offset int, limit int, LinkId_ int64, LinkUrl_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_url = ? and link_notes = ?", LinkId_, LinkUrl_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUrlAndLinkRss Get Linkses via LinkIdAndLinkUrlAndLinkRss
func GetLinksesByLinkIdAndLinkUrlAndLinkRss(offset int, limit int, LinkId_ int64, LinkUrl_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_url = ? and link_rss = ?", LinkId_, LinkUrl_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkNameAndLinkImage Get Linkses via LinkIdAndLinkNameAndLinkImage
func GetLinksesByLinkIdAndLinkNameAndLinkImage(offset int, limit int, LinkId_ int64, LinkName_ string, LinkImage_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_name = ? and link_image = ?", LinkId_, LinkName_, LinkImage_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkNameAndLinkTarget Get Linkses via LinkIdAndLinkNameAndLinkTarget
func GetLinksesByLinkIdAndLinkNameAndLinkTarget(offset int, limit int, LinkId_ int64, LinkName_ string, LinkTarget_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_name = ? and link_target = ?", LinkId_, LinkName_, LinkTarget_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkNameAndLinkDescription Get Linkses via LinkIdAndLinkNameAndLinkDescription
func GetLinksesByLinkIdAndLinkNameAndLinkDescription(offset int, limit int, LinkId_ int64, LinkName_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_name = ? and link_description = ?", LinkId_, LinkName_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkNameAndLinkVisible Get Linkses via LinkIdAndLinkNameAndLinkVisible
func GetLinksesByLinkIdAndLinkNameAndLinkVisible(offset int, limit int, LinkId_ int64, LinkName_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_name = ? and link_visible = ?", LinkId_, LinkName_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkNameAndLinkOwner Get Linkses via LinkIdAndLinkNameAndLinkOwner
func GetLinksesByLinkIdAndLinkNameAndLinkOwner(offset int, limit int, LinkId_ int64, LinkName_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_name = ? and link_owner = ?", LinkId_, LinkName_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkNameAndLinkRating Get Linkses via LinkIdAndLinkNameAndLinkRating
func GetLinksesByLinkIdAndLinkNameAndLinkRating(offset int, limit int, LinkId_ int64, LinkName_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_name = ? and link_rating = ?", LinkId_, LinkName_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkNameAndLinkUpdated Get Linkses via LinkIdAndLinkNameAndLinkUpdated
func GetLinksesByLinkIdAndLinkNameAndLinkUpdated(offset int, limit int, LinkId_ int64, LinkName_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_name = ? and link_updated = ?", LinkId_, LinkName_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkNameAndLinkRel Get Linkses via LinkIdAndLinkNameAndLinkRel
func GetLinksesByLinkIdAndLinkNameAndLinkRel(offset int, limit int, LinkId_ int64, LinkName_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_name = ? and link_rel = ?", LinkId_, LinkName_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkNameAndLinkNotes Get Linkses via LinkIdAndLinkNameAndLinkNotes
func GetLinksesByLinkIdAndLinkNameAndLinkNotes(offset int, limit int, LinkId_ int64, LinkName_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_name = ? and link_notes = ?", LinkId_, LinkName_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkNameAndLinkRss Get Linkses via LinkIdAndLinkNameAndLinkRss
func GetLinksesByLinkIdAndLinkNameAndLinkRss(offset int, limit int, LinkId_ int64, LinkName_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_name = ? and link_rss = ?", LinkId_, LinkName_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkImageAndLinkTarget Get Linkses via LinkIdAndLinkImageAndLinkTarget
func GetLinksesByLinkIdAndLinkImageAndLinkTarget(offset int, limit int, LinkId_ int64, LinkImage_ string, LinkTarget_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_image = ? and link_target = ?", LinkId_, LinkImage_, LinkTarget_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkImageAndLinkDescription Get Linkses via LinkIdAndLinkImageAndLinkDescription
func GetLinksesByLinkIdAndLinkImageAndLinkDescription(offset int, limit int, LinkId_ int64, LinkImage_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_image = ? and link_description = ?", LinkId_, LinkImage_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkImageAndLinkVisible Get Linkses via LinkIdAndLinkImageAndLinkVisible
func GetLinksesByLinkIdAndLinkImageAndLinkVisible(offset int, limit int, LinkId_ int64, LinkImage_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_image = ? and link_visible = ?", LinkId_, LinkImage_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkImageAndLinkOwner Get Linkses via LinkIdAndLinkImageAndLinkOwner
func GetLinksesByLinkIdAndLinkImageAndLinkOwner(offset int, limit int, LinkId_ int64, LinkImage_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_image = ? and link_owner = ?", LinkId_, LinkImage_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkImageAndLinkRating Get Linkses via LinkIdAndLinkImageAndLinkRating
func GetLinksesByLinkIdAndLinkImageAndLinkRating(offset int, limit int, LinkId_ int64, LinkImage_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_image = ? and link_rating = ?", LinkId_, LinkImage_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkImageAndLinkUpdated Get Linkses via LinkIdAndLinkImageAndLinkUpdated
func GetLinksesByLinkIdAndLinkImageAndLinkUpdated(offset int, limit int, LinkId_ int64, LinkImage_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_image = ? and link_updated = ?", LinkId_, LinkImage_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkImageAndLinkRel Get Linkses via LinkIdAndLinkImageAndLinkRel
func GetLinksesByLinkIdAndLinkImageAndLinkRel(offset int, limit int, LinkId_ int64, LinkImage_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_image = ? and link_rel = ?", LinkId_, LinkImage_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkImageAndLinkNotes Get Linkses via LinkIdAndLinkImageAndLinkNotes
func GetLinksesByLinkIdAndLinkImageAndLinkNotes(offset int, limit int, LinkId_ int64, LinkImage_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_image = ? and link_notes = ?", LinkId_, LinkImage_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkImageAndLinkRss Get Linkses via LinkIdAndLinkImageAndLinkRss
func GetLinksesByLinkIdAndLinkImageAndLinkRss(offset int, limit int, LinkId_ int64, LinkImage_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_image = ? and link_rss = ?", LinkId_, LinkImage_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkTargetAndLinkDescription Get Linkses via LinkIdAndLinkTargetAndLinkDescription
func GetLinksesByLinkIdAndLinkTargetAndLinkDescription(offset int, limit int, LinkId_ int64, LinkTarget_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_target = ? and link_description = ?", LinkId_, LinkTarget_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkTargetAndLinkVisible Get Linkses via LinkIdAndLinkTargetAndLinkVisible
func GetLinksesByLinkIdAndLinkTargetAndLinkVisible(offset int, limit int, LinkId_ int64, LinkTarget_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_target = ? and link_visible = ?", LinkId_, LinkTarget_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkTargetAndLinkOwner Get Linkses via LinkIdAndLinkTargetAndLinkOwner
func GetLinksesByLinkIdAndLinkTargetAndLinkOwner(offset int, limit int, LinkId_ int64, LinkTarget_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_target = ? and link_owner = ?", LinkId_, LinkTarget_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkTargetAndLinkRating Get Linkses via LinkIdAndLinkTargetAndLinkRating
func GetLinksesByLinkIdAndLinkTargetAndLinkRating(offset int, limit int, LinkId_ int64, LinkTarget_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_target = ? and link_rating = ?", LinkId_, LinkTarget_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkTargetAndLinkUpdated Get Linkses via LinkIdAndLinkTargetAndLinkUpdated
func GetLinksesByLinkIdAndLinkTargetAndLinkUpdated(offset int, limit int, LinkId_ int64, LinkTarget_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_target = ? and link_updated = ?", LinkId_, LinkTarget_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkTargetAndLinkRel Get Linkses via LinkIdAndLinkTargetAndLinkRel
func GetLinksesByLinkIdAndLinkTargetAndLinkRel(offset int, limit int, LinkId_ int64, LinkTarget_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_target = ? and link_rel = ?", LinkId_, LinkTarget_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkTargetAndLinkNotes Get Linkses via LinkIdAndLinkTargetAndLinkNotes
func GetLinksesByLinkIdAndLinkTargetAndLinkNotes(offset int, limit int, LinkId_ int64, LinkTarget_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_target = ? and link_notes = ?", LinkId_, LinkTarget_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkTargetAndLinkRss Get Linkses via LinkIdAndLinkTargetAndLinkRss
func GetLinksesByLinkIdAndLinkTargetAndLinkRss(offset int, limit int, LinkId_ int64, LinkTarget_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_target = ? and link_rss = ?", LinkId_, LinkTarget_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkDescriptionAndLinkVisible Get Linkses via LinkIdAndLinkDescriptionAndLinkVisible
func GetLinksesByLinkIdAndLinkDescriptionAndLinkVisible(offset int, limit int, LinkId_ int64, LinkDescription_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_description = ? and link_visible = ?", LinkId_, LinkDescription_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkDescriptionAndLinkOwner Get Linkses via LinkIdAndLinkDescriptionAndLinkOwner
func GetLinksesByLinkIdAndLinkDescriptionAndLinkOwner(offset int, limit int, LinkId_ int64, LinkDescription_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_description = ? and link_owner = ?", LinkId_, LinkDescription_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkDescriptionAndLinkRating Get Linkses via LinkIdAndLinkDescriptionAndLinkRating
func GetLinksesByLinkIdAndLinkDescriptionAndLinkRating(offset int, limit int, LinkId_ int64, LinkDescription_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_description = ? and link_rating = ?", LinkId_, LinkDescription_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkDescriptionAndLinkUpdated Get Linkses via LinkIdAndLinkDescriptionAndLinkUpdated
func GetLinksesByLinkIdAndLinkDescriptionAndLinkUpdated(offset int, limit int, LinkId_ int64, LinkDescription_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_description = ? and link_updated = ?", LinkId_, LinkDescription_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkDescriptionAndLinkRel Get Linkses via LinkIdAndLinkDescriptionAndLinkRel
func GetLinksesByLinkIdAndLinkDescriptionAndLinkRel(offset int, limit int, LinkId_ int64, LinkDescription_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_description = ? and link_rel = ?", LinkId_, LinkDescription_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkDescriptionAndLinkNotes Get Linkses via LinkIdAndLinkDescriptionAndLinkNotes
func GetLinksesByLinkIdAndLinkDescriptionAndLinkNotes(offset int, limit int, LinkId_ int64, LinkDescription_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_description = ? and link_notes = ?", LinkId_, LinkDescription_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkDescriptionAndLinkRss Get Linkses via LinkIdAndLinkDescriptionAndLinkRss
func GetLinksesByLinkIdAndLinkDescriptionAndLinkRss(offset int, limit int, LinkId_ int64, LinkDescription_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_description = ? and link_rss = ?", LinkId_, LinkDescription_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkVisibleAndLinkOwner Get Linkses via LinkIdAndLinkVisibleAndLinkOwner
func GetLinksesByLinkIdAndLinkVisibleAndLinkOwner(offset int, limit int, LinkId_ int64, LinkVisible_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_visible = ? and link_owner = ?", LinkId_, LinkVisible_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkVisibleAndLinkRating Get Linkses via LinkIdAndLinkVisibleAndLinkRating
func GetLinksesByLinkIdAndLinkVisibleAndLinkRating(offset int, limit int, LinkId_ int64, LinkVisible_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_visible = ? and link_rating = ?", LinkId_, LinkVisible_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkVisibleAndLinkUpdated Get Linkses via LinkIdAndLinkVisibleAndLinkUpdated
func GetLinksesByLinkIdAndLinkVisibleAndLinkUpdated(offset int, limit int, LinkId_ int64, LinkVisible_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_visible = ? and link_updated = ?", LinkId_, LinkVisible_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkVisibleAndLinkRel Get Linkses via LinkIdAndLinkVisibleAndLinkRel
func GetLinksesByLinkIdAndLinkVisibleAndLinkRel(offset int, limit int, LinkId_ int64, LinkVisible_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_visible = ? and link_rel = ?", LinkId_, LinkVisible_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkVisibleAndLinkNotes Get Linkses via LinkIdAndLinkVisibleAndLinkNotes
func GetLinksesByLinkIdAndLinkVisibleAndLinkNotes(offset int, limit int, LinkId_ int64, LinkVisible_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_visible = ? and link_notes = ?", LinkId_, LinkVisible_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkVisibleAndLinkRss Get Linkses via LinkIdAndLinkVisibleAndLinkRss
func GetLinksesByLinkIdAndLinkVisibleAndLinkRss(offset int, limit int, LinkId_ int64, LinkVisible_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_visible = ? and link_rss = ?", LinkId_, LinkVisible_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkOwnerAndLinkRating Get Linkses via LinkIdAndLinkOwnerAndLinkRating
func GetLinksesByLinkIdAndLinkOwnerAndLinkRating(offset int, limit int, LinkId_ int64, LinkOwner_ int64, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_owner = ? and link_rating = ?", LinkId_, LinkOwner_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkOwnerAndLinkUpdated Get Linkses via LinkIdAndLinkOwnerAndLinkUpdated
func GetLinksesByLinkIdAndLinkOwnerAndLinkUpdated(offset int, limit int, LinkId_ int64, LinkOwner_ int64, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_owner = ? and link_updated = ?", LinkId_, LinkOwner_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkOwnerAndLinkRel Get Linkses via LinkIdAndLinkOwnerAndLinkRel
func GetLinksesByLinkIdAndLinkOwnerAndLinkRel(offset int, limit int, LinkId_ int64, LinkOwner_ int64, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_owner = ? and link_rel = ?", LinkId_, LinkOwner_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkOwnerAndLinkNotes Get Linkses via LinkIdAndLinkOwnerAndLinkNotes
func GetLinksesByLinkIdAndLinkOwnerAndLinkNotes(offset int, limit int, LinkId_ int64, LinkOwner_ int64, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_owner = ? and link_notes = ?", LinkId_, LinkOwner_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkOwnerAndLinkRss Get Linkses via LinkIdAndLinkOwnerAndLinkRss
func GetLinksesByLinkIdAndLinkOwnerAndLinkRss(offset int, limit int, LinkId_ int64, LinkOwner_ int64, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_owner = ? and link_rss = ?", LinkId_, LinkOwner_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkRatingAndLinkUpdated Get Linkses via LinkIdAndLinkRatingAndLinkUpdated
func GetLinksesByLinkIdAndLinkRatingAndLinkUpdated(offset int, limit int, LinkId_ int64, LinkRating_ int, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_rating = ? and link_updated = ?", LinkId_, LinkRating_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkRatingAndLinkRel Get Linkses via LinkIdAndLinkRatingAndLinkRel
func GetLinksesByLinkIdAndLinkRatingAndLinkRel(offset int, limit int, LinkId_ int64, LinkRating_ int, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_rating = ? and link_rel = ?", LinkId_, LinkRating_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkRatingAndLinkNotes Get Linkses via LinkIdAndLinkRatingAndLinkNotes
func GetLinksesByLinkIdAndLinkRatingAndLinkNotes(offset int, limit int, LinkId_ int64, LinkRating_ int, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_rating = ? and link_notes = ?", LinkId_, LinkRating_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkRatingAndLinkRss Get Linkses via LinkIdAndLinkRatingAndLinkRss
func GetLinksesByLinkIdAndLinkRatingAndLinkRss(offset int, limit int, LinkId_ int64, LinkRating_ int, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_rating = ? and link_rss = ?", LinkId_, LinkRating_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUpdatedAndLinkRel Get Linkses via LinkIdAndLinkUpdatedAndLinkRel
func GetLinksesByLinkIdAndLinkUpdatedAndLinkRel(offset int, limit int, LinkId_ int64, LinkUpdated_ time.Time, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_updated = ? and link_rel = ?", LinkId_, LinkUpdated_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUpdatedAndLinkNotes Get Linkses via LinkIdAndLinkUpdatedAndLinkNotes
func GetLinksesByLinkIdAndLinkUpdatedAndLinkNotes(offset int, limit int, LinkId_ int64, LinkUpdated_ time.Time, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_updated = ? and link_notes = ?", LinkId_, LinkUpdated_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUpdatedAndLinkRss Get Linkses via LinkIdAndLinkUpdatedAndLinkRss
func GetLinksesByLinkIdAndLinkUpdatedAndLinkRss(offset int, limit int, LinkId_ int64, LinkUpdated_ time.Time, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_updated = ? and link_rss = ?", LinkId_, LinkUpdated_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkRelAndLinkNotes Get Linkses via LinkIdAndLinkRelAndLinkNotes
func GetLinksesByLinkIdAndLinkRelAndLinkNotes(offset int, limit int, LinkId_ int64, LinkRel_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_rel = ? and link_notes = ?", LinkId_, LinkRel_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkRelAndLinkRss Get Linkses via LinkIdAndLinkRelAndLinkRss
func GetLinksesByLinkIdAndLinkRelAndLinkRss(offset int, limit int, LinkId_ int64, LinkRel_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_rel = ? and link_rss = ?", LinkId_, LinkRel_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkNotesAndLinkRss Get Linkses via LinkIdAndLinkNotesAndLinkRss
func GetLinksesByLinkIdAndLinkNotesAndLinkRss(offset int, limit int, LinkId_ int64, LinkNotes_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_notes = ? and link_rss = ?", LinkId_, LinkNotes_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkNameAndLinkImage Get Linkses via LinkUrlAndLinkNameAndLinkImage
func GetLinksesByLinkUrlAndLinkNameAndLinkImage(offset int, limit int, LinkUrl_ string, LinkName_ string, LinkImage_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_name = ? and link_image = ?", LinkUrl_, LinkName_, LinkImage_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkNameAndLinkTarget Get Linkses via LinkUrlAndLinkNameAndLinkTarget
func GetLinksesByLinkUrlAndLinkNameAndLinkTarget(offset int, limit int, LinkUrl_ string, LinkName_ string, LinkTarget_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_name = ? and link_target = ?", LinkUrl_, LinkName_, LinkTarget_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkNameAndLinkDescription Get Linkses via LinkUrlAndLinkNameAndLinkDescription
func GetLinksesByLinkUrlAndLinkNameAndLinkDescription(offset int, limit int, LinkUrl_ string, LinkName_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_name = ? and link_description = ?", LinkUrl_, LinkName_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkNameAndLinkVisible Get Linkses via LinkUrlAndLinkNameAndLinkVisible
func GetLinksesByLinkUrlAndLinkNameAndLinkVisible(offset int, limit int, LinkUrl_ string, LinkName_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_name = ? and link_visible = ?", LinkUrl_, LinkName_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkNameAndLinkOwner Get Linkses via LinkUrlAndLinkNameAndLinkOwner
func GetLinksesByLinkUrlAndLinkNameAndLinkOwner(offset int, limit int, LinkUrl_ string, LinkName_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_name = ? and link_owner = ?", LinkUrl_, LinkName_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkNameAndLinkRating Get Linkses via LinkUrlAndLinkNameAndLinkRating
func GetLinksesByLinkUrlAndLinkNameAndLinkRating(offset int, limit int, LinkUrl_ string, LinkName_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_name = ? and link_rating = ?", LinkUrl_, LinkName_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkNameAndLinkUpdated Get Linkses via LinkUrlAndLinkNameAndLinkUpdated
func GetLinksesByLinkUrlAndLinkNameAndLinkUpdated(offset int, limit int, LinkUrl_ string, LinkName_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_name = ? and link_updated = ?", LinkUrl_, LinkName_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkNameAndLinkRel Get Linkses via LinkUrlAndLinkNameAndLinkRel
func GetLinksesByLinkUrlAndLinkNameAndLinkRel(offset int, limit int, LinkUrl_ string, LinkName_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_name = ? and link_rel = ?", LinkUrl_, LinkName_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkNameAndLinkNotes Get Linkses via LinkUrlAndLinkNameAndLinkNotes
func GetLinksesByLinkUrlAndLinkNameAndLinkNotes(offset int, limit int, LinkUrl_ string, LinkName_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_name = ? and link_notes = ?", LinkUrl_, LinkName_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkNameAndLinkRss Get Linkses via LinkUrlAndLinkNameAndLinkRss
func GetLinksesByLinkUrlAndLinkNameAndLinkRss(offset int, limit int, LinkUrl_ string, LinkName_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_name = ? and link_rss = ?", LinkUrl_, LinkName_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkImageAndLinkTarget Get Linkses via LinkUrlAndLinkImageAndLinkTarget
func GetLinksesByLinkUrlAndLinkImageAndLinkTarget(offset int, limit int, LinkUrl_ string, LinkImage_ string, LinkTarget_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_image = ? and link_target = ?", LinkUrl_, LinkImage_, LinkTarget_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkImageAndLinkDescription Get Linkses via LinkUrlAndLinkImageAndLinkDescription
func GetLinksesByLinkUrlAndLinkImageAndLinkDescription(offset int, limit int, LinkUrl_ string, LinkImage_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_image = ? and link_description = ?", LinkUrl_, LinkImage_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkImageAndLinkVisible Get Linkses via LinkUrlAndLinkImageAndLinkVisible
func GetLinksesByLinkUrlAndLinkImageAndLinkVisible(offset int, limit int, LinkUrl_ string, LinkImage_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_image = ? and link_visible = ?", LinkUrl_, LinkImage_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkImageAndLinkOwner Get Linkses via LinkUrlAndLinkImageAndLinkOwner
func GetLinksesByLinkUrlAndLinkImageAndLinkOwner(offset int, limit int, LinkUrl_ string, LinkImage_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_image = ? and link_owner = ?", LinkUrl_, LinkImage_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkImageAndLinkRating Get Linkses via LinkUrlAndLinkImageAndLinkRating
func GetLinksesByLinkUrlAndLinkImageAndLinkRating(offset int, limit int, LinkUrl_ string, LinkImage_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_image = ? and link_rating = ?", LinkUrl_, LinkImage_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkImageAndLinkUpdated Get Linkses via LinkUrlAndLinkImageAndLinkUpdated
func GetLinksesByLinkUrlAndLinkImageAndLinkUpdated(offset int, limit int, LinkUrl_ string, LinkImage_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_image = ? and link_updated = ?", LinkUrl_, LinkImage_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkImageAndLinkRel Get Linkses via LinkUrlAndLinkImageAndLinkRel
func GetLinksesByLinkUrlAndLinkImageAndLinkRel(offset int, limit int, LinkUrl_ string, LinkImage_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_image = ? and link_rel = ?", LinkUrl_, LinkImage_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkImageAndLinkNotes Get Linkses via LinkUrlAndLinkImageAndLinkNotes
func GetLinksesByLinkUrlAndLinkImageAndLinkNotes(offset int, limit int, LinkUrl_ string, LinkImage_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_image = ? and link_notes = ?", LinkUrl_, LinkImage_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkImageAndLinkRss Get Linkses via LinkUrlAndLinkImageAndLinkRss
func GetLinksesByLinkUrlAndLinkImageAndLinkRss(offset int, limit int, LinkUrl_ string, LinkImage_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_image = ? and link_rss = ?", LinkUrl_, LinkImage_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkTargetAndLinkDescription Get Linkses via LinkUrlAndLinkTargetAndLinkDescription
func GetLinksesByLinkUrlAndLinkTargetAndLinkDescription(offset int, limit int, LinkUrl_ string, LinkTarget_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_target = ? and link_description = ?", LinkUrl_, LinkTarget_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkTargetAndLinkVisible Get Linkses via LinkUrlAndLinkTargetAndLinkVisible
func GetLinksesByLinkUrlAndLinkTargetAndLinkVisible(offset int, limit int, LinkUrl_ string, LinkTarget_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_target = ? and link_visible = ?", LinkUrl_, LinkTarget_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkTargetAndLinkOwner Get Linkses via LinkUrlAndLinkTargetAndLinkOwner
func GetLinksesByLinkUrlAndLinkTargetAndLinkOwner(offset int, limit int, LinkUrl_ string, LinkTarget_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_target = ? and link_owner = ?", LinkUrl_, LinkTarget_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkTargetAndLinkRating Get Linkses via LinkUrlAndLinkTargetAndLinkRating
func GetLinksesByLinkUrlAndLinkTargetAndLinkRating(offset int, limit int, LinkUrl_ string, LinkTarget_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_target = ? and link_rating = ?", LinkUrl_, LinkTarget_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkTargetAndLinkUpdated Get Linkses via LinkUrlAndLinkTargetAndLinkUpdated
func GetLinksesByLinkUrlAndLinkTargetAndLinkUpdated(offset int, limit int, LinkUrl_ string, LinkTarget_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_target = ? and link_updated = ?", LinkUrl_, LinkTarget_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkTargetAndLinkRel Get Linkses via LinkUrlAndLinkTargetAndLinkRel
func GetLinksesByLinkUrlAndLinkTargetAndLinkRel(offset int, limit int, LinkUrl_ string, LinkTarget_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_target = ? and link_rel = ?", LinkUrl_, LinkTarget_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkTargetAndLinkNotes Get Linkses via LinkUrlAndLinkTargetAndLinkNotes
func GetLinksesByLinkUrlAndLinkTargetAndLinkNotes(offset int, limit int, LinkUrl_ string, LinkTarget_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_target = ? and link_notes = ?", LinkUrl_, LinkTarget_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkTargetAndLinkRss Get Linkses via LinkUrlAndLinkTargetAndLinkRss
func GetLinksesByLinkUrlAndLinkTargetAndLinkRss(offset int, limit int, LinkUrl_ string, LinkTarget_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_target = ? and link_rss = ?", LinkUrl_, LinkTarget_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkDescriptionAndLinkVisible Get Linkses via LinkUrlAndLinkDescriptionAndLinkVisible
func GetLinksesByLinkUrlAndLinkDescriptionAndLinkVisible(offset int, limit int, LinkUrl_ string, LinkDescription_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_description = ? and link_visible = ?", LinkUrl_, LinkDescription_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkDescriptionAndLinkOwner Get Linkses via LinkUrlAndLinkDescriptionAndLinkOwner
func GetLinksesByLinkUrlAndLinkDescriptionAndLinkOwner(offset int, limit int, LinkUrl_ string, LinkDescription_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_description = ? and link_owner = ?", LinkUrl_, LinkDescription_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkDescriptionAndLinkRating Get Linkses via LinkUrlAndLinkDescriptionAndLinkRating
func GetLinksesByLinkUrlAndLinkDescriptionAndLinkRating(offset int, limit int, LinkUrl_ string, LinkDescription_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_description = ? and link_rating = ?", LinkUrl_, LinkDescription_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkDescriptionAndLinkUpdated Get Linkses via LinkUrlAndLinkDescriptionAndLinkUpdated
func GetLinksesByLinkUrlAndLinkDescriptionAndLinkUpdated(offset int, limit int, LinkUrl_ string, LinkDescription_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_description = ? and link_updated = ?", LinkUrl_, LinkDescription_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkDescriptionAndLinkRel Get Linkses via LinkUrlAndLinkDescriptionAndLinkRel
func GetLinksesByLinkUrlAndLinkDescriptionAndLinkRel(offset int, limit int, LinkUrl_ string, LinkDescription_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_description = ? and link_rel = ?", LinkUrl_, LinkDescription_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkDescriptionAndLinkNotes Get Linkses via LinkUrlAndLinkDescriptionAndLinkNotes
func GetLinksesByLinkUrlAndLinkDescriptionAndLinkNotes(offset int, limit int, LinkUrl_ string, LinkDescription_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_description = ? and link_notes = ?", LinkUrl_, LinkDescription_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkDescriptionAndLinkRss Get Linkses via LinkUrlAndLinkDescriptionAndLinkRss
func GetLinksesByLinkUrlAndLinkDescriptionAndLinkRss(offset int, limit int, LinkUrl_ string, LinkDescription_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_description = ? and link_rss = ?", LinkUrl_, LinkDescription_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkVisibleAndLinkOwner Get Linkses via LinkUrlAndLinkVisibleAndLinkOwner
func GetLinksesByLinkUrlAndLinkVisibleAndLinkOwner(offset int, limit int, LinkUrl_ string, LinkVisible_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_visible = ? and link_owner = ?", LinkUrl_, LinkVisible_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkVisibleAndLinkRating Get Linkses via LinkUrlAndLinkVisibleAndLinkRating
func GetLinksesByLinkUrlAndLinkVisibleAndLinkRating(offset int, limit int, LinkUrl_ string, LinkVisible_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_visible = ? and link_rating = ?", LinkUrl_, LinkVisible_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkVisibleAndLinkUpdated Get Linkses via LinkUrlAndLinkVisibleAndLinkUpdated
func GetLinksesByLinkUrlAndLinkVisibleAndLinkUpdated(offset int, limit int, LinkUrl_ string, LinkVisible_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_visible = ? and link_updated = ?", LinkUrl_, LinkVisible_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkVisibleAndLinkRel Get Linkses via LinkUrlAndLinkVisibleAndLinkRel
func GetLinksesByLinkUrlAndLinkVisibleAndLinkRel(offset int, limit int, LinkUrl_ string, LinkVisible_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_visible = ? and link_rel = ?", LinkUrl_, LinkVisible_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkVisibleAndLinkNotes Get Linkses via LinkUrlAndLinkVisibleAndLinkNotes
func GetLinksesByLinkUrlAndLinkVisibleAndLinkNotes(offset int, limit int, LinkUrl_ string, LinkVisible_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_visible = ? and link_notes = ?", LinkUrl_, LinkVisible_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkVisibleAndLinkRss Get Linkses via LinkUrlAndLinkVisibleAndLinkRss
func GetLinksesByLinkUrlAndLinkVisibleAndLinkRss(offset int, limit int, LinkUrl_ string, LinkVisible_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_visible = ? and link_rss = ?", LinkUrl_, LinkVisible_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkOwnerAndLinkRating Get Linkses via LinkUrlAndLinkOwnerAndLinkRating
func GetLinksesByLinkUrlAndLinkOwnerAndLinkRating(offset int, limit int, LinkUrl_ string, LinkOwner_ int64, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_owner = ? and link_rating = ?", LinkUrl_, LinkOwner_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkOwnerAndLinkUpdated Get Linkses via LinkUrlAndLinkOwnerAndLinkUpdated
func GetLinksesByLinkUrlAndLinkOwnerAndLinkUpdated(offset int, limit int, LinkUrl_ string, LinkOwner_ int64, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_owner = ? and link_updated = ?", LinkUrl_, LinkOwner_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkOwnerAndLinkRel Get Linkses via LinkUrlAndLinkOwnerAndLinkRel
func GetLinksesByLinkUrlAndLinkOwnerAndLinkRel(offset int, limit int, LinkUrl_ string, LinkOwner_ int64, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_owner = ? and link_rel = ?", LinkUrl_, LinkOwner_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkOwnerAndLinkNotes Get Linkses via LinkUrlAndLinkOwnerAndLinkNotes
func GetLinksesByLinkUrlAndLinkOwnerAndLinkNotes(offset int, limit int, LinkUrl_ string, LinkOwner_ int64, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_owner = ? and link_notes = ?", LinkUrl_, LinkOwner_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkOwnerAndLinkRss Get Linkses via LinkUrlAndLinkOwnerAndLinkRss
func GetLinksesByLinkUrlAndLinkOwnerAndLinkRss(offset int, limit int, LinkUrl_ string, LinkOwner_ int64, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_owner = ? and link_rss = ?", LinkUrl_, LinkOwner_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkRatingAndLinkUpdated Get Linkses via LinkUrlAndLinkRatingAndLinkUpdated
func GetLinksesByLinkUrlAndLinkRatingAndLinkUpdated(offset int, limit int, LinkUrl_ string, LinkRating_ int, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_rating = ? and link_updated = ?", LinkUrl_, LinkRating_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkRatingAndLinkRel Get Linkses via LinkUrlAndLinkRatingAndLinkRel
func GetLinksesByLinkUrlAndLinkRatingAndLinkRel(offset int, limit int, LinkUrl_ string, LinkRating_ int, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_rating = ? and link_rel = ?", LinkUrl_, LinkRating_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkRatingAndLinkNotes Get Linkses via LinkUrlAndLinkRatingAndLinkNotes
func GetLinksesByLinkUrlAndLinkRatingAndLinkNotes(offset int, limit int, LinkUrl_ string, LinkRating_ int, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_rating = ? and link_notes = ?", LinkUrl_, LinkRating_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkRatingAndLinkRss Get Linkses via LinkUrlAndLinkRatingAndLinkRss
func GetLinksesByLinkUrlAndLinkRatingAndLinkRss(offset int, limit int, LinkUrl_ string, LinkRating_ int, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_rating = ? and link_rss = ?", LinkUrl_, LinkRating_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkUpdatedAndLinkRel Get Linkses via LinkUrlAndLinkUpdatedAndLinkRel
func GetLinksesByLinkUrlAndLinkUpdatedAndLinkRel(offset int, limit int, LinkUrl_ string, LinkUpdated_ time.Time, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_updated = ? and link_rel = ?", LinkUrl_, LinkUpdated_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkUpdatedAndLinkNotes Get Linkses via LinkUrlAndLinkUpdatedAndLinkNotes
func GetLinksesByLinkUrlAndLinkUpdatedAndLinkNotes(offset int, limit int, LinkUrl_ string, LinkUpdated_ time.Time, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_updated = ? and link_notes = ?", LinkUrl_, LinkUpdated_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkUpdatedAndLinkRss Get Linkses via LinkUrlAndLinkUpdatedAndLinkRss
func GetLinksesByLinkUrlAndLinkUpdatedAndLinkRss(offset int, limit int, LinkUrl_ string, LinkUpdated_ time.Time, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_updated = ? and link_rss = ?", LinkUrl_, LinkUpdated_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkRelAndLinkNotes Get Linkses via LinkUrlAndLinkRelAndLinkNotes
func GetLinksesByLinkUrlAndLinkRelAndLinkNotes(offset int, limit int, LinkUrl_ string, LinkRel_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_rel = ? and link_notes = ?", LinkUrl_, LinkRel_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkRelAndLinkRss Get Linkses via LinkUrlAndLinkRelAndLinkRss
func GetLinksesByLinkUrlAndLinkRelAndLinkRss(offset int, limit int, LinkUrl_ string, LinkRel_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_rel = ? and link_rss = ?", LinkUrl_, LinkRel_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkNotesAndLinkRss Get Linkses via LinkUrlAndLinkNotesAndLinkRss
func GetLinksesByLinkUrlAndLinkNotesAndLinkRss(offset int, limit int, LinkUrl_ string, LinkNotes_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_notes = ? and link_rss = ?", LinkUrl_, LinkNotes_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkImageAndLinkTarget Get Linkses via LinkNameAndLinkImageAndLinkTarget
func GetLinksesByLinkNameAndLinkImageAndLinkTarget(offset int, limit int, LinkName_ string, LinkImage_ string, LinkTarget_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_image = ? and link_target = ?", LinkName_, LinkImage_, LinkTarget_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkImageAndLinkDescription Get Linkses via LinkNameAndLinkImageAndLinkDescription
func GetLinksesByLinkNameAndLinkImageAndLinkDescription(offset int, limit int, LinkName_ string, LinkImage_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_image = ? and link_description = ?", LinkName_, LinkImage_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkImageAndLinkVisible Get Linkses via LinkNameAndLinkImageAndLinkVisible
func GetLinksesByLinkNameAndLinkImageAndLinkVisible(offset int, limit int, LinkName_ string, LinkImage_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_image = ? and link_visible = ?", LinkName_, LinkImage_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkImageAndLinkOwner Get Linkses via LinkNameAndLinkImageAndLinkOwner
func GetLinksesByLinkNameAndLinkImageAndLinkOwner(offset int, limit int, LinkName_ string, LinkImage_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_image = ? and link_owner = ?", LinkName_, LinkImage_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkImageAndLinkRating Get Linkses via LinkNameAndLinkImageAndLinkRating
func GetLinksesByLinkNameAndLinkImageAndLinkRating(offset int, limit int, LinkName_ string, LinkImage_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_image = ? and link_rating = ?", LinkName_, LinkImage_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkImageAndLinkUpdated Get Linkses via LinkNameAndLinkImageAndLinkUpdated
func GetLinksesByLinkNameAndLinkImageAndLinkUpdated(offset int, limit int, LinkName_ string, LinkImage_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_image = ? and link_updated = ?", LinkName_, LinkImage_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkImageAndLinkRel Get Linkses via LinkNameAndLinkImageAndLinkRel
func GetLinksesByLinkNameAndLinkImageAndLinkRel(offset int, limit int, LinkName_ string, LinkImage_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_image = ? and link_rel = ?", LinkName_, LinkImage_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkImageAndLinkNotes Get Linkses via LinkNameAndLinkImageAndLinkNotes
func GetLinksesByLinkNameAndLinkImageAndLinkNotes(offset int, limit int, LinkName_ string, LinkImage_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_image = ? and link_notes = ?", LinkName_, LinkImage_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkImageAndLinkRss Get Linkses via LinkNameAndLinkImageAndLinkRss
func GetLinksesByLinkNameAndLinkImageAndLinkRss(offset int, limit int, LinkName_ string, LinkImage_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_image = ? and link_rss = ?", LinkName_, LinkImage_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkTargetAndLinkDescription Get Linkses via LinkNameAndLinkTargetAndLinkDescription
func GetLinksesByLinkNameAndLinkTargetAndLinkDescription(offset int, limit int, LinkName_ string, LinkTarget_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_target = ? and link_description = ?", LinkName_, LinkTarget_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkTargetAndLinkVisible Get Linkses via LinkNameAndLinkTargetAndLinkVisible
func GetLinksesByLinkNameAndLinkTargetAndLinkVisible(offset int, limit int, LinkName_ string, LinkTarget_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_target = ? and link_visible = ?", LinkName_, LinkTarget_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkTargetAndLinkOwner Get Linkses via LinkNameAndLinkTargetAndLinkOwner
func GetLinksesByLinkNameAndLinkTargetAndLinkOwner(offset int, limit int, LinkName_ string, LinkTarget_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_target = ? and link_owner = ?", LinkName_, LinkTarget_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkTargetAndLinkRating Get Linkses via LinkNameAndLinkTargetAndLinkRating
func GetLinksesByLinkNameAndLinkTargetAndLinkRating(offset int, limit int, LinkName_ string, LinkTarget_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_target = ? and link_rating = ?", LinkName_, LinkTarget_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkTargetAndLinkUpdated Get Linkses via LinkNameAndLinkTargetAndLinkUpdated
func GetLinksesByLinkNameAndLinkTargetAndLinkUpdated(offset int, limit int, LinkName_ string, LinkTarget_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_target = ? and link_updated = ?", LinkName_, LinkTarget_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkTargetAndLinkRel Get Linkses via LinkNameAndLinkTargetAndLinkRel
func GetLinksesByLinkNameAndLinkTargetAndLinkRel(offset int, limit int, LinkName_ string, LinkTarget_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_target = ? and link_rel = ?", LinkName_, LinkTarget_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkTargetAndLinkNotes Get Linkses via LinkNameAndLinkTargetAndLinkNotes
func GetLinksesByLinkNameAndLinkTargetAndLinkNotes(offset int, limit int, LinkName_ string, LinkTarget_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_target = ? and link_notes = ?", LinkName_, LinkTarget_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkTargetAndLinkRss Get Linkses via LinkNameAndLinkTargetAndLinkRss
func GetLinksesByLinkNameAndLinkTargetAndLinkRss(offset int, limit int, LinkName_ string, LinkTarget_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_target = ? and link_rss = ?", LinkName_, LinkTarget_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkDescriptionAndLinkVisible Get Linkses via LinkNameAndLinkDescriptionAndLinkVisible
func GetLinksesByLinkNameAndLinkDescriptionAndLinkVisible(offset int, limit int, LinkName_ string, LinkDescription_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_description = ? and link_visible = ?", LinkName_, LinkDescription_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkDescriptionAndLinkOwner Get Linkses via LinkNameAndLinkDescriptionAndLinkOwner
func GetLinksesByLinkNameAndLinkDescriptionAndLinkOwner(offset int, limit int, LinkName_ string, LinkDescription_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_description = ? and link_owner = ?", LinkName_, LinkDescription_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkDescriptionAndLinkRating Get Linkses via LinkNameAndLinkDescriptionAndLinkRating
func GetLinksesByLinkNameAndLinkDescriptionAndLinkRating(offset int, limit int, LinkName_ string, LinkDescription_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_description = ? and link_rating = ?", LinkName_, LinkDescription_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkDescriptionAndLinkUpdated Get Linkses via LinkNameAndLinkDescriptionAndLinkUpdated
func GetLinksesByLinkNameAndLinkDescriptionAndLinkUpdated(offset int, limit int, LinkName_ string, LinkDescription_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_description = ? and link_updated = ?", LinkName_, LinkDescription_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkDescriptionAndLinkRel Get Linkses via LinkNameAndLinkDescriptionAndLinkRel
func GetLinksesByLinkNameAndLinkDescriptionAndLinkRel(offset int, limit int, LinkName_ string, LinkDescription_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_description = ? and link_rel = ?", LinkName_, LinkDescription_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkDescriptionAndLinkNotes Get Linkses via LinkNameAndLinkDescriptionAndLinkNotes
func GetLinksesByLinkNameAndLinkDescriptionAndLinkNotes(offset int, limit int, LinkName_ string, LinkDescription_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_description = ? and link_notes = ?", LinkName_, LinkDescription_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkDescriptionAndLinkRss Get Linkses via LinkNameAndLinkDescriptionAndLinkRss
func GetLinksesByLinkNameAndLinkDescriptionAndLinkRss(offset int, limit int, LinkName_ string, LinkDescription_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_description = ? and link_rss = ?", LinkName_, LinkDescription_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkVisibleAndLinkOwner Get Linkses via LinkNameAndLinkVisibleAndLinkOwner
func GetLinksesByLinkNameAndLinkVisibleAndLinkOwner(offset int, limit int, LinkName_ string, LinkVisible_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_visible = ? and link_owner = ?", LinkName_, LinkVisible_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkVisibleAndLinkRating Get Linkses via LinkNameAndLinkVisibleAndLinkRating
func GetLinksesByLinkNameAndLinkVisibleAndLinkRating(offset int, limit int, LinkName_ string, LinkVisible_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_visible = ? and link_rating = ?", LinkName_, LinkVisible_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkVisibleAndLinkUpdated Get Linkses via LinkNameAndLinkVisibleAndLinkUpdated
func GetLinksesByLinkNameAndLinkVisibleAndLinkUpdated(offset int, limit int, LinkName_ string, LinkVisible_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_visible = ? and link_updated = ?", LinkName_, LinkVisible_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkVisibleAndLinkRel Get Linkses via LinkNameAndLinkVisibleAndLinkRel
func GetLinksesByLinkNameAndLinkVisibleAndLinkRel(offset int, limit int, LinkName_ string, LinkVisible_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_visible = ? and link_rel = ?", LinkName_, LinkVisible_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkVisibleAndLinkNotes Get Linkses via LinkNameAndLinkVisibleAndLinkNotes
func GetLinksesByLinkNameAndLinkVisibleAndLinkNotes(offset int, limit int, LinkName_ string, LinkVisible_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_visible = ? and link_notes = ?", LinkName_, LinkVisible_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkVisibleAndLinkRss Get Linkses via LinkNameAndLinkVisibleAndLinkRss
func GetLinksesByLinkNameAndLinkVisibleAndLinkRss(offset int, limit int, LinkName_ string, LinkVisible_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_visible = ? and link_rss = ?", LinkName_, LinkVisible_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkOwnerAndLinkRating Get Linkses via LinkNameAndLinkOwnerAndLinkRating
func GetLinksesByLinkNameAndLinkOwnerAndLinkRating(offset int, limit int, LinkName_ string, LinkOwner_ int64, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_owner = ? and link_rating = ?", LinkName_, LinkOwner_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkOwnerAndLinkUpdated Get Linkses via LinkNameAndLinkOwnerAndLinkUpdated
func GetLinksesByLinkNameAndLinkOwnerAndLinkUpdated(offset int, limit int, LinkName_ string, LinkOwner_ int64, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_owner = ? and link_updated = ?", LinkName_, LinkOwner_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkOwnerAndLinkRel Get Linkses via LinkNameAndLinkOwnerAndLinkRel
func GetLinksesByLinkNameAndLinkOwnerAndLinkRel(offset int, limit int, LinkName_ string, LinkOwner_ int64, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_owner = ? and link_rel = ?", LinkName_, LinkOwner_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkOwnerAndLinkNotes Get Linkses via LinkNameAndLinkOwnerAndLinkNotes
func GetLinksesByLinkNameAndLinkOwnerAndLinkNotes(offset int, limit int, LinkName_ string, LinkOwner_ int64, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_owner = ? and link_notes = ?", LinkName_, LinkOwner_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkOwnerAndLinkRss Get Linkses via LinkNameAndLinkOwnerAndLinkRss
func GetLinksesByLinkNameAndLinkOwnerAndLinkRss(offset int, limit int, LinkName_ string, LinkOwner_ int64, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_owner = ? and link_rss = ?", LinkName_, LinkOwner_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkRatingAndLinkUpdated Get Linkses via LinkNameAndLinkRatingAndLinkUpdated
func GetLinksesByLinkNameAndLinkRatingAndLinkUpdated(offset int, limit int, LinkName_ string, LinkRating_ int, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_rating = ? and link_updated = ?", LinkName_, LinkRating_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkRatingAndLinkRel Get Linkses via LinkNameAndLinkRatingAndLinkRel
func GetLinksesByLinkNameAndLinkRatingAndLinkRel(offset int, limit int, LinkName_ string, LinkRating_ int, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_rating = ? and link_rel = ?", LinkName_, LinkRating_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkRatingAndLinkNotes Get Linkses via LinkNameAndLinkRatingAndLinkNotes
func GetLinksesByLinkNameAndLinkRatingAndLinkNotes(offset int, limit int, LinkName_ string, LinkRating_ int, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_rating = ? and link_notes = ?", LinkName_, LinkRating_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkRatingAndLinkRss Get Linkses via LinkNameAndLinkRatingAndLinkRss
func GetLinksesByLinkNameAndLinkRatingAndLinkRss(offset int, limit int, LinkName_ string, LinkRating_ int, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_rating = ? and link_rss = ?", LinkName_, LinkRating_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkUpdatedAndLinkRel Get Linkses via LinkNameAndLinkUpdatedAndLinkRel
func GetLinksesByLinkNameAndLinkUpdatedAndLinkRel(offset int, limit int, LinkName_ string, LinkUpdated_ time.Time, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_updated = ? and link_rel = ?", LinkName_, LinkUpdated_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkUpdatedAndLinkNotes Get Linkses via LinkNameAndLinkUpdatedAndLinkNotes
func GetLinksesByLinkNameAndLinkUpdatedAndLinkNotes(offset int, limit int, LinkName_ string, LinkUpdated_ time.Time, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_updated = ? and link_notes = ?", LinkName_, LinkUpdated_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkUpdatedAndLinkRss Get Linkses via LinkNameAndLinkUpdatedAndLinkRss
func GetLinksesByLinkNameAndLinkUpdatedAndLinkRss(offset int, limit int, LinkName_ string, LinkUpdated_ time.Time, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_updated = ? and link_rss = ?", LinkName_, LinkUpdated_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkRelAndLinkNotes Get Linkses via LinkNameAndLinkRelAndLinkNotes
func GetLinksesByLinkNameAndLinkRelAndLinkNotes(offset int, limit int, LinkName_ string, LinkRel_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_rel = ? and link_notes = ?", LinkName_, LinkRel_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkRelAndLinkRss Get Linkses via LinkNameAndLinkRelAndLinkRss
func GetLinksesByLinkNameAndLinkRelAndLinkRss(offset int, limit int, LinkName_ string, LinkRel_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_rel = ? and link_rss = ?", LinkName_, LinkRel_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkNotesAndLinkRss Get Linkses via LinkNameAndLinkNotesAndLinkRss
func GetLinksesByLinkNameAndLinkNotesAndLinkRss(offset int, limit int, LinkName_ string, LinkNotes_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_notes = ? and link_rss = ?", LinkName_, LinkNotes_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkTargetAndLinkDescription Get Linkses via LinkImageAndLinkTargetAndLinkDescription
func GetLinksesByLinkImageAndLinkTargetAndLinkDescription(offset int, limit int, LinkImage_ string, LinkTarget_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_target = ? and link_description = ?", LinkImage_, LinkTarget_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkTargetAndLinkVisible Get Linkses via LinkImageAndLinkTargetAndLinkVisible
func GetLinksesByLinkImageAndLinkTargetAndLinkVisible(offset int, limit int, LinkImage_ string, LinkTarget_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_target = ? and link_visible = ?", LinkImage_, LinkTarget_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkTargetAndLinkOwner Get Linkses via LinkImageAndLinkTargetAndLinkOwner
func GetLinksesByLinkImageAndLinkTargetAndLinkOwner(offset int, limit int, LinkImage_ string, LinkTarget_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_target = ? and link_owner = ?", LinkImage_, LinkTarget_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkTargetAndLinkRating Get Linkses via LinkImageAndLinkTargetAndLinkRating
func GetLinksesByLinkImageAndLinkTargetAndLinkRating(offset int, limit int, LinkImage_ string, LinkTarget_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_target = ? and link_rating = ?", LinkImage_, LinkTarget_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkTargetAndLinkUpdated Get Linkses via LinkImageAndLinkTargetAndLinkUpdated
func GetLinksesByLinkImageAndLinkTargetAndLinkUpdated(offset int, limit int, LinkImage_ string, LinkTarget_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_target = ? and link_updated = ?", LinkImage_, LinkTarget_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkTargetAndLinkRel Get Linkses via LinkImageAndLinkTargetAndLinkRel
func GetLinksesByLinkImageAndLinkTargetAndLinkRel(offset int, limit int, LinkImage_ string, LinkTarget_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_target = ? and link_rel = ?", LinkImage_, LinkTarget_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkTargetAndLinkNotes Get Linkses via LinkImageAndLinkTargetAndLinkNotes
func GetLinksesByLinkImageAndLinkTargetAndLinkNotes(offset int, limit int, LinkImage_ string, LinkTarget_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_target = ? and link_notes = ?", LinkImage_, LinkTarget_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkTargetAndLinkRss Get Linkses via LinkImageAndLinkTargetAndLinkRss
func GetLinksesByLinkImageAndLinkTargetAndLinkRss(offset int, limit int, LinkImage_ string, LinkTarget_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_target = ? and link_rss = ?", LinkImage_, LinkTarget_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkDescriptionAndLinkVisible Get Linkses via LinkImageAndLinkDescriptionAndLinkVisible
func GetLinksesByLinkImageAndLinkDescriptionAndLinkVisible(offset int, limit int, LinkImage_ string, LinkDescription_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_description = ? and link_visible = ?", LinkImage_, LinkDescription_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkDescriptionAndLinkOwner Get Linkses via LinkImageAndLinkDescriptionAndLinkOwner
func GetLinksesByLinkImageAndLinkDescriptionAndLinkOwner(offset int, limit int, LinkImage_ string, LinkDescription_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_description = ? and link_owner = ?", LinkImage_, LinkDescription_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkDescriptionAndLinkRating Get Linkses via LinkImageAndLinkDescriptionAndLinkRating
func GetLinksesByLinkImageAndLinkDescriptionAndLinkRating(offset int, limit int, LinkImage_ string, LinkDescription_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_description = ? and link_rating = ?", LinkImage_, LinkDescription_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkDescriptionAndLinkUpdated Get Linkses via LinkImageAndLinkDescriptionAndLinkUpdated
func GetLinksesByLinkImageAndLinkDescriptionAndLinkUpdated(offset int, limit int, LinkImage_ string, LinkDescription_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_description = ? and link_updated = ?", LinkImage_, LinkDescription_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkDescriptionAndLinkRel Get Linkses via LinkImageAndLinkDescriptionAndLinkRel
func GetLinksesByLinkImageAndLinkDescriptionAndLinkRel(offset int, limit int, LinkImage_ string, LinkDescription_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_description = ? and link_rel = ?", LinkImage_, LinkDescription_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkDescriptionAndLinkNotes Get Linkses via LinkImageAndLinkDescriptionAndLinkNotes
func GetLinksesByLinkImageAndLinkDescriptionAndLinkNotes(offset int, limit int, LinkImage_ string, LinkDescription_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_description = ? and link_notes = ?", LinkImage_, LinkDescription_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkDescriptionAndLinkRss Get Linkses via LinkImageAndLinkDescriptionAndLinkRss
func GetLinksesByLinkImageAndLinkDescriptionAndLinkRss(offset int, limit int, LinkImage_ string, LinkDescription_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_description = ? and link_rss = ?", LinkImage_, LinkDescription_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkVisibleAndLinkOwner Get Linkses via LinkImageAndLinkVisibleAndLinkOwner
func GetLinksesByLinkImageAndLinkVisibleAndLinkOwner(offset int, limit int, LinkImage_ string, LinkVisible_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_visible = ? and link_owner = ?", LinkImage_, LinkVisible_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkVisibleAndLinkRating Get Linkses via LinkImageAndLinkVisibleAndLinkRating
func GetLinksesByLinkImageAndLinkVisibleAndLinkRating(offset int, limit int, LinkImage_ string, LinkVisible_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_visible = ? and link_rating = ?", LinkImage_, LinkVisible_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkVisibleAndLinkUpdated Get Linkses via LinkImageAndLinkVisibleAndLinkUpdated
func GetLinksesByLinkImageAndLinkVisibleAndLinkUpdated(offset int, limit int, LinkImage_ string, LinkVisible_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_visible = ? and link_updated = ?", LinkImage_, LinkVisible_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkVisibleAndLinkRel Get Linkses via LinkImageAndLinkVisibleAndLinkRel
func GetLinksesByLinkImageAndLinkVisibleAndLinkRel(offset int, limit int, LinkImage_ string, LinkVisible_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_visible = ? and link_rel = ?", LinkImage_, LinkVisible_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkVisibleAndLinkNotes Get Linkses via LinkImageAndLinkVisibleAndLinkNotes
func GetLinksesByLinkImageAndLinkVisibleAndLinkNotes(offset int, limit int, LinkImage_ string, LinkVisible_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_visible = ? and link_notes = ?", LinkImage_, LinkVisible_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkVisibleAndLinkRss Get Linkses via LinkImageAndLinkVisibleAndLinkRss
func GetLinksesByLinkImageAndLinkVisibleAndLinkRss(offset int, limit int, LinkImage_ string, LinkVisible_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_visible = ? and link_rss = ?", LinkImage_, LinkVisible_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkOwnerAndLinkRating Get Linkses via LinkImageAndLinkOwnerAndLinkRating
func GetLinksesByLinkImageAndLinkOwnerAndLinkRating(offset int, limit int, LinkImage_ string, LinkOwner_ int64, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_owner = ? and link_rating = ?", LinkImage_, LinkOwner_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkOwnerAndLinkUpdated Get Linkses via LinkImageAndLinkOwnerAndLinkUpdated
func GetLinksesByLinkImageAndLinkOwnerAndLinkUpdated(offset int, limit int, LinkImage_ string, LinkOwner_ int64, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_owner = ? and link_updated = ?", LinkImage_, LinkOwner_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkOwnerAndLinkRel Get Linkses via LinkImageAndLinkOwnerAndLinkRel
func GetLinksesByLinkImageAndLinkOwnerAndLinkRel(offset int, limit int, LinkImage_ string, LinkOwner_ int64, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_owner = ? and link_rel = ?", LinkImage_, LinkOwner_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkOwnerAndLinkNotes Get Linkses via LinkImageAndLinkOwnerAndLinkNotes
func GetLinksesByLinkImageAndLinkOwnerAndLinkNotes(offset int, limit int, LinkImage_ string, LinkOwner_ int64, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_owner = ? and link_notes = ?", LinkImage_, LinkOwner_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkOwnerAndLinkRss Get Linkses via LinkImageAndLinkOwnerAndLinkRss
func GetLinksesByLinkImageAndLinkOwnerAndLinkRss(offset int, limit int, LinkImage_ string, LinkOwner_ int64, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_owner = ? and link_rss = ?", LinkImage_, LinkOwner_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkRatingAndLinkUpdated Get Linkses via LinkImageAndLinkRatingAndLinkUpdated
func GetLinksesByLinkImageAndLinkRatingAndLinkUpdated(offset int, limit int, LinkImage_ string, LinkRating_ int, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_rating = ? and link_updated = ?", LinkImage_, LinkRating_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkRatingAndLinkRel Get Linkses via LinkImageAndLinkRatingAndLinkRel
func GetLinksesByLinkImageAndLinkRatingAndLinkRel(offset int, limit int, LinkImage_ string, LinkRating_ int, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_rating = ? and link_rel = ?", LinkImage_, LinkRating_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkRatingAndLinkNotes Get Linkses via LinkImageAndLinkRatingAndLinkNotes
func GetLinksesByLinkImageAndLinkRatingAndLinkNotes(offset int, limit int, LinkImage_ string, LinkRating_ int, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_rating = ? and link_notes = ?", LinkImage_, LinkRating_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkRatingAndLinkRss Get Linkses via LinkImageAndLinkRatingAndLinkRss
func GetLinksesByLinkImageAndLinkRatingAndLinkRss(offset int, limit int, LinkImage_ string, LinkRating_ int, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_rating = ? and link_rss = ?", LinkImage_, LinkRating_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkUpdatedAndLinkRel Get Linkses via LinkImageAndLinkUpdatedAndLinkRel
func GetLinksesByLinkImageAndLinkUpdatedAndLinkRel(offset int, limit int, LinkImage_ string, LinkUpdated_ time.Time, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_updated = ? and link_rel = ?", LinkImage_, LinkUpdated_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkUpdatedAndLinkNotes Get Linkses via LinkImageAndLinkUpdatedAndLinkNotes
func GetLinksesByLinkImageAndLinkUpdatedAndLinkNotes(offset int, limit int, LinkImage_ string, LinkUpdated_ time.Time, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_updated = ? and link_notes = ?", LinkImage_, LinkUpdated_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkUpdatedAndLinkRss Get Linkses via LinkImageAndLinkUpdatedAndLinkRss
func GetLinksesByLinkImageAndLinkUpdatedAndLinkRss(offset int, limit int, LinkImage_ string, LinkUpdated_ time.Time, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_updated = ? and link_rss = ?", LinkImage_, LinkUpdated_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkRelAndLinkNotes Get Linkses via LinkImageAndLinkRelAndLinkNotes
func GetLinksesByLinkImageAndLinkRelAndLinkNotes(offset int, limit int, LinkImage_ string, LinkRel_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_rel = ? and link_notes = ?", LinkImage_, LinkRel_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkRelAndLinkRss Get Linkses via LinkImageAndLinkRelAndLinkRss
func GetLinksesByLinkImageAndLinkRelAndLinkRss(offset int, limit int, LinkImage_ string, LinkRel_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_rel = ? and link_rss = ?", LinkImage_, LinkRel_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkNotesAndLinkRss Get Linkses via LinkImageAndLinkNotesAndLinkRss
func GetLinksesByLinkImageAndLinkNotesAndLinkRss(offset int, limit int, LinkImage_ string, LinkNotes_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_notes = ? and link_rss = ?", LinkImage_, LinkNotes_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkDescriptionAndLinkVisible Get Linkses via LinkTargetAndLinkDescriptionAndLinkVisible
func GetLinksesByLinkTargetAndLinkDescriptionAndLinkVisible(offset int, limit int, LinkTarget_ string, LinkDescription_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_description = ? and link_visible = ?", LinkTarget_, LinkDescription_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkDescriptionAndLinkOwner Get Linkses via LinkTargetAndLinkDescriptionAndLinkOwner
func GetLinksesByLinkTargetAndLinkDescriptionAndLinkOwner(offset int, limit int, LinkTarget_ string, LinkDescription_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_description = ? and link_owner = ?", LinkTarget_, LinkDescription_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkDescriptionAndLinkRating Get Linkses via LinkTargetAndLinkDescriptionAndLinkRating
func GetLinksesByLinkTargetAndLinkDescriptionAndLinkRating(offset int, limit int, LinkTarget_ string, LinkDescription_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_description = ? and link_rating = ?", LinkTarget_, LinkDescription_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkDescriptionAndLinkUpdated Get Linkses via LinkTargetAndLinkDescriptionAndLinkUpdated
func GetLinksesByLinkTargetAndLinkDescriptionAndLinkUpdated(offset int, limit int, LinkTarget_ string, LinkDescription_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_description = ? and link_updated = ?", LinkTarget_, LinkDescription_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkDescriptionAndLinkRel Get Linkses via LinkTargetAndLinkDescriptionAndLinkRel
func GetLinksesByLinkTargetAndLinkDescriptionAndLinkRel(offset int, limit int, LinkTarget_ string, LinkDescription_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_description = ? and link_rel = ?", LinkTarget_, LinkDescription_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkDescriptionAndLinkNotes Get Linkses via LinkTargetAndLinkDescriptionAndLinkNotes
func GetLinksesByLinkTargetAndLinkDescriptionAndLinkNotes(offset int, limit int, LinkTarget_ string, LinkDescription_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_description = ? and link_notes = ?", LinkTarget_, LinkDescription_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkDescriptionAndLinkRss Get Linkses via LinkTargetAndLinkDescriptionAndLinkRss
func GetLinksesByLinkTargetAndLinkDescriptionAndLinkRss(offset int, limit int, LinkTarget_ string, LinkDescription_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_description = ? and link_rss = ?", LinkTarget_, LinkDescription_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkVisibleAndLinkOwner Get Linkses via LinkTargetAndLinkVisibleAndLinkOwner
func GetLinksesByLinkTargetAndLinkVisibleAndLinkOwner(offset int, limit int, LinkTarget_ string, LinkVisible_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_visible = ? and link_owner = ?", LinkTarget_, LinkVisible_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkVisibleAndLinkRating Get Linkses via LinkTargetAndLinkVisibleAndLinkRating
func GetLinksesByLinkTargetAndLinkVisibleAndLinkRating(offset int, limit int, LinkTarget_ string, LinkVisible_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_visible = ? and link_rating = ?", LinkTarget_, LinkVisible_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkVisibleAndLinkUpdated Get Linkses via LinkTargetAndLinkVisibleAndLinkUpdated
func GetLinksesByLinkTargetAndLinkVisibleAndLinkUpdated(offset int, limit int, LinkTarget_ string, LinkVisible_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_visible = ? and link_updated = ?", LinkTarget_, LinkVisible_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkVisibleAndLinkRel Get Linkses via LinkTargetAndLinkVisibleAndLinkRel
func GetLinksesByLinkTargetAndLinkVisibleAndLinkRel(offset int, limit int, LinkTarget_ string, LinkVisible_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_visible = ? and link_rel = ?", LinkTarget_, LinkVisible_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkVisibleAndLinkNotes Get Linkses via LinkTargetAndLinkVisibleAndLinkNotes
func GetLinksesByLinkTargetAndLinkVisibleAndLinkNotes(offset int, limit int, LinkTarget_ string, LinkVisible_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_visible = ? and link_notes = ?", LinkTarget_, LinkVisible_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkVisibleAndLinkRss Get Linkses via LinkTargetAndLinkVisibleAndLinkRss
func GetLinksesByLinkTargetAndLinkVisibleAndLinkRss(offset int, limit int, LinkTarget_ string, LinkVisible_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_visible = ? and link_rss = ?", LinkTarget_, LinkVisible_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkOwnerAndLinkRating Get Linkses via LinkTargetAndLinkOwnerAndLinkRating
func GetLinksesByLinkTargetAndLinkOwnerAndLinkRating(offset int, limit int, LinkTarget_ string, LinkOwner_ int64, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_owner = ? and link_rating = ?", LinkTarget_, LinkOwner_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkOwnerAndLinkUpdated Get Linkses via LinkTargetAndLinkOwnerAndLinkUpdated
func GetLinksesByLinkTargetAndLinkOwnerAndLinkUpdated(offset int, limit int, LinkTarget_ string, LinkOwner_ int64, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_owner = ? and link_updated = ?", LinkTarget_, LinkOwner_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkOwnerAndLinkRel Get Linkses via LinkTargetAndLinkOwnerAndLinkRel
func GetLinksesByLinkTargetAndLinkOwnerAndLinkRel(offset int, limit int, LinkTarget_ string, LinkOwner_ int64, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_owner = ? and link_rel = ?", LinkTarget_, LinkOwner_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkOwnerAndLinkNotes Get Linkses via LinkTargetAndLinkOwnerAndLinkNotes
func GetLinksesByLinkTargetAndLinkOwnerAndLinkNotes(offset int, limit int, LinkTarget_ string, LinkOwner_ int64, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_owner = ? and link_notes = ?", LinkTarget_, LinkOwner_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkOwnerAndLinkRss Get Linkses via LinkTargetAndLinkOwnerAndLinkRss
func GetLinksesByLinkTargetAndLinkOwnerAndLinkRss(offset int, limit int, LinkTarget_ string, LinkOwner_ int64, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_owner = ? and link_rss = ?", LinkTarget_, LinkOwner_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkRatingAndLinkUpdated Get Linkses via LinkTargetAndLinkRatingAndLinkUpdated
func GetLinksesByLinkTargetAndLinkRatingAndLinkUpdated(offset int, limit int, LinkTarget_ string, LinkRating_ int, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_rating = ? and link_updated = ?", LinkTarget_, LinkRating_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkRatingAndLinkRel Get Linkses via LinkTargetAndLinkRatingAndLinkRel
func GetLinksesByLinkTargetAndLinkRatingAndLinkRel(offset int, limit int, LinkTarget_ string, LinkRating_ int, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_rating = ? and link_rel = ?", LinkTarget_, LinkRating_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkRatingAndLinkNotes Get Linkses via LinkTargetAndLinkRatingAndLinkNotes
func GetLinksesByLinkTargetAndLinkRatingAndLinkNotes(offset int, limit int, LinkTarget_ string, LinkRating_ int, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_rating = ? and link_notes = ?", LinkTarget_, LinkRating_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkRatingAndLinkRss Get Linkses via LinkTargetAndLinkRatingAndLinkRss
func GetLinksesByLinkTargetAndLinkRatingAndLinkRss(offset int, limit int, LinkTarget_ string, LinkRating_ int, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_rating = ? and link_rss = ?", LinkTarget_, LinkRating_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkUpdatedAndLinkRel Get Linkses via LinkTargetAndLinkUpdatedAndLinkRel
func GetLinksesByLinkTargetAndLinkUpdatedAndLinkRel(offset int, limit int, LinkTarget_ string, LinkUpdated_ time.Time, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_updated = ? and link_rel = ?", LinkTarget_, LinkUpdated_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkUpdatedAndLinkNotes Get Linkses via LinkTargetAndLinkUpdatedAndLinkNotes
func GetLinksesByLinkTargetAndLinkUpdatedAndLinkNotes(offset int, limit int, LinkTarget_ string, LinkUpdated_ time.Time, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_updated = ? and link_notes = ?", LinkTarget_, LinkUpdated_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkUpdatedAndLinkRss Get Linkses via LinkTargetAndLinkUpdatedAndLinkRss
func GetLinksesByLinkTargetAndLinkUpdatedAndLinkRss(offset int, limit int, LinkTarget_ string, LinkUpdated_ time.Time, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_updated = ? and link_rss = ?", LinkTarget_, LinkUpdated_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkRelAndLinkNotes Get Linkses via LinkTargetAndLinkRelAndLinkNotes
func GetLinksesByLinkTargetAndLinkRelAndLinkNotes(offset int, limit int, LinkTarget_ string, LinkRel_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_rel = ? and link_notes = ?", LinkTarget_, LinkRel_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkRelAndLinkRss Get Linkses via LinkTargetAndLinkRelAndLinkRss
func GetLinksesByLinkTargetAndLinkRelAndLinkRss(offset int, limit int, LinkTarget_ string, LinkRel_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_rel = ? and link_rss = ?", LinkTarget_, LinkRel_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkNotesAndLinkRss Get Linkses via LinkTargetAndLinkNotesAndLinkRss
func GetLinksesByLinkTargetAndLinkNotesAndLinkRss(offset int, limit int, LinkTarget_ string, LinkNotes_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_notes = ? and link_rss = ?", LinkTarget_, LinkNotes_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkVisibleAndLinkOwner Get Linkses via LinkDescriptionAndLinkVisibleAndLinkOwner
func GetLinksesByLinkDescriptionAndLinkVisibleAndLinkOwner(offset int, limit int, LinkDescription_ string, LinkVisible_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_visible = ? and link_owner = ?", LinkDescription_, LinkVisible_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRating Get Linkses via LinkDescriptionAndLinkVisibleAndLinkRating
func GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRating(offset int, limit int, LinkDescription_ string, LinkVisible_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_visible = ? and link_rating = ?", LinkDescription_, LinkVisible_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkVisibleAndLinkUpdated Get Linkses via LinkDescriptionAndLinkVisibleAndLinkUpdated
func GetLinksesByLinkDescriptionAndLinkVisibleAndLinkUpdated(offset int, limit int, LinkDescription_ string, LinkVisible_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_visible = ? and link_updated = ?", LinkDescription_, LinkVisible_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRel Get Linkses via LinkDescriptionAndLinkVisibleAndLinkRel
func GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRel(offset int, limit int, LinkDescription_ string, LinkVisible_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_visible = ? and link_rel = ?", LinkDescription_, LinkVisible_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkVisibleAndLinkNotes Get Linkses via LinkDescriptionAndLinkVisibleAndLinkNotes
func GetLinksesByLinkDescriptionAndLinkVisibleAndLinkNotes(offset int, limit int, LinkDescription_ string, LinkVisible_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_visible = ? and link_notes = ?", LinkDescription_, LinkVisible_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRss Get Linkses via LinkDescriptionAndLinkVisibleAndLinkRss
func GetLinksesByLinkDescriptionAndLinkVisibleAndLinkRss(offset int, limit int, LinkDescription_ string, LinkVisible_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_visible = ? and link_rss = ?", LinkDescription_, LinkVisible_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRating Get Linkses via LinkDescriptionAndLinkOwnerAndLinkRating
func GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRating(offset int, limit int, LinkDescription_ string, LinkOwner_ int64, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_owner = ? and link_rating = ?", LinkDescription_, LinkOwner_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkOwnerAndLinkUpdated Get Linkses via LinkDescriptionAndLinkOwnerAndLinkUpdated
func GetLinksesByLinkDescriptionAndLinkOwnerAndLinkUpdated(offset int, limit int, LinkDescription_ string, LinkOwner_ int64, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_owner = ? and link_updated = ?", LinkDescription_, LinkOwner_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRel Get Linkses via LinkDescriptionAndLinkOwnerAndLinkRel
func GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRel(offset int, limit int, LinkDescription_ string, LinkOwner_ int64, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_owner = ? and link_rel = ?", LinkDescription_, LinkOwner_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkOwnerAndLinkNotes Get Linkses via LinkDescriptionAndLinkOwnerAndLinkNotes
func GetLinksesByLinkDescriptionAndLinkOwnerAndLinkNotes(offset int, limit int, LinkDescription_ string, LinkOwner_ int64, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_owner = ? and link_notes = ?", LinkDescription_, LinkOwner_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRss Get Linkses via LinkDescriptionAndLinkOwnerAndLinkRss
func GetLinksesByLinkDescriptionAndLinkOwnerAndLinkRss(offset int, limit int, LinkDescription_ string, LinkOwner_ int64, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_owner = ? and link_rss = ?", LinkDescription_, LinkOwner_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkRatingAndLinkUpdated Get Linkses via LinkDescriptionAndLinkRatingAndLinkUpdated
func GetLinksesByLinkDescriptionAndLinkRatingAndLinkUpdated(offset int, limit int, LinkDescription_ string, LinkRating_ int, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_rating = ? and link_updated = ?", LinkDescription_, LinkRating_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkRatingAndLinkRel Get Linkses via LinkDescriptionAndLinkRatingAndLinkRel
func GetLinksesByLinkDescriptionAndLinkRatingAndLinkRel(offset int, limit int, LinkDescription_ string, LinkRating_ int, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_rating = ? and link_rel = ?", LinkDescription_, LinkRating_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkRatingAndLinkNotes Get Linkses via LinkDescriptionAndLinkRatingAndLinkNotes
func GetLinksesByLinkDescriptionAndLinkRatingAndLinkNotes(offset int, limit int, LinkDescription_ string, LinkRating_ int, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_rating = ? and link_notes = ?", LinkDescription_, LinkRating_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkRatingAndLinkRss Get Linkses via LinkDescriptionAndLinkRatingAndLinkRss
func GetLinksesByLinkDescriptionAndLinkRatingAndLinkRss(offset int, limit int, LinkDescription_ string, LinkRating_ int, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_rating = ? and link_rss = ?", LinkDescription_, LinkRating_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkRel Get Linkses via LinkDescriptionAndLinkUpdatedAndLinkRel
func GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkRel(offset int, limit int, LinkDescription_ string, LinkUpdated_ time.Time, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_updated = ? and link_rel = ?", LinkDescription_, LinkUpdated_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkNotes Get Linkses via LinkDescriptionAndLinkUpdatedAndLinkNotes
func GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkNotes(offset int, limit int, LinkDescription_ string, LinkUpdated_ time.Time, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_updated = ? and link_notes = ?", LinkDescription_, LinkUpdated_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkRss Get Linkses via LinkDescriptionAndLinkUpdatedAndLinkRss
func GetLinksesByLinkDescriptionAndLinkUpdatedAndLinkRss(offset int, limit int, LinkDescription_ string, LinkUpdated_ time.Time, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_updated = ? and link_rss = ?", LinkDescription_, LinkUpdated_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkRelAndLinkNotes Get Linkses via LinkDescriptionAndLinkRelAndLinkNotes
func GetLinksesByLinkDescriptionAndLinkRelAndLinkNotes(offset int, limit int, LinkDescription_ string, LinkRel_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_rel = ? and link_notes = ?", LinkDescription_, LinkRel_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkRelAndLinkRss Get Linkses via LinkDescriptionAndLinkRelAndLinkRss
func GetLinksesByLinkDescriptionAndLinkRelAndLinkRss(offset int, limit int, LinkDescription_ string, LinkRel_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_rel = ? and link_rss = ?", LinkDescription_, LinkRel_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkNotesAndLinkRss Get Linkses via LinkDescriptionAndLinkNotesAndLinkRss
func GetLinksesByLinkDescriptionAndLinkNotesAndLinkRss(offset int, limit int, LinkDescription_ string, LinkNotes_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_notes = ? and link_rss = ?", LinkDescription_, LinkNotes_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkOwnerAndLinkRating Get Linkses via LinkVisibleAndLinkOwnerAndLinkRating
func GetLinksesByLinkVisibleAndLinkOwnerAndLinkRating(offset int, limit int, LinkVisible_ string, LinkOwner_ int64, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_owner = ? and link_rating = ?", LinkVisible_, LinkOwner_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkOwnerAndLinkUpdated Get Linkses via LinkVisibleAndLinkOwnerAndLinkUpdated
func GetLinksesByLinkVisibleAndLinkOwnerAndLinkUpdated(offset int, limit int, LinkVisible_ string, LinkOwner_ int64, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_owner = ? and link_updated = ?", LinkVisible_, LinkOwner_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkOwnerAndLinkRel Get Linkses via LinkVisibleAndLinkOwnerAndLinkRel
func GetLinksesByLinkVisibleAndLinkOwnerAndLinkRel(offset int, limit int, LinkVisible_ string, LinkOwner_ int64, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_owner = ? and link_rel = ?", LinkVisible_, LinkOwner_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkOwnerAndLinkNotes Get Linkses via LinkVisibleAndLinkOwnerAndLinkNotes
func GetLinksesByLinkVisibleAndLinkOwnerAndLinkNotes(offset int, limit int, LinkVisible_ string, LinkOwner_ int64, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_owner = ? and link_notes = ?", LinkVisible_, LinkOwner_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkOwnerAndLinkRss Get Linkses via LinkVisibleAndLinkOwnerAndLinkRss
func GetLinksesByLinkVisibleAndLinkOwnerAndLinkRss(offset int, limit int, LinkVisible_ string, LinkOwner_ int64, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_owner = ? and link_rss = ?", LinkVisible_, LinkOwner_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkRatingAndLinkUpdated Get Linkses via LinkVisibleAndLinkRatingAndLinkUpdated
func GetLinksesByLinkVisibleAndLinkRatingAndLinkUpdated(offset int, limit int, LinkVisible_ string, LinkRating_ int, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_rating = ? and link_updated = ?", LinkVisible_, LinkRating_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkRatingAndLinkRel Get Linkses via LinkVisibleAndLinkRatingAndLinkRel
func GetLinksesByLinkVisibleAndLinkRatingAndLinkRel(offset int, limit int, LinkVisible_ string, LinkRating_ int, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_rating = ? and link_rel = ?", LinkVisible_, LinkRating_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkRatingAndLinkNotes Get Linkses via LinkVisibleAndLinkRatingAndLinkNotes
func GetLinksesByLinkVisibleAndLinkRatingAndLinkNotes(offset int, limit int, LinkVisible_ string, LinkRating_ int, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_rating = ? and link_notes = ?", LinkVisible_, LinkRating_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkRatingAndLinkRss Get Linkses via LinkVisibleAndLinkRatingAndLinkRss
func GetLinksesByLinkVisibleAndLinkRatingAndLinkRss(offset int, limit int, LinkVisible_ string, LinkRating_ int, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_rating = ? and link_rss = ?", LinkVisible_, LinkRating_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkUpdatedAndLinkRel Get Linkses via LinkVisibleAndLinkUpdatedAndLinkRel
func GetLinksesByLinkVisibleAndLinkUpdatedAndLinkRel(offset int, limit int, LinkVisible_ string, LinkUpdated_ time.Time, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_updated = ? and link_rel = ?", LinkVisible_, LinkUpdated_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkUpdatedAndLinkNotes Get Linkses via LinkVisibleAndLinkUpdatedAndLinkNotes
func GetLinksesByLinkVisibleAndLinkUpdatedAndLinkNotes(offset int, limit int, LinkVisible_ string, LinkUpdated_ time.Time, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_updated = ? and link_notes = ?", LinkVisible_, LinkUpdated_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkUpdatedAndLinkRss Get Linkses via LinkVisibleAndLinkUpdatedAndLinkRss
func GetLinksesByLinkVisibleAndLinkUpdatedAndLinkRss(offset int, limit int, LinkVisible_ string, LinkUpdated_ time.Time, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_updated = ? and link_rss = ?", LinkVisible_, LinkUpdated_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkRelAndLinkNotes Get Linkses via LinkVisibleAndLinkRelAndLinkNotes
func GetLinksesByLinkVisibleAndLinkRelAndLinkNotes(offset int, limit int, LinkVisible_ string, LinkRel_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_rel = ? and link_notes = ?", LinkVisible_, LinkRel_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkRelAndLinkRss Get Linkses via LinkVisibleAndLinkRelAndLinkRss
func GetLinksesByLinkVisibleAndLinkRelAndLinkRss(offset int, limit int, LinkVisible_ string, LinkRel_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_rel = ? and link_rss = ?", LinkVisible_, LinkRel_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkNotesAndLinkRss Get Linkses via LinkVisibleAndLinkNotesAndLinkRss
func GetLinksesByLinkVisibleAndLinkNotesAndLinkRss(offset int, limit int, LinkVisible_ string, LinkNotes_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_notes = ? and link_rss = ?", LinkVisible_, LinkNotes_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkRatingAndLinkUpdated Get Linkses via LinkOwnerAndLinkRatingAndLinkUpdated
func GetLinksesByLinkOwnerAndLinkRatingAndLinkUpdated(offset int, limit int, LinkOwner_ int64, LinkRating_ int, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_rating = ? and link_updated = ?", LinkOwner_, LinkRating_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkRatingAndLinkRel Get Linkses via LinkOwnerAndLinkRatingAndLinkRel
func GetLinksesByLinkOwnerAndLinkRatingAndLinkRel(offset int, limit int, LinkOwner_ int64, LinkRating_ int, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_rating = ? and link_rel = ?", LinkOwner_, LinkRating_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkRatingAndLinkNotes Get Linkses via LinkOwnerAndLinkRatingAndLinkNotes
func GetLinksesByLinkOwnerAndLinkRatingAndLinkNotes(offset int, limit int, LinkOwner_ int64, LinkRating_ int, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_rating = ? and link_notes = ?", LinkOwner_, LinkRating_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkRatingAndLinkRss Get Linkses via LinkOwnerAndLinkRatingAndLinkRss
func GetLinksesByLinkOwnerAndLinkRatingAndLinkRss(offset int, limit int, LinkOwner_ int64, LinkRating_ int, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_rating = ? and link_rss = ?", LinkOwner_, LinkRating_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkUpdatedAndLinkRel Get Linkses via LinkOwnerAndLinkUpdatedAndLinkRel
func GetLinksesByLinkOwnerAndLinkUpdatedAndLinkRel(offset int, limit int, LinkOwner_ int64, LinkUpdated_ time.Time, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_updated = ? and link_rel = ?", LinkOwner_, LinkUpdated_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkUpdatedAndLinkNotes Get Linkses via LinkOwnerAndLinkUpdatedAndLinkNotes
func GetLinksesByLinkOwnerAndLinkUpdatedAndLinkNotes(offset int, limit int, LinkOwner_ int64, LinkUpdated_ time.Time, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_updated = ? and link_notes = ?", LinkOwner_, LinkUpdated_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkUpdatedAndLinkRss Get Linkses via LinkOwnerAndLinkUpdatedAndLinkRss
func GetLinksesByLinkOwnerAndLinkUpdatedAndLinkRss(offset int, limit int, LinkOwner_ int64, LinkUpdated_ time.Time, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_updated = ? and link_rss = ?", LinkOwner_, LinkUpdated_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkRelAndLinkNotes Get Linkses via LinkOwnerAndLinkRelAndLinkNotes
func GetLinksesByLinkOwnerAndLinkRelAndLinkNotes(offset int, limit int, LinkOwner_ int64, LinkRel_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_rel = ? and link_notes = ?", LinkOwner_, LinkRel_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkRelAndLinkRss Get Linkses via LinkOwnerAndLinkRelAndLinkRss
func GetLinksesByLinkOwnerAndLinkRelAndLinkRss(offset int, limit int, LinkOwner_ int64, LinkRel_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_rel = ? and link_rss = ?", LinkOwner_, LinkRel_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkNotesAndLinkRss Get Linkses via LinkOwnerAndLinkNotesAndLinkRss
func GetLinksesByLinkOwnerAndLinkNotesAndLinkRss(offset int, limit int, LinkOwner_ int64, LinkNotes_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_notes = ? and link_rss = ?", LinkOwner_, LinkNotes_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRatingAndLinkUpdatedAndLinkRel Get Linkses via LinkRatingAndLinkUpdatedAndLinkRel
func GetLinksesByLinkRatingAndLinkUpdatedAndLinkRel(offset int, limit int, LinkRating_ int, LinkUpdated_ time.Time, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rating = ? and link_updated = ? and link_rel = ?", LinkRating_, LinkUpdated_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRatingAndLinkUpdatedAndLinkNotes Get Linkses via LinkRatingAndLinkUpdatedAndLinkNotes
func GetLinksesByLinkRatingAndLinkUpdatedAndLinkNotes(offset int, limit int, LinkRating_ int, LinkUpdated_ time.Time, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rating = ? and link_updated = ? and link_notes = ?", LinkRating_, LinkUpdated_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRatingAndLinkUpdatedAndLinkRss Get Linkses via LinkRatingAndLinkUpdatedAndLinkRss
func GetLinksesByLinkRatingAndLinkUpdatedAndLinkRss(offset int, limit int, LinkRating_ int, LinkUpdated_ time.Time, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rating = ? and link_updated = ? and link_rss = ?", LinkRating_, LinkUpdated_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRatingAndLinkRelAndLinkNotes Get Linkses via LinkRatingAndLinkRelAndLinkNotes
func GetLinksesByLinkRatingAndLinkRelAndLinkNotes(offset int, limit int, LinkRating_ int, LinkRel_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rating = ? and link_rel = ? and link_notes = ?", LinkRating_, LinkRel_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRatingAndLinkRelAndLinkRss Get Linkses via LinkRatingAndLinkRelAndLinkRss
func GetLinksesByLinkRatingAndLinkRelAndLinkRss(offset int, limit int, LinkRating_ int, LinkRel_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rating = ? and link_rel = ? and link_rss = ?", LinkRating_, LinkRel_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRatingAndLinkNotesAndLinkRss Get Linkses via LinkRatingAndLinkNotesAndLinkRss
func GetLinksesByLinkRatingAndLinkNotesAndLinkRss(offset int, limit int, LinkRating_ int, LinkNotes_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rating = ? and link_notes = ? and link_rss = ?", LinkRating_, LinkNotes_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUpdatedAndLinkRelAndLinkNotes Get Linkses via LinkUpdatedAndLinkRelAndLinkNotes
func GetLinksesByLinkUpdatedAndLinkRelAndLinkNotes(offset int, limit int, LinkUpdated_ time.Time, LinkRel_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_updated = ? and link_rel = ? and link_notes = ?", LinkUpdated_, LinkRel_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUpdatedAndLinkRelAndLinkRss Get Linkses via LinkUpdatedAndLinkRelAndLinkRss
func GetLinksesByLinkUpdatedAndLinkRelAndLinkRss(offset int, limit int, LinkUpdated_ time.Time, LinkRel_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_updated = ? and link_rel = ? and link_rss = ?", LinkUpdated_, LinkRel_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUpdatedAndLinkNotesAndLinkRss Get Linkses via LinkUpdatedAndLinkNotesAndLinkRss
func GetLinksesByLinkUpdatedAndLinkNotesAndLinkRss(offset int, limit int, LinkUpdated_ time.Time, LinkNotes_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_updated = ? and link_notes = ? and link_rss = ?", LinkUpdated_, LinkNotes_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRelAndLinkNotesAndLinkRss Get Linkses via LinkRelAndLinkNotesAndLinkRss
func GetLinksesByLinkRelAndLinkNotesAndLinkRss(offset int, limit int, LinkRel_ string, LinkNotes_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rel = ? and link_notes = ? and link_rss = ?", LinkRel_, LinkNotes_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUrl Get Linkses via LinkIdAndLinkUrl
func GetLinksesByLinkIdAndLinkUrl(offset int, limit int, LinkId_ int64, LinkUrl_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_url = ?", LinkId_, LinkUrl_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkName Get Linkses via LinkIdAndLinkName
func GetLinksesByLinkIdAndLinkName(offset int, limit int, LinkId_ int64, LinkName_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_name = ?", LinkId_, LinkName_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkImage Get Linkses via LinkIdAndLinkImage
func GetLinksesByLinkIdAndLinkImage(offset int, limit int, LinkId_ int64, LinkImage_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_image = ?", LinkId_, LinkImage_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkTarget Get Linkses via LinkIdAndLinkTarget
func GetLinksesByLinkIdAndLinkTarget(offset int, limit int, LinkId_ int64, LinkTarget_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_target = ?", LinkId_, LinkTarget_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkDescription Get Linkses via LinkIdAndLinkDescription
func GetLinksesByLinkIdAndLinkDescription(offset int, limit int, LinkId_ int64, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_description = ?", LinkId_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkVisible Get Linkses via LinkIdAndLinkVisible
func GetLinksesByLinkIdAndLinkVisible(offset int, limit int, LinkId_ int64, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_visible = ?", LinkId_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkOwner Get Linkses via LinkIdAndLinkOwner
func GetLinksesByLinkIdAndLinkOwner(offset int, limit int, LinkId_ int64, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_owner = ?", LinkId_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkRating Get Linkses via LinkIdAndLinkRating
func GetLinksesByLinkIdAndLinkRating(offset int, limit int, LinkId_ int64, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_rating = ?", LinkId_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkUpdated Get Linkses via LinkIdAndLinkUpdated
func GetLinksesByLinkIdAndLinkUpdated(offset int, limit int, LinkId_ int64, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_updated = ?", LinkId_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkRel Get Linkses via LinkIdAndLinkRel
func GetLinksesByLinkIdAndLinkRel(offset int, limit int, LinkId_ int64, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_rel = ?", LinkId_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkNotes Get Linkses via LinkIdAndLinkNotes
func GetLinksesByLinkIdAndLinkNotes(offset int, limit int, LinkId_ int64, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_notes = ?", LinkId_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkIdAndLinkRss Get Linkses via LinkIdAndLinkRss
func GetLinksesByLinkIdAndLinkRss(offset int, limit int, LinkId_ int64, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ? and link_rss = ?", LinkId_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkName Get Linkses via LinkUrlAndLinkName
func GetLinksesByLinkUrlAndLinkName(offset int, limit int, LinkUrl_ string, LinkName_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_name = ?", LinkUrl_, LinkName_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkImage Get Linkses via LinkUrlAndLinkImage
func GetLinksesByLinkUrlAndLinkImage(offset int, limit int, LinkUrl_ string, LinkImage_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_image = ?", LinkUrl_, LinkImage_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkTarget Get Linkses via LinkUrlAndLinkTarget
func GetLinksesByLinkUrlAndLinkTarget(offset int, limit int, LinkUrl_ string, LinkTarget_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_target = ?", LinkUrl_, LinkTarget_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkDescription Get Linkses via LinkUrlAndLinkDescription
func GetLinksesByLinkUrlAndLinkDescription(offset int, limit int, LinkUrl_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_description = ?", LinkUrl_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkVisible Get Linkses via LinkUrlAndLinkVisible
func GetLinksesByLinkUrlAndLinkVisible(offset int, limit int, LinkUrl_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_visible = ?", LinkUrl_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkOwner Get Linkses via LinkUrlAndLinkOwner
func GetLinksesByLinkUrlAndLinkOwner(offset int, limit int, LinkUrl_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_owner = ?", LinkUrl_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkRating Get Linkses via LinkUrlAndLinkRating
func GetLinksesByLinkUrlAndLinkRating(offset int, limit int, LinkUrl_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_rating = ?", LinkUrl_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkUpdated Get Linkses via LinkUrlAndLinkUpdated
func GetLinksesByLinkUrlAndLinkUpdated(offset int, limit int, LinkUrl_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_updated = ?", LinkUrl_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkRel Get Linkses via LinkUrlAndLinkRel
func GetLinksesByLinkUrlAndLinkRel(offset int, limit int, LinkUrl_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_rel = ?", LinkUrl_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkNotes Get Linkses via LinkUrlAndLinkNotes
func GetLinksesByLinkUrlAndLinkNotes(offset int, limit int, LinkUrl_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_notes = ?", LinkUrl_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrlAndLinkRss Get Linkses via LinkUrlAndLinkRss
func GetLinksesByLinkUrlAndLinkRss(offset int, limit int, LinkUrl_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ? and link_rss = ?", LinkUrl_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkImage Get Linkses via LinkNameAndLinkImage
func GetLinksesByLinkNameAndLinkImage(offset int, limit int, LinkName_ string, LinkImage_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_image = ?", LinkName_, LinkImage_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkTarget Get Linkses via LinkNameAndLinkTarget
func GetLinksesByLinkNameAndLinkTarget(offset int, limit int, LinkName_ string, LinkTarget_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_target = ?", LinkName_, LinkTarget_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkDescription Get Linkses via LinkNameAndLinkDescription
func GetLinksesByLinkNameAndLinkDescription(offset int, limit int, LinkName_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_description = ?", LinkName_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkVisible Get Linkses via LinkNameAndLinkVisible
func GetLinksesByLinkNameAndLinkVisible(offset int, limit int, LinkName_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_visible = ?", LinkName_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkOwner Get Linkses via LinkNameAndLinkOwner
func GetLinksesByLinkNameAndLinkOwner(offset int, limit int, LinkName_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_owner = ?", LinkName_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkRating Get Linkses via LinkNameAndLinkRating
func GetLinksesByLinkNameAndLinkRating(offset int, limit int, LinkName_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_rating = ?", LinkName_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkUpdated Get Linkses via LinkNameAndLinkUpdated
func GetLinksesByLinkNameAndLinkUpdated(offset int, limit int, LinkName_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_updated = ?", LinkName_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkRel Get Linkses via LinkNameAndLinkRel
func GetLinksesByLinkNameAndLinkRel(offset int, limit int, LinkName_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_rel = ?", LinkName_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkNotes Get Linkses via LinkNameAndLinkNotes
func GetLinksesByLinkNameAndLinkNotes(offset int, limit int, LinkName_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_notes = ?", LinkName_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNameAndLinkRss Get Linkses via LinkNameAndLinkRss
func GetLinksesByLinkNameAndLinkRss(offset int, limit int, LinkName_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ? and link_rss = ?", LinkName_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkTarget Get Linkses via LinkImageAndLinkTarget
func GetLinksesByLinkImageAndLinkTarget(offset int, limit int, LinkImage_ string, LinkTarget_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_target = ?", LinkImage_, LinkTarget_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkDescription Get Linkses via LinkImageAndLinkDescription
func GetLinksesByLinkImageAndLinkDescription(offset int, limit int, LinkImage_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_description = ?", LinkImage_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkVisible Get Linkses via LinkImageAndLinkVisible
func GetLinksesByLinkImageAndLinkVisible(offset int, limit int, LinkImage_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_visible = ?", LinkImage_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkOwner Get Linkses via LinkImageAndLinkOwner
func GetLinksesByLinkImageAndLinkOwner(offset int, limit int, LinkImage_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_owner = ?", LinkImage_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkRating Get Linkses via LinkImageAndLinkRating
func GetLinksesByLinkImageAndLinkRating(offset int, limit int, LinkImage_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_rating = ?", LinkImage_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkUpdated Get Linkses via LinkImageAndLinkUpdated
func GetLinksesByLinkImageAndLinkUpdated(offset int, limit int, LinkImage_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_updated = ?", LinkImage_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkRel Get Linkses via LinkImageAndLinkRel
func GetLinksesByLinkImageAndLinkRel(offset int, limit int, LinkImage_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_rel = ?", LinkImage_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkNotes Get Linkses via LinkImageAndLinkNotes
func GetLinksesByLinkImageAndLinkNotes(offset int, limit int, LinkImage_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_notes = ?", LinkImage_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImageAndLinkRss Get Linkses via LinkImageAndLinkRss
func GetLinksesByLinkImageAndLinkRss(offset int, limit int, LinkImage_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ? and link_rss = ?", LinkImage_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkDescription Get Linkses via LinkTargetAndLinkDescription
func GetLinksesByLinkTargetAndLinkDescription(offset int, limit int, LinkTarget_ string, LinkDescription_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_description = ?", LinkTarget_, LinkDescription_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkVisible Get Linkses via LinkTargetAndLinkVisible
func GetLinksesByLinkTargetAndLinkVisible(offset int, limit int, LinkTarget_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_visible = ?", LinkTarget_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkOwner Get Linkses via LinkTargetAndLinkOwner
func GetLinksesByLinkTargetAndLinkOwner(offset int, limit int, LinkTarget_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_owner = ?", LinkTarget_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkRating Get Linkses via LinkTargetAndLinkRating
func GetLinksesByLinkTargetAndLinkRating(offset int, limit int, LinkTarget_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_rating = ?", LinkTarget_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkUpdated Get Linkses via LinkTargetAndLinkUpdated
func GetLinksesByLinkTargetAndLinkUpdated(offset int, limit int, LinkTarget_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_updated = ?", LinkTarget_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkRel Get Linkses via LinkTargetAndLinkRel
func GetLinksesByLinkTargetAndLinkRel(offset int, limit int, LinkTarget_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_rel = ?", LinkTarget_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkNotes Get Linkses via LinkTargetAndLinkNotes
func GetLinksesByLinkTargetAndLinkNotes(offset int, limit int, LinkTarget_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_notes = ?", LinkTarget_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTargetAndLinkRss Get Linkses via LinkTargetAndLinkRss
func GetLinksesByLinkTargetAndLinkRss(offset int, limit int, LinkTarget_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ? and link_rss = ?", LinkTarget_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkVisible Get Linkses via LinkDescriptionAndLinkVisible
func GetLinksesByLinkDescriptionAndLinkVisible(offset int, limit int, LinkDescription_ string, LinkVisible_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_visible = ?", LinkDescription_, LinkVisible_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkOwner Get Linkses via LinkDescriptionAndLinkOwner
func GetLinksesByLinkDescriptionAndLinkOwner(offset int, limit int, LinkDescription_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_owner = ?", LinkDescription_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkRating Get Linkses via LinkDescriptionAndLinkRating
func GetLinksesByLinkDescriptionAndLinkRating(offset int, limit int, LinkDescription_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_rating = ?", LinkDescription_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkUpdated Get Linkses via LinkDescriptionAndLinkUpdated
func GetLinksesByLinkDescriptionAndLinkUpdated(offset int, limit int, LinkDescription_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_updated = ?", LinkDescription_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkRel Get Linkses via LinkDescriptionAndLinkRel
func GetLinksesByLinkDescriptionAndLinkRel(offset int, limit int, LinkDescription_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_rel = ?", LinkDescription_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkNotes Get Linkses via LinkDescriptionAndLinkNotes
func GetLinksesByLinkDescriptionAndLinkNotes(offset int, limit int, LinkDescription_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_notes = ?", LinkDescription_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescriptionAndLinkRss Get Linkses via LinkDescriptionAndLinkRss
func GetLinksesByLinkDescriptionAndLinkRss(offset int, limit int, LinkDescription_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ? and link_rss = ?", LinkDescription_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkOwner Get Linkses via LinkVisibleAndLinkOwner
func GetLinksesByLinkVisibleAndLinkOwner(offset int, limit int, LinkVisible_ string, LinkOwner_ int64) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_owner = ?", LinkVisible_, LinkOwner_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkRating Get Linkses via LinkVisibleAndLinkRating
func GetLinksesByLinkVisibleAndLinkRating(offset int, limit int, LinkVisible_ string, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_rating = ?", LinkVisible_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkUpdated Get Linkses via LinkVisibleAndLinkUpdated
func GetLinksesByLinkVisibleAndLinkUpdated(offset int, limit int, LinkVisible_ string, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_updated = ?", LinkVisible_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkRel Get Linkses via LinkVisibleAndLinkRel
func GetLinksesByLinkVisibleAndLinkRel(offset int, limit int, LinkVisible_ string, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_rel = ?", LinkVisible_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkNotes Get Linkses via LinkVisibleAndLinkNotes
func GetLinksesByLinkVisibleAndLinkNotes(offset int, limit int, LinkVisible_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_notes = ?", LinkVisible_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisibleAndLinkRss Get Linkses via LinkVisibleAndLinkRss
func GetLinksesByLinkVisibleAndLinkRss(offset int, limit int, LinkVisible_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ? and link_rss = ?", LinkVisible_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkRating Get Linkses via LinkOwnerAndLinkRating
func GetLinksesByLinkOwnerAndLinkRating(offset int, limit int, LinkOwner_ int64, LinkRating_ int) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_rating = ?", LinkOwner_, LinkRating_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkUpdated Get Linkses via LinkOwnerAndLinkUpdated
func GetLinksesByLinkOwnerAndLinkUpdated(offset int, limit int, LinkOwner_ int64, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_updated = ?", LinkOwner_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkRel Get Linkses via LinkOwnerAndLinkRel
func GetLinksesByLinkOwnerAndLinkRel(offset int, limit int, LinkOwner_ int64, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_rel = ?", LinkOwner_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkNotes Get Linkses via LinkOwnerAndLinkNotes
func GetLinksesByLinkOwnerAndLinkNotes(offset int, limit int, LinkOwner_ int64, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_notes = ?", LinkOwner_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwnerAndLinkRss Get Linkses via LinkOwnerAndLinkRss
func GetLinksesByLinkOwnerAndLinkRss(offset int, limit int, LinkOwner_ int64, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ? and link_rss = ?", LinkOwner_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRatingAndLinkUpdated Get Linkses via LinkRatingAndLinkUpdated
func GetLinksesByLinkRatingAndLinkUpdated(offset int, limit int, LinkRating_ int, LinkUpdated_ time.Time) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rating = ? and link_updated = ?", LinkRating_, LinkUpdated_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRatingAndLinkRel Get Linkses via LinkRatingAndLinkRel
func GetLinksesByLinkRatingAndLinkRel(offset int, limit int, LinkRating_ int, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rating = ? and link_rel = ?", LinkRating_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRatingAndLinkNotes Get Linkses via LinkRatingAndLinkNotes
func GetLinksesByLinkRatingAndLinkNotes(offset int, limit int, LinkRating_ int, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rating = ? and link_notes = ?", LinkRating_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRatingAndLinkRss Get Linkses via LinkRatingAndLinkRss
func GetLinksesByLinkRatingAndLinkRss(offset int, limit int, LinkRating_ int, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rating = ? and link_rss = ?", LinkRating_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUpdatedAndLinkRel Get Linkses via LinkUpdatedAndLinkRel
func GetLinksesByLinkUpdatedAndLinkRel(offset int, limit int, LinkUpdated_ time.Time, LinkRel_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_updated = ? and link_rel = ?", LinkUpdated_, LinkRel_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUpdatedAndLinkNotes Get Linkses via LinkUpdatedAndLinkNotes
func GetLinksesByLinkUpdatedAndLinkNotes(offset int, limit int, LinkUpdated_ time.Time, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_updated = ? and link_notes = ?", LinkUpdated_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUpdatedAndLinkRss Get Linkses via LinkUpdatedAndLinkRss
func GetLinksesByLinkUpdatedAndLinkRss(offset int, limit int, LinkUpdated_ time.Time, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_updated = ? and link_rss = ?", LinkUpdated_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRelAndLinkNotes Get Linkses via LinkRelAndLinkNotes
func GetLinksesByLinkRelAndLinkNotes(offset int, limit int, LinkRel_ string, LinkNotes_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rel = ? and link_notes = ?", LinkRel_, LinkNotes_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRelAndLinkRss Get Linkses via LinkRelAndLinkRss
func GetLinksesByLinkRelAndLinkRss(offset int, limit int, LinkRel_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rel = ? and link_rss = ?", LinkRel_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNotesAndLinkRss Get Linkses via LinkNotesAndLinkRss
func GetLinksesByLinkNotesAndLinkRss(offset int, limit int, LinkNotes_ string, LinkRss_ string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_notes = ? and link_rss = ?", LinkNotes_, LinkRss_).Limit(limit, offset).Find(_Links)
	return _Links, err
}

// GetLinkses Get Linkses via field
func GetLinkses(offset int, limit int, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkId Get Linkss via LinkId
func GetLinksesByLinkId(offset int, limit int, LinkId_ int64, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_id = ?", LinkId_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUrl Get Linkss via LinkUrl
func GetLinksesByLinkUrl(offset int, limit int, LinkUrl_ string, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_url = ?", LinkUrl_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkName Get Linkss via LinkName
func GetLinksesByLinkName(offset int, limit int, LinkName_ string, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_name = ?", LinkName_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkImage Get Linkss via LinkImage
func GetLinksesByLinkImage(offset int, limit int, LinkImage_ string, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_image = ?", LinkImage_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkTarget Get Linkss via LinkTarget
func GetLinksesByLinkTarget(offset int, limit int, LinkTarget_ string, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_target = ?", LinkTarget_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkDescription Get Linkss via LinkDescription
func GetLinksesByLinkDescription(offset int, limit int, LinkDescription_ string, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_description = ?", LinkDescription_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkVisible Get Linkss via LinkVisible
func GetLinksesByLinkVisible(offset int, limit int, LinkVisible_ string, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_visible = ?", LinkVisible_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkOwner Get Linkss via LinkOwner
func GetLinksesByLinkOwner(offset int, limit int, LinkOwner_ int64, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_owner = ?", LinkOwner_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRating Get Linkss via LinkRating
func GetLinksesByLinkRating(offset int, limit int, LinkRating_ int, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rating = ?", LinkRating_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkUpdated Get Linkss via LinkUpdated
func GetLinksesByLinkUpdated(offset int, limit int, LinkUpdated_ time.Time, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_updated = ?", LinkUpdated_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRel Get Linkss via LinkRel
func GetLinksesByLinkRel(offset int, limit int, LinkRel_ string, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rel = ?", LinkRel_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkNotes Get Linkss via LinkNotes
func GetLinksesByLinkNotes(offset int, limit int, LinkNotes_ string, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_notes = ?", LinkNotes_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// GetLinksesByLinkRss Get Linkss via LinkRss
func GetLinksesByLinkRss(offset int, limit int, LinkRss_ string, field string) (*[]*Links, error) {
	var _Links = new([]*Links)
	err := Engine.Table("links").Where("link_rss = ?", LinkRss_).Limit(limit, offset).Desc(field).Find(_Links)
	return _Links, err
}

// HasLinksByLinkId Has Links via LinkId
func HasLinksByLinkId(iLinkId int64) bool {
	if has, err := Engine.Where("link_id = ?", iLinkId).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasLinksByLinkUrl Has Links via LinkUrl
func HasLinksByLinkUrl(iLinkUrl string) bool {
	if has, err := Engine.Where("link_url = ?", iLinkUrl).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasLinksByLinkName Has Links via LinkName
func HasLinksByLinkName(iLinkName string) bool {
	if has, err := Engine.Where("link_name = ?", iLinkName).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasLinksByLinkImage Has Links via LinkImage
func HasLinksByLinkImage(iLinkImage string) bool {
	if has, err := Engine.Where("link_image = ?", iLinkImage).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasLinksByLinkTarget Has Links via LinkTarget
func HasLinksByLinkTarget(iLinkTarget string) bool {
	if has, err := Engine.Where("link_target = ?", iLinkTarget).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasLinksByLinkDescription Has Links via LinkDescription
func HasLinksByLinkDescription(iLinkDescription string) bool {
	if has, err := Engine.Where("link_description = ?", iLinkDescription).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasLinksByLinkVisible Has Links via LinkVisible
func HasLinksByLinkVisible(iLinkVisible string) bool {
	if has, err := Engine.Where("link_visible = ?", iLinkVisible).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasLinksByLinkOwner Has Links via LinkOwner
func HasLinksByLinkOwner(iLinkOwner int64) bool {
	if has, err := Engine.Where("link_owner = ?", iLinkOwner).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasLinksByLinkRating Has Links via LinkRating
func HasLinksByLinkRating(iLinkRating int) bool {
	if has, err := Engine.Where("link_rating = ?", iLinkRating).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasLinksByLinkUpdated Has Links via LinkUpdated
func HasLinksByLinkUpdated(iLinkUpdated time.Time) bool {
	if has, err := Engine.Where("link_updated = ?", iLinkUpdated).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasLinksByLinkRel Has Links via LinkRel
func HasLinksByLinkRel(iLinkRel string) bool {
	if has, err := Engine.Where("link_rel = ?", iLinkRel).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasLinksByLinkNotes Has Links via LinkNotes
func HasLinksByLinkNotes(iLinkNotes string) bool {
	if has, err := Engine.Where("link_notes = ?", iLinkNotes).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasLinksByLinkRss Has Links via LinkRss
func HasLinksByLinkRss(iLinkRss string) bool {
	if has, err := Engine.Where("link_rss = ?", iLinkRss).Get(new(Links)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetLinksByLinkId Get Links via LinkId
func GetLinksByLinkId(iLinkId int64) (*Links, error) {
	var _Links = &Links{LinkId: iLinkId}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// GetLinksByLinkUrl Get Links via LinkUrl
func GetLinksByLinkUrl(iLinkUrl string) (*Links, error) {
	var _Links = &Links{LinkUrl: iLinkUrl}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// GetLinksByLinkName Get Links via LinkName
func GetLinksByLinkName(iLinkName string) (*Links, error) {
	var _Links = &Links{LinkName: iLinkName}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// GetLinksByLinkImage Get Links via LinkImage
func GetLinksByLinkImage(iLinkImage string) (*Links, error) {
	var _Links = &Links{LinkImage: iLinkImage}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// GetLinksByLinkTarget Get Links via LinkTarget
func GetLinksByLinkTarget(iLinkTarget string) (*Links, error) {
	var _Links = &Links{LinkTarget: iLinkTarget}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// GetLinksByLinkDescription Get Links via LinkDescription
func GetLinksByLinkDescription(iLinkDescription string) (*Links, error) {
	var _Links = &Links{LinkDescription: iLinkDescription}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// GetLinksByLinkVisible Get Links via LinkVisible
func GetLinksByLinkVisible(iLinkVisible string) (*Links, error) {
	var _Links = &Links{LinkVisible: iLinkVisible}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// GetLinksByLinkOwner Get Links via LinkOwner
func GetLinksByLinkOwner(iLinkOwner int64) (*Links, error) {
	var _Links = &Links{LinkOwner: iLinkOwner}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// GetLinksByLinkRating Get Links via LinkRating
func GetLinksByLinkRating(iLinkRating int) (*Links, error) {
	var _Links = &Links{LinkRating: iLinkRating}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// GetLinksByLinkUpdated Get Links via LinkUpdated
func GetLinksByLinkUpdated(iLinkUpdated time.Time) (*Links, error) {
	var _Links = &Links{LinkUpdated: iLinkUpdated}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// GetLinksByLinkRel Get Links via LinkRel
func GetLinksByLinkRel(iLinkRel string) (*Links, error) {
	var _Links = &Links{LinkRel: iLinkRel}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// GetLinksByLinkNotes Get Links via LinkNotes
func GetLinksByLinkNotes(iLinkNotes string) (*Links, error) {
	var _Links = &Links{LinkNotes: iLinkNotes}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// GetLinksByLinkRss Get Links via LinkRss
func GetLinksByLinkRss(iLinkRss string) (*Links, error) {
	var _Links = &Links{LinkRss: iLinkRss}
	has, err := Engine.Get(_Links)
	if has {
		return _Links, err
	} else {
		return nil, err
	}
}

// SetLinksByLinkId Set Links via LinkId
func SetLinksByLinkId(iLinkId int64, links *Links) (int64, error) {
	links.LinkId = iLinkId
	return Engine.Insert(links)
}

// SetLinksByLinkUrl Set Links via LinkUrl
func SetLinksByLinkUrl(iLinkUrl string, links *Links) (int64, error) {
	links.LinkUrl = iLinkUrl
	return Engine.Insert(links)
}

// SetLinksByLinkName Set Links via LinkName
func SetLinksByLinkName(iLinkName string, links *Links) (int64, error) {
	links.LinkName = iLinkName
	return Engine.Insert(links)
}

// SetLinksByLinkImage Set Links via LinkImage
func SetLinksByLinkImage(iLinkImage string, links *Links) (int64, error) {
	links.LinkImage = iLinkImage
	return Engine.Insert(links)
}

// SetLinksByLinkTarget Set Links via LinkTarget
func SetLinksByLinkTarget(iLinkTarget string, links *Links) (int64, error) {
	links.LinkTarget = iLinkTarget
	return Engine.Insert(links)
}

// SetLinksByLinkDescription Set Links via LinkDescription
func SetLinksByLinkDescription(iLinkDescription string, links *Links) (int64, error) {
	links.LinkDescription = iLinkDescription
	return Engine.Insert(links)
}

// SetLinksByLinkVisible Set Links via LinkVisible
func SetLinksByLinkVisible(iLinkVisible string, links *Links) (int64, error) {
	links.LinkVisible = iLinkVisible
	return Engine.Insert(links)
}

// SetLinksByLinkOwner Set Links via LinkOwner
func SetLinksByLinkOwner(iLinkOwner int64, links *Links) (int64, error) {
	links.LinkOwner = iLinkOwner
	return Engine.Insert(links)
}

// SetLinksByLinkRating Set Links via LinkRating
func SetLinksByLinkRating(iLinkRating int, links *Links) (int64, error) {
	links.LinkRating = iLinkRating
	return Engine.Insert(links)
}

// SetLinksByLinkUpdated Set Links via LinkUpdated
func SetLinksByLinkUpdated(iLinkUpdated time.Time, links *Links) (int64, error) {
	links.LinkUpdated = iLinkUpdated
	return Engine.Insert(links)
}

// SetLinksByLinkRel Set Links via LinkRel
func SetLinksByLinkRel(iLinkRel string, links *Links) (int64, error) {
	links.LinkRel = iLinkRel
	return Engine.Insert(links)
}

// SetLinksByLinkNotes Set Links via LinkNotes
func SetLinksByLinkNotes(iLinkNotes string, links *Links) (int64, error) {
	links.LinkNotes = iLinkNotes
	return Engine.Insert(links)
}

// SetLinksByLinkRss Set Links via LinkRss
func SetLinksByLinkRss(iLinkRss string, links *Links) (int64, error) {
	links.LinkRss = iLinkRss
	return Engine.Insert(links)
}

// AddLinks Add Links via all columns
func AddLinks(iLinkId int64, iLinkUrl string, iLinkName string, iLinkImage string, iLinkTarget string, iLinkDescription string, iLinkVisible string, iLinkOwner int64, iLinkRating int, iLinkUpdated time.Time, iLinkRel string, iLinkNotes string, iLinkRss string) error {
	_Links := &Links{LinkId: iLinkId, LinkUrl: iLinkUrl, LinkName: iLinkName, LinkImage: iLinkImage, LinkTarget: iLinkTarget, LinkDescription: iLinkDescription, LinkVisible: iLinkVisible, LinkOwner: iLinkOwner, LinkRating: iLinkRating, LinkUpdated: iLinkUpdated, LinkRel: iLinkRel, LinkNotes: iLinkNotes, LinkRss: iLinkRss}
	if _, err := Engine.Insert(_Links); err != nil {
		return err
	} else {
		return nil
	}
}

// PostLinks Post Links via iLinks
func PostLinks(iLinks *Links) (int64, error) {
	_, err := Engine.Insert(iLinks)
	return iLinks.LinkId, err
}

// PutLinks Put Links
func PutLinks(iLinks *Links) (int64, error) {
	_, err := Engine.Id(iLinks.LinkId).Update(iLinks)
	return iLinks.LinkId, err
}

// PutLinksByLinkId Put Links via LinkId
func PutLinksByLinkId(LinkId_ int64, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkId: LinkId_})
	return row, err
}

// PutLinksByLinkUrl Put Links via LinkUrl
func PutLinksByLinkUrl(LinkUrl_ string, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkUrl: LinkUrl_})
	return row, err
}

// PutLinksByLinkName Put Links via LinkName
func PutLinksByLinkName(LinkName_ string, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkName: LinkName_})
	return row, err
}

// PutLinksByLinkImage Put Links via LinkImage
func PutLinksByLinkImage(LinkImage_ string, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkImage: LinkImage_})
	return row, err
}

// PutLinksByLinkTarget Put Links via LinkTarget
func PutLinksByLinkTarget(LinkTarget_ string, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkTarget: LinkTarget_})
	return row, err
}

// PutLinksByLinkDescription Put Links via LinkDescription
func PutLinksByLinkDescription(LinkDescription_ string, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkDescription: LinkDescription_})
	return row, err
}

// PutLinksByLinkVisible Put Links via LinkVisible
func PutLinksByLinkVisible(LinkVisible_ string, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkVisible: LinkVisible_})
	return row, err
}

// PutLinksByLinkOwner Put Links via LinkOwner
func PutLinksByLinkOwner(LinkOwner_ int64, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkOwner: LinkOwner_})
	return row, err
}

// PutLinksByLinkRating Put Links via LinkRating
func PutLinksByLinkRating(LinkRating_ int, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkRating: LinkRating_})
	return row, err
}

// PutLinksByLinkUpdated Put Links via LinkUpdated
func PutLinksByLinkUpdated(LinkUpdated_ time.Time, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkUpdated: LinkUpdated_})
	return row, err
}

// PutLinksByLinkRel Put Links via LinkRel
func PutLinksByLinkRel(LinkRel_ string, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkRel: LinkRel_})
	return row, err
}

// PutLinksByLinkNotes Put Links via LinkNotes
func PutLinksByLinkNotes(LinkNotes_ string, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkNotes: LinkNotes_})
	return row, err
}

// PutLinksByLinkRss Put Links via LinkRss
func PutLinksByLinkRss(LinkRss_ string, iLinks *Links) (int64, error) {
	row, err := Engine.Update(iLinks, &Links{LinkRss: LinkRss_})
	return row, err
}

// UpdateLinksByLinkId via map[string]interface{}{}
func UpdateLinksByLinkId(iLinkId int64, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_id = ?", iLinkId).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateLinksByLinkUrl via map[string]interface{}{}
func UpdateLinksByLinkUrl(iLinkUrl string, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_url = ?", iLinkUrl).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateLinksByLinkName via map[string]interface{}{}
func UpdateLinksByLinkName(iLinkName string, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_name = ?", iLinkName).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateLinksByLinkImage via map[string]interface{}{}
func UpdateLinksByLinkImage(iLinkImage string, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_image = ?", iLinkImage).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateLinksByLinkTarget via map[string]interface{}{}
func UpdateLinksByLinkTarget(iLinkTarget string, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_target = ?", iLinkTarget).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateLinksByLinkDescription via map[string]interface{}{}
func UpdateLinksByLinkDescription(iLinkDescription string, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_description = ?", iLinkDescription).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateLinksByLinkVisible via map[string]interface{}{}
func UpdateLinksByLinkVisible(iLinkVisible string, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_visible = ?", iLinkVisible).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateLinksByLinkOwner via map[string]interface{}{}
func UpdateLinksByLinkOwner(iLinkOwner int64, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_owner = ?", iLinkOwner).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateLinksByLinkRating via map[string]interface{}{}
func UpdateLinksByLinkRating(iLinkRating int, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_rating = ?", iLinkRating).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateLinksByLinkUpdated via map[string]interface{}{}
func UpdateLinksByLinkUpdated(iLinkUpdated time.Time, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_updated = ?", iLinkUpdated).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateLinksByLinkRel via map[string]interface{}{}
func UpdateLinksByLinkRel(iLinkRel string, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_rel = ?", iLinkRel).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateLinksByLinkNotes via map[string]interface{}{}
func UpdateLinksByLinkNotes(iLinkNotes string, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_notes = ?", iLinkNotes).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateLinksByLinkRss via map[string]interface{}{}
func UpdateLinksByLinkRss(iLinkRss string, iLinksMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Links)).Where("link_rss = ?", iLinkRss).Update(iLinksMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteLinks Delete Links
func DeleteLinks(iLinkId int64) (int64, error) {
	row, err := Engine.Id(iLinkId).Delete(new(Links))
	return row, err
}

// DeleteLinksByLinkId Delete Links via LinkId
func DeleteLinksByLinkId(iLinkId int64) (err error) {
	var has bool
	var _Links = &Links{LinkId: iLinkId}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_id = ?", iLinkId).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteLinksByLinkUrl Delete Links via LinkUrl
func DeleteLinksByLinkUrl(iLinkUrl string) (err error) {
	var has bool
	var _Links = &Links{LinkUrl: iLinkUrl}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_url = ?", iLinkUrl).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteLinksByLinkName Delete Links via LinkName
func DeleteLinksByLinkName(iLinkName string) (err error) {
	var has bool
	var _Links = &Links{LinkName: iLinkName}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_name = ?", iLinkName).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteLinksByLinkImage Delete Links via LinkImage
func DeleteLinksByLinkImage(iLinkImage string) (err error) {
	var has bool
	var _Links = &Links{LinkImage: iLinkImage}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_image = ?", iLinkImage).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteLinksByLinkTarget Delete Links via LinkTarget
func DeleteLinksByLinkTarget(iLinkTarget string) (err error) {
	var has bool
	var _Links = &Links{LinkTarget: iLinkTarget}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_target = ?", iLinkTarget).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteLinksByLinkDescription Delete Links via LinkDescription
func DeleteLinksByLinkDescription(iLinkDescription string) (err error) {
	var has bool
	var _Links = &Links{LinkDescription: iLinkDescription}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_description = ?", iLinkDescription).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteLinksByLinkVisible Delete Links via LinkVisible
func DeleteLinksByLinkVisible(iLinkVisible string) (err error) {
	var has bool
	var _Links = &Links{LinkVisible: iLinkVisible}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_visible = ?", iLinkVisible).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteLinksByLinkOwner Delete Links via LinkOwner
func DeleteLinksByLinkOwner(iLinkOwner int64) (err error) {
	var has bool
	var _Links = &Links{LinkOwner: iLinkOwner}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_owner = ?", iLinkOwner).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteLinksByLinkRating Delete Links via LinkRating
func DeleteLinksByLinkRating(iLinkRating int) (err error) {
	var has bool
	var _Links = &Links{LinkRating: iLinkRating}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_rating = ?", iLinkRating).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteLinksByLinkUpdated Delete Links via LinkUpdated
func DeleteLinksByLinkUpdated(iLinkUpdated time.Time) (err error) {
	var has bool
	var _Links = &Links{LinkUpdated: iLinkUpdated}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_updated = ?", iLinkUpdated).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteLinksByLinkRel Delete Links via LinkRel
func DeleteLinksByLinkRel(iLinkRel string) (err error) {
	var has bool
	var _Links = &Links{LinkRel: iLinkRel}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_rel = ?", iLinkRel).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteLinksByLinkNotes Delete Links via LinkNotes
func DeleteLinksByLinkNotes(iLinkNotes string) (err error) {
	var has bool
	var _Links = &Links{LinkNotes: iLinkNotes}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_notes = ?", iLinkNotes).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteLinksByLinkRss Delete Links via LinkRss
func DeleteLinksByLinkRss(iLinkRss string) (err error) {
	var has bool
	var _Links = &Links{LinkRss: iLinkRss}
	if has, err = Engine.Get(_Links); (has == true) && (err == nil) {
		if row, err := Engine.Where("link_rss = ?", iLinkRss).Delete(new(Links)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
