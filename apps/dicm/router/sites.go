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

func InitDicmSite(container *restful.Container) {
	s := &api.DcimSiteApp{
		SiteApp: services.DcimSiteDao,
	}

	ws := new(restful.WebService)
	ws.Path("/dicm/sites").Produces(restful.MIME_JSON)
	tags := []string{"dicm", "站点"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取机柜分页列表").Handle(s.GetSiteList)
	}).
		Doc("获取机柜分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("name", "name").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("groupId", "groupId").DataType("string")).
		Param(ws.QueryParameter("regionId", "regionId").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{regionId}/{id}/{name}/{status}/{groupId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取机柜信息").Handle(s.GetSiteInfo)
	}).
		Doc("获取机柜信息").
		Param(ws.PathParameter("id", "id").DataType("string")).
		Param(ws.PathParameter("groupId", "groupId").DataType("string")).
		Param(ws.PathParameter("name", "name").DataType("string")).
		Param(ws.PathParameter("status", "status").DataType("string")).
		Param(ws.PathParameter("regionId", "regionId").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.DcimSite{}).
		Returns(200, "OK", entity.DcimSite{}).
		Returns(404, "Not Found", nil))

	container.Add(ws)
}
