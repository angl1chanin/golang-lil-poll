package app

import (
	"backend/config"
	"backend/internal/api/http/v1/handlers"
	"backend/internal/api/http/v1/routes"
	"backend/internal/repository/option"
	"backend/internal/repository/poll"
	option2 "backend/internal/service/option"
	poll2 "backend/internal/service/poll"
	"backend/internal/usecase"
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	// open db connection
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		fmt.Errorf("%w", err)
	}

	// Init repositories
	pollRepo, err := poll.NewPollRepository(db)
	if err != nil {
		fmt.Errorf("%w", err)
	}

	optionRepo, err := option.NewOptionRepository(db)
	if err != nil {
		fmt.Errorf("%w", err)
	}

	// Init services
	pollService := poll2.NewPollService(pollRepo)
	optionService := option2.NewOptionService(optionRepo)

	// Init usecase
	useCase := usecase.NewUseCase(pollService, optionService)

	// Init handlers
	handler := handlers.NewHandler(useCase)

	// HTTP server
	router := gin.Default()

	// Настройка CORS middleware
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowOrigins = []string{"*"}                                                                                                               // Разрешенный домен
	corsCfg.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}                                                                // Разрешенные методы
	corsCfg.AllowHeaders = []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With", "ngrok-skip-browser-warning"} // Разрешенные заголовки

	// Добавление CORS middleware в роутер
	router.Use(cors.New(corsCfg))

	// setup routes
	routes.Setup(router, handler)

	// start server
	err = router.Run(cfg.Address)
	if err != nil {
		fmt.Println(err)
	}
}
