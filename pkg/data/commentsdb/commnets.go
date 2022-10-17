package commentsdb

import (
	"context"
	"encoding/json"
	"time"

	"github.com/karim-w/go-clean-commments/database"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveComment(
	ctx context.Context,
	id string,
	userId string,
	body string,
	hash string,
	createdAt string,
	updatedAt string,
	deletedAt string,
	version int,
) error {
	tx := database.GetRedisInstance().TxPipeline()
	e := map[string]interface{}{
		"id":        id,
		"userId":    userId,
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
	_, err = tx.Set(ctx, "comments:"+id, byt, time.Hour).Result()
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
	ent := bson.D{
		{Key: "id", Value: id},
		{Key: "_id", Value: id},
		{Key: "userId", Value: userId},
		{Key: "body", Value: body},
		{Key: "hash", Value: hash},
		{Key: "createdAt", Value: createdAt},
		{Key: "updatedAt", Value: updatedAt},
		{Key: "deletedAt", Value: deletedAt},
		{Key: "version", Value: version},
	}
	_, err = sess.Client().Database("comments").Collection("comments").InsertOne(ctx, ent)
	if err != nil {
		return err
	}
	return sess.CommitTransaction(ctx)
}

func UpdateComment(
	ctx context.Context,
	id string,
	userId string,
	body string,
	hash string,
	createdAt string,
	updatedAt string,
	deletedAt string,
	version int,
) error {
	sess, err := database.GetDB().StartSession()
	if err != nil {
		return err
	}
	defer sess.EndSession(ctx)
	err = sess.StartTransaction()
	if err != nil {
		return err
	}
	ent := bson.D{
		{Key: "id", Value: id},
		{Key: "_id", Value: id},
		{Key: "userId", Value: userId},
		{Key: "body", Value: body},
		{Key: "hash", Value: hash},
		{Key: "createdAt", Value: createdAt},
		{Key: "updatedAt", Value: updatedAt},
		{Key: "deletedAt", Value: deletedAt},
		{Key: "version", Value: version},
	}
	_, err = sess.Client().Database("comments").Collection("comments").UpdateOne(ctx, map[string]interface{}{
		"id": id,
	}, ent)
	if err != nil {
		return err
	}
	return sess.CommitTransaction(ctx)
}

func GetComment(
	ctx context.Context,
	id string,
) (map[string]interface{}, error) {
	res := database.GetDB().Database("comments").Collection("comments").FindOne(ctx, map[string]interface{}{
		"id": id,
	})
	var comment map[string]interface{}
	err := res.Decode(&comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func GetCommentVersion(
	ctx context.Context,
	id string,
) (int, error) {
	res := database.GetDB().Database("comments").Collection("comments").FindOne(ctx, map[string]interface{}{
		"id": id,
	})
	var comment map[string]interface{}
	err := res.Decode(&comment)
	if err != nil {
		return 0, err
	}
	return comment["version"].(int), nil
}

func GetCommentHash(
	ctx context.Context,
	hash string,
) (string, error) {
	res := database.GetDB().Database("comments").Collection("comments").FindOne(ctx, bson.D{
		{Key: "hash", Value: hash},
	})
	var comment map[string]interface{}
	err := res.Decode(&comment)
	if err != nil {
		return "", err
	}
	return comment["hash"].(string), nil
}

func GetCommentById(
	ctx context.Context,
	id string,
) (map[string]interface{}, error) {
	var comment map[string]interface{}
	resStr, _ := database.GetRedisInstance().Get(ctx, "comments:"+id).Result()
	if len(resStr) > 2 {
		err := json.Unmarshal([]byte(resStr), &comment)
		if err == nil {
			return comment, nil
		}
	}
	res := database.GetDB().Database("comments").Collection("comments").FindOne(ctx, bson.D{
		{Key: "id", Value: id},
	})
	err := res.Decode(&comment)
	if err != nil {
		return nil, err
	}
	go func() {
		tx := database.GetRedisInstance().TxPipeline()
		byt, err := json.Marshal(comment)
		if err != nil {
			return
		}
		_, err = tx.Set(ctx, "comments:"+id, byt, time.Hour).Result()
		if err != nil {
			return
		}
		_, err = tx.Exec(ctx)
		if err != nil {
			return
		}
	}()
	return comment, nil
}

// Path: pkg/usecases/createComment.go
