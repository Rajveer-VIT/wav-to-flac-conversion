package main

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// Function to convert WAV data to FLAC using FFmpeg
func convertWavToFlac(wavData []byte) ([]byte, error) {
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "flac", "pipe:1")

	// Set up input and output for the FFmpeg command
	cmd.Stdin = bytes.NewReader(wavData)
	var flacData bytes.Buffer
	cmd.Stdout = &flacData

	// Run the command and check for errors
	if err := cmd.Run(); err != nil {
		log.Println("FFmpeg conversion error:", err)
		return nil, err
	}

	return flacData.Bytes(), nil
}

func handleWebSocket(c *websocket.Conn) {
	defer c.Close()

	for {
		// Read message from the client (expecting WAV data in binary format)
		_, wavData, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Log the received WAV data
		log.Printf("Received WAV data: %d bytes\n", len(wavData))

		// Convert WAV to FLAC
		flacData, err := convertWavToFlac(wavData)
		if err != nil {
			log.Println("Error converting WAV to FLAC:", err)
			break
		}

		// Send the converted FLAC data back to the client
		if err := c.WriteMessage(websocket.BinaryMessage, flacData); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Route for WebSocket connection
	app.Get("/convert", websocket.New(handleWebSocket))

	// Start the server on port 8080
	log.Fatal(app.Listen(":8080"))
}
