package tool

import (
	"testing"
)

func TestMfa(t *testing.T) {
	// 查找最新totp记录
	// TOTP := new(entity.AuthTotpWizard)
	// err := global.HrDb.Model(&entity.AuthTotpWizard{}).Where("user_id = ?", 2).Where("secret != ?", "false").Order("create_date desc").Find(&TOTP).Error
	// if err != nil {
	// 	fmt.Printf("错误：%+v\n", err)
	// }
	// fmt.Printf("记录：%+v\n", TOTP)

	// mfa := &TotpMfa{}
	// // 1、测试 生成密钥及二维码
	// r, _ := mfa.Initiate("create", 2)
	// fmt.Printf("%+v\n", r)

	// // 2、测试 激活 cg
	// totp, err := mfa.Enable(entity.LoginO{Login: "admin", Password: "admin"})
	// fmt.Printf("totp初始化=%+v\n发生错误：%+v\n", totp, err)

	// 3、测试 验证 cg
	// p, _ := mfa.Verify("743824", TOTP)
	// fmt.Printf("Result= %+v\n", p)

	// 4、测试 关闭 cg
	// d := mfa.Disable(32, 2)
	// fmt.Printf("%+v\n", d)
}
