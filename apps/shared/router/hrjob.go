package router

import (
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"

	// "pandax/apps/system/api"

	"pandax/apps/shared/api"
	"pandax/apps/shared/entity"
	"pandax/apps/shared/services"
	pservices "pandax/apps/system/services"
	odoorpc "pandax/pkg/device_rpc"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitHrJobRouter(container *restful.Container) {
	s := &api.HrJobApp{
		PostApp: services.HrJobModelDao,
		UserApp: services.ResUsersModelDao,
		RoleApp: pservices.SysRoleModelDao,
	}
	ws := new(restful.WebService)
	ws.Path("/system/post").Produces(restful.MIME_JSON)
	tags := []string{"system", "部门"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取岗位分页列表").Handle(s.GetPostList)
	}).
		Doc("获取岗位分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("departmentId", "departmentId").DataType("string")).
		Param(ws.QueryParameter("jobName", "jobName").DataType("string")).
		Param(ws.QueryParameter("jobCode", "jobCode").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{jobId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取岗位信息").Handle(s.GetPost)
	}).
		Doc("获取岗位信息").
		Param(ws.PathParameter("jobId", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.HrJob{}).
		Returns(200, "OK", entity.HrJob{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加岗位信息").Handle(s.InsertPost)
	}).
		Doc("添加岗位信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.HrJob{}).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改岗位信息").Handle(s.UpdatePost)
	}).
		Doc("修改岗位信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.HrJob{}).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	ws.Route(ws.GET("/publish/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("发布岗位信息").Handle(s.PublishPost)
	}).
		Doc("发布岗位信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "Id").DataType("int")).
		// Reads(entity.HrJob{}).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	// ws.Route(ws.DELETE("/{jobId}").To(func(request *restful.Request, response *restful.Response) {
	// 	restfulx.NewReqCtx(request, response).WithLog("删除岗位信息").Handle(s.DeletePost)
	// }).
	// 	Doc("删除岗位信息").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Param(ws.PathParameter("jobId", "多id 1,2,3").DataType("string")).
	// 	Writes(odoorpc.ResultData{}).
	// 	Returns(200, "OK", odoorpc.ResultData{}))

	container.Add(ws)
}
