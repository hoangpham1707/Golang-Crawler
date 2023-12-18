package admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"idist-core/app/collections"
	"idist-core/app/controllers"
	MongoConstant "idist-core/const/mongo-constant"
	"idist-core/const/response"
	"net/http"
)

func ListOrganizationUnits(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.OrganizationUnit{}
	entries := collections.OrganizationUnits{}
	pagination := controllers.BindRequestTable(c, "created_at")

	filter := pagination.CustomFilters(bson.M{})

	if entries, err = entry.Find(filter); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusInternalServerError, MongoConstant.QUERY_FAIL, err)
	} else if errors.Is(err, mongo.ErrNoDocuments) {
		controllers.ResponseError(c, http.StatusInternalServerError, MongoConstant.NOT_FOUND, err)
	}

	data["entries"] = entries
	data["pagination"] = pagination
	controllers.ResponseSuccess(c, http.StatusOK, response.GET_SUCCESS, data)
}

func CreateOrganizationUnit(c *gin.Context) {

}

func ReadOrganizationUnit(c *gin.Context) {

}

func UpdateOrganizationUnit(c *gin.Context) {

}

func DeleteOrganizationUnit(c *gin.Context) {

}
