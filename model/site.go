package model

type Site struct {
	Id     int64  `xorm:"BIGINT(20)"`
	Domain string `xorm:"not null default '' index(domain) VARCHAR(200)"`
	Path   string `xorm:"not null default '' index(domain) VARCHAR(100)"`
}

// GetSitesCount Sites' Count
func GetSitesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Site{})
	return total, err
}

// GetSiteCountById Get Site via Id
func GetSiteCountById(iId int64) int64 {
	n, _ := Engine.Where("id = ?", iId).Count(&Site{Id: iId})
	return n
}

// GetSiteCountByDomain Get Site via Domain
func GetSiteCountByDomain(iDomain string) int64 {
	n, _ := Engine.Where("domain = ?", iDomain).Count(&Site{Domain: iDomain})
	return n
}

// GetSiteCountByPath Get Site via Path
func GetSiteCountByPath(iPath string) int64 {
	n, _ := Engine.Where("path = ?", iPath).Count(&Site{Path: iPath})
	return n
}

// GetSitesByIdAndDomainAndPath Get Sites via IdAndDomainAndPath
func GetSitesByIdAndDomainAndPath(offset int, limit int, Id_ int64, Domain_ string, Path_ string) (*[]*Site, error) {
	var _Site = new([]*Site)
	err := Engine.Table("site").Where("id = ? and domain = ? and path = ?", Id_, Domain_, Path_).Limit(limit, offset).Find(_Site)
	return _Site, err
}

// GetSitesByIdAndDomain Get Sites via IdAndDomain
func GetSitesByIdAndDomain(offset int, limit int, Id_ int64, Domain_ string) (*[]*Site, error) {
	var _Site = new([]*Site)
	err := Engine.Table("site").Where("id = ? and domain = ?", Id_, Domain_).Limit(limit, offset).Find(_Site)
	return _Site, err
}

// GetSitesByIdAndPath Get Sites via IdAndPath
func GetSitesByIdAndPath(offset int, limit int, Id_ int64, Path_ string) (*[]*Site, error) {
	var _Site = new([]*Site)
	err := Engine.Table("site").Where("id = ? and path = ?", Id_, Path_).Limit(limit, offset).Find(_Site)
	return _Site, err
}

// GetSitesByDomainAndPath Get Sites via DomainAndPath
func GetSitesByDomainAndPath(offset int, limit int, Domain_ string, Path_ string) (*[]*Site, error) {
	var _Site = new([]*Site)
	err := Engine.Table("site").Where("domain = ? and path = ?", Domain_, Path_).Limit(limit, offset).Find(_Site)
	return _Site, err
}

// GetSites Get Sites via field
func GetSites(offset int, limit int, field string) (*[]*Site, error) {
	var _Site = new([]*Site)
	err := Engine.Table("site").Limit(limit, offset).Desc(field).Find(_Site)
	return _Site, err
}

// GetSitesById Get Sites via Id
func GetSitesById(offset int, limit int, Id_ int64, field string) (*[]*Site, error) {
	var _Site = new([]*Site)
	err := Engine.Table("site").Where("id = ?", Id_).Limit(limit, offset).Desc(field).Find(_Site)
	return _Site, err
}

// GetSitesByDomain Get Sites via Domain
func GetSitesByDomain(offset int, limit int, Domain_ string, field string) (*[]*Site, error) {
	var _Site = new([]*Site)
	err := Engine.Table("site").Where("domain = ?", Domain_).Limit(limit, offset).Desc(field).Find(_Site)
	return _Site, err
}

// GetSitesByPath Get Sites via Path
func GetSitesByPath(offset int, limit int, Path_ string, field string) (*[]*Site, error) {
	var _Site = new([]*Site)
	err := Engine.Table("site").Where("path = ?", Path_).Limit(limit, offset).Desc(field).Find(_Site)
	return _Site, err
}

