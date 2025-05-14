package services

import (
	"errors"
	"fmt"
	"pandax/apps/dicm/entity"
	"pandax/pkg/global"
	"time"
)

type (
	DcimRackModel interface {
		FindList(data entity.DcimRack) (*[]entity.DcimRack, error)
		FindListPage(page, pageSize int, data entity.DcimRack) (*[]entity.DcimRack, int64, error)
		FindReserveListPage(page, pageSize int, data entity.DcimRackreservation) (*[]entity.DcimRackreservation, int64, error)
		FindRoleListPage(page, pageSize int, data entity.DcimRackrole) (*[]entity.DcimRackrole, int64, error)
		FindOne(data entity.DcimRack) (*entity.DcimRack, error)
		Insert(data entity.DcimRack) (*entity.DcimRack, error)
		Update(data entity.DcimRack) error
	}

	dicmRackImpl struct {
		table string
	}
)

var DcimRackDao DcimRackModel = &dicmRackImpl{
	table: `dcim_rack`,
}

func (m *dicmRackImpl) FindList(data entity.DcimRack) (*[]entity.DcimRack, error) {
	list := make([]entity.DcimRack, 0)
	db := global.Db.Table(m.table)
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.RoleId != 0 {
		db = db.Where("role_id = ?", data.RoleId)
	}
	if data.SiteId != 0 {
		db = db.Where("site_id = ?", data.SiteId)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Type != "" {
		db = db.Where("type = ?", data.Type)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	err := db.Find(&list).Error
	return &list, err

}

// 机柜
func (m *dicmRackImpl) FindListPage(page, pageSize int, data entity.DcimRack) (*[]entity.DcimRack, int64, error) {
	list := make([]entity.DcimRack, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table("dcim_rack")
	// 此处填写 where参数判断
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.RoleId != 0 {
		db = db.Where("role_id = ?", data.RoleId)
	}
	if data.SiteId != 0 {
		db = db.Where("site_id = ?", data.SiteId)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Type != "" {
		db = db.Where("type = ?", data.Type)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error

	return &list, total, err
}

// 机柜预留
func (m *dicmRackImpl) FindReserveListPage(page, pageSize int, data entity.DcimRackreservation) (*[]entity.DcimRackreservation, int64, error) {
	list := make([]entity.DcimRackreservation, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table("dcim_rackreservation")
	// 此处填写 where参数判断
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.RackId != "" {
		db = db.Where("rack_id = ?", data.RackId)
	}
	if data.TenantId != "" {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.UserId != "" {
		db = db.Where("user_id = ?", data.UserId)
	}
	if data.Units != "" {
		db = db.Where("units = ?", data.Units)
	}
	if data.Description != "" {
		db = db.Where("description like ?", "%"+data.Description+"%")
	}
	if data.Comments != "" {
		db = db.Where("comments like ?", "%"+data.Comments+"%")
	}
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data like ?", "%"+data.CustomFieldData+"%")
	}
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error

	return &list, total, err
}

// 机柜类型、用户
func (m *dicmRackImpl) FindRoleListPage(page, pageSize int, data entity.DcimRackrole) (*[]entity.DcimRackrole, int64, error) {
	list := make([]entity.DcimRackrole, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Model(&entity.DcimRackrole{})
	// 此处填写 where参数判断
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Slug != "" {
		db = db.Where("slug like ?", "%"+data.Slug+"%")
	}
	if data.Color != "" {
		db = db.Where("color = ?", data.Color)
	}
	if data.Description != "" {
		db = db.Where("description like ?", "%"+data.Description+"%")
	}
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error

	return &list, total, err
}

func (m *dicmRackImpl) FindOne(data entity.DcimRack) (*entity.DcimRack, error) {
	resData := new(entity.DcimRack)
	db := global.Db.Table(m.table)
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.RoleId != 0 {
		db = db.Where("role_id = ?", data.RoleId)
	}
	if data.SiteId != 0 {
		db = db.Where("site_id = ?", data.SiteId)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.NickName != "" {
		db = db.Where("_name like ?", "%"+data.NickName+"%")
	}
	if data.Type != "" {
		db = db.Where("type = ?", data.Type)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	err := db.First(resData).Error
	return resData, err
}

func (m *dicmRackImpl) Insert(data entity.DcimRack) (*entity.DcimRack, error) {
	var count int64
	fmt.Println(data)
	global.Db.Debug().Table(m.table).Where("name = ? ", data.Name).Count(&count)
	if count != 0 {
		return nil, errors.New("该实例已存在！")
	}
	data.Created = time.Now()
	err := global.Db.Table(m.table).Create(&data).Error
	fmt.Printf("错误：%+v\n", err)
	// 错误：pq: 无效的类型 timestamp with time zone 输入语法: ""
	return &data, err
}

func (m *dicmRackImpl) Update(data entity.DcimRack) error {
	update := new(entity.DcimRack)
	err := global.Db.Table(m.table).First(update, data.Id).Error
	if err != nil {
		return err
	}

	if data.RoleId == 0 {
		data.RoleId = update.RoleId
	}

	data.LastUpdated = time.Now()
	return global.Db.Table(m.table).Updates(&data).Error
}
