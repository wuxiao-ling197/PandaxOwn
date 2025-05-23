package services

import (
	"database/sql"
	"errors"
	"log"
	"pandax/apps/dicm/entity"
	"pandax/pkg/global"
	"time"
)

type (
	DcimSiteModel interface {
		FindList(data entity.DcimSite) (*[]entity.DcimSite, error)
		FindListPage(page, pageSize int, data entity.DcimSite) (*[]entity.DcimSite, int64, error)
		FindOne(data entity.DcimSite) (*entity.DcimSite, error)
		Insert(data entity.DcimSite) (*entity.DcimSite, error)
		Update(data entity.DcimSite) error
		Delete(ids []int64) error
		JoinGroup(sites []int64, groupId int64) error
		GetSiteStructrue() ([]*entity.SiteNode, error)
		// 站点组
		FindListGroup(data entity.DcimSitegroup) (*[]entity.DcimSitegroup, error)
		FindListGroupPage(page, pageSize int, data entity.DcimSitegroup) (*[]entity.DcimSitegroup, int64, error)
		InsertGroup(data entity.DcimSitegroup) (*entity.DcimSitegroup, error)
		UpdateGroup(data entity.DcimSitegroup) error
		DeleteGroup(ids []int64) error
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
	if data.GroupId != nil {
		db = db.Where("group_id = ?", data.GroupId)
	}
	if data.RegionId != nil {
		db = db.Where("region_id = ?", data.RegionId)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.TenantId != nil {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.Slug != "" {
		db = db.Where("slug like ?", "%"+data.Slug+"%")
	}
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
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
	db := global.Db.Model(entity.DcimSite{})
	// Table(m.table)
	// 此处填写 where参数判断
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.GroupId != nil {
		db = db.Where("group_id = ?", data.GroupId)
	}
	if data.RegionId != nil {
		db = db.Where("region_id = ?", data.RegionId)
	}
	if data.TenantId != nil {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
	}
	if data.NickNameName != "" {
		db = db.Where("_name = ?", data.NickNameName)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.Slug != "" {
		db = db.Where("slug = ?", data.Slug)
	}
	if data.ShippingAddress != "" {
		db = db.Where("shipping_address = ?", data.ShippingAddress)
	}
	if data.PhysicalAddress != "" {
		db = db.Where("physical_address = ?", data.PhysicalAddress)
	}
	if data.TimeZone != "" {
		db = db.Where("time_zone = ?", data.TimeZone)
	}
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error

	return &list, total, err
}

func (m *dicmSiteImpl) FindOne(data entity.DcimSite) (*entity.DcimSite, error) {
	resData := new(entity.DcimSite)
	db := global.Db.Table(m.table)
	// if data.Id != 0 {
	// 	db = db.Where("id = ?", data.Id)
	// }
	// if data.GroupId != nil {
	// 	db = db.Where("group_id = ?", data.GroupId)
	// }
	// if data.RegionId != nil {
	// 	db = db.Where("region_id = ?", data.RegionId)
	// }
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	// if data.CustomFieldData != "" {
	// 	db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
	// }
	// if data.NickNameName != "" {
	// 	db = db.Where("_name = ?", data.NickNameName)
	// }
	// if data.TenantId != nil {
	// 	db = db.Where("tenant_id = ?", data.TenantId)
	// }
	// if data.Status != "" {
	// 	db = db.Where("status = ?", data.Status)
	// }
	// if data.Slug != "" {
	// 	db = db.Where("slug = ?", data.Slug)
	// }
	// if data.ShippingAddress != "" {
	// 	db = db.Where("shipping_address = ?", data.ShippingAddress)
	// }
	// if data.PhysicalAddress != "" {
	// 	db = db.Where("physical_address = ?", data.PhysicalAddress)
	// }
	// if data.TimeZone != "" {
	// 	db = db.Where("time_zone = ?", data.TimeZone)
	// }
	err := db.First(resData).Error
	return resData, err
}

func (m *dicmSiteImpl) Insert(data entity.DcimSite) (*entity.DcimSite, error) {
	var count int64
	global.Db.Table(m.table).Where("name = ? ", data.Name).Count(&count)
	if count != 0 {
		return nil, errors.New("该实例已存在！")
	}
	global.Db.Table(m.table).Where("slug = ? ", data.Slug).Count(&count)
	if count != 0 {
		return nil, errors.New("该标识符已存在！")
	}
	data.Created = time.Now()
	err := global.Db.Table(m.table).Create(&data).Error
	log.Printf("错误= %+v\n", err)
	return &data, err
}

func (m *dicmSiteImpl) Update(data entity.DcimSite) error {
	update := new(entity.DcimSite)
	err := global.Db.Table(m.table).First(update, data.Id).Error
	if err != nil {
		return err
	}
	// if data.GroupId != 0 {
	// 	update.GroupId = data.GroupId
	// }
	// if data.SiteId != 0 {
	// 	update.SiteId = data.SiteId
	// }

	data.LastUpdated = time.Now()
	return global.Db.Table(m.table).Updates(&data).Error
}

func (m *dicmSiteImpl) Delete(ids []int64) error {
	data := new(entity.DcimSite)
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

// 站点加入站点组
func (m *dicmSiteImpl) JoinGroup(sites []int64, groupId int64) error {
	var i int
	for i = 0; i < len(sites); i++ {
		result := global.Db.Model(&entity.DcimSite{}).Where("id = ?", sites[i]).Update("group_id", groupId)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// 站点组层级结构  需完善修改
func (m *dicmSiteImpl) GetSiteStructrue() ([]*entity.SiteNode, error) {
	// 1. 定义一个临时结构体，仅用于从数据库读取必要的字段
	type TempNode struct {
		Id       int64
		Name     string
		ParentId sql.NullInt64
	}
	var flatNodes []TempNode

	// 从数据库查询 id, name, parent_id
	global.Db.Model(&entity.DcimSitegroup{}).
		Select("id, name, parent_id").
		Find(&flatNodes)

	// if result.Error != nil {
	// 	log.Printf("Error fetching flat organization nodes: %v\n", result.Error)
	// 	return nil, result.Error
	// }

	// 如果没有数据，直接返回空
	if len(flatNodes) == 0 {
		return []*entity.SiteNode{}, nil
	}

	// 2. 构建层级结构
	//  创建一个map来存储所有节点，以ID为键，方便快速查找父节点
	nodeMap := make(map[int64]*entity.SiteNode)

	//  初始化所有节点，并将它们存入map
	for _, fn := range flatNodes {
		nodeMap[fn.Id] = &entity.SiteNode{
			Id:       fn.Id,
			Name:     fn.Name,
			ParentId: fn.ParentId,          // 保留ParentId字段
			Children: []*entity.SiteNode{}, // 初始化Children切片
		}
	}

	//  构建树形结构：遍历所有节点，将它们连接到其父节点
	var roots []*entity.SiteNode // 用于存储所有根节点 (没有父节点的节点)
	for _, fn := range flatNodes {
		node := nodeMap[fn.Id] // 获取当前处理的节点
		if fn.ParentId.Valid { // 如果存在父ID
			parentNode, exists := nodeMap[fn.ParentId.Int64] // 尝试从map中获取父节点
			if exists {
				// 如果父节点存在，将当前节点添加到父节点的Children列表中
				parentNode.Children = append(parentNode.Children, node)
			} else {
				// 如果父ID存在，但map中找不到对应的父节点 (数据可能不一致或父节点未被查询到)
				// 这种情况下，可以根据业务逻辑处理，例如将其视为一个孤立的根节点或报错
				// 为简单起见，这里也将其视为一个根节点（尽管它有一个指向不存在的父节点的ParentId）
				log.Printf("Warning: Node ID %d has ParentId %d, but parent node not found. Treating as a root.\n", node.Id, fn.ParentId.Int64)
				roots = append(roots, node)
			}
		} else {
			// 如果没有父ID (ParentId.Valid is false)，说明这是一个根节点
			roots = append(roots, node)
		}
	}

	return roots, nil
}

/** 站点组 */
func (m *dicmSiteImpl) FindListGroup(data entity.DcimSitegroup) (*[]entity.DcimSitegroup, error) {
	list := make([]entity.DcimSitegroup, 0)
	db := global.Db.Model(entity.DcimSitegroup{})
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.ParentId != nil {
		db = db.Where("parent_id = ?", data.ParentId)
	}
	if data.Rght != 0 {
		db = db.Where("rght = ?", data.Rght)
	}
	if data.Level != 0 {
		db = db.Where("level = ?", data.Level)
	}
	if data.TreeId != 0 {
		db = db.Where("tree_id = ?", data.TreeId)
	}
	if data.Slug != "" {
		db = db.Where("slug like ?", "%"+data.Slug+"%")
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
	}
	if data.Lft != 0 {
		db = db.Where("lft = ?", data.Lft)
	}
	err := db.Find(&list).Error
	return &list, err
}

func (m *dicmSiteImpl) FindListGroupPage(page, pageSize int, data entity.DcimSitegroup) (*[]entity.DcimSitegroup, int64, error) {
	list := make([]entity.DcimSitegroup, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Model(entity.DcimSitegroup{})
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.ParentId != nil {
		db = db.Where("parent_id = ?", data.ParentId)
	}
	if data.Rght != 0 {
		db = db.Where("rght = ?", data.Rght)
	}
	if data.Level != 0 {
		db = db.Where("level = ?", data.Level)
	}
	if data.TreeId != 0 {
		db = db.Where("tree_id = ?", data.TreeId)
	}
	if data.Slug != "" {
		db = db.Where("slug like ?", "%"+data.Slug+"%")
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
	}
	if data.Description != "" {
		db = db.Where("description like ?", "%"+data.Description+"%")
	}
	if data.Lft != 0 {
		db = db.Where("lft = ?", data.Lft)
	}
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error

	return &list, total, err
}

func (m *dicmSiteImpl) InsertGroup(data entity.DcimSitegroup) (*entity.DcimSitegroup, error) {
	var count int64
	global.Db.Model(&entity.DcimSitegroup{}).Where("name = ? ", data.Name).Count(&count)
	if count != 0 {
		return nil, errors.New("该实例已存在！")
	}
	global.Db.Model(&entity.DcimSitegroup{}).Where("slug = ? ", data.Slug).Count(&count)
	if count != 0 {
		return nil, errors.New("该标识符已存在！")
	}
	data.Created = time.Now()
	err := global.Db.Model(&entity.DcimSitegroup{}).Create(&data).Error
	log.Printf("错误=%+v\n", err)
	return &data, err
}

// 如果需要将ParentId从有效值设置回0即删除上下级关系不会有效修改，只能从默认值设置为有效值且只能在有效值之间修改
func (m *dicmSiteImpl) UpdateGroup(data entity.DcimSitegroup) error {
	update := new(entity.DcimSitegroup)
	err := global.Db.Model(&update).First(update, data.Id).Error
	if err != nil {
		return err
	}
	data.LastUpdated = time.Now()
	return global.Db.Model(&update).Updates(&data).Error
}

func (m *dicmSiteImpl) DeleteGroup(ids []int64) error {
	data := new(entity.DcimSitegroup)
	var i int
	for i = 0; i < len(ids); i++ {
		global.Db.Model(&entity.DcimSitegroup{}).First(data, ids[i])
		data.Deleted = time.Now()
		data.LastUpdated = time.Now()
		result := global.Db.Model(&entity.DcimSitegroup{}).Where("id = ?", ids[i]).Updates(&data)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
