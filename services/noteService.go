package services

import (
    db "github.com/ohnotes/api/database"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
)

func NoteService(c *gin.Context) {
    var response Note

    id := c.Param("id")
    pass, _ := c.GetQuery("pass")
    
    err := db.Notes.FindOne(db.Ctx, bson.M{"id": id}).Decode(&response)
    if err != nil {
        c.JSON(404, notFound)

        return
    }

    if response.Private {
        err := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(pass))
        if err != nil {
            c.JSON(403, forbidden)

            return
        }
    }

    c.JSON(200, gin.H {
        "id": response.ID,
        "name": response.Name,
        "text": response.Text,
        "observation": response.Observation,
        "private": response.Private,
        "destructive": response.Destructive,
        "shared": response.Shared,
    });
}
