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
				ReleaseClientHttp(data.Conn, message.GenerateReserveStockSuccessfulResponse(stockReserve.Id, data.ResourceId))
			} else {
				ReleaseClientHttp(data.Conn, message.GenerateReserveStockFailedResponse(data.ResourceId))
			}

		} else if data.Operation == core.OpReleaseStock {

		} else if data.Operation == core.OpCreateStock {

			quantity, _ := strconv.ParseInt(data.ExtraParams[0], 10, 64)
			app.CreateStock(data.ResourceId, quantity)

			ReleaseClientHttp(data.Conn, message.GenerateCreateStockResponse(data.ResourceId, data.ExtraParams[0]))
		}

		app.MutexForStocks.Unlock()
	}
}
