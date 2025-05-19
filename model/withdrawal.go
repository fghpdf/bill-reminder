package model

type Withdrawal struct {
	ID          string `json:"id"`
	Date        string `json:"date"`        // YYYY-MM-DD
	Description string `json:"description"` // 内容
	Amount      int    `json:"amount"`      // 单位：円
	CreatedAt   string `json:"created_at"`  // RFC3339
}
