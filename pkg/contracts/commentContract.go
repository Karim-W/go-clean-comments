package contracts

type CreateCommentContract struct {
	Body      string `json:"body"`
	UserId    string `json:"userId"`
	CreatedAt string `json:"createdAt"`
}
