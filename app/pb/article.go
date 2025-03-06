package pb

type CreateArticleReq struct {
	CateID string `json:"cateid"`
	Title  string `json:"title"`
}

type UpdateArticleReq struct {
	ID     string `json:"id"`
	CateID string `json:"cateid"`
	Title  string `json:"title"`
}

type DeleteArticleReq struct {
	ID string `json:"id"`
}
