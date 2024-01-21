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

type ProductRepository struct {
	client *mongo.Client
	db     *mongo.Database
	col    *mongo.Collection
}

func NewProductRepository(client *mongo.Client) *ProductRepository {
	return &ProductRepository{
		client: client,
		db:     client.Database("anotai"),
		col:    client.Database("anotai").Collection("products"),
	}
}

func (r *ProductRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
  cursor, err := r.col.Find(context.TODO(), bson.M{})
  if err!= nil {
    return nil, err
  }
  defer cursor.Close(context.TODO())
  for cursor.Next(context.TODO()) {
    var product model.Product
    if err := cursor.Decode(&product); err!= nil {
      return nil, err
    }
    products = append(products, *toEntityProduct(&product))
  }
  return products, nil
}

func (r *ProductRepository) Save(product *entity.Product) (*entity.Product, error) {
	result, err := r.col.InsertOne(context.Background(), fromModelProduct(product))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	
	product.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return product, nil
}

func (r *ProductRepository) FindByID(id string) (*entity.Product, error) {
	var product model.Product
	objectId, err := primitive.ObjectIDFromHex(id)
	if err!= nil {
    return nil, err
  }
	err = r.col.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&product)
	if err!= nil {
		return nil, err
	}

	return toEntityProduct(&product), nil
}

func (r *ProductRepository) Update(product *entity.Product) (*entity.Product, error) {
	objectId, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		return nil, err
	}
	model := fromModelProduct(product)
	model.ID = objectId
	_, err = r.col.UpdateOne(context.Background(), bson.M{"_id": objectId}, bson.M{"$set": model})
	if err != nil {
		return nil, err
	}

	return toEntityProduct(model), nil
}
 
func (r *ProductRepository) Delete(id string) error {
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

func fromModelProduct(product *entity.Product) *model.Product {
	return &model.Product{
		ID: primitive.NewObjectID(),
    Title:       product.Title,
    Price:       product.Price,
    Description: product.Description,
    OwnerID:     product.OwnerID,
    CategoryID:  product.CategoryID,
  }
}

func toEntityProduct(product *model.Product) *entity.Product {
	return &entity.Product{
		ID: product.ID.Hex(),
		Title:       product.Title,
		Price:       product.Price,
		Description: product.Description,
		OwnerID:     product.OwnerID,
		CategoryID:  product.CategoryID,
	}
}
