package vo

type WebhookRequest struct {
	Key string `json:"key" binding:"required"`
}
