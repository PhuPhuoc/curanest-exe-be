package userhandler

import (
	"net/http"

	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	eg := rg.Group("/users")
	{
		eg.GET("/hello", sayHello)
	}
}

// @BasePath		/api/v1
// @Summary		hello
// @Description	hello
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}	string	"say hi"
// @Router			/users/hello [get]
func sayHello(c *gin.Context) {
	utils.SendSuccess(c, http.StatusOK, "This is greeting from Phu Phuoc linux server", nil, nil)
}
