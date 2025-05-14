package entity

import "time"

// DcimSite
type DcimSite struct {
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
	GroupId         int64          `gorm:"column:group_id" json:"group_id"`
	RegionId        int64          `gorm:"column:region_id" json:"region_id"`
	TenantId        int64          `gorm:"column:tenant_id;" json:"tenant_id"` //,omitempty
	Tenant          *TenancyTenant `gorm:"foreignKey:TenantId;references:Id;belongsTo"`
}

// TableName 表名
func (d *DcimSite) TableName() string {
	return "dcim_site"
}
