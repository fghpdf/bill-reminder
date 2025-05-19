package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/fghpdf/bill-reminder/db"
)

func ICalHandler(c context.Context, ctx *app.RequestContext) {
	rows, err := db.DB.Query("SELECT date, description, amount FROM withdrawals")
	if err != nil {
		ctx.String(500, "query error")
		return
	}
	defer rows.Close()

	ical := "BEGIN:VCALENDAR\r\nVERSION:2.0\r\nPRODID:-//bill-reminder//EN\r\n"

	for rows.Next() {
		var dateStr, desc string
		var amt int
		rows.Scan(&dateStr, &desc, &amt)

		t, _ := time.Parse("2006-01-02", dateStr)
		dt := t.Format("20060102")

		ical += fmt.Sprintf("BEGIN:VEVENT\r\nSUMMARY:引落 %s (%d円)\r\nDTSTART;VALUE=DATE:%s\r\nDTEND;VALUE=DATE:%s\r\nDESCRIPTION:%s\r\nSTATUS:CONFIRMED\r\nEND:VEVENT\r\n", desc, amt, dt, dt, desc)
	}

	ical += "END:VCALENDAR\r\n"

	ctx.Response.Header.SetContentType("text/calendar; charset=utf-8")
	ctx.String(200, ical)
}
