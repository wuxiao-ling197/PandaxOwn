package services

import (
	"fmt"
	"log"

	// "pandax/apps/develop/entity"
	"pandax/apps/shared/entity"
	"pandax/pkg/global"
	"pandax/pkg/tool"
	"testing"
)

func TestService(t *testing.T) {
	// model := &resUsersModelImpl{} //创建resUsersModelImpl实例，作为调用接口的接收者
	// 1、测试Login cg
	// var lo = new(entity.LoginO)
	// lo.Login = "admin"
	// lo.Password = "admin"
	// u, _ := model.Login(*lo)
	// fmt.Printf("Login = %+v", u)

	// 2、测试FindList cg
	// data := new(entity.ResUsers)
	// err := global.HrDb.Table("res_users").Where("login = ? ", "测测结构体").Find(&data).Error
	// if err != nil {
	// 	log.Fatalf("Error querying users: %+v", err)
	// }
	// fmt.Printf("%T", *data) //打印参数类型
	// r, _ := model.FindList(*data)
	// fmt.Printf("FindList = %+v", r)

	// 3、测试FindOne cg
	// data := new(entity.HrEmployee)
	// err := global.HrDb.Table("hr_employee").Where("name = ? ", "Administrator").Find(&data).Error
	// if err != nil {
	// 	log.Fatalf("Error querying users: %+v", err)
	// }
	// o, _ := model.FindOne(*data)
	// fmt.Printf("FindOne = %+v", o)

	// 4、测试 Insert cg
	// data := entity.CreateUserDto{
	// 	Login:     "测测结构体",
	// 	Name:      "测测结构体",
	// 	CompanyId: 1,
	// 	Email:     "lindsay@tbird.com",
	// }
	// i, _ := model.Insert(data)
	// fmt.Printf("%+v", i)

	// 5、测试 Update cg
	// data := entity.HrEmployee{
	// 	ID:           36,
	// 	Name:         "测测修改",
	// 	DepartmentId: 2,
	// 	WorkEmail:    "dr60@tbird.com",
	// }
	// i := model.Update(data)
	// fmt.Printf("%+v", i)

	// 6、测试 归档用户 cg
	// data := []int64{32, 33, 37}
	// i := model.Delete(data)
	// fmt.Printf("%+v", i)

	// 7、测试 重置密码 cg
	// data := []int64{37, 42}
	// i := model.SetPwd(data)
	// fmt.Printf("%+v", i)

	// 8、测试归档员工 cg
	// data := []int64{30, 31, 35}
	// i := model.DeleteEmployee(data)
	// fmt.Printf("%+v", i)

	// 9、测试 分页列表 cg
	// r, n, _ := model.FindListPage(1, 15, entity.ResUsers{})
	// fmt.Printf("FindListPage = %+v\n,total=%v", r, n)

	hrmodel := &hrDepartmentModelImpl{}
	// 10、部门测试 FindList cg
	// result, _ := hrmodel.FindList(entity.HrDepartment{})
	// fmt.Printf("FindList=%+v\n", result)

	// 11、部门测试 FindOne cg
	// result, _ := hrmodel.FindOne(10)
	// fmt.Printf("FindOne=%+v\n", result)

	// 12、部门测试 Delete
	// data := []int64{10, 2}
	// result := hrmodel.Delete(data)
	// fmt.Printf("Delete=%+v\n", result)

	// 13、部门测试 Update cg
	// data := entity.HrDepartment{
	// 	ID:           2,
	// 	CompleteName: "研究||发展",
	// }
	// result := hrmodel.Update(data)
	// fmt.Printf("Update=%+v\n", result)

	// 13、部门测试 Insert cg
	// data := entity.HrDepartment{
	// 	Name: "后端测试",
	// }
	// result, _ := hrmodel.Insert(data)
	// fmt.Printf("Update=%+v\n", result)

	// 14、部门测试 FindListPage
	// r, n, _ := hrmodel.FindListPage(2, 4, entity.HrDepartment{})
	// fmt.Printf("FindListPage = %+v\n,total=%v", r, n)

	// 15、部门测试 组织书树
	list, err := hrmodel.FindList(entity.HrDepartment{})
	if err != nil {
		fmt.Print("查询失败\n")
	}
	// fmt.Printf("通过方法获取的数据：%T\n", *list)

	// sd := make([]entity.HrDepartment, 0)
	li := *list //[]entity.CompanyWithDepatment
	for j := 0; j < len(li); j++ {
		lip := li[j]
		// fmt.Printf("取出结果中的ResCompanyB数据：%T\n", lip.ResCompanyB) //entity.ResCompanyB
		// fmt.Printf("取出结果中的Departments数据：%T\n", lip.Departments) //[]entity.HrDepartment
		// fmt.Printf("谁的数据？：%+v\n", li[j]) //公司数据
		for i := 0; i < len(lip.Departments); i++ {
			if li[i].ParentId != 0 {
				continue
			}
			// info := DiguiA(list, li[i])

			// sd = append(sd, info)
		}
	}
}

