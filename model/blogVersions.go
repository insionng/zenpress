package model

import (
	"time"
)

type BlogVersions struct {
	BlogId      int64     `xorm:"not null pk default 0 BIGINT(20)"`
	DbVersion   string    `xorm:"not null default '' index VARCHAR(20)"`
	LastUpdated time.Time `xorm:"not null default '0000-00-00 00:00:00' DATETIME"`
}

// GetBlogVersionsesCount BlogVersionss' Count
func GetBlogVersionsesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&BlogVersions{})
	return total, err
}

// GetBlogVersionsCountViaBlogId Get BlogVersions via BlogId
func GetBlogVersionsCountViaBlogId(iBlogId int64) int64 {
	n, _ := Engine.Where("blog_id = ?", iBlogId).Count(&BlogVersions{BlogId: iBlogId})
	return n
}

// GetBlogVersionsCountViaDbVersion Get BlogVersions via DbVersion
func GetBlogVersionsCountViaDbVersion(iDbVersion string) int64 {
	n, _ := Engine.Where("db_version = ?", iDbVersion).Count(&BlogVersions{DbVersion: iDbVersion})
	return n
}

// GetBlogVersionsCountViaLastUpdated Get BlogVersions via LastUpdated
func GetBlogVersionsCountViaLastUpdated(iLastUpdated time.Time) int64 {
	n, _ := Engine.Where("last_updated = ?", iLastUpdated).Count(&BlogVersions{LastUpdated: iLastUpdated})
	return n
}

// GetBlogVersionsesByBlogIdAndDbVersionAndLastUpdated Get BlogVersionses via BlogIdAndDbVersionAndLastUpdated
func GetBlogVersionsesByBlogIdAndDbVersionAndLastUpdated(offset int, limit int, BlogId_ int64, DbVersion_ string, LastUpdated_ time.Time) (*[]*BlogVersions, error) {
	var _BlogVersions = new([]*BlogVersions)
	err := Engine.Table("blog_versions").Where("blog_id = ? and db_version = ? and last_updated = ?", BlogId_, DbVersion_, LastUpdated_).Limit(limit, offset).Find(_BlogVersions)
	return _BlogVersions, err
}

// GetBlogVersionsesByBlogIdAndDbVersion Get BlogVersionses via BlogIdAndDbVersion
func GetBlogVersionsesByBlogIdAndDbVersion(offset int, limit int, BlogId_ int64, DbVersion_ string) (*[]*BlogVersions, error) {
	var _BlogVersions = new([]*BlogVersions)
	err := Engine.Table("blog_versions").Where("blog_id = ? and db_version = ?", BlogId_, DbVersion_).Limit(limit, offset).Find(_BlogVersions)
	return _BlogVersions, err
}

// GetBlogVersionsesByBlogIdAndLastUpdated Get BlogVersionses via BlogIdAndLastUpdated
func GetBlogVersionsesByBlogIdAndLastUpdated(offset int, limit int, BlogId_ int64, LastUpdated_ time.Time) (*[]*BlogVersions, error) {
	var _BlogVersions = new([]*BlogVersions)
	err := Engine.Table("blog_versions").Where("blog_id = ? and last_updated = ?", BlogId_, LastUpdated_).Limit(limit, offset).Find(_BlogVersions)
	return _BlogVersions, err
}

// GetBlogVersionsesByDbVersionAndLastUpdated Get BlogVersionses via DbVersionAndLastUpdated
func GetBlogVersionsesByDbVersionAndLastUpdated(offset int, limit int, DbVersion_ string, LastUpdated_ time.Time) (*[]*BlogVersions, error) {
	var _BlogVersions = new([]*BlogVersions)
	err := Engine.Table("blog_versions").Where("db_version = ? and last_updated = ?", DbVersion_, LastUpdated_).Limit(limit, offset).Find(_BlogVersions)
	return _BlogVersions, err
}

// GetBlogVersionses Get BlogVersionses via field
func GetBlogVersionses(offset int, limit int, field string) (*[]*BlogVersions, error) {
	var _BlogVersions = new([]*BlogVersions)
	err := Engine.Table("blog_versions").Limit(limit, offset).Desc(field).Find(_BlogVersions)
	return _BlogVersions, err
}

