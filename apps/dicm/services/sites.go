package services

import (
	"pandax/apps/dicm/entity"
	"pandax/pkg/global"
)

type (
	DcimSiteModel interface {
		FindList(data entity.DcimSite) (*[]entity.DcimSite, error)
		FindListPage(page, pageSize int, data entity.DcimSite) (*[]entity.DcimSite, int64, error)
		FindOne(data entity.DcimSite) (*entity.DcimSite, error)
	}

	dicmSiteImpl struct {
		table string
	}
)

var DcimSiteDao DcimSiteModel = &dicmSiteImpl{
	table: `dcim_site`,
}

func (m *dicmSiteImpl) FindList(data entity.DcimSite) (*[]entity.DcimSite, error) {
	list := make([]entity.DcimSite, 0)
	db := global.Db.Table(m.table)
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.GroupId != 0 {
		db = db.Where("group_id = ?", data.GroupId)
	}
	if data.RegionId != 0 {
		db = db.Where("region_id = ?", data.RegionId)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.TenantId != 0 {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.Slug != "" {
		db = db.Where("slug like ?", "%"+data.Slug+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	err := db.Find(&list).Error
	return &list, err

}

func (m *dicmSiteImpl) FindListPage(page, pageSize int, data entity.DcimSite) (*[]entity.DcimSite, int64, error) {
	list := make([]entity.DcimSite, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.GroupId != 0 {
		db = db.Where("group_id = ?", data.GroupId)
	}
	if data.RegionId != 0 {
		db = db.Where("region_id = ?", data.RegionId)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.TenantId != 0 {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error

	return &list, total, err
}

func (m *dicmSiteImpl) FindOne(data entity.DcimSite) (*entity.DcimSite, error) {
	resData := new(entity.DcimSite)
	db := global.Db.Table(m.table)
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.GroupId != 0 {
		db = db.Where("group_id = ?", data.GroupId)
	}
	if data.RegionId != 0 {
		db = db.Where("region_id = ?", data.RegionId)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.TenantId != 0 {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	err := db.First(resData).Error
	return resData, err
}