// HasSiteById Has Site via Id
func HasSiteById(iId int64) bool {
	if has, err := Engine.Where("id = ?", iId).Get(new(Site)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSiteByDomain Has Site via Domain
func HasSiteByDomain(iDomain string) bool {
	if has, err := Engine.Where("domain = ?", iDomain).Get(new(Site)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSiteByPath Has Site via Path
func HasSiteByPath(iPath string) bool {
	if has, err := Engine.Where("path = ?", iPath).Get(new(Site)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetSiteById Get Site via Id
func GetSiteById(iId int64) (*Site, error) {
	var _Site = &Site{Id: iId}
	has, err := Engine.Get(_Site)
	if has {
		return _Site, err
	} else {
		return nil, err
	}
}

// GetSiteByDomain Get Site via Domain
func GetSiteByDomain(iDomain string) (*Site, error) {
	var _Site = &Site{Domain: iDomain}
	has, err := Engine.Get(_Site)
	if has {
		return _Site, err
	} else {
		return nil, err
	}
}

// GetSiteByPath Get Site via Path
func GetSiteByPath(iPath string) (*Site, error) {
	var _Site = &Site{Path: iPath}
	has, err := Engine.Get(_Site)
	if has {
		return _Site, err
	} else {
		return nil, err
	}
}

// SetSiteById Set Site via Id
func SetSiteById(iId int64, site *Site) (int64, error) {
	site.Id = iId
	return Engine.Insert(site)
}

// SetSiteByDomain Set Site via Domain
func SetSiteByDomain(iDomain string, site *Site) (int64, error) {
	site.Domain = iDomain
	return Engine.Insert(site)
}

// SetSiteByPath Set Site via Path
func SetSiteByPath(iPath string, site *Site) (int64, error) {
	site.Path = iPath
	return Engine.Insert(site)
}

// AddSite Add Site via all columns
func AddSite(iId int64, iDomain string, iPath string) error {
	_Site := &Site{Id: iId, Domain: iDomain, Path: iPath}
	if _, err := Engine.Insert(_Site); err != nil {
		return err
	} else {
		return nil
	}
}

// PostSite Post Site via iSite
func PostSite(iSite *Site) (int64, error) {
	_, err := Engine.Insert(iSite)
	return iSite.Id, err
}

// PutSite Put Site
func PutSite(iSite *Site) (int64, error) {
	_, err := Engine.Id(iSite.Id).Update(iSite)
	return iSite.Id, err
}

// PutSiteById Put Site via Id
func PutSiteById(Id_ int64, iSite *Site) (int64, error) {
	row, err := Engine.Update(iSite, &Site{Id: Id_})
	return row, err
}

// PutSiteByDomain Put Site via Domain
func PutSiteByDomain(Domain_ string, iSite *Site) (int64, error) {
	row, err := Engine.Update(iSite, &Site{Domain: Domain_})
	return row, err
}

// PutSiteByPath Put Site via Path
func PutSiteByPath(Path_ string, iSite *Site) (int64, error) {
	row, err := Engine.Update(iSite, &Site{Path: Path_})
	return row, err
}

// UpdateSiteById via map[string]interface{}{}
func UpdateSiteById(iId int64, iSiteMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Site)).Where("id = ?", iId).Update(iSiteMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSiteByDomain via map[string]interface{}{}
func UpdateSiteByDomain(iDomain string, iSiteMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Site)).Where("domain = ?", iDomain).Update(iSiteMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSiteByPath via map[string]interface{}{}
func UpdateSiteByPath(iPath string, iSiteMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Site)).Where("path = ?", iPath).Update(iSiteMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteSite Delete Site
func DeleteSite(iId int64) (int64, error) {
	row, err := Engine.Id(iId).Delete(new(Site))
	return row, err
}

// DeleteSiteById Delete Site via Id
func DeleteSiteById(iId int64) (err error) {
	var has bool
	var _Site = &Site{Id: iId}
	if has, err = Engine.Get(_Site); (has == true) && (err == nil) {
		if row, err := Engine.Where("id = ?", iId).Delete(new(Site)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSiteByDomain Delete Site via Domain
func DeleteSiteByDomain(iDomain string) (err error) {
	var has bool
	var _Site = &Site{Domain: iDomain}
	if has, err = Engine.Get(_Site); (has == true) && (err == nil) {
		if row, err := Engine.Where("domain = ?", iDomain).Delete(new(Site)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteSiteByPath Delete Site via Path
func DeleteSiteByPath(iPath string) (err error) {
	var has bool
	var _Site = &Site{Path: iPath}
	if has, err = Engine.Get(_Site); (has == true) && (err == nil) {
		if row, err := Engine.Where("path = ?", iPath).Delete(new(Site)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
