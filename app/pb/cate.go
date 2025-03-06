package pb

type CreateCateReq struct {
	Name string `json:"name"`
}

type UpdateCateReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DeleteCateReq struct {
	ID string `json:"id"`
}