func TestPassword(t *testing.T) {
	// login, _ := odoorpc.Authenticate()
	// fmt.Printf("odoo登录端口：%+v", login)
	// // t.Logf("%v \n%v", global.HrDb.ConnPool, global.Db.ConnPool)
	// var users []entity.ResUsers        // 定义切片存放查询结果
	// result := global.HrDb.Find(&users) //接收查询结果而非内容
	// // 检查错误
	// if result.Error != nil {
	// 	log.Fatalf("Error querying users: %v", result.Error)
	// }
	// fmt.Printf("查询成功")
	// for _, user := range users {
	// 	log.Printf("Wanted User: %+v\n", user)
	// }

	pwd := tool.HashPwd("odoo18")
	log.Print(pwd)
	rr := tool.VerifyPwd("odoo18", "$pbkdf2-sha512$600000$xvj/n9M659y7t7YWYizl/A$2YVgLexC1dsO1tTou90r13u7KBbHNsMEcIuP27Syy2FWKMqVVXrDddt/ZaalrBDRFOGAp2Pu2LMGeutvvByAgg")
	fmt.Printf("%v", rr)
}

func TestOdooList(t *testing.T) {
	// 只返回指定字符 cg
	// user := new(entity.ResUsersId)
	// er := global.HrDb.Table("res_users").Where("login = ?", "create").First(&user).Error
	// fmt.Printf("查询结果：%+v\n, %+v", user.ID, er)

	////简单查询 只查询了单个表中对象
	// user := new(entity.ResUsers)
	// err := global.HrDb.Table("res_users").Where("login = ? ", "odoo18").Find(user).Error
	// if err != nil {
	// 	log.Fatalf("Error querying users: %+v", err)
	// }
	// // log.Printf("登录用户Login User= %+v", user)
	// var users []entity.ResUsersView
	// result := global.HrDb.Find(&users)
	// if result == nil {
	// 	log.Fatalf("Error querying users: %+v", result.Error)
	// }
	// log.Printf("获取用户：%+v", users)
	// for _, user := range users {
	// 	log.Printf("Odoo User: %+v\n", user)
	// }

	////复杂查询，1对1，1对多，多对多查询，联合查询，条件查询
	// 1.cg 输出完整的用户+员工数据
	// employee := make([]entity.ResUsersView, 0)
	// err := global.HrDb.Preload("HrEmployee").Find(&employee).Error
	// if err != nil {
	// 	t.Fatalf("Error querying users: %+v", err)
	// }
	// fmt.Printf("获取员工用户：%+v", len(employee))

	//2.cg  只有预加载的时候才会存储employee的数据 而且这里new的模型实体必须在该结构体对应的模型实体中有外键指向 preload的是ResUsersPage中定义的对象
	// // 输出的是没有关联到员工的用户
	// view := make([]entity.ResUsersPage, 0)
	// wrong := global.HrDb.Table("res_users").Joins("left join hr_employee on hr_employee.user_id = res_users.id").Scan(&view).Error
	// if wrong != nil {
	// 	t.Fatalf("Error querying users: %+v", wrong)
	// }
	// fmt.Printf("获取员工用户（主员工数据）：%+v", len(view))

	//3.cg
	// have := make([]entity.ResUsers, 0)
	// //输出的是没有关联到员工的用户
	// ww := global.HrDb.Preload("Employee").Scan(&have).Error //Table("res_users").Joins("left join hr_employee on hr_employee.user_id = res_users.id")
	// if ww != nil {
	// 	t.Fatalf("Error querying users: %+v", ww)
	// }
	// fmt.Printf("获取员工部门：%+v", len(have))

	//4.条件查询 cg
	// eu := new(entity.ResUsersView)
	// cx := global.HrDb.Table("res_users").Joins("left join hr_employee on hr_employee.name = ?", "Administrator").First(&eu).Error //cg Preload("HrEmployee").where("login = ?", "xxx")也可以实现目标
	// cx := global.HrDb.Table("res_users").
	// Joins("left join hr_employee as employee on employee.user_id = res_users.id").
	// Where("employee.name = ?", "Administrator").
	// First(&eu).Error //cg
	// if cx != nil {
	// 	t.Fatalf("Error querying users: %+v", cx)
	// }
	// fmt.Printf("查询指定员工用户：%+v", eu)
	// 动态映射，上面都是按结构体返回
	// var ue []map[string]interface{} //Emma Granger  on employee.user_id = res_users.id
	// ue := new(entity.EmployeeWithUser)
	// cx := global.HrDb.Table("hr_employee").
	// 	Joins("left join res_users on res_users.id = hr_employee.user_id").
	// 	Joins("left join hr_department on hr_department.id = hr_employee.department_id").
	// 	Joins("left join res_company on res_company.id = hr_employee.company_id").
	// 	Where("hr_employee.name = ?", "Administrator").
	// 	Select(
	// 				"hr_employee.*",
	// 				"res_users.login as username",
	// 				"res_users.active as user_active",
	// 				"hr_department.name as department_name",
	// 				"res_company.name as company_name").
	// 	First(&ue).Error //Preload("HrEmployee").where("login = ?", "xxx")也可以实现目标
	// if cx != nil {
	// 	t.Fatalf("Error querying users: %+v", cx)
	// }
	// fmt.Printf("查询指定员工用户：%+v", ue)

	// 5.修改单项数据 cg
	// result := global.HrDb.Model(&entity.ResUsers{}).Where("id = ?", 37).Update("active", false)
	// fmt.Print(result)

	// 6.修改多项数据
	// params := map[string]interface{}{
	// 	"login":         "多项数据修改",
	// 	"pandax_secret": "false",
	// }
	// result := global.HrDb.Model(&entity.ResUsers{}).Where("id = ?", 37).Updates(params)
	result := global.HrDb.Model(&entity.ResUsers{}).Where("id = ?", 37).Updates(entity.ResUsersB{
		PandaxSecret: "",
	})
	fmt.Print(result.Error)
}

