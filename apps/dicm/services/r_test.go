package services

import (
	"log"
	"pandax/apps/dicm/entity"
	"testing"
)

func TestService(t *testing.T) {
	// rackModel := &dicmRackImpl{}
	// list, err := rackModel.FindList(entity.DcimRack{})
	// fmt.Println(err)
	// fmt.Println(list)

	// find, err := rackModel.FindOne(entity.DcimRack{Name: "机柜1"})
	// fmt.Println(err)
	// fmt.Println(find)

	// siteId := uint(1)
	// data := entity.DcimRack{
	// 	CustomFieldData: "{}",
	// 	Name:            "post测试",
	// 	NickName:        "post",
	// 	Status:          "xinjian",
	// 	Serial:          1,
	// 	Type:            "test",
	// 	Width:           180,
	// 	UHeight:         42,
	// 	DescUnits:       false,
	// 	OuterUnit:       "英寸",
	// 	Comments:        "评价",
	// 	SiteId:          &siteId,
	// 	WeightUnit:      "kg",
	// 	Description:     "描述",
	// 	StartingUnit:    1,
	// 	Created:         time.Now(),
	// }
	// result, err := rackModel.Insert(data)
	// fmt.Println(err)
	// fmt.Println(result)

	// result := rackModel.Update(entity.DcimRack{Id: 13, Status: "新建，即将投入使用ing..."})
	// fmt.Println(result)

	// siteModel := &dicmSiteImpl{}
	// list, err := siteModel.FindList(entity.DcimSite{})
	// fmt.Println(err)
	// fmt.Println(list)

	tenant := &dicmSiteImpl{}
	// data := int64(2)
	l, tt, e := tenant.FindListPage(1, 10, entity.DcimSite{})
	log.Println(l)
	log.Println(tt)
	log.Println(e)

	// r, ee := tenant.FindOne(entity.DcimSite{Id: 5})
	// log.Println(r)
	// log.Println(ee)

	// r, e := tenant.Insert(entity.TenancyTenant{
	// 	Name:            "1",
	// 	Slug:            "1",
	// 	Description:     "1",
	// 	Comments:        "1",
	// 	CustomFieldData: "{}",
	// })
	// log.Println(r)
	// log.Println(e)

	// 加入租户组
	// tenantIds := []int64{5, 6}
	// err := tenant.JoinGroup(tenantIds, 3)
	// log.Println(err)

	// r, e := tenant.InsertGroup(entity.TenancyTenantgroup{
	// 	Name:            "3",
	// 	Slug:            "3",
	// 	Description:     "1",
	// 	TreeId:          1,
	// 	Lft:             1,
	// 	Level:           1,
	// 	Rght:            1,
	// 	ParentId:        sql.NullInt64{Int64: 2, Valid: true},
	// 	CustomFieldData: "{}",
	// })
	// log.Println(r)
	// log.Println(e)

	// e := tenant.UpdateGroup(entity.TenancyTenantgroup{
	// 	Id: 5,
	// 	// Name:        "2",
	// 	// Slug:        "2",
	// 	// Description: "1",
	// 	// TreeId:      1,
	// 	// Lft:         1,
	// 	// Level:       1,
	// 	// Rght:        1,
	// 	ParentId: 0, //sql.NullInt64{Int64: 2, Valid: true},
	// 	// CustomFieldData: "{}",
	// })
	// log.Println(e)

	// 层级结构
	// r, e := tenant.GetGroupStructrue()
	// log.Printf("%+v\n%+v\n", r[0], r[0].Children[0])
	// log.Println(e)

}
