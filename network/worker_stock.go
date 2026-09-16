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
			//TODO:implement
		} else if data.Operation == core.OpCreateStock {

			quantity, _ := strconv.ParseInt(data.ExtraParams[0], 10, 64)
			app.CreateStock(data.ResourceId, quantity)

			ReleaseClient(data.Conn, message.GenerateCreateStockResponse(app, data.ResourceId, data.ExtraParams[0]))
		}

		app.MutexForStocks.Unlock()
	}
}
