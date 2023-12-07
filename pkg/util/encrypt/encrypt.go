package encrypt

import (
	"github.com/Campus-Hub/server/pkg/consts"
	"golang.org/x/crypto/bcrypt"
)

// SetPassword 设置密码，实现 staff 接口？
func SetPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+consts.Salt), consts.PassWordCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword 校验密码，实现 staff 接口？
// 有两个输入，string password 为接受验证的传入密码，
// string hash 为从数据库中读取到的密码哈希值
// 返回两个哈希值比较的结果
func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+consts.Salt))
	return err == nil
}
