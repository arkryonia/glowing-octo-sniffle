// Le paque middlewares est composé des middlewares gin
// Usage : router.USe(middlewares.Connect)

package middlewares

import (
	"net/http"

	"github.com/drxos/blog-api/db"
	"github.com/gin-gonic/gin"
)

// Connect est un middleware qui clone la session de la base de données pour chaque requête
// En rendant l'objet `db` disponible pour chaque gestionnaire
func Connect(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}

// ErrorHandler est un middleware qui gère les erreurs survues au cours d'une requête
func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
}
