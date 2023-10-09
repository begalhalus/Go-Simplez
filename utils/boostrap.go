package boostrap

import (
	"fmt"
	rolesHandler "go-simple/roles/handler"
	rolesRepo "go-simple/roles/repo"
	rolesUsecase "go-simple/roles/usecase"
	usersHandler "go-simple/users/handler"
	usersRepo "go-simple/users/repo"
	usersUsecase "go-simple/users/usecase"
	apps_config "go-simple/utils/config"
	"go-simple/utils/database"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BoostrapApp() {

	// load config env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error load .env file")
	}

	// init config apps
	apps_config.InitAppConfig()
	apps_config.InitDatabaseConfig()

	// database connection
	db := database.ConnectDatabase()

	// init config logging
	apps_config.DefaultLogging()

	// init routes
	router := gin.Default()

	// config cors
	router.Use(apps_config.InitCors())

	router.Static(apps_config.STATIC_ROUTE, apps_config.STATIC_DIR)

	// repository
	rolesRepo := rolesRepo.CreateRolesRepo(db)
	usersRepo := usersRepo.CreateUsersRepo(db)

	// usecase
	rolesUsecase := rolesUsecase.CreateRolesUsecase(rolesRepo)
	usersUsecase := usersUsecase.CreateUsersUsecase(usersRepo)

	// handler
	rolesHandler.CreateRolesHandler(router, rolesUsecase)
	usersHandler.CreateUsersHandler(router, usersUsecase)

	fmt.Printf("[%s] Restapps running on port: %s\n", time.Now().Format("2006-01-02 15:04:05"), apps_config.PORT)

	router.Run(":" + apps_config.PORT)
}
