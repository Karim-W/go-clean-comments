package replies

type Reply interface {
	GetId() string
	GetBody() string
	GetUserId() string
	GetParentId() string
	GetHash() string
	GetCreatedAt() string
	GetUpdatedAt() string
	GetDeletedAt() string
	GetVersion() int
	GetIsDeleted() bool
	GetDocument() interface{}
}

func (r *reply) GetId() string {
	return r.Id
}

func (r *reply) GetBody() string {
	return r.Body
}

func (r *reply) GetUserId() string {
	return r.UserId
}

func (r *reply) GetParentId() string {
	return r.ParentId
}

func (r *reply) GetHash() string {
	return r.Hash
}

func (r *reply) GetCreatedAt() string {
	return r.CreatedAt
}

func (r *reply) GetUpdatedAt() string {
	return r.UpdatedAt
}

func (r *reply) GetDeletedAt() string {
	return r.DeletedAt
}

func (r *reply) GetVersion() int {
	return r.Version
}

func (r *reply) GetIsDeleted() bool {
	return len(r.DeletedAt) > 0
}

func (r *reply) GetDocument() interface{} {
	return r
}
