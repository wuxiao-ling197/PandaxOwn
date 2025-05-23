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
	DcimLocationModel interface {
		// location
		FindList(data entity.DcimLocation) (*[]entity.DcimLocation, error)
		FindListPage(page, pageSize int, data entity.DcimLocation) (*[]entity.DcimLocation, int64, error)
		FindOne(data entity.DcimLocation) (*entity.DcimLocation, error)
		Insert(data entity.DcimLocation) (*entity.DcimLocation, error)
		Update(data entity.DcimLocation) error
		Delete(ids []int64) error
		GetLocationStructrue() ([]*entity.LocationNode, error)
		// region
		FindRegionList(data entity.DcimRegion) (*[]entity.DcimRegion, error)
		FindRegionListPage(page, pageSize int, data entity.DcimRegion) (*[]entity.DcimRegion, int64, error)
		FindRegionOne(data entity.DcimRegion) (*entity.DcimRegion, error)
		InsertRegion(data entity.DcimRegion) (*entity.DcimRegion, error)
		UpdateRegion(data entity.DcimRegion) error
		DeleteRegion(ids []int64) error
		GetRegionStructrue() ([]*entity.DcimRegion, error)
	}

	dicmLocationImpl struct {
		table string
	}
)

var DcimLocationDao DcimLocationModel = &dicmLocationImpl{
	table: `dcim_location`,
}

