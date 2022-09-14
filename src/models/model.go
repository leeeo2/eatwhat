package models

import (
	"context"
	"fmt"
	"log"
	"lxxxxxxxx.github.com/eatwhat/src/common"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type BaseModel struct {
	Id        string    `gorm:"type:varchar(255);not null;primaryKey"`
	CreatedAt time.Time `gorm:"type:datetime(6);not null;index:created_at;default:current_timestamp(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6);not null;index:updated_at;default:current_timestamp(6) on update current_timestamp(6)"`
	//DeletedAt gorm.DeletedAt `gorm:"type:datetime(6);index:deleted_at"`
}

func Setup(ctx context.Context) {
	config := common.GlobalConfig().Db
	fmt.Println("db config:", config)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Schema)
	dababase, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("open db failed.")
	}
	dababase.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Dish{})
	db = dababase
}

func GetDB() *gorm.DB {
	return db
}
