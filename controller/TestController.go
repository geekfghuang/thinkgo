package controller

import (
	"net/http"
	"thinkgo/access"
	"thinkgo/service"
	"thinkgo/response"
)

type TestController struct{
	BaseController
}

func init(){
	access.RegisterController(&TestController{})
}

func (this *TestController) HelloAction(w http.ResponseWriter, r *http.Request){
	svc, resp := new(service.TestService), new(response.Response)
	svc.HelloService(resp, w, r)
	this.ReturnJsonObj(resp, w, r)
}

func (this *TestController) RegisterAction(w http.ResponseWriter, r *http.Request){
	svc, resp := new(service.TestService), new(response.Response)
	svc.RegisterService(resp, w, r)
	this.ReturnJsonObj(resp, w, r)
}


func (this *TestController) LoginAction(w http.ResponseWriter, r *http.Request){
	svc, resp := new(service.TestService), new(response.Response)
	svc.LoginService(resp, w, r)
	this.ReturnJsonObj(resp, w, r)
}

func (this *TestController) LogoutAction(w http.ResponseWriter, r *http.Request){
	svc, resp := new(service.TestService), new(response.Response)
	svc.LogoutService(resp, w, r)
	this.ReturnJsonObj(resp, w, r)
}