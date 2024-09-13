package api

import (
	"context"
	"familyFormUi/config"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	tele "gopkg.in/telebot.v3"
)

func InitRoutes() {
	// Initialize Gin router
	router := gin.Default()
	runbot()
	router.LoadHTMLGlob("static/*")
	router.Static("/static", "./static/")
	//router.LoadHTMLFiles("./static/index.html")

	// Serve the index.html file at the root ("/")
	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"books": "books",
		})
	})

	// Define the route to handle habit form submissions
	router.POST("/save-habit", func(c *gin.Context) {
		var h Habit
		if err := c.ShouldBindJSON(&h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Save the h data in Redis
		err := h.Create()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Habit saved successfully!"})
	})

	// Start the server
	log.Println("Server starting on :8080")
	router.Run(":8080")
}

func (h *Habit) Create() (err error) {
	key := fmt.Sprintf("habitMember:%d", 12)
	err = config.Rdb.HSet(context.Background(), key, h).Err()
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("%v has joined the club, he is planning to do %v within %v, he will say xx", h.Name, h.HabitName, h.CommitmentPeriod)
	botID, _ := strconv.Atoi(os.Getenv("TestingBotID"))
	config.B.Send(tele.ChatID(botID), msg)
	
	return
}

func runbot() {
	config.B.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send("Hello world!")
	})
	config.B.Start()
}
