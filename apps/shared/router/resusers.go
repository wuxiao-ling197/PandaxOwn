package router

import (
	"pandax/apps/shared/api"
	"pandax/apps/shared/entity"
	"pandax/apps/shared/services"
	"pandax/apps/system/api/form"
	"pandax/apps/system/api/vo"
	server "pandax/apps/system/services"

	"github.com/PandaXGO/PandaKit/model"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"

	logServices "pandax/apps/log/services"
	odoorpc "pandax/pkg/device_rpc"

	"github.com/PandaXGO/PandaKit/restfulx"
)

// 初始化路由
func InitResUserRouter(container *restful.Container) {
	// 创建API处理实例，注册服务
	s := &api.UserApi{
		// RoleApp:         services.ResUsersModelDao,
		MenuApp:         server.SysMenuModelDao,
		RoleMenuApp:     server.SysRoleMenuModelDao,
		UserApp:         services.ResUsersModelDao,
		LogLogin:        logServices.LogLoginModelDao,
		HrDepartmentApp: services.HrDepartmentModelDao,
		HrJobApp:        services.HrJobModelDao,
	}
	ws := new(restful.WebService)
	ws.Path("/system/user").Produces(restful.MIME_JSON)
	tags := []string{"system", "用户"}

	// 获取登录用户信息
	// ws.Route(ws.Get("/getLoginUser").To(s.GetSysUser()))

	ws.Route(ws.GET("/getCaptcha").To(s.GenerateCaptcha).Doc("获取验证码"))

	// 系统外重启获取TOTP cg
	ws.Route(ws.POST("/enableTotp").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedToken(false).WithNeedCasbin(false).Handle(s.EnableTotp)
	}).
		Doc("激活双重验证").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(form.Login{}).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	// 验证TOTP cg
	ws.Route(ws.POST("/valideTotp").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedToken(false).WithNeedCasbin(false).Handle(s.ValideTotp)
	}).
		Doc("验证TOTP").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(form.Login{}).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	// 系统外重启重置TOTP cg
	ws.Route(ws.POST("/resetTotp").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedToken(false).WithNeedCasbin(false).Handle(s.ResetTOTP)
	}).
		Doc("重置双重验证").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(form.Login{}).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	// 登录 cg
	ws.Route(ws.POST("/login").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedToken(false).WithNeedCasbin(false).WithLog("登录").Handle(s.Login)
	}).
		Doc("登录").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(form.Login{}).
		Writes(vo.TokenVo{}).
		Returns(200, "OK", vo.TokenVo{}))

	// 登出 cg
	ws.Route(ws.POST("/logout").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedToken(false).WithNeedCasbin(false).WithLog("退出登录").Handle(s.LogOut)
	}).
		Doc("退出登录").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	// 获取权限 cg
	ws.Route(ws.GET("/auth").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("认证信息").Handle(s.Auth)
	}).
		Doc("认证信息").
		Param(ws.QueryParameter("username", "username").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(vo.AuthVo{}).
		Returns(200, "OK", vo.AuthVo{}))

	// cg
	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("得到用户分页列表").Handle(s.GetSysUserList)
	}).
		Doc("得到用户分页列表,默认查询的是状态为正常的数据").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("active", "active").DataType("string")).
		Param(ws.QueryParameter("username", "username").DataType("string")).
		Param(ws.QueryParameter("login", "login").DataType("string")).
		Param(ws.QueryParameter("departmentId", "departmentId").DataType("string")).
		Param(ws.QueryParameter("work_phone", "work_phone").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/me").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取个人信息").Handle(s.GetSysUserProfile)
	}).
		Doc("获取个人信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(vo.UserVo{}).
		Returns(200, "OK", vo.UserVo{}))

	ws.Route(ws.GET("/getById/{userId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取用户信息").Handle(s.GetSysUser)
	}).
		Doc("获取用户信息").
		Param(ws.PathParameter("userId", "Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(vo.UserVo{}).
		Returns(200, "OK", vo.UserVo{}).
		Returns(404, "Not Found", nil))

	// 获取初始化角色岗位信息
	ws.Route(ws.GET("/getInit").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取初始化角色岗位信息(添加用户初始化)").Handle(s.GetSysUserInit)
	}).
		Doc("获取初始化角色岗位信息(添加用户初始化)").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(vo.UserRolePost{}). // on the response
		Returns(200, "OK", vo.UserRolePost{}).
		Returns(404, "Not Found", nil))

	// 获取用户岗位信息 cg
	ws.Route(ws.GET("/getRoPo").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取用户角色岗位信息(添加用户初始化)").Handle(s.GetUserRolePost)
	}).
		Doc("获取用户角色岗位信息(添加用户初始化)").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", vo.UserRolePost{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/compTree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取公司组织结构").Handle(s.GetUserRolePost)
	}).
		Doc("获取公司组织结构").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", vo.UserRolePost{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/deptTree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取部门组织结构").Handle(s.GetUserRolePost)
	}).
		Doc("获取部门组织结构").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", vo.UserRolePost{}).
		Returns(404, "Not Found", nil))

	// 添加用户
	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加用户").Handle(s.InsertSysUser)
	}).
		Doc("添加用户").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ResUsers{}).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{})) // from the request

	// 添加员工
	ws.Route(ws.POST("/employee").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加员工").Handle(s.InsertEmployee)
	}).
		Doc("添加员工").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ResUsers{}).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{})) // from the request

	// 修改用户信息
	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改用户信息").Handle(s.UpdateSysUser)
	}).
		Doc("修改用户信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ResUsers{}).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	// 修改员工信息
	ws.Route(ws.PUT("/employee").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改员工信息").Handle(s.UpdateEmployee)
	}).
		Doc("修改员工信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.HrEmployee{}).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	// 个人信息修改
	ws.Route(ws.PUT("/profile").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("用户修改个人信息").Handle(s.UpdateSysUserSelf)
	}).
		Doc("修改用户个人信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ResUsers{}))

	// ws.Route(ws.PUT("/changeStatus").To(func(request *restful.Request, response *restful.Response) {
	// 	restfulx.NewReqCtx(request, response).WithLog("修改用户状态").Handle(s.UpdateSysUserStu)
	// }).
	// 	Doc("修改用户状态").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(entity.SysUser{}))

	// 用户归档
	ws.Route(ws.DELETE("/{userId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除用户信息").Handle(s.DeleteSysUser)
	}).
		Doc("删除用户信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("userId", "多id 1,2,3").DataType("string")))

	// ws.Route(ws.POST("/avatar").To(func(request *restful.Request, response *restful.Response) {
	// 	restfulx.NewReqCtx(request, response).WithLog("修改用户头像").Handle(s.InsetSysUserAvatar)
	// }).
	// 	Doc("修改用户头像").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags))

	// 重置密码 cg
	ws.Route(ws.PUT("resetPwd").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改用户密码").Handle(s.SysUserUpdatePwd)
	}).
		Doc("修改用户密码").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	// 启用TOTP双重验证 cg
	/**传入用户登录名 即需要启用该项功能的用户**/
	ws.Route(ws.GET("totp/enable/{userId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("启用TOTP双重验证").Handle(s.GenerateTOTP)
	}).
		Doc("启用TOTP双重验证").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("userId", "用户编号").DataType("string")).
		Param(ws.QueryParameter("name", "用户登录名").DataType("string")).
		Param(ws.QueryParameter("password", "验证密码").DataType("string")).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	// 禁用TOTP双重验证
	ws.Route(ws.DELETE("totp/disabale/{userId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("禁用TOTP双重验证").Handle(s.DisableTOTP)
	}).
		Doc("禁用TOTP双重验证").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("userId", "用户id").DataType("string")).
		Writes(odoorpc.ResultData{}).
		Returns(200, "OK", odoorpc.ResultData{}))

	// 数据导出
	ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("导出用户信息").Handle(s.ExportUser)
	}).
		Doc("导出用户信息").
		Param(ws.QueryParameter("filename", "filename").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("username", "username").DataType("string")).
		Param(ws.QueryParameter("phone", "phone").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)

}
