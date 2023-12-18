package admin

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/collections"
	"idist-core/app/controllers"
	"net/http"
)

//	@Summary	Lấy danh sách bản ghi log
//	@Schemes
//	@Tags		monitor
//	@Accept		json
//	@Method		GET
//	@Produce	json
//	@Params		page query int required
//	@Success	200	{object}	collections.LogRecords
//	@Router		/api/v1/admin/logs [get]
func ListLogRecords(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.LogRecord{}
	entries := collections.LogRecords{}

	pagination := BindRequestTable(c, "created_at")
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
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy danh sách bản ghi lịch sử thành công", data)
}
