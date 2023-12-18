package crawl

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/gocolly/colly"
	"github.com/olivere/elastic"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// Constants for Elasticsearch and MongoDB
const (
	elasticsearchURL      = "https://113.160.44.242:9200"
	elasticsearchUser     = "admin"
	elasticsearchPassword = "admin"
	mongoDBURL            = "mongodb+srv://hieu:123123abcabc@cluster0.mxjudnq.mongodb.net/"
)

// Product là một struct để biểu diễn thông tin sản phẩm
type New struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Url         string             `bson:"url"`
	Time        string             `bson:"time"`
	Description string             `bson:"description"`
	Content     string             `bson:"content"`
	Sentences   []SentenceInfo     `bson:"sentences"`
	SourceId    primitive.ObjectID `bson:"source_id"`
	CategoryId  string             `bson:"category_id"`
	KeywordId   []string           `bson:"keyword_id"`
	LabelId     string             `bson:"label_id"`
	CrawlTime   time.Time          `bson:"crawl_time"`
	Status      Status             `bson:"status"`
}
type SentenceInfo struct {
	LabelId string `bson:"label_id"`
	Name    string `bson:"name"`
}

type UrlNext struct {
	UrlStart           string `bson:"url_start"`
	BoxElement         string `bson:"box_element"`
	TitleElement       string `bson:"title_element"`
	LinkElement        string `bson:"link_element"`
	TimeElement        string `bson:"time_element"`
	CategoryElement    string `bson:"category_element"`
	DescriptionElement string `bson:"description_element"`
	ContentElement     string `bson:"content_element"`
	CheckTime          string `bson:"check_time"`
	CheckDesc          string `bson:"check_desc"`
	CheckCategory      string `bson:"check_category"`
	Status             Status `bson:"status"`
}
type Status struct {
	Likes   int `bson:"likes"`
	Shares  int `bson:"shares"`
	Reposts int `bson:"reposts"`
	Views   int `bson:"views"`
}

type StartLink struct {
	UrlStart   string `bson:"url_start"`
	CrawlCount int    `bson:"crawl_count"`
}

var categoryData struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

var sourcesData struct {
	ID         primitive.ObjectID `bson:"_id"`
	UrlStart   string             `bson:"url_start"`
	CrawlCount int                `bson:"crawl_count"`
}
var labelData struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

var visitedURLs sync.Map

var result1 []UrlNext
var result2 []StartLink
var visitedURLsList []string

var mu sync.Mutex
var sourceCrawlCount sync.Map

