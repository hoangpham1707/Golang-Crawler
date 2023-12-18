package admin

import (
	"idist-core/app/collections"
	"idist-core/app/controllers"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mention struct {
	Title       string         `bson:"title" json:"title"`
	Url         string         `bson:"url" json:"url"`
	Description string         `bson:"description" json:"description"`
	Content     string         `bson:"content" json:"content"`
	Sentences   []SentenceInfo `bson:"sentences" json:"sentences"`
	SourceId    string         `bson:"source_id" json:"source_id"`
	CategoryId  string         `bson:"category_id" json:"category_id"`
	KeywordId   []string       `bson:"keyword_id" json:"keyword_id"`
	LabelId     string         `bson:"label_id" json:"label_id"`
	CrawlTime   time.Time      `bson:"crawl_Time" json:"crawl_time"`
	Status      Status         `bson:"status" json:"status"`
}

type SentenceInfo struct {
	LabelId string `json:"label_id"`
	Name    string `json:"name"`
}
type Status struct {
	Likes     int    `bson:"likes"`
	Shares    int    `bson:"shares"`
	Sentiment string `bson:"sentiment"`
	Reposts   int    `bson:"reposts"`
	Views     int    `bson:"views"`
}

// @Summary	Xem danh sách Mention
// @Schemes
// @Description	Xem danh sách Mention
// @Tags			Mention
// @Accept			json
// @Method			GET
// @Produce		json
// @Success		200
// @Router			/api/v1/common/mentions [get]
func ListMention(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Mention{}
	entries := collections.Mentions{}

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

// @Summary	Xem danh sách Mention chưa gán nhãn
// @Schemes
// @Description	Xem danh sách Mention chưa gán nhãn
// @Tags			Mention
// @Accept			json
// @Method			GET
// @Produce		json
// @Success		200
// @Router			/api/v1/common/mentions/addLabel [get]
func ListMentionNoLabel(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Mention{}
	//entries = collections.Mentions{}

	paginationWeb := BindRequestTable(c, "crawl_time")
	filterWeb := paginationWeb.CustomFilters(bson.M{"label_id": bson.M{"$eq": ""},
		"source_id": bson.M{"$ne": "655f7463f4b65abe14a88f39"},
	})
	opts := paginationWeb.CustomOptions(options.Find())
	if _, err = entry.Find(filterWeb, opts); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Không tìm thấy dữ liệu", err)
	}

	paginationFace := BindRequestTable(c, "crawl_time")
	filterFace := paginationWeb.CustomFilters(bson.M{"label_id": bson.M{"$eq": ""},
		"source_id": bson.M{"$eq": "655f7463f4b65abe14a88f39"},
	})
	opts1 := paginationFace.CustomOptions(options.Find())
	if _, err = entry.Find(filterFace, opts1); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Không tìm thấy dữ liệu", err)
	}
	paginationWeb.Total, _ = entry.Count(filterWeb)
	paginationFace.Total, _ = entry.Count(filterFace)

	data["totalWeb"] = paginationWeb
	data["totalFace"] = paginationFace

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)

}

