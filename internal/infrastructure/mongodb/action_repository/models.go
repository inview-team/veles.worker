package action_repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type Action struct {
	ID               primitive.ObjectID `bson:"_id"`
	Name             string             `bson:"name"`
	Type             int                `bson:"type"`
	Input            map[string]Variable
	Output           map[string]Variable
	AdditionalParams map[string]interface{}
}
