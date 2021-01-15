package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"gitlab.com/bunlert274/go-service/src/delivery"
	"gitlab.com/bunlert274/go-service/src/middleware"
	"gitlab.com/bunlert274/go-service/src/migrate"
	"gitlab.com/bunlert274/go-service/src/repository"
	"gitlab.com/bunlert274/go-service/src/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("charset", "utf8")
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Bangkok")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	e := echo.New()
	mw := middleware.Initial()

	e.Pre(mw.RemoveTrailingSlash())
	e.Use(mw.CORS)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)
	sqlConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal(err)
	}
	migrate.Migrate(sqlConnection)
	defer func() {
		sql, err := sqlConnection.DB()
		if err != nil {
			e.Logger.Panic(err.Error())
		}
		sql.Close()
		e.Close()
	}()

	timeout := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := repository.NewUserRepository(sqlConnection)
	userUseCase := usecase.NewUserUseCase(userRepository, timeout)
	delivery.NewUserHandle(e, userUseCase)

	e.GET("/health-check", func(c echo.Context) error {
		// should be return from check in service (3th party, connection)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "available",
		})
	})

	log.Fatal(e.Start(viper.GetString(`server.address`)))
}
