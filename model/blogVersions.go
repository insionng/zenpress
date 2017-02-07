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

// GetBlogVersionsCountByBlogId Get BlogVersions via BlogId
func GetBlogVersionsCountByBlogId(iBlogId int64) int64 {
	n, _ := Engine.Where("blog_id = ?", iBlogId).Count(&BlogVersions{BlogId: iBlogId})
	return n
}

// GetBlogVersionsCountByDbVersion Get BlogVersions via DbVersion
func GetBlogVersionsCountByDbVersion(iDbVersion string) int64 {
	n, _ := Engine.Where("db_version = ?", iDbVersion).Count(&BlogVersions{DbVersion: iDbVersion})
	return n
}

// GetBlogVersionsCountByLastUpdated Get BlogVersions via LastUpdated
func GetBlogVersionsCountByLastUpdated(iLastUpdated time.Time) int64 {
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

// GetBlogVersionsesByBlogId Get BlogVersionss via BlogId
func GetBlogVersionsesByBlogId(offset int, limit int, BlogId_ int64, field string) (*[]*BlogVersions, error) {
	var _BlogVersions = new([]*BlogVersions)
	err := Engine.Table("blog_versions").Where("blog_id = ?", BlogId_).Limit(limit, offset).Desc(field).Find(_BlogVersions)
	return _BlogVersions, err
}

// GetBlogVersionsesByDbVersion Get BlogVersionss via DbVersion
func GetBlogVersionsesByDbVersion(offset int, limit int, DbVersion_ string, field string) (*[]*BlogVersions, error) {
	var _BlogVersions = new([]*BlogVersions)
	err := Engine.Table("blog_versions").Where("db_version = ?", DbVersion_).Limit(limit, offset).Desc(field).Find(_BlogVersions)
	return _BlogVersions, err
}

// GetBlogVersionsesByLastUpdated Get BlogVersionss via LastUpdated
func GetBlogVersionsesByLastUpdated(offset int, limit int, LastUpdated_ time.Time, field string) (*[]*BlogVersions, error) {
	var _BlogVersions = new([]*BlogVersions)
	err := Engine.Table("blog_versions").Where("last_updated = ?", LastUpdated_).Limit(limit, offset).Desc(field).Find(_BlogVersions)
	return _BlogVersions, err
}

// HasBlogVersionsByBlogId Has BlogVersions via BlogId
func HasBlogVersionsByBlogId(iBlogId int64) bool {
	if has, err := Engine.Where("blog_id = ?", iBlogId).Get(new(BlogVersions)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogVersionsByDbVersion Has BlogVersions via DbVersion
func HasBlogVersionsByDbVersion(iDbVersion string) bool {
	if has, err := Engine.Where("db_version = ?", iDbVersion).Get(new(BlogVersions)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasBlogVersionsByLastUpdated Has BlogVersions via LastUpdated
func HasBlogVersionsByLastUpdated(iLastUpdated time.Time) bool {
	if has, err := Engine.Where("last_updated = ?", iLastUpdated).Get(new(BlogVersions)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetBlogVersionsByBlogId Get BlogVersions via BlogId
func GetBlogVersionsByBlogId(iBlogId int64) (*BlogVersions, error) {
	var _BlogVersions = &BlogVersions{BlogId: iBlogId}
	has, err := Engine.Get(_BlogVersions)
	if has {
		return _BlogVersions, err
	} else {
		return nil, err
	}
}

// GetBlogVersionsByDbVersion Get BlogVersions via DbVersion
func GetBlogVersionsByDbVersion(iDbVersion string) (*BlogVersions, error) {
	var _BlogVersions = &BlogVersions{DbVersion: iDbVersion}
	has, err := Engine.Get(_BlogVersions)
	if has {
		return _BlogVersions, err
	} else {
		return nil, err
	}
}

// GetBlogVersionsByLastUpdated Get BlogVersions via LastUpdated
func GetBlogVersionsByLastUpdated(iLastUpdated time.Time) (*BlogVersions, error) {
	var _BlogVersions = &BlogVersions{LastUpdated: iLastUpdated}
	has, err := Engine.Get(_BlogVersions)
	if has {
		return _BlogVersions, err
	} else {
		return nil, err
	}
}

// SetBlogVersionsByBlogId Set BlogVersions via BlogId
func SetBlogVersionsByBlogId(iBlogId int64, blog_versions *BlogVersions) (int64, error) {
	blog_versions.BlogId = iBlogId
	return Engine.Insert(blog_versions)
}

// SetBlogVersionsByDbVersion Set BlogVersions via DbVersion
func SetBlogVersionsByDbVersion(iDbVersion string, blog_versions *BlogVersions) (int64, error) {
	blog_versions.DbVersion = iDbVersion
	return Engine.Insert(blog_versions)
}

// SetBlogVersionsByLastUpdated Set BlogVersions via LastUpdated
func SetBlogVersionsByLastUpdated(iLastUpdated time.Time, blog_versions *BlogVersions) (int64, error) {
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

// PutBlogVersionsByBlogId Put BlogVersions via BlogId
func PutBlogVersionsByBlogId(BlogId_ int64, iBlogVersions *BlogVersions) (int64, error) {
	row, err := Engine.Update(iBlogVersions, &BlogVersions{BlogId: BlogId_})
	return row, err
}

// PutBlogVersionsByDbVersion Put BlogVersions via DbVersion
func PutBlogVersionsByDbVersion(DbVersion_ string, iBlogVersions *BlogVersions) (int64, error) {
	row, err := Engine.Update(iBlogVersions, &BlogVersions{DbVersion: DbVersion_})
	return row, err
}

// PutBlogVersionsByLastUpdated Put BlogVersions via LastUpdated
func PutBlogVersionsByLastUpdated(LastUpdated_ time.Time, iBlogVersions *BlogVersions) (int64, error) {
	row, err := Engine.Update(iBlogVersions, &BlogVersions{LastUpdated: LastUpdated_})
	return row, err
}

// UpdateBlogVersionsByBlogId via map[string]interface{}{}
func UpdateBlogVersionsByBlogId(iBlogId int64, iBlogVersionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlogVersions)).Where("blog_id = ?", iBlogId).Update(iBlogVersionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogVersionsByDbVersion via map[string]interface{}{}
func UpdateBlogVersionsByDbVersion(iDbVersion string, iBlogVersionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(BlogVersions)).Where("db_version = ?", iDbVersion).Update(iBlogVersionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateBlogVersionsByLastUpdated via map[string]interface{}{}
func UpdateBlogVersionsByLastUpdated(iLastUpdated time.Time, iBlogVersionsMap *map[string]interface{}) error {
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

// DeleteBlogVersionsByBlogId Delete BlogVersions via BlogId
func DeleteBlogVersionsByBlogId(iBlogId int64) (err error) {
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

// DeleteBlogVersionsByDbVersion Delete BlogVersions via DbVersion
func DeleteBlogVersionsByDbVersion(iDbVersion string) (err error) {
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

// DeleteBlogVersionsByLastUpdated Delete BlogVersions via LastUpdated
func DeleteBlogVersionsByLastUpdated(iLastUpdated time.Time) (err error) {
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
