package tool

import (
	"fmt"
	"pandax/apps/shared/entity"
	"pandax/pkg/global"
	"testing"
)

func TestMfa(t *testing.T) {
	mfa := &TotpMfa{}
	// 1、测试 生成密钥及二维码
	r, _ := mfa.Initiate("create", 2)
	fmt.Printf("%+v\n", r)

	// 2、测试 激活 cg
	user := new(entity.ResUsers)
	e := global.HrDb.Model(&entity.ResUsers{}).Where("id = ?", 2).Find(&user)
	_ = mfa.Enable(user, r, "admin")
	fmt.Println(e)

	// 3、测试 验证 cg
	p, _ := mfa.Verify("449098", user)
	fmt.Printf("%+v\n", p)

	// 4、测试 关闭 cg
	d := mfa.Disable(32, 2)
	fmt.Printf("%+v\n", d)
}
