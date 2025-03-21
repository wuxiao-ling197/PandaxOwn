package services

import (
	"pandax/apps/device/entity"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
)

type (
	DeviceAlarmModel interface {
		Insert(data entity.DeviceAlarm) error
		FindOne(id string) (*entity.DeviceAlarm, error)
		FindOneByType(deviceId, ty, state string) (*entity.DeviceAlarm, error)
		FindListPage(page, pageSize int, data entity.DeviceAlarmForm) (*[]entity.DeviceAlarm, int64, error)
		Update(data entity.DeviceAlarm) error
		Delete(ids []string) error
		FindAlarmCount() (entity.DeviceCount, error)
		FindAlarmPanel() ([]entity.DeviceAlarmCount, error)
	}

	alarmModelImpl struct {
		table string
	}
)

var DeviceAlarmModelDao DeviceAlarmModel = &alarmModelImpl{
	table: `device_alarms`,
}

func (m *alarmModelImpl) Insert(data entity.DeviceAlarm) error {
	err := global.Db.Table(m.table).Create(&data).Error
	return err
}

func (m *alarmModelImpl) FindOne(id string) (*entity.DeviceAlarm, error) {
	resData := new(entity.DeviceAlarm)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	return resData, err
}

func (m *alarmModelImpl) FindOneByType(deviceId, ty, state string) (*entity.DeviceAlarm, error) {
	resData := new(entity.DeviceAlarm)
	db := global.Db.Table(m.table).Where("device_id = ?", deviceId).Where("type = ? ", ty).Where("state = ? ", state)
	err := db.First(resData).Error
	if err != nil {
		return nil, err
	}
	return resData, nil
}

func (m *alarmModelImpl) FindListPage(page, pageSize int, data entity.DeviceAlarmForm) (*[]entity.DeviceAlarm, int64, error) {
	list := make([]entity.DeviceAlarm, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	if data.DeviceId != "" {
		db = db.Where("device_id = ?", data.DeviceId)
	}
	if data.ProductId != "" {
		db = db.Where("product_id = ?", data.ProductId)
	}
	if data.Level != "" {
		db = db.Where("level = ?", data.Level)
	}
	if data.Type != "" {
		db = db.Where("type like ?", "%"+data.Type+"%")
	}
	if data.State != "" {
		db = db.Where("state = ?", data.State)
	}
	if data.StartTime != "" {
		db = db.Where("time > ?", data.StartTime)
	}
	if data.EndTime != "" {
		db = db.Where("time < ?", data.EndTime)
	}
	// 组织数据访问权限
	if err := model.OrgAuthSet(db, data.RoleId, data.Owner); err != nil {
		return nil, 0, err
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := db.Order("time").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *alarmModelImpl) Update(data entity.DeviceAlarm) error {
	err := global.Db.Table(m.table).
		Where("type = ?", data.Type).
		Where("device_id = ?", data.DeviceId).
		Updates(&data).Error
	return err
}

func (m *alarmModelImpl) Delete(id []string) error {
	return global.Db.Table(m.table).Delete(&entity.DeviceAlarm{}, "id in (?)", id).Error
}

// 获取告警数量统计
func (m *alarmModelImpl) FindAlarmCount() (count entity.DeviceCount, err error) {
	sql := `SELECT COUNT(*) AS total, (SELECT COUNT(*) FROM device_alarms WHERE DATE(time) = CURDATE()) AS today  FROM device_alarms`
	if global.Conf.Server.DbType == "postgresql" {
		sql = `SELECT COUNT(*) AS total, (SELECT COUNT(*) FROM device_alarms WHERE DATE(time) = current_date) AS today  FROM device_alarms`
	}
	err = global.Db.Raw(sql).Scan(&count).Error
	return
}
func (m *alarmModelImpl) FindAlarmPanel() (count []entity.DeviceAlarmCount, err error) {
	sql := `SELECT CAST(time AS DATE) AS date, COUNT(*) AS count FROM device_alarms WHERE time >= DATE_SUB(CURDATE(), INTERVAL 30 DAY) GROUP BY CAST(time AS DATE)`
	if global.Conf.Server.DbType == "postgresql" {
		sql = `SELECT CAST(time AS DATE) AS date, COUNT(*) AS count FROM device_alarms WHERE time >= CURRENT_DATE - INTERVAL '30 day' GROUP BY CAST(time AS DATE);`
	}
	err = global.Db.Raw(sql).Scan(&count).Error
	return
}
