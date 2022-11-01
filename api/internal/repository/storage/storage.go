package storage

import (
	"context"
	"fmt"
	"log"
	"pokeapi/api/internal/config"
	"pokeapi/api/internal/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewRepository(conf config.Storage) *Repository {
	uri := fmt.Sprintf("%s://%s:%s", conf.Connection, conf.Host, conf.Port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	return &Repository{
		client:     client,
		collection: client.Database(conf.Database).Collection(conf.Collection),
	}
}

func (r Repository) Get(id int) (entities.Pokemon, error) {
	ctx := context.Background()
	query := bson.D{
		primitive.E{Key: "_id", Value: id},
	}

	cursor, err := r.collection.Find(ctx, query)
	if err != nil {
		return entities.Pokemon{}, err
	}

	var p entities.Pokemon
	if cursor.Next(ctx) {
		err = cursor.Decode(&p)
	}
	return p, err
}

func (r Repository) Save(pokemon entities.Pokemon) error {
	ctx := context.Background()
	doc := bson.D{
		primitive.E{Key: "_id", Value: pokemon.ID},
		primitive.E{Key: "name", Value: pokemon.Name},
	}

	_, err := r.collection.InsertOne(ctx, doc)
	return err
}
