package network

import (
	"OnePiece/core"
	"OnePiece/message"
	"strconv"
)

func ProcessStocks(app *core.App) {
	for data := range app.ChanStocks {

		app.MutexForStocks.Lock()

		if data.Operation == core.OpReserveStock {

			stockReserve, success := app.ReserveStock(data.ResourceId)

			if success {
				ReleaseClient(data.Conn, message.GenerateReserveStockSuccessfulResponse(app, stockReserve.Id, data.ResourceId))
			} else {
				ReleaseClient(data.Conn, message.GenerateReserveStockFailedResponse(app, data.ResourceId))
			}

		} else if data.Operation == core.OpReleaseStock {

			app.ReleaseStock(data.ResourceId, data.ExtraParams[0])
			msg := message.GenerateReleaseStockSuccessfulResponse(app, data.ResourceId)
			ReleaseClient(data.Conn, msg)

		} else if data.Operation == core.OpCreateStock {

			quantity, _ := strconv.ParseInt(data.ExtraParams[0], 10, 64)

			success := app.CreateStock(data.ResourceId, quantity)

			var msg string
			if success {
				msg = message.GenerateCreateStockSuccessfulResponse(app, data.ResourceId, data.ExtraParams[0])
			} else {
				msg = message.GenerateCreateStockFailedResponse(app, data.ResourceId)
			}

			ReleaseClient(data.Conn, msg)
		}

		app.MutexForStocks.Unlock()
	}
}
