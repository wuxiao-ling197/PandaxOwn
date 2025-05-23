package entity

import (
	"database/sql"
	"time"
)

// DcimSite
type DcimSite struct {
	Deleted         time.Time      `gorm:"column:deleted" json:"deleted"`
	Created         time.Time      `gorm:"column:created" json:"created"`
	LastUpdated     time.Time      `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string         `gorm:"column:custom_field_data;NOT NULL" json:"custom_field_data"`
	Id              int64          `gorm:"primaryKey;column:id;autoIncrement;NOT NULL" json:"id"` //autoIncrement
	Name            string         `gorm:"column:name;NOT NULL" json:"name"`
	NickNameName    string         `gorm:"column:_name;NOT NULL" json:"_name"`
	Slug            string         `gorm:"column:slug;NOT NULL" json:"slug"`
	Status          string         `gorm:"column:status;NOT NULL" json:"status"`
	Facility        string         `gorm:"column:facility;NOT NULL" json:"facility"`
	TimeZone        string         `gorm:"column:time_zone;NOT NULL" json:"time_zone"`
	Description     string         `gorm:"column:description;NOT NULL" json:"description"`
	PhysicalAddress string         `gorm:"column:physical_address;NOT NULL" json:"physical_address"`
	ShippingAddress string         `gorm:"column:shipping_address;NOT NULL" json:"shipping_address"`
	Latitude        string         `gorm:"column:latitude" json:"latitude"`
	Longitude       string         `gorm:"column:longitude" json:"longitude"`
	Comments        string         `gorm:"column:comments;NOT NULL" json:"comments"`
	GroupId         *int64         `gorm:"column:group_id" json:"group_id"`
	Group           *DcimSitegroup `gorm:"foreignKey:GroupId;references:Id;belongsTo"`
	RegionId        *int64         `gorm:"column:region_id" json:"region_id"`
	Region          *DcimRegion    `gorm:"foreignKey:RegionId;references:Id;belongsTo"`
	TenantId        *int64         `gorm:"column:tenant_id;" json:"tenant_id"` //,omitempty
	Tenant          *TenancyTenant `gorm:"foreignKey:TenantId;references:Id;belongsTo"`
}

// TableName 表名
func (d *DcimSite) TableName() string {
	return "dcim_site"
}

// DcimSitegroup
type DcimSitegroup struct {
	Created         time.Time      `gorm:"column:created" json:"created"`
	LastUpdated     time.Time      `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string         `gorm:"column:custom_field_data;NOT NULL" json:"custom_field_data"`
	Id              int64          `gorm:"column:id;autoIncrement;NOT NULL" json:"id"`
	Name            string         `gorm:"column:name;NOT NULL" json:"name"`
	Slug            string         `gorm:"column:slug;NOT NULL" json:"slug"`
	Description     string         `gorm:"column:description;NOT NULL" json:"description"`
	Lft             int32          `gorm:"column:lft;NOT NULL" json:"lft,string"`
	Rght            int32          `gorm:"column:rght;NOT NULL" json:"rght,string"`
	TreeId          int32          `gorm:"column:tree_id;NOT NULL" json:"tree_id,string"`
	Level           int32          `gorm:"column:level;NOT NULL" json:"level,string"`
	ParentId        *int64         `gorm:"column:parent_id" json:"parent_id"`
	Children        *DcimSitegroup `gorm:"foreignKey:ParentId;references:Id;belongsTo"`
	Deleted         time.Time      `gorm:"column:deleted" json:"deleted"`
}

// TableName 表名
func (d *DcimSitegroup) TableName() string {
	return "dcim_sitegroup"
}

// 处理站点、站点组所属关系
type Site2Group struct {
	SiteIds []int64 `json:"site_ids"`
	GroupId int64   `json:"group_id"`
}

type SiteNode struct {
	Id       int64         `gorm:"column:id;NOT NULL" json:"id"`
	Name     string        `gorm:"column:name;NOT NULL" json:"name"`
	ParentId sql.NullInt64 `gorm:"column:parent_id" json:"parent_id"`
	Children []*SiteNode   `json:"children"`
}
