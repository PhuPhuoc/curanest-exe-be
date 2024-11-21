package api

import (
	"fmt"

	appointmenthandler "github.com/PhuPhuoc/curanest_exe_be/controller/appointment_services/handler"
	customerhandler "github.com/PhuPhuoc/curanest_exe_be/controller/customer_services/handler"
	nursehandler "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/handler"
	patienthandler "github.com/PhuPhuoc/curanest_exe_be/controller/patient_services/handler"
	techniquehandler "github.com/PhuPhuoc/curanest_exe_be/controller/technique_services/handler"
	userhandler "github.com/PhuPhuoc/curanest_exe_be/controller/user_services/handler"
	docs "github.com/PhuPhuoc/curanest_exe_be/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server struct {
	address string
	db      *sqlx.DB
}

func InitServer(addr string, db *sqlx.DB) *server {
	return &server{
		address: addr,
		db:      db,
	}
}

func (sv *server) RunApp() error {
	docs.SwaggerInfo.BasePath = "/api/v1"
	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/swagger/*any"),
		gin.Recovery(),
		cors.New(config),
	)

	v1 := router.Group("/api/v1")
	sv.registerRoutes(v1)
	sv.runLog()
	return router.Run(sv.address)
}

func (sv *server) runLog() {
	fmt.Printf("\nFor development: http://localhost%s/swagger/index.html\n\n", sv.address)
}

func (sv *server) registerRoutes(v1 *gin.RouterGroup) {
	userhandler.RegisterUserRoutes(v1, sv.db)
	customerhandler.RegisterCustomerRoutes(v1, sv.db)
	techniquehandler.RegisterTechniqueRoutes(v1, sv.db)
	nursehandler.RegisterNurseRoutes(v1, sv.db)
	patienthandler.RegisterPatientRoutes(v1, sv.db)
	appointmenthandler.RegisterNurseRoutes(v1, sv.db)
}
