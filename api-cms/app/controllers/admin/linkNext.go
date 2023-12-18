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

type LinkNext struct {
	BoxElement         string `form:"box_element" json:"box_element"`
	TitleElement       string `form:"title_element" json:"title_element"`
	CategoryElement    string `form:"category_element" json:"category_element"`
	ContentElement     string `form:"content_element" json:"content_element"`
	LinkElement        string `form:"link_element" json:"link_element"`
	TimeElement        string `form:"time_element" json:"time_element"`
	DescriptionElement string `form:"description_element" json:"description_element"`
	UrlStart           string `form:"url_start" json:"url_start"`
	CheckTime          string `form:"check_time" json:"check_time"`
	CheckDesc          string `form:"check_desc" json:"check_desc"`
	CheckCategory      string `form:"check_category" json:"check_category"`
}

// @Summary	Xem danh sách LinkNext
// @Schemes
// @Description	Xem danh sách LinkNext
// @Tags			LinkNext
// @Accept			json
// @Method			GET
// @Produce		json
// @Success		200
// @Router			/api/v1/common/linkNexts [get]
func ListLinkNext(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.LinkNext{}
	entries := collections.LinkNexts{}

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

// @Summary	Tạo mới LinkNext
// @Schemes
// @Description	Tạo mới LinkNext
// @Tags			LinkNext
// @Accept			json
// @Method			POST
// @Parameters     LinkNext
// @Produce		json
// @Param       LinkNextForm	body		LinkNext	true	"Name"
// @Success		200
// @Router			/api/v1/common/linkNexts [post]
func CreateLinkNext(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.LinkNext{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu gửi lên không chính xác", err)
	}

	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Xử lý dữ liệu lỗi", err)
	}

	data["entry"] = entry

	controllers.ResponseSuccess(c, http.StatusOK, "Tạo mới dữ liệu thành công", data)
}

// @Summary	Xem chi tiết LinkNext
// @Schemes
// @Description	Xem chi tiết LinkNext
// @Tags			LinkNext
// @Accept			json
// @Method			GET
// @Produce		json
// @Param          linkNext_id   path    string	true	"linkNext_id"
// @Success		200
// @Router			/api/v1/common/linkNexts/{linkNext_id} [get]
func GetLinkNext(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.LinkNext{}

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

// @Summary	Xóa LinkNext
// @Schemes
// @Description	Xóa LinkNext
// @Tags			LinkNext
// @Accept			json
// @Method			DELETE
// @Produce		json
// @Param          linkNext_id   path    string	true	"linkNext_id"
// @Success		200
// @Router			/api/v1/common/linkNexts/{linkNext_id} [delete]
func DeleteLinkNext(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.LinkNext{}

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

// @Summary	Cập nhật LinkNext
// @Schemes
// @Description	Cập nhật LinkNext
// @Tags			LinkNext
// @Accept			json
// @Method			PUT
// @Produce		json
// @Param       linkNext_id	path	string	true	"linkNext_id của LinkNext cần cập nhật"
// @Param       LinkNextForm	body	LinkNext	true	"Thông tin cập nhật của LinkNext"
// @Success		200
// @Router			/api/v1/common/linkNexts/{linkNext_id} [put]
func UpdateLinkNext(c *gin.Context) {
	data := bson.M{}
	var err error
	var entry collections.LinkNext

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
