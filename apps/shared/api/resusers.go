package api

import (
	"fmt"
	"log"
	logEntity "pandax/apps/log/entity"
	logServices "pandax/apps/log/services"
	"pandax/apps/shared/entity"
	"pandax/apps/shared/services"
	"pandax/apps/system/api/form"
	"pandax/apps/system/api/vo"
	pentity "pandax/apps/system/entity"
	pservices "pandax/apps/system/services"
	"pandax/pkg/global"
	"pandax/pkg/tool"
	"strings"
	"time"

	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kakuilan/kgo"
	"github.com/mssola/user_agent"

	odoorpc "pandax/pkg/device_rpc"

	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
)

type UserApi struct {
	UserApp         services.ResUsersModel
	MenuApp         pservices.SysMenuModel
	HrJobApp        services.HrJobModel
	RoleApp         pservices.SysRoleModel
	RoleMenuApp     pservices.SysRoleMenuModel
	HrDepartmentApp services.HrDepartmentModel
	LogLogin        logServices.LogLoginModel
}

// GenerateCaptcha 获取验证码
// func (u *UserApi) GenerateCaptcha(request *restful.Request, response *restful.Response) {
// 	id, image, _ := captcha.Generate()
// 	response.WriteEntity(vo.CaptchaVo{Base64Captcha: image, CaptchaId: id})
// }

// GenerateTOTP 获取并激活TOTP 用户管理页面
func (u *UserApi) GenerateTOTP(rc *restfulx.ReqCtx) {
	password := restfulx.QueryParam(rc, "password")
	name := restfulx.QueryParam(rc, "name")
	// fmt.Printf("totp后端接收参数：%v,%v", name, password)
	mfa := &tool.TotpMfa{}
	user := new(entity.ResUsers)
	global.HrDb.Table("res_users").Where("id = ?", rc.LoginAccount.UserId).Find(&user)
	b := tool.VerifyPwd(password, user.Password)
	if !b {
		rc.ResData = odoorpc.ResultData{
			Code:    400,
			Message: "密码验证失败，请重新输入",
		}
	}
	totp, _ := mfa.Initiate(name, rc.LoginAccount.UserId)
	// fmt.Printf("totp后端初始化：%+v\n", totp)
	rc.ResData = odoorpc.ResultData{
		Code:    200,
		Message: "已成功启用",
		Data:    totp,
	}
}

// 激活totp 仅在第一次认证时调用 登录页面
func (u *UserApi) EnableTotp(rc *restfulx.ReqCtx) {
	var login form.Login
	mfa := &tool.TotpMfa{}
	restfulx.BindJsonAndValid(rc, &login)
	// fmt.Printf("EnableTotp后端路由接收login：%+v\n", login)
	totp, _ := mfa.Enable(entity.LoginO{Login: login.Username, Password: login.Password})
	// fmt.Printf("EnableTotp后端路由接收totp记录：%+v\n", totp)
	// result_data := {
	// 	totp.Secret,
	// 	tototp.Qrcode,
	// }
	if totp.Secret != "false" {
		rc.ResData = totp.Qrcode
	} else {
		rc.ResData = "用户已激活双重验证"
	}

}

// 验证 TOTP 登录页面
func (u *UserApi) ValideTotp(rc *restfulx.ReqCtx) {
	var login form.Login
	mfa := &tool.TotpMfa{}
	// passcode := restfulx.QueryParam(rc, "passcode")
	restfulx.BindJsonAndValid(rc, &login)
	user := new(entity.ResUsers)
	global.HrDb.Table("res_users").Where("login = ?", login.Username).Find(&user)
	// fmt.Printf("\nValideTotp后端验证totp:用户信息：%+v\n passcode=%v\n", user, login.Passcode)
	wizard, _ := mfa.Enable(entity.LoginO{Login: login.Username, Password: login.Password})
	if user.PandaxSecret == "false" && wizard.Secret == "false" {
		biz.NewBizErr("用户已禁用双重验证，请联系管理员启用")
	}
	result, _ := mfa.Verify(login.Passcode, wizard)
	biz.IsTrue(result, "验证码认证失败")
	rc.ResData = result
}

