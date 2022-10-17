package controllers

const (
	endpoint_COMMENTS string = "/comments"
	endpoint_COMMENT  string = "/comments/:id"
	endpoint_REPLIES  string = endpoint_COMMENT + "/replies"
	endpoint_REPLY    string = endpoint_REPLIES + "/:id"
)
