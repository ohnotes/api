package services

import (
    "fmt"

    db "github.com/ohnotes/api/database"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
)

type Note struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Observation string `json:"observation,omitempty"`
    Private bool `json:"private"`
    Text string `json:"text"`
    CreatedAt string `json:"createdAt"`
    Owner string `json:"owner"`
    Password string `json:"password"`
    Destructive bool `json:"destructive"`
    Turns int `json:"turns"`
}

func NewService(c *gin.Context) {
    var create Note

    token := c.MustGet("token")
    id := uuid.New().String()

    if err := c.BindJSON(&create); err != nil {
        c.JSON(400, gin.H {
            "message": "Failed to create a new note.",
        })

        return
    }

    var password []byte

    if create.Private {
        password, _ = bcrypt.GenerateFromPassword([]byte(create.Password), 12)

    }

    _, err := db.Notes.InsertOne(db.Ctx, Note {
        ID: id,
        Name: create.Name,
        Observation: create.Observation,
        Private: create.Private,
        Text: create.Text,
        CreatedAt: create.CreatedAt,
        Owner: fmt.Sprintf("%v", token),
        Password: string(password),
        Destructive: create.Destructive,
        Turns: create.Turns,
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
