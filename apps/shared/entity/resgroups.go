package entity

// ResGroups
type ResGroups struct {
	Id             int64               `gorm:"column:id;default:nextval(res_groups_id_seq::regclass);NOT NULL" json:"id,omitempty"`
	Name           string              `gorm:"column:name;NOT NULL" json:"name,omitempty"`
	CategoryId     int64               `gorm:"column:category_id" json:"category_id,omitempty"`
	Color          int64               `gorm:"column:color" json:"color,omitempty"`
	CreateUid      int64               `gorm:"column:create_uid" json:"create_uid,omitempty"`
	WriteUid       int64               `gorm:"column:write_uid" json:"write_uid,omitempty"`
	Comment        string              `gorm:"column:comment" json:"comment,omitempty"`
	Share          string              `gorm:"column:share" json:"share,omitempty"`
	CreateDate     string              `gorm:"column:create_date" json:"create_date,omitempty"`
	WriteDate      string              `gorm:"column:write_date" json:"write_date,omitempty"`
	ApiKeyDuration string              `gorm:"column:api_key_duration" json:"api_key_duration,omitempty"`
	Users          []ResGroupsUsersRel `gorm:"foreignKey:Gid" json:"users,omitempty"`
}

func (r *ResGroups) TableName() string {
	return "res_groups"
}

// ResGroupsUsersRel
type ResGroupsUsersRel struct {
	Gid int64 `gorm:"column:gid;NOT NULL" json:"gid,omitempty"`
	Uid int64 `gorm:"column:uid;NOT NULL" json:"uid,omitempty"`
}

func (ResGroupsUsersRel) TableName() string {
	return "res_groups_users_rel"
}
