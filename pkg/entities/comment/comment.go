package comment

import "time"

type comment struct {
	Id        string `json:"id"`
	Body      string `json:"body"`
	Hash      string `json:"hash"`
	UserId    string `json:"userId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
	Version   int    `json:"version"`
}

func MakeComment(id func() string, body string, userId string, hash func(string) string) (Comment, error) {
	if len(body) == 0 {
		return nil, Err_EMPTY_BODY
	}
	if len(userId) == 0 {
		return nil, Err_NO_USER_ID
	}
	now := time.Now().UTC().Format(time.RFC3339)
	return &comment{
		Id:        id(),
		Body:      body,
		UserId:    userId,
		Hash:      hash(userId + body),
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: "",
		Version:   0,
	}, nil
}

func UpdateComment(
	id string,
	body string,
	userId string,
	hasher func(string) string,
	createdAt string,
	deletedAt string,
	version int,
) (Comment, error) {
	if len(body) == 0 {
		return nil, Err_EMPTY_BODY
	}
	if len(userId) == 0 {
		return nil, Err_NO_USER_ID
	}
	now := time.Now().UTC().Format(time.RFC3339)
	return &comment{
		Id:        id,
		Body:      body,
		UserId:    userId,
		Hash:      hasher(userId + body),
		CreatedAt: createdAt,
		UpdatedAt: now,
		DeletedAt: deletedAt,
		Version:   version,
	}, nil
}

func DeleteComment(
	id string,
	body string,
	userId string,
	hash string,
	createdAt string,
	updatedAt string,
	version int,
) (Comment, error) {
	now := time.Now().UTC().Format(time.RFC3339)
	return &comment{
		Id:        id,
		Body:      body,
		UserId:    userId,
		Hash:      hash,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: now,
		Version:   version,
	}, nil
}
