package config

import(
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv(){

}

func ConnectDB(){
	dsn := os.Getenv("DATABASE_URL")
	db,err:=gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("DB FAILED")
	}

	DB = db
	print("CONNECTION SUCCESSFUL")
}