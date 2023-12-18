package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/helpers"
	"sync"
	"time"
)

type Account struct {
	ID        primitive.ObjectID  `bson:"_id" json:"id"`
	Name      string              `bson:"name" json:"name"`
	Link      string              `bson:"link" json:"link"`
	Mention   int                 `bson:"mention" json:"mention"`
	Follower  int                 `bson:"follower" json:"follower"`
	CreatedBy primitive.ObjectID  `bson:"created_by" json:"-"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at"`
	UpdatedBy primitive.ObjectID  `bson:"updated_by" json:"-"`
	UpdatedAt time.Time           `bson:"updated_at" json:"updated_at"`
	DeletedBy *primitive.ObjectID `bson:"deleted_by" json:"-"`
	DeletedAt *time.Time          `bson:"deleted_at" json:"-"`

	PlatformId *primitive.ObjectID `bson:"platform_id" json:"platform_id"`

	//Pre-load
	Platform *Platform `bson:"platform" json:"platform"`
}

type Accounts []Account

func (a *Account) CollectionName() string {
	return "accounts"
}

func (a *Account) Find(filter interface{}, opts ...*options.FindOptions) (Accounts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Accounts, 0)
	if cursor, err := DB().Collection(a.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Account
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

func (a *Account) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(a.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&a)
	}
}

func (a *Account) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	a.DeletedAt = helpers.PNow()
	if _, err := DB().Collection(a.CollectionName()).UpdateOne(ctx, bson.M{"_id": a.ID}, bson.M{
		"$set": a,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (a *Account) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	a.ID = primitive.NewObjectID()
	a.CreatedAt = helpers.Now()
	a.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(a.CollectionName()).InsertOne(ctx, a); err != nil {
		return err
	} else {
		return nil
	}
}

func (a *Account) Preload(properties ...string) {
	var wg sync.WaitGroup
	for _, property := range properties {
		if property == "platform" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				entry := Platform{}
				_ = entry.First(bson.M{"_id": a.PlatformId, "deleted_at": nil})
				a.Platform = &entry
			}()
		}
	}
	wg.Wait()
}
