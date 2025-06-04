package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/PhipattanachaiDev/golang_api-migration/internal/domain"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(col *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{collection: col}
}

func (r *MongoUserRepository) Create(user *domain.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}
func (r *MongoUserRepository) GetByID(id string) (*domain.User, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	var user domain.User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	return &user, err
}
func (r *MongoUserRepository) GetAll() ([]*domain.User, error) {
	cur, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*domain.User
	for cur.Next(context.Background()) {
		var user domain.User
		cur.Decode(&user)
		users = append(users, &user)
	}
	return users, nil
}
func (r *MongoUserRepository) Update(user *domain.User) error {
	objID, _ := primitive.ObjectIDFromHex(user.ID)
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": user})
	return err
}
func (r *MongoUserRepository) Delete(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}