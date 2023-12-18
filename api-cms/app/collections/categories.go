package collections

import (
	"context"
	"idist-core/helpers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Category struct {
	ID   primitive.ObjectID `bson:"_id" json:"_id"`
	Name string             `bson:"name" json:"name"`

	CreatedBy primitive.ObjectID  `bson:"created_by" json:"-"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at"`
	UpdatedBy primitive.ObjectID  `bson:"updated_by" json:"-"`
	UpdatedAt time.Time           `bson:"updated_at" json:"updated_at"`
	DeletedBy *primitive.ObjectID `bson:"deleted_by" json:"-"`
	DeletedAt *time.Time          `bson:"deleted_at" json:"-"`
}

type Categories []Category

func (u *Category) CollectionName() string {
	return "categories"
}

func (u *Category) Find(filter interface{}, opts ...*options.FindOptions) (Categories, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Categories, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Category
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

func (u *Category) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Category) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).DeleteOne(ctx, bson.M{"_id": u.ID}, options.Delete()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Category) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	u.CreatedAt = helpers.Now()
	u.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Category) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}

}

func (u *Category) FindByField(field string, value interface{}) (*Category, error) {
	filter := bson.M{field: value}
	options := options.FindOne()

	// You can add additional options if needed
	// options.SetProjection(bson.M{"field_name": 1})

	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()

	var result Category
	err := DB().Collection(u.CollectionName()).FindOne(ctx, filter, options).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
