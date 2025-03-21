package tool

import (
	"gopkg.in/hlandau/passlib.v1/abstract"
	"gopkg.in/hlandau/passlib.v1/hash/pbkdf2"
)

type Context struct {
	Schemes []abstract.Scheme
}

var defualtSchemes = []abstract.Scheme{
	pbkdf2.SHA512Crypter,
}

// 加密算法
func HashPwd(password string) string {
	var hash string
	for _, scheme := range defualtSchemes {
		ctx := Context{Schemes: []abstract.Scheme{scheme}}
		hash, _ = ctx.Schemes[0].Hash(password) //casdoor加密后hashf("hash=%v\n", hash)
	}
	return hash
}

// 验证算法
func VerifyPwd(password string, hashedPwd string) bool {
	flag := true
	for _, scheme := range defualtSchemes {
		ctx := Context{Schemes: []abstract.Scheme{scheme}}
		st := ctx.Schemes[0].Verify(password, hashedPwd)
		if st != nil {
			flag = false
		}
	}
	return flag
}
