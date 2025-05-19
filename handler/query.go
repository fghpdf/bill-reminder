package handler

import (
	"context"

	"github.com/fghpdf/bill-reminder/db"
	"github.com/fghpdf/bill-reminder/model"

	"github.com/cloudwego/hertz/pkg/app"
)

func QueryHandler(c context.Context, ctx *app.RequestContext) {
	from := string(ctx.Query("from"))
	to := string(ctx.Query("to"))

	rows, err := db.DB.Query(`
		SELECT id, date, description, amount, created_at
		FROM withdrawals
		WHERE date BETWEEN ? AND ?
		ORDER BY date ASC`, from, to)

	if err != nil {
		ctx.String(500, "query failed")
		return
	}

	var results []model.Withdrawal
	for rows.Next() {
		var w model.Withdrawal
		rows.Scan(&w.ID, &w.Date, &w.Description, &w.Amount, &w.CreatedAt)
		results = append(results, w)
	}
	ctx.JSON(200, results)
}
