package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"idist-core/app/collections"
	"idist-core/app/controllers"
	MongoConstant "idist-core/const/mongo-constant"
	"idist-core/const/response"
	"net/http"
	"strconv"
)

//	@Summary	lấy danh sách tỉnh thành
//	@Schemes
//	@Description	do ping
//	@Tags			user
//	@Accept			json
//	@Method			GET
//	@Produce		json
//	@Router			/api/v1/admin/provinces [get]
func ListProvinces(c *gin.Context) {
	data := bson.M{}
	var err error
	entries := collections.Provinces{}
	entry := collections.Province{}

	pagination := controllers.BindRequestTable(c, "_id")

	filter := pagination.CustomFilters(bson.M{})
	if entries, err = entry.Find(filter); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, MongoConstant.NOT_FOUND, err)
	}
	for i := 0; i < len(entries); i++ {
		entries[i].Preload("districts")
	}
	data["entries"] = entries
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

//	@Summary	lấy thông tin chi tiết của tỉnh thành
//	@Schemes
//	@Description	do ping
//	@Tags			user
//	@Accept			json
//	@Method			GET
//	@Param			province_id	query	primitive.ObjectID	true	"province_id"
//	@Produce		json
//	@Router			/api/v1/admin/provinces/{{province_id}} [get]
func ReadProvince(c *gin.Context) {
	data := bson.M{}
	var err error
	entryId, _ := strconv.Atoi(c.Param("id"))
	entry := collections.Province{}
	filter := bson.M{
		"_id": entryId,
	}

	if err = entry.First(filter); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusNotFound, MongoConstant.NOT_FOUND, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

//	@Summary	cập nhật thông tin tỉnh thành
//	@Schemes
//	@Description	do ping
//	@Tags			user
//	@Accept			json
//	@Method			GET
//	@Param			province_id	query	primitive.ObjectID	true	"province_id"
//	@Produce		json
//	@Router			/api/v1/admin/provinces/{{province_id}} [post]
func UpdateProvince(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Province{}
	exist := collections.Province{}
	entryId, _ := strconv.Atoi(c.Param("id"))

	filter := bson.M{
		"_id":        entryId,
		"deleted_at": nil,
	}

	if err = exist.First(filter); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, response.GET_FAIL, err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusNotFound, MongoConstant.NOT_FOUND, err)
	}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, response.GET_FAIL, err)
	}

	if err = entry.Update(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, response.UPDATE_FAIL, err)
	}

	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, response.UPDATE_SUCCESS, data)

}