func TestDepCom(t *testing.T) {
	global.HrDb = global.HrDb.Debug() // 启用调试模式
	// 查询公司
	// var com = new(entity.ResCompany)
	// ress := global.HrDb.Preload("Departments").
	// 	// Table("res_company").
	// 	// Select(
	// 	// 	"res_company.id",
	// 	// 	"res_company.name",
	// 	// 	"res_company.parent_id",
	// 	// 	"res_company.email",
	// 	// 	"res_company.phone",
	// 	// 	"res_company.active").
	// 	Find(&com).Error
	// if ress == nil {
	// 	t.Fatalf("公司信息查询失败: %+v", ress)
	// }
	// fmt.Printf("公司信息：%+v\n", com)

	//cg 查询主要公司信息及所属部门信息 通过主表查外键表
	var comdep = make([]entity.CompanyWithDepatment, 0)
	// params := "管理"
	res := global.HrDb.Preload("Departments").
		Joins("left join hr_department on hr_department.company_id = res_company.id").
		// Where("hr_department.complete_name like ?", "%"+params+"%").
		Order("res_company.id").
		Select("DISTINCT res_company.*").
		Find(&comdep).Error
	if res != nil {
		t.Fatalf("常见公司信息查询失败: %+v", res)
	}
	fmt.Printf("常见公司信息->部门：%+v\n", comdep)

	//cg 查询公司及所属部门
	// var dep = new(entity.ResCompany)
	// result := global.HrDb.Preload("Departments").Find(&dep).Error
	// if result == nil {
	// 	t.Fatalf("公司->部门查询失败: %+v", result)
	// }
	// fmt.Printf("公司->部门：%+v", dep) //只有公司数据 没有部门数据

	// 查询部门以及上级公司
	// var partmen = new(entity.HrDepartment)
	// jieguo := global.HrDb.Preload("Employees").Find(&partmen).Error
	// if jieguo != nil {
	// 	t.Fatalf("部门->+公司查询失败: %+v", jieguo)
	// }
	// fmt.Printf("部门->+公司：%+v\n", partmen)

	//cg 查询部门及管理员工
	// var departmentWithEmployee = make([]entity.HrDepartment, 0)
	// data := global.HrDb.Preload("Employees").Find(&departmentWithEmployee).Error
	// if data != nil {
	// 	t.Fatalf("部门->员工查询失败: %+v", data)
	// }
	// fmt.Printf("部门->员工：%+v\n", departmentWithEmployee) //nil

}

// func DiguiA(fulldata *[]entity.CompanyWithDepatment, menu entity.HrDepartment, company entity.ResCompany) entity.HrDepartment {
// 	list := *fulldata
// 	// 创建接收？
// 	min := make([]entity.HrDepartment, 0)
// 	// 第一次 遍历公司
// 	for j := 0; j < len(list); j++ {
// 		//公司从属
// 		if menu.ID != list[j].ParentId {
// 			continue
// 		}
// 		// 创建模型
// 		mi := entity.HrDepartment{}
// 		// list[j] = 单条公司+所属部门数据
// 		mi.ID = list[j].ID //属于公司ID ResCompanyB中的数据可以直接访问
// 		mi.ParentId = list[j].ParentId
// 		// mi.ParentPath = list[j].Departments.ParentPath
// 		mi.Name = list[j].Name
// 		// mi.ManagerId = list[j].ManagerId
// 		// mi.Property = list[j].Property
// 		// mi.DutyId = list[j].DutyId
// 		// mi.Virtual = list[j].Virtual
// 		// mi.Code =list[j].Code
// 		mi.Active = list[j].Active
// 		// mi.CreateDate = list[j].CreateDate
// 		// mi.WriteDate = list[j].WriteDate
// 		mi.Children = []entity.HrDepartment{}
// 		ms := DiguiA(fulldata, mi)
// 		min = append(min, ms)
// 	}
// 	menu.Children = min
// 	return menu
// }
