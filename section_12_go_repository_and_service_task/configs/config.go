package configs

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConfig() (*gorm.DB, error) {
	var (
		db_host     = os.Getenv("DB_HOST")
		db_username = os.Getenv("DB_USERNAME")
		db_password = os.Getenv("DB_PASSWORD")
		db_name     = os.Getenv("DB_NAME")
		db_port     = os.Getenv("DB_PORT")
	)

	dsn := db_username + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
