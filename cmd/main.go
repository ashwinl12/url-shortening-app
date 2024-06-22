package main

import (
	"fmt"
	"go-url-shortner/config"
	internalPkg "go-url-shortner/internal"
	"go-url-shortner/internal/server/url"
	"go-url-shortner/internal/server/user"
	"go-url-shortner/pkg/logger"
	log "go-url-shortner/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	// "github.com/golang-migrate/migrate/v4"
	// "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	fmt.Println("Hello from main")

	// Parse flags received through CLI
	newCnfg, err := config.ParseFlags()
	if err != nil {
		logger.LogError(err, "error inside main.go: ParseFlags", nil)
		return
	}

	// Make a DB connection to MongoDB
	client, ctx, cancel, err := config.ConnectToDB(log.Logger)
	if err != nil {
		log.LogFatal(err, "error inside main.go: error while connecting to DB", nil)
		return
	}
	fmt.Println("Db Name: ", *(newCnfg.DB.DbName))
	db := client.Database(*newCnfg.DB.DbName)

	defer cancel()
	defer client.Disconnect(ctx)

	// Migration code
	// driver, err := mongodb.WithInstance(client, &mongodb.Config{DatabaseName: "tiny_url_db"})
	// if err != nil {
	// 	logger.LogFatal(err, "error inside main.go: Failed to create MongoDB driver instance", nil)
	// 	return
	// }

	// // Create migrate instance
	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file://scripts/migrations", // Make sure this path is correct
	// 	"tiny_url_db",
	// 	driver,
	// )
	// if err != nil {
	// 	logger.LogFatal(err, "error inside main.go: failed to create migrate instance", nil)
	// 	return
	// }

	// // Run migrations
	// log.Logger.Info("Starting migrations...")
	// err = m.Up()
	// if err != nil {
	// 	if err == migrate.ErrNoChange {
	// 		log.Logger.Info("No migrations to apply")
	// 	} else {
	// 		logger.LogFatal(err, "error inside main.go: failed to run migrations", nil)
	// 		return
	// 	}
	// } else {
	// 	log.Logger.Info("Migrations applied successfully!")
	// }

	// Initialize Gin server
	server := gin.Default()

	// Initialize handlers
	InitializeHandlers(log.Logger, db)

	// Load routes
	internalPkg.LoadRoutes(server)

	server.Run()
}

func InitializeHandlers(logger *logrus.Logger, db *mongo.Database) {
	internalPkg.Usrh = user.NewUserSrvPkg(logger, db)
	internalPkg.Urlh = url.NewUrlSrvPkg(logger, db)
}