func StartCrawler() {
	esClient, err := connectToElasticsearch()
	if err != nil {
		log.Fatal("Error connecting to Elasticsearch:", err)
	}

	mongoClient, err := connectToMongoDB()
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer mongoClient.Disconnect(context.Background())

	c := colly.NewCollector()

	// Kết nối đến MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://hieu:123123abcabc@cluster0.mxjudnq.mongodb.net/"))
	if err != nil {
		fmt.Println("Lỗi kết nối đến MongoDB:", err)
		return
	}
	defer client.Disconnect(context.TODO())

	// Lựa chọn cơ sở dữ liệu và bảng (collection)
	collection := client.Database("Social_Listening").Collection("dataWeb")
	categoryCollection := client.Database("Social_Listening").Collection("categories")
	sourcesCollection := client.Database("Social_Listening").Collection("sources")
	keywordCollection := client.Database("Social_Listening").Collection("keywordCrawl")
	// collection1 := client.Database("Social_Listening").Collection("linkNexts")
	// collection2 := client.Database("Social_Listening").Collection("sources")
	collection1 := client.Database("Social_Listening").Collection("b")
	collection2 := client.Database("Social_Listening").Collection("a")
	// Tìm kiếm dữ liệu trong MongoDB: UrlStart
	filter2 := map[string]interface{}{}
	cur2, err := collection2.Find(context.TODO(), filter2)
	if err != nil {
		log.Fatal(err)
	}
	defer cur2.Close(context.TODO())

	// Lặp qua kết quả và lấy các trường
	for cur2.Next(context.TODO()) {
		var results StartLink
		err := cur2.Decode(&results)
		if err != nil {
			log.Fatal(err)
		}
		result2 = append(result2, results)
	}
	// Số lượng URL tối đa mà bạn muốn truy cập từ mỗi UrlStart
	maxURLsPerStart := 300
	// Biến đếm số lượng URL đã được xử lý từ mỗi UrlStart
	urlsProcessed := make(map[string]int)

	// Lặp qua các URL Start và in ra các URL đã truy cập cho mỗi URL Start
	for _, startUrl := range result2 {
		urlsProcessed[startUrl.UrlStart] = 0
		visitedURLsList = append(visitedURLsList, startUrl.UrlStart)

		// Định nghĩa cách xử lý khi tìm thấy một trang mới
		c.OnHTML("li a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if !strings.HasPrefix(link, "/") {
				link = e.Attr("href")
			} else {
				link = startUrl.UrlStart + e.Attr("href")
			}

			if isValidURL(link) && urlsProcessed[startUrl.UrlStart] < maxURLsPerStart {
				urlsProcessed[startUrl.UrlStart]++
				e.Request.Visit(link)
				visitedURLsList = append(visitedURLsList, link)
			}
		})

		// Tìm kiếm dữ liệu trong MongoDB: UrlNext
		filter := map[string]interface{}{}
		cur, err := collection1.Find(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(context.TODO())

		// Lặp qua kết quả và lấy cac trường
		for cur.Next(context.TODO()) {
			var result UrlNext
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			result1 = append(result1, result)
		}
		fmt.Println("Url: " + startUrl.UrlStart)
		initialCrawlCount, err := loadInitialCrawlCount(sourcesCollection, startUrl.UrlStart)
		if err != nil {
			fmt.Println("Error loading initial crawl count:", err)
			return
		}
		// LoadOrStore with the initial crawl count
		_, exists := sourceCrawlCount.LoadOrStore(startUrl.UrlStart, initialCrawlCount)
		if !exists {
			fmt.Printf("Initial Crawl count for %s: %d\n", startUrl.UrlStart, initialCrawlCount)
		}
		// Định nghĩa cách xử lý khi bạn đã truy cập một trang web và muốn crawl dữ liệu
		c.OnHTML("html", func(e *colly.HTMLElement) {
			for _, indexUrlNext := range result1 {
				if startUrl.UrlStart == indexUrlNext.UrlStart {
					e.ForEach(indexUrlNext.BoxElement, func(_ int, el *colly.HTMLElement) {
						//Title
						title := el.ChildText(indexUrlNext.TitleElement)

						//URL
						url := el.ChildAttr(indexUrlNext.LinkElement, "href")
						if !strings.HasPrefix(url, "/") {
							url = el.ChildAttr(indexUrlNext.LinkElement, "href")
						} else {
							url = indexUrlNext.UrlStart + el.ChildAttr(indexUrlNext.LinkElement, "href")
						}

						//Desc
						description := ""
						if indexUrlNext.CheckDesc == "1" {
							description = el.ChildText(indexUrlNext.DescriptionElement)
						} else {
							description = getDetailPage(url, indexUrlNext.DescriptionElement)
						}

						//Content
						content := getDetailPage(url, indexUrlNext.ContentElement)
						var sentenceInfos []SentenceInfo
						var sentences []string

						sentences = splitSentences(content)
						// Tách thành các câu

						for _, sentence := range sentences {
							trimmedSentence := strings.TrimSpace(sentence)
							if trimmedSentence != "" {
								sentenceInfo := SentenceInfo{
									LabelId: "",
									Name:    sentence,
								}
								sentenceInfos = append(sentenceInfos, sentenceInfo)
							}
						}

						//TIME
						timer := ""
						if indexUrlNext.CheckTime == "0" {
							//timer = el.ChildAttr(indexUrlNext.TimeElement, "title")
							timer = el.ChildText(indexUrlNext.TimeElement)
							timer = strings.TrimSpace(strings.TrimPrefix(timer, "-"))
						} else if indexUrlNext.CheckTime == "1" {
							timer = el.ChildText(indexUrlNext.TimeElement)
						} else {
							timer = getDetailPage(url, indexUrlNext.TimeElement)
							//timer = strings.TrimSpace(strings.TrimPrefix(timer, "-"))
							//fmt.Println("trang chi tiết")
						}

						//SourceId
						sources := indexUrlNext.UrlStart
						err = sourcesCollection.FindOne(context.Background(), bson.M{"url_start": sources}).Decode(&sourcesData)
						sourceId := sourcesData.ID

						//Category-->CategoryId
						category := ""
						if indexUrlNext.CheckCategory == "1" {
							category = el.ChildText(indexUrlNext.CategoryElement)
						} else {
							category = getDetailPage(url, indexUrlNext.CategoryElement)
						}

						// Thêm dữ liệu vào MongoDB
						if !dataExists(client, title, url, timer, description) {
							saveDataToMongoDB(title, url, timer, description, content, category, sources, sourceId, sentences, sentenceInfos, &visitedURLs, collection, categoryCollection, keywordCollection, sourcesCollection, esClient)
						} else {
							fmt.Println("Dữ liệu đã tồn tại, không thêm vào MongoDB.")
						}
					})
				}
			}
		})

		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
		})

		c.Visit(startUrl.UrlStart)
		c.Wait()
	}

	fmt.Println("Các URL đã truy cập:")
	for _, url := range visitedURLsList {
		fmt.Println(url)
	}
	fmt.Println("Thông tin sản phẩm đã được lưu vào MongoDB")
}

