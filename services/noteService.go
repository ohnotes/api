package services

import (
    db "github.com/ohnotes/api/database"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
)

type NoteResponse struct {
    ID string `json:"id"`
    Text string `json:"text"`
    Name string `json:"name"`
    Observation string `json:"observation"`
    Private bool `json:"private"`
    IsOwner bool `json:"isOwner"`
    Owner string `json:"owner"`
    Password string `json:"password"`
    Destructive bool `json:"destructive"`
    Times int `json:"times"`
}

func NoteService(c *gin.Context) {
    var response NoteResponse

    id := c.Param("id")
    token := c.MustGet("token")
    pass, _ := c.GetQuery("pass")
    
    err := db.Notes.FindOne(db.Ctx, bson.M{"id": id}).Decode(&response)
    if err != nil {
        c.JSON(404, notFound)

        return
    }

    if response.Owner != token {
        if response.Private {
            err := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(pass))
            if err != nil {
                c.JSON(403, forbidden)

                return
            }
        }

    } else {
        response.IsOwner = true

    }

    c.JSON(200, gin.H {
        "id": response.ID,
        "name": response.Name,
        "text": response.Text,
        "observation": response.Observation,
        "private": response.Private,
        "isOwner": response.IsOwner,
        "destructive": response.Destructive,
    });
}
