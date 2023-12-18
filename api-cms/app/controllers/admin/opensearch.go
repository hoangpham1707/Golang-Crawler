package admin

import (
	"encoding/json"
	"fmt"
	"idist-core/app/controllers"
	"log"

	"idist-core/app/providers/mongoProvider"
	openSearchProvider "idist-core/app/providers/opensearchProvider"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListIndexes(c *gin.Context) {
	data := bson.M{}
	client := openSearchProvider.GetOpensearch()

	content := strings.NewReader(`{
		"size": 5,
		"query": {
			"multi_match": {
				"query": "miller",
				"fields": ["title^2", "director"]
			}
		}
	}`)

	search := opensearchapi.SearchRequest{
		Index: []string{"ptit"},
		Body:  content,
	}

	searchResponse, err := search.Do(c, client)
	if err != nil {
		controllers.ResponseError(c, 500, "error", err)
		return
	}

	// Chuyển kết quả tìm kiếm thành cấu trúc dữ liệu JSON
	var searchData map[string]interface{}
	err = json.NewDecoder(searchResponse.Body).Decode(&searchData)
	if err != nil {
		controllers.ResponseError(c, 500, "Parse error", err)
		return
	}

	data["result"] = searchData
	controllers.ResponseSuccess(c, http.StatusOK, "Get info opensearch", data)
}

// func InsertData(c *gin.Context) {
// 	//data := bson.M{}
// 	client := openSearchProvider.GetOpensearch()

// 	document := strings.NewReader(`{
//         "title": "Ptit",
//         "director": "OK",
//         "year": "2011"
//     }`)

// 	docId := "1"
// 	req := opensearchapi.IndexRequest{
// 		Index:      "ptit",
// 		DocumentID: docId,
// 		Body:       document,
// 	}

// 	// Thực hiện InsertRequest và kiểm tra lỗi
// 	insertResponse, err := req.Do(c, client)
// 	if err != nil {
// 		// In lỗi để xem lỗi cụ thể
// 		fmt.Println("Error:", err)
// 		controllers.ResponseError(c, 500, "Insert error", err.Error())
// 		return
// 	}
// 	if err != nil {
// 		fmt.Println("failed to insert document ", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("Inserting a document")
// 	fmt.Println(insertResponse)
// 	defer insertResponse.Body.Close()

// }

func DeleteIndex(c *gin.Context) {

	client := openSearchProvider.GetOpensearch()
	index := c.Query("index")
	deleteIndex := opensearchapi.IndicesDeleteRequest{
		Index: []string{index},
	}

	deleteIndexResponse, err := deleteIndex.Do(c, client)
	if err != nil {
		fmt.Println("failed to delete index ", err)

	}
	fmt.Println("Deleting the index")
	fmt.Println(deleteIndexResponse)
}

func InsertData(c *gin.Context) {
	data := bson.M{}
	mongoDB := mongoProvider.GetMongoDB()
	client := openSearchProvider.GetOpensearch()

	collection := c.Query("collection")
	index := c.Query("index")

	// Assuming you have a MongoDB collection object named "your_collection"
	cursor, err := mongoDB.Collection(collection).Find(c, bson.M{})
	if err != nil {
		controllers.ResponseError(c, http.StatusInternalServerError, "MongoDB query error", err)
		return
	}
	//defer cursor.Close(c)

	for cursor.Next(c) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			controllers.ResponseError(c, http.StatusInternalServerError, "MongoDB decoding error", err)
			return
		}
		delete(result, "_id")
		// Convert the BSON _id to a string
		data, err := json.Marshal(result)
		if err != nil {
			fmt.Printf("Lỗi khi chuyển đổi sang JSON: %v", err)
			controllers.ResponseError(c, http.StatusInternalServerError, "Lỗi khi chuyển đổi sang JSON", err)
			return
		}

		content := strings.NewReader(string(data))
		//fmt.Println("Data: %s", content)
		req := opensearchapi.IndexRequest{
			Index:      index,
			DocumentID: primitive.NewObjectID().Hex(),
			Body:       content,
		}

		// Execute the index request
		_, err = req.Do(c, client)
		if err != nil {
			log.Printf("Lỗi khi chỉ mục hóa tài liệu vào OpenSearch: %v", err)
			controllers.ResponseError(c, http.StatusInternalServerError, "Lỗi khi chỉ mục hóa tài liệu vào OpenSearch", err)
			return
		}
		//log.Printf("Phản hồi từ OpenSearch: %s", insertResponse)
		// Check the response from OpenSearch
	}

	// Respond with success
	controllers.ResponseSuccess(c, http.StatusOK, "MongoDB data inserted into OpenSearch", data)
}

func parseInsertResponse(body io.ReadCloser) (map[string]interface{}, error) {
	var insertedData map[string]interface{}
	err := json.NewDecoder(body).Decode(&insertedData)
	if err != nil {
		return nil, err
	}
	return insertedData, nil
}
