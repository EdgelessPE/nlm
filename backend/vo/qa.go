package vo

type QaResult struct {
	Scope            string `json:"scope"`
	TaskName         string `json:"taskName"`
	IsSuccess        bool   `json:"isSuccess"`
	ResultStorageKey string `json:"resultStorageKey"`
}
