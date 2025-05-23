package api

import (
	"fmt"
	"log"
	"pandax/apps/dicm/entity"
	"pandax/apps/dicm/services"
	"strings"

	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
)

type TenantApp struct {
	TApp services.TenantModel
}

// 分页查询列表 cg
func (s *TenantApp) GetTenantList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	name := restfulx.QueryParam(rc, "name")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	groupId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "groupId"))
	slug := restfulx.QueryParam(rc, "slug")
	// site := entity.TenancyTenant{Name: name, Id: id, GroupId: sql.NullInt64{Int64: groupId, Valid: true}, Slug: slug}
	site := entity.TenancyTenant{Name: name, Id: id, GroupId: groupId, Slug: slug}
	list, total, err := s.TApp.FindListPage(pageNum, pageSize, site)
	biz.ErrIsNil(err, "查询租户列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// 查询某个数据 cg
func (s *TenantApp) GetTenantInfo(rc *restfulx.ReqCtx) {
	name := restfulx.PathParam(rc, "name")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	groupId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "groupId"))
	slug := restfulx.QueryParam(rc, "slug")
	// data, err := s.TApp.FindOne(entity.TenancyTenant{Name: name, Id: id, GroupId: sql.NullInt64{Int64: groupId, Valid: true}, Slug: slug})
	data, err := s.TApp.FindOne(entity.TenancyTenant{Name: name, Id: id, GroupId: groupId, Slug: slug})

	biz.ErrIsNil(err, "查询租户信息失败")
	rc.ResData = data
}

// InsertTenant 创建  cg
func (s *TenantApp) InsertTenant(rc *restfulx.ReqCtx) {
	var tenant entity.TenancyTenant
	// restfulx.BindJsonAndValid(rc, &tenant)  //将解析json，即要求前端传参类型为json
	// if tenant.GroupId.Int64 != 0 {
	// 	tenant.GroupId = sql.NullInt64{Int64: tenant.GroupId.Int64, Valid: true}
	// }
	if err := restfulx.BindJsonAndValid(rc, &tenant); err != nil {
		fmt.Printf("BindJsonAndValid error: %+v\n", err)

		biz.ErrIsNil(err, "请求参数绑定或校验失败") // Use a more specific message or let biz.ErrIsNil handle it
		return                           // Important to return after error
	}
	// fmt.Printf("解析参数：%+v\n", tenant)
	_, err := s.TApp.Insert(tenant)
	biz.ErrIsNil(err, "添加租户失败")
}

// UpdateTenant 修改数据 cg
func (s *TenantApp) UpdateTenant(rc *restfulx.ReqCtx) {
	var tenant entity.TenancyTenant
	restfulx.BindJsonAndValid(rc, &tenant)
	// if tenant.GroupId.Int64 != 0 {
	// 	tenant.GroupId = sql.NullInt64{Int64: tenant.GroupId.Int64, Valid: true}
	// }
	err := s.TApp.Update(tenant)
	biz.ErrIsNil(err, "修改用户失败")
}

func (s *TenantApp) DeleteTenant(rc *restfulx.ReqCtx) {
	ids := restfulx.PathParam(rc, "ids")
	data := []int64{}
	if ids != "" {
		sp := strings.Split(ids, ",")
		for i := 0; i < len(sp); i++ {
			data = append(data, kgo.KConv.Str2Int64(sp[i]))
		}
	}
	err := s.TApp.Delete(data)
	biz.ErrIsNil(err, "删除租户失败")
}

// 租户加入租户组 cg
func (s *TenantApp) JoinTenantGroup(rc *restfulx.ReqCtx) {
	// name := restfulx.PathParam(rc, "name")
	var data entity.Own2Group
	restfulx.BindJsonAndValid(rc, &data)
	err := s.TApp.JoinGroup(data.TenantIds, data.GroupId)
	biz.ErrIsNil(err, "租户组添加租户失败")
}