// GetBlogVersionsesViaBlogId Get BlogVersionss via BlogId
func GetBlogVersionsesViaBlogId(offset int, limit int, BlogId_ int64, field string) (*[]*BlogVersions, error) {
	var _BlogVersions = new([]*BlogVersions)
	err := Engine.Table("blog_versions").Where("blog_id = ?", BlogId_).Limit(limit, offset).Desc(field).Find(_BlogVersions)
	return _BlogVersions, err
}

// GetBlogVersionsesViaDbVersion Get BlogVersionss via DbVersion
func GetBlogVersionsesViaDbVersion(offset int, limit int, DbVersion_ string, field string) (*[]*BlogVersions, error) {
	var _BlogVersions = new([]*BlogVersions)
	err := Engine.Table("blog_versions").Where("db_version = ?", DbVersion_).Limit(limit, offset).Desc(field).Find(_BlogVersions)
	return _BlogVersions, err
}

// GetBlogVersionsesViaLastUpdated Get BlogVersionss via LastUpdated
func GetBlogVersionsesViaLastUpdated(offset int, limit int, LastUpdated_ time.Time, field string) (*[]*BlogVersions, error) {
	var _BlogVersions = new([]*BlogVersions)
	err := Engine.Table("blog_versions").Where("last_updated = ?", LastUpdated_).Limit(limit, offset).Desc(field).Find(_BlogVersions)
	return _BlogVersions, err
}

