package model

import (
	"time"
)

type Blogs struct {
	BlogId      int64     `xorm:"not null pk autoincr BIGINT(20)"`
	SiteId      int64     `xorm:"not null default 0 BIGINT(20)"`
	Domain      string    `xorm:"not null default '' index(domain) VARCHAR(200)"`
	Path        string    `xorm:"not null default '' index(domain) VARCHAR(100)"`
	Registered  time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	LastUpdated time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
	Public      int       `xorm:"not null default 1 TINYINT(2)"`
	Archived    int       `xorm:"not null default 0 TINYINT(2)"`
	Mature      int       `xorm:"not null default 0 TINYINT(2)"`
	Spam        int       `xorm:"not null default 0 TINYINT(2)"`
	Deleted     int       `xorm:"not null default 0 TINYINT(2)"`
	LangId      int       `xorm:"not null default 0 index INT(11)"`
}

// GetBlogsesCount Blogss' Count
func GetBlogsesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Blogs{})
	return total, err
}

// GetBlogsCountViaBlogId Get Blogs via BlogId
func GetBlogsCountViaBlogId(iBlogId int64) int64 {
	n, _ := Engine.Where("blog_id = ?", iBlogId).Count(&Blogs{BlogId: iBlogId})
	return n
}

// GetBlogsCountViaSiteId Get Blogs via SiteId
func GetBlogsCountViaSiteId(iSiteId int64) int64 {
	n, _ := Engine.Where("site_id = ?", iSiteId).Count(&Blogs{SiteId: iSiteId})
	return n
}

// GetBlogsCountViaDomain Get Blogs via Domain
func GetBlogsCountViaDomain(iDomain string) int64 {
	n, _ := Engine.Where("domain = ?", iDomain).Count(&Blogs{Domain: iDomain})
	return n
}

// GetBlogsCountViaPath Get Blogs via Path
func GetBlogsCountViaPath(iPath string) int64 {
	n, _ := Engine.Where("path = ?", iPath).Count(&Blogs{Path: iPath})
	return n
}

// GetBlogsCountViaRegistered Get Blogs via Registered
func GetBlogsCountViaRegistered(iRegistered time.Time) int64 {
	n, _ := Engine.Where("registered = ?", iRegistered).Count(&Blogs{Registered: iRegistered})
	return n
}

// GetBlogsCountViaLastUpdated Get Blogs via LastUpdated
func GetBlogsCountViaLastUpdated(iLastUpdated time.Time) int64 {
	n, _ := Engine.Where("last_updated = ?", iLastUpdated).Count(&Blogs{LastUpdated: iLastUpdated})
	return n
}

// GetBlogsCountViaPublic Get Blogs via Public
func GetBlogsCountViaPublic(iPublic int) int64 {
	n, _ := Engine.Where("public = ?", iPublic).Count(&Blogs{Public: iPublic})
	return n
}

// GetBlogsCountViaArchived Get Blogs via Archived
func GetBlogsCountViaArchived(iArchived int) int64 {
	n, _ := Engine.Where("archived = ?", iArchived).Count(&Blogs{Archived: iArchived})
	return n
}

// GetBlogsCountViaMature Get Blogs via Mature
func GetBlogsCountViaMature(iMature int) int64 {
	n, _ := Engine.Where("mature = ?", iMature).Count(&Blogs{Mature: iMature})
	return n
}

// GetBlogsCountViaSpam Get Blogs via Spam
func GetBlogsCountViaSpam(iSpam int) int64 {
	n, _ := Engine.Where("spam = ?", iSpam).Count(&Blogs{Spam: iSpam})
	return n
}

// GetBlogsCountViaDeleted Get Blogs via Deleted
func GetBlogsCountViaDeleted(iDeleted int) int64 {
	n, _ := Engine.Where("deleted = ?", iDeleted).Count(&Blogs{Deleted: iDeleted})
	return n
}

// GetBlogsCountViaLangId Get Blogs via LangId
func GetBlogsCountViaLangId(iLangId int) int64 {
	n, _ := Engine.Where("lang_id = ?", iLangId).Count(&Blogs{LangId: iLangId})
	return n
}

