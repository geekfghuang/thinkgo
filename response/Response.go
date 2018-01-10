package response

const (
	SHOULDNOTEMPTY	=	"不能为空"
	WRONG			=	"错误"
	FAILED			=	"失败"
	SUCCEED			=	"成功"
	OK				=	"OK"

	ALREADYLOGGEDIN	=	"你已经登录，请先注销"
	NOTYETLOGGEDIN	=	"你还未登录，请先登录"
)

type Response struct {
	Code 	int
	Msg 	string
	Data 	interface{}
}