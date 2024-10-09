package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Response struct {
	GeneratedText string `json:"generated_text"`
}

func loading(delay float32) {
	remainingTime := int(delay) + 1
	dots := remainingTime / 2
	stars := 0
	progress := strings.Repeat(".", remainingTime/2)
	fmt.Printf("{%s}", progress)
	for i := 1; i <= remainingTime; i++ {
		time.Sleep(2 * time.Second)
		stars++
		dots--
		progressStars := strings.Repeat("*", stars)
		progressDots := strings.Repeat(".", dots)
		progress = progressStars + progressDots
		fmt.Printf("\r{%s}", progress)
	}
	fmt.Println()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	huggingFaceAPIURL := os.Getenv("API")
	apiKey := os.Getenv("KEY")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter input: ")
		if !scanner.Scan() {
			break
		}
		inputText := scanner.Text()
		/*
			payload := map[string]interface{}{
					"messages": []map[string]string{
						{
							"role":    "user",
							"content": inputText,
						},
					},
				}
		*/
		payload := map[string]interface{}{
			"inputs": inputText,
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
		func() {
			req, err := http.NewRequest("POST", huggingFaceAPIURL, bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response:", err)
				return
			}
			//fmt.Printf("Raw response body: %s\n", body)
			var response []Response
			err = json.Unmarshal(body, &response)
			if err != nil {
				fmt.Println("Error unmarshalling JSON:", err)
				return
			}
			fmt.Println("Generated Text: ", response[0].GeneratedText)
		}()
	}
}
