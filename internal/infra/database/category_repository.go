package database

import (
	"context"
	"log"

	"github.com/lidiagaldino/desafio-backend/internal/domain/entity"
	"github.com/lidiagaldino/desafio-backend/internal/infra/database/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	client *mongo.Client
	db     *mongo.Database
	col    *mongo.Collection
}

func NewCategoryRepository(client *mongo.Client) *CategoryRepository {
	return &CategoryRepository{
    client: client,
    db:     client.Database("anotai"),
    col:    client.Database("anotai").Collection("categories"),
  }
}

func (r *CategoryRepository) FindAll() ([]entity.Category, error) {
	var categories []entity.Category
	cursor, err := r.col.Find(context.TODO(), bson.M{})
	if err!= nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var category model.Category
		if err := cursor.Decode(&category); err!= nil {
			return nil, err
		}
		categories = append(categories, *toEntityCategory(&category))
	}
	return categories, nil
}

func (r *CategoryRepository) Save(category *entity.Category) (*entity.Category, error) {
	result, err := r.col.InsertOne(context.Background(), fromModelCategory(category))
  if err!= nil {
    log.Fatal(err)
    return nil, err
  }
  
  category.ID = result.InsertedID.(primitive.ObjectID).Hex()
  return category, nil
}

func (r *CategoryRepository) FindByID(id string) (*entity.Category, error) {
	var category model.Category
  objectId, err := primitive.ObjectIDFromHex(id)
  if err!= nil {
    return nil, err
  }
  err = r.col.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&category)
  if err!= nil {
    return nil, err
  }

  return toEntityCategory(&category), nil
}

func (r *CategoryRepository) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err!= nil {
		return err
	}
	_, err = r.col.DeleteOne(context.Background(), bson.M{"_id": objectId})
	if err!= nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) Update(category *entity.Category) (*entity.Category, error) {
	objectId, err := primitive.ObjectIDFromHex(category.ID)
  if err!= nil {
    return nil, err
  }
  model := fromModelCategory(category)
  model.ID = objectId
  _, err = r.col.UpdateOne(context.Background(), bson.M{"_id": objectId}, bson.M{"$set": model})
  if err!= nil {
    return nil, err
  }

  return toEntityCategory(model), nil
}

func fromModelCategory(category *entity.Category) *model.Category {
	return &model.Category{
    ID:   primitive.NewObjectID(),
    Title: category.Title,
		Description: category.Description,
		OwnerID: category.OwnerID,
  }
}

func toEntityCategory(category *model.Category) *entity.Category {
	return &entity.Category{
		ID:   category.ID.Hex(),
		Title: category.Title,
		Description: category.Description,
		OwnerID: category.OwnerID,
	}
}
