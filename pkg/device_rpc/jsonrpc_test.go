package devicerpc

import (

	// dentity "pandax/apps/develop/entity"

	"encoding/json"
	"fmt"
	"pandax/apps/shared/entity"
	"testing"
)

func TestRpc(t *testing.T) {
	// 验证
	// userID, err := Authenticate()
	// if err != nil {
	// 	fmt.Printf("登录验证失败：%v", err)
	// } else {
	// 	fmt.Printf("获取的用户 ID 是：%v\n", userID.Data)
	// }

	// odoorpc中没有提供列表查询
	// users, err := AllUser()
	// if err != nil {
	// 	fmt.Errorf("获取用户失败: %v", err)
	// } else {
	// 	fmt.Printf("查询记录成功: %v", users.Data)
	// }

	// 条件查询1
	// users, err := FindOne(12)
	// if err != nil {
	// 	fmt.Printf("获取用户失败: %v", err)
	// } else {
	// 	fmt.Printf("AAA查询记录成功: %+v", users.Data)
	// }

	// // 条件查询2
	// domain := [][]interface{}{
	// 	{"active", "=", true},
	// }
	// result, err := Search(domain)
	// if err != nil {
	// 	fmt.Printf("查询用户失败: %v", err)
	// }
	// fmt.Printf("BBB查询用户成功: %+v", result)

	// 定义object 对象
	// createData := map[string]interface{}{
	// 	"login":             "210",
	// 	"password":          "odoo18",
	// 	"notification_type": "email",
	// 	"company_id":        1,
	// 	"partner_id":        10,
	// 	// 可以添加其他字段
	// }
	// createResult, err := OnChange(createData, "res.users") //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	// if err != nil {
	// 	t.Errorf("创建记录失败: %v", err)
	// } else {
	// 	t.Logf("创建记录成功: %v", createResult)
	// }

	//对象创建用户 同步创建员工和联系人
	// createData := map[string]interface{}{
	// 	// "sequence":   10,
	// 	"company_id": 1,
	// 	// "contract_type_id":  4,
	// 	// "no_of_recruitment": 10,
	// 	"login": "object_test",
	// 	"name":  "object_test",
	// 	// "department_id": 1,
	// 	// "partner_id": 23,
	// }
	// fmt.Print("对象传参\n")
	// result, _ := Create(createData, "res.users")
	// fmt.Printf("标准传参= %+v\n 结果：%+v\n", createData, result)

	// 创建部门
	createData := entity.HrDepartment{
		// "sequence":   10,
		CompanyId: 1,
		// "contract_type_id":  4,
		// "no_of_recruitment": 10,
		ParentId: 2,
		Name:     "测试",
		// Name:"{"en_US": "子公司管理部", "zh_CN": "子公司管理部"}",
		CompleteName: "测试",
		// "department_id": 1,
		// "partner_id": 23,
	}
	var create map[string]interface{}
	data, _ := json.Marshal(createData)
	_ = json.Unmarshal(data, &create)
	result, _ := Create(create, "hr.department")
	fmt.Printf("标准传参= %+v\n 结果：%+v\n", create, result)

	// cg 定义了用户创建时的变量
	// emp := entity.CreateUserDto{
	// 	Login:     "测测结构体传参",
	// 	Name:      "测测结构体传参",
	// 	CompanyId: 1,
	// 	Email:     "lindsay@tbird.com",
	// }
	// data, _ := json.Marshal(emp)
	// var create map[string]interface{}
	// _ = json.Unmarshal(data, &create)
	// // fmt.Printf("类型：%T\n,数据：%+v\n", create, create)
	// createResult, err := Create(create, "res.users") // 调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	// if err != nil {
	// 	t.Errorf("创建记录失败: %v", err)
	// }
	// fmt.Printf("创建记录成功: %+v", createResult)

	// 修改
	// writeData := map[string]interface{}{
	// 	"name": "天天开心",
	// }
	// writeResult, err := Write(writeData, "hr.employee", 11) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	// if err != nil {
	// 	t.Errorf("修改记录失败: %v", err)
	// } else {
	// 	t.Logf("修改记录成功: %+v", writeResult)
	// }

	// var uID int64
	// data := [][]interface{}{
	// 	{"login", "=", "测测结构体"}, //like
	// }
	// result, _ := Search(data, "res.users")
	// fmt.Printf("查询%T\n", result.Data)
	// if value, ok := result.Data.(float64); ok {
	// 	fmt.Printf("float64: %v\n", value)
	// } else {
	// 	fmt.Println("not an int")
	// }
	// for value, uid := range result.Data.([]interface{}) {
	// 	fmt.Printf("%T, %+v\n", uid, uid)
	// 	fmt.Print(value)
	// }
	// writeuData := map[string]interface{}{
	// 	"": "天天开心",
	// }

	// 修改用户
	// write := entity.ResUsers{
	// 	ResUsersB: entity.ResUsersB{
	// 		// Active:      true,
	// 		// Share:       false,
	// 		// TourEnabled: false,
	// 		Karma: "666689",
	// 	},
	// }
	// dd, _ := json.Marshal(write)
	// var writeuData map[string]interface{}
	// _ = json.Unmarshal(dd, &writeuData)
	// fmt.Println(writeuData)
	// writeuResult, err := Write(writeuData, "res.users", 37) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	// if err != nil {
	// 	t.Errorf("修改记录失败: %v", err)
	// } else {
	// 	t.Logf("修改记录成功: %+v", writeuResult)
	// }

	// 修改员工
	// write := entity.HrEmployee{
	// 	DepartmentId: 0, //让该字段赋值为null
	// 	Active:       false,
	// 	// UserId:       0,
	// 	JobTitle: "医务秘书", //删除保存的内容
	// }
	// dd, _ := json.Marshal(write)
	// var writeuData map[string]interface{}
	// _ = json.Unmarshal(dd, &writeuData)
	// writeuResult, err := Write(writeuData, "hr.employee", 36) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	// if err != nil {
	// 	t.Errorf("修改记录失败: %v", err)
	// } else {
	// 	t.Logf("修改记录成功: %+v", writeuResult)
	// }

	// 修改部门
	// write := entity.HrDepartment{
	// 	Active:       true,
	// 	CompleteName: "研究和发展",
	// }
	// dd, _ := json.Marshal(write)
	// var writeuData map[string]interface{}
	// _ = json.Unmarshal(dd, &writeuData)
	// fmt.Println(writeuData)
	// writeuResult, err := Write(writeuData, "hr.department", 2) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	// if err != nil {
	// 	t.Errorf("修改记录失败: %v", err)
	// } else {
	// 	t.Logf("修改记录成功: %+v", writeuResult)
	// }

	//为员工创建关联用户
	// cru, err := CreateRalativeUser(15)
	// if err != nil {
	// 	t.Errorf("重置密码失败: %v", err)
	// } else {
	// 	t.Logf("创建关联用户成功: %+v", cru)
	// }

	// 重置密码
	// data := []int{37, 42}
	// rpResult, err := ResetPassword(37) //调用create rpc创建时同时创建了employee，删除时存在外键约束
	// if err != nil {
	// 	t.Errorf("重置密码失败: %v", err)
	// } else {
	// 	t.Logf("重置密码成功: %+v", rpResult)
	// }

	// 删除
	/**  删除员工不会删除关联的用户和联系人
	删除用户不会删除关联的联系人
	*/
	// deleteResult, err := Delete(37, "res.users") //调用create rpc创建时同时创建了employee，删除时存在外键约束
	// if err != nil {
	// 	t.Errorf("删除记录失败: %v", err)
	// } else {
	// 	t.Logf("删除记录成功: %+v", deleteResult)
	// }
}
