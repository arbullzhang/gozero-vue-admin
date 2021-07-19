package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/common/utils"
	"gozero-vue-admin/service/user/api/internal/config"
	"gozero-vue-admin/service/user/model"
)

type CustomClaims struct {
	UUID 		uuid.UUID
	ID 			uint
	Username    string
	NickName    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}

func NewJwtMiddleware(c config.Config) *JwtMiddleware {
	return &JwtMiddleware{
		Config: c,
		SigningKey:[]byte(c.Auth.AccessSecret),
	}
}

type JwtMiddleware struct {
	Config config.Config
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从header头中取出x-token。go-zero中的jwt组件，默认是从Authorization中取token的
		token := r.Header.Get("x-token")
		if token == "" {
			utils.FailWithDetailed(utils.H{"reload": true}, "未登录或非法访问", w)
			return
		}

		if model.IsBlacklist(token) {
			utils.FailWithDetailed(utils.H{"reload": true}, "您的帐户异地登陆或令牌失效", w)
			return
		}

		claims, err := m.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				utils.FailWithDetailed(utils.H{"reload": true}, "授权已过期", w)
				return
			}

			utils.FailWithDetailed(utils.H{"reload": true}, err.Error(), w)
			return
		}

		if err, _= model.FindUserByUuid(claims.UUID.String()); err != nil {
			model.JsonInBlacklist(model.JwtBlacklist{Jwt: token})
			utils.FailWithDetailed(utils.H{"reload": true}, err.Error(), w)
			return
		}

		if claims.ExpiresAt - time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + m.Config.Auth.AccessExpire
			newToken, _ := m.CreateToken(*claims)
			newClaims, _ := m.ParseToken(newToken)
			w.Header().Set("new-token", newToken)
			w.Header().Set("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if m.Config.System.UseMultipoint {
				err, RedisJwtToken := model.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.ZapLog.Error("get redis jwt failed", zap.Any("err", err))
				} else { // 当之前的取成功时才进行拉黑操作
					model.JsonInBlacklist(model.JwtBlacklist{Jwt: RedisJwtToken})
				}

				model.SetRedisJWT(newToken, newClaims.Username, m.Config.Auth.AccessExpire)
			}
		}
		next(w, r)
	}
}

// 创建一个token
func (m *JwtMiddleware) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.SigningKey)
}

// 解析token
func (m *JwtMiddleware) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return m.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}