// GetBlogsesByBlogIdAndSiteIdAndDomain Get Blogses via BlogIdAndSiteIdAndDomain
func GetBlogsesByBlogIdAndSiteIdAndDomain(offset int, limit int, BlogId_ int64, SiteId_ int64, Domain_ string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and site_id = ? and domain = ?", BlogId_, SiteId_, Domain_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSiteIdAndPath Get Blogses via BlogIdAndSiteIdAndPath
func GetBlogsesByBlogIdAndSiteIdAndPath(offset int, limit int, BlogId_ int64, SiteId_ int64, Path_ string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and site_id = ? and path = ?", BlogId_, SiteId_, Path_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSiteIdAndRegistered Get Blogses via BlogIdAndSiteIdAndRegistered
func GetBlogsesByBlogIdAndSiteIdAndRegistered(offset int, limit int, BlogId_ int64, SiteId_ int64, Registered_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and site_id = ? and registered = ?", BlogId_, SiteId_, Registered_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSiteIdAndLastUpdated Get Blogses via BlogIdAndSiteIdAndLastUpdated
func GetBlogsesByBlogIdAndSiteIdAndLastUpdated(offset int, limit int, BlogId_ int64, SiteId_ int64, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and site_id = ? and last_updated = ?", BlogId_, SiteId_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSiteIdAndPublic Get Blogses via BlogIdAndSiteIdAndPublic
func GetBlogsesByBlogIdAndSiteIdAndPublic(offset int, limit int, BlogId_ int64, SiteId_ int64, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and site_id = ? and public = ?", BlogId_, SiteId_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSiteIdAndArchived Get Blogses via BlogIdAndSiteIdAndArchived
func GetBlogsesByBlogIdAndSiteIdAndArchived(offset int, limit int, BlogId_ int64, SiteId_ int64, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and site_id = ? and archived = ?", BlogId_, SiteId_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSiteIdAndMature Get Blogses via BlogIdAndSiteIdAndMature
func GetBlogsesByBlogIdAndSiteIdAndMature(offset int, limit int, BlogId_ int64, SiteId_ int64, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and site_id = ? and mature = ?", BlogId_, SiteId_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSiteIdAndSpam Get Blogses via BlogIdAndSiteIdAndSpam
func GetBlogsesByBlogIdAndSiteIdAndSpam(offset int, limit int, BlogId_ int64, SiteId_ int64, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and site_id = ? and spam = ?", BlogId_, SiteId_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSiteIdAndDeleted Get Blogses via BlogIdAndSiteIdAndDeleted
func GetBlogsesByBlogIdAndSiteIdAndDeleted(offset int, limit int, BlogId_ int64, SiteId_ int64, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and site_id = ? and deleted = ?", BlogId_, SiteId_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSiteIdAndLangId Get Blogses via BlogIdAndSiteIdAndLangId
func GetBlogsesByBlogIdAndSiteIdAndLangId(offset int, limit int, BlogId_ int64, SiteId_ int64, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and site_id = ? and lang_id = ?", BlogId_, SiteId_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndDomainAndPath Get Blogses via BlogIdAndDomainAndPath
func GetBlogsesByBlogIdAndDomainAndPath(offset int, limit int, BlogId_ int64, Domain_ string, Path_ string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and domain = ? and path = ?", BlogId_, Domain_, Path_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndDomainAndRegistered Get Blogses via BlogIdAndDomainAndRegistered
func GetBlogsesByBlogIdAndDomainAndRegistered(offset int, limit int, BlogId_ int64, Domain_ string, Registered_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and domain = ? and registered = ?", BlogId_, Domain_, Registered_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndDomainAndLastUpdated Get Blogses via BlogIdAndDomainAndLastUpdated
func GetBlogsesByBlogIdAndDomainAndLastUpdated(offset int, limit int, BlogId_ int64, Domain_ string, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and domain = ? and last_updated = ?", BlogId_, Domain_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndDomainAndPublic Get Blogses via BlogIdAndDomainAndPublic
func GetBlogsesByBlogIdAndDomainAndPublic(offset int, limit int, BlogId_ int64, Domain_ string, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and domain = ? and public = ?", BlogId_, Domain_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndDomainAndArchived Get Blogses via BlogIdAndDomainAndArchived
func GetBlogsesByBlogIdAndDomainAndArchived(offset int, limit int, BlogId_ int64, Domain_ string, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and domain = ? and archived = ?", BlogId_, Domain_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndDomainAndMature Get Blogses via BlogIdAndDomainAndMature
func GetBlogsesByBlogIdAndDomainAndMature(offset int, limit int, BlogId_ int64, Domain_ string, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and domain = ? and mature = ?", BlogId_, Domain_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndDomainAndSpam Get Blogses via BlogIdAndDomainAndSpam
func GetBlogsesByBlogIdAndDomainAndSpam(offset int, limit int, BlogId_ int64, Domain_ string, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and domain = ? and spam = ?", BlogId_, Domain_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndDomainAndDeleted Get Blogses via BlogIdAndDomainAndDeleted
func GetBlogsesByBlogIdAndDomainAndDeleted(offset int, limit int, BlogId_ int64, Domain_ string, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and domain = ? and deleted = ?", BlogId_, Domain_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndDomainAndLangId Get Blogses via BlogIdAndDomainAndLangId
func GetBlogsesByBlogIdAndDomainAndLangId(offset int, limit int, BlogId_ int64, Domain_ string, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and domain = ? and lang_id = ?", BlogId_, Domain_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPathAndRegistered Get Blogses via BlogIdAndPathAndRegistered
func GetBlogsesByBlogIdAndPathAndRegistered(offset int, limit int, BlogId_ int64, Path_ string, Registered_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and path = ? and registered = ?", BlogId_, Path_, Registered_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPathAndLastUpdated Get Blogses via BlogIdAndPathAndLastUpdated
func GetBlogsesByBlogIdAndPathAndLastUpdated(offset int, limit int, BlogId_ int64, Path_ string, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and path = ? and last_updated = ?", BlogId_, Path_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPathAndPublic Get Blogses via BlogIdAndPathAndPublic
func GetBlogsesByBlogIdAndPathAndPublic(offset int, limit int, BlogId_ int64, Path_ string, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and path = ? and public = ?", BlogId_, Path_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPathAndArchived Get Blogses via BlogIdAndPathAndArchived
func GetBlogsesByBlogIdAndPathAndArchived(offset int, limit int, BlogId_ int64, Path_ string, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and path = ? and archived = ?", BlogId_, Path_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPathAndMature Get Blogses via BlogIdAndPathAndMature
func GetBlogsesByBlogIdAndPathAndMature(offset int, limit int, BlogId_ int64, Path_ string, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and path = ? and mature = ?", BlogId_, Path_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPathAndSpam Get Blogses via BlogIdAndPathAndSpam
func GetBlogsesByBlogIdAndPathAndSpam(offset int, limit int, BlogId_ int64, Path_ string, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and path = ? and spam = ?", BlogId_, Path_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPathAndDeleted Get Blogses via BlogIdAndPathAndDeleted
func GetBlogsesByBlogIdAndPathAndDeleted(offset int, limit int, BlogId_ int64, Path_ string, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and path = ? and deleted = ?", BlogId_, Path_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPathAndLangId Get Blogses via BlogIdAndPathAndLangId
func GetBlogsesByBlogIdAndPathAndLangId(offset int, limit int, BlogId_ int64, Path_ string, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and path = ? and lang_id = ?", BlogId_, Path_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndRegisteredAndLastUpdated Get Blogses via BlogIdAndRegisteredAndLastUpdated
func GetBlogsesByBlogIdAndRegisteredAndLastUpdated(offset int, limit int, BlogId_ int64, Registered_ time.Time, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and registered = ? and last_updated = ?", BlogId_, Registered_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndRegisteredAndPublic Get Blogses via BlogIdAndRegisteredAndPublic
func GetBlogsesByBlogIdAndRegisteredAndPublic(offset int, limit int, BlogId_ int64, Registered_ time.Time, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and registered = ? and public = ?", BlogId_, Registered_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndRegisteredAndArchived Get Blogses via BlogIdAndRegisteredAndArchived
func GetBlogsesByBlogIdAndRegisteredAndArchived(offset int, limit int, BlogId_ int64, Registered_ time.Time, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and registered = ? and archived = ?", BlogId_, Registered_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndRegisteredAndMature Get Blogses via BlogIdAndRegisteredAndMature
func GetBlogsesByBlogIdAndRegisteredAndMature(offset int, limit int, BlogId_ int64, Registered_ time.Time, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and registered = ? and mature = ?", BlogId_, Registered_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndRegisteredAndSpam Get Blogses via BlogIdAndRegisteredAndSpam
func GetBlogsesByBlogIdAndRegisteredAndSpam(offset int, limit int, BlogId_ int64, Registered_ time.Time, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and registered = ? and spam = ?", BlogId_, Registered_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndRegisteredAndDeleted Get Blogses via BlogIdAndRegisteredAndDeleted
func GetBlogsesByBlogIdAndRegisteredAndDeleted(offset int, limit int, BlogId_ int64, Registered_ time.Time, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and registered = ? and deleted = ?", BlogId_, Registered_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndRegisteredAndLangId Get Blogses via BlogIdAndRegisteredAndLangId
func GetBlogsesByBlogIdAndRegisteredAndLangId(offset int, limit int, BlogId_ int64, Registered_ time.Time, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and registered = ? and lang_id = ?", BlogId_, Registered_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndLastUpdatedAndPublic Get Blogses via BlogIdAndLastUpdatedAndPublic
func GetBlogsesByBlogIdAndLastUpdatedAndPublic(offset int, limit int, BlogId_ int64, LastUpdated_ time.Time, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and last_updated = ? and public = ?", BlogId_, LastUpdated_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndLastUpdatedAndArchived Get Blogses via BlogIdAndLastUpdatedAndArchived
func GetBlogsesByBlogIdAndLastUpdatedAndArchived(offset int, limit int, BlogId_ int64, LastUpdated_ time.Time, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and last_updated = ? and archived = ?", BlogId_, LastUpdated_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndLastUpdatedAndMature Get Blogses via BlogIdAndLastUpdatedAndMature
func GetBlogsesByBlogIdAndLastUpdatedAndMature(offset int, limit int, BlogId_ int64, LastUpdated_ time.Time, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and last_updated = ? and mature = ?", BlogId_, LastUpdated_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndLastUpdatedAndSpam Get Blogses via BlogIdAndLastUpdatedAndSpam
func GetBlogsesByBlogIdAndLastUpdatedAndSpam(offset int, limit int, BlogId_ int64, LastUpdated_ time.Time, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and last_updated = ? and spam = ?", BlogId_, LastUpdated_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndLastUpdatedAndDeleted Get Blogses via BlogIdAndLastUpdatedAndDeleted
func GetBlogsesByBlogIdAndLastUpdatedAndDeleted(offset int, limit int, BlogId_ int64, LastUpdated_ time.Time, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and last_updated = ? and deleted = ?", BlogId_, LastUpdated_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndLastUpdatedAndLangId Get Blogses via BlogIdAndLastUpdatedAndLangId
func GetBlogsesByBlogIdAndLastUpdatedAndLangId(offset int, limit int, BlogId_ int64, LastUpdated_ time.Time, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and last_updated = ? and lang_id = ?", BlogId_, LastUpdated_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPublicAndArchived Get Blogses via BlogIdAndPublicAndArchived
func GetBlogsesByBlogIdAndPublicAndArchived(offset int, limit int, BlogId_ int64, Public_ int, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and public = ? and archived = ?", BlogId_, Public_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPublicAndMature Get Blogses via BlogIdAndPublicAndMature
func GetBlogsesByBlogIdAndPublicAndMature(offset int, limit int, BlogId_ int64, Public_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and public = ? and mature = ?", BlogId_, Public_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPublicAndSpam Get Blogses via BlogIdAndPublicAndSpam
func GetBlogsesByBlogIdAndPublicAndSpam(offset int, limit int, BlogId_ int64, Public_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and public = ? and spam = ?", BlogId_, Public_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPublicAndDeleted Get Blogses via BlogIdAndPublicAndDeleted
func GetBlogsesByBlogIdAndPublicAndDeleted(offset int, limit int, BlogId_ int64, Public_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and public = ? and deleted = ?", BlogId_, Public_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPublicAndLangId Get Blogses via BlogIdAndPublicAndLangId
func GetBlogsesByBlogIdAndPublicAndLangId(offset int, limit int, BlogId_ int64, Public_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and public = ? and lang_id = ?", BlogId_, Public_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndArchivedAndMature Get Blogses via BlogIdAndArchivedAndMature
func GetBlogsesByBlogIdAndArchivedAndMature(offset int, limit int, BlogId_ int64, Archived_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and archived = ? and mature = ?", BlogId_, Archived_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndArchivedAndSpam Get Blogses via BlogIdAndArchivedAndSpam
func GetBlogsesByBlogIdAndArchivedAndSpam(offset int, limit int, BlogId_ int64, Archived_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and archived = ? and spam = ?", BlogId_, Archived_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndArchivedAndDeleted Get Blogses via BlogIdAndArchivedAndDeleted
func GetBlogsesByBlogIdAndArchivedAndDeleted(offset int, limit int, BlogId_ int64, Archived_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and archived = ? and deleted = ?", BlogId_, Archived_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndArchivedAndLangId Get Blogses via BlogIdAndArchivedAndLangId
func GetBlogsesByBlogIdAndArchivedAndLangId(offset int, limit int, BlogId_ int64, Archived_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and archived = ? and lang_id = ?", BlogId_, Archived_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndMatureAndSpam Get Blogses via BlogIdAndMatureAndSpam
func GetBlogsesByBlogIdAndMatureAndSpam(offset int, limit int, BlogId_ int64, Mature_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and mature = ? and spam = ?", BlogId_, Mature_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndMatureAndDeleted Get Blogses via BlogIdAndMatureAndDeleted
func GetBlogsesByBlogIdAndMatureAndDeleted(offset int, limit int, BlogId_ int64, Mature_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and mature = ? and deleted = ?", BlogId_, Mature_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndMatureAndLangId Get Blogses via BlogIdAndMatureAndLangId
func GetBlogsesByBlogIdAndMatureAndLangId(offset int, limit int, BlogId_ int64, Mature_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and mature = ? and lang_id = ?", BlogId_, Mature_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSpamAndDeleted Get Blogses via BlogIdAndSpamAndDeleted
func GetBlogsesByBlogIdAndSpamAndDeleted(offset int, limit int, BlogId_ int64, Spam_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and spam = ? and deleted = ?", BlogId_, Spam_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSpamAndLangId Get Blogses via BlogIdAndSpamAndLangId
func GetBlogsesByBlogIdAndSpamAndLangId(offset int, limit int, BlogId_ int64, Spam_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and spam = ? and lang_id = ?", BlogId_, Spam_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndDeletedAndLangId Get Blogses via BlogIdAndDeletedAndLangId
func GetBlogsesByBlogIdAndDeletedAndLangId(offset int, limit int, BlogId_ int64, Deleted_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and deleted = ? and lang_id = ?", BlogId_, Deleted_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndDomainAndPath Get Blogses via SiteIdAndDomainAndPath
func GetBlogsesBySiteIdAndDomainAndPath(offset int, limit int, SiteId_ int64, Domain_ string, Path_ string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and domain = ? and path = ?", SiteId_, Domain_, Path_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndDomainAndRegistered Get Blogses via SiteIdAndDomainAndRegistered
func GetBlogsesBySiteIdAndDomainAndRegistered(offset int, limit int, SiteId_ int64, Domain_ string, Registered_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and domain = ? and registered = ?", SiteId_, Domain_, Registered_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndDomainAndLastUpdated Get Blogses via SiteIdAndDomainAndLastUpdated
func GetBlogsesBySiteIdAndDomainAndLastUpdated(offset int, limit int, SiteId_ int64, Domain_ string, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and domain = ? and last_updated = ?", SiteId_, Domain_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndDomainAndPublic Get Blogses via SiteIdAndDomainAndPublic
func GetBlogsesBySiteIdAndDomainAndPublic(offset int, limit int, SiteId_ int64, Domain_ string, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and domain = ? and public = ?", SiteId_, Domain_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndDomainAndArchived Get Blogses via SiteIdAndDomainAndArchived
func GetBlogsesBySiteIdAndDomainAndArchived(offset int, limit int, SiteId_ int64, Domain_ string, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and domain = ? and archived = ?", SiteId_, Domain_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndDomainAndMature Get Blogses via SiteIdAndDomainAndMature
func GetBlogsesBySiteIdAndDomainAndMature(offset int, limit int, SiteId_ int64, Domain_ string, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and domain = ? and mature = ?", SiteId_, Domain_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndDomainAndSpam Get Blogses via SiteIdAndDomainAndSpam
func GetBlogsesBySiteIdAndDomainAndSpam(offset int, limit int, SiteId_ int64, Domain_ string, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and domain = ? and spam = ?", SiteId_, Domain_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndDomainAndDeleted Get Blogses via SiteIdAndDomainAndDeleted
func GetBlogsesBySiteIdAndDomainAndDeleted(offset int, limit int, SiteId_ int64, Domain_ string, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and domain = ? and deleted = ?", SiteId_, Domain_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndDomainAndLangId Get Blogses via SiteIdAndDomainAndLangId
func GetBlogsesBySiteIdAndDomainAndLangId(offset int, limit int, SiteId_ int64, Domain_ string, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and domain = ? and lang_id = ?", SiteId_, Domain_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPathAndRegistered Get Blogses via SiteIdAndPathAndRegistered
func GetBlogsesBySiteIdAndPathAndRegistered(offset int, limit int, SiteId_ int64, Path_ string, Registered_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and path = ? and registered = ?", SiteId_, Path_, Registered_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPathAndLastUpdated Get Blogses via SiteIdAndPathAndLastUpdated
func GetBlogsesBySiteIdAndPathAndLastUpdated(offset int, limit int, SiteId_ int64, Path_ string, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and path = ? and last_updated = ?", SiteId_, Path_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPathAndPublic Get Blogses via SiteIdAndPathAndPublic
func GetBlogsesBySiteIdAndPathAndPublic(offset int, limit int, SiteId_ int64, Path_ string, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and path = ? and public = ?", SiteId_, Path_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPathAndArchived Get Blogses via SiteIdAndPathAndArchived
func GetBlogsesBySiteIdAndPathAndArchived(offset int, limit int, SiteId_ int64, Path_ string, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and path = ? and archived = ?", SiteId_, Path_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPathAndMature Get Blogses via SiteIdAndPathAndMature
func GetBlogsesBySiteIdAndPathAndMature(offset int, limit int, SiteId_ int64, Path_ string, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and path = ? and mature = ?", SiteId_, Path_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPathAndSpam Get Blogses via SiteIdAndPathAndSpam
func GetBlogsesBySiteIdAndPathAndSpam(offset int, limit int, SiteId_ int64, Path_ string, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and path = ? and spam = ?", SiteId_, Path_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPathAndDeleted Get Blogses via SiteIdAndPathAndDeleted
func GetBlogsesBySiteIdAndPathAndDeleted(offset int, limit int, SiteId_ int64, Path_ string, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and path = ? and deleted = ?", SiteId_, Path_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPathAndLangId Get Blogses via SiteIdAndPathAndLangId
func GetBlogsesBySiteIdAndPathAndLangId(offset int, limit int, SiteId_ int64, Path_ string, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and path = ? and lang_id = ?", SiteId_, Path_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndRegisteredAndLastUpdated Get Blogses via SiteIdAndRegisteredAndLastUpdated
func GetBlogsesBySiteIdAndRegisteredAndLastUpdated(offset int, limit int, SiteId_ int64, Registered_ time.Time, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and registered = ? and last_updated = ?", SiteId_, Registered_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndRegisteredAndPublic Get Blogses via SiteIdAndRegisteredAndPublic
func GetBlogsesBySiteIdAndRegisteredAndPublic(offset int, limit int, SiteId_ int64, Registered_ time.Time, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and registered = ? and public = ?", SiteId_, Registered_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndRegisteredAndArchived Get Blogses via SiteIdAndRegisteredAndArchived
func GetBlogsesBySiteIdAndRegisteredAndArchived(offset int, limit int, SiteId_ int64, Registered_ time.Time, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and registered = ? and archived = ?", SiteId_, Registered_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndRegisteredAndMature Get Blogses via SiteIdAndRegisteredAndMature
func GetBlogsesBySiteIdAndRegisteredAndMature(offset int, limit int, SiteId_ int64, Registered_ time.Time, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and registered = ? and mature = ?", SiteId_, Registered_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndRegisteredAndSpam Get Blogses via SiteIdAndRegisteredAndSpam
func GetBlogsesBySiteIdAndRegisteredAndSpam(offset int, limit int, SiteId_ int64, Registered_ time.Time, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and registered = ? and spam = ?", SiteId_, Registered_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndRegisteredAndDeleted Get Blogses via SiteIdAndRegisteredAndDeleted
func GetBlogsesBySiteIdAndRegisteredAndDeleted(offset int, limit int, SiteId_ int64, Registered_ time.Time, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and registered = ? and deleted = ?", SiteId_, Registered_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndRegisteredAndLangId Get Blogses via SiteIdAndRegisteredAndLangId
func GetBlogsesBySiteIdAndRegisteredAndLangId(offset int, limit int, SiteId_ int64, Registered_ time.Time, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and registered = ? and lang_id = ?", SiteId_, Registered_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndLastUpdatedAndPublic Get Blogses via SiteIdAndLastUpdatedAndPublic
func GetBlogsesBySiteIdAndLastUpdatedAndPublic(offset int, limit int, SiteId_ int64, LastUpdated_ time.Time, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and last_updated = ? and public = ?", SiteId_, LastUpdated_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndLastUpdatedAndArchived Get Blogses via SiteIdAndLastUpdatedAndArchived
func GetBlogsesBySiteIdAndLastUpdatedAndArchived(offset int, limit int, SiteId_ int64, LastUpdated_ time.Time, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and last_updated = ? and archived = ?", SiteId_, LastUpdated_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndLastUpdatedAndMature Get Blogses via SiteIdAndLastUpdatedAndMature
func GetBlogsesBySiteIdAndLastUpdatedAndMature(offset int, limit int, SiteId_ int64, LastUpdated_ time.Time, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and last_updated = ? and mature = ?", SiteId_, LastUpdated_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndLastUpdatedAndSpam Get Blogses via SiteIdAndLastUpdatedAndSpam
func GetBlogsesBySiteIdAndLastUpdatedAndSpam(offset int, limit int, SiteId_ int64, LastUpdated_ time.Time, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and last_updated = ? and spam = ?", SiteId_, LastUpdated_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndLastUpdatedAndDeleted Get Blogses via SiteIdAndLastUpdatedAndDeleted
func GetBlogsesBySiteIdAndLastUpdatedAndDeleted(offset int, limit int, SiteId_ int64, LastUpdated_ time.Time, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and last_updated = ? and deleted = ?", SiteId_, LastUpdated_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndLastUpdatedAndLangId Get Blogses via SiteIdAndLastUpdatedAndLangId
func GetBlogsesBySiteIdAndLastUpdatedAndLangId(offset int, limit int, SiteId_ int64, LastUpdated_ time.Time, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and last_updated = ? and lang_id = ?", SiteId_, LastUpdated_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPublicAndArchived Get Blogses via SiteIdAndPublicAndArchived
func GetBlogsesBySiteIdAndPublicAndArchived(offset int, limit int, SiteId_ int64, Public_ int, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and public = ? and archived = ?", SiteId_, Public_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPublicAndMature Get Blogses via SiteIdAndPublicAndMature
func GetBlogsesBySiteIdAndPublicAndMature(offset int, limit int, SiteId_ int64, Public_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and public = ? and mature = ?", SiteId_, Public_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPublicAndSpam Get Blogses via SiteIdAndPublicAndSpam
func GetBlogsesBySiteIdAndPublicAndSpam(offset int, limit int, SiteId_ int64, Public_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and public = ? and spam = ?", SiteId_, Public_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPublicAndDeleted Get Blogses via SiteIdAndPublicAndDeleted
func GetBlogsesBySiteIdAndPublicAndDeleted(offset int, limit int, SiteId_ int64, Public_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and public = ? and deleted = ?", SiteId_, Public_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPublicAndLangId Get Blogses via SiteIdAndPublicAndLangId
func GetBlogsesBySiteIdAndPublicAndLangId(offset int, limit int, SiteId_ int64, Public_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and public = ? and lang_id = ?", SiteId_, Public_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndArchivedAndMature Get Blogses via SiteIdAndArchivedAndMature
func GetBlogsesBySiteIdAndArchivedAndMature(offset int, limit int, SiteId_ int64, Archived_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and archived = ? and mature = ?", SiteId_, Archived_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndArchivedAndSpam Get Blogses via SiteIdAndArchivedAndSpam
func GetBlogsesBySiteIdAndArchivedAndSpam(offset int, limit int, SiteId_ int64, Archived_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and archived = ? and spam = ?", SiteId_, Archived_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndArchivedAndDeleted Get Blogses via SiteIdAndArchivedAndDeleted
func GetBlogsesBySiteIdAndArchivedAndDeleted(offset int, limit int, SiteId_ int64, Archived_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and archived = ? and deleted = ?", SiteId_, Archived_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndArchivedAndLangId Get Blogses via SiteIdAndArchivedAndLangId
func GetBlogsesBySiteIdAndArchivedAndLangId(offset int, limit int, SiteId_ int64, Archived_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and archived = ? and lang_id = ?", SiteId_, Archived_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndMatureAndSpam Get Blogses via SiteIdAndMatureAndSpam
func GetBlogsesBySiteIdAndMatureAndSpam(offset int, limit int, SiteId_ int64, Mature_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and mature = ? and spam = ?", SiteId_, Mature_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndMatureAndDeleted Get Blogses via SiteIdAndMatureAndDeleted
func GetBlogsesBySiteIdAndMatureAndDeleted(offset int, limit int, SiteId_ int64, Mature_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and mature = ? and deleted = ?", SiteId_, Mature_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndMatureAndLangId Get Blogses via SiteIdAndMatureAndLangId
func GetBlogsesBySiteIdAndMatureAndLangId(offset int, limit int, SiteId_ int64, Mature_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and mature = ? and lang_id = ?", SiteId_, Mature_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndSpamAndDeleted Get Blogses via SiteIdAndSpamAndDeleted
func GetBlogsesBySiteIdAndSpamAndDeleted(offset int, limit int, SiteId_ int64, Spam_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and spam = ? and deleted = ?", SiteId_, Spam_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndSpamAndLangId Get Blogses via SiteIdAndSpamAndLangId
func GetBlogsesBySiteIdAndSpamAndLangId(offset int, limit int, SiteId_ int64, Spam_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and spam = ? and lang_id = ?", SiteId_, Spam_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndDeletedAndLangId Get Blogses via SiteIdAndDeletedAndLangId
func GetBlogsesBySiteIdAndDeletedAndLangId(offset int, limit int, SiteId_ int64, Deleted_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and deleted = ? and lang_id = ?", SiteId_, Deleted_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPathAndRegistered Get Blogses via DomainAndPathAndRegistered
func GetBlogsesByDomainAndPathAndRegistered(offset int, limit int, Domain_ string, Path_ string, Registered_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and path = ? and registered = ?", Domain_, Path_, Registered_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPathAndLastUpdated Get Blogses via DomainAndPathAndLastUpdated
func GetBlogsesByDomainAndPathAndLastUpdated(offset int, limit int, Domain_ string, Path_ string, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and path = ? and last_updated = ?", Domain_, Path_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPathAndPublic Get Blogses via DomainAndPathAndPublic
func GetBlogsesByDomainAndPathAndPublic(offset int, limit int, Domain_ string, Path_ string, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and path = ? and public = ?", Domain_, Path_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPathAndArchived Get Blogses via DomainAndPathAndArchived
func GetBlogsesByDomainAndPathAndArchived(offset int, limit int, Domain_ string, Path_ string, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and path = ? and archived = ?", Domain_, Path_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPathAndMature Get Blogses via DomainAndPathAndMature
func GetBlogsesByDomainAndPathAndMature(offset int, limit int, Domain_ string, Path_ string, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and path = ? and mature = ?", Domain_, Path_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPathAndSpam Get Blogses via DomainAndPathAndSpam
func GetBlogsesByDomainAndPathAndSpam(offset int, limit int, Domain_ string, Path_ string, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and path = ? and spam = ?", Domain_, Path_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPathAndDeleted Get Blogses via DomainAndPathAndDeleted
func GetBlogsesByDomainAndPathAndDeleted(offset int, limit int, Domain_ string, Path_ string, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and path = ? and deleted = ?", Domain_, Path_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPathAndLangId Get Blogses via DomainAndPathAndLangId
func GetBlogsesByDomainAndPathAndLangId(offset int, limit int, Domain_ string, Path_ string, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and path = ? and lang_id = ?", Domain_, Path_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndRegisteredAndLastUpdated Get Blogses via DomainAndRegisteredAndLastUpdated
func GetBlogsesByDomainAndRegisteredAndLastUpdated(offset int, limit int, Domain_ string, Registered_ time.Time, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and registered = ? and last_updated = ?", Domain_, Registered_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndRegisteredAndPublic Get Blogses via DomainAndRegisteredAndPublic
func GetBlogsesByDomainAndRegisteredAndPublic(offset int, limit int, Domain_ string, Registered_ time.Time, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and registered = ? and public = ?", Domain_, Registered_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndRegisteredAndArchived Get Blogses via DomainAndRegisteredAndArchived
func GetBlogsesByDomainAndRegisteredAndArchived(offset int, limit int, Domain_ string, Registered_ time.Time, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and registered = ? and archived = ?", Domain_, Registered_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndRegisteredAndMature Get Blogses via DomainAndRegisteredAndMature
func GetBlogsesByDomainAndRegisteredAndMature(offset int, limit int, Domain_ string, Registered_ time.Time, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and registered = ? and mature = ?", Domain_, Registered_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndRegisteredAndSpam Get Blogses via DomainAndRegisteredAndSpam
func GetBlogsesByDomainAndRegisteredAndSpam(offset int, limit int, Domain_ string, Registered_ time.Time, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and registered = ? and spam = ?", Domain_, Registered_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndRegisteredAndDeleted Get Blogses via DomainAndRegisteredAndDeleted
func GetBlogsesByDomainAndRegisteredAndDeleted(offset int, limit int, Domain_ string, Registered_ time.Time, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and registered = ? and deleted = ?", Domain_, Registered_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndRegisteredAndLangId Get Blogses via DomainAndRegisteredAndLangId
func GetBlogsesByDomainAndRegisteredAndLangId(offset int, limit int, Domain_ string, Registered_ time.Time, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and registered = ? and lang_id = ?", Domain_, Registered_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndLastUpdatedAndPublic Get Blogses via DomainAndLastUpdatedAndPublic
func GetBlogsesByDomainAndLastUpdatedAndPublic(offset int, limit int, Domain_ string, LastUpdated_ time.Time, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and last_updated = ? and public = ?", Domain_, LastUpdated_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndLastUpdatedAndArchived Get Blogses via DomainAndLastUpdatedAndArchived
func GetBlogsesByDomainAndLastUpdatedAndArchived(offset int, limit int, Domain_ string, LastUpdated_ time.Time, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and last_updated = ? and archived = ?", Domain_, LastUpdated_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndLastUpdatedAndMature Get Blogses via DomainAndLastUpdatedAndMature
func GetBlogsesByDomainAndLastUpdatedAndMature(offset int, limit int, Domain_ string, LastUpdated_ time.Time, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and last_updated = ? and mature = ?", Domain_, LastUpdated_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndLastUpdatedAndSpam Get Blogses via DomainAndLastUpdatedAndSpam
func GetBlogsesByDomainAndLastUpdatedAndSpam(offset int, limit int, Domain_ string, LastUpdated_ time.Time, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and last_updated = ? and spam = ?", Domain_, LastUpdated_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndLastUpdatedAndDeleted Get Blogses via DomainAndLastUpdatedAndDeleted
func GetBlogsesByDomainAndLastUpdatedAndDeleted(offset int, limit int, Domain_ string, LastUpdated_ time.Time, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and last_updated = ? and deleted = ?", Domain_, LastUpdated_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndLastUpdatedAndLangId Get Blogses via DomainAndLastUpdatedAndLangId
func GetBlogsesByDomainAndLastUpdatedAndLangId(offset int, limit int, Domain_ string, LastUpdated_ time.Time, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and last_updated = ? and lang_id = ?", Domain_, LastUpdated_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPublicAndArchived Get Blogses via DomainAndPublicAndArchived
func GetBlogsesByDomainAndPublicAndArchived(offset int, limit int, Domain_ string, Public_ int, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and public = ? and archived = ?", Domain_, Public_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPublicAndMature Get Blogses via DomainAndPublicAndMature
func GetBlogsesByDomainAndPublicAndMature(offset int, limit int, Domain_ string, Public_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and public = ? and mature = ?", Domain_, Public_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPublicAndSpam Get Blogses via DomainAndPublicAndSpam
func GetBlogsesByDomainAndPublicAndSpam(offset int, limit int, Domain_ string, Public_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and public = ? and spam = ?", Domain_, Public_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPublicAndDeleted Get Blogses via DomainAndPublicAndDeleted
func GetBlogsesByDomainAndPublicAndDeleted(offset int, limit int, Domain_ string, Public_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and public = ? and deleted = ?", Domain_, Public_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPublicAndLangId Get Blogses via DomainAndPublicAndLangId
func GetBlogsesByDomainAndPublicAndLangId(offset int, limit int, Domain_ string, Public_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and public = ? and lang_id = ?", Domain_, Public_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndArchivedAndMature Get Blogses via DomainAndArchivedAndMature
func GetBlogsesByDomainAndArchivedAndMature(offset int, limit int, Domain_ string, Archived_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and archived = ? and mature = ?", Domain_, Archived_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndArchivedAndSpam Get Blogses via DomainAndArchivedAndSpam
func GetBlogsesByDomainAndArchivedAndSpam(offset int, limit int, Domain_ string, Archived_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and archived = ? and spam = ?", Domain_, Archived_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndArchivedAndDeleted Get Blogses via DomainAndArchivedAndDeleted
func GetBlogsesByDomainAndArchivedAndDeleted(offset int, limit int, Domain_ string, Archived_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and archived = ? and deleted = ?", Domain_, Archived_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndArchivedAndLangId Get Blogses via DomainAndArchivedAndLangId
func GetBlogsesByDomainAndArchivedAndLangId(offset int, limit int, Domain_ string, Archived_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and archived = ? and lang_id = ?", Domain_, Archived_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndMatureAndSpam Get Blogses via DomainAndMatureAndSpam
func GetBlogsesByDomainAndMatureAndSpam(offset int, limit int, Domain_ string, Mature_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and mature = ? and spam = ?", Domain_, Mature_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndMatureAndDeleted Get Blogses via DomainAndMatureAndDeleted
func GetBlogsesByDomainAndMatureAndDeleted(offset int, limit int, Domain_ string, Mature_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and mature = ? and deleted = ?", Domain_, Mature_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndMatureAndLangId Get Blogses via DomainAndMatureAndLangId
func GetBlogsesByDomainAndMatureAndLangId(offset int, limit int, Domain_ string, Mature_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and mature = ? and lang_id = ?", Domain_, Mature_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndSpamAndDeleted Get Blogses via DomainAndSpamAndDeleted
func GetBlogsesByDomainAndSpamAndDeleted(offset int, limit int, Domain_ string, Spam_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and spam = ? and deleted = ?", Domain_, Spam_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndSpamAndLangId Get Blogses via DomainAndSpamAndLangId
func GetBlogsesByDomainAndSpamAndLangId(offset int, limit int, Domain_ string, Spam_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and spam = ? and lang_id = ?", Domain_, Spam_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndDeletedAndLangId Get Blogses via DomainAndDeletedAndLangId
func GetBlogsesByDomainAndDeletedAndLangId(offset int, limit int, Domain_ string, Deleted_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and deleted = ? and lang_id = ?", Domain_, Deleted_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndRegisteredAndLastUpdated Get Blogses via PathAndRegisteredAndLastUpdated
func GetBlogsesByPathAndRegisteredAndLastUpdated(offset int, limit int, Path_ string, Registered_ time.Time, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and registered = ? and last_updated = ?", Path_, Registered_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndRegisteredAndPublic Get Blogses via PathAndRegisteredAndPublic
func GetBlogsesByPathAndRegisteredAndPublic(offset int, limit int, Path_ string, Registered_ time.Time, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and registered = ? and public = ?", Path_, Registered_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndRegisteredAndArchived Get Blogses via PathAndRegisteredAndArchived
func GetBlogsesByPathAndRegisteredAndArchived(offset int, limit int, Path_ string, Registered_ time.Time, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and registered = ? and archived = ?", Path_, Registered_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndRegisteredAndMature Get Blogses via PathAndRegisteredAndMature
func GetBlogsesByPathAndRegisteredAndMature(offset int, limit int, Path_ string, Registered_ time.Time, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and registered = ? and mature = ?", Path_, Registered_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndRegisteredAndSpam Get Blogses via PathAndRegisteredAndSpam
func GetBlogsesByPathAndRegisteredAndSpam(offset int, limit int, Path_ string, Registered_ time.Time, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and registered = ? and spam = ?", Path_, Registered_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndRegisteredAndDeleted Get Blogses via PathAndRegisteredAndDeleted
func GetBlogsesByPathAndRegisteredAndDeleted(offset int, limit int, Path_ string, Registered_ time.Time, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and registered = ? and deleted = ?", Path_, Registered_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndRegisteredAndLangId Get Blogses via PathAndRegisteredAndLangId
func GetBlogsesByPathAndRegisteredAndLangId(offset int, limit int, Path_ string, Registered_ time.Time, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and registered = ? and lang_id = ?", Path_, Registered_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndLastUpdatedAndPublic Get Blogses via PathAndLastUpdatedAndPublic
func GetBlogsesByPathAndLastUpdatedAndPublic(offset int, limit int, Path_ string, LastUpdated_ time.Time, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and last_updated = ? and public = ?", Path_, LastUpdated_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndLastUpdatedAndArchived Get Blogses via PathAndLastUpdatedAndArchived
func GetBlogsesByPathAndLastUpdatedAndArchived(offset int, limit int, Path_ string, LastUpdated_ time.Time, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and last_updated = ? and archived = ?", Path_, LastUpdated_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndLastUpdatedAndMature Get Blogses via PathAndLastUpdatedAndMature
func GetBlogsesByPathAndLastUpdatedAndMature(offset int, limit int, Path_ string, LastUpdated_ time.Time, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and last_updated = ? and mature = ?", Path_, LastUpdated_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndLastUpdatedAndSpam Get Blogses via PathAndLastUpdatedAndSpam
func GetBlogsesByPathAndLastUpdatedAndSpam(offset int, limit int, Path_ string, LastUpdated_ time.Time, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and last_updated = ? and spam = ?", Path_, LastUpdated_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndLastUpdatedAndDeleted Get Blogses via PathAndLastUpdatedAndDeleted
func GetBlogsesByPathAndLastUpdatedAndDeleted(offset int, limit int, Path_ string, LastUpdated_ time.Time, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and last_updated = ? and deleted = ?", Path_, LastUpdated_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndLastUpdatedAndLangId Get Blogses via PathAndLastUpdatedAndLangId
func GetBlogsesByPathAndLastUpdatedAndLangId(offset int, limit int, Path_ string, LastUpdated_ time.Time, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and last_updated = ? and lang_id = ?", Path_, LastUpdated_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndPublicAndArchived Get Blogses via PathAndPublicAndArchived
func GetBlogsesByPathAndPublicAndArchived(offset int, limit int, Path_ string, Public_ int, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and public = ? and archived = ?", Path_, Public_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndPublicAndMature Get Blogses via PathAndPublicAndMature
func GetBlogsesByPathAndPublicAndMature(offset int, limit int, Path_ string, Public_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and public = ? and mature = ?", Path_, Public_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndPublicAndSpam Get Blogses via PathAndPublicAndSpam
func GetBlogsesByPathAndPublicAndSpam(offset int, limit int, Path_ string, Public_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and public = ? and spam = ?", Path_, Public_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndPublicAndDeleted Get Blogses via PathAndPublicAndDeleted
func GetBlogsesByPathAndPublicAndDeleted(offset int, limit int, Path_ string, Public_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and public = ? and deleted = ?", Path_, Public_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndPublicAndLangId Get Blogses via PathAndPublicAndLangId
func GetBlogsesByPathAndPublicAndLangId(offset int, limit int, Path_ string, Public_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and public = ? and lang_id = ?", Path_, Public_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndArchivedAndMature Get Blogses via PathAndArchivedAndMature
func GetBlogsesByPathAndArchivedAndMature(offset int, limit int, Path_ string, Archived_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and archived = ? and mature = ?", Path_, Archived_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndArchivedAndSpam Get Blogses via PathAndArchivedAndSpam
func GetBlogsesByPathAndArchivedAndSpam(offset int, limit int, Path_ string, Archived_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and archived = ? and spam = ?", Path_, Archived_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndArchivedAndDeleted Get Blogses via PathAndArchivedAndDeleted
func GetBlogsesByPathAndArchivedAndDeleted(offset int, limit int, Path_ string, Archived_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and archived = ? and deleted = ?", Path_, Archived_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndArchivedAndLangId Get Blogses via PathAndArchivedAndLangId
func GetBlogsesByPathAndArchivedAndLangId(offset int, limit int, Path_ string, Archived_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and archived = ? and lang_id = ?", Path_, Archived_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndMatureAndSpam Get Blogses via PathAndMatureAndSpam
func GetBlogsesByPathAndMatureAndSpam(offset int, limit int, Path_ string, Mature_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and mature = ? and spam = ?", Path_, Mature_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndMatureAndDeleted Get Blogses via PathAndMatureAndDeleted
func GetBlogsesByPathAndMatureAndDeleted(offset int, limit int, Path_ string, Mature_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and mature = ? and deleted = ?", Path_, Mature_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndMatureAndLangId Get Blogses via PathAndMatureAndLangId
func GetBlogsesByPathAndMatureAndLangId(offset int, limit int, Path_ string, Mature_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and mature = ? and lang_id = ?", Path_, Mature_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndSpamAndDeleted Get Blogses via PathAndSpamAndDeleted
func GetBlogsesByPathAndSpamAndDeleted(offset int, limit int, Path_ string, Spam_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and spam = ? and deleted = ?", Path_, Spam_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndSpamAndLangId Get Blogses via PathAndSpamAndLangId
func GetBlogsesByPathAndSpamAndLangId(offset int, limit int, Path_ string, Spam_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and spam = ? and lang_id = ?", Path_, Spam_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndDeletedAndLangId Get Blogses via PathAndDeletedAndLangId
func GetBlogsesByPathAndDeletedAndLangId(offset int, limit int, Path_ string, Deleted_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and deleted = ? and lang_id = ?", Path_, Deleted_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndLastUpdatedAndPublic Get Blogses via RegisteredAndLastUpdatedAndPublic
func GetBlogsesByRegisteredAndLastUpdatedAndPublic(offset int, limit int, Registered_ time.Time, LastUpdated_ time.Time, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and last_updated = ? and public = ?", Registered_, LastUpdated_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndLastUpdatedAndArchived Get Blogses via RegisteredAndLastUpdatedAndArchived
func GetBlogsesByRegisteredAndLastUpdatedAndArchived(offset int, limit int, Registered_ time.Time, LastUpdated_ time.Time, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and last_updated = ? and archived = ?", Registered_, LastUpdated_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndLastUpdatedAndMature Get Blogses via RegisteredAndLastUpdatedAndMature
func GetBlogsesByRegisteredAndLastUpdatedAndMature(offset int, limit int, Registered_ time.Time, LastUpdated_ time.Time, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and last_updated = ? and mature = ?", Registered_, LastUpdated_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndLastUpdatedAndSpam Get Blogses via RegisteredAndLastUpdatedAndSpam
func GetBlogsesByRegisteredAndLastUpdatedAndSpam(offset int, limit int, Registered_ time.Time, LastUpdated_ time.Time, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and last_updated = ? and spam = ?", Registered_, LastUpdated_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndLastUpdatedAndDeleted Get Blogses via RegisteredAndLastUpdatedAndDeleted
func GetBlogsesByRegisteredAndLastUpdatedAndDeleted(offset int, limit int, Registered_ time.Time, LastUpdated_ time.Time, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and last_updated = ? and deleted = ?", Registered_, LastUpdated_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndLastUpdatedAndLangId Get Blogses via RegisteredAndLastUpdatedAndLangId
func GetBlogsesByRegisteredAndLastUpdatedAndLangId(offset int, limit int, Registered_ time.Time, LastUpdated_ time.Time, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and last_updated = ? and lang_id = ?", Registered_, LastUpdated_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndPublicAndArchived Get Blogses via RegisteredAndPublicAndArchived
func GetBlogsesByRegisteredAndPublicAndArchived(offset int, limit int, Registered_ time.Time, Public_ int, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and public = ? and archived = ?", Registered_, Public_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndPublicAndMature Get Blogses via RegisteredAndPublicAndMature
func GetBlogsesByRegisteredAndPublicAndMature(offset int, limit int, Registered_ time.Time, Public_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and public = ? and mature = ?", Registered_, Public_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndPublicAndSpam Get Blogses via RegisteredAndPublicAndSpam
func GetBlogsesByRegisteredAndPublicAndSpam(offset int, limit int, Registered_ time.Time, Public_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and public = ? and spam = ?", Registered_, Public_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndPublicAndDeleted Get Blogses via RegisteredAndPublicAndDeleted
func GetBlogsesByRegisteredAndPublicAndDeleted(offset int, limit int, Registered_ time.Time, Public_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and public = ? and deleted = ?", Registered_, Public_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndPublicAndLangId Get Blogses via RegisteredAndPublicAndLangId
func GetBlogsesByRegisteredAndPublicAndLangId(offset int, limit int, Registered_ time.Time, Public_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and public = ? and lang_id = ?", Registered_, Public_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndArchivedAndMature Get Blogses via RegisteredAndArchivedAndMature
func GetBlogsesByRegisteredAndArchivedAndMature(offset int, limit int, Registered_ time.Time, Archived_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and archived = ? and mature = ?", Registered_, Archived_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndArchivedAndSpam Get Blogses via RegisteredAndArchivedAndSpam
func GetBlogsesByRegisteredAndArchivedAndSpam(offset int, limit int, Registered_ time.Time, Archived_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and archived = ? and spam = ?", Registered_, Archived_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndArchivedAndDeleted Get Blogses via RegisteredAndArchivedAndDeleted
func GetBlogsesByRegisteredAndArchivedAndDeleted(offset int, limit int, Registered_ time.Time, Archived_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and archived = ? and deleted = ?", Registered_, Archived_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndArchivedAndLangId Get Blogses via RegisteredAndArchivedAndLangId
func GetBlogsesByRegisteredAndArchivedAndLangId(offset int, limit int, Registered_ time.Time, Archived_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and archived = ? and lang_id = ?", Registered_, Archived_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndMatureAndSpam Get Blogses via RegisteredAndMatureAndSpam
func GetBlogsesByRegisteredAndMatureAndSpam(offset int, limit int, Registered_ time.Time, Mature_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and mature = ? and spam = ?", Registered_, Mature_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndMatureAndDeleted Get Blogses via RegisteredAndMatureAndDeleted
func GetBlogsesByRegisteredAndMatureAndDeleted(offset int, limit int, Registered_ time.Time, Mature_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and mature = ? and deleted = ?", Registered_, Mature_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndMatureAndLangId Get Blogses via RegisteredAndMatureAndLangId
func GetBlogsesByRegisteredAndMatureAndLangId(offset int, limit int, Registered_ time.Time, Mature_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and mature = ? and lang_id = ?", Registered_, Mature_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndSpamAndDeleted Get Blogses via RegisteredAndSpamAndDeleted
func GetBlogsesByRegisteredAndSpamAndDeleted(offset int, limit int, Registered_ time.Time, Spam_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and spam = ? and deleted = ?", Registered_, Spam_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndSpamAndLangId Get Blogses via RegisteredAndSpamAndLangId
func GetBlogsesByRegisteredAndSpamAndLangId(offset int, limit int, Registered_ time.Time, Spam_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and spam = ? and lang_id = ?", Registered_, Spam_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndDeletedAndLangId Get Blogses via RegisteredAndDeletedAndLangId
func GetBlogsesByRegisteredAndDeletedAndLangId(offset int, limit int, Registered_ time.Time, Deleted_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and deleted = ? and lang_id = ?", Registered_, Deleted_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndPublicAndArchived Get Blogses via LastUpdatedAndPublicAndArchived
func GetBlogsesByLastUpdatedAndPublicAndArchived(offset int, limit int, LastUpdated_ time.Time, Public_ int, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and public = ? and archived = ?", LastUpdated_, Public_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndPublicAndMature Get Blogses via LastUpdatedAndPublicAndMature
func GetBlogsesByLastUpdatedAndPublicAndMature(offset int, limit int, LastUpdated_ time.Time, Public_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and public = ? and mature = ?", LastUpdated_, Public_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndPublicAndSpam Get Blogses via LastUpdatedAndPublicAndSpam
func GetBlogsesByLastUpdatedAndPublicAndSpam(offset int, limit int, LastUpdated_ time.Time, Public_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and public = ? and spam = ?", LastUpdated_, Public_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndPublicAndDeleted Get Blogses via LastUpdatedAndPublicAndDeleted
func GetBlogsesByLastUpdatedAndPublicAndDeleted(offset int, limit int, LastUpdated_ time.Time, Public_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and public = ? and deleted = ?", LastUpdated_, Public_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndPublicAndLangId Get Blogses via LastUpdatedAndPublicAndLangId
func GetBlogsesByLastUpdatedAndPublicAndLangId(offset int, limit int, LastUpdated_ time.Time, Public_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and public = ? and lang_id = ?", LastUpdated_, Public_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndArchivedAndMature Get Blogses via LastUpdatedAndArchivedAndMature
func GetBlogsesByLastUpdatedAndArchivedAndMature(offset int, limit int, LastUpdated_ time.Time, Archived_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and archived = ? and mature = ?", LastUpdated_, Archived_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndArchivedAndSpam Get Blogses via LastUpdatedAndArchivedAndSpam
func GetBlogsesByLastUpdatedAndArchivedAndSpam(offset int, limit int, LastUpdated_ time.Time, Archived_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and archived = ? and spam = ?", LastUpdated_, Archived_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndArchivedAndDeleted Get Blogses via LastUpdatedAndArchivedAndDeleted
func GetBlogsesByLastUpdatedAndArchivedAndDeleted(offset int, limit int, LastUpdated_ time.Time, Archived_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and archived = ? and deleted = ?", LastUpdated_, Archived_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndArchivedAndLangId Get Blogses via LastUpdatedAndArchivedAndLangId
func GetBlogsesByLastUpdatedAndArchivedAndLangId(offset int, limit int, LastUpdated_ time.Time, Archived_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and archived = ? and lang_id = ?", LastUpdated_, Archived_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndMatureAndSpam Get Blogses via LastUpdatedAndMatureAndSpam
func GetBlogsesByLastUpdatedAndMatureAndSpam(offset int, limit int, LastUpdated_ time.Time, Mature_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and mature = ? and spam = ?", LastUpdated_, Mature_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndMatureAndDeleted Get Blogses via LastUpdatedAndMatureAndDeleted
func GetBlogsesByLastUpdatedAndMatureAndDeleted(offset int, limit int, LastUpdated_ time.Time, Mature_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and mature = ? and deleted = ?", LastUpdated_, Mature_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndMatureAndLangId Get Blogses via LastUpdatedAndMatureAndLangId
func GetBlogsesByLastUpdatedAndMatureAndLangId(offset int, limit int, LastUpdated_ time.Time, Mature_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and mature = ? and lang_id = ?", LastUpdated_, Mature_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndSpamAndDeleted Get Blogses via LastUpdatedAndSpamAndDeleted
func GetBlogsesByLastUpdatedAndSpamAndDeleted(offset int, limit int, LastUpdated_ time.Time, Spam_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and spam = ? and deleted = ?", LastUpdated_, Spam_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndSpamAndLangId Get Blogses via LastUpdatedAndSpamAndLangId
func GetBlogsesByLastUpdatedAndSpamAndLangId(offset int, limit int, LastUpdated_ time.Time, Spam_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and spam = ? and lang_id = ?", LastUpdated_, Spam_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndDeletedAndLangId Get Blogses via LastUpdatedAndDeletedAndLangId
func GetBlogsesByLastUpdatedAndDeletedAndLangId(offset int, limit int, LastUpdated_ time.Time, Deleted_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and deleted = ? and lang_id = ?", LastUpdated_, Deleted_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndArchivedAndMature Get Blogses via PublicAndArchivedAndMature
func GetBlogsesByPublicAndArchivedAndMature(offset int, limit int, Public_ int, Archived_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and archived = ? and mature = ?", Public_, Archived_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndArchivedAndSpam Get Blogses via PublicAndArchivedAndSpam
func GetBlogsesByPublicAndArchivedAndSpam(offset int, limit int, Public_ int, Archived_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and archived = ? and spam = ?", Public_, Archived_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndArchivedAndDeleted Get Blogses via PublicAndArchivedAndDeleted
func GetBlogsesByPublicAndArchivedAndDeleted(offset int, limit int, Public_ int, Archived_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and archived = ? and deleted = ?", Public_, Archived_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndArchivedAndLangId Get Blogses via PublicAndArchivedAndLangId
func GetBlogsesByPublicAndArchivedAndLangId(offset int, limit int, Public_ int, Archived_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and archived = ? and lang_id = ?", Public_, Archived_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndMatureAndSpam Get Blogses via PublicAndMatureAndSpam
func GetBlogsesByPublicAndMatureAndSpam(offset int, limit int, Public_ int, Mature_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and mature = ? and spam = ?", Public_, Mature_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndMatureAndDeleted Get Blogses via PublicAndMatureAndDeleted
func GetBlogsesByPublicAndMatureAndDeleted(offset int, limit int, Public_ int, Mature_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and mature = ? and deleted = ?", Public_, Mature_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndMatureAndLangId Get Blogses via PublicAndMatureAndLangId
func GetBlogsesByPublicAndMatureAndLangId(offset int, limit int, Public_ int, Mature_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and mature = ? and lang_id = ?", Public_, Mature_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndSpamAndDeleted Get Blogses via PublicAndSpamAndDeleted
func GetBlogsesByPublicAndSpamAndDeleted(offset int, limit int, Public_ int, Spam_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and spam = ? and deleted = ?", Public_, Spam_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndSpamAndLangId Get Blogses via PublicAndSpamAndLangId
func GetBlogsesByPublicAndSpamAndLangId(offset int, limit int, Public_ int, Spam_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and spam = ? and lang_id = ?", Public_, Spam_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndDeletedAndLangId Get Blogses via PublicAndDeletedAndLangId
func GetBlogsesByPublicAndDeletedAndLangId(offset int, limit int, Public_ int, Deleted_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and deleted = ? and lang_id = ?", Public_, Deleted_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByArchivedAndMatureAndSpam Get Blogses via ArchivedAndMatureAndSpam
func GetBlogsesByArchivedAndMatureAndSpam(offset int, limit int, Archived_ int, Mature_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("archived = ? and mature = ? and spam = ?", Archived_, Mature_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByArchivedAndMatureAndDeleted Get Blogses via ArchivedAndMatureAndDeleted
func GetBlogsesByArchivedAndMatureAndDeleted(offset int, limit int, Archived_ int, Mature_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("archived = ? and mature = ? and deleted = ?", Archived_, Mature_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByArchivedAndMatureAndLangId Get Blogses via ArchivedAndMatureAndLangId
func GetBlogsesByArchivedAndMatureAndLangId(offset int, limit int, Archived_ int, Mature_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("archived = ? and mature = ? and lang_id = ?", Archived_, Mature_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByArchivedAndSpamAndDeleted Get Blogses via ArchivedAndSpamAndDeleted
func GetBlogsesByArchivedAndSpamAndDeleted(offset int, limit int, Archived_ int, Spam_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("archived = ? and spam = ? and deleted = ?", Archived_, Spam_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByArchivedAndSpamAndLangId Get Blogses via ArchivedAndSpamAndLangId
func GetBlogsesByArchivedAndSpamAndLangId(offset int, limit int, Archived_ int, Spam_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("archived = ? and spam = ? and lang_id = ?", Archived_, Spam_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByArchivedAndDeletedAndLangId Get Blogses via ArchivedAndDeletedAndLangId
func GetBlogsesByArchivedAndDeletedAndLangId(offset int, limit int, Archived_ int, Deleted_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("archived = ? and deleted = ? and lang_id = ?", Archived_, Deleted_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByMatureAndSpamAndDeleted Get Blogses via MatureAndSpamAndDeleted
func GetBlogsesByMatureAndSpamAndDeleted(offset int, limit int, Mature_ int, Spam_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("mature = ? and spam = ? and deleted = ?", Mature_, Spam_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByMatureAndSpamAndLangId Get Blogses via MatureAndSpamAndLangId
func GetBlogsesByMatureAndSpamAndLangId(offset int, limit int, Mature_ int, Spam_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("mature = ? and spam = ? and lang_id = ?", Mature_, Spam_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByMatureAndDeletedAndLangId Get Blogses via MatureAndDeletedAndLangId
func GetBlogsesByMatureAndDeletedAndLangId(offset int, limit int, Mature_ int, Deleted_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("mature = ? and deleted = ? and lang_id = ?", Mature_, Deleted_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySpamAndDeletedAndLangId Get Blogses via SpamAndDeletedAndLangId
func GetBlogsesBySpamAndDeletedAndLangId(offset int, limit int, Spam_ int, Deleted_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("spam = ? and deleted = ? and lang_id = ?", Spam_, Deleted_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSiteId Get Blogses via BlogIdAndSiteId
func GetBlogsesByBlogIdAndSiteId(offset int, limit int, BlogId_ int64, SiteId_ int64) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and site_id = ?", BlogId_, SiteId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndDomain Get Blogses via BlogIdAndDomain
func GetBlogsesByBlogIdAndDomain(offset int, limit int, BlogId_ int64, Domain_ string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and domain = ?", BlogId_, Domain_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPath Get Blogses via BlogIdAndPath
func GetBlogsesByBlogIdAndPath(offset int, limit int, BlogId_ int64, Path_ string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and path = ?", BlogId_, Path_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndRegistered Get Blogses via BlogIdAndRegistered
func GetBlogsesByBlogIdAndRegistered(offset int, limit int, BlogId_ int64, Registered_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and registered = ?", BlogId_, Registered_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndLastUpdated Get Blogses via BlogIdAndLastUpdated
func GetBlogsesByBlogIdAndLastUpdated(offset int, limit int, BlogId_ int64, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and last_updated = ?", BlogId_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndPublic Get Blogses via BlogIdAndPublic
func GetBlogsesByBlogIdAndPublic(offset int, limit int, BlogId_ int64, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and public = ?", BlogId_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndArchived Get Blogses via BlogIdAndArchived
func GetBlogsesByBlogIdAndArchived(offset int, limit int, BlogId_ int64, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and archived = ?", BlogId_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndMature Get Blogses via BlogIdAndMature
func GetBlogsesByBlogIdAndMature(offset int, limit int, BlogId_ int64, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and mature = ?", BlogId_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndSpam Get Blogses via BlogIdAndSpam
func GetBlogsesByBlogIdAndSpam(offset int, limit int, BlogId_ int64, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and spam = ?", BlogId_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndDeleted Get Blogses via BlogIdAndDeleted
func GetBlogsesByBlogIdAndDeleted(offset int, limit int, BlogId_ int64, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and deleted = ?", BlogId_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByBlogIdAndLangId Get Blogses via BlogIdAndLangId
func GetBlogsesByBlogIdAndLangId(offset int, limit int, BlogId_ int64, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ? and lang_id = ?", BlogId_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndDomain Get Blogses via SiteIdAndDomain
func GetBlogsesBySiteIdAndDomain(offset int, limit int, SiteId_ int64, Domain_ string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and domain = ?", SiteId_, Domain_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPath Get Blogses via SiteIdAndPath
func GetBlogsesBySiteIdAndPath(offset int, limit int, SiteId_ int64, Path_ string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and path = ?", SiteId_, Path_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndRegistered Get Blogses via SiteIdAndRegistered
func GetBlogsesBySiteIdAndRegistered(offset int, limit int, SiteId_ int64, Registered_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and registered = ?", SiteId_, Registered_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndLastUpdated Get Blogses via SiteIdAndLastUpdated
func GetBlogsesBySiteIdAndLastUpdated(offset int, limit int, SiteId_ int64, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and last_updated = ?", SiteId_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndPublic Get Blogses via SiteIdAndPublic
func GetBlogsesBySiteIdAndPublic(offset int, limit int, SiteId_ int64, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and public = ?", SiteId_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndArchived Get Blogses via SiteIdAndArchived
func GetBlogsesBySiteIdAndArchived(offset int, limit int, SiteId_ int64, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and archived = ?", SiteId_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndMature Get Blogses via SiteIdAndMature
func GetBlogsesBySiteIdAndMature(offset int, limit int, SiteId_ int64, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and mature = ?", SiteId_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndSpam Get Blogses via SiteIdAndSpam
func GetBlogsesBySiteIdAndSpam(offset int, limit int, SiteId_ int64, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and spam = ?", SiteId_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndDeleted Get Blogses via SiteIdAndDeleted
func GetBlogsesBySiteIdAndDeleted(offset int, limit int, SiteId_ int64, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and deleted = ?", SiteId_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySiteIdAndLangId Get Blogses via SiteIdAndLangId
func GetBlogsesBySiteIdAndLangId(offset int, limit int, SiteId_ int64, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ? and lang_id = ?", SiteId_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPath Get Blogses via DomainAndPath
func GetBlogsesByDomainAndPath(offset int, limit int, Domain_ string, Path_ string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and path = ?", Domain_, Path_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndRegistered Get Blogses via DomainAndRegistered
func GetBlogsesByDomainAndRegistered(offset int, limit int, Domain_ string, Registered_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and registered = ?", Domain_, Registered_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndLastUpdated Get Blogses via DomainAndLastUpdated
func GetBlogsesByDomainAndLastUpdated(offset int, limit int, Domain_ string, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and last_updated = ?", Domain_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndPublic Get Blogses via DomainAndPublic
func GetBlogsesByDomainAndPublic(offset int, limit int, Domain_ string, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and public = ?", Domain_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndArchived Get Blogses via DomainAndArchived
func GetBlogsesByDomainAndArchived(offset int, limit int, Domain_ string, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and archived = ?", Domain_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndMature Get Blogses via DomainAndMature
func GetBlogsesByDomainAndMature(offset int, limit int, Domain_ string, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and mature = ?", Domain_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndSpam Get Blogses via DomainAndSpam
func GetBlogsesByDomainAndSpam(offset int, limit int, Domain_ string, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and spam = ?", Domain_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndDeleted Get Blogses via DomainAndDeleted
func GetBlogsesByDomainAndDeleted(offset int, limit int, Domain_ string, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and deleted = ?", Domain_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDomainAndLangId Get Blogses via DomainAndLangId
func GetBlogsesByDomainAndLangId(offset int, limit int, Domain_ string, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ? and lang_id = ?", Domain_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndRegistered Get Blogses via PathAndRegistered
func GetBlogsesByPathAndRegistered(offset int, limit int, Path_ string, Registered_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and registered = ?", Path_, Registered_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndLastUpdated Get Blogses via PathAndLastUpdated
func GetBlogsesByPathAndLastUpdated(offset int, limit int, Path_ string, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and last_updated = ?", Path_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndPublic Get Blogses via PathAndPublic
func GetBlogsesByPathAndPublic(offset int, limit int, Path_ string, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and public = ?", Path_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndArchived Get Blogses via PathAndArchived
func GetBlogsesByPathAndArchived(offset int, limit int, Path_ string, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and archived = ?", Path_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndMature Get Blogses via PathAndMature
func GetBlogsesByPathAndMature(offset int, limit int, Path_ string, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and mature = ?", Path_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndSpam Get Blogses via PathAndSpam
func GetBlogsesByPathAndSpam(offset int, limit int, Path_ string, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and spam = ?", Path_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndDeleted Get Blogses via PathAndDeleted
func GetBlogsesByPathAndDeleted(offset int, limit int, Path_ string, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and deleted = ?", Path_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPathAndLangId Get Blogses via PathAndLangId
func GetBlogsesByPathAndLangId(offset int, limit int, Path_ string, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ? and lang_id = ?", Path_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndLastUpdated Get Blogses via RegisteredAndLastUpdated
func GetBlogsesByRegisteredAndLastUpdated(offset int, limit int, Registered_ time.Time, LastUpdated_ time.Time) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and last_updated = ?", Registered_, LastUpdated_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndPublic Get Blogses via RegisteredAndPublic
func GetBlogsesByRegisteredAndPublic(offset int, limit int, Registered_ time.Time, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and public = ?", Registered_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndArchived Get Blogses via RegisteredAndArchived
func GetBlogsesByRegisteredAndArchived(offset int, limit int, Registered_ time.Time, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and archived = ?", Registered_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndMature Get Blogses via RegisteredAndMature
func GetBlogsesByRegisteredAndMature(offset int, limit int, Registered_ time.Time, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and mature = ?", Registered_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndSpam Get Blogses via RegisteredAndSpam
func GetBlogsesByRegisteredAndSpam(offset int, limit int, Registered_ time.Time, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and spam = ?", Registered_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndDeleted Get Blogses via RegisteredAndDeleted
func GetBlogsesByRegisteredAndDeleted(offset int, limit int, Registered_ time.Time, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and deleted = ?", Registered_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByRegisteredAndLangId Get Blogses via RegisteredAndLangId
func GetBlogsesByRegisteredAndLangId(offset int, limit int, Registered_ time.Time, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ? and lang_id = ?", Registered_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndPublic Get Blogses via LastUpdatedAndPublic
func GetBlogsesByLastUpdatedAndPublic(offset int, limit int, LastUpdated_ time.Time, Public_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and public = ?", LastUpdated_, Public_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndArchived Get Blogses via LastUpdatedAndArchived
func GetBlogsesByLastUpdatedAndArchived(offset int, limit int, LastUpdated_ time.Time, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and archived = ?", LastUpdated_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndMature Get Blogses via LastUpdatedAndMature
func GetBlogsesByLastUpdatedAndMature(offset int, limit int, LastUpdated_ time.Time, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and mature = ?", LastUpdated_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndSpam Get Blogses via LastUpdatedAndSpam
func GetBlogsesByLastUpdatedAndSpam(offset int, limit int, LastUpdated_ time.Time, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and spam = ?", LastUpdated_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndDeleted Get Blogses via LastUpdatedAndDeleted
func GetBlogsesByLastUpdatedAndDeleted(offset int, limit int, LastUpdated_ time.Time, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and deleted = ?", LastUpdated_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByLastUpdatedAndLangId Get Blogses via LastUpdatedAndLangId
func GetBlogsesByLastUpdatedAndLangId(offset int, limit int, LastUpdated_ time.Time, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ? and lang_id = ?", LastUpdated_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndArchived Get Blogses via PublicAndArchived
func GetBlogsesByPublicAndArchived(offset int, limit int, Public_ int, Archived_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and archived = ?", Public_, Archived_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndMature Get Blogses via PublicAndMature
func GetBlogsesByPublicAndMature(offset int, limit int, Public_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and mature = ?", Public_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndSpam Get Blogses via PublicAndSpam
func GetBlogsesByPublicAndSpam(offset int, limit int, Public_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and spam = ?", Public_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndDeleted Get Blogses via PublicAndDeleted
func GetBlogsesByPublicAndDeleted(offset int, limit int, Public_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and deleted = ?", Public_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByPublicAndLangId Get Blogses via PublicAndLangId
func GetBlogsesByPublicAndLangId(offset int, limit int, Public_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ? and lang_id = ?", Public_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByArchivedAndMature Get Blogses via ArchivedAndMature
func GetBlogsesByArchivedAndMature(offset int, limit int, Archived_ int, Mature_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("archived = ? and mature = ?", Archived_, Mature_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByArchivedAndSpam Get Blogses via ArchivedAndSpam
func GetBlogsesByArchivedAndSpam(offset int, limit int, Archived_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("archived = ? and spam = ?", Archived_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByArchivedAndDeleted Get Blogses via ArchivedAndDeleted
func GetBlogsesByArchivedAndDeleted(offset int, limit int, Archived_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("archived = ? and deleted = ?", Archived_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByArchivedAndLangId Get Blogses via ArchivedAndLangId
func GetBlogsesByArchivedAndLangId(offset int, limit int, Archived_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("archived = ? and lang_id = ?", Archived_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByMatureAndSpam Get Blogses via MatureAndSpam
func GetBlogsesByMatureAndSpam(offset int, limit int, Mature_ int, Spam_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("mature = ? and spam = ?", Mature_, Spam_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByMatureAndDeleted Get Blogses via MatureAndDeleted
func GetBlogsesByMatureAndDeleted(offset int, limit int, Mature_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("mature = ? and deleted = ?", Mature_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByMatureAndLangId Get Blogses via MatureAndLangId
func GetBlogsesByMatureAndLangId(offset int, limit int, Mature_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("mature = ? and lang_id = ?", Mature_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySpamAndDeleted Get Blogses via SpamAndDeleted
func GetBlogsesBySpamAndDeleted(offset int, limit int, Spam_ int, Deleted_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("spam = ? and deleted = ?", Spam_, Deleted_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesBySpamAndLangId Get Blogses via SpamAndLangId
func GetBlogsesBySpamAndLangId(offset int, limit int, Spam_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("spam = ? and lang_id = ?", Spam_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesByDeletedAndLangId Get Blogses via DeletedAndLangId
func GetBlogsesByDeletedAndLangId(offset int, limit int, Deleted_ int, LangId_ int) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("deleted = ? and lang_id = ?", Deleted_, LangId_).Limit(limit, offset).Find(_Blogs)
	return _Blogs, err
}

// GetBlogses Get Blogses via field
func GetBlogses(offset int, limit int, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesViaBlogId Get Blogss via BlogId
func GetBlogsesViaBlogId(offset int, limit int, BlogId_ int64, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("blog_id = ?", BlogId_).Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesViaSiteId Get Blogss via SiteId
func GetBlogsesViaSiteId(offset int, limit int, SiteId_ int64, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("site_id = ?", SiteId_).Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesViaDomain Get Blogss via Domain
func GetBlogsesViaDomain(offset int, limit int, Domain_ string, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("domain = ?", Domain_).Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesViaPath Get Blogss via Path
func GetBlogsesViaPath(offset int, limit int, Path_ string, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("path = ?", Path_).Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesViaRegistered Get Blogss via Registered
func GetBlogsesViaRegistered(offset int, limit int, Registered_ time.Time, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("registered = ?", Registered_).Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesViaLastUpdated Get Blogss via LastUpdated
func GetBlogsesViaLastUpdated(offset int, limit int, LastUpdated_ time.Time, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("last_updated = ?", LastUpdated_).Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesViaPublic Get Blogss via Public
func GetBlogsesViaPublic(offset int, limit int, Public_ int, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("public = ?", Public_).Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesViaArchived Get Blogss via Archived
func GetBlogsesViaArchived(offset int, limit int, Archived_ int, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("archived = ?", Archived_).Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesViaMature Get Blogss via Mature
func GetBlogsesViaMature(offset int, limit int, Mature_ int, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("mature = ?", Mature_).Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesViaSpam Get Blogss via Spam
func GetBlogsesViaSpam(offset int, limit int, Spam_ int, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("spam = ?", Spam_).Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesViaDeleted Get Blogss via Deleted
func GetBlogsesViaDeleted(offset int, limit int, Deleted_ int, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("deleted = ?", Deleted_).Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// GetBlogsesViaLangId Get Blogss via LangId
func GetBlogsesViaLangId(offset int, limit int, LangId_ int, field string) (*[]*Blogs, error) {
	var _Blogs = new([]*Blogs)
	err := Engine.Table("blogs").Where("lang_id = ?", LangId_).Limit(limit, offset).Desc(field).Find(_Blogs)
	return _Blogs, err
}

// HasBlogsViaBlogId Has Blogs via BlogId
func HasBlogsViaBlogId(iBlogId int64) bool {
	if has, err := Engine.Where("blog_id = ?", iBlogId).Get(new(Blogs)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogsViaSiteId Has Blogs via SiteId
func HasBlogsViaSiteId(iSiteId int64) bool {
	if has, err := Engine.Where("site_id = ?", iSiteId).Get(new(Blogs)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogsViaDomain Has Blogs via Domain
func HasBlogsViaDomain(iDomain string) bool {
	if has, err := Engine.Where("domain = ?", iDomain).Get(new(Blogs)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogsViaPath Has Blogs via Path
func HasBlogsViaPath(iPath string) bool {
	if has, err := Engine.Where("path = ?", iPath).Get(new(Blogs)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogsViaRegistered Has Blogs via Registered
func HasBlogsViaRegistered(iRegistered time.Time) bool {
	if has, err := Engine.Where("registered = ?", iRegistered).Get(new(Blogs)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogsViaLastUpdated Has Blogs via LastUpdated
func HasBlogsViaLastUpdated(iLastUpdated time.Time) bool {
	if has, err := Engine.Where("last_updated = ?", iLastUpdated).Get(new(Blogs)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogsViaPublic Has Blogs via Public
func HasBlogsViaPublic(iPublic int) bool {
	if has, err := Engine.Where("public = ?", iPublic).Get(new(Blogs)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogsViaArchived Has Blogs via Archived
func HasBlogsViaArchived(iArchived int) bool {
	if has, err := Engine.Where("archived = ?", iArchived).Get(new(Blogs)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogsViaMature Has Blogs via Mature
func HasBlogsViaMature(iMature int) bool {
	if has, err := Engine.Where("mature = ?", iMature).Get(new(Blogs)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogsViaSpam Has Blogs via Spam
func HasBlogsViaSpam(iSpam int) bool {
	if has, err := Engine.Where("spam = ?", iSpam).Get(new(Blogs)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogsViaDeleted Has Blogs via Deleted
func HasBlogsViaDeleted(iDeleted int) bool {
	if has, err := Engine.Where("deleted = ?", iDeleted).Get(new(Blogs)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogsViaLangId Has Blogs via LangId
func HasBlogsViaLangId(iLangId int) bool {
	if has, err := Engine.Where("lang_id = ?", iLangId).Get(new(Blogs)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetBlogsViaBlogId Get Blogs via BlogId
func GetBlogsViaBlogId(iBlogId int64) (*Blogs, error) {
	var _Blogs = &Blogs{BlogId: iBlogId}
	has, err := Engine.Get(_Blogs)
	if has {
		return _Blogs, err
	} else {
		return nil, err
	}
}

// GetBlogsViaSiteId Get Blogs via SiteId
func GetBlogsViaSiteId(iSiteId int64) (*Blogs, error) {
	var _Blogs = &Blogs{SiteId: iSiteId}
	has, err := Engine.Get(_Blogs)
	if has {
		return _Blogs, err
	} else {
		return nil, err
	}
}

// GetBlogsViaDomain Get Blogs via Domain
func GetBlogsViaDomain(iDomain string) (*Blogs, error) {
	var _Blogs = &Blogs{Domain: iDomain}
	has, err := Engine.Get(_Blogs)
	if has {
		return _Blogs, err
	} else {
		return nil, err
	}
}

// GetBlogsViaPath Get Blogs via Path
func GetBlogsViaPath(iPath string) (*Blogs, error) {
	var _Blogs = &Blogs{Path: iPath}
	has, err := Engine.Get(_Blogs)
	if has {
		return _Blogs, err
	} else {
		return nil, err
	}
}

// GetBlogsViaRegistered Get Blogs via Registered
func GetBlogsViaRegistered(iRegistered time.Time) (*Blogs, error) {
	var _Blogs = &Blogs{Registered: iRegistered}
	has, err := Engine.Get(_Blogs)
	if has {
		return _Blogs, err
	} else {
		return nil, err
	}
}

// GetBlogsViaLastUpdated Get Blogs via LastUpdated
func GetBlogsViaLastUpdated(iLastUpdated time.Time) (*Blogs, error) {
	var _Blogs = &Blogs{LastUpdated: iLastUpdated}
	has, err := Engine.Get(_Blogs)
	if has {
		return _Blogs, err
	} else {
		return nil, err
	}
}

// GetBlogsViaPublic Get Blogs via Public
func GetBlogsViaPublic(iPublic int) (*Blogs, error) {
	var _Blogs = &Blogs{Public: iPublic}
	has, err := Engine.Get(_Blogs)
	if has {
		return _Blogs, err
	} else {
		return nil, err
	}
}

// GetBlogsViaArchived Get Blogs via Archived
func GetBlogsViaArchived(iArchived int) (*Blogs, error) {
	var _Blogs = &Blogs{Archived: iArchived}
	has, err := Engine.Get(_Blogs)
	if has {
		return _Blogs, err
	} else {
		return nil, err
	}
}

// GetBlogsViaMature Get Blogs via Mature
func GetBlogsViaMature(iMature int) (*Blogs, error) {
	var _Blogs = &Blogs{Mature: iMature}
	has, err := Engine.Get(_Blogs)
	if has {
		return _Blogs, err
	} else {
		return nil, err
	}
}

// GetBlogsViaSpam Get Blogs via Spam
func GetBlogsViaSpam(iSpam int) (*Blogs, error) {
	var _Blogs = &Blogs{Spam: iSpam}
	has, err := Engine.Get(_Blogs)
	if has {
		return _Blogs, err
	} else {
		return nil, err
	}
}

// GetBlogsViaDeleted Get Blogs via Deleted
func GetBlogsViaDeleted(iDeleted int) (*Blogs, error) {
	var _Blogs = &Blogs{Deleted: iDeleted}
	has, err := Engine.Get(_Blogs)
	if has {
		return _Blogs, err
	} else {
		return nil, err
	}
}

// GetBlogsViaLangId Get Blogs via LangId
func GetBlogsViaLangId(iLangId int) (*Blogs, error) {
	var _Blogs = &Blogs{LangId: iLangId}
	has, err := Engine.Get(_Blogs)
	if has {
		return _Blogs, err
	} else {
		return nil, err
	}
}

// SetBlogsViaBlogId Set Blogs via BlogId
func SetBlogsViaBlogId(iBlogId int64, blogs *Blogs) (int64, error) {
	blogs.BlogId = iBlogId
	return Engine.Insert(blogs)
}

// SetBlogsViaSiteId Set Blogs via SiteId
func SetBlogsViaSiteId(iSiteId int64, blogs *Blogs) (int64, error) {
	blogs.SiteId = iSiteId
	return Engine.Insert(blogs)
}

// SetBlogsViaDomain Set Blogs via Domain
func SetBlogsViaDomain(iDomain string, blogs *Blogs) (int64, error) {
	blogs.Domain = iDomain
	return Engine.Insert(blogs)
}

// SetBlogsViaPath Set Blogs via Path
func SetBlogsViaPath(iPath string, blogs *Blogs) (int64, error) {
	blogs.Path = iPath
	return Engine.Insert(blogs)
}

// SetBlogsViaRegistered Set Blogs via Registered
func SetBlogsViaRegistered(iRegistered time.Time, blogs *Blogs) (int64, error) {
	blogs.Registered = iRegistered
	return Engine.Insert(blogs)
}

// SetBlogsViaLastUpdated Set Blogs via LastUpdated
func SetBlogsViaLastUpdated(iLastUpdated time.Time, blogs *Blogs) (int64, error) {
	blogs.LastUpdated = iLastUpdated
	return Engine.Insert(blogs)
}

// SetBlogsViaPublic Set Blogs via Public
func SetBlogsViaPublic(iPublic int, blogs *Blogs) (int64, error) {
	blogs.Public = iPublic
	return Engine.Insert(blogs)
}

// SetBlogsViaArchived Set Blogs via Archived
func SetBlogsViaArchived(iArchived int, blogs *Blogs) (int64, error) {
	blogs.Archived = iArchived
	return Engine.Insert(blogs)
}

// SetBlogsViaMature Set Blogs via Mature
func SetBlogsViaMature(iMature int, blogs *Blogs) (int64, error) {
	blogs.Mature = iMature
	return Engine.Insert(blogs)
}

// SetBlogsViaSpam Set Blogs via Spam
func SetBlogsViaSpam(iSpam int, blogs *Blogs) (int64, error) {
	blogs.Spam = iSpam
	return Engine.Insert(blogs)
}

// SetBlogsViaDeleted Set Blogs via Deleted
func SetBlogsViaDeleted(iDeleted int, blogs *Blogs) (int64, error) {
	blogs.Deleted = iDeleted
	return Engine.Insert(blogs)
}

// SetBlogsViaLangId Set Blogs via LangId
func SetBlogsViaLangId(iLangId int, blogs *Blogs) (int64, error) {
	blogs.LangId = iLangId
	return Engine.Insert(blogs)
}

// AddBlogs Add Blogs via all columns
func AddBlogs(iBlogId int64, iSiteId int64, iDomain string, iPath string, iRegistered time.Time, iLastUpdated time.Time, iPublic int, iArchived int, iMature int, iSpam int, iDeleted int, iLangId int) error {
	_Blogs := &Blogs{BlogId: iBlogId, SiteId: iSiteId, Domain: iDomain, Path: iPath, Registered: iRegistered, LastUpdated: iLastUpdated, Public: iPublic, Archived: iArchived, Mature: iMature, Spam: iSpam, Deleted: iDeleted, LangId: iLangId}
	if _, err := Engine.Insert(_Blogs); err != nil {
		return err
	} else {
		return nil
	}
}

// PostBlogs Post Blogs via iBlogs
func PostBlogs(iBlogs *Blogs) (int64, error) {
	_, err := Engine.Insert(iBlogs)
	return iBlogs.BlogId, err
}

// PutBlogs Put Blogs
func PutBlogs(iBlogs *Blogs) (int64, error) {
	_, err := Engine.Id(iBlogs.BlogId).Update(iBlogs)
	return iBlogs.BlogId, err
}

// PutBlogsViaBlogId Put Blogs via BlogId
func PutBlogsViaBlogId(BlogId_ int64, iBlogs *Blogs) (int64, error) {
	row, err := Engine.Update(iBlogs, &Blogs{BlogId: BlogId_})
	return row, err
}

// PutBlogsViaSiteId Put Blogs via SiteId
func PutBlogsViaSiteId(SiteId_ int64, iBlogs *Blogs) (int64, error) {
	row, err := Engine.Update(iBlogs, &Blogs{SiteId: SiteId_})
	return row, err
}

// PutBlogsViaDomain Put Blogs via Domain
func PutBlogsViaDomain(Domain_ string, iBlogs *Blogs) (int64, error) {
	row, err := Engine.Update(iBlogs, &Blogs{Domain: Domain_})
	return row, err
}

// PutBlogsViaPath Put Blogs via Path
func PutBlogsViaPath(Path_ string, iBlogs *Blogs) (int64, error) {
	row, err := Engine.Update(iBlogs, &Blogs{Path: Path_})
	return row, err
}

// PutBlogsViaRegistered Put Blogs via Registered
func PutBlogsViaRegistered(Registered_ time.Time, iBlogs *Blogs) (int64, error) {
	row, err := Engine.Update(iBlogs, &Blogs{Registered: Registered_})
	return row, err
}

// PutBlogsViaLastUpdated Put Blogs via LastUpdated
func PutBlogsViaLastUpdated(LastUpdated_ time.Time, iBlogs *Blogs) (int64, error) {
	row, err := Engine.Update(iBlogs, &Blogs{LastUpdated: LastUpdated_})
	return row, err
}

// PutBlogsViaPublic Put Blogs via Public
func PutBlogsViaPublic(Public_ int, iBlogs *Blogs) (int64, error) {
	row, err := Engine.Update(iBlogs, &Blogs{Public: Public_})
	return row, err
}

// PutBlogsViaArchived Put Blogs via Archived
func PutBlogsViaArchived(Archived_ int, iBlogs *Blogs) (int64, error) {
	row, err := Engine.Update(iBlogs, &Blogs{Archived: Archived_})
	return row, err
}

// PutBlogsViaMature Put Blogs via Mature
func PutBlogsViaMature(Mature_ int, iBlogs *Blogs) (int64, error) {
	row, err := Engine.Update(iBlogs, &Blogs{Mature: Mature_})
	return row, err
}

// PutBlogsViaSpam Put Blogs via Spam
func PutBlogsViaSpam(Spam_ int, iBlogs *Blogs) (int64, error) {
	row, err := Engine.Update(iBlogs, &Blogs{Spam: Spam_})
	return row, err
}

// PutBlogsViaDeleted Put Blogs via Deleted
func PutBlogsViaDeleted(Deleted_ int, iBlogs *Blogs) (int64, error) {
	row, err := Engine.Update(iBlogs, &Blogs{Deleted: Deleted_})
	return row, err
}

// PutBlogsViaLangId Put Blogs via LangId
func PutBlogsViaLangId(LangId_ int, iBlogs *Blogs) (int64, error) {
	row, err := Engine.Update(iBlogs, &Blogs{LangId: LangId_})
	return row, err
}

// UpdateBlogsViaBlogId via map[string]interface{}{}
func UpdateBlogsViaBlogId(iBlogId int64, iBlogsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Blogs)).Where("blog_id = ?", iBlogId).Update(iBlogsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogsViaSiteId via map[string]interface{}{}
func UpdateBlogsViaSiteId(iSiteId int64, iBlogsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Blogs)).Where("site_id = ?", iSiteId).Update(iBlogsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogsViaDomain via map[string]interface{}{}
func UpdateBlogsViaDomain(iDomain string, iBlogsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Blogs)).Where("domain = ?", iDomain).Update(iBlogsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogsViaPath via map[string]interface{}{}
func UpdateBlogsViaPath(iPath string, iBlogsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Blogs)).Where("path = ?", iPath).Update(iBlogsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogsViaRegistered via map[string]interface{}{}
func UpdateBlogsViaRegistered(iRegistered time.Time, iBlogsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Blogs)).Where("registered = ?", iRegistered).Update(iBlogsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogsViaLastUpdated via map[string]interface{}{}
func UpdateBlogsViaLastUpdated(iLastUpdated time.Time, iBlogsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Blogs)).Where("last_updated = ?", iLastUpdated).Update(iBlogsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogsViaPublic via map[string]interface{}{}
func UpdateBlogsViaPublic(iPublic int, iBlogsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Blogs)).Where("public = ?", iPublic).Update(iBlogsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogsViaArchived via map[string]interface{}{}
func UpdateBlogsViaArchived(iArchived int, iBlogsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Blogs)).Where("archived = ?", iArchived).Update(iBlogsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogsViaMature via map[string]interface{}{}
func UpdateBlogsViaMature(iMature int, iBlogsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Blogs)).Where("mature = ?", iMature).Update(iBlogsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogsViaSpam via map[string]interface{}{}
func UpdateBlogsViaSpam(iSpam int, iBlogsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Blogs)).Where("spam = ?", iSpam).Update(iBlogsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogsViaDeleted via map[string]interface{}{}
func UpdateBlogsViaDeleted(iDeleted int, iBlogsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Blogs)).Where("deleted = ?", iDeleted).Update(iBlogsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogsViaLangId via map[string]interface{}{}
func UpdateBlogsViaLangId(iLangId int, iBlogsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Blogs)).Where("lang_id = ?", iLangId).Update(iBlogsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteBlogs Delete Blogs
func DeleteBlogs(iBlogId int64) (int64, error) {
	row, err := Engine.Id(iBlogId).Delete(new(Blogs))
	return row, err
}

// DeleteBlogsViaBlogId Delete Blogs via BlogId
func DeleteBlogsViaBlogId(iBlogId int64) (err error) {
	var has bool
	var _Blogs = &Blogs{BlogId: iBlogId}
	if has, err = Engine.Get(_Blogs); (has == true) && (err == nil) {
		if row, err := Engine.Where("blog_id = ?", iBlogId).Delete(new(Blogs)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogsViaSiteId Delete Blogs via SiteId
func DeleteBlogsViaSiteId(iSiteId int64) (err error) {
	var has bool
	var _Blogs = &Blogs{SiteId: iSiteId}
	if has, err = Engine.Get(_Blogs); (has == true) && (err == nil) {
		if row, err := Engine.Where("site_id = ?", iSiteId).Delete(new(Blogs)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogsViaDomain Delete Blogs via Domain
func DeleteBlogsViaDomain(iDomain string) (err error) {
	var has bool
	var _Blogs = &Blogs{Domain: iDomain}
	if has, err = Engine.Get(_Blogs); (has == true) && (err == nil) {
		if row, err := Engine.Where("domain = ?", iDomain).Delete(new(Blogs)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogsViaPath Delete Blogs via Path
func DeleteBlogsViaPath(iPath string) (err error) {
	var has bool
	var _Blogs = &Blogs{Path: iPath}
	if has, err = Engine.Get(_Blogs); (has == true) && (err == nil) {
		if row, err := Engine.Where("path = ?", iPath).Delete(new(Blogs)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogsViaRegistered Delete Blogs via Registered
func DeleteBlogsViaRegistered(iRegistered time.Time) (err error) {
	var has bool
	var _Blogs = &Blogs{Registered: iRegistered}
	if has, err = Engine.Get(_Blogs); (has == true) && (err == nil) {
		if row, err := Engine.Where("registered = ?", iRegistered).Delete(new(Blogs)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogsViaLastUpdated Delete Blogs via LastUpdated
func DeleteBlogsViaLastUpdated(iLastUpdated time.Time) (err error) {
	var has bool
	var _Blogs = &Blogs{LastUpdated: iLastUpdated}
	if has, err = Engine.Get(_Blogs); (has == true) && (err == nil) {
		if row, err := Engine.Where("last_updated = ?", iLastUpdated).Delete(new(Blogs)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogsViaPublic Delete Blogs via Public
func DeleteBlogsViaPublic(iPublic int) (err error) {
	var has bool
	var _Blogs = &Blogs{Public: iPublic}
	if has, err = Engine.Get(_Blogs); (has == true) && (err == nil) {
		if row, err := Engine.Where("public = ?", iPublic).Delete(new(Blogs)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogsViaArchived Delete Blogs via Archived
func DeleteBlogsViaArchived(iArchived int) (err error) {
	var has bool
	var _Blogs = &Blogs{Archived: iArchived}
	if has, err = Engine.Get(_Blogs); (has == true) && (err == nil) {
		if row, err := Engine.Where("archived = ?", iArchived).Delete(new(Blogs)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogsViaMature Delete Blogs via Mature
func DeleteBlogsViaMature(iMature int) (err error) {
	var has bool
	var _Blogs = &Blogs{Mature: iMature}
	if has, err = Engine.Get(_Blogs); (has == true) && (err == nil) {
		if row, err := Engine.Where("mature = ?", iMature).Delete(new(Blogs)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogsViaSpam Delete Blogs via Spam
func DeleteBlogsViaSpam(iSpam int) (err error) {
	var has bool
	var _Blogs = &Blogs{Spam: iSpam}
	if has, err = Engine.Get(_Blogs); (has == true) && (err == nil) {
		if row, err := Engine.Where("spam = ?", iSpam).Delete(new(Blogs)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogsViaDeleted Delete Blogs via Deleted
func DeleteBlogsViaDeleted(iDeleted int) (err error) {
	var has bool
	var _Blogs = &Blogs{Deleted: iDeleted}
	if has, err = Engine.Get(_Blogs); (has == true) && (err == nil) {
		if row, err := Engine.Where("deleted = ?", iDeleted).Delete(new(Blogs)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogsViaLangId Delete Blogs via LangId
func DeleteBlogsViaLangId(iLangId int) (err error) {
	var has bool
	var _Blogs = &Blogs{LangId: iLangId}
	if has, err = Engine.Get(_Blogs); (has == true) && (err == nil) {
		if row, err := Engine.Where("lang_id = ?", iLangId).Delete(new(Blogs)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
