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
		restfulx.NewReqCtx(request, response).WithLog("获取站点分页列表").Handle(s.GetSiteList)
	}).
		Doc("获取站点分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("name", "name").DataType("string")).
		Param(ws.QueryParameter("_name", "别名").DataType("string")).
		Param(ws.QueryParameter("slug", "短标识符").DataType("string")).
		Param(ws.QueryParameter("physical_address", "物理地址，联系地址").DataType("string")).
		Param(ws.QueryParameter("shipping_address", "物流地址，快递地址").DataType("string")).
		Param(ws.QueryParameter("custom_field_data", "配置数据").DataType("string")).
		Param(ws.QueryParameter("time_zone", "时区").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("group_id", "groupId").DataType("string")).
		Param(ws.QueryParameter("region_id", "regionId").DataType("string")).
		Param(ws.QueryParameter("tenant_id", "tenantId").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{name}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取站点详细信息").Handle(s.GetSiteInfo)
	}).
		Doc("获取站点详细信息").
		Param(ws.PathParameter("name", "name").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.DcimSite{}).
		Returns(200, "OK", entity.DcimSite{}).
		Returns(404, "Not Found", nil))

	// 添加站点实例
	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加站点实例").Handle(s.InsertSite)
	}).
		Doc("添加站点实例").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimSite{})) // from the request，Reads定义得是Body中参数

	// 修改站点信息
	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改站点信息").Handle(s.UpdateSite)
	}).
		Doc("修改站点信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimSitegroup{}))

	// 删除站点信息
	ws.Route(ws.PUT("/{ids}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除站点信息").Handle(s.DeleteSite)
	}).
		Doc("删除站点信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("ids", "多id 1,2,3").DataType("string")))

	// 站点归属到站点组
	ws.Route(ws.PUT("/add2group").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("站点归属到站点组").Handle(s.JoinSiteGroup)
	}).
		Doc("站点归属到站点组").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		// Param(ws.PathParameter("name", "name").DataType("string")).
		Reads(entity.Site2Group{}))

	/** 站点组 **/
	ws.Route(ws.GET("/group/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取站点组分页列表").Handle(s.GetSiteGroupList)
	}).
		Doc("获取站点组分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("name", "name").DataType("string")).
		Param(ws.QueryParameter("slug", "slug").DataType("string")).
		Param(ws.QueryParameter("id", "id").DataType("string")).
		Param(ws.QueryParameter("parentId", "parentId").DataType("string")).
		Param(ws.QueryParameter("level", "level").DataType("string")).
		Param(ws.QueryParameter("lft", "lft").DataType("string")).
		Param(ws.QueryParameter("rght", "rght").DataType("string")).
		Param(ws.QueryParameter("treeId", "treeId").DataType("string")).
		Param(ws.QueryParameter("description", "description").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	// 添加站点组实例
	ws.Route(ws.POST("/group").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加站点组实例").Handle(s.InsertSiteGroup)
	}).
		Doc("添加站点组实例").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimSitegroup{})) // from the request，Reads定义得是Body中参数

	// 修改站点组信息
	ws.Route(ws.PUT("/group").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改站点组信息").Handle(s.UpdateSiteGroup)
	}).
		Doc("修改站点组信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DcimSitegroup{}))

	// 删除站点组信息
	ws.Route(ws.PUT("/group/{ids}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除站点组信息").Handle(s.DeleteSiteGroup)
	}).
		Doc("删除站点组信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("ids", "多id 1,2,3").DataType("string")))

	// 站点组添加站点
	ws.Route(ws.PUT("/group/{name}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("站点加入站点组").Handle(s.JoinSiteGroup)
	}).
		Doc("站点加入站点组").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("name", "name").DataType("string")).
		Reads(entity.Site2Group{}))

	// 渲染站点组层级结构
	ws.Route(ws.GET("/groupTree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取站点组层级结构").Handle(s.GetGroupTree)
	}).
		Doc("获取站点组层级结构").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SiteNode{}).
		Returns(200, "OK", entity.SiteNode{}))

	container.Add(ws)
}
