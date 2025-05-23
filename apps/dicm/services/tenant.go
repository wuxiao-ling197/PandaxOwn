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
	TenantModel interface {
		/* 租户 */
		FindList(data entity.TenancyTenant) (*[]entity.TenancyTenant, error)
		FindListPage(page, pageSize int, data entity.TenancyTenant) (*[]entity.TenancyTenant, int64, error)
		FindOne(data entity.TenancyTenant) (*entity.TenancyTenant, error)
		Insert(data entity.TenancyTenant) (*entity.TenancyTenant, error)
		Update(data entity.TenancyTenant) error
		Delete(ids []int64) error
		// 租户加入租户组
		JoinGroup(tenants []int64, groupId int64) error
		/* 租户组 */
		FindGroupList(data entity.TenancyTenantgroup) (*[]entity.TenancyTenantgroup, error)
		FindListGroupPage(page, pageSize int, data entity.TenancyTenantgroup) (*[]entity.TenancyTenantgroup, int64, error)
		FindGroupOne(data entity.TenancyTenantgroup) (*entity.TenancyTenantgroup, error)
		InsertGroup(data entity.TenancyTenantgroup) (*entity.TenancyTenantgroup, error)
		UpdateGroup(data entity.TenancyTenantgroup) error
		DeleteGroup(ids []int64) error
		// 租户组层级结构
		GetGroupStructrue() ([]*entity.GroupNode, error)
	}

	tenantImpl struct {
		table string
	}
)

var TenantDao TenantModel = &tenantImpl{
	table: `tenancy_tenant`,
}