func connectToElasticsearch() (*elastic.Client, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	return elastic.NewClient(
		elastic.SetURL(elasticsearchURL),
		elastic.SetBasicAuth(elasticsearchUser, elasticsearchPassword),
		elastic.SetSniff(false),
		elastic.SetHttpClient(httpClient),
	)
}

// connectToMongoDB establishes a connection to MongoDB and returns the client
func connectToMongoDB() (*mongo.Client, error) {
	return mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDBURL))
}

// Hàm kiểm tra dữ liệu đã tồn tại trong MongoDB
func dataExists(client *mongo.Client, title, url, time, description string) bool {
	collection := client.Database("Social_Listening").Collection("dataWeb")

	filter := bson.M{
		"title": title,
		"url":   url,
	}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return false
	}

	if count > 0 {
		return true
	}

	return false
}

func isValidURL(url string) bool {
	for _, urlNext := range result1 {
		if strings.HasPrefix(url, urlNext.UrlStart+"/") {
			return true
		}
	}
	return false
}

func saveDataToMongoDB(title, url, timer, description, content, category, sources string, sourceId primitive.ObjectID, sentences []string, sentenceInfo []SentenceInfo, visitedURLs *sync.Map, collection *mongo.Collection, categoryCollection *mongo.Collection, keywordCollection *mongo.Collection, sourcesCollection *mongo.Collection, esClient *elastic.Client) {
	if title != "" && url != "" && description != "" && timer != "" {
		// Kiểm tra xem URL đã được lưu vào MongoDB chưa
		if _, ok := visitedURLs.Load(url); !ok {

			// Lấy thời gian crawl hiện tại
			crawlTime := time.Now()

			// Lấy categoryId từ collectionCategory
			err := categoryCollection.FindOne(context.Background(), bson.M{"name": category}).Decode(&categoryData)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					fmt.Println("Không tìm thấy category với name =", category)
					// Không tìm thấy data
					return
				} else {
					fmt.Println("Lỗi khi lấy category từ collectionCategory:", err)
					return
				}
			}
			// Chuyển đổi ObjectID thành chuỗi
			categoryId := categoryData.ID.Hex()

			// Extract keywords from title and description
			keywordIds := extractKeywords(title, description, keywordCollection)
			if len(keywordIds) == 0 {
				fmt.Println("Không có keyword phù hợp, không lưu vào MongoDB.")
				return
			}
			// Lưu dữ liệu vào MongoDB
			pageData := New{
				Title:       title,
				Url:         url,
				Time:        timer,
				Description: description,
				Content:     content,
				Sentences:   sentenceInfo,
				CategoryId:  categoryId,
				SourceId:    sourceId,
				KeywordId:   keywordIds,
				CrawlTime:   crawlTime,
			}

			_, err = collection.InsertOne(context.TODO(), pageData)
			if err != nil {
				fmt.Println("Lỗi khi lưu sản phẩm vào MongoDB:", err)
			}
			mu.Lock()
			crawlCount, ok := sourceCrawlCount.Load(sources)
			if ok {
				updatedCount := crawlCount.(int) + 1
				sourceCrawlCount.Store(sources, updatedCount)
				fmt.Printf("Crawl count for %s: %d\n", sources, updatedCount)
				filter := bson.M{"url_start": sources}
				update := bson.M{"$set": bson.M{"crawl_count": updatedCount}}

				_, err := sourcesCollection.UpdateOne(context.Background(), filter, update)
				if err != nil {
					fmt.Printf("Error updating crawl count for %s in MongoDB: %v\n", sources, err)
				}
			}
			mu.Unlock()

			//Insert từng sentences vào opensearch
			for _, sentence := range sentences {
				var mentionData New

				err = collection.FindOne(context.Background(), bson.M{"url": url}).Decode(&mentionData)
				type SentenceInfo struct {
					Name      string `bson:"name"`
					MentionId string `bson:"mention_id"`
				}
				sentenceInfo1 := SentenceInfo{
					MentionId: mentionData.ID.Hex(),
					Name:      sentence,
				}
				_, err = esClient.Index().
					Index("sentences").
					Type("_doc").
					BodyJson(sentenceInfo1).
					Do(context.Background())
			}
			// Insert dữ liệu vào OpenSearch
			_, err = esClient.Index().
				Index("mentions").
				Type("_doc").
				BodyJson(pageData).
				Do(context.Background())
			if err != nil {
				log.Fatal(err)
			}
		}
		// Đánh dấu URL đã được lưu
		visitedURLs.Store(url, struct{}{})

	}
}

