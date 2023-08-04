package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func logMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"Method": r.Method,
			"Time":   time.Now().Format(time.ANSIC),
		}).Info("MiddleWareOutMethod")
		next.ServeHTTP(w, r)
	})
}
