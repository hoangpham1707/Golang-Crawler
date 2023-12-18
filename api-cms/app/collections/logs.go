package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/helpers"
	"time"
)

type LogRecord struct {
	ID          string    `bson:"_id" json:"request_id"`
	Method      string    `bson:"method" json:"method"`
	Latency     string    `bson:"latency" json:"latency"`
	StatusCode  int       `bson:"status_code" json:"status_code"`
	ContentType string    `bson:"content_type" json:"content_type"`
	URI         string    `bson:"uri" json:"uri"`
	IP          string    `bson:"ip" json:"ip"`
	Request     string    `bson:"request" json:"request"`
	Response    string    `bson:"response" json:"response"`
	PreferUrl   string    `bson:"prefer_url" json:"prefer_url"`
	UserAgent   string    `bson:"user_agent" json:"user_agent"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

type LogRecords []LogRecord

func (u *LogRecord) CollectionName() string {
	return "log_records"
}

func (u *LogRecord) First(filter bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *LogRecord) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.CreatedAt = helpers.Now()
	u.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *LogRecord) Update() error {
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
func (u *LogRecord) Find(filter bson.M, opts ...*options.FindOptions) (LogRecords, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(LogRecords, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem LogRecord
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

func (u *LogRecord) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).DeleteOne(ctx, bson.M{"_id": u.ID}); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *LogRecord) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}
