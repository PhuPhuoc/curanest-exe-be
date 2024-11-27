package api

import (
	"fmt"
	"strings"

	appointmenthandler "github.com/PhuPhuoc/curanest_exe_be/controller/appointment_services/handler"
	customerhandler "github.com/PhuPhuoc/curanest_exe_be/controller/customer_services/handler"
	nursehandler "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/handler"
	patienthandler "github.com/PhuPhuoc/curanest_exe_be/controller/patient_services/handler"
	paymenthandler "github.com/PhuPhuoc/curanest_exe_be/controller/payment_services/handler"
	reviewhandler "github.com/PhuPhuoc/curanest_exe_be/controller/review_services/handler"
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

// func (sv *server) RunApp() error {
// 	docs.SwaggerInfo.BasePath = "/api/v1"
// 	router := gin.New()
// 	config := cors.DefaultConfig()
// 	config.AllowAllOrigins = true
// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
// 	router.Use(
// 		cors.New(config),
// 		gin.LoggerWithWriter(gin.DefaultWriter, "/swagger/*any"),
// 		gin.Recovery(),
// 	)

// 	v1 := router.Group("/api/v1")
// 	sv.registerRoutes(v1)
// 	sv.runLog()
// 	return router.Run(sv.address)
// }

func (sv *server) RunApp() error {
	docs.SwaggerInfo.BasePath = "/api/v1"
	router := gin.New()

	// Cấu hình CORS chi tiết và linh hoạt
	config := cors.DefaultConfig()
	config.AllowOriginFunc = func(origin string) bool {
		// Danh sách các origin được phép
		allowedOrigins := []string{
			// Web localhost
			"http://localhost:3000",

			// Domain production frontend
			"https://curanest.com.vn",
			"https://www.curanest.com.vn",

			// Flutter mobile development origins
			"http://localhost",
			"capacitor://localhost",
			"ionic://localhost",
			"http://127.0.0.1",

			// Các domain mobile app (nếu có)
			// "https://your-mobile-app-domain.com",
		}

		// Kiểm tra origin
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				return true
			}
		}

		// Cho phép subdomain của domain chính
		return strings.Contains(origin, "curanest.com.cn")
	}

	// Cấu hình headers và methods
	config.AllowMethods = []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"OPTIONS",
		"PATCH",
	}
	config.AllowHeaders = []string{
		"Origin",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"accept",
		"origin",
		"Cache-Control",
		"X-Requested-With",
		"Access-Control-Allow-Origin",
	}
	config.AllowCredentials = true

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Sử dụng middleware CORS
	router.Use(
		cors.New(config),
		gin.LoggerWithWriter(gin.DefaultWriter, "/swagger/*any"),
		gin.Recovery(),
	)

	// Thêm middleware xử lý OPTIONS request nếu cần
	router.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
			c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Đăng ký routes
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
	paymenthandler.RegisterPaymentRoutes(v1, sv.db)
	reviewhandler.RegisterReviewRoutes(v1, sv.db)
}
