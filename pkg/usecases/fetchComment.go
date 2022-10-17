package usecases

import (
	"context"

	"github.com/karim-w/go-clean-commments/pkg/data/commentsdb"
)

func FetchComment(ctx context.Context, id string) (int, interface{}) {
	res, err := commentsdb.GetCommentById(ctx, id)
	if err != nil {
		return 404, map[string]interface{}{
			"error": err.Error(),
		}
	}
	return 200, res
}
