package model

type WebhookInput struct {
	ID                   string `json:"id"`
	ScheduledDate        string `json:"scheduled_date"`         // e.g. 2025年03月27日
	WithdrawItemsAmount  string `json:"withdraw_items_amount"`  // e.g. "95208.00"
	WithdrawItemsContent string `json:"withdraw_items_content"` // e.g. 内容　　：　ペイペイカード
}
