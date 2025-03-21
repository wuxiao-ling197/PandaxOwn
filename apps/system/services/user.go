package services

import (
	"errors"
	"pandax/apps/system/entity"
	"pandax/pkg/global"

	"github.com/kakuilan/kgo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/hlandau/passlib.v1/abstract"
	"gopkg.in/hlandau/passlib.v1/hash/pbkdf2"
)

type (
	SysUserModel interface {
		Login(u entity.Login) (*entity.SysUser, error)
		Insert(data entity.SysUser) (*entity.SysUser, error)
		FindOne(data entity.SysUser) (resData *entity.SysUserView, err error)
		FindListPage(page, pageSize int, data entity.SysUser) (list *[]entity.SysUserPage, total int64, err error)
		FindList(data entity.SysUser) (list *[]entity.SysUserView, err error)
		Update(data entity.SysUser) error
		Delete(userId []int64) error
		SetPwd(data entity.SysUser, pwd entity.SysUserPwd) error
	}

	sysUserModelImpl struct {
		table string
	}
)

var SysUserModelDao SysUserModel = &sysUserModelImpl{
	table: `sys_users`,
}

type Context struct {
	Schemes []abstract.Scheme
}

var defualtSchemes = []abstract.Scheme{
	pbkdf2.SHA512Crypter,
}

