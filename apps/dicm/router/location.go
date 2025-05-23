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

func InitDicmLocation(container *restful.Container) {
	s := &api.DcimLocationApp{
		LocationApp: services.DcimLocationDao,
	}

	ws := new(restful.WebService)
	ws.Path("/dicm").Produces(restful.MIME_JSON)
	tags := []string{"dicm", "地区"}

	/** Location **/
	ws.Route(ws.GET("/locations/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取物理位置分页列表").Handle(s.GetLocationList)
	}).
		Doc("获取物理位置分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("name", "name").DataType("string")).
		Param(ws.QueryParameter("slug", "slug").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("level", "level").DataType("string")).
		Param(ws.QueryParameter("treeId", "treeId").DataType("string")).
		Param(ws.QueryParameter("tenantId", "tenantId").DataType("string")).
		Param(ws.QueryParameter("parentId", "parentId").DataType("string")).
		Param(ws.QueryParameter("siteId", "siteId").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/locations/{name}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取物理位置信息").Handle(s.GetLocationInfo)
	}).
		Doc("获取物理位置信息").
		Param(ws.PathParameter("name", "name").DataType("string")).
		Param(ws.PathParameter("slug", "slug").DataType("string")).
		Param(ws.PathParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("level", "level").DataType("string")).
		Param(ws.QueryParameter("treeId", "treeId").DataType("string")).
		Param(ws.QueryParameter("tenantId", "tenantId").DataType("string")).
		Param(ws.QueryParameter("parentId", "parentId").DataType("string")).
		Param(ws.QueryParameter("siteId", "siteId").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.DcimLocation{}).
		Returns(200, "OK", entity.DcimLocation{}).
		Returns(404, "Not Found", nil))

	// 添加实例
	ws.Route(ws.POST("/locations").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加物理位置实例").Handle(s.InsertLocation)
	}).
		Doc("添加物理位置实例").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimLocation{}))

	// 修改信息
	ws.Route(ws.PUT("/locations").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改物理位置信息").Handle(s.UpdateLocation)
	}).
		Doc("修改物理位置信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimLocation{}))

	// location信息归档
	ws.Route(ws.PUT("/locations/{ids}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("location信息归档").Handle(s.DeleteLocation)
	}).
		Doc("location信息归档").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("ids", "多id 1,2,3").DataType("string")))

	// 渲染location层级结构
	ws.Route(ws.GET("/locations/tree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("渲染location层级结构").Handle(s.GetLocationTree)
	}).
		Doc("渲染location层级结构").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.LocationNode{}).
		Returns(200, "OK", entity.LocationNode{}))

	/** Region **/
	ws.Route(ws.GET("/region/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取地区分页列表").Handle(s.GetRegionList)
	}).
		Doc("获取地区分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("name", "name").DataType("string")).
		Param(ws.QueryParameter("slug", "slug").DataType("string")).
		Param(ws.QueryParameter("custom_field_data", "custom_field_data").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("level", "level").DataType("string")).
		Param(ws.QueryParameter("treeId", "treeId").DataType("string")).
		Param(ws.QueryParameter("parentId", "parentId").DataType("string")).
		Param(ws.QueryParameter("description", "description").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/region/{name}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取地区信息").Handle(s.GetRegionInfo)
	}).
		Doc("获取地区信息").
		Param(ws.PathParameter("name", "name").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.DcimRegion{}).
		Returns(200, "OK", entity.DcimRegion{}).
		Returns(404, "Not Found", nil))

	// 添加实例
	ws.Route(ws.POST("/region").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加地区实例").Handle(s.InsertRegion)
	}).
		Doc("添加地区实例").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimRegion{}))

	// 修改信息
	ws.Route(ws.PUT("/region").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改地区信息").Handle(s.UpdateRegion)
	}).
		Doc("修改地区信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimRegion{}))

	// 归档
	ws.Route(ws.PUT("/region/{ids}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("region信息归档").Handle(s.DeleteRegion)
	}).
		Doc("region信息归档").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("ids", "多id 1,2,3").DataType("string")))

	// 渲染层级结构
	ws.Route(ws.GET("/region/tree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("渲染地区层级结构").Handle(s.GetRegionTree)
	}).
		Doc("渲染地区层级结构").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.DcimRegion{}).
		Returns(200, "OK", entity.DcimRegion{}))

	container.Add(ws)
}
