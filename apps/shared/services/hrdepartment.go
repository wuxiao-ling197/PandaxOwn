package services

import (
	"encoding/json"
	"fmt"
	"pandax/apps/shared/entity"
	odoorpc "pandax/pkg/device_rpc"
	"pandax/pkg/global"
)

type (
	HrDepartmentModel interface {
		// 新建部门 注意只能在所属公司下操作本公司的下属部门
		Insert(data entity.HrDepartment) (odoorpc.ResultData, error)
		// 查询部门详细数据 先加载公司下部门数据再公司数据
		FindOne(Id int64) (*entity.CompanyWithDepatment, error)
		// 查询指定公司 携带员工数据
		FindEmployee(Id int64) (*entity.HrDepartment, error)
		// 分页查询 携带员工数据
		FindListPage(page, pageSize int, data entity.HrDepartment) (*[]entity.HrDepartment, int64, error)
		// 常见公司信息+部门数据
		FindList(data entity.HrDepartment) (*[]entity.ResCompanyB, error)
		// 只加载部门数据
		FindListDepartment(data entity.HrDepartment) (*[]entity.HrDepartment, error)
		// 修改部门数据
		Update(data entity.HrDepartment) odoorpc.ResultData
		// 归档部门数据 部门中原有员工需要处理
		Delete(users []int64, Ids []int64, transferDepId int64) odoorpc.ResultData
		// 部门归档 没有所属员工
		DeleteNoEmployee(Ids []int64) odoorpc.ResultData
		//
		SelectOrganization(data entity.HrDepartment) ([]entity.ResCompanyB, error)
		//
		SelectOrganizationLable(data entity.HrDepartment) ([]entity.DepartmentLable, error)
		//
		SelectOrganizationIds(data entity.HrDepartment) ([]int64, error)
	}

	hrDepartmentModelImpl struct {
		table string
	}
)

var HrDepartmentModelDao HrDepartmentModel = &hrDepartmentModelImpl{
	table: `hr_department`,
}

/**注意操作用户的权限 只能操作本公司下属的部门新建等操作 否则会超时失败*/
func (m *hrDepartmentModelImpl) Insert(data entity.HrDepartment) (odoorpc.ResultData, error) {
	// if data.CompanyId !=
	var create map[string]interface{}
	createData, _ := json.Marshal(data)
	_ = json.Unmarshal(createData, &create)
	result, err := odoorpc.Create(create, "hr.department")
	return result, err
}

func (m *hrDepartmentModelImpl) FindOne(Id int64) (*entity.CompanyWithDepatment, error) {
	resData := new(entity.CompanyWithDepatment)
	err := global.HrDb.Preload("Departments").
		Joins("left join hr_department on res_company.id = hr_department.company_id").
		Where("id = ?", Id).First(&resData).Error
	return resData, err
}

func (m *hrDepartmentModelImpl) FindEmployee(Id int64) (*entity.HrDepartment, error) {
	resData := new(entity.HrDepartment)
	err := global.HrDb.Preload("Employees").
		Where("id = ?", Id).First(&resData).Error
	return resData, err
}

