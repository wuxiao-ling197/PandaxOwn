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

type DcimLocationApp struct {
	LocationApp services.DcimLocationModel
}

/** Location **/
// 分页查询列表
func (s *DcimLocationApp) GetLocationList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	name := restfulx.QueryParam(rc, "name")
	slug := restfulx.QueryParam(rc, "slug")
	status := restfulx.QueryParam(rc, "status")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	level := kgo.KConv.Str2Int32(restfulx.QueryParam(rc, "level"))
	treeId := kgo.KConv.Str2Int32(restfulx.QueryParam(rc, "treeId"))
	tenantId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "tenantId"))
	parentId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "parentId"))
	siteId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "siteId"))
	// site := entity.DcimLocation{Name: name, Id: id, ParentId: &parentId, SiteId: siteId, Status: status, TreeId: treeId, Level: level, TenantId: tenantId, Slug: slug}
	site := entity.DcimLocation{Name: name, Id: id, ParentId: &parentId, SiteId: siteId, Status: status, TreeId: treeId, Level: level, TenantId: tenantId, Slug: slug}
	if parentId != 0 {
		site.ParentId = &parentId
	} else {
		site.ParentId = nil
	}
	list, total, err := s.LocationApp.FindListPage(pageNum, pageSize, site)
	biz.ErrIsNil(err, "查询物理位置列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// 查询某个数据
func (s *DcimLocationApp) GetLocationInfo(rc *restfulx.ReqCtx) {
	name := restfulx.PathParam(rc, "name")
	status := restfulx.PathParam(rc, "status")
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	slug := restfulx.QueryParam(rc, "slug")
	level := kgo.KConv.Str2Int32(restfulx.QueryParam(rc, "level"))
	treeId := kgo.KConv.Str2Int32(restfulx.QueryParam(rc, "treeId"))
	tenantId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "tenantId"))
	parentId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "parentId"))
	siteId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "siteId"))
	// data, err := s.LocationApp.FindOne(entity.DcimLocation{Name: name, Id: id, Slug: slug, Level: level, TreeId: treeId, TenantId: tenantId, SiteId: siteId, Status: status, ParentId: &parentId})
	site := entity.DcimLocation{Name: name, Id: id, ParentId: &parentId, SiteId: siteId, Status: status, TreeId: treeId, Level: level, TenantId: tenantId, Slug: slug}
	if parentId != 0 {
		site.ParentId = &parentId
	} else {
		site.ParentId = nil
	}
	data, err := s.LocationApp.FindOne(site)
	biz.ErrIsNil(err, "查询物理位置信息失败")
	rc.ResData = data
}

// 添加实例
func (s *DcimLocationApp) InsertLocation(rc *restfulx.ReqCtx) {
	var data entity.DcimLocation
	// restfulx.BindJsonAndValid(rc, &tenant)  //将解析json，即要求前端传参类型为json
	// if tenant.GroupId.Int64 != 0 {
	// 	tenant.GroupId = sql.NullInt64{Int64: tenant.GroupId.Int64, Valid: true}
	// }
	if err := restfulx.BindJsonAndValid(rc, &data); err != nil {
		fmt.Printf("BindJsonAndValid error: %+v\n", err)

		biz.ErrIsNil(err, "请求参数绑定或校验失败") // Use a more specific message or let biz.ErrIsNil handle it
		return                           // Important to return after error
	}
	fmt.Printf("解析参数：%+v\n", data)
	_, err := s.LocationApp.Insert(data)
	biz.ErrIsNil(err, "添加物理位置失败")
}

// 修改实例
func (s *DcimLocationApp) UpdateLocation(rc *restfulx.ReqCtx) {
	var data entity.DcimLocation
	// restfulx.BindJsonAndValid(rc, &tenant)  //将解析json，即要求前端传参类型为json
	// if tenant.GroupId.Int64 != 0 {
	// 	tenant.GroupId = sql.NullInt64{Int64: tenant.GroupId.Int64, Valid: true}
	// }
	if err := restfulx.BindJsonAndValid(rc, &data); err != nil {
		fmt.Printf("BindJsonAndValid error: %+v\n", err)

		biz.ErrIsNil(err, "请求参数绑定或校验失败") // Use a more specific message or let biz.ErrIsNil handle it
		return                           // Important to return after error
	}
	// fmt.Printf("解析参数：%+v\n", tenant)
	err := s.LocationApp.Update(data)
	biz.ErrIsNil(err, "修改物理位置失败")
}

