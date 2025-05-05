package service

import (
	"log"
)

func TriggerWebhook(key string, params interface{}) (string, error) {
	log.Println("Triggering webhook with key:", key)
	return key, nil
}
