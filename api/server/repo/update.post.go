package repo

import (
	"context"
	"log"

	model "github.com/trantho123/CRUD-gRPC/api/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i *impl) UpdatePost(ctx context.Context, id primitive.ObjectID, post *model.UpdatePost) error {
	doc, err := ToDoc(post)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": doc,
	}

	_, err = i.client.UpdateByID(ctx, id, update)
	if err != nil {
		log.Println("repo-UpdatePost", err)
		return err
	}

	return nil
}

func ToDoc(v interface{}) (bson.M, error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}

	var doc bson.M
	err = bson.Unmarshal(data, &doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