// GenerateTOTP 获取并激活TOTP 用户管理页面
func (u *UserApi) ResetTOTP(rc *restfulx.ReqCtx) {
	var login form.Login
	mfa := &tool.TotpMfa{}
	restfulx.BindJsonAndValid(rc, &login)
	user := new(entity.ResUsers)
	global.HrDb.Table("res_users").Where("login = ?", login.Username).Find(&user)
	b := tool.VerifyPwd(login.Password, user.Password)
	biz.IsTrue(b, "密码错误，请重新输入")
	mfa.Disable(user.ID, user.ID)
	mfa.Initiate(login.Username, user.ID)
	rc.ResData = "用户双重验证已重置"
}

// 禁用TOTP
func (u *UserApi) DisableTOTP(rc *restfulx.ReqCtx) {
	userId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "userId"))
	mfa := &tool.TotpMfa{}
	result := mfa.Disable(userId, rc.LoginAccount.UserId)
	if result != nil {
		rc.ResData = result.Error()
	}
	rc.ResData = odoorpc.ResultData{
		Code:    200,
		Message: "已禁用TOTP",
	}
}

// todo: 登录页重置totp 需要完善
// func (u *UserApi) ResetTotp(request *restful.Request, response *restful.Response) {
// 	mfa := &tool.TotpMfa{}
// 	// mfa.Disable()
// 	totp, _ := mfa.Initiate("create", 2)
// 	response.WriteEntity(totp)
// }

// RefreshToken 刷新token
func (u *UserApi) RefreshToken(rc *restfulx.ReqCtx) {
	tokenStr := rc.Request.Request.Header.Get("X-TOKEN")
	// 如果token为空，从请求参数中获取
	if tokenStr == "" {
		tokenStr = rc.Request.Request.URL.Query().Get("token")
	}
	j := token.NewJWT("", []byte(global.Conf.Jwt.Key), jwt.SigningMethodHS256)
	token, err := j.RefreshToken(tokenStr)
	biz.ErrIsNil(err, "刷新token失败")
	rc.ResData = vo.TokenVo{
		Token:  token,
		Expire: time.Now().Unix() + global.Conf.Jwt.ExpireTime,
	}
}

// Login 用户登录 totp处理
//
//	func (u *UserApi) Login(rc *restfulx.ReqCtx) {
//		log.Printf("后端用户API接收数据：%+v\n", rc)
//		var l form.Login
//		mfa := &tool.TotpMfa{}
//		restfulx.BindJsonAndValid(rc, &l)
//		valid, errr := mfa.Verify(l.Code, entity.LoginO{Login: l.Username})
//		if errr != nil {
//			biz.ErrIsNil(errr, "mfa验证失败！")
//		}
//		biz.IsTrue(valid, "验证码认证失败")
//		login, err := u.UserApp.Login(entity.LoginO{Login: l.Username, Password: l.Password})
//		biz.ErrIsNil(err, "登录失败,用户不存在！")
//		// role, err := u.RoleApp.FindOne(login.RoleId)
//		// biz.ErrIsNil(err, "用户所属角色查询失败")
//		j := token.NewJWT("", []byte(global.Conf.Jwt.Key), jwt.SigningMethodHS256)
//		token, err := j.CreateToken(token.Claims{
//			UserId:         login.ID,
//			UserName:       login.Login,
//			RoleId:         2,
//			RoleKey:        "admin",
//			OrganizationId: login.Employee.DepartmentId,
//			PostId:         login.Employee.JobId,
//			RegisteredClaims: jwt.RegisteredClaims{
//				NotBefore: jwt.NewNumericDate(time.Now()),                                                            // 签名生效时间
//				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Conf.Jwt.ExpireTime) * time.Hour)), // 过期时间 7天  配置文件
//				Issuer:    "PandaX",                                                                                  // 签名的发行者
//			},
//		})
//		biz.ErrIsNil(err, "生成Token失败")
//		rc.ResData = vo.TokenVo{
//			Token:  token,
//			Expire: time.Now().Unix() + global.Conf.Jwt.ExpireTime,
//		}
//		go func() {
//			var loginLog logEntity.LogLogin
//			ua := user_agent.New(rc.Request.Request.UserAgent())
//			loginLog.Ipaddr = rc.Request.Request.RemoteAddr
//			loginLog.LoginLocation = utils.GetRealAddressByIP(rc.Request.Request.RemoteAddr)
//			loginLog.LoginTime = time.Now()
//			loginLog.Status = "0"
//			loginLog.Remark = rc.Request.Request.UserAgent()
//			browserName, browserVersion := ua.Browser()
//			loginLog.Browser = browserName + " " + browserVersion
//			loginLog.Os = ua.OS()
//			loginLog.Platform = ua.Platform()
//			loginLog.Username = login.Login
//			loginLog.Msg = "登录成功"
//			loginLog.CreateBy = fmt.Sprint(login.CreateUid)
//			u.LogLogin.Insert(loginLog)
//		}()
//	}

