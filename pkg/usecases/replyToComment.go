package usecases

import (
	"context"

	"github.com/karim-w/go-clean-commments/pkg/data/commentsdb"
	"github.com/karim-w/go-clean-commments/pkg/data/repliesdb"
	"github.com/karim-w/go-clean-commments/pkg/entities/replies"
	"github.com/karim-w/go-clean-commments/pkg/utils"
)

func ReplyToComment(
	ctx context.Context,
	commentId string,
	body string,
	userId string,
) (int, interface{}) {
	comment, err := commentsdb.GetCommentById(ctx, commentId)
	if err != nil {
		return 404, map[string]interface{}{
			"error": err.Error(),
		}
	}
	if comment == nil {
		return 404, map[string]interface{}{
			"error": "Comment Not Found",
		}
	}
	repl, err := replies.MakeReply(utils.CreateId, body, userId, commentId, utils.CreateMd5Hash)
	if err != nil {
		return 400, map[string]interface{}{
			"error": err.Error(),
		}
	}
	hash, err := repliesdb.GetReplyHash(ctx, repl.GetHash())
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return 400, map[string]interface{}{
				"error": err.Error(),
			}
		}
	}
	if hash != "" {
		return 409, map[string]interface{}{
			"error": "Reply Already Exists",
		}
	}
	err = repliesdb.CreateReply(
		ctx,
		repl.GetId(),
		repl.GetUserId(),
		repl.GetParentId(),
		repl.GetBody(),
		repl.GetHash(),
		repl.GetCreatedAt(),
		repl.GetUpdatedAt(),
		repl.GetDeletedAt(),
		repl.GetVersion(),
	)
	if err != nil {
		return 400, map[string]interface{}{
			"error": err.Error(),
		}
	}
	return 201, repl.GetDocument()
}
