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

type Label struct {
	Name string `bson:"name" json:"name"`
	Slug string `bson:"name" json:"slug"`
}

// @Summary	Xem danh sách Label
// @Schemes
// @Description	Xem danh sách Label
// @Tags			Label
// @Accept			json
// @Method			GET
// @Produce		json
// @Success		200
// @Router			/api/v1/common/labels [get]
func ListLabel(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Label{}
	entries := collections.Labels{}

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

// @Summary	Tạo mới Label
// @Schemes
// @Description	Tạo mới Label
// @Tags			Label
// @Accept			json
// @Method			POST
// @Parameters     ListLabel
// @Produce		json
// @Param       LabelForm	body		Label	true	"name"
// @Success		200
// @Router			/api/v1/common/labels [post]
func CreateLabel(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Label{}

	// Bind the request body to the entry variable
	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu gửi lên không chính xác", err)
		return
	}

	// Check if any field already exists
	existingNameLabel, err := entry.FindByField("name", entry.Name)
	existingSlugLabel, err := entry.FindByField("slug", entry.Slug)
	if existingNameLabel != nil || existingSlugLabel != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu đã tồn tại", nil)
		return
	}

	// If the label is unique, proceed to create it
	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Xử lý dữ liệu lỗi", err)
		return
	}

	data["entry"] = entry

	controllers.ResponseSuccess(c, http.StatusOK, "Tạo mới từ khoá thành công", data)
}

// @Summary	Xem chi tiết Label
// @Schemes
// @Description	Xem chi tiết Label
// @Tags			Label
// @Accept			json
// @Method			GET
// @Produce		json
// @Param          label_id   path    string	true	"label_id"
// @Success		200
// @Router			/api/v1/common/labels/{label_id} [get]
func GetLabel(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Label{}

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

// @Summary	Update Label
// @Schemes
// @Description	Sửa Label
// @Tags			Label
// @Accept			json
// @Method			PUT
// @Produce		json
// @Param          label_id   path    string	true	"label_id"
// @Param       LabelForm	body	Label	true	"Thông tin cập nhật của Label"
// @Success		200
// @Router			/api/v1/common/labels/{label_id} [put]
func UpdateLabel(c *gin.Context) {
	data := bson.M{}
	var err error

	entry := collections.Label{}
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

// @Summary	Xóa Label
// @Schemes
// @Description	Xóa Label
// @Tags			Label
// @Accept			json
// @Method			DELETE
// @Produce		json
// @Param          label_id   path    string	true	"label_id"
// @Success		200
// @Router			/api/v1/common/labels/{label_id} [delete]
func DeleteLabel(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Label{}

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
	controllers.ResponseSuccess(c, http.StatusOK, "Xoá từ khoá thành công", data)
}