// 1 Login
func (u *UserApi) Login(rc *restfulx.ReqCtx) {
	var l form.Login
	restfulx.BindJsonAndValid(rc, &l)
	login, err := u.UserApp.Login(entity.LoginO{Login: l.Username, Password: l.Password})
	biz.ErrIsNil(err, "登录失败,用户不存在！")
	// role, err := u.RoleApp.FindOne(1)
	// log.Printf("登录用户角色: %+v\n", role)
	// biz.ErrIsNil(err, "用户所属角色查询失败")
	j := token.NewJWT("", []byte(global.Conf.Jwt.Key), jwt.SigningMethodHS256)
	token, err := j.CreateToken(token.Claims{
		UserId:         login.ID,
		UserName:       login.Login,
		RoleId:         1,
		RoleKey:        "admin",
		OrganizationId: login.Employee.DepartmentId,
		PostId:         login.Employee.JobId,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                                                            // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Conf.Jwt.ExpireTime) * time.Hour)), // 过期时间 7天  配置文件
			Issuer:    "PandaX",                                                                                  // 签名的发行者
		},
	})
	biz.ErrIsNil(err, "生成Token失败")

	rc.ResData = vo.TokenVo{
		Token:  token,
		Expire: time.Now().Unix() + global.Conf.Jwt.ExpireTime,
	}
	go func() {
		var loginLog logEntity.LogLogin
		ua := user_agent.New(rc.Request.Request.UserAgent())
		loginLog.Ipaddr = rc.Request.Request.RemoteAddr
		loginLog.LoginLocation = utils.GetRealAddressByIP(rc.Request.Request.RemoteAddr)
		loginLog.LoginTime = time.Now()
		loginLog.Status = "0"
		loginLog.Remark = rc.Request.Request.UserAgent()
		browserName, browserVersion := ua.Browser()
		loginLog.Browser = browserName + " " + browserVersion
		loginLog.Os = ua.OS()
		loginLog.Platform = ua.Platform()
		loginLog.Username = login.Login
		loginLog.Msg = "登录成功"
		loginLog.CreateBy = fmt.Sprint(login.CreateUid)
		u.LogLogin.Insert(loginLog)
	}()
}

// 2 Auth 用户权限信息 login进入的第一个路由 同时保存登录用户信息
func (u *UserApi) Auth(rc *restfulx.ReqCtx) {
	userName := restfulx.QueryParam(rc, "username")
	biz.NotEmpty(userName, "用户名必传")
	var user entity.ResUsers
	user.Login = userName
	userData, err := u.UserApp.FindProfile(user)
	// global.Log.Infof("验证用户权限-查询用户数据= %v", userData)
	biz.ErrIsNil(err, "用户可能不存在！")
	// role, err := u.RoleApp.FindOne(1)
	// fmt.Printf("Auth 角色=%+v\n", role)
	biz.ErrIsNil(err, "用户所属角色查询失败")
	//前端权限
	// 把角色写死
	permis, _ := u.RoleMenuApp.GetPermis(1)
	menus, _ := u.MenuApp.SelectMenuRole("admin")
	global.Log.Infof("Auth验证用户权限-获取前端菜单路由权限=%v", menus)
	// global.Log.Infof("验证用户权限-上下文数据= %v", rc.LogInfo)

	rc.ResData = vo.AuthVoB{
		User: *userData,
		// Role:        *role,
		Permissions: permis,
		Menus:       Build(*menus),
	}
}

