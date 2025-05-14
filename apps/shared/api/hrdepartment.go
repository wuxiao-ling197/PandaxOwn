package api

import (
	"errors"
	"fmt"
	"pandax/apps/shared/entity"
	"pandax/apps/shared/services"
	pservices "pandax/apps/system/services"
	"pandax/pkg/global"
	"strconv"

	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	"github.com/kakuilan/kgo"
)

type HrDepartmentApi struct {
	HrDepartmentApp services.HrDepartmentModel
	UserApp         services.ResUsersModel
	RoleApp         pservices.SysRoleModel
}

// 组织树
// func (m *HrDepartmentApi) GetOrganizationTreeRoleSelect(rc *restfulx.ReqCtx) {
// 	// roleId := restfulx.PathParamInt(rc, "roleId")
// 	var organization entity.HrDepartment
// 	result, err := m.HrDepartmentApp.SelectOrganizationLable(organization)
// 	biz.ErrIsNil(err, "查询组织树失败")
// 	organizationIds := make([]int64, 0)
// 	// if roleId != 0 {
// 	// 	organizationIds, err = m.RoleApp.GetRoleOrganizationId(pentity.SysRole{RoleId: 2})
// 	// 	biz.ErrIsNil(err, "查询角色组织失败")
// 	// }
// 	organizationIds, err = m.RoleApp.GetRoleOrganizationId(pentity.SysRole{RoleId: 2})
// 	biz.ErrIsNil(err, "查询角色组织失败")
// 	rc.ResData = vo.OrganizationTreeVo{
// 		Organizations: result,
// 		CheckedKeys:   organizationIds,
// 	}
// }

// 部门列表 实际需要公司+部门完整组织数据
func (a *HrDepartmentApi) GetOrganizationList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	departmentName := restfulx.QueryParam(rc, "departmentName")
	status := restfulx.QueryParam(rc, "active")
	departmentId := restfulx.QueryInt(rc, "departmentId", 0)
	companyId := restfulx.QueryInt(rc, "companyId", 0)
	organization := entity.HrDepartment{Name: departmentName, Active: status == "true", ID: int64(departmentId), CompanyId: int64(companyId)}
	// log.Printf("组织列表解析参数：%+v\n", organization)
	// if organization.Name == "" {
	// 	data, err := a.HrDepartmentApp.FindListPage(pageNum, pageSize, organization)
	// 	biz.ErrIsNil(err, "查询部门列表失败")
	// 	rc.ResData = model.ResultPage{
	// 		Total:    total,
	// 		PageNum:  int64(pageNum),
	// 		PageSize: int64(pageSize),
	// 		Data:     data,
	// 	}
	// } else {
	// data, err := a.HrDepartmentApp.FindListPage(organization)
	// biz.ErrIsNil(err, "查询组织列表失败")
	// rc.ResData = data
	data, total, err := a.HrDepartmentApp.FindListPage(pageNum, pageSize, organization)
	biz.ErrIsNil(err, "查询部门列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     data,
		// }
	}
}

func (a *HrDepartmentApi) GetCompanyList(rc *restfulx.ReqCtx) {
	var organization entity.ResCompanyB
	name := restfulx.QueryParam(rc, "name")
	email := restfulx.QueryParam(rc, "email")
	phone := restfulx.QueryParam(rc, "phone")
	id := restfulx.QueryInt(rc, "id", 0)
	// manager := restfulx.QueryInt(rc, "manager", 0)
	organization.ID = int64(id)
	organization.Name = name
	organization.Phone = phone
	organization.Email = email
	data, err := a.HrDepartmentApp.FindCompanyList(organization)
	biz.ErrIsNil(err, "查询组织列表失败")
	rc.ResData = data
}

// 用户侧栏组织结构————公司和部门完整组织树
func (a *HrDepartmentApi) GetOrganizationTree(rc *restfulx.ReqCtx) {
	organizationName := restfulx.QueryParam(rc, "organizationName")
	status := kgo.KConv.Str2Bool(restfulx.QueryParam(rc, "status"))
	organizationId := restfulx.QueryInt(rc, "organizationId", 0)
	organization := entity.HrDepartment{Name: organizationName, Active: status, ID: int64(organizationId)}
	data, err := a.HrDepartmentApp.SelectOrganization(organization)
	// company, e := a.HrDepartmentApp.FindListDepartment(organization)
	// biz.ErrIsNil(e, "数据库查询失败.")
	// department, r := a.HrDepartmentApp.FindListDepartment(organization)
	// biz.ErrIsNil(r, "数据库查询失败..")
	// data, err := tool.BuildFullOrganizationTree(*company, *department)
	biz.ErrIsNil(err, "查询组织树失败")
	rc.ResData = data
}

