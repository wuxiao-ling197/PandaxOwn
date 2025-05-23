package api

import (
	"fmt"
	"pandax/apps/dicm/entity"
	"pandax/apps/dicm/services"
	"strings"

	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
)

type DcimSiteApp struct {
	SiteApp services.DcimSiteModel
}

// 分页查询列表 对于无值的query项转换结果为Id:0【int64,Str2Int64】 Name: PhysicalAddress:716【查询字段】  Comments: GroupId:<nil>【*int64,外键，有处理】 Group:<nil>
func (s *DcimSiteApp) GetSiteList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	name := restfulx.QueryParam(rc, "name")
	nickname := restfulx.QueryParam(rc, "_name")
	slug := restfulx.QueryParam(rc, "slug")
	physical_address := restfulx.QueryParam(rc, "physical_address")
	shipping_address := restfulx.QueryParam(rc, "shipping_address")
	custom_field_data := restfulx.QueryParam(rc, "custom_field_data")
	tz := restfulx.QueryParam(rc, "time_zone")
	// created := restfulx.QueryParam(rc, "created")
	// updated := restfulx.QueryParam(rc, "last_updated")
	// deleted := restfulx.QueryParam(rc, "deleted")
	status := restfulx.QueryParam(rc, "status")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	groupId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "group_id"))
	regionId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "region_id"))
	tenantId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "tenant_id"))
	site := entity.DcimSite{
		Name: name, NickNameName: nickname, Id: id, Status: status, Slug: slug,
		CustomFieldData: custom_field_data, PhysicalAddress: physical_address,
		ShippingAddress: shipping_address, TimeZone: tz}
	// Created: created,LastUpdated: updated,Deleted: deleted}
	if groupId != 0 {
		site.GroupId = &groupId
	} else {
		site.GroupId = nil
	}
	if regionId != 0 {
		site.RegionId = &regionId
	} else {
		site.RegionId = nil
	}
	if tenantId != 0 {
		site.TenantId = &tenantId
	} else {
		site.TenantId = nil
	}
	/**
	site={Deleted:0001-01-01 00:00:00 +0000 UTC Created:0001-01-01 00:00:00 +0000 UTC LastUpdated:0001-01-01 00:00:00 +0000 UTC
	CustomFieldData: Id:0 Name: NickNameName: Slug: Status: Facility: TimeZone: Description:
	PhysicalAddress:716 ShippingAddress: Latitude: Longitude: Comments: GroupId:<nil> Group:<nil>
	RegionId:<nil> Region:<nil> TenantId:<nil> Tenant:<nil>}
	*/
	list, total, err := s.SiteApp.FindListPage(pageNum, pageSize, site)
	biz.ErrIsNil(err, "查询站点列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// 通过名称查询站点数据详情
func (s *DcimSiteApp) GetSiteInfo(rc *restfulx.ReqCtx) {
	name := restfulx.PathParam(rc, "name")
	// status := restfulx.PathParam(rc, "status")
	// id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	// groupId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "groupId"))
	// regionId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "regionId"))
	site := entity.DcimSite{Name: name}
	// , Id: id, Status: status}
	// if groupId != 0 {
	// 	site.GroupId = &groupId
	// } else {
	// 	site.GroupId = nil
	// }
	// if regionId != 0 {
	// 	site.RegionId = &regionId
	// } else {
	// 	site.RegionId = nil
	// }
	data, err := s.SiteApp.FindOne(site)
	biz.ErrIsNil(err, "查询站点信息失败")
	rc.ResData = data
}

func (s *DcimSiteApp) InsertSite(rc *restfulx.ReqCtx) {
	var group entity.DcimSite
	// restfulx.BindJsonAndValid(rc, &tenant)
	if err := restfulx.BindJsonAndValid(rc, &group); err != nil {
		fmt.Printf("BindJsonAndValid error: %+v\n", err)

		biz.ErrIsNil(err, "请求参数绑定或校验失败") // Use a more specific message or let biz.ErrIsNil handle it
		return                           // Important to return after error
	}
	// if tenant.ParentId.Int64 != 0 {
	// 	tenant.ParentId = sql.NullInt64{Int64: tenant.ParentId.Int64, Valid: true}
	// }
	_, err := s.SiteApp.Insert(group)
	biz.ErrIsNil(err, "添加站点实例失败")
}

// UpdateTenant 修改数据
func (s *DcimSiteApp) UpdateSite(rc *restfulx.ReqCtx) {
	var group entity.DcimSite
	// restfulx.BindJsonAndValid(rc, &group)
	// if tenant.ParentId.Int64 != 0 {
	// 	tenant.ParentId = sql.NullInt64{Int64: tenant.ParentId.Int64, Valid: true}
	// }
	if err := restfulx.BindJsonAndValid(rc, &group); err != nil {
		fmt.Printf("BindJsonAndValid error: %+v\n", err)

		biz.ErrIsNil(err, "请求参数绑定或校验失败") // Use a more specific message or let biz.ErrIsNil handle it
		return                           // Important to return after error
	}
	err := s.SiteApp.Update(group)
	biz.ErrIsNil(err, "修改站点信息失败")
}

func (s *DcimSiteApp) DeleteSite(rc *restfulx.ReqCtx) {
	ids := restfulx.PathParam(rc, "ids")
	data := []int64{}
	if ids != "" {
		sp := strings.Split(ids, ",")
		// log.Println(sp)
		for i := 0; i < len(sp); i++ {
			data = append(data, kgo.KConv.Str2Int64(sp[i]))
		}
	}
	err := s.SiteApp.Delete(data)
	biz.ErrIsNil(err, "归档站点失败")
}

/** 站点组 **/
// 分页查询列表
func (s *DcimSiteApp) GetSiteGroupList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	name := restfulx.QueryParam(rc, "name")
	slug := restfulx.QueryParam(rc, "slug")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	parentId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "parentId"))
	level := kgo.KConv.Str2Int32(restfulx.QueryParam(rc, "level"))
	lft := kgo.KConv.Str2Int32(restfulx.QueryParam(rc, "lft"))
	rght := kgo.KConv.Str2Int32(restfulx.QueryParam(rc, "rght"))
	treeId := kgo.KConv.Str2Int32(restfulx.QueryParam(rc, "treeId"))
	description := restfulx.QueryParam(rc, "description")
	data := entity.DcimSitegroup{Name: name, Id: id, ParentId: &parentId, Slug: slug, Level: level, Lft: lft, Rght: rght, TreeId: treeId, Description: description}
	if parentId != 0 {
		data.ParentId = &parentId
	} else {
		data.ParentId = nil
	}
	list, total, err := s.SiteApp.FindListGroupPage(pageNum, pageSize, data)
	biz.ErrIsNil(err, "查询站点组列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (s *DcimSiteApp) JoinSiteGroup(rc *restfulx.ReqCtx) {
	// name := restfulx.PathParam(rc, "name")
	var data entity.Site2Group
	restfulx.BindJsonAndValid(rc, &data)
	err := s.SiteApp.JoinGroup(data.SiteIds, data.GroupId)
	biz.ErrIsNil(err, "站点组添加站点失败")
}

