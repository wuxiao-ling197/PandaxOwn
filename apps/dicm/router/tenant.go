package router

import (
	"pandax/apps/dicm/api"
	"pandax/apps/dicm/entity"
	"pandax/apps/dicm/services"

	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitTenantSite(container *restful.Container) {
	s := &api.TenantApp{
		TApp: services.TenantDao,
	}

	ws := new(restful.WebService)
	ws.Path("/tenant").Produces(restful.MIME_JSON)
	tags := []string{"tenant", "租户"}

	// 租户列表 cg
	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取租户分页列表").Handle(s.GetTenantList)
	}).
		Doc("获取租户分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("name", "name").DataType("string")).
		Param(ws.QueryParameter("slug", "slug").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("groupId", "groupId").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	// 租户详情 cg
	ws.Route(ws.GET("/{name}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取租户信息").Handle(s.GetTenantInfo)
	}).
		Doc("获取租户信息").
		Param(ws.PathParameter("id", "id").DataType("string")).
		Param(ws.PathParameter("groupId", "groupId").DataType("string")).
		Param(ws.PathParameter("name", "name").DataType("string")).
		Param(ws.PathParameter("slug", "slug").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.TenancyTenant{}).
		Returns(200, "OK", entity.TenancyTenant{}).
		Returns(404, "Not Found", nil))

	// 添加租户实例 cg
	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加租户实例").Handle(s.InsertTenant)
	}).
		Doc("添加租户实例").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.TenancyTenant{})) // from the request

	// 修改租户信息 cg
	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改租户信息").Handle(s.UpdateTenant)
	}).
		Doc("修改租户信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.TenancyTenant{}))

	// 删除租户信息 多选删除需要实现
	ws.Route(ws.PUT("/{ids}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除租户信息").Handle(s.DeleteTenant)
	}).
		Doc("删除租户信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("ids", "多id 1,2,3").DataType("string")))

	// 租户归属到租户组
	ws.Route(ws.PUT("/add2group").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("租户归属到租户组").Handle(s.JoinTenantGroup)
	}).
		Doc("租户归属到租户组").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// Param(ws.PathParameter("name", "name").DataType("string")).
		Reads(entity.Own2Group{}))

	/** 租户组 */
	ws.Route(ws.GET("/group/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取租户组分页列表").Handle(s.GetTenantGroupList)
	}).
		Doc("获取租户组分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("name", "name").DataType("string")).
		Param(ws.QueryParameter("slug", "slug").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.PathParameter("parentId", "parentId").DataType("string")).
		Param(ws.PathParameter("level", "level").DataType("string")).
		Param(ws.PathParameter("treeId", "treeId").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/group/{name}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取租户组信息").Handle(s.GetTenantGroupInfo)
	}).
		Doc("获取租户组信息").
		Param(ws.PathParameter("id", "id").DataType("string")).
		Param(ws.PathParameter("parentId", "parentId").DataType("string")).
		Param(ws.PathParameter("name", "name").DataType("string")).
		Param(ws.PathParameter("slug", "slug").DataType("string")).
		Param(ws.PathParameter("level", "level").DataType("string")).
		Param(ws.PathParameter("treeId", "treeId").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.TenancyTenantgroup{}).
		Returns(200, "OK", entity.TenancyTenantgroup{}).
		Returns(404, "Not Found", nil))

	// 添加租户组实例
	ws.Route(ws.POST("/group").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加租户组实例").Handle(s.InsertTenantGroup)
	}).
		Doc("添加租户组实例").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.TenancyTenantgroup{})) // from the request，Reads定义得是Body中参数

	// 修改租户组信息
	ws.Route(ws.PUT("/group").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改租户组信息").Handle(s.UpdateTenantGroup)
	}).
		Doc("修改租户组信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.TenancyTenantgroup{}))

	// 删除租户组信息
	ws.Route(ws.PUT("/group/{ids}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除租户组信息").Handle(s.DeleteTenantGroup)
	}).
		Doc("删除租户组信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("ids", "多id 1,2,3").DataType("string")))

	// 租户组添加租户
	ws.Route(ws.PUT("/group/{name}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("租户加入租户组").Handle(s.JoinTenantGroup)
	}).
		Doc("租户加入租户组").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("name", "name").DataType("string")).
		Reads(entity.Own2Group{}))

	// 渲染租户组层级结构
	ws.Route(ws.GET("/groupTree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取租户组层级结构").Handle(s.GetGroupTree)
	}).
		Doc("获取租户组层级结构").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.GroupNode{}).
		Returns(200, "OK", entity.GroupNode{}))

	container.Add(ws)
}