// HasBlogVersionsViaBlogId Has BlogVersions via BlogId
func HasBlogVersionsViaBlogId(iBlogId int64) bool {
	if has, err := Engine.Where("blog_id = ?", iBlogId).Get(new(BlogVersions)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogVersionsViaDbVersion Has BlogVersions via DbVersion
func HasBlogVersionsViaDbVersion(iDbVersion string) bool {
	if has, err := Engine.Where("db_version = ?", iDbVersion).Get(new(BlogVersions)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogVersionsViaLastUpdated Has BlogVersions via LastUpdated
func HasBlogVersionsViaLastUpdated(iLastUpdated time.Time) bool {
	if has, err := Engine.Where("last_updated = ?", iLastUpdated).Get(new(BlogVersions)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetBlogVersionsViaBlogId Get BlogVersions via BlogId
func GetBlogVersionsViaBlogId(iBlogId int64) (*BlogVersions, error) {
	var _BlogVersions = &BlogVersions{BlogId: iBlogId}
	has, err := Engine.Get(_BlogVersions)
	if has {
		return _BlogVersions, err
	} else {
		return nil, err
	}
}

// GetBlogVersionsViaDbVersion Get BlogVersions via DbVersion
func GetBlogVersionsViaDbVersion(iDbVersion string) (*BlogVersions, error) {
	var _BlogVersions = &BlogVersions{DbVersion: iDbVersion}
	has, err := Engine.Get(_BlogVersions)
	if has {
		return _BlogVersions, err
	} else {
		return nil, err
	}
}

// GetBlogVersionsViaLastUpdated Get BlogVersions via LastUpdated
func GetBlogVersionsViaLastUpdated(iLastUpdated time.Time) (*BlogVersions, error) {
	var _BlogVersions = &BlogVersions{LastUpdated: iLastUpdated}
	has, err := Engine.Get(_BlogVersions)
	if has {
		return _BlogVersions, err
	} else {
		return nil, err
	}
}

// SetBlogVersionsViaBlogId Set BlogVersions via BlogId
func SetBlogVersionsViaBlogId(iBlogId int64, blog_versions *BlogVersions) (int64, error) {
	blog_versions.BlogId = iBlogId
	return Engine.Insert(blog_versions)
}

// SetBlogVersionsViaDbVersion Set BlogVersions via DbVersion
func SetBlogVersionsViaDbVersion(iDbVersion string, blog_versions *BlogVersions) (int64, error) {
	blog_versions.DbVersion = iDbVersion
	return Engine.Insert(blog_versions)
}

// SetBlogVersionsViaLastUpdated Set BlogVersions via LastUpdated
func SetBlogVersionsViaLastUpdated(iLastUpdated time.Time, blog_versions *BlogVersions) (int64, error) {
	blog_versions.LastUpdated = iLastUpdated
	return Engine.Insert(blog_versions)
}

// AddBlogVersions Add BlogVersions via all columns
func AddBlogVersions(iBlogId int64, iDbVersion string, iLastUpdated time.Time) error {
	_BlogVersions := &BlogVersions{BlogId: iBlogId, DbVersion: iDbVersion, LastUpdated: iLastUpdated}
	if _, err := Engine.Insert(_BlogVersions); err != nil {
		return err
	} else {
		return nil
	}
}

// PostBlogVersions Post BlogVersions via iBlogVersions
func PostBlogVersions(iBlogVersions *BlogVersions) (int64, error) {
	_, err := Engine.Insert(iBlogVersions)
	return iBlogVersions.BlogId, err
}

// PutBlogVersions Put BlogVersions
func PutBlogVersions(iBlogVersions *BlogVersions) (int64, error) {
	_, err := Engine.Id(iBlogVersions.BlogId).Update(iBlogVersions)
	return iBlogVersions.BlogId, err
}

// PutBlogVersionsViaBlogId Put BlogVersions via BlogId
func PutBlogVersionsViaBlogId(BlogId_ int64, iBlogVersions *BlogVersions) (int64, error) {
	row, err := Engine.Update(iBlogVersions, &BlogVersions{BlogId: BlogId_})
	return row, err
}

// PutBlogVersionsViaDbVersion Put BlogVersions via DbVersion
func PutBlogVersionsViaDbVersion(DbVersion_ string, iBlogVersions *BlogVersions) (int64, error) {
	row, err := Engine.Update(iBlogVersions, &BlogVersions{DbVersion: DbVersion_})
	return row, err
}

// PutBlogVersionsViaLastUpdated Put BlogVersions via LastUpdated
func PutBlogVersionsViaLastUpdated(LastUpdated_ time.Time, iBlogVersions *BlogVersions) (int64, error) {
	row, err := Engine.Update(iBlogVersions, &BlogVersions{LastUpdated: LastUpdated_})
	return row, err
}

// UpdateBlogVersionsViaBlogId via map[string]interface{}{}
func UpdateBlogVersionsViaBlogId(iBlogId int64, iBlogVersionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlogVersions)).Where("blog_id = ?", iBlogId).Update(iBlogVersionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogVersionsViaDbVersion via map[string]interface{}{}
func UpdateBlogVersionsViaDbVersion(iDbVersion string, iBlogVersionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlogVersions)).Where("db_version = ?", iDbVersion).Update(iBlogVersionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogVersionsViaLastUpdated via map[string]interface{}{}
func UpdateBlogVersionsViaLastUpdated(iLastUpdated time.Time, iBlogVersionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlogVersions)).Where("last_updated = ?", iLastUpdated).Update(iBlogVersionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteBlogVersions Delete BlogVersions
func DeleteBlogVersions(iBlogId int64) (int64, error) {
	row, err := Engine.Id(iBlogId).Delete(new(BlogVersions))
	return row, err
}

// DeleteBlogVersionsViaBlogId Delete BlogVersions via BlogId
func DeleteBlogVersionsViaBlogId(iBlogId int64) (err error) {
	var has bool
	var _BlogVersions = &BlogVersions{BlogId: iBlogId}
	if has, err = Engine.Get(_BlogVersions); (has == true) && (err == nil) {
		if row, err := Engine.Where("blog_id = ?", iBlogId).Delete(new(BlogVersions)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogVersionsViaDbVersion Delete BlogVersions via DbVersion
func DeleteBlogVersionsViaDbVersion(iDbVersion string) (err error) {
	var has bool
	var _BlogVersions = &BlogVersions{DbVersion: iDbVersion}
	if has, err = Engine.Get(_BlogVersions); (has == true) && (err == nil) {
		if row, err := Engine.Where("db_version = ?", iDbVersion).Delete(new(BlogVersions)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteBlogVersionsViaLastUpdated Delete BlogVersions via LastUpdated
func DeleteBlogVersionsViaLastUpdated(iLastUpdated time.Time) (err error) {
	var has bool
	var _BlogVersions = &BlogVersions{LastUpdated: iLastUpdated}
	if has, err = Engine.Get(_BlogVersions); (has == true) && (err == nil) {
		if row, err := Engine.Where("last_updated = ?", iLastUpdated).Delete(new(BlogVersions)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
