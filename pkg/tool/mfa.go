// Copyright 2023 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tool

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image/png"
	"time"

	"pandax/apps/shared/entity"
	"pandax/pkg/global"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	qrcode "github.com/skip2/go-qrcode"
)

type MfaProps struct {
	Enabled bool   `json:"enabled"`
	Secret  string `json:"secret,omitempty"`
	URL     string `json:"url,omitempty"`
}

const (
	MfaTotpPeriodInSeconds = 30
)

type TotpMfa struct {
	*entity.AuthTotpWizard
	period     uint
	secretSize uint
	digits     otp.Digits //生成的验证码位数
}

// auth_totp_wizard 生成 TOTP记录
func (mfa *TotpMfa) Initiate(login string, operator int64) (*entity.AuthTotpWizard, error) {
	issuer := "Pandax"
	// 生成totp密钥以及二维码扫描图像URL
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: login,
		Period:      mfa.period,
		SecretSize:  mfa.secretSize, //默认为32
		Digits:      mfa.digits,     //默认为6位验证码
	})
	if err != nil {
		return nil, err
	}
	// 生成 QR 码图片并写入响应
	qrCode, err := qrcode.New(key.URL(), qrcode.Low)
	if err != nil {
		return &entity.AuthTotpWizard{}, err
	}
	// 将二维码编码为PNG格式
	var buf bytes.Buffer
	err = png.Encode(&buf, qrCode.Image(256))
	if err != nil {
		return nil, err
	}
	// 将PNG格式的二维码转换为Base64编码
	qrCodeBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	// fmt.Printf("转换成功: %v\n ", qrCodeBase64)
	// fmt.Print("渲染二维码\n")
	// fmt.Println(qrCode.WriteFile(256, "qrcode.jpg"))
	// qrCode.WriteFile(256, key.URL())
	// mfaProps := MfaProps{
	// 	Secret: key.Secret(),
	// 	URL:    key.URL(),
	// }
	user := new(entity.ResUsersId)
	er := global.HrDb.Table("res_users").Where("login = ?", login).First(&user).Error
	if er != nil {
		return nil, errors.New("不存在该用户，请检查登录账号是否准确！")
	}
	// 保存到数据库
	data := entity.AuthTotpWizard{
		UserId:     user.ID,
		Secret:     key.Secret(),
		Url:        key.URL(),
		Qrcode:     "data:image/png;base64," + qrCodeBase64, //否则将不被识别为base64类型字段
		CreateUid:  operator,
		WriteUid:   operator,
		CreateDate: time.Now().Format(time.RFC3339),
		WriteDate:  time.Now().Format(time.RFC3339),
	}
	// 该odoo对应模型中不存在create方法
	// var create map[string]interface{}
	// createData, _ := json.Marshal(data)
	// _ = json.Unmarshal(createData, &create)
	// r, _ := odoorpc.Create(create, "auth_totp.wizard")
	// 通过grom保存
	r := global.HrDb.Model(&entity.AuthTotpWizard{}).Create(&data).Error
	return &data, r
}

