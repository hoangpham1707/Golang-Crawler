package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"idist-core/app/collections"
	"idist-core/app/controllers"
	"net/http"
)

type ChangePassword struct {
	OldPassword     string `json:"old-password"`
	NewPassword     string `json:"new-password"`
	ConfirmPassword string `json:"confirm-password"`
	OTP             string `json:"otp"`
}

func ChangePasswordRequest(c *gin.Context) {
	user := c.MustGet("user").(collections.User)
	data := bson.M{}
	var err error
	entry := ChangePassword{}

	if err = c.ShouldBindBodyWith(&entry, binding.JSON); err != nil {
		controllers.ResponseSuccess(c, http.StatusConflict, "Dữ liệu gửi lên không chính xác", err)
	}

	_ = user.First(bson.M{"_id": user.ID})
	// Kiểm tra mật khẩu cũ
	if user.ComparePassword(entry.OldPassword) == false {
		controllers.ResponseError(c, http.StatusConflict, "Mật khẩu cũ không chính xác", err)
	}

	// Kiểm tra mật khẩu lịch sử
	if user.ComparePasswordHistory(entry.NewPassword) == false {
		controllers.ResponseError(c, http.StatusConflict, "Mật khẩu mới không được trùng với mật khẩu lịch sử", err)
	}

	// Tạo mã OTP
	//helpers.CreateOTP()

	// gửi OTP

	controllers.ResponseSuccess(c, http.StatusOK, "Gửi OTP thành công", data)
}

func ChangePasswordConfirm(c *gin.Context) {
	data := bson.M{}
	controllers.ResponseSuccess(c, http.StatusOK, "API chưa xử lý logic", data)
}
