package api

import (
	"pandax/apps/dicm/entity"
	"pandax/apps/dicm/services"

	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
)

type DcimSiteApp struct {
	SiteApp services.DcimSiteModel
}

// 分页查询列表
func (s *DcimSiteApp) GetSiteList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	name := restfulx.QueryParam(rc, "name")
	status := restfulx.QueryParam(rc, "status")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	groupId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "groupId"))
	regionId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "regionId"))
	site := entity.DcimSite{Name: name, Id: id, GroupId: groupId, Status: status, RegionId: regionId}
	list, total, err := s.SiteApp.FindListPage(pageNum, pageSize, site)
	biz.ErrIsNil(err, "查询机柜列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// 查询某个数据
func (s *DcimSiteApp) GetSiteInfo(rc *restfulx.ReqCtx) {
	name := restfulx.PathParam(rc, "name")
	status := restfulx.PathParam(rc, "status")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	groupId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "groupId"))
	regionId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "regionId"))
	data, err := s.SiteApp.FindOne(entity.DcimSite{Name: name, Id: id, GroupId: groupId, Status: status, RegionId: regionId})
	biz.ErrIsNil(err, "查询机柜信息失败")
	rc.ResData = data
}
