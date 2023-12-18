package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type OrganizationUnit struct {
	ID        int64     `bson:"_id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Alias     string    `bson:"alias" json:"alias"`
	Priority  int64     `bson:"priority" json:"priority"`
	IsShow    bool      `bson:"is_show" json:"is_show"`
	UpdatedAt time.Time `bson:"updated_at" json:"-"`
	Wards     []Ward    `bson:"-" json:"wards"`
}

type OrganizationUnits []OrganizationUnit

func (u *OrganizationUnit) CollectionName() string {
	return "organization_units"
}

func (u *OrganizationUnit) Find(filter interface{}, opts ...*options.FindOptions) (OrganizationUnits, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(OrganizationUnits, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem OrganizationUnit
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

func (u *OrganizationUnit) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *OrganizationUnit) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *OrganizationUnit) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *OrganizationUnit) Preload(properties ...string) {
	var wg sync.WaitGroup
	for _, property := range properties {
		if property == "wards" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				entry := Ward{}
				entries := Wards{}
				entries, _ = entry.Find(bson.M{"OrganizationUnit_id": u.ID})
				u.Wards = entries
			}()
		}
	}
	wg.Wait()

}
