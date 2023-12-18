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
)

type Category struct {
	Name string `bson:"name" json:"name"`
}

// @Summary	Xem danh sách Category
// @Schemes
// @Description	Xem danh sách Category
// @Tags			Category
// @Accept			json
// @Method			GET
// @Produce		json
// @Success		200
// @Router			/api/v1/common/categories [get]
func ListCategory(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Category{}
	entries := collections.Categories{}
	// pagination := BindRequestTable(c, "updated_at")
	// filter := pagination.CustomFilters(bson.M{})
	// opts := pagination.CustomOptions(options.Find())
	filter := bson.M{}
	if entries, err = entry.Find(filter); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Không tìm thấy dữ liệu", err)
	}

	data["entries"] = entries
	// data["pagination"] = pagination

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

// @Summary	Tạo mới Category
// @Schemes
// @Description	Tạo mới Category
// @Tags			Category
// @Accept			json
// @Method			POST
// @Parameters     Category
// @Produce		json
// @Param       AccountForm	body		Category	true	"Name"
// @Success		200
// @Router			/api/v1/common/categories [post]
func CreateCategory(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Category{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu gửi lên không chính xác", err)
	}

	// Check if a category with the same name or identifier already exists
	existingCategory, err := entry.FindByField("name", entry.Name)
	// If a category with the same name or identifier exists, return a conflict response
	if existingCategory != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu đã tồn tại", nil)
		return
	}

	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Xử lý dữ liệu lỗi", err)
	}

	data["entry"] = entry

	controllers.ResponseSuccess(c, http.StatusOK, "Tạo mới dữ liệu thành công", data)
}

// @Summary	Xem chi tiết Category
// @Schemes
// @Description	Xem chi tiết Category
// @Tags			Category
// @Accept			json
// @Method			GET
// @Produce		json
// @Param          category_id   path    string	true	"category_id"
// @Success		200
// @Router			/api/v1/common/categories/{category_id} [get]
func GetCategory(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Category{}

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

// @Summary	Xóa Category
// @Schemes
// @Description	Xóa Category
// @Tags			Category
// @Accept			json
// @Method			DELETE
// @Produce		json
// @Param          category_id   path    string	true	"category_id"
// @Success		200
// @Router			/api/v1/common/categories/{category_id} [delete]
func DeleteCategory(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Category{}

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

// @Summary	Cập nhật Category
// @Description	Cập nhật Category
// @Tags		Category
// @Accept		json
// @Method		PUT
// @Produce		json
// @Param		category_id		path	string	true	"ID of the Category"
// @Param		CategoryForm	body Category		true	"Thông tin cập nhật của Category"
// @Success		200
// @Router		/api/v1/common/categories/{category_id} [put]
func UpdateCategory(c *gin.Context) {
	data := bson.M{}
	var err error
	var entry collections.Category

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
		controllers.ResponseError(c, http.StatusBadRequest, "Dữ liệu gửi lên không đúng", err)
		return
	}

	if err = entry.Update(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, "Cập nhật dữ liệu thành công", data)

}
