package controllers

import (
	"encoding/json"
	"go-boiler-plate/initializers"
	"go-boiler-plate/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func PostsCreate(c *gin.Context) {

	//Get Records
	var b struct {
		Title string
		Body  string
	}
	c.Bind(&b)

	//Create Records
	post := models.Post{Title: b.Title, Body: b.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	const cacheKey = "post_index"

	val, err := initializers.RDB.Get(initializers.CTX, cacheKey).Result()

	if err == redis.Nil {
		//put the redis cache code here
		//Fetch data from db
		var posts []models.Post
		initializers.DB.Find(&posts)

		//cache fetched data
		postJSON, _ := json.Marshal(posts)
		initializers.RDB.Set(initializers.CTX, cacheKey, postJSON, 10*time.Minute)
		c.JSON(http.StatusOK, gin.H{
			"posts": posts,
		})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to find the records",
		})
	} else {
		var posts []models.Post
		json.Unmarshal([]byte(val), &posts)
		c.JSON(200, gin.H{
			"post": posts,
		})
	}
}

func PostsShow(c *gin.Context) {
	//Get Query string
	id := c.Param("id")

	val, err := initializers.RDB.Get(initializers.CTX, id).Result()

	if err == redis.Nil {
		//cache miss, fetch from the db
		var post models.Post
		initializers.DB.Find(&post, id)

		//Cache the result
		postJSON, _ := json.Marshal(post)
		initializers.RDB.Set(initializers.CTX, id, postJSON, 10*time.Minute)
		c.JSON(200, gin.H{
			"post": post,
		})
	} else if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to get post from Redis"},
		)
	} else {
		//Cache hit
		var post models.Post
		json.Unmarshal([]byte(val), &post)
		c.JSON(http.StatusOK, gin.H{
			"post": post,
		})
	}

}

func PostsUpdate(c *gin.Context) {
	//Get Query string
	id := c.Param("id")

	//Get the updated data
	var b struct {
		Title string
		Body  string
	}
	c.Bind(&b)
	//Find the post were updating
	var posts models.Post
	initializers.DB.Find(&posts, id)

	//Update it
	initializers.DB.Model(&posts).Updates(models.Post{Title: b.Title, Body: b.Body})

	//response
	c.JSON(200, gin.H{
		"post": posts,
	})

}

func PostsDelete(c *gin.Context) {
	//Get Query string
	id := c.Param("id")

	//Action
	initializers.DB.Delete(&models.Post{}, id)

	//response
	c.JSON(200, gin.H{
		"msg": "Post has been deleted",
	})

}
