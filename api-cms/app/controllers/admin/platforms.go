package admin

import (
	"idist-core/app/collections"
	"idist-core/app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Platform struct {
	Name string `form:"name" json:"name"`
	Link string `form:"link" json:"link"`
}

type PlatformID struct {
	ID string `form:"id" json:"id"`
}

func ListPlatform(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Platform{}
	entries := collections.Platforms{}

	pagination := BindRequestTable(c, "updated_at")
	filter := pagination.CustomFilters(bson.M{})
	opts := pagination.CustomOptions(options.Find())
	if entries, err = entry.Find(filter, opts); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Không tìm thấy dữ liệu", err)
	}

	data["entries"] = entries
	data["pagination"] = pagination

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func CreatePlatform(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Platform{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu gửi lên không chính xác", err)
	}

	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Xử lý dữ liệu lỗi", err)
	}

	data["entry"] = entry

	controllers.ResponseSuccess(c, http.StatusOK, "Tạo mới dữ liệu thành công", data)
}

func GetPlatform(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Platform{}

	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{
		"deleted_at": nil,
		"_id":        entryId,
	}
	if err = entry.First(filter); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Không tìm thấy dữ liệu", err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func DeletePlatform(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Platform{}

	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{
		"deleted_at": nil,
		"_id":        entryId,
	}
	if err = entry.First(filter); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Không tìm thấy dữ liệu", err)
	}

	if err = entry.Delete(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	}
	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, "Xoá dữ liệu thành công", data)
}