// LogOut 退出登录
func (u *UserApi) LogOut(rc *restfulx.ReqCtx) {
	var loginLog logEntity.LogLogin
	ua := user_agent.New(rc.Request.Request.UserAgent())
	loginLog.Ipaddr = rc.Request.Request.RemoteAddr
	loginLog.LoginTime = time.Now()
	loginLog.Status = "0"
	loginLog.Remark = rc.Request.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	loginLog.Browser = browserName + " " + browserVersion
	loginLog.Os = ua.OS()
	loginLog.Platform = ua.Platform()
	loginLog.Username = rc.LoginAccount.UserName
	loginLog.Msg = "退出成功"
	u.LogLogin.Insert(loginLog)
}

// GetSysUserList 列表数据(条件查询) cg
func (u *UserApi) GetSysUserList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	active := restfulx.QueryParam(rc, "active")
	username := restfulx.QueryParam(rc, "username")
	login := restfulx.QueryParam(rc, "login")
	departmentId := restfulx.QueryParam(rc, "departmentId")
	work_phone := restfulx.QueryParam(rc, "work_phone")

	// organizationId := restfulx.QueryInt(rc, "organizationId", 0)
	var user entity.ResUsersPage
	if active != "" {
		user.Active = kgo.KConv.Str2Bool(active)
	} else {
		user.Active = true
	}
	user.Login = login
	user.Employee.WorkPhone = work_phone
	user.Employee.Name = username
	user.Employee.DepartmentId = kgo.KConv.Str2Int64(departmentId)
	// user.OrganizationId = int64(organizationId)

	// log.Println("发送数据库请求————请求api: ", )
	tt := time.Now()
	list, total, err := u.UserApp.FindListPage(pageNum, pageSize, user)
	// fmt.Printf("后端api获取user查询参数= %+v\n", user)
	biz.ErrIsNil(err, "查询用户分页列表失败")
	log.Println("发送数据库请求————耗时: ", time.Since(tt))
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetSysUserProfile 获取当前登录用户
func (u *UserApi) GetSysUserProfile(rc *restfulx.ReqCtx) {

	sysUser := entity.HrEmployee{}
	sysUser.UserId = rc.LoginAccount.UserId
	user, err := u.UserApp.FindOne(sysUser)
	biz.ErrIsNil(err, "用户可能不存在！")
	//获取角色列表
	// roleList, _ := u.RoleApp.FindList(entity.SysRole{RoleId: rc.LoginAccount.RoleId})
	//岗位列表
	postList, _ := u.HrJobApp.FindList(entity.HrJob{Id: rc.LoginAccount.PostId})
	//获取组织列表
	// organizationList, _ := u.HrDepartmentApp.FindList(entity.HrDepartment{ID: rc.LoginAccount.OrganizationId})
	organizationList, _ := u.HrDepartmentApp.FindOne(rc.LoginAccount.OrganizationId)

	postIds := make([]int64, 0)
	postIds = append(postIds, rc.LoginAccount.PostId)

	roleIds := make([]int64, 0)
	roleIds = append(roleIds, rc.LoginAccount.RoleId)

	rc.ResData = vo.UserProfileVoB{
		Data:    user,
		PostIds: postIds,
		RoleIds: roleIds,
		// Roles:        *roleList,
		Posts:      *postList,
		Department: *organizationList,
	}
}

// todo InsetSysUserAvatar 修改头像
//	func (u *UserApi) InsetSysUserAvatar(rc *restfulx.ReqCtx) {
//		form := rc.Request.Request.MultipartForm
//		files := form.File["upload[]"]
//		guid, _ := kgo.KStr.UuidV4()
//		filPath := "static/uploadfile/" + guid + ".jpg"
//		for _, file := range files {
//			global.Log.Info(file.Filename)
//			// 上传文件至指定目录
//			biz.ErrIsNil(filek.SaveUploadedFile(file, filPath), "保存头像失败")
//		}
//		sysuser := entity.HrEmployee{}
//		sysuser.UserId = rc.LoginAccount.UserId
//		sysuser.Avatar = "/" + filPath
//		sysuser.UpdateBy = rc.LoginAccount.UserName
//		err := u.UserApp.Update(sysuser)
//		biz.ErrIsNil(err, "修改头像失败")
//	}

