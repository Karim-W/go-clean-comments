package repliesdb

import (
	"context"
	"encoding/json"
	"time"

	"github.com/karim-w/go-clean-commments/database"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateReply(
	ctx context.Context,
	id string,
	userId string,
	commentId string,
	body string,
	hash string,
	createdAt string,
	updatedAt string,
	deletedAt string,
	version int,
) error {
	e := map[string]interface{}{
		"id":        id,
		"_id":       id,
		"userId":    userId,
		"commentId": commentId,
		"body":      body,
		"hash":      hash,
		"createdAt": createdAt,
		"updatedAt": updatedAt,
		"deletedAt": deletedAt,
		"version":   version,
	}
	byt, err := json.Marshal(e)
	if err != nil {
		return err
	}
	tx := database.GetRedisInstance().TxPipeline()
	_, err = tx.Set(ctx, "replies:"+id, byt, time.Hour).Result()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx)
	if err != nil {
		return err
	}
	sess, err := database.GetDB().StartSession()
	if err != nil {
		return err
	}
	defer sess.EndSession(ctx)
	err = sess.StartTransaction()
	if err != nil {
		return err
	}
	_, err = sess.Client().Database("comments").Collection("replies").InsertOne(ctx, e)
	if err != nil {
		return err
	}
	err = sess.CommitTransaction(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetReplyHash(ctx context.Context, hash string) (string, error) {
	res := database.GetDB().Database("comments").Collection("replies").FindOne(ctx, bson.D{
		{Key: "hash", Value: hash},
	}, nil)

	var comment map[string]interface{}
	err := res.Decode(&comment)
	if err != nil {
		return "", err
	}
	return comment["hash"].(string), nil
}
