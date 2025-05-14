package services

import (
	"encoding/json"
	"fmt"
	"pandax/apps/shared/entity"
	odoorpc "pandax/pkg/device_rpc"
	"pandax/pkg/global"
)

type (
	HrJobModel interface {
		Insert(data entity.HrJob) (odoorpc.ResultData, error)
		FindOne(postId int64) (*entity.HrJob, error)
		FindListPage(page, pageSize int, data entity.HrJob) (*[]entity.HrJob, int64, error)
		FindList(data entity.HrJob) (*[]entity.HrJob, error)
		Update(data entity.HrJob) odoorpc.ResultData
		Publish(id int64) odoorpc.ResultData
		Delete(postId []int64) odoorpc.ResultData
	}

	hrJobModelImpl struct {
		table string
	}
)

var HrJobModelDao HrJobModel = &hrJobModelImpl{
	table: `hr_job`,
}

func (m *hrJobModelImpl) Insert(data entity.HrJob) (odoorpc.ResultData, error) {
	// err := global.Db.Table(m.table).Create(&data).Error
	// return &data, err
	var create map[string]interface{}
	createData, _ := json.Marshal(data)
	_ = json.Unmarshal(createData, &create)
	result, err := odoorpc.Create(create, "hr.job")
	return result, err
}

func (m *hrJobModelImpl) FindOne(postId int64) (*entity.HrJob, error) {
	resData := new(entity.HrJob)
	err := global.HrDb.Table(m.table).Where("id = ?", postId).First(resData).Error
	return resData, err
}

func (m *hrJobModelImpl) FindListPage(page, pageSize int, data entity.HrJob) (*[]entity.HrJob, int64, error) {
	list := make([]entity.HrJob, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.HrDb.Table("hr_job")
	// 此处填写 where参数判断
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.Name != "" {
		db = db.Where("name::TEXT like ?", "%"+data.Name+"%") //数据库字段为jsonb，要实现模糊查询请修改查询条件
	}
	if data.DepartmentId != 0 {
		db = db.Where("department_id = ?", data.DepartmentId)
	}
	// if data.Active != false {
	// 	db = db.Where("active = ?", data.Active)
	// }
	// db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	// info, _ := global.Db.DB()
	// stats := info.Stats()
	// fmt.Printf("OpenConnections: %d, InUse: %d, Idle: %d",
	// 	stats.OpenConnections, stats.InUse, stats.Idle) //vault: OpenConnections: 0, InUse: 0, Idle: 0  初始：OpenConnections: 2, InUse: 1, Idle: 1
	return &list, total, err
}

func (m *hrJobModelImpl) FindList(data entity.HrJob) (*[]entity.HrJob, error) {
	list := make([]entity.HrJob, 0)
	db := global.HrDb.Table(m.table)
	// 此处填写 where参数判断
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	if data.Name != "" {
		db = db.Where("name = ?", data.Name)
	}
	if data.DepartmentId != 0 {
		db = db.Where("department_id = ?", data.DepartmentId)
	}
	// if data.Active != "" {
	// 	db = db.Where("active = ?", data.Active)
	// }
	err := db.Find(&list).Error
	return &list, err
}

func (m *hrJobModelImpl) Update(data entity.HrJob) odoorpc.ResultData {
	// return global.Db.Table(m.table).Updates(&data).Error
	// job, e := m.FindOne(data.Id)
	// if e != nil {
	// 	return odoorpc.ResultData{Status: "failed", Message: "公司中没有该部门"}
	// }
	// 判断部门是否为该公司所属
	if data.DepartmentId != 0 {
		department := &hrDepartmentModelImpl{}
		com, _ := department.FindOne(data.DepartmentId)
		list := *com
		fmt.Printf("当前部门信息：%+v\n", list)
		// if err != nil || list ==  {
		// 	return odoorpc.ResultData{Status: "failed", Message: "公司中没有该部门"}
		// }
		// 判断该部门树是否属于岗位中原定义的公司下,只修改了部门，对应也要修改公司信息
		if data.CompanyId == 0 {
			data.CompanyId = list.ID
		} else {
			if list.ID != data.CompanyId {
				return odoorpc.ResultData{Status: "failed", Message: "该部门不属于该公司，请检查"}
			}
		}
	}
	fmt.Printf("当前修改信息：%+v\n", data)
	dd, _ := json.Marshal(data)
	var writeuData map[string]interface{}
	_ = json.Unmarshal(dd, &writeuData)
	writeuResult, err := odoorpc.Write(writeuData, "hr.job", data.Id) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	if err != nil {
		return odoorpc.ResultData{Status: "fail", Code: 500, Message: "修改失败！"}
	}
	return writeuResult
}

func (m *hrJobModelImpl) Publish(id int64) odoorpc.ResultData {
	data := entity.HrJob{
		Id:          id,
		IsPublished: true,
	}
	dd, _ := json.Marshal(data)
	var writeuData map[string]interface{}
	_ = json.Unmarshal(dd, &writeuData)
	writeuResult, err := odoorpc.Write(writeuData, "hr.job", id) //调用内部 res.users create方法时报错partner_id约束错误，但是web页面新建时根本不需要
	if err != nil {
		return odoorpc.ResultData{Status: "fail", Code: 500, Message: "修改失败！"}
	}
	return writeuResult
}

// 部门归档 部门中原有员工需要处理
func (m *hrJobModelImpl) Delete(Ids []int64) odoorpc.ResultData {
	// return global.Db.Table(m.table).Delete(&entity.HrJob{}, "post_id in (?)", postIds).Error
	for i := 0; i < len(Ids); i++ {
		result := global.HrDb.Model(&entity.HrJob{}).
			Where("id = ?", Ids[i]).
			Update("active", false)
		// Updates(entity.ResUsersB{
		// 	WriteUid:     user.ID,
		// 	WriteDate:    time.Now().Format(time.RFC3339),
		// 	PandaxSecret: wizard.Secret,
		// })
		if result.Error != nil {
			return odoorpc.ResultData{Status: "fail", Code: 500, Message: "归档失败！"}
		}

	}
	return odoorpc.ResultData{Status: "success", Code: 200, Message: "成功归档！"}
}
