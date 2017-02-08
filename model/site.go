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

// GetSiteCountViaId Get Site via Id
func GetSiteCountViaId(iId int64) int64 {
	n, _ := Engine.Where("id = ?", iId).Count(&Site{Id: iId})
	return n
}

// GetSiteCountViaDomain Get Site via Domain
func GetSiteCountViaDomain(iDomain string) int64 {
	n, _ := Engine.Where("domain = ?", iDomain).Count(&Site{Domain: iDomain})
	return n
}

// GetSiteCountViaPath Get Site via Path
func GetSiteCountViaPath(iPath string) int64 {
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

// GetSitesViaId Get Sites via Id
func GetSitesViaId(offset int, limit int, Id_ int64, field string) (*[]*Site, error) {
	var _Site = new([]*Site)
	err := Engine.Table("site").Where("id = ?", Id_).Limit(limit, offset).Desc(field).Find(_Site)
	return _Site, err
}

// GetSitesViaDomain Get Sites via Domain
func GetSitesViaDomain(offset int, limit int, Domain_ string, field string) (*[]*Site, error) {
	var _Site = new([]*Site)
	err := Engine.Table("site").Where("domain = ?", Domain_).Limit(limit, offset).Desc(field).Find(_Site)
	return _Site, err
}

// GetSitesViaPath Get Sites via Path
func GetSitesViaPath(offset int, limit int, Path_ string, field string) (*[]*Site, error) {
	var _Site = new([]*Site)
	err := Engine.Table("site").Where("path = ?", Path_).Limit(limit, offset).Desc(field).Find(_Site)
	return _Site, err
}

// HasSiteViaId Has Site via Id
func HasSiteViaId(iId int64) bool {
	if has, err := Engine.Where("id = ?", iId).Get(new(Site)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSiteViaDomain Has Site via Domain
func HasSiteViaDomain(iDomain string) bool {
	if has, err := Engine.Where("domain = ?", iDomain).Get(new(Site)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasSiteViaPath Has Site via Path
func HasSiteViaPath(iPath string) bool {
	if has, err := Engine.Where("path = ?", iPath).Get(new(Site)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetSiteViaId Get Site via Id
func GetSiteViaId(iId int64) (*Site, error) {
	var _Site = &Site{Id: iId}
	has, err := Engine.Get(_Site)
	if has {
		return _Site, err
	} else {
		return nil, err
	}
}

// GetSiteViaDomain Get Site via Domain
func GetSiteViaDomain(iDomain string) (*Site, error) {
	var _Site = &Site{Domain: iDomain}
	has, err := Engine.Get(_Site)
	if has {
		return _Site, err
	} else {
		return nil, err
	}
}

// GetSiteViaPath Get Site via Path
func GetSiteViaPath(iPath string) (*Site, error) {
	var _Site = &Site{Path: iPath}
	has, err := Engine.Get(_Site)
	if has {
		return _Site, err
	} else {
		return nil, err
	}
}

// SetSiteViaId Set Site via Id
func SetSiteViaId(iId int64, site *Site) (int64, error) {
	site.Id = iId
	return Engine.Insert(site)
}

// SetSiteViaDomain Set Site via Domain
func SetSiteViaDomain(iDomain string, site *Site) (int64, error) {
	site.Domain = iDomain
	return Engine.Insert(site)
}

// SetSiteViaPath Set Site via Path
func SetSiteViaPath(iPath string, site *Site) (int64, error) {
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

// PutSiteViaId Put Site via Id
func PutSiteViaId(Id_ int64, iSite *Site) (int64, error) {
	row, err := Engine.Update(iSite, &Site{Id: Id_})
	return row, err
}

// PutSiteViaDomain Put Site via Domain
func PutSiteViaDomain(Domain_ string, iSite *Site) (int64, error) {
	row, err := Engine.Update(iSite, &Site{Domain: Domain_})
	return row, err
}

// PutSiteViaPath Put Site via Path
func PutSiteViaPath(Path_ string, iSite *Site) (int64, error) {
	row, err := Engine.Update(iSite, &Site{Path: Path_})
	return row, err
}

// UpdateSiteViaId via map[string]interface{}{}
func UpdateSiteViaId(iId int64, iSiteMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Site)).Where("id = ?", iId).Update(iSiteMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSiteViaDomain via map[string]interface{}{}
func UpdateSiteViaDomain(iDomain string, iSiteMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Site)).Where("domain = ?", iDomain).Update(iSiteMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateSiteViaPath via map[string]interface{}{}
func UpdateSiteViaPath(iPath string, iSiteMap *map[string]interface{}) error {
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

// DeleteSiteViaId Delete Site via Id
func DeleteSiteViaId(iId int64) (err error) {
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

// DeleteSiteViaDomain Delete Site via Domain
func DeleteSiteViaDomain(iDomain string) (err error) {
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

// DeleteSiteViaPath Delete Site via Path
func DeleteSiteViaPath(iPath string) (err error) {
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
