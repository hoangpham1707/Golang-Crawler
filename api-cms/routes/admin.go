package routes

import (
	"idist-core/app/controllers/admin"
	"idist-core/app/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.AuthMiddleware().MiddlewareFunc())
	router.GET("/refresh-token", middlewares.AuthMiddleware().RefreshHandler)
	router.POST("/logout", middlewares.AuthMiddleware().LogoutHandler)
	router.GET("/profile", admin.GetProfile)
	router.POST("/profile", admin.UpdateProfile)

	// Auth
	router.POST("/profile/change-password", middlewares.Gate("", ""), admin.ChangePasswordRequest)
	router.PUT("/profile/change-password", middlewares.Gate("", ""), admin.ChangePasswordConfirm)

	// Basic
	router.GET("/provinces", middlewares.Gate("", ""), admin.ListProvinces)
	router.GET("/provinces/:id", middlewares.Gate("", ""), admin.ReadProvince)
	router.PUT("/provinces/:id", middlewares.Gate("", ""), admin.UpdateProvince)

	router.GET("/districts", middlewares.Gate("", ""), admin.ListDistricts)
	router.GET("/districts/id", middlewares.Gate("", ""), admin.ReadDistrict)
	router.PUT("/districts/id", middlewares.Gate("", ""), admin.UpdateDistrict)

	router.GET("/wards", middlewares.Gate("", ""), admin.ListWards)
	router.GET("/wards/id", middlewares.Gate("", ""), admin.ReadWard)
	router.PUT("/wards/id", middlewares.Gate("", ""), admin.UpdateWard)

	router.GET("/roles", middlewares.Gate("", ""), admin.ListRoles)
	router.POST("/roles", middlewares.Gate("", ""), admin.CreateRole)
	router.GET("/roles/:id", middlewares.Gate("", ""), admin.ReadRole)
	router.PUT("/roles/:id", middlewares.Gate("", ""), admin.UpdateRole)
	router.DELETE("/roles/:id", middlewares.Gate("", ""), admin.DeleteRole)

	router.GET("/users", middlewares.Gate("", ""), admin.ListUsers)
	router.POST("/users", middlewares.Gate("", ""), admin.CreateUser)
	router.GET("/users/:id", middlewares.Gate("", ""), admin.ReadUser)
	router.PUT("/users/:id", middlewares.Gate("", ""), admin.UpdateUser)
	router.DELETE("/users/:id", middlewares.Gate("", ""), admin.DeleteUser)

	// Categories
	// router.GET("/categories", middlewares.Gate("", ""), admin.ListCategories)
	// router.POST("/categories", middlewares.Gate("", ""), admin.CreateCategory)
	// router.GET("/categories/:id", middlewares.Gate("", ""), admin.GetCategory)
	// router.PUT("/categories/:id", middlewares.Gate("", ""), admin.UpdateCategory)
	// router.DELETE("/categories/:id", middlewares.Gate("", ""), admin.DeleteCategory)

	// Tags
	router.GET("/tags", middlewares.Gate("", ""), admin.ListTags)
	router.POST("/tags", middlewares.Gate("", ""), admin.CreateTag)
	router.GET("/tags/:id", middlewares.Gate("", ""), admin.GetTag)
	router.PUT("/tags/:id", middlewares.Gate("", ""), admin.UpdateTag)
	router.DELETE("/tags/:id", middlewares.Gate("", ""), admin.DeleteTag)

	// Articles
	router.GET("/articles", middlewares.Gate("", ""), admin.ListArticles)
	router.POST("/articles", middlewares.Gate("", ""), admin.CreateArticle)
	router.GET("/articles/:id", middlewares.Gate("", ""), admin.GetArticle)
	router.PUT("/articles/:id", middlewares.Gate("", ""), admin.UpdateArticle)
	router.DELETE("/articles/:id", middlewares.Gate("", ""), admin.DeleteArticle)

	// Schools
	// router.GET("/schools", middlewares.Gate("", ""), admin.ListSchools)
	// router.POST("/schools", middlewares.Gate("", ""), admin.CreateSchool)
	// router.GET("/schools/:id", middlewares.Gate("", ""), admin.ReadSchool)
	// router.PUT("/schools/:id", middlewares.Gate("", ""), admin.UpdateSchool)
	// router.DELETE("/schools/:id", middlewares.Gate("", ""), admin.DeleteSchool)

	// Schools
	// router.GET("/admissions", middlewares.Gate("", ""), admin.ListAdmissions)
	// router.POST("/admissions", middlewares.Gate("", ""), admin.CreateAdmission)
	// router.GET("/admissions/:id", middlewares.Gate("", ""), admin.ReadAdmission)
	// router.PUT("/admissions/:id", middlewares.Gate("", ""), admin.UpdateAdmission)
	// router.DELETE("/admissions/:id", middlewares.Gate("", ""), admin.DeleteAdmission)

	// Logs
	router.GET("/logs", middlewares.Gate("", ""), admin.ListLogRecords)

	// Organization Units
	router.GET("/organization-units", middlewares.Gate("", ""), admin.ListOrganizationUnits)
	router.POST("/organization-units", middlewares.Gate("", ""), admin.CreateOrganizationUnit)
	router.GET("/organization-units/:id", middlewares.Gate("", ""), admin.ReadOrganizationUnit)
	router.PUT("/organization-units/:id", middlewares.Gate("", ""), admin.UpdateOrganizationUnit)
	router.DELETE("/organization-units/:id", middlewares.Gate("", ""), admin.DeleteOrganizationUnit)

	// Platforms
	router.GET("/platforms", middlewares.Gate("", ""), admin.ListPlatform)
	router.POST("/platforms", middlewares.Gate("", ""), admin.CreatePlatform)
	router.GET("/platforms/:id", middlewares.Gate("", ""), admin.GetPlatform)
	router.DELETE("/platforms/:id", middlewares.Gate("", ""), admin.DeletePlatform)

	//Accounts
	router.GET("/accounts", middlewares.Gate("", ""), admin.ListAccount)
	router.POST("/accounts", middlewares.Gate("", ""), admin.CreateAccount)
	router.GET("/accounts/:id", middlewares.Gate("", ""), admin.GetAccount)
	router.DELETE("/accounts/:id", middlewares.Gate("", ""), admin.DeleteAccount)

	// //Keywords
	// router.GET("/keywords", middlewares.Gate("", ""), admin.ListKeyword)
	// router.POST("/keywords", middlewares.Gate("", ""), admin.CreateKeyword)
	// router.GET("/keywords/:id", middlewares.Gate("", ""), admin.GetKeyword)
	// router.DELETE("/keywords/:id", middlewares.Gate("", ""), admin.DeleteKeyword)

	//News
	// router.GET("/news", middlewares.Gate("", ""), admin.ListNews)
	// router.POST("/news", middlewares.Gate("", ""), admin.CreateNews)
	// router.GET("/news/:id", middlewares.Gate("", ""), admin.GetNews)
	// router.DELETE("/news/delete/:id", middlewares.Gate("", ""), admin.DeleteNews)
	// router.DELETE("/news/soft-delete/:id", middlewares.Gate("", ""), admin.SoftDeleteNews)

	// Demo opensearch
	router.GET("/opensearch", middlewares.Gate("", ""), admin.ListIndexes)

}
