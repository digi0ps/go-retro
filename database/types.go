package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type Card struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt int64              `bson:"created_at", json:"created_at"`
}

type Column struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Cards     []Card             `bson:"cards" json:"cards"`
	CreatedAt int64              `bson:"created_at", json:"created_at"`
}

type Board struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Columns   []Column           `bson:"columns" json:"columns"`
	CreatedAt int64              `bson:"created_at", json:"created_at"`
}
