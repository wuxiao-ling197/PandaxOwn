package entity

import (
	"time"
)

// DcimRack
type DcimRack struct {
	Deleted         time.Time      `gorm:"column:deleted" json:"deleted"`
	Created         time.Time      `gorm:"column:created" json:"created"` //,omitempty
	LastUpdated     time.Time      `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string         `gorm:"column:custom_field_data;NOT NULL" json:"custom_field_data"`
	Id              int64          `gorm:"primaryKey;autoIncrement;column:id;NOT NULL" json:"id"` //autoIncrement
	Name            string         `gorm:"column:name;NOT NULL" json:"name"`
	NickName        string         `gorm:"column:_name;NOT NULL" json:"_name"`
	FacilityId      string         `gorm:"column:facility_id" json:"facility_id"`
	Status          string         `gorm:"column:status;NOT NULL" json:"status"`
	Serial          int32          `gorm:"column:serial;AUTO_INCREMENT;NOT NULL" json:"serial,string"`
	AssetTag        string         `gorm:"column:asset_tag;comment:'资产标签，确保唯一性'" json:"asset_tag"`
	Type            string         `gorm:"column:type;NOT NULL" json:"type"`
	Width           int16          `gorm:"column:width;NOT NULL" json:"width,string"`
	UHeight         int16          `gorm:"column:u_height;NOT NULL" json:"u_height,string"`
	DescUnits       bool           `gorm:"column:desc_units;NOT NULL" json:"desc_units"`
	OuterWidth      int16          `gorm:"column:outer_width" json:"outer_width,string"`
	OuterDepth      int16          `gorm:"column:outer_depth" json:"outer_depth,string"`
	OuterUnit       string         `gorm:"column:outer_unit;NOT NULL" json:"outer_unit"`
	Comments        string         `gorm:"column:comments;NOT NULL" json:"comments"`
	LocationId      int64          `gorm:"column:location_id;" json:"location_id"`
	Location        *DcimLocation  `gorm:"foreignKey:LocationId;references:Id;belongsTo"`
	RoleId          int64          `gorm:"column:role_id;" json:"role_id"`
	Role            *DcimRackrole  `gorm:"foreignKey:RoleId;references:Id;belongsTo"`
	SiteId          int64          `gorm:"column:site_id;NOT NULL;" json:"site_id"`
	Site            *DcimSite      `gorm:"foreignKey:SiteId;references:Id;belongsTo"`
	TenantId        int64          `gorm:"column:tenant_id;" json:"tenant_id"`
	Tenant          *TenancyTenant `gorm:"foreignKey:TenantId;references:Id;belongsTo"`
	Weight          float64        `gorm:"column:weight" json:"weight,string"`
	MaxWeight       int32          `gorm:"column:max_weight" json:"max_weight,string"`
	WeightUnit      string         `gorm:"column:weight_unit;NOT NULL" json:"weight_unit"`
	AbsWeight       int64          `gorm:"column:_abs_weight" json:"_abs_weight,string"`
	AbsMaxWeight    int64          `gorm:"column:_abs_max_weight" json:"_abs_max_weight,string"`
	MountingDepth   int16          `gorm:"column:mounting_depth" json:"mounting_depth,string"`
	Description     string         `gorm:"column:description;NOT NULL" json:"description"`
	StartingUnit    int16          `gorm:"column:starting_unit;NOT NULL" json:"starting_unit,string"`
}

// TableName 表名
func (d *DcimRack) TableName() string {
	return "dcim_rack"
}

// DcimRackrole
type DcimRackrole struct {
	Deleted         time.Time `gorm:"column:deleted" json:"deleted"`
	Created         time.Time `gorm:"column:created" json:"created"`
	LastUpdated     time.Time `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string    `gorm:"column:custom_field_data;NOT NULL" json:"custom_field_data"`
	Id              int64     `gorm:"column:id;autoIncrement;NOT NULL" json:"id"` //default:AS;
	Name            string    `gorm:"column:name;NOT NULL" json:"name"`
	Slug            string    `gorm:"column:slug;NOT NULL" json:"slug"`
	Color           string    `gorm:"column:color;NOT NULL" json:"color"`
	Description     string    `gorm:"column:description;NOT NULL" json:"description"`
}

// TableName 表名
func (d *DcimRackrole) TableName() string {
	return "dcim_rackrole"
}

// DcimRackreservation 机柜预留
type DcimRackreservation struct {
	Deleted         time.Time      `gorm:"column:deleted" json:"deleted"`
	Created         time.Time      `gorm:"column:created" json:"created"`
	LastUpdated     time.Time      `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string         `gorm:"column:custom_field_data;NOT NULL" json:"custom_field_data"`
	Id              int64          `gorm:"column:id;autoIncrement;NOT NULL" json:"id"`
	Units           int16          `gorm:"column:units;NOT NULL" json:"units"`
	Description     string         `gorm:"column:description;NOT NULL" json:"description"`
	RackId          *int64         `gorm:"column:rack_id;NOT NULL" json:"rack_id"`
	Rack            *DcimRack      `gorm:"foreignKey:RackId;references:Id;belongsTo"`
	TenantId        *int64         `gorm:"column:tenant_id" json:"tenant_id"`
	Tenant          *TenancyTenant `gorm:"foreignKey:TenantId;references:Id;belongsTo"`
	UserId          *int64         `gorm:"column:user_id;NOT NULL" json:"user_id"`
	Comments        string         `gorm:"column:comments;NOT NULL" json:"comments"`
}

// TableName 表名
func (d *DcimRackreservation) TableName() string {
	return "dcim_rackreservation"
}
