package auth

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"idist-core/app/collections"
	"idist-core/app/controllers"
	"time"
)

type RegisterAccount struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// @Summary	Đăng ký tài khoản
// @Schemes
// @Description	Tạo tài khoản mới
// @Tags			user
// @Accept			json
// @Method			POST
// @Parameters		RegisterAccount
// @Produce		json
// @Param			RegisterForm	body		RegisterAccount	true	"username password"
// @Success		200				{object}	LoginResponse
// @Router			/api/v1/auth/register [post]
func Register(c *gin.Context) {
	// Tạo Account từ POST
	var account RegisterAccount
	var err error
	if err = c.ShouldBind(&account); err != nil {
		controllers.ResponseError(c, 400, "Dữ liệu gửi lên không chính xác", err)
		return
	}

	// Kiểm tra xem tồn tại username
	existingUser := collections.User{}
	filter := bson.M{"username": account.Username}
	if err := existingUser.First(filter); err == nil {
		controllers.ResponseError(c, 409, "Tên người dùng đã tồn tại", nil)
		return
	}

	// Mã hóa password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		controllers.ResponseError(c, 500, "Xử lý mật khẩu lỗi", err)
		return
	}

	// Tạo user mới
	newUser := collections.User{
		Username:  account.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Lưu user
	if err := newUser.Create(); err != nil {
		controllers.ResponseError(c, 500, "Xử lý dữ liệu lỗi", err)
		return
	}

	controllers.ResponseSuccess(c, 201, "Tạo tài khoản thành công", nil)
}
