package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"pandax/apps/shared/entity"
	odoorpc "pandax/pkg/device_rpc"
	"pandax/pkg/global"
	"pandax/pkg/tool"
)

type (
	ResUsersModel interface {
		Login(u entity.LoginO) (*entity.ResUsersPage, error)
		// 新增用户 同步创建员工和联系人
		InsertUser(data entity.CreateUserDto) (odoorpc.ResultData, error)
		// 新增员工 同步创建联系人
		InsertEmployee(data entity.HrEmployee) (odoorpc.ResultData, error)
		// 员工用户详情
		FindOne(data entity.HrEmployee) (resData *entity.EmployeeWithUser, err error)
		// 用户个人详情
		FindProfile(data entity.ResUsers) (*entity.ResUsersPage, error)
		// 分页查询用户员工数据 findlist的分页版 查询员工为先 返回关联用户数据
		FindListPage(page, pageSize int, data entity.ResUsersPage) (*[]entity.ResUsersPage, int64, error)
		// 用户员工数据全连接详细
		FindList(data entity.ResUsers) (list *[]entity.ResUsersView, err error)
		// 更新员工
		UpdateEmloyee(data entity.HrEmployee) odoorpc.ResultData
		// 更新用户
		UpdateUsers(data entity.ResUsers) odoorpc.ResultData
		// 归档用户
		DeleteUser(IDs []int64) odoorpc.ResultData
		// 归档员工
		DeleteEmployee(IDs []int64) odoorpc.ResultData
		// 发送邮件以邀请或重置密码
		SetPwd(IDs []int64) odoorpc.ResultData
		// 员工离职
		EmployeeResign(ID int64, employee entity.HrEmployee) odoorpc.ResultData
	}

	resUsersModelImpl struct {
		table string
	}
)

var ResUsersModelDao ResUsersModel = &resUsersModelImpl{
	table: `res_users`,
}

