package control

import (
	"context"
	"temp/app/model"
	"temp/app/pb"

	"temp/internal/code"
	jwtcomp "temp/internal/component/jwt"

	"github.com/dobyte/due/component/http/v2"
	eventbusredis "github.com/dobyte/due/eventbus/redis/v2"

	"github.com/dobyte/due/v2/codes"
	"golang.org/x/crypto/bcrypt"

	"github.com/dobyte/due/transport/grpc/v2"

	rpc "google.golang.org/grpc"
)

// 登陆接口
// @Summary 登陆接口
// @Tags 用户
// @Schemes
// @Accept json
// @Produce json
// @Param request body pb.LoginReq true "请求参数"
// @Response 200 {object} http.Resp{Data=pb.LoginResp} "响应参数"
// @Router /login [post]
func LoginHandler(ctx http.Context) error {

	req := &pb.LoginReq{}

	if err := ctx.Bind().Form(req); err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	if req.Username == "" {
		return ctx.Failure(code.ErrorUserUsernameEmpty)
	}
	if req.Password == "" {
		return ctx.Failure(code.ErrorUserPasswordNot)
	}

	// 查找用户
	user, err := model.GetUserByUsername(req.Username)
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}
	if user == nil {
		return ctx.Failure(code.ErrorUserPasswordNot)
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return ctx.Failure(code.ErrorUserPasswordNot)
	}

	jwt := jwtcomp.Instance()
	token, err := jwt.GenerateToken(jwtcomp.Payload{
		jwt.IdentityKey(): user.UUID,
	})
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	// 登陆成功
	resp := &pb.LoginResp{
		Token:    token.Token,
		Username: user.Username,
	}
	return ctx.Success(resp)
}

// 注册接口
// @Summary 注册接口
// @Tags 用户
// @Schemes
// @Accept json
// @Produce json
// @Param request body pb.RegisterReq true "请求参数"
// @Response 200 {object} http.Resp{Data=model.UUser} "响应参数"
// @Router /register [post]
func RegisterHandler(ctx http.Context) error {

	req := &pb.RegisterReq{}

	if err := ctx.Bind().Form(req); err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	if req.Username == "" || req.Password == "" {
		return ctx.Failure(code.ErrorUserNotFound)
	}

	// 查找用户
	user, err := model.GetUserByUsername(req.Username)
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}
	if user != nil {
		return ctx.Failure(code.ErrorUserFound)
	}

	// 保存用户
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Failure(code.ErrorUserPassword)
	}
	uuser, err := model.RegisterUser(req.Username, password)
	if err != nil {
		return ctx.Failure(code.ErrorUserRegister)
	}

	return ctx.Success(uuser)
}

// 订阅发布测试
// @Summary 订阅发布测试
// @Tags 订阅发布
// @Schemes
// @Accept json
// @Produce json
// @Param request body pb.PubReq true "请求参数"
// @Response 200 {object} http.Resp{Data=string} "响应参数"
// @Router /pub [post]
func PubHandler(ctx http.Context) error {
	req := &pb.PubReq{}

	if err := ctx.Bind().Form(req); err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	eb := eventbusredis.NewEventbus()

	eb.Publish(context.Background(), "speak", req.Message)
	return ctx.Success(codes.OK)
}

func RpcHandler(ctx http.Context) error {

	target := "direct://127.0.0.1:12233"

	transporter := grpc.NewTransporter()
	client, err := transporter.NewClient(target)
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	cli := pb.NewGreeterClient(client.Client().(rpc.ClientConnInterface))

	reply, err := cli.Hello(context.Background(), &pb.HelloArgs{Name: "web"})
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	return ctx.Success(reply)
}
