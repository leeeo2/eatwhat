package models

import (
	"context"
	"log"
)

type Dish struct {
	BaseModel
	Description string
	Type        string
}

func (d *Dish) Save(ctx context.Context) {
	db := GetDB()
	err := db.Create(d).Error
	if err != nil {
		log.Println("create user failed.", "err:", err.Error())
	}
}

func GetAllDishes(ctx context.Context) (dishes []Dish, err error) {
	db := GetDB()
	err = db.Order("type").Find(&dishes).Error
	if err != nil {
		log.Println("get all error:", err)
		return nil, err
	}
	return dishes, nil
}

func RandTakeDish(ctx context.Context, t string, limit int) (dishes []Dish, err error) {
	db := GetDB()
	err = db.Limit(limit).Order("RAND()").Where("type like ?", "%"+t+"%").Find(&dishes).Error
	if err != nil {
		log.Println("randTake error:", err)
		return nil, err
	}
	return dishes, nil
}

func DeleteDishById(ctx context.Context, id string) error {
	db := GetDB()
	err := db.Where("Id = ?", id).Delete(&Dish{}).Error
	if err != nil {
		log.Printf("delete dish %s error:%v ", id, err)
		return err
	}
	return nil
}
