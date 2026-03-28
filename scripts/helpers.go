package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func NewConfig() *redis.Options {
	redisURI := os.Getenv("REDIS_URI")
	if redisURI == "" {
		log.Fatal("REDIS_URI is not set")
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	if redisPassword == "" {
		log.Fatal("REDIS_PASSWORD is not set")
	}

	redisDB := os.Getenv("REDIS_DB")
	if redisDB == "" {
		log.Fatal("REDIS_DB is not set")
	}

	db, err := strconv.Atoi(redisDB)
	if err != nil {
		log.Fatal(err)
	}

	return &redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       db,
	}
}

func NewClient() (*redis.Client, func()) {
	config := NewConfig()
	client := redis.NewClient(config)

	return client, func() {
		ctx := context.Background()
		err := client.Close()
		if err != nil {
			log.Println(err)
		}
	}
}

func GetClient() (*redis.Client, error) {
	client := redis.NewClient(NewConfig())
	return client, nil
}

func NewLexicon() map[string]string {
	example := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	return example
}

func GetLexicon() (map[string]string, error) {
	lexicon := NewLexicon()
	return lexicon, nil
}

func NewExample() map[string]interface{} {
	example := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	return example
}

func GetExample() (map[string]interface{}, error) {
	example := NewExample()
	return example, nil
}

func NewData() []map[string]interface{} {
	example := []map[string]interface{}{
		{
			"key1": "value1",
			"key2": "value2",
		},
		{
			"key3": "value3",
			"key4": "value4",
		},
	}
	return example
}

func GetData() ([]map[string]interface{}, error) {
	data := NewData()
	return data, nil
}

func GetDump() (string, error) {
	dump := fmt.Sprintf(`{
	"key1": "value1",
	"key2": "value2"
}`)
	return dump, nil
}