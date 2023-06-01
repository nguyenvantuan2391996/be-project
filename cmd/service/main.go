package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nguyenvantuan2391996/be-project/config"
	_ "github.com/nguyenvantuan2391996/be-project/docs"
	"github.com/nguyenvantuan2391996/be-project/handler"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/usecase"
	"github.com/nguyenvantuan2391996/be-project/internal/infrastructure/repository"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	cfg, err := config.LoadConfig("build")
	if err != nil {
		log.Fatal("cannot load config:", err)
		return
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(cfg.DBSource), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("failed to open database:", err)
		return
	}

	// validator
	validate := validator.New()

	userRepo := repository.NewUserRepository(db)
	standardRepo := repository.NewStandardRepository(db)
	scoreRatingRepo := repository.NewScoreRatingRepository(db)

	userDomain := usecase.NewUserDomain(userRepo)
	standardDomain := usecase.NewStandardDomain(standardRepo)
	scoreRatingDomain := usecase.NewScoreRatingDomain(scoreRatingRepo)
	consultDomain := usecase.NewConsultDomain(userRepo, standardRepo, scoreRatingRepo)

	initRoutes(cfg, userDomain, standardDomain, scoreRatingDomain, consultDomain, validate)
}

func initRoutes(cfg *config.Config, userDomain *usecase.UserDomain, standardDomain *usecase.StandardDomain, scoreRatingDomain *usecase.ScoreRatingDomain, consultDomain *usecase.ConsultDomain, validate *validator.Validate) {
	// init handler
	h := handler.NewHandler(userDomain, standardDomain, scoreRatingDomain, consultDomain, validate)

	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := r.Group("v1/api")
	{
		// API users
		api.POST("/users", h.CreateUser)

		// API standards
		api.POST("/standards", h.CreateStandard)
		api.GET("/standards/:user_id", h.GetStandards)
		api.DELETE("/standards/:id", h.DeleteStandard)

		// API score_ratings
		api.POST("/score-ratings", h.BulkCreateScoreRating)
		api.GET("/score-ratings/:user_id", h.GetScoreRatings)
		api.PUT("/score-ratings", h.UpdateScoreRating)
		api.DELETE("/score-ratings/:id", h.DeleteScoreRating)

		// API Consult
		api.POST("/consult/:user_id", h.Consult)
	}

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":" + cfg.Port)
	if err != nil {
		return
	}
}
