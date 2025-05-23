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

func InitDicmRack(container *restful.Container) {
	s := &api.DcimRackApp{
		RackApp: services.DcimRackDao,
	}

	ws := new(restful.WebService)
	ws.Path("/dicm").Produces(restful.MIME_JSON)
	tags := []string{"dicm", "机柜"}

	/** 机柜 QueryParameter中第一个参数需与前端传过来的字段匹配 */
	ws.Route(ws.GET("/racks/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取机柜分页列表").Handle(s.GetRackList)
	}).
		Doc("获取机柜分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("name", "name").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("_name", "_name").DataType("string")).
		Param(ws.QueryParameter("type", "type").DataType("string")).
		Param(ws.QueryParameter("serial", "serial").DataType("string")).
		Param(ws.QueryParameter("width", "width").DataType("string")).
		Param(ws.QueryParameter("u_height", "u_height").DataType("string")).
		Param(ws.QueryParameter("starting_unit", "starting_unit").DataType("string")).
		Param(ws.QueryParameter("mounting_depth", "mounting_depth").DataType("string")).
		Param(ws.QueryParameter("outer_depth", "outer_depth").DataType("string")).
		Param(ws.QueryParameter("outer_width", "outer_width").DataType("string")).
		Param(ws.QueryParameter("asset_tag", "asset_tag").DataType("string")).
		Param(ws.QueryParameter("custom_field_data", "custom_field_data").DataType("string")).
		Param(ws.QueryParameter("weight", "weight").DataType("string")).
		Param(ws.QueryParameter("facility_id", "facility_id").DataType("string")).
		Param(ws.QueryParameter("role_id", "role_id").DataType("string")).
		Param(ws.QueryParameter("location_id", "location_id").DataType("string")).
		Param(ws.QueryParameter("tenant_id", "tenant_id").DataType("string")).
		Param(ws.QueryParameter("description", "description").DataType("string")).
		Param(ws.QueryParameter("comments", "comments").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	// 查询机柜详情
	ws.Route(ws.GET("/racks/{name}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取机柜信息").Handle(s.GetRackInfo)
	}).
		Doc("获取机柜信息").
		// Param(ws.PathParameter("id", "id").DataType("string")).
		// Param(ws.PathParameter("siteId", "siteId").DataType("string")).
		Param(ws.PathParameter("name", "name").DataType("string")).
		// Param(ws.PathParameter("status", "status").DataType("string")).
		// Param(ws.PathParameter("type", "type").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.DcimRack{}).
		Returns(200, "OK", entity.DcimRack{}).
		Returns(404, "Not Found", nil))

	// 添加机柜实例
	ws.Route(ws.POST("/racks").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加机柜实例").Handle(s.InsertRack)
	}).
		Doc("添加机柜实例").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimRack{})) // from the request

	// 修改机柜信息
	ws.Route(ws.PUT("/racks").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改机柜信息").Handle(s.UpdateRack)
	}).
		Doc("修改机柜信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimRack{}))

	// 删除机柜信息 多选删除需要实现
	ws.Route(ws.PUT("/racks/{ids}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除机柜信息").Handle(s.DeleteRack)
	}).
		Doc("删除机柜信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("ids", "多id 1,2,3").DataType("string")))

	/** 机柜预留 */
	ws.Route(ws.GET("/rackreserve/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取机柜预留列表").Handle(s.GetRackReserveList)
	}).
		Doc("获取机柜预留列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("rack_id", "rack_id").DataType("string")).
		Param(ws.QueryParameter("tenant_id", "tenant_id").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("user_id", "user_id").DataType("string")).
		Param(ws.QueryParameter("units", "units").DataType("string")).
		Param(ws.QueryParameter("comments", "comments").DataType("string")).
		Param(ws.QueryParameter("custom_field_data", "custom_field_data").DataType("string")).
		Param(ws.QueryParameter("description", "description").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	// 添加实例
	ws.Route(ws.POST("/rackreserve").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加机柜预留实例").Handle(s.InsertReservation)
	}).
		Doc("添加机柜预留实例").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimRackreservation{})) // from the request

	// 修改机柜信息
	ws.Route(ws.PUT("/rackreserve").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改机柜预留信息").Handle(s.UpdateReservation)
	}).
		Doc("修改机柜预留信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimRackreservation{}))

	// 删除机柜信息 多选删除需要实现
	ws.Route(ws.PUT("/rackreserve/{ids}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除机柜预留信息").Handle(s.DeleteReservation)
	}).
		Doc("删除机柜预留信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("ids", "多id 1,2,3").DataType("string")))

	/** 机柜角色 */
	ws.Route(ws.GET("/rackrole/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取机柜类型列表").Handle(s.GetRackRoleList)
	}).
		Doc("获取机柜类型列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("name", "name").DataType("string")).
		Param(ws.QueryParameter("color", "color").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("slug", "slug").DataType("string")).
		Param(ws.QueryParameter("custom_field_data", "custom_field_data").DataType("string")).
		Param(ws.QueryParameter("description", "description").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	// 添加机柜实例
	ws.Route(ws.POST("/rackrole").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加机柜角色实例").Handle(s.InsertRackRole)
	}).
		Doc("添加机柜角色实例").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimRackrole{})) // from the request

	// 修改机柜信息
	ws.Route(ws.PUT("/rackrole").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改机柜角色信息").Handle(s.UpdateRackRole)
	}).
		Doc("修改机柜角色信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimRackrole{}))

	// 删除机柜信息 多选删除需要实现
	ws.Route(ws.PUT("/rackrole/{ids}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除机柜角色信息").Handle(s.DeleteRackRole)
	}).
		Doc("删除机柜角色信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("ids", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