// 归档
func (s *DcimLocationApp) DeleteLocation(rc *restfulx.ReqCtx) {
	ids := restfulx.PathParam(rc, "ids")
	data := []int64{}
	if ids != "" {
		sp := strings.Split(ids, ",")
		log.Println(sp)
		for i := 0; i < len(sp); i++ {
			data = append(data, kgo.KConv.Str2Int64(sp[i]))
		}
	}
	err := s.LocationApp.Delete(data)
	biz.ErrIsNil(err, "归档失败")
}

func (s *DcimLocationApp) GetLocationTree(rc *restfulx.ReqCtx) {
	tree, err := s.LocationApp.GetLocationStructrue()
	biz.ErrIsNil(err, "获取location结构失败")
	rc.ResData = tree
}

/** Region **/
// 分页查询列表
func (s *DcimLocationApp) GetRegionList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	id := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "id"))
	name := restfulx.QueryParam(rc, "name")
	slug := restfulx.QueryParam(rc, "slug")
	level := kgo.KConv.Str2Int32(restfulx.QueryParam(rc, "level"))
	cfd := restfulx.QueryParam(rc, "custom_field_data")
	parentId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "parentId"))
	treeId := kgo.KConv.Str2Int32(restfulx.QueryParam(rc, "treeId"))
	description := restfulx.QueryParam(rc, "description")
	site := entity.DcimRegion{Name: name, Id: id, Description: description, CustomFieldData: cfd, TreeId: treeId, Level: level, Slug: slug}
	if parentId != 0 {
		site.ParentId = &parentId
	} else {
		site.ParentId = nil
	}
	list, total, err := s.LocationApp.FindRegionListPage(pageNum, pageSize, site)
	biz.ErrIsNil(err, "查询地区列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// 查询某个数据
func (s *DcimLocationApp) GetRegionInfo(rc *restfulx.ReqCtx) {
	name := restfulx.PathParam(rc, "name")
	site := entity.DcimRegion{Name: name}
	data, err := s.LocationApp.FindRegionOne(site)
	biz.ErrIsNil(err, "查询地区信息失败")
	rc.ResData = data
}

// 添加实例
func (s *DcimLocationApp) InsertRegion(rc *restfulx.ReqCtx) {
	var data entity.DcimRegion
	if err := restfulx.BindJsonAndValid(rc, &data); err != nil {
		fmt.Printf("BindJsonAndValid error: %+v\n", err)
		biz.ErrIsNil(err, "请求参数绑定或校验失败")
		return
	}
	fmt.Printf("解析参数：%+v\n", data)
	_, err := s.LocationApp.InsertRegion(data)
	biz.ErrIsNil(err, "添加地区失败")
}

// 修改实例
func (s *DcimLocationApp) UpdateRegion(rc *restfulx.ReqCtx) {
	var data entity.DcimRegion
	if err := restfulx.BindJsonAndValid(rc, &data); err != nil {
		fmt.Printf("BindJsonAndValid error: %+v\n", err)

		biz.ErrIsNil(err, "请求参数绑定或校验失败")
		return
	}
	err := s.LocationApp.UpdateRegion(data)
	biz.ErrIsNil(err, "修改地区失败")
}

// 归档
func (s *DcimLocationApp) DeleteRegion(rc *restfulx.ReqCtx) {
	ids := restfulx.PathParam(rc, "ids")
	data := []int64{}
	if ids != "" {
		sp := strings.Split(ids, ",")
		log.Println(sp)
		for i := 0; i < len(sp); i++ {
			data = append(data, kgo.KConv.Str2Int64(sp[i]))
		}
	}
	err := s.LocationApp.DeleteRegion(data)
	biz.ErrIsNil(err, "归档失败")
}

func (s *DcimLocationApp) GetRegionTree(rc *restfulx.ReqCtx) {
	tree, err := s.LocationApp.GetRegionStructrue()
	biz.ErrIsNil(err, "获取region结构失败")
	rc.ResData = tree
}
