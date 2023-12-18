package collections

import (
	"context"
	"idist-core/helpers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LinkNext struct {
	ID                 primitive.ObjectID `bson:"_id" json:"id"`
	BoxElement         string             `bson:"box_element" json:"box_element"`
	TitleElement       string             `bson:"title_element" json:"title_element"`
	CategoryElement    string             `bson:"category_element" json:"category_element"`
	ContentElement     string             `bson:"content_element" json:"content_element"`
	LinkElement        string             `bson:"link_element" json:"link_element"`
	TimeElement        string             `bson:"time_element" json:"time_element"`
	DescriptionElement string             `bson:"description_element" json:"description_element"`
	UrlStart           string             `bson:"url_start" json:"url_start"`
	CheckTime          string             `bson:"check_time" json:"check_time"`
	CheckDesc          string             `bson:"check_desc" json:"check_desc"`
	CheckCategory      string             `bson:"check_category" json:"check_category"`

	CreatedBy primitive.ObjectID  `bson:"created_by" json:"-"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at"`
	UpdatedBy primitive.ObjectID  `bson:"updated_by" json:"-"`
	UpdatedAt time.Time           `bson:"updated_at" json:"updated_at"`
	DeletedBy *primitive.ObjectID `bson:"deleted_by" json:"-"`
	DeletedAt *time.Time          `bson:"deleted_at" json:"-"`
}

type LinkNexts []LinkNext

func (u *LinkNext) CollectionName() string {
	return "linkNexts"
}

func (u *LinkNext) Find(filter interface{}, opts ...*options.FindOptions) (LinkNexts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(LinkNexts, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem LinkNext
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

func (u *LinkNext) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *LinkNext) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).DeleteOne(ctx, bson.M{"_id": u.ID}, options.Delete()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *LinkNext) Create() error {
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

func (u *LinkNext) Update() error {
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

func (n *LinkNext) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(n.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}
