package routes

import (
	"github.com/anggunpermata/patreon-clone/internal/controller"
	"github.com/labstack/echo/v4"
)

func NewRoutes(e *echo.Echo, handler *controller.BackendHandler) {
	e.GET("/healthcheck", handler.HealthcheckHandler)
	e.GET("/upload", handler.RouteUploadFile)
	e.POST("/upload/images", handler.UploadFile)
	e.GET("/file/show/:file_name", handler.GetFile)
	e.GET("/signup", handler.UserSignup)
	e.POST("/signup", handler.UserSignup)
	e.GET("/login", handler.UserLogin)
	e.POST("/login", handler.UserLogin)
	e.POST("/logout", handler.UserLogout)
	e.GET("/logout", handler.UserLogout)
	e.GET("/posts/create", handler.CreateAPost)
	e.POST("/posts/create", handler.CreateAPost)
	e.GET("/users/:targeted_user_id/browse/posts", handler.GetAllPostsByUserID)
	e.GET("/users/:targeted_user_id", handler.UserProfiles)
	e.POST("/users/:targeted_user_id/follow", handler.FollowAUser)
	e.GET("/users/:targeted_user_id/browse/followers", handler.GetAllFollower)
	e.GET("/users/:targeted_user_id/browse/following", handler.GetAllFollowing)
	// dashboard
	e.GET("/dashboard", handler.ShowUserDashboard)
	e.GET("/*", handler.HomePage)
}
