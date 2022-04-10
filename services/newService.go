package services

import (
    "fmt"

    db "github.com/ohnotes/api/database"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
)

type CreateNote struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Observation string `json:"observation,omitempty"`
    Private bool `json:"private"`
    Text string `json:"text"`
    CreatedAt string `json:"createdAt"`
    Owner string `json:"owner"`
    Password string `json:"password"`
    Destructive bool `json:"destructive"`
    Times int `json:"times"`
}

func NewService(c *gin.Context) {
    var create CreateNote

    token := c.MustGet("token")
    id := uuid.New().String()

    if err := c.BindJSON(&create); err != nil {
        c.JSON(400, gin.H {
            "message": "Failed to create a new note.",
        })

        return
    }

    hashed, err := bcrypt.GenerateFromPassword([]byte(create.Password), 12)
    if err != nil {
        c.JSON(400, gin.H {
            "message": "Failed to create a new note.",
        })

        return
    }

    _, err = db.Notes.InsertOne(db.Ctx, CreateNote {
        ID: id,
        Name: create.Name,
        Observation: create.Observation,
        Private: create.Private,
        Text: create.Text,
        CreatedAt: create.CreatedAt,
        Owner: fmt.Sprintf("%v", token),
        Password: string(hashed),
        Destructive: create.Destructive,
        Times: create.Times,
    })
    if err != nil {
        c.JSON(400, gin.H {
            "message": "An error was occurred during insert.",
        })

        return
    }

    c.JSON(200, gin.H {
        "id": id,
    })
}
