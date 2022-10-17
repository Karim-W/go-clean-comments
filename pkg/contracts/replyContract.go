package contracts

type CreateReplyContract struct {
	Body     string `json:"body"`
	UserId   string `json:"userId"`
	ParentId string `json:"commentId"`
}
