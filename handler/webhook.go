package handler

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/fghpdf/bill-reminder/db"
	"github.com/fghpdf/bill-reminder/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type WebhookInput struct {
	ID        string `json:"id"`
	Date      string `json:"scheduled_date"`
	AmountStr string `json:"withdraw_items_amount"`
	Content   string `json:"withdraw_items_content"`
}

func WebhookHandler(c context.Context, ctx *app.RequestContext) {
	var input model.WebhookInput
	if err := ctx.BindAndValidate(&input); err != nil {
		ctx.String(400, "Invalid input: %v", err)
		return
	}

	// 转换金额
	amountF, err := strconv.ParseFloat(input.WithdrawItemsAmount, 64)
	if err != nil {
		ctx.String(400, "Invalid amount")
		return
	}
	amount := int(amountF)

	// 提取描述（切掉“内容 ：”前缀）
	desc := extractDescription(input.WithdrawItemsContent)

	// 处理日期：2025年03月27日 → 2025-03-27
	date := convertDate(input.ScheduledDate)

	record := model.Withdrawal{
		ID:          input.ID + "_" + input.WithdrawItemsContent,
		Date:        date,
		Description: desc,
		Amount:      amount,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	_, err = db.DB.Exec(`
		INSERT OR IGNORE INTO withdrawals (id, date, description, amount, created_at)
		VALUES (?, ?, ?, ?, ?)`,
		record.ID, record.Date, record.Description, record.Amount, record.CreatedAt)

	if err != nil {
		ctx.String(500, "DB Error: %v", err)
		return
	}

	ctx.JSON(200, record)
}

func convertDate(raw string) string {
	t, err := time.Parse("2006年01月02日", raw)
	if err != nil {
		return raw // fallback
	}
	return t.Format("2006-01-02")
}

func extractDescription(raw string) string {
	// 提取“内容　　：　...”之后的部分
	parts := strings.Split(raw, "：")
	if len(parts) > 1 {
		return strings.TrimSpace(parts[1])
	}
	return strings.TrimSpace(raw)
}
