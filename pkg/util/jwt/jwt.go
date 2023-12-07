package jwt

import (
	"fmt"
	"time"

	"github.com/Campus-Hub/server/pkg/consts"
	"github.com/golang-jwt/jwt/v5"
)

type HubClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateToken(userId int64) (string, error) {

	var (
		expireTime = time.Now().Add(24 * time.Hour)
		nowTime    = time.Now()
	)

	claims := HubClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			Issuer:    "campus-hub",
			// Subject 写入职位 or 部门
			// Subject:   "somebody",
			// Audience 可选字段，用于鉴权和范围限定
			// Audience:  []string{"/protect/api/v1/:company/:dept/:uid"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(consts.JwtSecret)

	return tokenStr, err
}

func ParseToken(tokenStr string) (bool, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return consts.JwtSecret, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println(claims["user_id"], claims["exp"])
	} else {
		fmt.Println(err)
	}

	return ok, err
}

// tokenProcess test function
func tokenProcess(userId int64) bool {
	tokenStr, _ := CreateToken(10)
	ok, _ := ParseToken(tokenStr)
	return ok
}
