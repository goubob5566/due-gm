package code

import "github.com/dobyte/due/v2/codes"

var (
	OK               = codes.OK
	Canceled         = codes.Canceled
	Unknown          = codes.Unknown
	InvalidArgument  = codes.InvalidArgument
	DeadlineExceeded = codes.DeadlineExceeded
	NotFound         = codes.NotFound
	InternalError    = codes.InternalError
	Unauthorized     = codes.Unauthorized
	IllegalInvoke    = codes.IllegalInvoke
	IllegalRequest   = codes.IllegalRequest

	ErrorUserFound         = codes.NewCode(100, "用户名已存在")
	ErrorUserNotFound      = codes.NewCode(101, "用户名或密码不能为空")
	ErrorUserPassword      = codes.NewCode(102, "密码格式错误")
	ErrorUserRegister      = codes.NewCode(103, "注册失败")
	ErrorUserUsernameEmpty = codes.NewCode(104, "用户名不能为空")
	ErrorUserUsernameNot   = codes.NewCode(105, "用户不存在")
	ErrorUserPasswordNot   = codes.NewCode(106, "用户名或密码错误")

	ErrorCateEmpty    = codes.NewCode(110, "分类名不能为空")
	ErrorCateUpdate   = codes.NewCode(111, "分类修改失败")
	ErrorCateNotFound = codes.NewCode(112, "分类不存在")

	ErrorArticleEmpty        = codes.NewCode(120, "标题不能为空")
	ErrorArticleCateEmpty    = codes.NewCode(121, "分类id不能为空")
	ErrorArticleCateNotFound = codes.NewCode(122, "分类不存在")
)
