package request

type ReqUpdateAllId struct {
	IDs []int `json:"ids" binding:"required"`
}
