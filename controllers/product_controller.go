package controllers

import (
	"belajar-middleware/database"
	"belajar-middleware/helpers"
	"belajar-middleware/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"fmt"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)

	product := models.Product{}
	user := c.MustGet("user").(jwt.MapClaims)

	print(fmt.Sprintf("%v",user))
	userID := uint(user["id"].(float64))

	// error
	var err error

	if contentType != "application/json" {
		err = c.ShouldBindJSON(&product)
	} else {
		err = c.ShouldBind(&product)
	}

	if err != nil {
		panic(err)
	}

	product.UserID = userID

	err = db.Debug().Create(&product).Error

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to create",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success to create",
		"data":    product,
	})

}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)

	product := models.Product{}
	user := c.MustGet("user").(jwt.MapClaims)
	print(user);
	userID := uint(user["id"].(float64))

	// error
	var err error

	productID, err := strconv.Atoi(c.Param("productID"))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Parameter",
			"error":   err.Error(),
		})
	}

	if contentType != "application/json" {
		err = c.ShouldBindJSON(&product)
	} else {
		err = c.ShouldBind(&product)
	}

	if err != nil {
		panic(err)
	}

	product.ID = uint(productID)
	product.UserID = userID

	err = db.Debug().Where("id = ?", productID).Updates(models.Product{
		Title:       product.Title,
		Description: product.Description,
	}).Error

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to update",
			"error":   err.Error(),
		})
		return
	}

	db.Preload("User").First(&product)

	c.JSON(200, gin.H{
		"message": "Success to update",
		"data":    product,
	})

}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()


	product := models.Product{}
	user := c.MustGet("user").(jwt.MapClaims)
	print(user);

	// error
	var err error

	productID, err := strconv.Atoi(c.Param("productID"))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Parameter",
			"error":   err.Error(),
		})
	}


	if err != nil {
		panic(err)
	}

	product.ID = uint(productID)

	err = db.Debug().Delete(&product).Error

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success to delete",
		"data":    product,
	})

}

func GetProduct(c *gin.Context) {
	db := database.GetDB()


	product := models.Product{}
	user := c.MustGet("user").(jwt.MapClaims)
	print(user);

	// error
	var err error

	productID, err := strconv.Atoi(c.Param("productID"))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Parameter",
			"error":   err.Error(),
		})
	}

	product.ID = uint(productID)

	res := db.Preload("User").First(&product)

	if res.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to find product",
			"error":   res.Error,
		})
		return
	}

	if res.RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "product not found",
			"error":   "",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success to find",
		"data": product,
	})

}

func GetProducts(c *gin.Context) {
	db := database.GetDB()
	product := []models.Product{}
	user := c.MustGet("user").(jwt.MapClaims)
	print(user);


	res := db.Debug().Preload("User").Find(&product)

	if res.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed to find product",
			"error":   res.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
		"data": product,
	})

}
