package vo

type WebhookRequest struct {
	Key    string      `json:"key" binding:"required"`
	Params interface{} `json:"params" binding:"required"`
	Token  string      `json:"token" binding:"required"`
}
