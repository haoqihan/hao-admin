package app

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"hao-admin/global"
	"time"
)

type Claims struct {
	UUID        uuid.UUID
	ID          uint32
	Username    string
	NickName    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}

// GetJWTSecret 获取该项目的 secret 秘钥信息
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken 生成jwt token信息
func GenerateToken(UUID uuid.UUID, ID uint32, NickName, Username, AuthorityId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		UUID:       UUID,
		ID:          ID,
		NickName:    NickName,
		Username:    Username,
		AuthorityId: AuthorityId,
		BufferTime:  60 * 60 * 24, // 缓冲时间1天，缓冲时间内会获得新的token令牌，此时用户会有两个有效令牌
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),        // 过期时间
			Issuer:    global.JWTSetting.Issuer, // 签名发行者
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
