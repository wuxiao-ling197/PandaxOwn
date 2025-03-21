package services

import (
	"encoding/json"
	"pandax/apps/shared/entity"
	odoorpc "pandax/pkg/device_rpc"
	"pandax/pkg/global"
)

type (
	HrDepartmentModel interface {
		// 新建部门 注意只能在所属公司下操作本公司的下属部门
		Insert(data entity.HrDepartment) (odoorpc.ResultData, error)
		// 查询部门详细数据
		FindOne(Id int64) (*entity.HrDepartment, error)
		// 分页查询 携带员工数据
		FindListPage(page, pageSize int, data entity.HrDepartment) (*[]entity.HrDepartment, int64, error)
		// 常见公司信息+部门数据
		FindList(data entity.HrDepartment) (*[]entity.CompanyWithDepatment, error)
		// 修改部门数据
		Update(data entity.HrDepartment) odoorpc.ResultData
		// 归档部门数据
		Delete(Ids []int64) odoorpc.ResultData
		// SelectOrganization(data entity.HrDepartment) ([]entity.HrDepartment, error)
		// SelectOrganizationLable(data entity.HrDepartment) ([]entity.DepartmentLable, error)
		// SelectOrganizationIds(data entity.HrDepartment) ([]int64, error)
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

func (m *hrDepartmentModelImpl) FindOne(Id int64) (*entity.HrDepartment, error) {
	resData := new(entity.HrDepartment)
	err := global.HrDb.Table("hr_department").Where("id = ?", Id).First(&resData).Error
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

// 部门详细数据+常见公司数据
func (m *hrDepartmentModelImpl) FindList(data entity.HrDepartment) (*[]entity.CompanyWithDepatment, error) {
	list := make([]entity.CompanyWithDepatment, 0)
	db := global.HrDb.Preload("Departments").
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
	// db.Where("delete_time IS NULL")
	err := db.Order("res_company.id").Select("DISTINCT res_company.*").Find(&list).Error
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

// 部门归档
func (m *hrDepartmentModelImpl) Delete(Ids []int64) odoorpc.ResultData {
	var i int
	for i = 0; i < len(Ids); i++ {
		result := global.HrDb.Model(&entity.HrDepartment{}).Where("id = ?", Ids[i]).Update("active", false)
		if result.Error != nil {
			return odoorpc.ResultData{Status: "fail", Code: 500, Message: "归档失败！"}
		}

	}
	return odoorpc.ResultData{Status: "success", Code: 200, Message: "成功归档！"}
	// return global.Db.Table(m.table).Delete(&entity.Hr{}, "organization_id in (?)", organizationIds).Error
}

// 组织树处理代码 公司-子公司-部门-子部门
// func (m *hrDepartmentModelImpl) SelectOrganization(data entity.HrDepartment) ([]entity.HrDepartment, error) {
// 	list, err := m.FindList(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	sd := make([]entity.CompanyWithDepatment, 0)
// 	li := *list
// 	for j = 0; j < len(li.Departments).len; j++ {
// 		lip := li.Departments[j]
// 		for i := 0; i < len(li); i++ {
// 			if li[i].ParentId != 0 {
// 				continue
// 			}
// 			info := Digui(list, li[i])

// 			sd = append(sd, info)
// 		}
// 	}
// 	// for i := 0; i < len(li); i++ {
// 	// 	if li[i].ParentId != 0 {
// 	// 		continue
// 	// 	}
// 	// 	info := Digui(list, li[i])
// 	// 	sd = append(sd, info)
// 	// }
// 	return sd, nil
// }

// func (m *hrDepartmentModelImpl) SelectOrganizationLable(data entity.HrDepartment) ([]entity.DepartmentLable, error) {
// 	// organizationlist, err := m.FindList(data)
// 	organizationlist := make([]entity.HrDepartment, 0)
// 	err := global.HrDb.Find(&organizationlist).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	dl := make([]entity.DepartmentLable, 0)
// 	organizationl := organizationlist
// 	for i := 0; i < len(organizationl); i++ {
// 		if organizationl[i].ParentId != 0 {
// 			continue
// 		}
// 		e := entity.DepartmentLable{}
// 		e.DepartmentId = organizationl[i].ID
// 		e.DepartmentName = organizationl[i].Name
// 		organizationsInfo := DiguiOrganizationLable(&organizationlist, e)

// 		dl = append(dl, organizationsInfo)
// 	}
// 	return dl, nil
// }

// func (m *hrDepartmentModelImpl) SelectOrganizationIds(data entity.HrDepartment) ([]int64, error) {
// 	organizationlist := make([]entity.HrDepartment, 0)
// 	err := global.HrDb.Find(&organizationlist).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	dl := make([]int64, 0)
// 	organizationl := organizationlist
// 	for i := 0; i < len(organizationl); i++ {
// 		if organizationl[i].ParentId != 0 {
// 			continue
// 		}
// 		dl = append(dl, organizationl[i].ID)
// 		e := entity.DepartmentLable{}
// 		e.DepartmentId = organizationl[i].ID
// 		e.DepartmentName = organizationl[i].Name
// 		id := DiguiOrganizationId(&organizationlist, e)
// 		dl = append(dl, id...)
// 	}
// 	return dl, nil
// }

// func Digui(organizationlist *[]entity.CompanyWithDepatment, menu entity.HrDepartment) entity.HrDepartment {
// 	list := *organizationlist
// 	min := make([]entity.HrDepartment, 0)
// 	for j := 0; j < len(list); j++ {

// 		if menu.ID != list[j].ParentId {
// 			continue
// 		}
// 		mi := entity.HrDepartment{}
// 		mi.ID = list[j].ID
// 		mi.ParentId = list[j].ParentId
// 		mi.ParentPath = list[j].Departments.ParentPath
// 		mi.Name = list[j].Name
// 		mi.ManagerId = list[j].ManagerId
// 		// mi.Property = list[j].Property
// 		// mi.DutyId = list[j].DutyId
// 		// mi.Virtual = list[j].Virtual
// 		// mi.Code =list[j].Code
// 		mi.Active = list[j].Active
// 		mi.CreateDate = list[j].CreateDate
// 		mi.WriteDate = list[j].WriteDate
// 		mi.Children = []entity.HrDepartment{}
// 		ms := Digui(organizationlist, mi)
// 		min = append(min, ms)
// 	}
// 	menu.Children = min
// 	return menu
// }
// func DiguiOrganizationLable(organizationlist *[]entity.HrDepartment, organization entity.DepartmentLable) entity.DepartmentLable {
// 	list := *organizationlist
// 	min := make([]entity.DepartmentLable, 0)
// 	for j := 0; j < len(list); j++ {

// 		if organization.DepartmentId != list[j].ParentId {
// 			continue
// 		}
// 		mi := entity.DepartmentLable{list[j].ID, list[j].Name, []entity.DepartmentLable{}}
// 		ms := DiguiOrganizationLable(organizationlist, mi)
// 		min = append(min, ms)
// 	}
// 	organization.Children = min
// 	return organization
// }

// func DiguiOrganizationId(organizationlist *[]entity.HrDepartment, organization entity.DepartmentLable) []int64 {
// 	list := *organizationlist
// 	min := make([]int64, 0)
// 	for j := 0; j < len(list); j++ {
// 		if organization.DepartmentId != list[j].ParentId {
// 			continue
// 		}
// 		min = append(min, list[j].ID)
// 		mi := entity.DepartmentLable{list[j].ID, list[j].Name, []entity.DepartmentLable{}}
// 		id := DiguiOrganizationId(organizationlist, mi)
// 		min = append(min, id...)
// 	}
// 	return min
// }
