package jobs

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"idist-core/app/collections"
	"idist-core/app/providers/loggerProvider"
	"idist-core/helpers"
	"time"
)

type CleanRequestLog struct {
}

func (u *CleanRequestLog) Run() {
	loggerProvider.GetLogger().Info("CleanRequestLog: Starting clean log record...")
	getEntries := collections.LogRecords{}
	otherEntries := collections.LogRecords{}
	entry := collections.LogRecord{}
	var err error
	filter := bson.M{
		"method": "GET",
		"created_at": bson.M{
			"$lte": helpers.Now().Add(time.Minute * time.Duration(-24*60)),
		},
	}
	opts := options.Find()
	if getEntries, err = entry.Find(filter, opts); err != nil {
		loggerProvider.GetLogger().Error("CleanRequestLog: find dữ liệu lỗi", zap.Error(err))
	}

	// Xóa các bản ghi GET method sau 1 ngày
	for _, getEntry := range getEntries {
		// remove file
		_ = getEntry.Delete()
	}

	// Xóa các bản ghi khác GET method sau 1 tháng
	filterOther := bson.M{
		"method": bson.M{
			"$ne": "GET",
		},
		"created_at": bson.M{
			"$lte": helpers.Now().Add(time.Hour * time.Duration(-24*30)),
		},
	}
	if otherEntries, err = entry.Find(filterOther, opts); err != nil {
		loggerProvider.GetLogger().Error("CleanRequestLog: find dữ liệu lỗi", zap.Error(err))
	}
	for _, otherEntry := range otherEntries {
		_ = otherEntry.Delete()
	}
}