// SysUserUpdatePwd 修改密码 cg
func (u *UserApi) SysUserUpdatePwd(rc *restfulx.ReqCtx) {
	// var pws entity.ResUsers
	// restfulx.BindJsonAndValid(rc, &pws)
	user := entity.ResUsers{}
	user.ID = rc.LoginAccount.UserId
	data := []int64{user.ID}
	rc.ResData = u.UserApp.SetPwd(data)
	// biz.ErrIsNil(err, "修改密码失败")
}

// GetSysUser 获取用户
func (u *UserApi) GetSysUser(rc *restfulx.ReqCtx) {
	userId := restfulx.PathParamInt(rc, "userId")
	user := entity.HrEmployee{}
	user.UserId = int64(userId)
	result, err := u.UserApp.FindOne(user)
	biz.ErrIsNil(err, "用户可能不存在！")
	// var role entity.SysRole
	// var post entity.HrJob
	// var organization entity.HrDepartment

	// roles, _ := u.RoleApp.FindList(role)
	posts, _ := u.HrJobApp.FindOne(result.JobId)
	// orgs, _ := u.HrDepartmentApp.SelectOrganization(organization)
	orgs, _ := u.HrDepartmentApp.FindOne(result.DepartmentId)

	rc.ResData = vo.UserVoB{
		Data:    result,
		PostIds: result.JobId,
		// RoleIds:       result.RoleIds,
		// Roles:         *roles,
		Posts:     *posts,
		Depatment: *orgs,
	}
}

// GetSysUserInit 获取添加用户角色和职位
func (u *UserApi) GetSysUserInit(rc *restfulx.ReqCtx) {

	// var role entity.SysRole
	// roles, _ := u.RoleApp.FindList(role)
	var post entity.HrJob
	posts, _ := u.HrJobApp.FindList(post)
	rc.ResData = vo.UserRolePostB{
		// Roles: *roles,
		Posts: *posts,
	}
}

// GetUserRolePost 获取添加用户角色和职位
func (u *UserApi) GetUserRolePost(rc *restfulx.ReqCtx) {
	var user entity.HrEmployee
	user.UserId = rc.LoginAccount.UserId

	resData, err := u.UserApp.FindOne(user)
	biz.ErrIsNil(err, "用户可能不存在！")
	// roles := make([]entity.SysRole, 0)
	posts := make([]entity.HrJob, 0)
	// for _, roleId := range strings.Split(resData.RoleIds, ",") {
	// 	ro, err := u.RoleApp.FindOne(kgo.KConv.Str2Int64(roleId))
	// 	if err != nil {
	// 		continue
	// 	}
	// 	roles = append(roles, *ro)
	// }
	for _, postId := range strings.Split(kgo.KConv.Int2Str(resData.JobId), ",") {
		po, err := u.HrJobApp.FindOne(kgo.KConv.Str2Int64(postId))
		if err != nil {
			continue
		}
		posts = append(posts, *po)
	}
	rc.ResData = vo.UserRolePostB{
		// Roles: roles,
		Posts: posts,
	}
}

// InsertSysUser 创建用户，同步创建员工和联系人
func (u *UserApi) InsertSysUser(rc *restfulx.ReqCtx) {
	var sysUser entity.CreateUserDto
	restfulx.BindJsonAndValid(rc, &sysUser)
	// sysUser.CreateBy = rc.LoginAccount.UserName
	result, err := u.UserApp.InsertUser(sysUser)
	biz.ErrIsNil(err, "添加用户失败")
	rc.ResData = result
}