func (m *hrDepartmentModelImpl) FindListPage(page, pageSize int, data entity.HrDepartment) (*[]entity.HrDepartment, int64, error) {
	list := make([]entity.HrDepartment, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.HrDb.Preload("Employees")
	if data.ID != 0 {
		db = db.Where("hr_department.id = ?", data.ID)
	}
	if data.Name != "" {
		db = db.Where("hr_department.name like ?", "%"+data.Name+"%")
	}
	if data.CompleteName != "" {
		db = db.Where("hr_department.complete_name like ?", "%"+data.CompleteName+"%")
	}
	if data.ParentId != 0 {
		db = db.Where("hr_department.parent_id = ?", data.ParentId)
	}
	if data.ManagerId != 0 {
		db = db.Where("hr_department.manager_id = ?", data.ManagerId)
	}
	err := global.HrDb.Table("hr_department").Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

// 部门详细数据+常见公司数据 公司->部门
/**返回结果：
FindList=&[
    // 公司
    {ID:1 Name:总部 ParentId:0 ParentPath:1/ Email:wlzyl4136@qq.com Phone:+1 555-555-5556 Active:true
        // 公司的子公司
        Children:[
                {ID:4 Name:成都分区 ParentId:1 ParentPath:1/4/ Email: Phone: Active:true
                    Children:[]
                    Departments:[]}
                {ID:2 Name:一级子公司 ParentId:1 ParentPath:2/ Email: Phone: Active:true
                    Children:[]
                    Departments:[]}]
        // 部门
        Departments:[
            {ID:2 CompanyId:1 ParentId:0 ManagerId:1 Color:10 MasterDepartmentId:2 CreateUid:2 WriteUid:2 CompleteName:研究||发展 ParentPath:2/ Name:{"en_US": "Research & Development", "zh_CN": "研究和发展"} Note: Active:true CreateDate:2024-10-09T06:26:32.393177Z WriteDate:2025-03-21T05:52:35.554288Z
                Employees:[]
                // 部门的子部门
                Children:[
                    {ID:22 CompanyId:1 ParentId:2 ManagerId:0 Color:0 MasterDepartmentId:2 CreateUid:2 WriteUid:2 CompleteName:研究||发展 / 后端测试 ParentPath:2/22/ Name:{"en_US": "后端测试"} Note: Active:true CreateDate:2025-03-19T07:00:46.466526Z WriteDate:2025-03-27T05:29:31.240497Z
                        Employees:[]
                        Children:[]}]}
            {ID:22 CompanyId:1 ParentId:2 ManagerId:0 Color:0 MasterDepartmentId:2 CreateUid:2 WriteUid:2 CompleteName:研究||发展 / 后端测试 ParentPath:2/22/ Name:{"en_US": "后端测试"} Note: Active:true CreateDate:2025-03-19T07:00:46.466526Z WriteDate:2025-03-27T05:29:31.240497Z
                Employees:[]
                Children:[]}
            {ID:3 CompanyId:1 ParentId:1 ManagerId:2 Color:5 MasterDepartmentId:1 CreateUid:2 WriteUid:2 CompleteName:系统管理 / 人事资源 ParentPath:1/3/ Name:{"en_US": "Management", "zh_CN": "人事资源"} Note: Active:true CreateDate:2024-10-09T06:26:32.393177Z WriteDate:2025-03-27T05:30:30.500677Z
                Employees:[]
                Children:[]}
            {ID:1 ParentId:0 ManagerId:1 Color:3 MasterDepartmentId:1 CreateUid:1 WriteUid:2 CompleteName:系统管理 ParentPath:1/ Name:{"en_US": "系统管理", "zh_CN": "系统管理"} Note: Active:true CreateDate:2024-10-09T06:10:50.623733Z WriteDate:2025-03-27T05:30:52.98679Z
                Employees:[]
                Children:[
                    {ID:3 CompanyId:1 ParentId:1 ManagerId:2 Color:5 MasterDepartmentId:1 CreateUid:2 WriteUid:2 CompleteName:系统管理 / 人事资源 ParentPath:1/3/ Name:{"en_US": "Management", "zh_CN": "人事资源"} Note: Active:true CreateDate:2024-10-09T06:26:32.393177Z WriteDate:2025-03-27T05:30:30.500677Z
                        Employees:[]
                        Children:[]}
                    {ID:23 CompanyId:1 ParentId:1 ManagerId:0 Color:9 MasterDepartmentId:1 CreateUid:2 WriteUid:2 CompleteName:系统管理 / 统统 ParentPath:1/23/ Name:{"en_US": "统统", "zh_CN": "统统"} Note: Active:true CreateDate:2025-03-19T07:05:29.840439Z WriteDate:2025-03-27T05:31:02.692637Z
                        Employees:[]
                        Children:[]}]}
            {ID:23 CompanyId:1 ParentId:1 ManagerId:0 Color:9 MasterDepartmentId:1 CreateUid:2 WriteUid:2 CompleteName:系统管理 / 统统 ParentPath:1/23/ Name:{"en_US": "统统", "zh_CN": "统统"} Note: Active:true CreateDate:2025-03-19T07:05:29.840439Z WriteDate:2025-03-27T05:31:02.692637Z
                Employees:[]
                Children:[]}]}
    {ID:2 Name:一级子公司 ParentId:1 ParentPath:2/ Email: Phone: Active:true
        Children:[]
        Departments:[
            {ID:10 CompanyId:2 ParentId:0 ManagerId:0 Color:0 MasterDepartmentId:10 CreateUid:2 WriteUid:2 CompleteName:子公司管理部 ParentPath:10/ Name:{"en_US": "子公司管理部", "zh_CN": "子公司管理部"} Note: Active:true CreateDate:2025-03-19T05:18:40.161385Z WriteDate:2025-03-19T05:18:40.161385Z
                Employees:[]
                Children:[]}]}
    {ID:3 Name:香港分部 ParentId:0 ParentPath:3/ Email: Phone: Active:true
        Children:[]
        Departments:[
            {ID:25 CompanyId:3 ParentId:0 ManagerId:0 Color:0 MasterDepartmentId:25 CreateUid:2 WriteUid:2 CompleteName:独立工作室 ParentPath:25/ Name:{"en_US": "独立工作室", "zh_CN": "独立工作室"} Note: Active:true CreateDate:2025-03-19T09:23:21.070181Z WriteDate:2025-03-19T09:23:21.070181Z
                Employees:[]
                Children:[]}]}
    {ID:4 Name:成都分区 ParentId:1 ParentPath:1/4/ Email: Phone: Active:true
        Children:[]
        Departments:[]}]
*/
func (m *hrDepartmentModelImpl) FindList(data entity.HrDepartment) (*[]entity.ResCompanyB, error) {
	list := make([]entity.ResCompanyB, 0)
	db := global.HrDb.Preload("Departments").Preload("Children").Preload("Departments.Children").
		Joins("left join hr_department on res_company.id = hr_department.company_id")
	// 此处填写 where参数判断
	if data.ID != 0 {
		db = db.Where("hr_department.id = ?", data.ID)
	}
	if data.CompleteName != "" {
		db = db.Where("hr_department.complete_name like ?", "%"+data.CompleteName+"%")
	}
	// if data.Active != "" {
	// 	db = db.Where("status = ?", data.Status)
	// }
	err := db.Order("res_company.id").Select("DISTINCT res_company.*").Find(&list).Error
	return &list, err
}

// 查询所有部门
func (m *hrDepartmentModelImpl) FindListDepartment(data entity.HrDepartment) (*[]entity.HrDepartment, error) {
	list := make([]entity.HrDepartment, 0)
	db := global.HrDb.Preload("Children")
	// 此处填写 where参数判断
	if data.ID != 0 {
		db = db.Where("id = ?", data.ID)
	}
	if data.CompleteName != "" {
		db = db.Where("complete_name like ?", "%"+data.CompleteName+"%")
	}
	// if data.Active != "" {
	// 	db = db.Where("status = ?", data.Status)
	// }
	err := db.Order("hr_department.id").Select("DISTINCT hr_department.*").Find(&list).Error
	return &list, err
}

// 部门数据修改
func (m *hrDepartmentModelImpl) Update(data entity.HrDepartment) odoorpc.ResultData {
	dd, _ := json.Marshal(data)
	var writeuData map[string]interface{}
	_ = json.Unmarshal(dd, &writeuData)
	writeuResult, err := odoorpc.Write(writeuData, "hr.department", data.ID) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	if err != nil {
		return odoorpc.ResultData{Status: "fail", Code: 500, Message: "修改失败！"}
	}
	return writeuResult
}

// 部门归档 没有所属员工
func (m *hrDepartmentModelImpl) DeleteNoEmployee(Ids []int64) odoorpc.ResultData {
	for i := 0; i < len(Ids); i++ {
		result := global.HrDb.Model(&entity.HrDepartment{}).Where("id = ?", Ids[i]).Update("active", false)
		if result.Error != nil {
			return odoorpc.ResultData{Status: "fail", Code: 500, Message: "归档失败！"}
		}

	}
	return odoorpc.ResultData{Status: "success", Code: 200, Message: "成功归档！"}
}

// 部门归档 先对所属员工进行部门调动，而后修改操作目标部门状态
func (m *hrDepartmentModelImpl) Delete(users []int64, Ids []int64, transferDepId int64) odoorpc.ResultData {
	// user := make([]entity.HrEmployee, 0)
	// err := global.HrDb.Table("hr_employee").Select("id").Where("department_id IN ?", Ids).Find(&user).Error
	// if err != nil {
	// 	return odoorpc.ResultData{Status: "fail", Code: 500, Message: "原有员工处理失败！"}
	// }
	// fmt.Printf("users = %+v\n", user)

	// 需要先删除员工绑定的部门id
	dd, _ := json.Marshal(entity.HrEmployee{DepartmentId: transferDepId})
	var writeuData map[string]interface{}
	_ = json.Unmarshal(dd, &writeuData)
	for i := 0; i < len(users); i++ {
		// result := global.HrDb.Model(&entity.HrEmployee{}).Where("id = ?", users[i]).Update("department_id", transferDepId)
		// if result.Error != nil {
		// 	return odoorpc.ResultData{Status: "fail", Code: 500, Message: "归档失败！"}
		// }
		_, err := odoorpc.Write(writeuData, "hr.employee", users[i]) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
		if err != nil {
			return odoorpc.ResultData{Status: "fail", Code: 500, Message: "原有员工部门调动失败！"}
		}
	}

	// 再将部门状态设置为false
	for i := 0; i < len(Ids); i++ {
		result := global.HrDb.Model(&entity.HrDepartment{}).Where("id = ?", Ids[i]).Update("active", false)
		if result.Error != nil {
			return odoorpc.ResultData{Status: "fail", Code: 500, Message: "归档失败！"}
		}

	}
	return odoorpc.ResultData{Status: "success", Code: 200, Message: "成功归档！"}
	// return global.Db.Table(m.table).Delete(&entity.Hr{}, "organization_id in (?)", organizationIds).Error
}

// 个人组织树 处理代码 公司-子公司-部门-子部门
func (m *hrDepartmentModelImpl) SelectOrganization(data entity.HrDepartment) ([]entity.ResCompanyB, error) {
	list, err := m.FindList(data)
	if err != nil {
		return nil, err
	}
	sd := make([]entity.ResCompanyB, 0)
	li := *list
	// 开始根据公司遍历
	for j := 0; j < len(li); j++ {
		lip := li[j] //进入每一个公司下
		fmt.Printf("进入每一个公司下:%+v\n", lip)
		// 开始遍历公司下属部门
		for i := 0; i < len(lip.Departments); i++ {
			dep := lip.Departments[i]
			if dep.ParentId != 0 {
				continue
			}
			info := Digui(list, li[i])

			sd = append(sd, info)
		}
	}
	// for i := 0; i < len(li); i++ {
	// 	if li[i].ParentId != 0 {
	// 		continue
	// 	}
	// 	info := Digui(list, li[i])
	// 	sd = append(sd, info)
	// }
	return sd, nil
}

func (m *hrDepartmentModelImpl) SelectOrganizationLable(data entity.HrDepartment) ([]entity.DepartmentLable, error) {
	organizationlist, err := m.FindList(data)
	// organizationlist := make([]entity.HrDepartment, 0)
	// err := global.HrDb.Find(&organizationlist).Error
	if err != nil {
		return nil, err
	}
	dl := make([]entity.DepartmentLable, 0)
	organizationl := *organizationlist
	for i := 0; i < len(organizationl); i++ {
		if organizationl[i].ParentId != 0 {
			continue
		}
		e := entity.DepartmentLable{}
		e.DepartmentId = organizationl[i].ID
		e.DepartmentName = organizationl[i].Name
		organizationsInfo := DiguiOrganizationLable(organizationlist, e)

		dl = append(dl, organizationsInfo)
	}
	return dl, nil
}

func (m *hrDepartmentModelImpl) SelectOrganizationIds(data entity.HrDepartment) ([]int64, error) {
	organizationlist := make([]entity.HrDepartment, 0)
	err := global.HrDb.Find(&organizationlist).Error
	if err != nil {
		return nil, err
	}
	dl := make([]int64, 0)
	organizationl := organizationlist
	for i := 0; i < len(organizationl); i++ {
		if organizationl[i].ParentId != 0 {
			continue
		}
		dl = append(dl, organizationl[i].ID)
		e := entity.DepartmentLable{}
		e.DepartmentId = organizationl[i].ID
		e.DepartmentName = organizationl[i].Name
		id := DiguiOrganizationId(&organizationlist, e)
		dl = append(dl, id...)
	}
	return dl, nil
}

// menu 指定父组织的信息,递归过程中，menu会不断被修改，最终构成完整的组织层级结构。
func Digui(organizationlist *[]entity.ResCompanyB, menu entity.ResCompanyB) entity.ResCompanyB {
	list := *organizationlist
	min := make([]entity.ResCompanyB, 0)
	for j := 0; j < len(list); j++ {
		// 检查当前遍历的组织是否是当前节点的子组织，如果不是continue跳出循环
		if menu.ID != list[j].ParentId {
			continue
		}
		mi := entity.ResCompanyB{}
		mi.ID = list[j].ID
		mi.ParentId = list[j].ParentId
		mi.ParentPath = list[j].ParentPath
		mi.Name = list[j].Name
		mi.Active = list[j].Active
		mi.Email = list[j].Email
		mi.Phone = list[j].Phone
		// 开始接收公司下属部门 完成公司下部门组织树构建
		department := make([]entity.HrDepartment, 0)
		for i := 0; i < len(list[j].Departments); i++ {
			dep := entity.HrDepartment{}
			dep.ManagerId = list[j].Departments[i].ManagerId
			dep.ParentId = list[j].Departments[i].ParentId
			dep.ParentPath = list[j].Departments[i].ParentPath
			dep.ID = list[j].Departments[i].ID
			dep.Name = list[j].Departments[i].Name
			dep.Active = list[j].Departments[i].Active
			dep.CreateDate = list[j].Departments[i].CreateDate
			dep.Children = []entity.HrDepartment{} // 初始化该部门子组织列表

			department = append(department, dep)
		}
		// mi.Property = list[j].Property
		// mi.DutyId = list[j].DutyId
		// mi.Virtual = list[j].Virtual
		// mi.Code =list[j].Code
		// mi.CreateDate = list[j].CreateDate
		// mi.WriteDate = list[j].WriteDate

		mi.Children = []entity.ResCompanyB{}
		// 递归调用，构建当前子组织的子树
		ms := Digui(organizationlist, mi)
		// 将构建好的子树添加到当前节点的子组织列表
		min = append(min, ms) //在min中追加ms数据并返回
	}
	menu.Children = min
	return menu
}

func DiguiOrganizationLable(organizationlist *[]entity.ResCompanyB, organization entity.DepartmentLable) entity.DepartmentLable {
	list := *organizationlist
	min := make([]entity.DepartmentLable, 0)
	for j := 0; j < len(list); j++ {

		if organization.DepartmentId != list[j].ParentId {
			continue
		}
		mi := entity.DepartmentLable{CompanyId: list[j].ID, CompanyName: list[j].Name, Children: []entity.DepartmentLable{}}
		ms := DiguiOrganizationLable(organizationlist, mi)
		min = append(min, ms)
	}
	organization.Children = min
	return organization
}

// 递归组织ID
func DiguiOrganizationId(organizationlist *[]entity.HrDepartment, organization entity.DepartmentLable) []int64 {
	list := *organizationlist
	min := make([]int64, 0)
	for j := 0; j < len(list); j++ {
		if organization.DepartmentId != list[j].ParentId {
			continue
		}
		min = append(min, list[j].ID)
		mi := entity.DepartmentLable{CompanyId: list[j].ID, CompanyName: list[j].Name, Children: []entity.DepartmentLable{}}
		id := DiguiOrganizationId(organizationlist, mi)
		min = append(min, id...)
	}
	return min
}
