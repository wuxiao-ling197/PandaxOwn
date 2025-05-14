package router

import (
	// "pandax/apps/system/api"

	// "pandax/apps/system/entity"
	// "pandax/apps/system/services"
	"pandax/apps/shared/api"
	"pandax/apps/shared/entity"
	"pandax/apps/shared/services"
	pservices "pandax/apps/system/services"
	odoorpc "pandax/pkg/device_rpc"

	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

// 初始化路由
func InitHeDepartmentRouter(container *restful.Container) {
	// 创建api处理实例，注入服务
	s := &api.HrDepartmentApi{
		HrDepartmentApp: services.HrDepartmentModelDao,
		RoleApp:         pservices.SysRoleModelDao,
		UserApp:         services.ResUsersModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/system/organization").Produces(restful.MIME_JSON)
	tags := []string{"system", "组织"}

	// ws.Route(ws.GET("/roleOrganizationTreeSelect/{roleId}").To(func(request *restful.Request, response *restful.Response) {
	// 	restfulx.NewReqCtx(request, response).WithLog("获取角色组织树").Handle(s.GetOrganizationTreeRoleSelect)
	// }).
	// 	Doc("获取角色组织树").
	// 	Param(ws.PathParameter("roleId", "角色Id").DataType("int").DefaultValue("1")).
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Writes(vo.OrganizationTreeVo{}).
	// 	Returns(200, "OK", vo.OrganizationTreeVo{}).
	// 	Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/organizationTree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取所有组织树").Handle(s.GetOrganizationTree)
	}).
		Doc("获取所有组织树").
		Param(ws.QueryParameter("organizationName", "organizationName").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("organizationId", "organizationId").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.ResCompanyB{}).
		Returns(200, "OK", []entity.ResCompanyB{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/departmentTree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取部门组织树").Handle(s.GetDepartmentTree)
	}).
		Doc("获取部门组织树").
		Param(ws.QueryParameter("departmentName", "departmentName").DataType("string")).
		Param(ws.QueryParameter("active", "active").DataType("string")).
		Param(ws.QueryParameter("departmentId", "departmentId").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.ResCompanyB{}).
		Returns(200, "OK", []entity.ResCompanyB{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取组织列表").Handle(s.GetOrganizationList)
	}).
		Doc("获取组织列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("departmentName", "departmentName").DataType("string")).
		Param(ws.QueryParameter("active", "active").DataType("string")).
		Param(ws.QueryParameter("departmentId", "departmentId").DataType("int")).
		Param(ws.QueryParameter("companyId", "companyId").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]model.ResultPage{}).
		Returns(200, "OK", []model.ResultPage{}))

	// 查询公司信息
	ws.Route(ws.GET("/company").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取公司信息").Handle(s.GetCompanyList)
	}).
		Doc("获取公司信息").
		Param(ws.QueryParameter("name", "name").DataType("string")).
		Param(ws.QueryParameter("email", "email").DataType("string")).
		Param(ws.QueryParameter("phone", "phone").DataType("int")).
		Param(ws.QueryParameter("id", "id").DataType("int")).
		// Param(ws.QueryParameter("manager", "manager").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.ResCompanyB{}).
		Returns(200, "OK", []entity.ResCompanyB{}))

	ws.Route(ws.GET("/{organizationId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取组织信息").Handle(s.GetOrganization)
	}).
		Doc("获取组织信息").
		Param(ws.PathParameter("organizationId", "组织Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.HrDepartment{}). // on the response
		Returns(200, "OK", entity.HrDepartment{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加组织信息").Handle(s.InsertOrganization)
	}).
		Doc("添加组织信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.HrDepartment{}).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改组织信息").Handle(s.UpdateOrganization)
	}).
		Doc("修改组织信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.HrDepartment{}).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	ws.Route(ws.DELETE("/{organizationId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除组织信息").Handle(s.DeleteOrganization)
	}).
		Doc("删除组织信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("organizationId", "多id 1,2,3").DataType("int")).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	container.Add(ws)

}
