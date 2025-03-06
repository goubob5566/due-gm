package middleware

import (
	jwtcomp "temp/internal/component/jwt"

	"github.com/dobyte/due/component/http/v2"
	"github.com/dobyte/due/v2/codes"
)

type TokenReq struct {
	Token string `json:"token"`
}

func AuthHandler(ctx http.Context) error {

	req := &TokenReq{}
	err := ctx.Bind().Query(req)
	if err == nil {
		jwt := jwtcomp.Instance()
		// 解析token
		identity, err := jwt.ExtractIdentity(req.Token)
		if err != nil {
			return ctx.Failure(codes.InvalidArgument)
		}
		ctx.Set("token", identity.(string))
		return ctx.Next()
	} else {
		return ctx.Failure(codes.InvalidArgument)
	}

}
