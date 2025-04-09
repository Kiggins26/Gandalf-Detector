package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

    utils "./gandalf-detector"

    "github.com/gin-gonic/gin"
	"github.com/bwmarrin/discordgo"
    "github.com/joho/godotenv"
)

var discordSession *discordgo.Session

// We'll store message & user IDs for matching
var trackedMessages = make(map[string]string)

func main() {

    err := godotenv.Load(".env")
    if err != nil{
        log.Panic("Error loading .env file: ", err)
    }
    BotToken := os.Getenv("DiscordToken")

	// Initialize Discord bot
	var err error
	discordSession, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v", err)
	}

	discordSession.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsDirectMessages |
		discordgo.IntentsGuildMessageReactions |
		discordgo.IntentsMessageContent

	discordSession.AddHandler(reactionAdd)

	err = discordSession.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %v", err)
	}
	fmt.Println("🤖 Discord bot is running...")

	// Start Gin server in a goroutine
	go startAPI()

	// Handle shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	discordSession.Close()
}

// Gin API starts here
func startAPI() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		var req struct {
			wallet_address string `json:"wallet_address"` // Discord user ID
		}

		if err := c.BindJSON(&req); err != nil || req.wallet_address == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		}

		go func(wallet_address, text string) {
            userId := utils.GetDiscordNameForWallet("rpcClient", wallet_address)
			channel, err := discordSession.UserChannelCreate(userId)
			if err != nil {
				log.Panic("Failed to create DM channel: ", err)
			}

			message = "Please react with 👍 to confirm your recent transaction"

			msg, err := discordSession.ChannelMessageSend(channel.ID, message)
			if err != nil {
				log.Panic("Failed to send DM: ", err)
			}

			trackedMessages[msg.ID] = userID

			// Add thumbs up reaction
			_ = discordSession.MessageReactionAdd(channel.ID, msg.ID, "👍")
		}(req.UserID, req.Text)

		c.JSON(http.StatusOK, gin.H{"status": "DM sent"})
	})

	log.Println("🚀 Gin API is running on http://localhost:8080")
	r.Run(":8080")
}

// Reaction event
func reactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	userID, ok := trackedMessages[r.MessageID]
	if ok && r.UserID == userID && r.Emoji.Name == "👍" {
		channel, err := s.UserChannelCreate(r.UserID)
		if err != nil {
			log.Printf("Error creating DM channel: %v", err)
			return
		}

		_, _ = s.ChannelMessageSend(channel.ID, "✅ Confirmed! Taking action now...")


		// Clean up
		delete(trackedMessages, r.MessageID)
	}
}

