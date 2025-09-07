package main

import (
	"fmt"
	"log"
	"time"

	"github.com/qyroai/qyro-go-sdk"
)

const (
	BASE_URL       = "http://localhost:3000"
	API_KEY_ID     = "68bd1f1fe5bcc3beb6628a2d"
	API_KEY_SECRET = "GipH9XMNTBJkaQGD9ExQGdixiXWgvgZgbCuu35igwBAttbUPuqCkHpwkrQdrKmQVzDZ1kF9xxpzkLibAXcsOJVjy6fHD0aXeAdyxbzSUJmWpetPY10o6fD21VsfWgZ69"
	ASSISTANT_ID   = "68bcf8d9ef381744eced70b4"
)

func main() {
	baseURL := BASE_URL
	apiKeyID := API_KEY_ID
	apiKeySecret := API_KEY_SECRET
	assistantID := ASSISTANT_ID

	// ---------------------------
	// Using Server Client
	// ---------------------------
	serverClient, err := qyro.NewQyroServerClient(baseURL, apiKeyID, apiKeySecret, 30*time.Second)
	if err != nil {
		log.Fatal("Error creating server client:", err)
	}

	fmt.Println("👉 Creating session via server API...")
	session, err := serverClient.CreateSession(assistantID, map[string]interface{}{
		"user_id": "user_abc",
		"context": "first test run",
	})
	if err != nil {
		log.Fatal("Error creating session:", err)
	}
	fmt.Println("✅ Session created:", session.ID)

	fmt.Println("👉 Sending message to assistant (server)...")
	messages, err := serverClient.Chat(assistantID, session.ID, "Hello from Go Server Client!")
	if err != nil {
		log.Fatal("Error sending chat:", err)
	}
	for _, m := range messages {
		fmt.Printf("[%s] %s\n", m.Role, m.Content)
	}

	// ---------------------------
	// Using Client SDK (token-based)
	// ---------------------------
	// Generate a client token
	tokenGen := qyro.NewClientTokenGenerator(apiKeyID, apiKeySecret)
	clientToken, err := tokenGen.Generate(map[string]interface{}{
		"user_id": "123",
	})
	if err != nil {
		log.Fatal("Error generating client token:", err)
	}

	client, err := qyro.NewQyroClient(baseURL, clientToken, 30*time.Second)
	if err != nil {
		log.Fatal("Error creating client:", err)
	}

	fmt.Println("👉 Creating session via client API...")
	clientSession, err := client.CreateSession(assistantID, map[string]interface{}{
		"context": "from client SDK",
	})
	if err != nil {
		log.Fatal("Error creating client session:", err)
	}
	fmt.Println("✅ Client session created:", clientSession.ID)

	fmt.Println("👉 Sending message to assistant (client)...")
	clientMessages, err := client.Chat(assistantID, clientSession.ID, "Hello from Go Client SDK!")
	if err != nil {
		log.Fatal("Error sending chat:", err)
	}
	for _, m := range clientMessages {
		fmt.Printf("[%s] %s\n", m.Role, m.Content)
	}
}
