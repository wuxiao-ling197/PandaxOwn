package entity

import "time"

// TenancyTenant
type TenancyTenant struct {
	Created         time.Time `gorm:"column:created" json:"created"`
	LastUpdated     time.Time `gorm:"column:last_updated" json:"last_updated"`
	CustomFieldData string    `gorm:"column:custom_field_data;NOT NULL" json:"custom_field_data"`
	Id              int64     `gorm:"column:id;autoIncrement;NOT NULL" json:"id"` //autoIncrement
	Name            string    `gorm:"column:name;NOT NULL" json:"name"`
	Slug            string    `gorm:"column:slug;NOT NULL" json:"slug"`
	Description     string    `gorm:"column:description;NOT NULL" json:"description"`
	Comments        string    `gorm:"column:comments;NOT NULL" json:"comments"`
	GroupId         *int64    `gorm:"column:group_id;foreignKey:GroupId;references:Id" json:"group_id"`
}

// TableName 表名
func (t *TenancyTenant) TableName() string {
	return "tenancy_tenant"
}
