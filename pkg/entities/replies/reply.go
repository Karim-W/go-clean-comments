package replies

import "time"

type reply struct {
	Id        string `json:"id"`
	Body      string `json:"body"`
	UserId    string `json:"userId"`
	ParentId  string `json:"commentId"`
	Hash      string `json:"hash"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
	Version   int    `json:"version"`
}

func MakeReply(id func() string, body string, userId string, parentId string, hash func(string) string) (Reply, error) {
	if len(body) == 0 {
		return nil, Err_EMPTY_BODY
	}
	if len(userId) == 0 {
		return nil, Err_NO_USER_ID
	}
	if len(parentId) == 0 {
		return nil, Err_NO_PARENT_ID
	}
	now := time.Now().UTC().Format(time.RFC3339)
	return &reply{
		Id:        id(),
		Body:      body,
		UserId:    userId,
		ParentId:  parentId,
		Hash:      hash(userId + body + parentId),
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: "",
		Version:   0,
	}, nil
}

func UpdateReply(
	id string,
	body string,
	userId string,
	parentId string,
	hasher func(string) string,
	createdAt string,
	deletedAt string,
	version int,
) (Reply, error) {
	if len(body) == 0 {
		return nil, Err_EMPTY_BODY
	}
	if len(userId) == 0 {
		return nil, Err_NO_USER_ID
	}
	if len(parentId) == 0 {
		return nil, Err_NO_PARENT_ID
	}
	now := time.Now().UTC().Format(time.RFC3339)
	return &reply{
		Id:        id,
		Body:      body,
		UserId:    userId,
		ParentId:  parentId,
		Hash:      hasher(userId + body),
		CreatedAt: createdAt,
		UpdatedAt: now,
		DeletedAt: deletedAt,
		Version:   version,
	}, nil
}

func DeleteReply(
	id string,
	body string,
	userId string,
	parentId string,
	hasher func(string) string,
	createdAt string,
	updatedAt string,
	version int,
) (Reply, error) {
	if len(body) == 0 {
		return nil, Err_EMPTY_BODY
	}
	if len(userId) == 0 {
		return nil, Err_NO_USER_ID
	}
	if len(parentId) == 0 {
		return nil, Err_NO_PARENT_ID
	}
	now := time.Now().UTC().Format(time.RFC3339)
	return &reply{
		Id:        id,
		Body:      body,
		UserId:    userId,
		ParentId:  parentId,
		Hash:      hasher(userId + body),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: now,
		Version:   version,
	}, nil
}
