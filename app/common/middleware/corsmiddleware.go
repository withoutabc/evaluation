package middleware

import "net/http"

type CORSMIDDLEWAREMiddleware struct {
}

func NewCORSMIDDLEWAREMiddleware() *CORSMIDDLEWAREMiddleware {
	return &CORSMIDDLEWAREMiddleware{}
}

func (m *CORSMIDDLEWAREMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 设置跨域请求头
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token, x-access-token")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}
