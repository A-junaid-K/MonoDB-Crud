package controllers

import (
	helper "github.com/Crud-Mongodb/helpers"
	models "github.com/Crud-Mongodb/model"
	"github.com/gin-gonic/gin"
)

func AddMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.Bind(&movie); err != nil {
		c.JSON(500, gin.H{
			"error": "Binding error",
		})
		return
	}
	res,err := helper.InsertMovie(movie)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully inserted movie",
		"id":      res,
	})
}

func ViewAllMovies(c *gin.Context) {
	list, err := helper.ListAllMovies()

	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"movies": list,
	})
}

func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie
	if err := c.Bind(&movie); err != nil {
		c.JSON(500, gin.H{
			"error": "binding error",
		})
		return
	}
	count, err := helper.UpdateMovie(movie, id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully updated",
		"count":   count,
	})

}

func GetOneMovie(c *gin.Context) {
	id := c.Param("id")
	movie, err := helper.GetOneMovie(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"movie": movie,
	})
}

func DeleteOneMovie(c *gin.Context) {
	id := c.Param("id")
	err := helper.DeleteOneMovie(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully deleted " + id + " movie",
	})
}

func DeleteAllMovies(c *gin.Context) {
	err := helper.DeleteAllMovie()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully deleted all data",
	})
}
