package entity

import (
	"database/sql"
	"time"
)

// DcimLocation
type DcimLocation struct {
	Deleted         time.Time      `gorm:"column:deleted" json:"deleted"`
	Created         time.Time      `gorm:"column:created" json:"created"`
	LastUpdated     time.Time      `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string         `gorm:"column:custom_field_data;NOT NULL" json:"custom_field_data"`
	Id              int64          `gorm:"column:id;autoIncrement;NOT NULL" json:"id"` //autoIncrement
	Name            string         `gorm:"column:name;NOT NULL" json:"name"`
	Slug            string         `gorm:"column:slug;NOT NULL" json:"slug"`
	Description     string         `gorm:"column:description;NOT NULL" json:"description"`
	Lft             int32          `gorm:"column:lft;NOT NULL" json:"lft"`
	Rght            int32          `gorm:"column:rght;NOT NULL" json:"rght"`
	TreeId          int32          `gorm:"column:tree_id;NOT NULL" json:"tree_id"`
	Level           int32          `gorm:"column:level;NOT NULL" json:"level"`
	ParentId        *int64         `gorm:"column:parent_id" json:"parent_id"`
	Parent          *DcimLocation  `gorm:"foreignKey:ParentId;references:Id;belongsTo"`
	SiteId          int64          `gorm:"column:site_id;NOT NULL" json:"site_id"`
	Site            *DcimSite      `gorm:"foreignKey:SiteId;references:Id;belongsTo"`
	TenantId        int64          `gorm:"column:tenant_id;" json:"tenant_id"`
	Tenant          *TenancyTenant `gorm:"foreignKey:TenantId;references:Id;belongsTo"`
	Status          string         `gorm:"column:status;NOT NULL" json:"status"`
	Facility        string         `gorm:"column:facility;NOT NULL" json:"facility"`
}

// TableName 表名
func (d *DcimLocation) TableName() string {
	return "dcim_location"
}

type LocationNode struct {
	Id       int64           `gorm:"column:id;NOT NULL" json:"id"`
	Name     string          `gorm:"column:name;NOT NULL" json:"name"`
	ParentId sql.NullInt64   `gorm:"column:parent_id" json:"parent_id"`
	Children []*LocationNode `json:"children"`
}

// DcimRegion
type DcimRegion struct {
	Created         time.Time     `gorm:"column:created" json:"created"`
	LastUpdated     time.Time     `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string        `gorm:"column:custom_field_data;NOT NULL" json:"custom_field_data"`
	Id              int64         `gorm:"column:id;autoIncrement;NOT NULL" json:"id"`
	Name            string        `gorm:"column:name;NOT NULL" json:"name"`
	Slug            string        `gorm:"column:slug;NOT NULL" json:"slug"`
	Description     string        `gorm:"column:description;NOT NULL" json:"description"`
	Lft             int32         `gorm:"column:lft;NOT NULL" json:"lft"`
	Rght            int32         `gorm:"column:rght;NOT NULL" json:"rght"`
	TreeId          int32         `gorm:"column:tree_id;NOT NULL" json:"tree_id"`
	Level           int32         `gorm:"column:level;NOT NULL" json:"level"`
	ParentId        *int64        `gorm:"column:parent_id" json:"parent_id"`
	Children        []*DcimRegion `json:"children"`
	Deleted         time.Time     `gorm:"column:deleted" json:"deleted"`
}

// TableName 表名
func (d *DcimRegion) TableName() string {
	return "dcim_region"
}
