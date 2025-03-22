package trigger

import "fmt"

func TriggerWebhook(key string) (string, error) {
	fmt.Println("Triggering webhook with key:", key)
	return key, nil
}
