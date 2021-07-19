package middleware

import (
	"bytes"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"gozero-vue-admin/common/global"
	"gozero-vue-admin/service/user/model"
)

type OperationRecordMiddleware struct {
	jwt *JwtMiddleware
}

func NewOperationRecordMiddleware(jwtObject *JwtMiddleware) *OperationRecordMiddleware {
	return &OperationRecordMiddleware{
		jwt: jwtObject,
	}
}

func (m *OperationRecordMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body []byte
		var userId int
		if r.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(r.Body)
			if err != nil {
				global.ZapLog.Error("read body from request error:", zap.Any("err", err))
				return
			} else {
				r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}

		token := r.Header.Get("x-token")
		claims, err := m.jwt.ParseToken(token)
		if err != nil {
			return
		} else {
			if claims.ID != 0 {
				userId = int(claims.ID)
			} else {
				id, err := strconv.Atoi(r.Header.Get("x-user-id"))
				if err != nil {
					userId = 0
				}
				userId = id
			}
		}

		record := model.SysOperationRecord{
			Ip:     r.Host,
			Method: r.Method,
			Path:   r.URL.Path,
			Agent:  r.UserAgent(),
			Body:   string(body),
			UserID: userId,
		}

		now := time.Now()

		// Passthrough to next handler if need
		next(w, r)

		record.Latency = time.Now().Sub(now)
		//record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		//record.Status = c.Writer.Status()
		//record.Resp = writer.body.String()

		if err := model.CreateSysOperationRecord(record); err != nil {
			global.ZapLog.Error("create operation record error:", zap.Any("err", err))
		}
	}
}

type responseBodyWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