func (s *DcimSiteApp) GetGroupTree(rc *restfulx.ReqCtx) {
	tree, err := s.SiteApp.GetSiteStructrue()
	biz.ErrIsNil(err, "获取站点组层级结构失败")
	rc.ResData = tree
}

func (s *DcimSiteApp) InsertSiteGroup(rc *restfulx.ReqCtx) {
	var group entity.DcimSitegroup
	// restfulx.BindJsonAndValid(rc, &tenant)
	if err := restfulx.BindJsonAndValid(rc, &group); err != nil {
		fmt.Printf("BindJsonAndValid error: %+v\n", err)

		biz.ErrIsNil(err, "请求参数绑定或校验失败") // Use a more specific message or let biz.ErrIsNil handle it
		return                           // Important to return after error
	}
	// if tenant.ParentId.Int64 != 0 {
	// 	tenant.ParentId = sql.NullInt64{Int64: tenant.ParentId.Int64, Valid: true}
	// }
	_, err := s.SiteApp.InsertGroup(group)
	biz.ErrIsNil(err, "添加站点组实例失败")
}

// UpdateTenant 修改数据
func (s *DcimSiteApp) UpdateSiteGroup(rc *restfulx.ReqCtx) {
	var group entity.DcimSitegroup
	// restfulx.BindJsonAndValid(rc, &group)
	// if tenant.ParentId.Int64 != 0 {
	// 	tenant.ParentId = sql.NullInt64{Int64: tenant.ParentId.Int64, Valid: true}
	// }
	if err := restfulx.BindJsonAndValid(rc, &group); err != nil {
		fmt.Printf("BindJsonAndValid error: %+v\n", err)

		biz.ErrIsNil(err, "请求参数绑定或校验失败") // Use a more specific message or let biz.ErrIsNil handle it
		return                           // Important to return after error
	}
	err := s.SiteApp.UpdateGroup(group)
	biz.ErrIsNil(err, "修改站点组信息失败")
}

func (s *DcimSiteApp) DeleteSiteGroup(rc *restfulx.ReqCtx) {
	ids := restfulx.PathParam(rc, "ids")
	data := []int64{}
	if ids != "" {
		sp := strings.Split(ids, ",")
		// log.Println(sp)
		for i := 0; i < len(sp); i++ {
			data = append(data, kgo.KConv.Str2Int64(sp[i]))
		}
	}
	err := s.SiteApp.DeleteGroup(data)
	biz.ErrIsNil(err, "归档站点组失败")
}

// 查询某个数据
// func (s *DcimSiteApp) GetSiteGroupInfo(rc *restfulx.ReqCtx) {
// 	name := restfulx.PathParam(rc, "name")
// 	status := restfulx.PathParam(rc, "status")
// 	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
// 	groupId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "groupId"))
// 	regionId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "regionId"))
// 	data, err := s.SiteApp.FindOne(entity.DcimSite{Name: name, Id: id, GroupId: groupId, Status: status, RegionId: regionId})
// 	biz.ErrIsNil(err, "查询站点组信息失败")
// 	rc.ResData = data
// }
