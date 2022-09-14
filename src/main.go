package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"lxxxxxxxx.github.com/eatwhat/src/common"
	"lxxxxxxxx.github.com/eatwhat/src/models"
	"net/http"
	"strconv"
	"strings"
)

var (
	configPath = flag.String("c", "./src/etc/benckend.yml", "config file path.")
)

type dishRes struct {
	Id   string
	Type string
	Desc string
}

func main() {
	flag.Parse()
	fmt.Println("config path:", *configPath)

	if err := common.InitConfig(*configPath); err != nil {
		fmt.Errorf("init config failed,error:%w", err)
	}
	ctx := context.Background()

	models.Setup(ctx)

	router := gin.Default()

	router.GET("/AddDish", func(c *gin.Context) {
		desc := c.Query("Description")
		t := c.Query("Type")
		//for i := 0; i < len(t); i++ {
		//	if t[i] != 'a' && t[i] != 'b' && t[i] != 'c' && t[i] != 'd' {
		//		fmt.Println("dish type not support:", t[i])
		//		continue
		//	}
		id, _ := uuid.NewUUID()
		dish := models.Dish{
			BaseModel:   models.BaseModel{Id: id.String()},
			Description: desc,
			Type:        string(t),
		}
		dish.Save(ctx)
		//}
		c.JSON(http.StatusOK, gin.H{"message": "add success"})
	})

	router.GET("/EatWhat", func(c *gin.Context) {
		count, err := strconv.Atoi(c.Query("Count"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "parameter `Count` must be a integer",
			})
			return
		}
		t := c.Query("Type")
		dishes, err := models.RandTakeDish(ctx, t, count)
		fmt.Println("dishes:", dishes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "query database failed",
			})
			return
		}

		res := make([]dishRes, 0)
		for _, dish := range dishes {
			t := ""
			if strings.Contains(dish.Type, "a") {
				t += "早"
			}
			if strings.Contains(dish.Type, "b") {
				t += "中"
			}
			if strings.Contains(dish.Type, "c") {
				t += "晚"
			}
			if strings.Contains(dish.Type, "d") {
				t += "宵夜"
			}
			res = append(res, dishRes{
				Id:   dish.Id,
				Desc: dish.Description,
				Type: t,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"Dishes": res,
		})
	})

	router.GET("/DescribeDishes", func(c *gin.Context) {
		dishes, err := models.GetAllDishes(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "query database failed",
			})
			return
		}

		res := make([]dishRes, 0)
		for _, dish := range dishes {
			t := ""
			if strings.Contains(dish.Type, "a") {
				t += "早"
			}
			if strings.Contains(dish.Type, "b") {
				t += "中"
			}
			if strings.Contains(dish.Type, "c") {
				t += "晚"
			}
			if strings.Contains(dish.Type, "d") {
				t += "宵夜"
			}
			res = append(res, dishRes{
				Id:   dish.Id,
				Desc: dish.Description,
				Type: t,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"Dishes": res,
		})
	})

	router.GET("/DeleteDish", func(c *gin.Context) {
		id := c.Query("Id")
		err := models.DeleteDishById(ctx, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "delete from database failed:" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "delete success",
		})
	})

	// 静态资源加载，本例为css,js以及资源图片
	router.StaticFS("/eatwhat", http.Dir("website/static"))

	// Listen and serve on 0.0.0.0:80
	router.Run(common.GlobalConfig().Server.ListenAddr)
}
