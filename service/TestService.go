package service

import (
	"net/http"
	"thinkgo/response"
	"thinkgo/utils"
	"time"
)

type TestService struct {
	BaseService
}

func (this *TestService) HelloService(resp *response.Response , w http.ResponseWriter, r *http.Request){
	resp.Code = 0
	resp.Msg = response.OK
	resp.Data = "hello thinkgo~"
}

func (this *TestService) RegisterService(resp *response.Response , w http.ResponseWriter, r *http.Request){
	username, password := r.FormValue("username"), r.FormValue("password")
	if len(username) == 0 {
		resp.Code = -1
		resp.Msg = "用户名" + response.SHOULDNOTEMPTY
		return
	}
	if len(password) == 0 {
		resp.Code = -1
		resp.Msg = "密码" + response.SHOULDNOTEMPTY
		return
	}

	//裸写sql操作，不走orm
	db_sql := "insert into t_user(username, password, join_time) values(?, ?, ?)"
	_, err := utils.Insert("root", "pipigui", "db_thinkgo", db_sql, username, password, time.Now())
	if err != nil {
		resp.Code = -1
		resp.Msg = "注册" + response.FAILED
	}else{
		resp.Code = 0
		resp.Msg = response.OK
		resp.Data = "注册" + response.SUCCEED
	}
}

func (this *TestService) LoginService(resp *response.Response , w http.ResponseWriter, r *http.Request){
	//先检查是否已存在session
	sess := utils.GetSession(w, r)
	idinterface := sess.Get("id")
	if idinterface != nil {
		resp.Code = -1
		resp.Msg = response.ALREADYLOGGEDIN
		return
	}

	username, password := r.FormValue("username"), r.FormValue("password")
	if len(username) == 0 {
		resp.Code = -1
		resp.Msg = "用户名" + response.SHOULDNOTEMPTY
		return
	}
	if len(password) == 0 {
		resp.Code = -1
		resp.Msg = "密码" + response.SHOULDNOTEMPTY
		return
	}

	//裸写sql操作，不走orm
	db_sql := "select id, join_time from t_user where username = ? and password = ?"
	rows, err := utils.Select("root", "pipigui", "db_thinkgo", db_sql, username, password)
	if err != nil {
		resp.Code = -1
		resp.Msg = "登录" + response.FAILED
		return
	}
	var rowscount int
	var id int
	var joinTime string//不能是time.Time
	for rows.Next() {
		rows.Scan(&id ,&joinTime)
		rowscount++
	}
	if rowscount == 0 {
		resp.Code = -1
		resp.Msg = "用户名或密码" + response.WRONG
	}else {
		resp.Code = 0
		resp.Msg = response.OK
		resp.Data = "登录" + response.SUCCEED

		//session操作
		sess.Set("id", id)
		sess.Set("username", username)
	}
}

func (this *TestService) LogoutService(resp *response.Response , w http.ResponseWriter, r *http.Request){
	sess := utils.GetSession(w, r)
	idinterface := sess.Get("id")
	if idinterface == nil {
		resp.Code = -1
		resp.Msg = response.NOTYETLOGGEDIN
		//utils.GetSession就会在服务端生成或读取已有session（但里面没键值），在客户端生成sessionid，所以要destory，客户端cookie也会被清除
		utils.DestorySession(w, r)
		return
	}
	utils.DestorySession(w, r)
	resp.Code = 0
	resp.Msg = response.OK
	resp.Data = "注销" + response.SUCCEED
}