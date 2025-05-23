package services

import (
	"errors"
	"fmt"
	"log"
	"pandax/apps/dicm/entity"
	"pandax/pkg/global"
	"time"
)

type (
	DcimRackModel interface {
		/** 机柜 **/
		FindList(data entity.DcimRack) (*[]entity.DcimRack, error)
		FindListPage(page, pageSize int, data entity.DcimRack) (*[]entity.DcimRack, int64, error)
		FindOne(data entity.DcimRack) (*entity.DcimRack, error)
		Insert(data entity.DcimRack) (*entity.DcimRack, error)
		Update(data entity.DcimRack) error
		Delete(ids []int64) error
		/** 机柜预留 **/
		FindReserveListPage(page, pageSize int, data entity.DcimRackreservation) (*[]entity.DcimRackreservation, int64, error)
		InsertReserve(data entity.DcimRackreservation) (*entity.DcimRackreservation, error)
		UpdateReserve(data entity.DcimRackreservation) error
		DeleteReserve(ids []int64) error
		/** 机柜角色 **/
		FindRoleListPage(page, pageSize int, data entity.DcimRackrole) (*[]entity.DcimRackrole, int64, error)
		InsertRackRole(data entity.DcimRackrole) (*entity.DcimRackrole, error)
		UpdateRackRole(data entity.DcimRackrole) error
		DeleteRackRole(ids []int64) error
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
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
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
	if data.LocationId != 0 {
		db = db.Where("location_id = ?", data.LocationId)
	}
	if data.TenantId != 0 {
		db = db.Where("tenant_id = ?", data.TenantId)
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
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
	}
	if data.Type != "" {
		db = db.Where("type like ?", "%"+data.Type+"%")
	}
	if data.Status != "" {
		db = db.Where("status like ?", "%"+data.Status+"%")
	}
	if data.Serial != 0 {
		db = db.Where("serial = ?", data.Serial)
	}
	if data.Weight != 0 {
		db = db.Where("weight = ?", data.Weight)
	}
	if data.UHeight != 0 {
		db = db.Where("u_height = ?", data.UHeight)
	}
	if data.StartingUnit != 0 {
		db = db.Where("starting_unit = ?", data.StartingUnit)
	}
	if data.MountingDepth != 0 {
		db = db.Where("mounting_depth = ?", data.MountingDepth)
	}
	if data.OuterDepth != 0 {
		db = db.Where("outer_depth = ?", data.OuterDepth)
	}
	if data.OuterWidth != 0 {
		db = db.Where("outer_width = ?", data.OuterWidth)
	}
	if data.AssetTag != "" {
		db = db.Where("asset_tag like ?", "%"+data.AssetTag+"%")
	}
	if data.Weight != 0 {
		db = db.Where("weight = ?", data.Weight)
	}
	if data.FacilityId != "" {
		db = db.Where("facility_id like ?", "%"+data.FacilityId+"%")
	}
	if data.Description != "" {
		db = db.Where("description like ?", "%"+data.Description+"%")
	}
	if data.Comments != "" {
		db = db.Where("comments like ?", "%"+data.Comments+"%")
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
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
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
	// fmt.Println(data)
	global.Db.Debug().Table(m.table).Where("name = ? ", data.Name).Count(&count)
	if count != 0 {
		return nil, errors.New("该实例已存在！")
	}
	data.Created = time.Now()
	err := global.Db.Table(m.table).Create(&data).Error
	log.Printf("错误：%+v\n", err)
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
		update.RoleId = data.RoleId
	}

	data.LastUpdated = time.Now()
	return global.Db.Table(m.table).Updates(&data).Error
}

func (m *dicmRackImpl) Delete(ids []int64) error {
	data := new(entity.DcimRack)
	var i int
	for i = 0; i < len(ids); i++ {
		global.Db.Table(m.table).First(data, ids[i])
		data.Deleted = time.Now()
		data.LastUpdated = time.Now()
		data.Status = "abandon"
		result := global.Db.Table(m.table).Where("id = ?", ids[i]).Updates(&data)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// 机柜预留
func (m *dicmRackImpl) FindReserveListPage(page, pageSize int, data entity.DcimRackreservation) (*[]entity.DcimRackreservation, int64, error) {
	list := make([]entity.DcimRackreservation, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Model(entity.DcimRackreservation{})
	// .Table("dcim_rackreservation")
	// 此处填写 where参数判断
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.RackId != nil {
		db = db.Where("rack_id = ?", data.RackId)
	}
	if data.TenantId != nil {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.UserId != nil {
		db = db.Where("user_id = ?", data.UserId)
	}
	if data.Units != 0 {
		db = db.Where("units = ?", data.Units)
	}
	if data.Description != "" {
		db = db.Where("description like ?", "%"+data.Description+"%")
	}
	if data.Comments != "" {
		db = db.Where("comments like ?", "%"+data.Comments+"%")
	}
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
	}
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error

	return &list, total, err
}

func (m *dicmRackImpl) InsertReserve(data entity.DcimRackreservation) (*entity.DcimRackreservation, error) {
	// var count int64
	// fmt.Println(data)
	// global.Db.Model(entity.DcimRackreservation{}).Where("name = ? ", data.Name).Count(&count)
	// if count != 0 {
	// 	return nil, errors.New("该实例已存在！")
	// }
	data.Created = time.Now()
	err := global.Db.Model(entity.DcimRackreservation{}).Create(&data).Error
	log.Printf("错误：%+v\n", err)
	return &data, err
}

func (m *dicmRackImpl) UpdateReserve(data entity.DcimRackreservation) error {
	update := new(entity.DcimRackreservation)
	err := global.Db.Model(entity.DcimRackreservation{}).First(update, data.Id).Error
	if err != nil {
		return err
	}
	// if data.RackId != nil {
	// 	update.RackId = data.RackId
	// }
	data.LastUpdated = time.Now()
	return global.Db.Model(entity.DcimRackreservation{}).Updates(&data).Error
}

func (m *dicmRackImpl) DeleteReserve(ids []int64) error {
	data := new(entity.DcimRackreservation)
	var i int
	for i = 0; i < len(ids); i++ {
		global.Db.Model(entity.DcimRackreservation{}).First(data, ids[i])
		data.Deleted = time.Now()
		data.LastUpdated = time.Now()
		result := global.Db.Model(entity.DcimRackreservation{}).Where("id = ?", ids[i]).Updates(&data)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// 机柜角色
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
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
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

func (m *dicmRackImpl) InsertRackRole(data entity.DcimRackrole) (*entity.DcimRackrole, error) {
	var count int64
	fmt.Println(data)
	global.Db.Debug().Model(entity.DcimRackrole{}).Where("name = ? ", data.Name).Count(&count)
	if count != 0 {
		return nil, errors.New("该实例已存在！")
	}
	data.Created = time.Now()
	err := global.Db.Model(entity.DcimRackrole{}).Create(&data).Error
	fmt.Printf("错误：%+v\n", err)
	// 错误：pq: 无效的类型 timestamp with time zone 输入语法: ""
	return &data, err
}

func (m *dicmRackImpl) UpdateRackRole(data entity.DcimRackrole) error {
	update := new(entity.DcimRackrole)
	err := global.Db.Model(entity.DcimRackrole{}).First(update, data.Id).Error
	if err != nil {
		return err
	}
	// if data.RackId != nil {
	// 	update.RackId = data.RackId
	// }
	data.LastUpdated = time.Now()
	return global.Db.Model(entity.DcimRackrole{}).Updates(&data).Error
}
func (m *dicmRackImpl) DeleteRackRole(ids []int64) error {
	data := new(entity.DcimRackrole)
	var i int
	for i = 0; i < len(ids); i++ {
		global.Db.Model(entity.DcimRackrole{}).First(data, ids[i])
		data.Deleted = time.Now()
		data.LastUpdated = time.Now()
		result := global.Db.Model(entity.DcimRackrole{}).Where("id = ?", ids[i]).Updates(&data)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
