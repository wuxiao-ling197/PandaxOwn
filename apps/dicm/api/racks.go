package api

import (
	"pandax/apps/dicm/entity"
	"pandax/apps/dicm/services"

	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
)

type DcimRackApp struct {
	RackApp services.DcimRackModel
}

// 分页查询列表
func (r *DcimRackApp) GetRackList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	name := restfulx.QueryParam(rc, "name")
	status := restfulx.QueryParam(rc, "status")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	siteId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "siteId"))
	job := entity.DcimRack{Name: name, Id: id, SiteId: siteId, Status: status}
	list, total, err := r.RackApp.FindListPage(pageNum, pageSize, job)
	biz.ErrIsNil(err, "查询机柜列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// 机柜预留列表
func (r *DcimRackApp) GetRackReserveList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	rackId := restfulx.QueryParam(rc, "rackId")
	tenantId := restfulx.QueryParam(rc, "tenantId")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	userId := restfulx.QueryParam(rc, "userId") //kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "departmentId"))
	data := entity.DcimRackreservation{RackId: rackId, Id: id, TenantId: tenantId, UserId: userId}
	list, total, err := r.RackApp.FindReserveListPage(pageNum, pageSize, data)
	biz.ErrIsNil(err, "查询机柜预留数据失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// 机柜用途列表
func (r *DcimRackApp) GetRackRoleList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	name := restfulx.QueryParam(rc, "name")
	color := restfulx.QueryParam(rc, "color")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	slug := restfulx.QueryParam(rc, "slug") //kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "departmentId"))
	data := entity.DcimRackrole{Name: name, Id: id, Color: color, Slug: slug}
	list, total, err := r.RackApp.FindRoleListPage(pageNum, pageSize, data)
	biz.ErrIsNil(err, "查询机柜类型数据失败")
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
	biz.ErrIsNil(err, "添加用户失败")
}

// UpdateRack 修改数据
func (r *DcimRackApp) UpdateRack(rc *restfulx.ReqCtx) {
	var dicmRack entity.DcimRack
	restfulx.BindJsonAndValid(rc, &dicmRack)
	err := r.RackApp.Update(dicmRack)
	biz.ErrIsNil(err, "修改用户失败")
}
