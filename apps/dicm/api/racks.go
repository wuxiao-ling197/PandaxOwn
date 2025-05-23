package api

import (
	"pandax/apps/dicm/entity"
	"pandax/apps/dicm/services"
	"strings"

	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
)

type DcimRackApp struct {
	RackApp services.DcimRackModel
}

// 机柜分页查询列表
func (r *DcimRackApp) GetRackList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	name := restfulx.QueryParam(rc, "name")
	status := restfulx.QueryParam(rc, "status")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	siteId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "siteId"))
	nickname := restfulx.QueryParam(rc, "_name")
	typee := restfulx.QueryParam(rc, "type")
	serial := kgo.KConv.Str2Int32(restfulx.QueryParam(rc, "serial"))
	width := kgo.KConv.Str2Int16(restfulx.QueryParam(rc, "width"))
	u_height := kgo.KConv.Str2Int16(restfulx.QueryParam(rc, "u_height"))
	starting_unit := kgo.KConv.Str2Int16(restfulx.QueryParam(rc, "starting_unit"))
	mounting_depth := kgo.KConv.Str2Int16(restfulx.QueryParam(rc, "mounting_depth"))
	outer_depth := kgo.KConv.Str2Int16(restfulx.QueryParam(rc, "outer_depth"))
	outer_width := kgo.KConv.Str2Int16(restfulx.QueryParam(rc, "outer_width"))
	asset_tag := restfulx.QueryParam(rc, "asset_tag")
	custom_field_data := restfulx.QueryParam(rc, "custom_field_data")
	weight := kgo.KConv.Str2Float64(restfulx.QueryParam(rc, "weight"))
	facilityId := restfulx.QueryParam(rc, "facility_id")
	// facilityId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "facility_id"))
	// facilityId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "facility_id"))
	// facilityId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "facility_id"))
	// facilityId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "facility_id"))
	roleId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "role_id"))
	locationId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "location_id"))
	tenantId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "tenant_id"))
	description := restfulx.QueryParam(rc, "description")
	comments := restfulx.QueryParam(rc, "comments")
	job := entity.DcimRack{Name: name, Id: id, SiteId: siteId, Status: status, NickName: nickname, Weight: weight,
		Type: typee, Serial: serial, Width: width, UHeight: u_height, StartingUnit: starting_unit, MountingDepth: mounting_depth,
		OuterDepth: outer_depth, OuterWidth: outer_width, AssetTag: asset_tag, CustomFieldData: custom_field_data,
		FacilityId: facilityId, RoleId: roleId, LocationId: locationId, TenantId: tenantId, Description: description, Comments: comments,
	}
	list, total, err := r.RackApp.FindListPage(pageNum, pageSize, job)
	biz.ErrIsNil(err, "查询机柜列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// 查询某个数据
func (r *DcimRackApp) GetRackInfo(rc *restfulx.ReqCtx) {
	name := restfulx.PathParam(rc, "name")
	// status := restfulx.PathParam(rc, "status")
	// id := kgo.KConv.Str2Int64(restfulx.PathParam(rc, "id"))
	// siteId := kgo.KConv.Str2Int64(restfulx.PathParam(rc, "siteId"))
	// typee := restfulx.PathParam(rc, "type")
	data, err := r.RackApp.FindOne(entity.DcimRack{Name: name}) //Id: id, SiteId: siteId, Status: status, Type: typee
	biz.ErrIsNil(err, "查询机柜信息失败")
	rc.ResData = data
}

// InsertRack 创建
func (r *DcimRackApp) InsertRack(rc *restfulx.ReqCtx) {
	var dicmRack entity.DcimRack
	restfulx.BindJsonAndValid(rc, &dicmRack)
	_, err := r.RackApp.Insert(dicmRack)
	biz.ErrIsNil(err, "添加机柜失败")
}

// UpdateRack 修改数据
func (r *DcimRackApp) UpdateRack(rc *restfulx.ReqCtx) {
	var dicmRack entity.DcimRack
	restfulx.BindJsonAndValid(rc, &dicmRack)
	err := r.RackApp.Update(dicmRack)
	biz.ErrIsNil(err, "修改机柜信息失败")
}

// 机柜归档
func (r *DcimRackApp) DeleteRack(rc *restfulx.ReqCtx) {
	ids := restfulx.PathParam(rc, "ids")
	data := []int64{}
	if ids != "" {
		sp := strings.Split(ids, ",")
		// log.Println(sp)
		for i := 0; i < len(sp); i++ {
			data = append(data, kgo.KConv.Str2Int64(sp[i]))
		}
	}
	err := r.RackApp.Delete(data)
	biz.ErrIsNil(err, "归档机柜失败")
}

/** 机柜预留 */
// 机柜预留列表
func (r *DcimRackApp) GetRackReserveList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	rackId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "rack_id"))
	tenantId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "tenant_id"))
	userId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "user_id"))
	units := kgo.KConv.Str2Int16(restfulx.QueryParam(rc, "units"))
	description := restfulx.QueryParam(rc, "description")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	comments := restfulx.QueryParam(rc, "comments") //kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "departmentId"))
	cfd := restfulx.QueryParam(rc, "custom_field_data")
	data := entity.DcimRackreservation{Id: id, Units: units, Description: description, Comments: comments, CustomFieldData: cfd}
	if rackId != 0 {
		data.RackId = &rackId
	} else {
		data.RackId = nil
	}
	if tenantId != 0 {
		data.TenantId = &tenantId
	} else {
		data.TenantId = nil
	}
	if userId != 0 {
		data.UserId = &userId
	} else {
		data.UserId = nil
	}
	list, total, err := r.RackApp.FindReserveListPage(pageNum, pageSize, data)
	biz.ErrIsNil(err, "查询机柜预留数据失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// 创建
func (r *DcimRackApp) InsertReservation(rc *restfulx.ReqCtx) {
	var dicmRack entity.DcimRackreservation
	restfulx.BindJsonAndValid(rc, &dicmRack)
	_, err := r.RackApp.InsertReserve(dicmRack)
	biz.ErrIsNil(err, "添加机柜预留实例失败")
}

// 修改数据
func (r *DcimRackApp) UpdateReservation(rc *restfulx.ReqCtx) {
	var dicmRack entity.DcimRackreservation
	restfulx.BindJsonAndValid(rc, &dicmRack)
	err := r.RackApp.UpdateReserve(dicmRack)
	biz.ErrIsNil(err, "修改机柜预留信息失败")
}

// 机柜归档
func (r *DcimRackApp) DeleteReservation(rc *restfulx.ReqCtx) {
	ids := restfulx.PathParam(rc, "ids")
	data := []int64{}
	if ids != "" {
		sp := strings.Split(ids, ",")
		// log.Println(sp)
		for i := 0; i < len(sp); i++ {
			data = append(data, kgo.KConv.Str2Int64(sp[i]))
		}
	}
	err := r.RackApp.DeleteReserve(data)
	biz.ErrIsNil(err, "归档机柜预留失败")
}

/** 机柜角色 */
// 机柜用途列表
func (r *DcimRackApp) GetRackRoleList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	name := restfulx.QueryParam(rc, "name")
	color := restfulx.QueryParam(rc, "color")
	description := restfulx.QueryParam(rc, "description")
	cfd := restfulx.QueryParam(rc, "custom_field_data")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	slug := restfulx.QueryParam(rc, "slug") //kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "departmentId"))
	data := entity.DcimRackrole{Name: name, Id: id, Color: color, Slug: slug, Description: description, CustomFieldData: cfd}
	list, total, err := r.RackApp.FindRoleListPage(pageNum, pageSize, data)
	biz.ErrIsNil(err, "查询机柜类型数据失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// 创建
func (r *DcimRackApp) InsertRackRole(rc *restfulx.ReqCtx) {
	var dicmRack entity.DcimRackrole
	restfulx.BindJsonAndValid(rc, &dicmRack)
	_, err := r.RackApp.InsertRackRole(dicmRack)
	biz.ErrIsNil(err, "添加机柜角色实例失败")
}

// 修改数据
func (r *DcimRackApp) UpdateRackRole(rc *restfulx.ReqCtx) {
	var dicmRack entity.DcimRackrole
	restfulx.BindJsonAndValid(rc, &dicmRack)
	err := r.RackApp.UpdateRackRole(dicmRack)
	biz.ErrIsNil(err, "修改机柜角色信息失败")
}

// 机柜归档
func (r *DcimRackApp) DeleteRackRole(rc *restfulx.ReqCtx) {
	ids := restfulx.PathParam(rc, "ids")
	data := []int64{}
	if ids != "" {
		sp := strings.Split(ids, ",")
		// log.Println(sp)
		for i := 0; i < len(sp); i++ {
			data = append(data, kgo.KConv.Str2Int64(sp[i]))
		}
	}
	err := r.RackApp.DeleteRackRole(data)
	biz.ErrIsNil(err, "归档机柜角色失败")
}
