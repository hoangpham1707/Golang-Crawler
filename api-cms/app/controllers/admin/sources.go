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

type Source struct {
	UrlStart string       `form:"url_start" json:"url_start"`
	Avatar   string       `form:"avatar" json:"avatar"`
	Type     string       `form:"typeSource" json:"typeSource"`
	Record   SourceRecord `form:"record" json:"record"`
}

type SourceRecord struct {
	Blocked bool `bson:"blocked" json:"blocked"`
}
type SourceBlocked struct {
	Record SourceRecord `bson:"record" json:"record"`
}

// @Summary	Xem danh sách Source
// @Schemes
// @Description	Xem danh sách Source
// @Tags			Source
// @Accept			json
// @Method			GET
// @Produce		json
// @Success		200
// @Router			/api/v1/common/sources [get]
func ListSource(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Source{}

	pipelineNew := []bson.M{
		{"$match": bson.M{"typeSource": "New"}},
		{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": "$crawl_count"}}},
	}

	pipelineWeb := []bson.M{
		{"$match": bson.M{"typeSource": "Web"}},
		{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": "$crawl_count"}}},
	}

	// Add similar pipelines for other typeSource values

	pipelineFace := []bson.M{
		{"$match": bson.M{"typeSource": "Facebook"}},
		{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": "$crawl_count"}}},
	}

	pipelineVideo := []bson.M{
		{"$match": bson.M{"typeSource": "Video"}},
		{"$group": bson.M{"_id": nil, "count": bson.M{"$sum": "$crawl_count"}}},
	}

	countNew, err := entry.AggregateCountByType(pipelineNew)
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
		return
	}

	countWeb, err := entry.AggregateCountByType(pipelineWeb)
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
		return
	}

	// Add similar code for other typeSource values

	countFace, err := entry.AggregateCountByType(pipelineFace)
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
		return
	}

	countVideo, err := entry.AggregateCountByType(pipelineVideo)
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
		return
	}

	data["countNew"] = countNew
	data["countWeb"] = countWeb
	// Add similar entries for other typeSource values
	data["countFace"] = countFace
	data["countVideo"] = countVideo

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

// @Summary	Xem danh sách Source theo nguồn
// @Schemes
// @Description	Xem danh sách Source theo nguồn
// @Tags			Source
// @Accept			json
// @Method			GET
// @Produce		json
// @Param			typeSource query string false "Type Source"
// @Success		200
// @Router			/api/v1/common/sources/type [get]
func ListTypeSource(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Source{}
	entries := collections.Sources{}

	pagination := BindRequestTable(c, "updated_at")
	filter := pagination.CustomFilters(bson.M{})

	// Thêm điều kiện tìm kiếm theo typeSource
	typeSource := c.Query("typeSource")
	if typeSource != "" {
		filter["typeSource"] = typeSource
	}

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

// @Summary	Tạo mới Source
// @Schemes
// @Description	Tạo mới Source
// @Tags			Source
// @Accept			json
// @Method			POST
// @Parameters     Source
// @Produce		json
// @Param       AccountForm	body		Source	true	"name"
// @Success		200
// @Router			/api/v1/common/sources [post]
func CreateSource(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Source{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu gửi lên không chính xác", err)
	}

	existingSource, err := entry.FindByField("url_start", entry.UrlStart)
	if existingSource != nil {
		controllers.ResponseError(c, http.StatusConflict, "Dữ liệu đã tồn tại", nil)
		return
	}
	if err = entry.Create(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Xử lý dữ liệu lỗi", err)
	}

	data["entry"] = entry

	controllers.ResponseSuccess(c, http.StatusOK, "Tạo mới dữ liệu thành công", data)
}

// @Summary	Xem chi tiết Source
// @Schemes
// @Description	Xem chi tiết Source
// @Tags			Source
// @Accept			json
// @Method			GET
// @Produce		json
// @Param          source_id   path    string	true	"source_id"
// @Success		200
// @Router			/api/v1/common/sources/{source_id} [get]
func GetSource(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Source{}

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

// @Summary	Xóa Source
// @Schemes
// @Description	Xóa Source
// @Tags			Source
// @Accept			json
// @Method			DELETE
// @Produce		json
// @Param          source_id   path    string	true	"source_id"
// @Success		200
// @Router			/api/v1/common/sources/{source_id} [delete]
func DeleteSource(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Source{}

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

// @Summary	Cập nhật Source
// @Description	Cập nhật Source
// @Tags		Source
// @Accept		json
// @Method		PUT
// @Produce		json
// @Param		source_id		path	string	true	"ID of the SOurce"
// @Param		SourceForm	body Source		true	"Thông tin cập nhật của Source"
// @Success		200
// @Router		/api/v1/common/sources/{source_id} [put]
func UpdateSource(c *gin.Context) {
	data := bson.M{}
	var err error
	var entry collections.Source

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
