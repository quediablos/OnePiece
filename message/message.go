package message

import "OnePiece/core"

func GenerateAcquireLockResponse(app *core.App, resourceId string) string {
	if app.Config.ListenHttp {
		return GenerateAcquireLockHttpResponse(resourceId)
	}
	return GenerateAcquireLockTcpResponse(resourceId)
}

func GenerateReleaseLockResponse(app *core.App, resourceId string) string {
	if app.Config.ListenHttp {
		return GenerateReleaseLockHttpResponse(resourceId)
	}
	return GenerateReleaseLockTcpResponse(resourceId)
}

func GenerateReserveStockSuccessfulResponse(app *core.App, reserveId string, resourceId string) string {
	if app.Config.ListenHttp {
		return GenerateReserveStockSuccessfulHttpResponse(reserveId, resourceId)
	}
	return GenerateReserveStockSuccessfulTcpResponse(reserveId, resourceId)
}

func GenerateReserveStockFailedResponse(app *core.App, resourceId string) string {
	if app.Config.ListenHttp {
		return GenerateReserveStockFailedHttpResponse(resourceId)
	}
	return GenerateReserveStockFailedTcpResponse(resourceId)
}

func GenerateCreateStockResponse(app *core.App, resourceId string, quantity string) string {
	if app.Config.ListenHttp {
		return GenerateCreateStockHttpResponse(resourceId, quantity)
	}
	return GenerateCreateStockTcpResponse(resourceId, quantity)
}

func GenerateGenericErrorResponse(app *core.App) string {
	if app.Config.ListenHttp {
		return GenerateGenericErrorHttpResponse()
	}
	return GenerateGenericErrorTcpResponse()
}

func GenerateErrorResponse(app *core.App, err string) string {
	if app.Config.ListenHttp {
		return GenerateErrorHttpResponse(err)
	}
	return GenerateErrorTcpResponse(err)
}