/** 租户组 */
func (s *TenantApp) GetTenantGroupList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	name := restfulx.QueryParam(rc, "name")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	parentId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "parentId"))
	slug := restfulx.QueryParam(rc, "slug")
	level := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "level"))
	treeId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "treeId"))
	// site := entity.TenancyTenantgroup{Name: name, Id: id, ParentId: sql.NullInt64{Int64: parentId, Valid: true}, Slug: slug, Level: level, TreeId: treeId}
	site := entity.TenancyTenantgroup{Name: name, Id: id, Slug: slug, Level: level, TreeId: treeId}
	if parentId != 0 {
		site.ParentId = &parentId
	} else {
		site.ParentId = nil
	}
	list, total, err := s.TApp.FindListGroupPage(pageNum, pageSize, site)
	biz.ErrIsNil(err, "查询租户组列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (s *TenantApp) GetTenantGroupInfo(rc *restfulx.ReqCtx) {
	name := restfulx.PathParam(rc, "name")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	parentId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "parentId"))
	slug := restfulx.QueryParam(rc, "slug")
	level := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "level"))
	treeId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "treeId"))
	// site := entity.TenancyTenantgroup{Name: name, Id: id, ParentId: sql.NullInt64{Int64: parentId, Valid: true}, Slug: slug, Level: level, TreeId: treeId}
	site := entity.TenancyTenantgroup{Name: name, Id: id, Slug: slug, Level: level, TreeId: treeId}
	if parentId != 0 {
		site.ParentId = &parentId
	} else {
		site.ParentId = nil
	}
	data, err := s.TApp.FindGroupOne(site)
	biz.ErrIsNil(err, "查询租户组信息失败")
	rc.ResData = data
}

// InsertTenantGroup 创建
func (s *TenantApp) InsertTenantGroup(rc *restfulx.ReqCtx) {
	var tenant entity.TenancyTenantgroup
	// restfulx.BindJsonAndValid(rc, &tenant)
	if err := restfulx.BindJsonAndValid(rc, &tenant); err != nil {
		fmt.Printf("BindJsonAndValid error: %+v\n", err)

		biz.ErrIsNil(err, "请求参数绑定或校验失败") // Use a more specific message or let biz.ErrIsNil handle it
		return                           // Important to return after error
	}
	// fmt.Printf("解析参数：%+v\n", tenant)
	// if tenant.ParentId.Int64 != 0 {
	// 	tenant.ParentId = sql.NullInt64{Int64: tenant.ParentId.Int64, Valid: true}
	// }
	_, err := s.TApp.InsertGroup(tenant)
	biz.ErrIsNil(err, "添加租户组实例失败")
}

// UpdateTenant 修改数据
func (s *TenantApp) UpdateTenantGroup(rc *restfulx.ReqCtx) {
	var tenant entity.TenancyTenantgroup
	// restfulx.BindJsonAndValid(rc, &tenant)
	// if tenant.ParentId.Int64 != 0 {
	// 	tenant.ParentId = sql.NullInt64{Int64: tenant.ParentId.Int64, Valid: true}
	// }
	if err := restfulx.BindJsonAndValid(rc, &tenant); err != nil {
		fmt.Printf("BindJsonAndValid error: %+v\n", err)

		biz.ErrIsNil(err, "请求参数绑定或校验失败") // Use a more specific message or let biz.ErrIsNil handle it
		return                           // Important to return after error
	}
	err := s.TApp.UpdateGroup(tenant)
	biz.ErrIsNil(err, "修改租户组信息失败")
}

func (s *TenantApp) DeleteTenantGroup(rc *restfulx.ReqCtx) {
	ids := restfulx.PathParam(rc, "ids")
	data := []int64{}
	if ids != "" {
		sp := strings.Split(ids, ",")
		log.Println(sp)
		for i := 0; i < len(sp); i++ {
			data = append(data, kgo.KConv.Str2Int64(sp[i]))
		}
	}
	err := s.TApp.DeleteGroup(data)
	biz.ErrIsNil(err, "删除租户组失败")
}

func (s *TenantApp) GetGroupTree(rc *restfulx.ReqCtx) {
	tree, err := s.TApp.GetGroupStructrue()
	biz.ErrIsNil(err, "获取location层级结构失败")
	rc.ResData = tree
}