/** Location **/
func (m *dicmLocationImpl) FindList(data entity.DcimLocation) (*[]entity.DcimLocation, error) {
	list := make([]entity.DcimLocation, 0)
	db := global.Db.Table(m.table)
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.Slug != "" {
		db = db.Where("slug = ?", data.Slug)
	}
	if data.Level != 0 {
		db = db.Where("level = ?", data.Level)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
	}
	if data.ParentId != nil {
		db = db.Where("parent_id = ?", data.ParentId)
	}
	if data.TenantId != 0 {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.TreeId != 0 {
		db = db.Where("tree_id = ?", data.TreeId)
	}
	if data.SiteId != 0 {
		db = db.Where("site_id = ?", data.SiteId)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	err := db.Find(&list).Error
	return &list, err
}

func (m *dicmLocationImpl) FindListPage(page, pageSize int, data entity.DcimLocation) (*[]entity.DcimLocation, int64, error) {
	list := make([]entity.DcimLocation, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.Slug != "" {
		db = db.Where("slug = ?", data.Slug)
	}
	if data.Level != 0 {
		db = db.Where("level = ?", data.Level)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
	}
	if data.SiteId != 0 {
		db = db.Where("site_id = ?", data.SiteId)
	}
	if data.ParentId != nil {
		db = db.Where("parent_id = ?", data.ParentId)
	}
	if data.TenantId != 0 {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.TreeId != 0 {
		db = db.Where("tree_id = ?", data.TreeId)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error

	return &list, total, err
}

func (m *dicmLocationImpl) FindOne(data entity.DcimLocation) (*entity.DcimLocation, error) {
	resData := new(entity.DcimLocation)
	db := global.Db.Table(m.table)
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.Slug != "" {
		db = db.Where("slug = ?", data.Slug)
	}
	if data.Level != 0 {
		db = db.Where("level = ?", data.Level)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
	}
	if data.SiteId != 0 {
		db = db.Where("site_id = ?", data.SiteId)
	}
	if data.ParentId != nil {
		db = db.Where("parent_id = ?", data.ParentId)
	}
	if data.TenantId != 0 {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.TreeId != 0 {
		db = db.Where("tree_id = ?", data.TreeId)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	err := db.First(resData).Error
	return resData, err
}

func (m *dicmLocationImpl) Insert(data entity.DcimLocation) (*entity.DcimLocation, error) {
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

func (m *dicmLocationImpl) Update(data entity.DcimLocation) error {
	update := new(entity.DcimLocation)
	err := global.Db.Table(m.table).First(update, data.Id).Error
	if err != nil {
		return err
	}
	if data.ParentId != nil {
		update.ParentId = data.ParentId
	}
	if data.SiteId != 0 {
		update.SiteId = data.SiteId
	}

	data.LastUpdated = time.Now()
	return global.Db.Table(m.table).Updates(&data).Error
}

func (m *dicmLocationImpl) Delete(ids []int64) error {
	data := new(entity.DcimLocation)
	var i int
	for i = 0; i < len(ids); i++ {
		global.Db.Table(m.table).First(data, ids[i])
		data.Deleted = time.Now()
		data.LastUpdated = time.Now()
		result := global.Db.Table(m.table).Where("id = ?", ids[i]).Updates(&data)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// 层级结构
func (m *dicmLocationImpl) GetLocationStructrue() ([]*entity.LocationNode, error) {
	// 1. 定义一个临时结构体，仅用于从数据库读取必要的字段
	type TempNode struct {
		Id       int64
		Name     string
		ParentId sql.NullInt64
	}
	var flatNodes []TempNode

	// 从数据库查询 id, name, parent_id
	global.Db.Model(&entity.DcimLocation{}).
		Select("id, name, parent_id").
		Find(&flatNodes)

	// if result.Error != nil {
	// 	log.Printf("Error fetching flat organization nodes: %v\n", result.Error)
	// 	return nil, result.Error
	// }

	// 如果没有数据，直接返回空
	if len(flatNodes) == 0 {
		return []*entity.LocationNode{}, nil
	}

	// 2. 构建层级结构
	//  创建一个map来存储所有节点，以ID为键，方便快速查找父节点
	nodeMap := make(map[int64]*entity.LocationNode)

	//  初始化所有节点，并将它们存入map
	for _, fn := range flatNodes {
		nodeMap[fn.Id] = &entity.LocationNode{
			Id:       fn.Id,
			Name:     fn.Name,
			ParentId: fn.ParentId,              // 保留ParentId字段
			Children: []*entity.LocationNode{}, // 初始化Children切片
		}
	}

	//  构建树形结构：遍历所有节点，将它们连接到其父节点
	var roots []*entity.LocationNode // 用于存储所有根节点 (没有父节点的节点)
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

/** Region **/
func (m *dicmLocationImpl) FindRegionList(data entity.DcimRegion) (*[]entity.DcimRegion, error) {
	list := make([]entity.DcimRegion, 0)
	db := global.Db.Model(entity.DcimRegion{})
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Slug != "" {
		db = db.Where("slug = ?", data.Slug)
	}
	if data.Level != 0 {
		db = db.Where("level = ?", data.Level)
	}
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
	}
	if data.ParentId != nil {
		db = db.Where("parent_id = ?", data.ParentId)
	}
	if data.TreeId != 0 {
		db = db.Where("tree_id = ?", data.TreeId)
	}
	if data.Description != "" {
		db = db.Where("description like ?", "%"+data.Description+"%")
	}
	err := db.Find(&list).Error
	return &list, err
}

func (m *dicmLocationImpl) FindRegionListPage(page, pageSize int, data entity.DcimRegion) (*[]entity.DcimRegion, int64, error) {
	list := make([]entity.DcimRegion, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Model(entity.DcimRegion{})
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Slug != "" {
		db = db.Where("slug = ?", data.Slug)
	}
	if data.Level != 0 {
		db = db.Where("level = ?", data.Level)
	}
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
	}
	if data.ParentId != nil {
		db = db.Where("parent_id = ?", data.ParentId)
	}
	if data.TreeId != 0 {
		db = db.Where("tree_id = ?", data.TreeId)
	}
	if data.Description != "" {
		db = db.Where("description like ?", "%"+data.Description+"%")
	}
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error

	return &list, total, err
}

func (m *dicmLocationImpl) FindRegionOne(data entity.DcimRegion) (*entity.DcimRegion, error) {
	resData := new(entity.DcimRegion)
	db := global.Db.Model(entity.DcimRegion{})
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Slug != "" {
		db = db.Where("slug = ?", data.Slug)
	}
	if data.Level != 0 {
		db = db.Where("level = ?", data.Level)
	}
	if data.CustomFieldData != "" {
		db = db.Where("custom_field_data::TEXT like ?", "%"+data.CustomFieldData+"%")
	}
	if data.ParentId != nil {
		db = db.Where("parent_id = ?", data.ParentId)
	}
	if data.TreeId != 0 {
		db = db.Where("tree_id = ?", data.TreeId)
	}
	if data.Description != "" {
		db = db.Where("description like ?", "%"+data.Description+"%")
	}
	err := db.First(resData).Error
	return resData, err
}

func (m *dicmLocationImpl) InsertRegion(data entity.DcimRegion) (*entity.DcimRegion, error) {
	var count int64
	global.Db.Model(entity.DcimRegion{}).Where("name = ? ", data.Name).Count(&count)
	if count != 0 {
		return nil, errors.New("该实例已存在！")
	}
	global.Db.Model(entity.DcimRegion{}).Where("slug = ? ", data.Slug).Count(&count)
	if count != 0 {
		return nil, errors.New("该标识符已存在！")
	}
	data.Created = time.Now()
	err := global.Db.Model(entity.DcimRegion{}).Create(&data).Error
	log.Printf("错误= %+v\n", err)
	return &data, err
}

func (m *dicmLocationImpl) UpdateRegion(data entity.DcimRegion) error {
	update := new(entity.DcimRegion)
	err := global.Db.Model(entity.DcimRegion{}).First(update, data.Id).Error
	if err != nil {
		return err
	}
	if data.ParentId != nil {
		update.ParentId = data.ParentId
	}

	data.LastUpdated = time.Now()
	return global.Db.Model(entity.DcimRegion{}).Updates(&data).Error
}

func (m *dicmLocationImpl) DeleteRegion(ids []int64) error {
	data := new(entity.DcimLocation)
	var i int
	for i = 0; i < len(ids); i++ {
		global.Db.Model(entity.DcimRegion{}).First(data, ids[i])
		data.Deleted = time.Now()
		data.LastUpdated = time.Now()
		result := global.Db.Model(entity.DcimRegion{}).Where("id = ?", ids[i]).Updates(&data)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (m *dicmLocationImpl) GetRegionStructrue() ([]*entity.DcimRegion, error) {
	// 1. 定义一个临时结构体，仅用于从数据库读取必要的字段
	type TempNode struct {
		Id       int64
		Name     string
		ParentId sql.NullInt64
	}
	var flatNodes []entity.DcimRegion

	// 从数据库查询 id, name, parent_id
	global.Db.Model(&entity.DcimRegion{}).
		Select("id, name, parent_id").
		Find(&flatNodes)

	// 如果没有数据，直接返回空
	if len(flatNodes) == 0 {
		return []*entity.DcimRegion{}, nil
	}

	// 2. 构建层级结构
	//  创建一个map来存储所有节点，以ID为键，方便快速查找父节点
	nodeMap := make(map[int64]*entity.DcimRegion)

	//  初始化所有节点，并将它们存入map
	for _, fn := range flatNodes {
		nodeMap[fn.Id] = &entity.DcimRegion{
			Id:       fn.Id,
			Name:     fn.Name,
			ParentId: fn.ParentId,            // 保留ParentId字段
			Children: []*entity.DcimRegion{}, // 初始化Children切片
		}
	}

	//  构建树形结构：遍历所有节点，将它们连接到其父节点
	var roots []*entity.DcimRegion // 用于存储所有根节点 (没有父节点的节点)
	for _, fn := range flatNodes {
		node := nodeMap[fn.Id]  // 获取当前处理的节点
		if fn.ParentId != nil { // 如果存在父ID
			parentNode, exists := nodeMap[*fn.ParentId] // 尝试从map中获取父节点
			if exists {
				// 如果父节点存在，将当前节点添加到父节点的Children列表中
				parentNode.Children = append(parentNode.Children, node)
			} else {
				// 如果父ID存在，但map中找不到对应的父节点 (数据可能不一致或父节点未被查询到)
				// 这种情况下，可以根据业务逻辑处理，例如将其视为一个孤立的根节点或报错
				// 为简单起见，这里也将其视为一个根节点（尽管它有一个指向不存在的父节点的ParentId）
				log.Printf("Warning: Node ID %d has ParentId %d, but parent node not found. Treating as a root.\n", node.Id, fn.ParentId)
				roots = append(roots, node)
			}
		} else {
			// 如果没有父ID (ParentId.Valid is false)，说明这是一个根节点
			roots = append(roots, node)
		}
	}

	return roots, nil
}
