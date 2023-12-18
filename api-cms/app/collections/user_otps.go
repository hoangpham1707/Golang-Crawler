package collections

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/helpers"
	"time"
)

type UserOTP struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	OTP       string             `bson:"otp_hash" json:"-"`
	TimeLive  int64              `bson:"time_live" json:"time_live"`
	ExpiredAt time.Time          `bson:"expired_at" json:"expired_at"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	DeletedAt *time.Time         `bson:"deleted_at" json:"-"`
	Role      *Role              `bson:"-" json:"role"`
}

type UserOTPs []UserOTP

func (u *UserOTP) CollectionName() string {
	return "user_otps"
}

func (u *UserOTP) String() string {
	uJSON, _ := json.Marshal(u)
	return string(uJSON)
}

func (u *UserOTP) First(filter bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *UserOTP) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	u.CreatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *UserOTP) Find(filter bson.M, opts ...*options.FindOptions) (UserOTPs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(UserOTPs, 0)

	/* Lấy danh sách bản ghi */
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem UserOTP
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

func (u *UserOTP) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *UserOTP) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).DeleteOne(ctx, bson.M{"_id": u.ID}, options.Delete()); err != nil {
		return err
	} else {
		return nil
	}
}
