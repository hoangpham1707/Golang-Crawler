package collections

import (
	"context"
	"idist-core/helpers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Source struct {
	ID         primitive.ObjectID  `bson:"_id" json:"id"`
	UrlStart   string              `bson:"url_start" json:"url_start"`
	Avatar     string              `bson:"avatar" json:"avatar"`
	Type       string              `bson:"typeSource" json:"typeSource"`
	Record     SourceRecord        `bson:"record" json:"record"`
	CrawlCount int                 `bson:"crawl_count" json:"crawl_count"`
	CreatedBy  primitive.ObjectID  `bson:"created_by" json:"-"`
	CreatedAt  time.Time           `bson:"created_at" json:"created_at"`
	UpdatedBy  primitive.ObjectID  `bson:"updated_by" json:"-"`
	UpdatedAt  time.Time           `bson:"updated_at" json:"updated_at"`
	DeletedBy  *primitive.ObjectID `bson:"deleted_by" json:"-"`
	DeletedAt  *time.Time          `bson:"deleted_at" json:"-"`
}
type SourceRecord struct {
	Views        int     `bson:"views" json:"views"`
	Blocked      bool    `bson:"blocked" json:"blocked"`
	TotalLikes   int     `bson:"total_likes" json:"total_likes"`
	TotalShares  int     `bson:"total_shares" json:"total_shares"`
	TotalReposts int     `bson:"total_reposts" json:"total_reposts"`
	TotalViews   int     `bson:"total_views" json:"total_views"`
	ReactRate    float32 `bson:"react_rate" json:"react_rate"`
}
type Sources []Source

func (u *Source) CollectionName() string {
	return "sources"
}

func (u *Source) Find(filter interface{}, opts ...*options.FindOptions) (Sources, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Sources, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Source
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

func (u *Source) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Source) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).DeleteOne(ctx, bson.M{"_id": u.ID}, options.Delete()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Source) Update() error {
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

func (u *Source) Create() error {
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
func (n *Source) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(n.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *Source) FindByField(field string, value interface{}) (*Source, error) {
	filter := bson.M{field: value}
	options := options.FindOne()

	// You can add additional options if needed
	// options.SetProjection(bson.M{"field_name": 1})

	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()

	var result Source
	err := DB().Collection(u.CollectionName()).FindOne(ctx, filter, options).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (u *Source) AggregateCountByType(pipeline interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()

	cursor, err := DB().Collection(u.CollectionName()).Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var result struct {
		Count int64 `bson:"count"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, err
		}
	}

	return result.Count, nil
}
