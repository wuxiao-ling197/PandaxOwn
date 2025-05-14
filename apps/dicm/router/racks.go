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

	ws.Route(ws.GET("/racks/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取机柜分页列表").Handle(s.GetRackList)
	}).
		Doc("获取机柜分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("name", "name").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("siteId", "siteId").DataType("string")).
		// Param(ws.QueryParameter("siteId", "siteId").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/rackreserve/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取机柜预留列表").Handle(s.GetRackReserveList)
	}).
		Doc("获取机柜预留列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("rackId", "rackId").DataType("string")).
		Param(ws.QueryParameter("tenantId", "tenantId").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("userId", "userId").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/rackrole/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取机柜类型列表").Handle(s.GetRackRoleList)
	}).
		Doc("获取机柜类型列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("name", "racknameId").DataType("string")).
		Param(ws.QueryParameter("color", "color").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("slug", "slug").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	// /{siteId}/{id}/{status}/{type}
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

	container.Add(ws)
}
