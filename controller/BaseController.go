// 控制器层，通过各Action调用业务逻辑层各模块，生成Json、流文件等格式数据完成请求
package controller

import (
	"encoding/json"
	"net/http"
	"log"
	"thinkgo/response"
	"thinkgo/utils"
)

type BaseController struct {

}

func (this *BaseController) ReturnJsonObj(resp *response.Response, w http.ResponseWriter, r *http.Request) {
	reply, err := json.Marshal(resp)
	if err != nil {
		ctl, action := utils.CtlAction(r.URL.Path)
		log.Fatalf(ctl + "/" + action + " reply := json.Marshal() error => %v\n", err)
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(reply)
}