package trigger

import "log"

func TriggerWebhook(key string) (string, error) {
	log.Println("Triggering webhook with key:", key)
	return key, nil
}
