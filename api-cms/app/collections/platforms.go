package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/helpers"
	"time"
)

type Platform struct {
	ID        primitive.ObjectID  `bson:"_id" json:"id"`
	Name      string              `bson:"name" json:"name"`
	Link      string              `bson:"link" json:"link"`
	Score     float64             `bson:"score" json:"score"`
	Access    float64             `bson:"access" json:"access"`
	CreatedBy primitive.ObjectID  `bson:"created_by" json:"-"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at"`
	UpdatedBy primitive.ObjectID  `bson:"updated_by" json:"-"`
	UpdatedAt time.Time           `bson:"updated_at" json:"updated_at"`
	DeletedBy *primitive.ObjectID `bson:"deleted_by" json:"-"`
	DeletedAt *time.Time          `bson:"deleted_at" json:"-"`
}

type Platforms []Platform

func (p *Platform) CollectionName() string {
	return "platforms"
}

func (p *Platform) Find(filter interface{}, opts ...*options.FindOptions) (Platforms, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Platforms, 0)
	if cursor, err := DB().Collection(p.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Platform
			if err = cursor.Decode(&elem); err != nil {
				return data, err
			}
			data = append(data, elem)
		}
		if err = cursor.Err(); err != nil {
			return data, err
		}
		return data, cursor.Close(ctx)
	} else {
		return data, err
	}
}

func (p *Platform) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(p.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&p)
	}
}

func (p *Platform) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	p.DeletedAt = helpers.PNow()
	if _, err := DB().Collection(p.CollectionName()).UpdateOne(ctx, bson.M{"_id": p.ID}, bson.M{
		"$set": p,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (p *Platform) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	p.ID = primitive.NewObjectID()
	p.CreatedAt = helpers.Now()
	p.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(p.CollectionName()).InsertOne(ctx, p); err != nil {
		return err
	} else {
		return nil
	}
}
