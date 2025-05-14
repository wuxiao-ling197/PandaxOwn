package services

import (
	"fmt"
	"pandax/apps/dicm/entity"
	"testing"
)

func TestService(t *testing.T) {
	rackModel := &dicmRackImpl{}
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

	result := rackModel.Update(entity.DcimRack{Id: 13, Status: "新建，即将投入使用ing..."})
	fmt.Println(result)

	// siteModel := &dicmSiteImpl{}
	// list, err := siteModel.FindList(entity.DcimSite{})
	// fmt.Println(err)
	// fmt.Println(list)

}
