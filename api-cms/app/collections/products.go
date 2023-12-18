package collections

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	Name          string             `bson:"name" json:"name"`
	Description   string             `bson:"description" json:"description"`
	Thumbnail     string             `bson:"thumbnail" json:"thumbnail"`
	RawPrice      float64            `bson:"raw_price" json:"raw_price"`
	DiscountPrice float64            `bson:"discount_price" json:"discount_price"`
	Discount      float64            `bson:"discount" json:"discount"`
	Quantity      int64              `bson:"quantity" json:"quantity"`
	Unit          string             `bson:"unit" json:"unit"`
	Brand         string             `bson:"brand" json:"brand"`
	Origin        string             `bson:"origin" json:"origin"`
	Images        []string           `bson:"images" json:"images"`
	IsNew         bool               `bson:"is_new" json:"is_new"`
	IsHot         bool               `bson:"is_hot" json:"is_hot"`
	IsSale        bool               `bson:"is_sale" json:"is_sale"`
	IsBestSeller  bool               `bson:"is_best_seller" json:"is_best_seller"`
	IsAvailable   bool               `bson:"is_available" json:"is_available"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt     *time.Time         `bson:"deleted_at" json:"-"`
}

type Products []Product

func (p *Product) CollectionName() string {
	return "Products"
}

func (p *Product) String() string {
	uJSON, _ := json.Marshal(p)
	return string(uJSON)
}
