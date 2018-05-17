package handlers

import (
	"net/http"

	"github.com/drxos/blog-api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const coll = models.CollectionArticle

// Create Article
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	article := models.Article{}
	err := c.Bind(&article) // Bind incoming json article to previous article variable
	if err != nil {
		c.Error(err)
		return
	}
	err = db.C(coll).Insert(article)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": err,
		})
	}
	err = db.C(coll).Find(nil).Sort("-_id").One(&article)
	c.JSON(http.StatusCreated, article)
}

// Read an article
func Read(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	article := models.Article{}
	oID := bson.ObjectIdHex(c.Param("id"))
	err := db.C(coll).FindId(oID).One(&article)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, article)
}

// Update an article
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	article := models.Article{}
	err := c.Bind(&article)
	if err != nil {
		c.Error(err)
		return
	}

	oID := bson.ObjectIdHex(c.Param("id"))
	query := bson.M{"_id": oID}

	doc := bson.M{
		"article": article.Title,
	}

	err = db.C(coll).Update(query, doc)
	err = db.C(coll).FindId(oID).One(&article)

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, article)
}

// Delete an article
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"_id": bson.ObjectIdHex(c.Param("id"))}
	err := db.C(coll).Remove(query)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, "Deleted")
}

// List an article
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	articles := []models.Article{}
	err := db.C(coll).Find(nil).Sort("-_id").All(&articles)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}
