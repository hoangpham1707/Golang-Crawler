package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/collections"
	"idist-core/app/controllers"
	"net/http"
)

type Account struct {
	Name     string `form:"name" json:"name"`
	Platform string `form:"platform" json:"platform"`
	Link     string `form:"link" json:"link"`
}

// @Summary     	Xem danh sách account
// @Schemes
// @Description	Xem danh sách account
// @Tags			account
// @Accept			json
// @Method			GET
// @Produce		json
// @Success		200
// @Router			/api/v1/admin/accounts [get]
func ListAccount(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Account{}
	entries := collections.Accounts{}

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

// @Summary	Tạo mới account
// @Schemes
// @Description	Tạo mới account
// @Tags			account
// @Accept			json
// @Method			POST
// @Parameters     Account
// @Produce		json
// @Param       AccountForm	body		Account	true	"name platform link"
// @Success		200
// @Router			/api/v1/admin/accounts [post]
func CreateAccount(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Account{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu gửi lên không chính xác", err)
	}

	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Xử lý dữ liệu lỗi", err)
	}

	data["entry"] = entry

	controllers.ResponseSuccess(c, http.StatusOK, "Tạo mới dữ liệu thành công", data)
}

// @Summary	Xem chi tiết account
// @Schemes
// @Description	Xem chi tiết account
// @Tags			account
// @Accept			json
// @Method			GET
// @Produce		json
// @Param          account_id   path    string	true	"account_id"
// @Success		200
// @Router			/api/v1/admin/accounts/{account_id} [get]
func GetAccount(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Account{}

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

// @Summary	Xóa account
// @Schemes
// @Description	Xóa account
// @Tags			account
// @Accept			json
// @Method			DELETE
// @Produce		json
// @Parameters     account_id
// @Param          account_id   path    string	true	"account_id"
// @Success		200
// @Router			/api/v1/admin/accounts/{account_id} [delete]
func DeleteAccount(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Account{}

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