// 登录 后续需要添加totp mfa
func (m *resUsersModelImpl) Login(u entity.LoginO) (*entity.ResUsersPage, error) {
	user := new(entity.ResUsersPage)
	err := global.HrDb.Preload("Employee").Where("login = ? ", u.Login).Find(&user).Error
	if err != nil {
		return nil, errors.New("查询用户信息失败")
	}
	b := tool.VerifyPwd(u.Password, user.Password)
	if !b {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

// 插入or新建 管理员新建用户顺带员工和联系人,同时默认为内部用户自动添加res_groups和res_users关系到res_groups_users_rel表中
// data entity.CreateUserDto
func (m *resUsersModelImpl) InsertUser(data entity.CreateUserDto) (odoorpc.ResultData, error) {
	// bytes := tool.HashPwd(data.Password) //kgo.KEncr.PasswordHash([]byte(data.Password), bcrypt.DefaultCost)
	// data.Password = string(bytes)
	// check 用户名
	var count int64
	global.HrDb.Table("hr_employee").Where("name = ?", data.Name).Count(&count)
	if count != 0 {
		return odoorpc.ResultData{Status: "fail", Code: 500, Message: "该员工已存在！"}, errors.New("该员工已存在！")
	}
	global.HrDb.Table("res_users").Where("login = ?", data.Login).Count(&count)
	if count != 0 {
		return odoorpc.ResultData{Status: "fail", Code: 500, Message: "该用户已存在！"}, errors.New("该用户已存在！")
	}
	// 通过odoorpc 实现创建
	createData, _ := json.Marshal(data)
	// fmt.Printf("===%T\n", string(data))
	var create map[string]interface{}
	_ = json.Unmarshal(createData, &create)
	createResult, err := odoorpc.Create(create, "res.users") //调用内部 res.users create方法时需要定义name以提供员工和联系人参数
	if err != nil {
		return odoorpc.ResultData{Status: "fail", Code: 500, Message: "新建操作失败"}, errors.New("新建操作失败")
	}
	return createResult, nil
}

// 插入or新建 管理员新建员工顺带联系人
func (m *resUsersModelImpl) InsertEmployee(data entity.HrEmployee) (odoorpc.ResultData, error) {
	// bytes := tool.HashPwd(data.Password) //kgo.KEncr.PasswordHash([]byte(data.Password), bcrypt.DefaultCost)
	// data.Password = string(bytes)
	// check 用户名
	var count int64
	global.HrDb.Table("hr_employee").Where("name = ?", data.Name).Count(&count)
	if count != 0 {
		return odoorpc.ResultData{Status: "fail", Code: 500, Message: "该员工已存在！"}, errors.New("该员工已存在！")
	}
	// 通过odoorpc 实现创建
	createData, _ := json.Marshal(data)
	// fmt.Printf("===%T\n", string(data))
	var create map[string]interface{}
	_ = json.Unmarshal(createData, &create)
	createResult, err := odoorpc.Create(create, "hr.employee") //调用内部 res.users create方法时需要定义name以提供员工和联系人参数
	if err != nil {
		return odoorpc.ResultData{Status: "fail", Code: 500, Message: "新建操作失败"}, errors.New("新建操作失败")
	}
	return createResult, nil
}

// 条件查询  员工详情
func (m *resUsersModelImpl) FindOne(data entity.HrEmployee) (*entity.EmployeeWithUser, error) {
	// fmt.Print(data.Name.String, data.ID)
	resData := new(entity.EmployeeWithUser)
	db := global.HrDb.Table("hr_employee").
		Joins("left join res_users on res_users.id = hr_employee.user_id").
		Joins("left join hr_department on hr_department.id = hr_employee.department_id").
		Joins("left join res_company on res_company.id = hr_employee.company_id")
	if data.ID != 0 {
		db = db.Where("hr_employee.id = ?", data.ID)
	}
	if data.Name != "" {
		db = db.Where("hr_employee.name = ?", data.Name)
	}
	if data.DepartmentId != 0 {
		db = db.Where("hr_employee.department_id = ?", data.DepartmentId)
	}
	if data.UserId != 0 {
		db = db.Where("hr_employee.user_id = ?", data.UserId)
	}
	// if data.RoleId != 0 {
	// 	db = db.Where("role_id = ?", data.RoleId)
	// }
	// if data.OrganizationId != 0 {
	// 	db = db.Where("organization_id = ?", data.OrganizationId)
	// }
	// if data.PostId != 0 {
	// 	db = db.Where("post_id = ?", data.PostId)
	// }
	err := db.Select(
		"hr_employee.*",
		"res_users.login as username",
		"res_users.active as user_active",
		"hr_department.name as department_name",
		"res_company.name as company_name").
		First(&resData).Error
	return resData, err
}

func (m *resUsersModelImpl) FindProfile(data entity.ResUsers) (*entity.ResUsersPage, error) {
	user := new(entity.ResUsersPage)
	err := global.HrDb.Preload("Employee").Where("login = ? ", data.Login).Find(&user).Error
	if err != nil {
		return nil, errors.New("查询用户信息失败")
	}
	return user, nil
}

func (m *resUsersModelImpl) FindListPage(page, pageSize int, data entity.ResUsersPage) (*[]entity.ResUsersPage, int64, error) {
	list := make([]entity.ResUsersPage, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	// db := global.HrDb.Debug()
	// tt := time.Now()
	db := global.HrDb.Preload("Employee").
		Joins("left join hr_employee on res_users.id = hr_employee.user_id")
	// Joins("left join hr_department on hr_department.id = hr_employee.department_id").
	// Joins("left join res_company on res_company.id = hr_employee.company_id")
	if data.ID != 0 {
		db = db.Where("res_users.id = ?", data.ID)
	}
	if data.Login != "" {
		db = db.Where("res_users.login like ?", "%"+data.Login+"%")
	}
	if data.Employee.WorkPhone != "" {
		db = db.Where("hr_employee.work_phone = ?", data.Employee.WorkPhone)
	}
	if data.Employee.Name != "" {
		db = db.Where("hr_employee.name like ?", "%"+data.Employee.Name+"%")
	}
	if data.Employee.DepartmentId != 0 {
		db = db.Where("hr_employee.department_id = ?", data.Employee.DepartmentId)
	}
	// if data.Active {
	// 	db = db.Where("res_users.active = ?", data.Active)
	// }
	// if data.RoleId != 0 {
	// 	db = db.Where("role_id = ?", data.RoleId)
	// }
	// if data.PostId != 0 {
	// 	db = db.Where("post_id = ?", data.PostId)
	// }
	// tt = time.Now()
	err := global.HrDb.Table("res_users").Count(&total).Error
	err = db.
		Where("res_users.active = ?", data.Active).
		// Select(
		// 	"hr_employee.*",
		// 	"res_users.login as username",
		// 	"res_users.active as user_active",
		// 	"hr_department.name as department_name",
		// 	"res_company.name as company_name").
		Limit(pageSize).Offset(offset).Find(&list).Error

	return &list, total, err
}

// 用户列表+员工数据
func (m *resUsersModelImpl) FindList(data entity.ResUsers) (*[]entity.ResUsersView, error) {
	list := make([]entity.ResUsersView, 0)
	// 预加载员工数据 主要体现员工数据 用户数据只需要编号以及登录名
	db := global.HrDb.Preload("Employee")
	if data.ID != 0 {
		db = db.Where("id = ?", data.ID)
	}
	if data.Login != "" {
		db = db.Where("login = ?", data.Login)
	}
	db.Where("active = ?", data.Active)
	//存疑 先测试
	// if data.HrEmployee.Name.String != "" {
	// 	db = db.Where("hr_employee.name = ?", data.HrEmployee.Name.String)
	// }

	err := db.Find(&list).Error
	return &list, err
}

// 更新员工数据
func (m *resUsersModelImpl) UpdateEmloyee(data entity.HrEmployee) odoorpc.ResultData {
	dd, _ := json.Marshal(data)
	var writeuData map[string]interface{}
	_ = json.Unmarshal(dd, &writeuData)
	writeuResult, err := odoorpc.Write(writeuData, "hr.employee", data.ID) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	if err != nil {
		return odoorpc.ResultData{Status: "fail", Code: 500, Message: "修改失败！"}
	}
	return writeuResult
}

// 更新用户数据
func (m *resUsersModelImpl) UpdateUsers(data entity.ResUsers) odoorpc.ResultData {
	//如果传入了bool字段需要更改为false，直接操作数据库更加方便
	if !data.Active {
		result := global.HrDb.Model(&entity.ResUsers{}).Where("id = ?", data.ID).Update("active", false)
		if result.Error != nil {
			return odoorpc.ResultData{Status: "fail", Code: 500, Message: "归档失败！"}
		}
	}
	if !data.TourEnabled {
		result := global.HrDb.Model(&entity.ResUsers{}).Where("id = ?", data.ID).Update("tour_enabled", false)
		if result.Error != nil {
			return odoorpc.ResultData{Status: "fail", Code: 500, Message: "归档失败！"}
		}
	}
	if !data.OdoobotFailed {
		result := global.HrDb.Model(&entity.ResUsers{}).Where("id = ?", data.ID).Update("odoobot_failed", false)
		if result.Error != nil {
			return odoorpc.ResultData{Status: "fail", Code: 500, Message: "归档失败！"}
		}
	}
	//调用odoorpc
	dd, _ := json.Marshal(data)
	var writeuData map[string]interface{}
	_ = json.Unmarshal(dd, &writeuData)
	writeuResult, err := odoorpc.Write(writeuData, "res.users", data.ID) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	if err != nil {
		return odoorpc.ResultData{Status: "fail", Code: 500, Message: "修改失败！"}
	}

	return writeuResult
}

// 归档用户
func (m *resUsersModelImpl) DeleteUser(IDs []int64) odoorpc.ResultData {
	// 直接调用odoorpc归档
	// data := entity.ResUsers{
	// 	ResUsersB: entity.ResUsersB{
	// 		Active: false,
	// 	},
	// }
	// var writeuResult odoorpc.ResultData
	// var err error
	// dd, _ := json.Marshal(data)
	// var writeuData map[string]interface{}
	// _ = json.Unmarshal(dd, &writeuData)
	// var i int
	// for i = 0; i < len(IDs); i++ {
	// 	writeuResult, err = odoorpc.Write(writeuData, "res.users", IDs[i]) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	// 	if err != nil {
	// 		return odoorpc.ResultData{Status: "fail", Code: 500, Message: "归档失败！"}
	// 	}

	// }
	// return writeuResult

	// 直接数据库操作归档
	var i int
	for i = 0; i < len(IDs); i++ {
		result := global.HrDb.Model(&entity.ResUsers{}).Where("id = ?", IDs[i]).Update("active", false)
		if result.Error != nil {
			return odoorpc.ResultData{Status: "fail", Code: 500, Message: "归档失败！"}
		}

	}
	return odoorpc.ResultData{Status: "success", Code: 200, Message: "成功归档！"}
}

// 归档员工
func (m *resUsersModelImpl) DeleteEmployee(IDs []int64) odoorpc.ResultData {
	var i int
	for i = 0; i < len(IDs); i++ {
		result := global.HrDb.Model(&entity.HrEmployee{}).Where("id = ?", IDs[i]).Update("active", false)
		if result.Error != nil {
			return odoorpc.ResultData{Status: "fail", Code: 500, Message: "归档失败！"}
		}
	}
	return odoorpc.ResultData{Status: "success", Code: 200, Message: "成功归档！"}
}

// 重置密码
func (m *resUsersModelImpl) SetPwd(IDs []int64) odoorpc.ResultData {
	var rpResult odoorpc.ResultData
	var err error
	var i int
	for i = 0; i < len(IDs); i++ {
		// fmt.Print(IDs[i])
		rpResult, err = odoorpc.ResetPassword(IDs[i]) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
		if err != nil {
			return odoorpc.ResultData{Status: "fail", Code: 500, Message: "发送邀请文件！"}
		}

	}

	return rpResult
}

// 员工离职
func (m *resUsersModelImpl) EmployeeResign(ID int64, employee entity.HrEmployee) odoorpc.ResultData {
	dd, _ := json.Marshal(employee)
	var writeuData map[string]interface{}
	_ = json.Unmarshal(dd, &writeuData)
	fmt.Printf("员工离职传入参数：%+v\n", writeuData)
	writeuResult, err := odoorpc.Write(writeuData, "hr.employee", employee.ID) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	if err != nil {
		return odoorpc.ResultData{Status: "fail", Code: 500, Message: "修改失败！"}
	}
	m.DeleteEmployee([]int64{employee.ID})
	if employee.UserId != 0 {
		m.DeleteEmployee([]int64{employee.UserId})
	}
	return writeuResult
}