// @Summary	Xem danh sách Mention Web chưa gán nhãn
// @Schemes
// @Description	Xem danh sách Mention chưa gán nhãn
// @Tags			Mention
// @Accept			json
// @Method			GET
// @Produce		json
// @Success		200
// @Router			/api/v1/common/mentions/addLabelWeb [get]
func ListMentionWebNoLabel(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Mention{}
	entries := collections.Mentions{}

	pagination := BindRequestTable(c, "crawl_time")
	sourceIDToExclude, err := primitive.ObjectIDFromHex("655f7463f4b65abe14a88f39")
	filter := pagination.CustomFilters(bson.M{"label_id": bson.M{"$eq": ""},
		"source_id": bson.M{"$ne": sourceIDToExclude},
	})
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

// @Summary	Xem danh sách Mention Face chưa gán nhãn
// @Schemes
// @Description	Xem danh sách Mention chưa gán nhãn
// @Tags			Mention
// @Accept			json
// @Method			GET
// @Produce		json
// @Param			account_id query string false "Account ID"
// @Success		200
// @Router			/api/v1/common/mentions/addLabelFace [get]
func ListMentionFaceNoLabel(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Mention{}
	entries := collections.Mentions{}

	pagination := BindRequestTable(c, "crawl_time")

	accountID := c.Query("account_id")

	// Add the condition for label_id and other filters as needed
	filter := pagination.CustomFilters(bson.M{
		"label_id":   bson.M{"$eq": ""},
		"account_id": accountID,
	})

	opts := pagination.CustomOptions(options.Find())
	if entries, err = entry.Find(filter, opts); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
		return
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Không tìm thấy dữ liệu", err)
		return
	}

	pagination.Total, _ = entry.Count(filter)
	data["entries"] = entries
	data["pagination"] = pagination

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

// @Summary	Xem chi tiết Mention
// @Schemes
// @Description	Xem chi tiết Mention
// @Tags			Mention
// @Accept			json
// @Method			GET
// @Produce		json
// @Param          mention_id   path    string	true	"mention_id"
// @Success		200
// @Router			/api/v1/common/mentions/{mention_id} [get]
func GetNew(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Mention{}

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
	//sentences := entry.Sentences
	data["entry"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

// @Summary	Xóa Mention
// @Schemes
// @Description	Xóa Mention
// @Tags			Mention
// @Accept			json
// @Method			DELETE
// @Produce		json
// @Param          mention_id   path    string	true	"mention_id"
// @Success		200
// @Router			/api/v1/common/mentions/{mention_id} [delete]
func DeleteNew(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Mention{}

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

func GetSentencesFromNew(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Mention{}

	entryId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	pagination := BindRequestTable(c, "updated_at")
	filter := pagination.CustomFilters(bson.M{"deleted_at": nil,
		"_id": entryId})

	if err = entry.First(filter); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Không tìm thấy dữ liệu", err)
	}
	//sentences := entry.Sentences
	data["entry"] = entry.Sentences
	data["pagination"] = pagination
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

// @Summary	Gán nhãn cho Mention
// @Description	Gán nhãn cho Mention
// @Tags		Mention
// @Accept		json
// @Method		PUT
// @Produce		json
// @Param		mention_id		path	string	true	"ID of the Mention"
// @Param		label_name	query	string	true	"Name of the Label to assign"
// @Success		200
// @Router		/api/v1/common/mentions/{mention_id}/assign-label [put]
func AssignLabel(c *gin.Context) {
	// newID := c.Param("new_id")
	labelName := c.Query("label_name")
	data := bson.M{}
	var err error

	entry := collections.Mention{}
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

	// Find the Label by name
	labelEntry := collections.Label{}
	if err := labelEntry.First(bson.M{"slug": labelName}); err != nil {
		if err == mongo.ErrNoDocuments {
			controllers.ResponseError(c, http.StatusNotFound, "Label not found", nil)
		} else {
			controllers.ResponseError(c, http.StatusInternalServerError, "Error retrieving Label", err)
		}
		return
	}

	// Assign Label ID to New
	entry.LabelId = labelEntry.Slug

	// Update the New in the database
	if err := entry.Update(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Error updating New", err)
		return
	}
	data["label"] = labelEntry
	//data["mention"] = entry
	controllers.ResponseSuccess(c, http.StatusOK, "Label assigned successfully", data)
}

// @Summary	Gán nhãn cho Sentences
// @Description	Gán nhãn cho Sentences
// @Tags		Mention
// @Accept		json
// @Method		PUT
// @Produce		json
// @Param		mention_id		path	string	true	"ID of the Mention"
// @Param		sentence_index	query	int	true	"Index of the sentence in the Sentences array"
// @Param		label_name	query	string	true	"Name of the Label to assign"
// @Success		200
// @Router		/api/v1/common/mentions/{mention_id}/assign-label-sentences [put]
func AssignLabelSentences(c *gin.Context) {
	// newID := c.Param("new_id")
	labelName := c.Query("label_name")
	sentenceIndex := c.Query("sentence_index")

	//	data := bson.M{}
	var err error
	entry := collections.Mention{}

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

	// Find the Label by name
	labelEntry := collections.Label{}
	if err := labelEntry.First(bson.M{"slug": labelName}); err != nil {
		if err == mongo.ErrNoDocuments {
			controllers.ResponseError(c, http.StatusNotFound, "Label not found", nil)
		} else {
			controllers.ResponseError(c, http.StatusInternalServerError, "Error retrieving Label", err)
		}
		return
	}

	index, err := strconv.Atoi(sentenceIndex)
	if err != nil {
		controllers.ResponseError(c, http.StatusBadRequest, "Invalid sentence index", nil)
		return
	}
	// Assign Label ID to New
	entry.Sentences[index].LabelId = labelEntry.Slug

	// Update the New in the database
	if err := entry.UpdateSentence(); err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Error updating New", err)
		return
	}
	//data["sentences"] = entry.Sentences
	controllers.ResponseSuccess(c, http.StatusOK, "Label assigned successfully", nil)
}

// @Summary	Đếm danh sách Mention đã gán
// @Schemes
// @Description	Xem danh sách Mention
// @Tags			Mention
// @Accept			json
// @Method			GET
// @Produce		json
// @Success		200
// @Router			/api/v1/common/mentions/countLabelId [get]
func ListMentionAddLabel(c *gin.Context) {
	data := bson.M{}

	entry := collections.Mention{}

	pagination := BindRequestTable(c, "updated_at")
	filter := pagination.CustomFilters(bson.M{"label_id": bson.M{"$ne": ""}}) // Chỉ lấy những entry có label_id khác ""

	pagination.Total, _ = entry.Count(filter)
	data["Count"] = pagination
	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}

func CountMention(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Mention{}
	//entries := collections.Mentions{}

	sourceWebIDToExclude, err := primitive.ObjectIDFromHex("655f7463f4b65abe14a88f39")
	paginationWeb := BindRequestTable(c, "updated_at")
	filterWeb := paginationWeb.CustomFilters(bson.M{"label_id": bson.M{"$eq": ""},
		"source_id": bson.M{"$ne": sourceWebIDToExclude},
	})
	optWeb := paginationWeb.CustomOptions(options.Find())
	if _, err = entry.Find(filterWeb, optWeb); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Không tìm thấy dữ liệu", err)
	}

	sourceFaceIDToExclude, err := primitive.ObjectIDFromHex("655f7463f4b65abe14a88f39")
	paginationFace := BindRequestTable(c, "updated_at")
	filterFace := paginationFace.CustomFilters(bson.M{"label_id": bson.M{"$eq": ""},
		"source_id": bson.M{"$eq": sourceFaceIDToExclude},
	})
	optFace := paginationFace.CustomOptions(options.Find())
	if _, err = entry.Find(filterFace, optFace); err != nil && err != mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Truy vấn dữ liệu lỗi", err)
	} else if err == mongo.ErrNoDocuments {
		controllers.ResponseError(c, http.StatusInternalServerError, "Không tìm thấy dữ liệu", err)
	}
	// totalFacebookCrawlCount := 0
	// totalWebCrawlCount := 0

	// for _, entry := range entries {
	// 	if strings.HasPrefix(entry.UrlStart, "https://www.facebook.com") {
	// 		totalFacebookCrawlCount = entry.CrawlCount
	// 	} else {
	// 		totalWebCrawlCount += entry.CrawlCount
	// 	}
	// }
	paginationWeb.Total, _ = entry.Count(filterWeb)
	paginationFace.Total, _ = entry.Count(filterFace)
	data["totalFacebook"] = paginationFace
	data["totalWeb"] = paginationWeb

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)

}

