package admin

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"idist-core/app/collections"
	"idist-core/app/controllers"
	"net/http"
)

//	@Summary	lấy thông tin profile đang login
//	@Schemes
//	@Description	do ping
//	@Tags			user
//	@Accept			json
//	@Method			GET
//	@Produce		json
//	@Router			/api/v1/admin/profile [get]
func GetProfile(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.User{}
	claims := jwt.ExtractClaims(c)
	userId, _ := primitive.ObjectIDFromHex(claims["id"].(string))

	filter := bson.M{
		"_id":        userId,
		"deleted_at": nil,
	}

	if err = entry.First(filter); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Trích xuất dữ liệu lỗi", nil)
	}
	data["entry"] = entry

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func UpdateProfile(c *gin.Context) {
	data := bson.M{}
	controllers.ResponseSuccess(c, http.StatusOK, "API chưa xử lý logic", data)
}
