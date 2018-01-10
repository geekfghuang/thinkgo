//公共模块，一些公用函数
package utils

import (
	"strings"
)

//根据URL返回controller、action的全名
func CtlAction(url string) (ctl, action string) {//形如/test/hello
	path := strings.ToLower(url)
	ctlaction := strings.Split(path, "/")
	ctl = string(ctlaction[1][0] - ('a' - 'A')) + ctlaction[1][1:] + "Controller"
	action = string(ctlaction[2][0] - ('a' - 'A')) + ctlaction[2][1:] + "Action"
	return
}

