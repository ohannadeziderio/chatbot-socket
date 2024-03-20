package chat

import (
	"math/rand"
	"strings"
	"time"

	"github.com/ohannadeziderio/chatbot-socket/pkg/references"
)

type Chatbot struct {
	Name string
}

func NewChatbot() *Chatbot {
	return &Chatbot{Name: "bot"}
}

func (bot Chatbot) Answer(message string) string {
	message = strings.ToLower(message)

	switch {
	case strings.Contains(message, "movie"):
		return ProcessMessage(message, references.GeekMovies)
	case strings.Contains(message, "music"):
		return ProcessMessage(message, references.Music)
	case strings.Contains(message, "book"):
		return ProcessMessage(message, references.Books)
	case strings.Contains(message, "meme"):
		return RandMeme(references.Memes)
	default:
		return ProcessMessage(message, references.Greetings)
	}
}

func ProcessMessage(message string, ref map[string]string) string {
	for key, value := range ref {
		if strings.Contains(message, key) {
			return value
		}
	}

	return "I think you spoke in Glogulese, because I didn't understand anything."
}

func RandMeme(ref []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := r.Intn(len(ref))

	return ref[randomIndex]
}
