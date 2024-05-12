package job_repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"worker/internal/domain/entities"
	"worker/internal/infrastructure/mongodb"
	"worker/internal/infrastructure/mongodb/common"
)

type jobRepository struct {
	client *mongodb.DBClient
	coll   *mongo.Collection
}

func NewJobRepository(client *mongodb.DBClient) entities.JobRepository {
	coll := client.Collection("assistant", common.JobCollectionName)
	return &jobRepository{
		client: client,
		coll:   coll,
	}
}

func (j jobRepository) Create(job entities.Job) error {
	//TODO implement me
	panic("implement me")
}

func (j jobRepository) GetByID(id string) (entities.Job, error) {
	//TODO implement me
	panic("implement me")
}

func (j jobRepository) NextID() entities.JobID {
	return entities.JobID(primitive.NewObjectID().Hex())
}