// 部门组织树
func (a *HrDepartmentApi) GetDepartmentTree(rc *restfulx.ReqCtx) {
	organizationName := restfulx.QueryParam(rc, "departmentName")
	status := kgo.KConv.Str2Bool(restfulx.QueryParam(rc, "active"))
	organizationId := restfulx.QueryInt(rc, "departmentId", 0)
	organization := entity.HrDepartment{Name: organizationName, Active: status, ID: int64(organizationId)}
	department, r := a.HrDepartmentApp.FindListDepartment(organization)
	// data := tool.BuildDepartmentTree(*department, make(map[int64]*entity.HrDepartment))
	biz.ErrIsNil(r, "数据库查询失败..")
	rc.ResData = department
}

// 查询指定部门数据
func (a *HrDepartmentApi) GetOrganization(rc *restfulx.ReqCtx) {
	organizationId := restfulx.PathParamInt(rc, "organizationId")
	data, err := a.HrDepartmentApp.FindOne(int64(organizationId))
	biz.ErrIsNil(err, "查询组织失败")
	rc.ResData = data
}

func (a *HrDepartmentApi) InsertOrganization(rc *restfulx.ReqCtx) {
	var organization entity.HrDepartment
	restfulx.BindJsonAndValid(rc, &organization)
	// organization.CreateBy = rc.LoginAccount.UserName
	_, err := a.HrDepartmentApp.Insert(organization)
	biz.ErrIsNil(err, "添加组织失败")
}

func (a *HrDepartmentApi) UpdateOrganization(rc *restfulx.ReqCtx) {
	var organization entity.HrDepartment
	restfulx.BindJsonAndValid(rc, &organization)
	// organization.UpdateBy = rc.LoginAccount.UserName
	result := a.HrDepartmentApp.Update(organization)
	// biz.ErrIsNil(err, "修改组织失败")
	rc.ResData = result
}

// 部门归档 无员工部门
func (a *HrDepartmentApi) DeleteOrganizationNoEmployee(rc *restfulx.ReqCtx) {
	organizationId := restfulx.PathParam(rc, "organizationId")
	// 需要归档或取消的部门,整数数组接收
	organizationIds := utils.IdsStrToIdsIntGroup(organizationId)
	// 接收可删除的组织ID
	deList := make([]int64, 0)
	for _, id := range organizationIds {
		// 先查询部门中所属员工
		list, err := a.HrDepartmentApp.FindEmployee(id)
		if err != nil {
			continue
		}
		employee := list.Employees
		// 如果没有所属员工，可以直接删除，否则报错
		if len(employee) == 0 {
			deList = append(deList, id)
		} else {
			global.Log.Info(fmt.Sprintf("部门编号: 【%d】名称为【%v】中存在用户绑定无法删除", id, list.CompleteName))
		}
	}
	if len(deList) == 0 {
		biz.ErrIsNil(errors.New("所有组织都已绑定用户无法删除"), "所有组织都已绑定用户，无法删除")
	}
	a.HrDepartmentApp.DeleteNoEmployee(deList)
}

// 部门归档 有员工部门 先对员工进行部门调动
func (a *HrDepartmentApi) DeleteOrganization(rc *restfulx.ReqCtx) {
	organizationId := restfulx.PathParam(rc, "organizationId")
	// 原有员工调动到目标部门
	transferDepId, _ := strconv.ParseInt(restfulx.PathParam(rc, "transferDepId"), 10, 64)
	// 需要归档或取消的部门,整数数组接收
	organizationIds := utils.IdsStrToIdsIntGroup(organizationId)
	// 接收需要进行部门调动的组织ID
	empList := make([]int64, 0)
	for _, id := range organizationIds {
		// 先查询部门中所属员工
		list, err := a.HrDepartmentApp.FindEmployee(id)
		if err != nil {
			continue
		}
		employee := list.Employees
		// 如果没有所属员工，可以直接删除，否则报错
		if len(employee) != 0 {
			for emp := 0; emp < len(employee); emp++ {
				empList = append(empList, employee[emp].ID)
			}
		} else {
			global.Log.Info(fmt.Sprintln("该部门没有所属员工，请直接删除！"))
		}
	}
	if len(empList) == 0 {
		biz.ErrIsNil(errors.New("该部门没有绑定用户，请直接删除"), "该部门没有绑定用户，请直接删除")
	}
	result := a.HrDepartmentApp.Delete(empList, organizationIds, transferDepId)
	rc.ResData = result
}
