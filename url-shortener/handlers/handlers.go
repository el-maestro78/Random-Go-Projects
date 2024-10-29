package handlers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"os"
	"url-shortener/models"
)

const filename = "urls.json"

func saveJson(url models.Url) error {
	var file *os.File
	var err error
	if _, err = os.Stat(filename); err == nil {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	} else if os.IsNotExist(err) {
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ") //format for readability
	return encoder.Encode(url)
}

func Shortener(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hash := hasher.Sum(nil)[:6]
	shortUrl := base64.URLEncoding.EncodeToString(hash)
	urlToSave := models.Url{
		Original: url,
		ShortUrl: shortUrl,
	}
	err := saveJson(urlToSave)
	if err != nil {
		return err.Error()
	}
	return shortUrl
}

func GetOriginalUrl(shortUrl string) (string, error) {
	if _, err := os.Stat(filename); err == nil {
		file, err := os.Open(filename)
		if err != nil {
			return shortUrl, err
		}
		defer file.Close()
		decoder := json.NewDecoder(file)

		// Read the opening bracket of the JSON array
		if _, err := decoder.Token(); err != nil {
			return shortUrl, err
		}
		// Iterating through each object
		for decoder.More() {
			var record models.Url
			if err := decoder.Decode(&record); err != nil {
				return shortUrl, err
			}
			if record.ShortUrl == shortUrl {
				return record.Original, nil
			}
		}
		return shortUrl, nil
	} else {
		return shortUrl, err
	}
}