// InsertSysUser 创建员工，同步创建联系人
func (u *UserApi) InsertEmployee(rc *restfulx.ReqCtx) {
	var sysUser entity.HrEmployee
	restfulx.BindJsonAndValid(rc, &sysUser)
	// sysUser.CreateBy = rc.LoginAccount.UserName
	result, err := u.UserApp.InsertEmployee(sysUser)
	biz.ErrIsNil(err, "添加员工失败")
	rc.ResData = result
}

// UpdateEmployee 修改员工数据
func (u *UserApi) UpdateEmployee(rc *restfulx.ReqCtx) {
	var employee entity.HrEmployee
	restfulx.BindJsonAndValid(rc, &employee)
	// sysUser.CreateBy = rc.LoginAccount.UserName
	result := u.UserApp.UpdateEmloyee(employee)
	// biz.ErrIsNil(result.Message, "修改用户失败")
	rc.ResData = result
}

// UpdateSysUser 修改用户数据
func (u *UserApi) UpdateSysUser(rc *restfulx.ReqCtx) {
	var user entity.ResUsers
	restfulx.BindJsonAndValid(rc, &user)
	result := u.UserApp.UpdateUsers(user)
	rc.ResData = result
}

// UpdateSysUserSelf 用户修改数据
func (u *UserApi) UpdateSysUserSelf(rc *restfulx.ReqCtx) {
	var sysUser entity.ResUsers
	restfulx.BindJsonAndValid(rc, &sysUser)
	sysUser.ID = rc.LoginAccount.UserId
	u.UserApp.UpdateUsers(sysUser)
	// biz.ErrIsNil(err, "修改用户数据失败")
}

// UpdateSysUserStu 修改用户状态  用户状态只有0、1，只有正常、归档
// func (u *UserApi) UpdateSysUserStu(rc *restfulx.ReqCtx) {
// 	var sysUser entity.SysUser
// 	restfulx.BindJsonAndValid(rc, &sysUser)
// 	sysUser.CreateBy = rc.LoginAccount.UserName
// 	err := u.UserApp.Update(sysUser)
// 	biz.ErrIsNil(err, "修改用户状态失败")
// }

// DeleteSysUser 删除用户数据
func (u *UserApi) DeleteSysUser(rc *restfulx.ReqCtx) {
	userIds := restfulx.PathParam(rc, "userId")
	u.UserApp.DeleteUser(utils.IdsStrToIdsIntGroup(userIds))
	// biz.ErrIsNil(err, "删除用户失败")
}

// ExportUser 导出用户
func (u *UserApi) ExportUser(rc *restfulx.ReqCtx) {
	filename := restfulx.QueryParam(rc, "filename")
	status := restfulx.QueryParam(rc, "status")
	username := restfulx.QueryParam(rc, "username")
	// phone := restfulx.QueryParam(rc, "phone")

	var user entity.ResUsers
	user.Active = status == "true"
	user.Login = username
	// user.Phone = phone

	list, err := u.UserApp.FindList(user)
	biz.ErrIsNil(err, "用户列表查询失败")
	// 对设置的文件名进行处理
	fileName := utils.GetFileName(global.Conf.Server.ExcelDir, filename)
	utils.InterfaceToExcel(*list, fileName)
	rc.Download(fileName)
}

// Build 构建前端路由
func Build(menus []pentity.SysMenu) []vo.RouterVo {
	equals := func(a string, b string) bool {
		return a == b
	}
	rvs := make([]vo.RouterVo, 0)
	for _, ms := range menus {
		var rv vo.RouterVo
		rv.Name = ms.Path
		rv.Path = ms.Path
		rv.Component = ms.Component
		auth := make([]string, 0)
		if ms.Permission != "" {
			auth = strings.Split(ms.Permission, ",")
		}
		rv.Meta = vo.MetaVo{
			Title:       ms.MenuName,
			IsLink:      ms.IsLink,
			IsHide:      equals("1", ms.IsHide),
			IsKeepAlive: equals("0", ms.IsKeepAlive),
			IsAffix:     equals("0", ms.IsAffix),
			IsIframe:    equals("0", ms.IsIframe),
			Auth:        auth,
			Icon:        ms.Icon,
		}
		rv.Children = Build(ms.Children)
		rvs = append(rvs, rv)
	}

	return rvs
}
