package handler

import (
	"database/sql"

	"bimbo/internal/config"
	"bimbo/internal/repository"
	"bimbo/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetEndpoints(cfg *config.Config, db *sql.DB, logger *logrus.Logger, group *gin.RouterGroup) {
	repos := repository.RepositoryInit(db)
	// struct inside - all repos
	services := service.ServiceInit(repos, cfg) // all service init

	handler := NewHandler(services, logger, cfg)

	auth := group.Group("/auth")
	{
		auth.POST("/signup", handler.SignUp)
		auth.POST("/signin", handler.SignIn)
	}

	admin := group.Group("/admin")

	admin.POST("/company", handler.CreateCompany)
	admin.GET("/company", handler.GetListCompany)

	admin.POST("/departament", handler.CreateDepartament)
	admin.GET("/departament", handler.GetListDepartament)

	admin.POST("/position", handler.CreatePosition)
	admin.GET("/position", handler.GetListPosition)

	admin.POST("/role", handler.CreateRole)
	admin.GET("/role", handler.GetListRole)

	// refresh := group.Group("/auth/refresh", middleware.AuthorizeJWT("refreshx")) // todo: env config
	{
		// refresh.POST("", handler.RefreshJwt)
	}

	// logout := group.Group("/auth/logout", middleware.AuthorizeJWT("accessx"))
	{
		// logout.POST("", handler.Logout)
	}
}
