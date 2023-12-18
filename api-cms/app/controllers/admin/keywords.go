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

type Keyword struct {
	Name string `form:"name" json:"name"`
}

// @Summary	Xem danh sách Keyword
// @Schemes
// @Description	Xem danh sách Keyword
// @Tags			Keyword
// @Accept			json
// @Method			GET
// @Produce		json
// @Success		200
// @Router			/api/v1/common/keywords [get]
func ListKeyword(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Keyword{}
	entries := collections.Keywords{}

	pagination := BindRequestTable(c, "updated_at")
	filter := pagination.CustomFilters(bson.M{})
	opts := pagination.CustomOptions(options.Find())
	if entries, err = entry.Find(filter, opts); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Không tìm thấy dữ liệu", err)
	}
	pagination.Total, _ = entry.Count(filter)
	data["entries"] = entries
	data["pagination"] = pagination

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

// @Summary	Tạo mới Keyword
// @Schemes
// @Description	Tạo mới Keyword
// @Tags			Keyword
// @Accept			json
// @Method			POST
// @Parameters     Keyword
// @Produce		json
// @Param       AccountForm	body		Keyword	true	"Name"
// @Success		200
// @Router			/api/v1/common/keywords [post]
func CreateKeyword(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Keyword{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu gửi lên không chính xác", err)
	}

	existingKeyword, err := entry.FindByField("name", entry.Name)
	if existingKeyword != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu đã tồn tại", nil)
		return
	}

	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Xử lý dữ liệu lỗi", err)
	}

	data["entry"] = entry

	controllers.ResponseSuccess(c, http.StatusOK, "Tạo mới dữ liệu thành công", data)
}

// @Summary	Xem chi tiết Keyword
// @Schemes
// @Description	Xem chi tiết Keyword
// @Tags			Keyword
// @Accept			json
// @Method			GET
// @Produce		json
// @Param          keyword_id   path    string	true	"Keyword_id"
// @Success		200
// @Router			/api/v1/common/keywords/{keyword_id} [get]
func GetKeyword(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Keyword{}

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

// @Summary	Update Keyword
// @Schemes
// @Description	Sửa Keyword
// @Tags			Keyword
// @Accept			json
// @Method			PUT
// @Produce		json
// @Param          keyword_id   path    string	true	"Keyword_id"
// @Param       KeywordForm	body	Keyword	true	"Thông tin cập nhật của Keyword"
// @Success		200
// @Router			/api/v1/common/keywords/{keyword_id} [put]
func UpdateKeyword(c *gin.Context) {
	data := bson.M{}
	var err error
	var entry collections.Keyword

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

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu gửi lên không đúng", err)
	}

	if err = entry.Update(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, "Cập nhật dữ liệu thành công", data)

}

// @Summary	Xóa Keyword
// @Schemes
// @Description	Xóa Keyword
// @Tags			Keyword
// @Accept			json
// @Method			DELETE
// @Produce		json
// @Param          keyword_id   path    string	true	"Keyword_id"
// @Success		200
// @Router			/api/v1/common/keywords/{keyword_id} [delete]
func DeleteKeyword(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Keyword{}

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
