package entity

import (
	"database/sql"

	"github.com/PandaXGO/PandaKit/casbin"
	"github.com/PandaXGO/PandaKit/model"
)

const (
	SELFDATASCOPE   = "0"
	ALLDATASCOPE    = "1"
	DIYDATASCOPE    = "2"
	ORGDATASCOPE    = "3"
	ORGALLDATASCOPE = "4"
)

type SysRole struct {
	model.BaseModel
	RoleId          int64               `json:"roleId" gorm:"primary_key;AUTO_INCREMENT"`
	RoleName        string              `json:"roleName" gorm:"type:varchar(128);comment:角色名称"`
	Status          string              `json:"status" gorm:"type:varchar(1);comment:状态"`
	RoleKey         string              `json:"roleKey" gorm:"type:varchar(128);comment:角色代码"`
	RoleSort        int64               `json:"roleSort" gorm:"type:int;comment:角色排序"`
	DataScope       string              `json:"dataScope" gorm:"type:varchar(1);comment:数据范围（0: 本人数据 1：全部数据权限 2：自定数据权限 3：本组织数据权限 4：本组织及以下数据权限）"`
	CreateBy        string              `json:"createBy" gorm:"type:varchar(128);comment:创建人"`
	UpdateBy        string              `json:"updateBy" gorm:"type:varchar(128);comment:修改人"`
	Remark          string              `json:"remark" gorm:"type:varchar(255);comment:备注"`
	ApiIds          []casbin.CasbinRule `json:"apiIds" gorm:"-"`
	MenuIds         []int64             `json:"menuIds" gorm:"-"`
	OrganizationIds []int64             `json:"organizationIds" gorm:"-"`
}
type SysRoleAuth struct {
	Org       string `json:"org" gorm:"column:org"`
	DataScope string `json:"dataScope"`
}
type MenuIdList struct {
	MenuId int64 `json:"menuId"`
}

type OrganizationIdList struct {
	OrganizationId int64 `json:"organizationId"`
}

// Odoo User-Type
type ResGroups struct {
	Id             string         `gorm:"column:id;default:nextval(res_groups_id_seq::regclass);NOT NULL"`
	Name           string         `gorm:"column:name;NOT NULL"`
	CategoryId     sql.NullString `gorm:"column:category_id;comment:'application'"`
	Color          sql.NullString `gorm:"column:color;comment:'color index'"`
	CreateUid      sql.NullString `gorm:"column:create_uid;comment:'created by'"`
	WriteUid       sql.NullString `gorm:"column:write_uid;comment:'last updated by'"`
	Comment        sql.NullString `gorm:"column:comment;comment:'comment'"`
	Share          sql.NullBool   `gorm:"column:share;comment:'share group'"`
	CreateDate     sql.NullString `gorm:"column:create_date;comment:'created on'"`
	WriteDate      sql.NullString `gorm:"column:write_date;comment:'last updated on'"`
	ApiKeyDuration sql.NullString `gorm:"column:api_key_duration;comment:'api keys maximum duration days'"`
}
