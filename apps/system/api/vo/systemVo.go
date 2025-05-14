package vo

import (
	rentity "pandax/apps/shared/entity"
	"pandax/apps/system/entity"
)

/**
 * @Description
 * @Author 熊猫
 * @Date 2022/8/4 15:25
 **/

type OrganizationTreeVo struct {
	Organizations []entity.OrganizationLable `json:"organizations"`
	CheckedKeys   []int64                    `json:"checkedKeys"`
}

type MenuTreeVo struct {
	Menus       []entity.MenuLable `json:"menus"`
	CheckedKeys []int64            `json:"checkedKeys"`
}

type MenuPermisVo struct {
	Menus       []RouterVo `json:"menus"`
	Permissions []string   `json:"permissions"`
}

type CaptchaVo struct {
	Base64Captcha string `json:"base64Captcha"`
	CaptchaId     string `json:"captchaId"`
}

type TokenVo struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type AuthVo struct {
	User        entity.SysUserView `json:"user"`
	Role        entity.SysRole     `json:"role"`
	Permissions []string           `json:"permissions"`
	Menus       []RouterVo         `json:"menus"`
}

type AuthVoB struct {
	User rentity.ResUsersPage `json:"user"`
	// Role        entity.SysRole       `json:"role"`
	Permissions []string   `json:"permissions"`
	Menus       []RouterVo `json:"menus"`
}

type UserProfileVo struct {
	Data         any                      `json:"data"`
	PostIds      []int64                  `json:"postIds"`
	RoleIds      []int64                  `json:"roleIds"`
	Roles        []entity.SysRole         `json:"roles"`
	Posts        []entity.SysPost         `json:"posts"`
	Organization []entity.SysOrganization `json:"organization"`
}

type UserProfileVoB struct {
	Data    any     `json:"data"`
	PostIds []int64 `json:"postIds"`
	RoleIds []int64 `json:"roleIds"`
	// Roles        []entity.SysRole         `json:"roles"`
	Posts      []rentity.HrJob              `json:"posts"`
	Department rentity.CompanyWithDepatment `json:"department"`
}

type UserVo struct {
	Data          any                      `json:"data"`
	PostIds       string                   `json:"postIds"`
	RoleIds       string                   `json:"roleIds"`
	Roles         []entity.SysRole         `json:"roles"`
	Posts         []entity.SysPost         `json:"posts"`
	Organizations []entity.SysOrganization `json:"organizations"`
}

type UserVoB struct {
	Data    any   `json:"data"`
	PostIds int64 `json:"postIds"`
	// RoleIds string  `json:"roleIds"`
	// Roles         []entity.SysRole         `json:"roles"`
	Posts     rentity.HrJob                `json:"posts"`
	Depatment rentity.CompanyWithDepatment `json:"organizations"`
}

type UserRolePost struct {
	Roles []entity.SysRole `json:"roles"`
	Posts []entity.SysPost `json:"posts"`
}
type UserRolePostB struct {
	// Roles []entity.SysRole `json:"roles"`
	Posts []rentity.HrJob `json:"posts"`
}