func (m *tenantImpl) FindList(data entity.TenancyTenant) (*[]entity.TenancyTenant, error) {
	list := make([]entity.TenancyTenant, 0)
	db := global.Db.Table(m.table)
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	// if data.GroupId.Valid == true {
	// 	db = db.Where("group_id = ?", data.GroupId)
	// }
	if data.GroupId != 0 {
		db = db.Where("group_id = ?", data.GroupId)
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
	err := db.Find(&list).Error
	return &list, err

}

func (m *tenantImpl) FindListPage(page, pageSize int, data entity.TenancyTenant) (*[]entity.TenancyTenant, int64, error) {
	list := make([]entity.TenancyTenant, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	// if data.GroupId.Int64 != 0 {
	// 	db = db.Where("group_id = ?", data.GroupId.Int64)
	// }
	if data.GroupId != 0 {
		db = db.Where("group_id = ?", data.GroupId)
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
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error

	return &list, total, err
}

func (m *tenantImpl) FindOne(data entity.TenancyTenant) (*entity.TenancyTenant, error) {
	resData := new(entity.TenancyTenant)
	db := global.Db.Table(m.table)
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	// if data.GroupId.Int64 != 0 {
	// 	db = db.Where("group_id = ?", data.GroupId)
	// }
	if data.GroupId != 0 {
		db = db.Where("group_id = ?", data.GroupId)
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
	err := db.First(resData).Error
	return resData, err
}

func (m *tenantImpl) Insert(data entity.TenancyTenant) (*entity.TenancyTenant, error) {
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
	log.Printf("错误=%+v\n", err)
	return &data, err
}

func (m *tenantImpl) Update(data entity.TenancyTenant) error {
	update := new(entity.TenancyTenant)
	err := global.Db.Table(m.table).First(update, data.Id).Error
	if err != nil {
		return err
	}

	// if data.GroupId.Valid == true {
	// 	data.GroupId = update.GroupId
	// }
	if data.GroupId != 0 {
		data.GroupId = update.GroupId
	}

	data.LastUpdated = time.Now()
	return global.Db.Table(m.table).Updates(&data).Error
}

func (m *tenantImpl) Delete(ids []int64) error {
	data := new(entity.TenancyTenant)
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

// 租户加入租户组
func (m *tenantImpl) JoinGroup(tenants []int64, groupId int64) error {
	var i int
	for i = 0; i < len(tenants); i++ {
		result := global.Db.Model(&entity.TenancyTenant{}).Where("id = ?", tenants[i]).Update("group_id", groupId)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

/** 租户组 */
func (m *tenantImpl) FindGroupList(data entity.TenancyTenantgroup) (*[]entity.TenancyTenantgroup, error) {
	list := make([]entity.TenancyTenantgroup, 0)
	db := global.Db.Model(&entity.TenancyTenantgroup{})
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	// if data.ParentId.Int64 != 0 {
	// 	db = db.Where("parent_id = ?", data.ParentId)
	// }
	if data.ParentId != nil {
		db = db.Where("parent_id = ?", data.ParentId)
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
	err := db.Find(&list).Error
	return &list, err

}

func (m *tenantImpl) FindListGroupPage(page, pageSize int, data entity.TenancyTenantgroup) (*[]entity.TenancyTenantgroup, int64, error) {
	list := make([]entity.TenancyTenantgroup, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Model(&entity.TenancyTenantgroup{})
	// 此处填写 where参数判断
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	// if data.ParentId.Int64 != 0 {
	// 	db = db.Where("parent_id = ?", data.ParentId.Int64)
	// }
	if data.ParentId != nil {
		db = db.Where("parent_id = ?", data.ParentId)
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
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error

	return &list, total, err
}

func (m *tenantImpl) FindGroupOne(data entity.TenancyTenantgroup) (*entity.TenancyTenantgroup, error) {
	resData := new(entity.TenancyTenantgroup)
	db := global.Db.Model(&entity.TenancyTenantgroup{})
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	// if data.ParentId.Int64 != 0 {
	// 	db = db.Where("parent_id = ?", data.ParentId)
	// }
	if data.ParentId != nil {
		db = db.Where("parent_id = ?", data.ParentId)
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
	err := db.First(resData).Error
	return resData, err
}

// 新建时由于 pq: 插入或更新表 "tenancy_tenantgroup" 违反外键约束 "tenancy_tenantgroup_parent_id_2542fc18_fk_tenancy_t" parentId不能为空必须设置有效值,可以通过设置*int64指针类型解决
func (m *tenantImpl) InsertGroup(data entity.TenancyTenantgroup) (*entity.TenancyTenantgroup, error) {
	var count int64
	global.Db.Model(&entity.TenancyTenantgroup{}).Where("name = ? ", data.Name).Count(&count)
	if count != 0 {
		return nil, errors.New("该实例已存在！")
	}
	global.Db.Model(&entity.TenancyTenantgroup{}).Where("slug = ? ", data.Slug).Count(&count)
	if count != 0 {
		return nil, errors.New("该标识符已存在！")
	}
	data.Created = time.Now()
	err := global.Db.Model(&entity.TenancyTenantgroup{}).Create(&data).Error
	log.Printf("错误=%+v\n", err)
	return &data, err
}

// 如果需要将ParentId从有效值设置回0即删除上下级关系不会有效修改，只能从默认值设置为有效值且只能在有效值之间修改
func (m *tenantImpl) UpdateGroup(data entity.TenancyTenantgroup) error {
	update := new(entity.TenancyTenantgroup)
	err := global.Db.Model(&update).First(update, data.Id).Error
	if err != nil {
		return err
	}

	// if data.ParentId == 0 {
	// 	data.ParentId = update.ParentId
	// }

	data.LastUpdated = time.Now()
	return global.Db.Model(&update).Updates(&data).Error
}

func (m *tenantImpl) DeleteGroup(ids []int64) error {
	data := new(entity.TenancyTenantgroup)
	var i int
	for i = 0; i < len(ids); i++ {
		global.Db.Model(&entity.TenancyTenantgroup{}).First(data, ids[i])
		data.Deleted = time.Now()
		data.LastUpdated = time.Now()
		result := global.Db.Model(&entity.TenancyTenantgroup{}).Where("id = ?", ids[i]).Updates(&data)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// 获取租户组层级结构
func (m *tenantImpl) GetGroupStructrue() ([]*entity.GroupNode, error) {
	// 1. 定义一个临时结构体，仅用于从数据库读取必要的字段
	type TempNode struct {
		Id       int64
		Name     string
		ParentId sql.NullInt64
	}
	var flatNodes []TempNode

	// 从数据库查询 id, name, parent_id
	global.Db.Model(&entity.TenancyTenantgroup{}).
		Select("id, name, parent_id").
		Find(&flatNodes)

	// if result.Error != nil {
	// 	log.Printf("Error fetching flat organization nodes: %v\n", result.Error)
	// 	return nil, result.Error
	// }

	// 如果没有数据，直接返回空
	if len(flatNodes) == 0 {
		return []*entity.GroupNode{}, nil
	}

	// 2. 构建层级结构
	//  创建一个map来存储所有节点，以ID为键，方便快速查找父节点
	nodeMap := make(map[int64]*entity.GroupNode)

	//  初始化所有节点，并将它们存入map
	for _, fn := range flatNodes {
		nodeMap[fn.Id] = &entity.GroupNode{
			Id:       fn.Id,
			Name:     fn.Name,
			ParentId: fn.ParentId,           // 保留ParentId字段
			Children: []*entity.GroupNode{}, // 初始化Children切片
		}
	}

	//  构建树形结构：遍历所有节点，将它们连接到其父节点
	var roots []*entity.GroupNode // 用于存储所有根节点 (没有父节点的节点)
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
