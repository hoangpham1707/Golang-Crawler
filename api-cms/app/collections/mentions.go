package collections

import (
	"context"
	"idist-core/helpers"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mention struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Title       string             `bson:"title" json:"title"`
	Url         string             `bson:"url" json:"url"`
	Description string             `bson:"description" json:"description"`
	Content     string             `bson:"content" json:"content"`
	Sentences   []SentenceInfo     `bson:"sentences" json:"sentences"`
	SourceId    string             `bson:"source_id" json:"source_id"`
	CategoryId  string             `bson:"category_id" json:"category_id"`
	KeywordId   []string           `bson:"keyword_id" json:"keyword_id"`
	LabelId     string             `bson:"label_id" json:"label_id"`
	CrawlTime   time.Time          `bson:"crawl_Time" json:"crawl_time"`
	Status      Status             `bson:"status" json:"status"`

	UpdatedBy primitive.ObjectID  `bson:"updated_by" json:"-"`
	UpdatedAt time.Time           `bson:"updated_at" json:"updated_at"`
	DeletedBy *primitive.ObjectID `bson:"deleted_by" json:"-"`
	DeletedAt *time.Time          `bson:"deleted_at" json:"-"`
}
type SentenceInfo struct {
	LabelId string `bson:"label_id"`
	Name    string `bson:"name"`
}
type Status struct {
	Likes   int `bson:"likes"`
	Shares  int `bson:"shares"`
	Reposts int `bson:"reposts"`
	Views   int `bson:"views"`
}

type Mentions []Mention

func (u *Mention) CollectionName() string {
	return "mentions"
}

func (u *Mention) GetSentencesByID(id primitive.ObjectID) ([]SentenceInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()

	var data []SentenceInfo
	filter := bson.M{"_id": id}
	if err := DB().Collection(u.CollectionName()).FindOne(ctx, filter).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func (u *Mention) AssignLabels(id primitive.ObjectID, sentences []SentenceInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"sentences": sentences}}

	_, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (u *Mention) Find(filter interface{}, opts ...*options.FindOptions) (Mentions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Mentions, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Mention
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

func (u *Mention) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Mention) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.DeletedAt = helpers.PNow()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Mention) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	u.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Mention) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()

	_, err := DB().Collection(u.CollectionName()).UpdateOne(
		ctx,
		bson.M{"_id": u.ID},
		bson.M{"$set": bson.M{"label_id": u.LabelId}},
	)

	return err
}

func (u *Mention) UpdateSentence() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()

	_, err := DB().Collection(u.CollectionName()).UpdateOne(
		ctx,
		bson.M{"_id": u.ID},
		bson.M{"$set": bson.M{"sentences": u.Sentences}},
	)

	return err
}
func (n *Mention) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(n.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}
