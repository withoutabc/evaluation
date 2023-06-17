package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"rpc/app/common/config/mycasbin"
	"strconv"
)

type AuthMIDDLEWAREMiddleware struct {
}

func NewAuthMIDDLEWAREMiddleware() *AuthMIDDLEWAREMiddleware {
	return &AuthMIDDLEWAREMiddleware{}
}

func (m *AuthMIDDLEWAREMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value("user_id").(int64)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": 2007,
				"msg":    "没有获取到用户id",
			})
			return
		}
		log.Println(userID)
		e := mycasbin.InitEnforcer()
		b, err := e.Enforce(strconv.FormatInt(userID, 10), "admin_data", "write")
		if err != nil || b != true {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": 2008,
				"msg":    "没有对应权限",
			})
			return
		}
		/*
			b, err := e.AddGroupingPolicy(userId, "admin_group")
		*/
		next(w, r)
	}
}
