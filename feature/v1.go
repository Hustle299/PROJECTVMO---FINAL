package feature

import (
	authStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/storage"
	authTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/transport"
	authUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/usecase"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/middleware"
	"github.com/gin-gonic/gin"

	userStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
	userTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/transport"
	userUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"

	countryStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/country/storage"
	countryTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/country/transport"
	countryUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/country/usecase"

	departmentStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/department/storage"
	departmentTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/department/transport"
	departmentUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/department/usecase"

	roleStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/role/storage"
	roleTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/role/transport"
	roleUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/role/usecase"

	appliIdentityStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/storage"
	appliIdentityTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/transport"
	appliIdentityUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/usecase"

	"github.com/cesc1802/onboarding-and-volunteer-service/cmd/config"
	"github.com/cesc1802/onboarding-and-volunteer-service/database"
)

// @host localhost:8080
// @BasePath /api/v1
func NewRouter() *gin.Engine {

	router := gin.Default()
	secretKey := authStorage.GetSecretKey()
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	db, err := database.NewDatabase(cfg)
	if err != nil {
		panic(err)
	}

	v1 := router.Group("/api/v1")
	// Initialize repository
	authRepo := authStorage.NewAuthenticationRepository(db)
	userRepo := userStorage.NewAdminRepository(db)
	applicantRepo := userStorage.NewApplicantRepository(db)
	applicantRequestRepo := userStorage.NewApplicantRequestRepository(db)
	applicantIdentityRepo := appliIdentityStorage.NewUserIdentityRepository(db)
	countryRepo := countryStorage.NewCountryRepository(db)
	departmentRepo := departmentStorage.NewDepartmentRepository(db)
	roleRepo := roleStorage.NewRoleRepository(db)

	// Initialize usecase
	authUseCase := authUsecase.NewUserUsecase(authRepo, secretKey)
	userUseCase := userUsecase.NewAdminUsecase(userRepo)
	applicantUseCase := userUsecase.NewApplicantUsecase(applicantRepo)
	applicantRequestUseCase := userUsecase.NewApplicantRequestUsecase(applicantRequestRepo)
	applicantIdenityUseCase := appliIdentityUsecase.NewUserIdentityUsecase(applicantIdentityRepo)
	countryUseCase := countryUsecase.NewCountryUsecase(countryRepo)
	departmentUseCase := departmentUsecase.NewDepartmentUsecase(departmentRepo)
	roleUseCase := roleUsecase.NewRoleUsecase(roleRepo)

	// Initialize handler
	authHandler := authTransport.NewAuthenticationHandler(authUseCase)
	userHandler := userTransport.NewAuthenticationHandler(userUseCase)
	applicantHandler := userTransport.NewApplicantHandler(applicantUseCase)
	applicantRequestHandler := userTransport.NewApplicantRequestHandler(applicantRequestUseCase)
	applicantIdentityHandler := appliIdentityTransport.NewUserIdentityHandler(applicantIdenityUseCase)
	countryHandler := countryTransport.NewCountryHandler(countryUseCase)
	departmentHandler := departmentTransport.NewDepartmentHandler(departmentUseCase)
	roleHandler := roleTransport.NewRoleHandler(roleUseCase)

	auth := v1.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)

		auth.POST("/register", authHandler.Register)
	}

	country := v1.Group("/country")
	{
		country.POST("/", countryHandler.CreateCountry)
		country.GET("/:id", countryHandler.GetCountryByID)
		country.PUT("/:id", countryHandler.UpdateCountry)
		country.DELETE("/:id", countryHandler.DeleteCountry)
	}
	department := v1.Group("/department")
	{
		department.POST("/", departmentHandler.CreateDepartment)
		department.GET("/:id", departmentHandler.GetDepartmentByID)
		department.PUT("/:id", departmentHandler.UpdateDepartment)
		department.DELETE("/:id", departmentHandler.DeleteDepartment)
	}
	role := v1.Group("/role")
	{
		role.POST("/", roleHandler.CreateRole)
		role.GET("/:id", roleHandler.GetRoleByID)
		role.PUT("/:id", roleHandler.UpdateRole)
		role.DELETE("/:id", roleHandler.DeleteRole)
	}

	admin := v1.Group("/admin")
	admin.Use(middleware.AuthMiddleware(secretKey))
	{
		admin.GET("/list-request", userHandler.GetListRequest)
		admin.GET("/request/:id", userHandler.GetRequestById)
		admin.GET("/list-pending-request", userHandler.GetListPendingRequest)
		admin.GET("/pending-request/:id", userHandler.GetPendingRequestById)
		admin.POST("/approve-request/:id", userHandler.ApproveRequest)
		admin.POST("/reject-request/:id", userHandler.RejectRequest)
		admin.POST("/add-reject-notes/:id", userHandler.AddRejectNotes)
		admin.DELETE("/delete-request/:id", userHandler.DeleteRequest)
	}

	applicant := v1.Group("/user")
	{
		applicant.PUT("/:id/basic", applicantHandler.UpdateApplicantBasic)
		applicant.PUT("/:id", applicantHandler.UpdateApplicant)
		applicant.DELETE("/:id", applicantHandler.DeleteApplicant)
		applicant.GET("/:id", applicantHandler.FindApplicantByID)
	}

	appliRequest := v1.Group("/user-request")
	{
		appliRequest.POST("/", applicantRequestHandler.CreateApplicantRequest)
	}

	appliIdentity := v1.Group("user-identity")
	{
		appliIdentity.POST("/", applicantIdentityHandler.CreateUserIdentity)
		appliIdentity.GET("/:id", applicantIdentityHandler.FindUserIdentity)
		appliIdentity.PUT("/:id", applicantIdentityHandler.UpdateUserIdentity)
	}
	return router
}
