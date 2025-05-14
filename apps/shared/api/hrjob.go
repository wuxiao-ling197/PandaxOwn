package api

import (
	"fmt"
	"pandax/apps/shared/entity"
	"pandax/apps/shared/services"
	pservices "pandax/apps/system/services"

	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
)

type HrJobApp struct {
	PostApp services.HrJobModel
	UserApp services.ResUsersModel
	RoleApp pservices.SysRoleModel
}

// GetPostList 职位列表数据
func (p *HrJobApp) GetPostList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	jobName := restfulx.QueryParam(rc, "jobName")
	jobCode := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "jobCode"))
	departmentId := kgo.KConv.Str2Int64(restfulx.QueryParam(rc, "departmentId"))
	job := entity.HrJob{Name: jobName, Id: jobCode, DepartmentId: departmentId}
	list, total, err := p.PostApp.FindListPage(pageNum, pageSize, job)
	biz.ErrIsNil(err, "查询岗位列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetPost 获取职位
func (p *HrJobApp) GetPost(rc *restfulx.ReqCtx) {
	jobId := restfulx.PathParamInt(rc, "jobId")
	data, err := p.PostApp.FindOne(int64(jobId))
	biz.ErrIsNil(err, "查询岗位失败")
	rc.ResData = data
}

// InsertPost 添加职位
func (p *HrJobApp) InsertPost(rc *restfulx.ReqCtx) {
	var post entity.HrJob
	restfulx.BindJsonAndValid(rc, &post)
	post.CreateUid = rc.LoginAccount.UserId
	result, err := p.PostApp.Insert(post)
	biz.ErrIsNil(err, "添加部门失败")
	rc.ResData = result
}

// UpdatePost 修改职位
func (p *HrJobApp) UpdatePost(rc *restfulx.ReqCtx) {
	var post entity.HrJob
	restfulx.BindJsonAndValid(rc, &post)

	post.CreateUid = rc.LoginAccount.UserId
	result := p.PostApp.Update(post)
	rc.ResData = result.Data
	// biz.ErrIsNil(err, "修改部门失败")
}

// PublishPost 发布招聘
func (p *HrJobApp) PublishPost(rc *restfulx.ReqCtx) {
	jobId := restfulx.PathParamInt(rc, "id")
	fmt.Println(jobId)
	result := p.PostApp.Publish(int64(jobId))
	rc.ResData = result
}

// DeletePost 删除职位
// func (p *HrJobApp) DeletePost(rc *restfulx.ReqCtx) {
// 	jobId := restfulx.PathParam(rc, "jobId")
// 	postIds := utils.IdsStrToIdsIntGroup(jobId)

// 	deList := make([]int64, 0)
// for _, id := range postIds {
// 	user := entity.HrJob{}
// 	user.PostId = id
// 	list, err := p.UserApp.FindList(user)
// 	if err != nil {
// 		continue
// 	}
// 	if len(*list) == 0 {
// 		deList = append(deList, id)
// 	} else {
// 		global.Log.Info(fmt.Sprintf("dictId: %d 存在岗位绑定用户无法删除", id))
// 	}
// }
// if len(deList) == 0 {
// 	biz.ErrIsNil(errors.New("所有岗位都已绑定用户，无法删除"), "所有岗位都已绑定用户，无法删除")
// }
// err := p.PostApp.Delete(deList)
// biz.ErrIsNil(err, "删除部门失败")
// }
