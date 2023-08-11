package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rodrigoherera/know-vegan-service/src/api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func init() {
	InitMySQL()
}

func InitMySQL() {
	db, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{
		Logger: getLogger(),
	})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if config.MODE == "test" {
		db.Debug()
	}

	DB = db
}

//Returns logger interface
func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color
		},
	)
}

//Returns data-source-name
func getDsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUSER,
		config.DBPASSWORD,
		config.DBHOST,
		config.DBPORT,
		config.DBNAME)
}
