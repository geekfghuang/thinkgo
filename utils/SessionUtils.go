package utils

import (
	"github.com/astaxie/beego/session"
	"net/http"
	"log"
)

var sessionManager *session.Manager

func init() {
	config := &session.ManagerConfig{CookieName: "thinkgosessionid", EnableSetCookie: true, Maxlifetime: 60 * 2, CookieLifeTime: 60 * 2, Gclifetime: 60 * 2}
	sessionManager, _ = session.NewManager("memory", config)
	go sessionManager.GC()
}

func GetSession(w http.ResponseWriter, r *http.Request) session.Store{
	sess, err := sessionManager.SessionStart(w, r)
	if err != nil {
		log.Fatalf("utils GetSession() sess, err := SessionManager.SessionStart() error => %v\n", err)
	}
	return sess
}

func DestorySession(w http.ResponseWriter, r *http.Request){
	sessionManager.SessionDestroy(w, r)
}
