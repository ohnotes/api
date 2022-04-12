package services

import (
    db "github.com/ohnotes/api/database"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
)

type UpdateSettingsResponse struct {
    Name string `json:"name" bson:"name"`
    Observation string `json:"observation" bson:"observation"`
    Private bool `json:"private" bson:"private"`
    Password string `json:"password" bson:"password"`
}

func UpdateSettingsService(c *gin.Context) {
    var response UpdateSettingsResponse
    var note Note

    token := c.MustGet("token")
    
    err := c.BindJSON(&response)
    if err != nil {
        c.JSON(400, gin.H {
            "message": "An error was occurred during settings update.",
        })

        return
    }

    if err = db.Notes.FindOne(db.Ctx, bson.M{"id": c.Param("id")}).Decode(&note); err != nil {
        c.JSON(404, notFound)

        return
    }

    if token != note.Owner {
        c.JSON(403, forbidden)

        return
    }

    hashed, err := bcrypt.GenerateFromPassword([]byte(response.Password), 12)
    if err != nil {
        c.JSON(400, gin.H {
            "message": "Failed to create a valid hash.",
        })

        return
    }

    filter := bson.M{"id": note.ID}
    db.Notes.UpdateOne(db.Ctx, filter, bson.M{"$set": bson.M{"name": response.Name}})
    db.Notes.UpdateOne(db.Ctx, filter, bson.M{"$set": bson.M{"observation": response.Observation}})
    db.Notes.UpdateOne(db.Ctx, filter, bson.M{"$set": bson.M{"private": response.Private}})
    db.Notes.UpdateOne(db.Ctx, filter, bson.M{"$set": bson.M{"password": string(hashed)}})

    c.JSON(200, gin.H {
        "message": "Succesfully updated.",
    })
}
