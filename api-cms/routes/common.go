package routes

import (
	"idist-core/app/controllers/admin"
	"idist-core/app/middlewares"

	"github.com/gin-gonic/gin"
)

func CommonRoutes(router *gin.RouterGroup) {
	router.GET("/provinces", middlewares.Gate("", ""), admin.ListProvinces)
	router.GET("/provinces/id", middlewares.Gate("", ""), admin.ReadProvince)
	router.PUT("/provinces/id", middlewares.Gate("", ""), admin.UpdateProvince)

	// router.POST("/tuyen-sinh", middlewares.Gate("", ""), admin.CreateTuyenSinh)

	//Opensearch
	router.GET("/opensearch", middlewares.Gate("", ""), admin.ListIndexes)
	router.POST("/insertOpensearch", middlewares.Gate("", ""), admin.InsertData)
	router.DELETE("/deleteOpensearch", middlewares.Gate("", ""), admin.DeleteIndex)

	//Link: Url Start
	// router.GET("/links", middlewares.Gate("", ""), admin.ListLink)
	// router.POST("/links", middlewares.Gate("", ""), admin.CreateLink)
	// router.GET("/links/:id", middlewares.Gate("", ""), admin.GetLink)
	// router.DELETE("/links/:id", middlewares.Gate("", ""), admin.DeleteLink)

	//Mention
	router.GET("/mentions", middlewares.Gate("", ""), admin.ListMention)
	router.GET("/mentions/addLabel", middlewares.Gate("", ""), admin.ListMentionNoLabel)
	router.GET("/mentions/addLabelWeb", middlewares.Gate("", ""), admin.ListMentionWebNoLabel)
	router.GET("/mentions/addLabelFace", middlewares.Gate("", ""), admin.ListMentionFaceNoLabel)
	router.GET("/mentions/:id", middlewares.Gate("", ""), admin.GetNew)
	router.GET("/mentions/count", middlewares.Gate("", ""), admin.CountMention)
	router.GET("/mentions/countLabelId", middlewares.Gate("", ""), admin.ListMentionAddLabel)
	router.GET("/mentions/:id/sentences", middlewares.Gate("", ""), admin.GetSentencesFromNew)
	router.GET("/mentions/countLabel", middlewares.Gate("", ""), admin.CountLabelMention)
	router.DELETE("/mentions/:id", middlewares.Gate("", ""), admin.DeleteNew)
	router.PUT("/mentions/:id/assign-label", middlewares.Gate("", ""), admin.AssignLabel)
	router.PUT("/mentions/:id/assign-label-sentences", middlewares.Gate("", ""), admin.AssignLabelSentences)

	//Source
	router.GET("/sources", middlewares.Gate("", ""), admin.ListSource)
	router.GET("/sources/type", middlewares.Gate("", ""), admin.ListTypeSource)
	router.GET("/sources/:id", middlewares.Gate("", ""), admin.GetSource)
	router.POST("/sources", middlewares.Gate("", ""), admin.CreateSource)
	router.PUT("/sources/:id", middlewares.Gate("", ""), admin.UpdateSource)
	router.DELETE("/sources/:id", middlewares.Gate("", ""), admin.DeleteSource)

	//LinkNext (Form Html)
	router.GET("/linkNexts", middlewares.Gate("", ""), admin.ListLinkNext)
	router.POST("/linkNexts", middlewares.Gate("", ""), admin.CreateLinkNext)
	router.GET("/linkNexts/:id", middlewares.Gate("", ""), admin.GetLinkNext)
	router.DELETE("/linkNexts/:id", middlewares.Gate("", ""), admin.DeleteLinkNext)
	router.PUT("/linkNexts/:id", middlewares.Gate("", ""), admin.UpdateLinkNext)

	// Keywords
	router.GET("/keywords", middlewares.Gate("", ""), admin.ListKeyword)
	router.POST("/keywords", middlewares.Gate("", ""), admin.CreateKeyword)
	router.GET("/keywords/:id", middlewares.Gate("", ""), admin.GetKeyword)
	router.PUT("/keywords/:id", middlewares.Gate("", ""), admin.UpdateKeyword)
	router.DELETE("/keywords/:id", middlewares.Gate("", ""), admin.DeleteKeyword)

	//Category
	router.GET("/categories", middlewares.Gate("", ""), admin.ListCategory)
	router.POST("/categories", middlewares.Gate("", ""), admin.CreateCategory)
	router.GET("/categories/:id", middlewares.Gate("", ""), admin.GetCategory)
	router.PUT("/categories/:id", middlewares.Gate("", ""), admin.UpdateCategory)
	router.DELETE("/categories/:id", middlewares.Gate("", ""), admin.DeleteCategory)

	//Label
	router.GET("/labels", middlewares.Gate("", ""), admin.ListLabel)
	router.POST("/labels", middlewares.Gate("", ""), admin.CreateLabel)
	router.GET("/labels/:id", middlewares.Gate("", ""), admin.GetLabel)
	router.PUT("/labels/:id", middlewares.Gate("", ""), admin.UpdateLabel)
	router.DELETE("/labels/:id", middlewares.Gate("", ""), admin.DeleteLabel)
}
