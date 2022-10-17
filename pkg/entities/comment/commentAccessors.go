package comment

type Comment interface {
	GetId() string
	GetBody() string
	GetUserId() string
	GetHash() string
	GetCreatedAt() string
	GetUpdatedAt() string
	GetDeletedAt() string
	GetVersion() int
	GetIsDeleted() bool
	GetDocument() interface{}
}

func (c *comment) GetId() string {
	return c.Id
}

func (c *comment) GetBody() string {
	return c.Body
}

func (c *comment) GetUserId() string {
	return c.UserId
}

func (c *comment) GetHash() string {
	return c.Hash
}

func (c *comment) GetCreatedAt() string {
	return c.CreatedAt
}

func (c *comment) GetUpdatedAt() string {
	return c.UpdatedAt
}

func (c *comment) GetDeletedAt() string {
	return c.DeletedAt
}

func (c *comment) GetVersion() int {
	return c.Version
}

func (c *comment) GetIsDeleted() bool {
	return len(c.DeletedAt) > 0
}

func (c *comment) GetDocument() interface{} {
	return c
}