// Pandax 原实现
func (m *sysUserModelImpl) Login(u entity.Login) (*entity.SysUser, error) {
	user := new(entity.SysUser)
	err := global.Db.Table(m.table).Where("username = ? ", u.Username).Find(user).Error
	if err != nil {
		return nil, errors.New("查询用户信息失败")
	}
	global.Log.Infof("用户后端服务login user= %+v", user)
	// 验证密码 bcrypt.hash
	b := kgo.KEncr.PasswordVerify([]byte(u.Password), []byte(user.Password))
	if !b {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (m *sysUserModelImpl) Insert(data entity.SysUser) (*entity.SysUser, error) {
	bytes, _ := kgo.KEncr.PasswordHash([]byte(data.Password), bcrypt.DefaultCost)
	data.Password = string(bytes)

	// check 用户名
	var count int64
	global.Db.Table(m.table).Where("username = ? and delete_time IS NULL", data.Username).Count(&count)
	if count != 0 {
		return nil, errors.New("账户已存在！")
	}
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *sysUserModelImpl) FindOne(data entity.SysUser) (*entity.SysUserView, error) {
	resData := new(entity.SysUserView)

	db := global.Db.Table(m.table).Select([]string{"sys_users.*", "sys_roles.role_name"})
	db = db.Joins("left join sys_roles on sys_users.role_id=sys_roles.role_id")
	if data.UserId != 0 {
		db = db.Where("user_id = ?", data.UserId)
	}
	if data.Username != "" {
		db = db.Where("username = ?", data.Username)
	}
	if data.Password != "" {
		db = db.Where("password = ?", data.Password)
	}
	if data.RoleId != 0 {
		db = db.Where("role_id = ?", data.RoleId)
	}
	if data.OrganizationId != 0 {
		db = db.Where("organization_id = ?", data.OrganizationId)
	}
	if data.PostId != 0 {
		db = db.Where("post_id = ?", data.PostId)
	}
	err := db.First(resData).Error

	return resData, err
}

func (m *sysUserModelImpl) FindListPage(page, pageSize int, data entity.SysUser) (*[]entity.SysUserPage, int64, error) {
	list := make([]entity.SysUserPage, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	// fmt.Printf("user global.Db=  %v  连接成功", global.Db)
	// global.Log.Infof("user global.Db=  %v  连接成功", global.Db)
	db := global.Db.Table(m.table).Select("sys_users.*,sys_organizations.organization_name")
	db = db.Joins("left join sys_organizations on sys_organizations.organization_id = sys_users.organization_id")
	// 此处填写 where参数判断
	if data.Username != "" {
		db = db.Where("sys_users.username = ?", data.Username)
	}
	if data.NickName != "" {
		db = db.Where("sys_users.nick_name like ?", "%"+data.NickName+"%")
	}

	if data.Status != "" {
		db = db.Where("sys_users.status = ?", data.Status)
	}

	if data.Phone != "" {
		db = db.Where("sys_users.phone like ?", "%"+data.Phone+"%")
	}
	if data.OrganizationId != 0 {
		db = db.Where("sys_users.organization_id = ?", data.OrganizationId)
	}
	db.Where("sys_users.delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	global.Log.Infof("用户后端服务 FindListPage %+v", &list)
	return &list, total, err
}

func (m *sysUserModelImpl) FindList(data entity.SysUser) (*[]entity.SysUserView, error) {
	list := make([]entity.SysUserView, 0)
	// 此处填写 where参数判断
	db := global.Db.Table(m.table).Select([]string{"sys_users.*", "sys_roles.role_name"})
	db = db.Joins("left join sys_roles on sys_users.role_id=sys_roles.role_id")
	if data.UserId != 0 {
		db = db.Where("user_id = ?", data.UserId)
	}
	if data.Username != "" {
		db = db.Where("username = ?", data.Username)
	}

	if data.Password != "" {
		db = db.Where("password = ?", data.Password)
	}

	if data.RoleId != 0 {
		db = db.Where("sys_users.role_id = ?", data.RoleId)
	}

	if data.OrganizationId != 0 {
		db = db.Where("sys_users.organization_id = ?", data.OrganizationId)
	}

	if data.PostId != 0 {
		db = db.Where("sys_users.post_id = ?", data.PostId)
	}
	if data.Status != "" {
		db = db.Where("sys_users.status = ?", data.Status)
	}
	db.Where("sys_users.delete_time IS NULL")

	err := db.Find(&list).Error
	global.Log.Infof("用户后端服务 FindList %+v", &list)
	return &list, err
}

func (m *sysUserModelImpl) Update(data entity.SysUser) error {
	if data.Password != "" {
		bytes, _ := kgo.KEncr.PasswordHash([]byte(data.Password), bcrypt.DefaultCost)
		data.Password = string(bytes)
	}
	update := new(entity.SysUser)
	err := global.Db.Table(m.table).First(update, data.UserId).Error
	if err != nil {
		return err
	}

	if data.RoleId == 0 {
		data.RoleId = update.RoleId
	}

	return global.Db.Table(m.table).Updates(&data).Error
}

func (m *sysUserModelImpl) Delete(userIds []int64) error {
	return global.Db.Table(m.table).Delete(&entity.SysUser{}, "user_id in (?)", userIds).Error
}

func (m *sysUserModelImpl) SetPwd(data entity.SysUser, pwd entity.SysUserPwd) error {
	user, err := m.FindOne(data)
	if err != nil {
		return err
	}
	bl := kgo.KEncr.PasswordVerify([]byte(pwd.OldPassword), []byte(user.Password))
	if !bl {
		return errors.New("旧密码输入错误")
	}

	data.Password = pwd.NewPassword
	return m.Update(data)
}

// odoo migrate
// func (m *sysUserModelImpl) LoginO(ro entity.LoginO) (*entity.ResUsers, error) {
// 	user := new(entity.ResUsers)
// 	err := global.HrDb.Table(m.table).Where("login = ? ", ro.Login).Find(user).Error
// 	if err != nil {
// 		return nil, errors.New("用户查询失败")
// 	}
// 	global.Log.Infof("Login User= %+v", user)
// 	//验证密码 pbkdf2-sha512
// 	b := kgo.KEncr.PasswordVerify([]byte(ro.Password), []byte(user.Password))
// 	if !b {
// 		return nil, errors.New("密码错误")
// 	}
// 	return user, nil
// }
// // 加密算法
// func HashPwd(password string) string {
// 	var hash string
// 	for _, scheme := range defualtSchemes {
// 		ctx := Context{Schemes: []abstract.Scheme{scheme}}
// 		hash, _ = ctx.Schemes[0].Hash(password) //casdoor加密后hashf("hash=%v\n", hash)
// 	}
// 	return hash
// }
// 验证算法
// func VerifyPwd(data entity.LoginO, hashedPwd string) bool {
// 	flag := true
// 	for _, scheme := range defualtSchemes {
// 		ctx := Context{Schemes: []abstract.Scheme{scheme}}
// 		st := ctx.Schemes[0].Verify(data.Password, hashedPwd)
// 		if st != nil {
// 			flag = false
// 		}
// 	}
// 	return flag
// }