// @Summary	Đếm label của Mention
// @Schemes
// @Description	Đếm label của Mention
// @Tags			Mention
// @Accept			json
// @Method			GET
// @Produce		json
// @Success		200
// @Router			/api/v1/common/mentions/countLabel [get]
func CountLabelMention(c *gin.Context) {
	data := bson.M{}
	var err error
	entry := collections.Mention{}

	countByLabel := func(label string, filter bson.M) (int64, error) {
		pagination := BindRequestTable(c, "updated_at")
		opt := pagination.CustomOptions(options.Find())

		if _, err := entry.Find(filter, opt); err != nil && err != mongo.ErrNoDocuments {
			return 0, err
		} else if err == mongo.ErrNoDocuments {
			return 0, nil
		}

		pagination.Total, _ = entry.Count(filter)
		return pagination.Total, nil
	}

	// Đếm cho label tích cực
	countPos, err := countByLabel("1", bson.M{"label_id": "1"})
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Lỗi khi đếm số lượng", err)
		return
	}
	data["Pos"] = countPos

	// Đếm cho label trung tính
	countNeu, err := countByLabel("2", bson.M{"label_id": "2"})
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Lỗi khi đếm số lượng", err)
		return
	}
	data["Neu"] = countNeu

	// Đếm cho label tiêu cực
	countNeg, err := countByLabel("0", bson.M{"label_id": "0"})
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "Lỗi khi đếm số lượng", err)
		return
	}
	data["Neg"] = countNeg

	controllers.ResponseSuccess(c, http.StatusOK, "Lấy dữ liệu thành công", data)
}
