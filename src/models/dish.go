package models

import (
	"context"
	"log"
)

type Dish struct {
	BaseModel
	Name        string
	Description string
}

func (d *Dish) Save(ctx context.Context) {
	db := GetDB()
	err := db.Create(d).Error
	if err != nil {
		log.Println("create user failed.", "err:", err.Error())
	}
}

func SaveUserInBatch(ctx context.Context, dishes *[]Dish) error {
	db := GetDB()
	err := db.Create(dishes).Error
	if err != nil {
		log.Println("create user failed.", "err:", err.Error())
	}
	return nil
}

func GetAllDishes(ctx context.Context) (dishes []Dish, err error) {
	db := GetDB()
	err = db.Find(&dishes).Error
	if err != nil {
		log.Println("get all error:", err)
		return nil, err
	}
	return dishes, nil
}

func GetAllDishesByNames(ctx context.Context, names []string) (dishes []Dish, err error) {
	db := GetDB()
	err = db.Where("name IN ?", names).Find(&dishes).Error
	if err != nil {
		log.Println("get all by names error:", err)
		return nil, err
	}
	return dishes, nil
}

func RandTakeDish(ctx context.Context, limit int) (dishes []Dish, err error) {
	db := GetDB()
	err = db.Limit(limit).Order("RAND()").Find(&dishes).Error
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
