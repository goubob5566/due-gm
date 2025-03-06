package control

import (
	"temp/app/model"
	"temp/app/pb"
	"temp/internal/code"

	"github.com/dobyte/due/component/http/v2"
	"github.com/dobyte/due/v2/codes"
)

// 创建分类
// @Summary 创建分类
// @Tags 分类
// @Schemes
// @Accept json
// @Produce json
// @Param request body pb.CreateCateReq true "请求参数"
// @Response 200 {object} http.Resp{Data=model.Cate} "响应参数"
// @Router /createcate [post]
func CreateCateHandler(ctx http.Context) error {
	req := &pb.CreateCateReq{}

	if err := ctx.Bind().Form(req); err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	if req.Name == "" {
		return ctx.Failure(code.ErrorCateEmpty)
	}

	// 添加分类
	cate, err := model.CreateCate(req.Name)
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	// 读取分类
	return ctx.Success(cate)
}

// 获取分类
// @Summary 获取分类
// @Tags 分类
// @Schemes
// @Accept json
// @Produce json
// @Response 200 {object} http.Resp{Data=[]model.Cate} "响应参数"
// @Router /readcate [get]
func ReadCateHandler(ctx http.Context) error {

	// 获取所有分类
	cates, err := model.ReadCateAll()
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	return ctx.Success(cates)
}

// 更新分类
// @Summary 更新分类
// @Tags 分类
// @Schemes
// @Accept json
// @Produce json
// @Param request body pb.UpdateCateReq true "请求参数"
// @Response 200 {object} http.Resp{Data=model.Cate} "响应参数"
// @Router /updatecate [post]
func UpdateCateHandler(ctx http.Context) error {
	req := &pb.UpdateCateReq{}

	if err := ctx.Bind().Form(req); err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	if req.Name == "" || req.ID == "" {
		return ctx.Failure(code.ErrorCateUpdate)
	}

	// 获取分类
	cate, err := model.ReadCateOne(req.ID)
	if err != nil || cate == nil {
		return ctx.Failure(code.ErrorCateNotFound)
	}

	// 修改分类
	cate, err = model.UpdateCate(req.ID, req.Name)
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	return ctx.Success(cate)
}

// 删除分类
// @Summary 删除分类
// @Tags 分类
// @Schemes
// @Accept json
// @Produce json
// @Param request body pb.UpdateCateReq true "请求参数"
// @Response 200 {object} http.Resp{Data=string} "响应参数"
// @Router /deletecate [post]
func DeleteCateHandler(ctx http.Context) error {

	req := &pb.UpdateCateReq{}

	if err := ctx.Bind().Form(req); err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	// 删除分类
	err := model.DeleteCate(req.ID)
	if err != nil {
		return ctx.Failure(code.ErrorCateNotFound)
	}

	return ctx.Success(codes.OK)
}
