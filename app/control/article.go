package control

import (
	"temp/app/model"
	"temp/app/pb"
	"temp/internal/code"

	"github.com/dobyte/due/component/http/v2"
	"github.com/dobyte/due/v2/codes"
)

// 创建文章
// @Summary 创建文章
// @Tags 文章
// @Schemes
// @Accept json
// @Produce json
// @Param request body pb.CreateArticleReq true "请求参数"
// @Response 200 {object} http.Resp{Data=model.Article} "响应参数"
// @Router /createarticle [post]
func CreateArticleHandler(ctx http.Context) error {
	req := &pb.CreateArticleReq{}

	if err := ctx.Bind().Form(req); err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	if req.Title == "" {
		return ctx.Failure(code.ErrorArticleEmpty)
	}
	if req.CateID == "" {
		return ctx.Failure(code.ErrorArticleCateEmpty)
	}

	// 读取分类
	cate, err := model.ReadCateOne(req.CateID)
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}
	if cate == nil {
		return ctx.Failure(code.ErrorArticleCateNotFound)
	}

	// 添加文章
	am, err := model.NewArticleModel()
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	article, err := am.CreateArticle(req.CateID, req.Title)
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	// 返回文章
	return ctx.Success(article)
}

// 读取文章
// @Summary 读取文章
// @Tags 文章
// @Schemes
// @Accept json
// @Produce json
// @Response 200 {object} http.Resp{Data=[]model.Article} "响应参数"
// @Router /readarticle [get]
func ReadArticleHandler(ctx http.Context) error {
	list, err := model.NewArticleModel()
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	return ctx.Success(list.GetArticle())
}

// 更新文章
// @Summary 更新文章
// @Tags 文章
// @Schemes
// @Accept json
// @Produce json
// @Param request body pb.UpdateArticleReq true "请求参数"
// @Response 200 {object} http.Resp{Data=model.Article} "响应参数"
// @Router /updatearticle [post]
func UpdateArticleHandler(ctx http.Context) error {
	req := &pb.UpdateArticleReq{}

	if err := ctx.Bind().Form(req); err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	if req.Title == "" {
		return ctx.Failure(code.ErrorArticleEmpty)
	}
	if req.CateID == "" {
		return ctx.Failure(code.ErrorArticleCateEmpty)
	}

	// 读取分类
	cate, err := model.ReadCateOne(req.CateID)
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}
	if cate == nil {
		return ctx.Failure(code.ErrorArticleCateNotFound)
	}

	// 修改文章
	am, err := model.NewArticleModel()
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	article, err := am.UpdateArticle(req.ID, req.CateID, req.Title)
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	// 返回文章
	return ctx.Success(article)
}

// 删除文章
// @Summary 删除文章
// @Tags 文章
// @Schemes
// @Accept json
// @Produce json
// @Param request body pb.DeleteArticleReq true "请求参数"
// @Response 200 {object} http.Resp{Data=string} "响应参数"
// @Router /deletearticle [post]
func DeleteArticleHandler(ctx http.Context) error {

	req := &pb.DeleteArticleReq{}

	if err := ctx.Bind().Form(req); err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	am, err := model.NewArticleModel()
	if err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	// 删除文章
	err = am.DeleteArticle(req.ID)
	if err != nil {
		return ctx.Failure(code.ErrorCateNotFound)
	}

	return ctx.Success(codes.OK)
}
