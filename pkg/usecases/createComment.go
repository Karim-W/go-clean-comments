package usecases

import (
	"context"

	"github.com/karim-w/go-clean-commments/pkg/data/commentsdb"
	"github.com/karim-w/go-clean-commments/pkg/entities/comment"
	"github.com/karim-w/go-clean-commments/pkg/utils"
)

func CreateComment(
	ctx context.Context,
	body string,
	userId string,
) (int, interface{}) {
	//Create Base entity
	com, err := comment.MakeComment(utils.CreateId, body, userId, utils.CreateMd5Hash)
	if err != nil {
		return 400, err
	}
	// run Validations
	hash, err := commentsdb.GetCommentHash(ctx, com.GetHash())
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return 400, map[string]interface{}{
				"error": err.Error(),
			}
		}
	}
	if hash != "" {
		return 409, map[string]interface{}{
			"error": "Comment Already Exists",
		}
	}
	err = commentsdb.SaveComment(
		ctx,
		com.GetId(),
		com.GetUserId(),
		com.GetBody(),
		com.GetHash(),
		com.GetCreatedAt(),
		com.GetUpdatedAt(),
		com.GetDeletedAt(),
		com.GetVersion(),
	)
	if err != nil {
		return 400, map[string]interface{}{
			"error": err.Error(),
		}
	}
	return 0, com.GetDocument()
}