// 激活 TOTP，最终版 只返回totp初始化记录 数据更改在verify中实现
// 用户登录时调用，具体实现就是将secret数据保存到用户表 user应该是当前操作的用户【即被启用totp的用户，在初始登录】 wizard是totp初始化记录
// 默认调用前并未激活，调用前应先判断用户pandax_secret值是否存在，作为首次验证的判断条件，之后调用verify进行验证
func (mfa *TotpMfa) Enable(login entity.LoginO) (*entity.AuthTotpWizard, error) {
	// 验证当前操作用户密码
	user := entity.ResUsers{}
	TOTP := entity.AuthTotpWizard{}
	totpDB := global.HrDb.Model(&entity.AuthTotpWizard{})
	err := global.HrDb.Table("res_users").Where("login = ?", login.Login).Find(&user).Error
	// fmt.Printf("用户信息：%+v\n", user)
	if err != nil {
		return nil, errors.New("用户不存在，请重新输入")
	}
	b := VerifyPwd(login.Password, user.Password)
	if !b {
		return nil, errors.New("密码错误，请重新输入")
	}
	// fmt.Printf("验证密码：%v\n", b)
	// 更改数据库数据
	// result := global.HrDb.Model(&entity.AuthTotpWizard{}).Where("id = ?", user.ID).Update("secret", "")
	// if result.Error != nil {
	// 	return result.Error
	// 	// return odoorpc.ResultData{Status: "fail", Code: 500, Message: "归档失败！"}
	// }

	if user.PandaxSecret == "" || user.PandaxSecret == "false" {
		// 如果没有secret值,查询最新的totp记录
		err = totpDB.Where("user_id = ?", user.ID).Where("secret != ?", "false").Order("create_date desc").Find(&TOTP).Error
		if err != nil {
			return nil, err
		}
		// 保存ResUsers表中secret
		// r := global.HrDb.Model(&entity.ResUsers{}).Where("id = ?", user.ID).Updates(entity.ResUsersB{
		// 	WriteUid:     user.ID,
		// 	WriteDate:    time.Now().Format(time.RFC3339),
		// 	PandaxSecret: TOTP.Secret,
		// })
		// if r.Error != nil {
		// 	return nil, r.Error
		// }
		// // 清空AuthTotpWizard表中secret
		// totpDB := global.HrDb.Model(&entity.AuthTotpWizard{})
		// d := totpDB.Where("id = ?", TOTP.Id).Updates(entity.AuthTotpWizard{
		// 	WriteUid:  TOTP.UserId,
		// 	WriteDate: time.Now().Format(time.RFC3339),
		// 	Secret:    "false",
		// })
		// if d.Error != nil {
		// 	return nil, d.Error
		// }
		return &TOTP, nil
	}

	newtotp := new(entity.AuthTotpWizard)
	totpDB.Where("user_id = ?", user.ID).Where("url like ?", "%"+user.PandaxSecret+"%").Find(&newtotp)
	return newtotp, nil
}

// 验证 TOTP, 登录时调用
func (mfa *TotpMfa) Verify(passcode string, wizard *entity.AuthTotpWizard) (bool, error) {
	// fmt.Printf("mfa方法调用，%+v\n", wizard)
	user := entity.ResUsers{}
	// 认证
	if wizard.Secret == "false" {
		global.HrDb.Table("res_users").Where("id = ?", wizard.UserId).Find(&user)
		result, err := totp.ValidateCustom(passcode, user.PandaxSecret, time.Now().UTC(), totp.ValidateOpts{
			Period:    MfaTotpPeriodInSeconds,
			Skew:      1,
			Digits:    otp.DigitsSix,
			Algorithm: otp.AlgorithmSHA1,
		})
		if err != nil || !result {
			return result, errors.New("mfa认证码错误！")
		}
		return result, nil
	} else {
		// 第一次认证
		result, err := totp.ValidateCustom(passcode, wizard.Secret, time.Now().UTC(), totp.ValidateOpts{
			Period:    MfaTotpPeriodInSeconds,
			Skew:      1,
			Digits:    otp.DigitsSix,
			Algorithm: otp.AlgorithmSHA1,
		})
		if err != nil || !result {
			return result, errors.New("mfa认证码错误！")
		}
		// 保存ResUsers表中secret
		r := global.HrDb.Model(&entity.ResUsers{}).Where("id = ?", wizard.UserId).Updates(entity.ResUsersB{
			WriteUid:     wizard.UserId,
			WriteDate:    time.Now().Format(time.RFC3339),
			PandaxSecret: wizard.Secret,
		})
		if r.Error != nil {
			return false, r.Error
		}
		// 清空AuthTotpWizard表中secret
		totpDB := global.HrDb.Model(&entity.AuthTotpWizard{})
		d := totpDB.Where("id = ?", wizard.Id).Updates(entity.AuthTotpWizard{
			WriteUid:  wizard.UserId,
			WriteDate: time.Now().Format(time.RFC3339),
			Secret:    "false",
		})
		if d.Error != nil {
			return false, d.Error
		}

		return result, nil
	}
}

// 关闭 TOTP
func (mfa *TotpMfa) Disable(userId int64, operator int64) error {
	// fmt.Println("关闭TOTP")
	// result := global.HrDb.Model(&entity.ResUsers{}).Where("id = ?", userId).Update("pandax_secret", "false")
	result := global.HrDb.Model(&entity.ResUsers{}).Where("id = ?", userId).Updates(entity.ResUsersB{
		WriteUid:     operator,
		WriteDate:    time.Now().Format(time.RFC3339),
		PandaxSecret: "false",
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
