package service

import (
	"errors"
	"log"
	"nlm/config"
)

func TriggerWebhook(key string, params interface{}, token string) (string, error) {
	if token != config.ENV.WEBHOOK_TOKEN {
		return "", errors.New("invalid token")
	}

	log.Println("Triggering webhook with key:", key)
	return key, nil
}
