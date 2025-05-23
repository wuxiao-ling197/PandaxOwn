package entity

import (
	"database/sql"
	"time"
)

// TenancyTenant
type TenancyTenant struct {
	Deleted         time.Time `gorm:"column:deleted" json:"deleted"`
	Created         time.Time `gorm:"column:created" json:"created"`
	LastUpdated     time.Time `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string    `gorm:"column:custom_field_data;NOT NULL" json:"custom_field_data"`
	Id              int64     `gorm:"column:id;autoIncrement;NOT NULL" json:"id"` //autoIncrement
	Name            string    `gorm:"column:name;NOT NULL" json:"name"`
	Slug            string    `gorm:"column:slug;NOT NULL" json:"slug"`
	Description     string    `gorm:"column:description;NOT NULL" json:"description"`
	Comments        string    `gorm:"column:comments;NOT NULL" json:"comments"`
	// GroupId         sql.NullInt64       `gorm:"column:group_id;default:null" json:"group_id"`
	GroupId     int64               `gorm:"column:group_id;default:null" json:"group_id"`
	TenantGroup *TenancyTenantgroup `gorm:"foreignKey:GroupId;references:Id;belongsTo"`
}

// TableName 表名
func (t *TenancyTenant) TableName() string {
	return "tenancy_tenant"
}

// 处理租户、租户组所属关系
type Own2Group struct {
	TenantIds []int64 `json:"tenant_ids"`
	GroupId   int64   `json:"group_id"`
}

// TenancyTenantgroup
type TenancyTenantgroup struct {
	// gorm.Model
	Deleted         time.Time           `gorm:"column:deleted" json:"deleted"`
	Created         time.Time           `gorm:"column:created" json:"created"`
	LastUpdated     time.Time           `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string              `gorm:"column:custom_field_data;NOT NULL" json:"custom_field_data"`
	Id              int64               `gorm:"column:id;autoIncrement;NOT NULL" json:"id"`
	Name            string              `gorm:"column:name;NOT NULL" json:"name"`
	Slug            string              `gorm:"column:slug;NOT NULL" json:"slug"`
	Description     string              `gorm:"column:description;NOT NULL" json:"description"`
	Lft             int64               `gorm:"column:lft;NOT NULL" json:"lft,string"`
	Rght            int64               `gorm:"column:rght;NOT NULL" json:"rght,string"`
	TreeId          int64               `gorm:"column:tree_id;NOT NULL" json:"tree_id,string"`
	Level           int64               `gorm:"column:level;NOT NULL" json:"level,string"`
	ParentId        *int64              `gorm:"column:parent_id" json:"parent_id"` //sql.NullInt64
	Group           *TenancyTenantgroup `gorm:"foreignKey:ParentId;references:Id;belongsTo"`
}

// TableName 表名
func (t *TenancyTenantgroup) TableName() string {
	return "tenancy_tenantgroup"
}

type GroupNode struct {
	Id       int64         `gorm:"column:id;NOT NULL" json:"id"`
	Name     string        `gorm:"column:name;NOT NULL" json:"name"`
	ParentId sql.NullInt64 `gorm:"column:parent_id" json:"parent_id"`
	Children []*GroupNode  `json:"children"`
}
