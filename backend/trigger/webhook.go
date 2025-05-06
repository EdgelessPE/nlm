package trigger

import (
	"fmt"
	"log"
	"nlm/pipeline"
)

func TriggerWebhook(event string) (string, error) {
	log.Println("Triggering webhook with event:", event)
	switch event {
	case "ping":
		return "pong", nil
	case "release":
		ctx := pipeline.RunEptPipeline()
		return ctx.Id, nil
	default:
		return "", fmt.Errorf("invalid webhook event: %s", event)
	}
}
