package action_repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"worker/internal/domain/entities"
	"worker/internal/infrastructure/mongodb"
	"worker/internal/infrastructure/mongodb/common"
)

type actionRepository struct {
	client *mongodb.DBClient
	coll   *mongo.Collection
}

func NewActionRepository(client *mongodb.DBClient) entities.ActionRepository {
	coll := client.Collection("assistant", common.ActionCollectionName)
	return &actionRepository{
		client: client,
		coll:   coll,
	}
}

func (a actionRepository) Create(action entities.Action) error {
	//TODO implement me
	panic("implement me")
}

func (a actionRepository) GetByID(id string) (entities.Action, error) {
	//TODO implement me
	panic("implement me")
}

func (a actionRepository) NextID() entities.ActionID {
	return entities.ActionID(primitive.NewObjectID().Hex())
}
