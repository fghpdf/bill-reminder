package main

import (
	"github.com/fghpdf/bill-reminder/db"
	"github.com/fghpdf/bill-reminder/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(server.WithHostPorts(":8080"))
	db.InitDB()

	h.POST("/webhook", handler.WebhookHandler)
	h.GET("/withdrawals", handler.QueryHandler)
	h.GET("/ical", handler.ICalHandler)

	h.Spin()
}
