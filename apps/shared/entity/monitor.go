package entity

// AssetMonitor
type AssetMonitor struct {
	Id         int32  `gorm:"primaryKey;column:id;default:nextval(asset_monitor_id_seq::regclass);NOT NULL" json:"id,omitempty"`
	CreateUid  int32  `gorm:"column:create_uid;comment:'created by'" json:"create_uid,omitempty"`
	WriteUid   int32  `gorm:"column:write_uid;comment:'last updated by'" json:"write_uid,omitempty"`
	Name       string `gorm:"column:name;comment:'监控设备编号'" json:"name,omitempty"`
	Status     string `gorm:"column:status;comment:'设备状态'" json:"status,omitempty"`
	Note       string `gorm:"column:note;comment:'说明'" json:"note,omitempty"`
	IsUse      string `gorm:"column:is_use;comment:'是否运转'" json:"is_use,omitempty"`
	CreateDate string `gorm:"column:create_date;comment:'created on'" json:"create_date,omitempty"`
	WriteDate  string `gorm:"column:write_date;comment:'last updated on'" json:"write_date,omitempty"`
}

// TableName 表名
func (a *AssetMonitor) TableName() string {
	return "asset_monitor"
}