var keywordData struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func extractKeywords(title, description string, keywordCollection *mongo.Collection) []string {
	var keywordIds []string

	// Fetch all keywords from the keyword collection
	filter := bson.M{} // Empty filter to get all keywords
	cur, err := keywordCollection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		if err := cur.Decode(&keywordData); err != nil {
			log.Printf("Lỗi giải mã keywordData: %v\n", err)
			continue
		}

		keyword := keywordData.Name
		keywordId := keywordData.ID
		//fmt.Printf("Title: %s, Description: %s, KeywordData: %+v\n", title, description, keywordData)

		// Check if the keyword exists in the title or description
		if strings.Contains(strings.ToLower(title), strings.ToLower(keyword)) || strings.Contains(strings.ToLower(description), strings.ToLower(keyword)) {
			keywordIds = append(keywordIds, keywordId.Hex())
		}
	}

	return keywordIds
}

func getDetailPage(detailURL, elementSelector string) string {
	detailCollector := colly.NewCollector()

	var result string

	detailCollector.OnHTML(elementSelector, func(e *colly.HTMLElement) {
		result = strings.TrimSpace(e.Text)
	})

	// Truy cập trang chi tiết
	detailCollector.Visit(detailURL)

	return result
}

func loadInitialCrawlCount(sourcesCollection *mongo.Collection, sourceURL string) (int, error) {
	err := sourcesCollection.FindOne(context.Background(), bson.M{"url_start": sourceURL}).Decode(&sourcesData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Printf("Không tìm thấy crawl count cho %s, khởi tạo với 0.\n", sourceURL)
			return 0, nil
		}
		return 0, err
	}
	fmt.Printf("Loaded initial crawl count for %s: %d\n", sourceURL, sourcesData.CrawlCount)
	return sourcesData.CrawlCount, nil
}

func splitSentences(text string) []string {
	text = strings.TrimSpace(text)

	re := regexp.MustCompile(`[.]+|[!?]|\n`)

	// Tách câu bằng cách sử dụng biểu thức chính quy
	sentences := re.Split(text, -1)

	// Kiểm tra và ghép nối các số thập phân với câu trước đó
	for i := 0; i < len(sentences)-1; i++ {
		if len(sentences[i]) > 0 && len(sentences[i+1]) > 0 &&
			unicode.IsDigit(rune(sentences[i][len(sentences[i])-1])) && unicode.IsDigit(rune(sentences[i+1][0])) {
			sentences[i] += "." + sentences[i+1]
			sentences = append(sentences[:i+1], sentences[i+2:]...)
			i--
		}
	}

	// Kiểm tra xem slice có phần tử nào không trước khi loại bỏ khoảng trắng
	if len(sentences) == 0 {
		return sentences
	}

	// Loại bỏ các khoảng trắng đầu và cuối của từng câu
	for i, sentence := range sentences {
		sentences[i] = strings.TrimSpace(sentence)
	}

	return sentences
}
