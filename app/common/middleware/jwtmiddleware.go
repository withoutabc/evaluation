package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"rpc/utils/secret"
	"strings"
)

type JWTMIDDLEWAREMiddleware struct {
}

func NewJWTMIDDLEWAREMiddleware() *JWTMIDDLEWAREMiddleware {
	return &JWTMIDDLEWAREMiddleware{}
}

func (m *JWTMIDDLEWAREMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": 2003,
				"msg":    "请求头中auth为空",
			})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": 2004,
				"msg":    "请求头中auth格式有误",
			})
			return
		}
		mc, err := secret.ParseToken(parts[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": 2005,
				"msg":    "无效的Token",
			})
			return
		}
		if mc.Type != "access" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": 2006,
				"msg":    "错误的Token类型",
			})
			return
		}
		// 将解析出的用户ID和角色信息设置到请求上下文中
		ctx := context.WithValue(r.Context(), "user_id", mc.ID)
		ctx = context.WithValue(ctx, "role", mc.Role)
		// 调用下一个处理器
		next(w, r.WithContext(ctx))
		// 下一个处理器获取方式
		// userID, ok := r.Context().Value("user_id").(string)
	}
}
