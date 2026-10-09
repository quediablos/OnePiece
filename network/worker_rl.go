package network

import (
	"OnePiece/core"
	"OnePiece/message"
	"strconv"
	"time"
)

func ProcessRl(app *core.App) {
	for data := range app.ChanRl {

		app.MutexForRl.Lock()

		if data.Operation == core.OpSetupRateLimiter {
			handleSetupRateLimiter(app, data)
		} else if data.Operation == core.OpWaitForRateLimiter {
			handleWaitForRateLimiter(app, data)
		}

		app.MutexForRl.Unlock()
	}
}

func handleSetupRateLimiter(app *core.App, data core.OperationData) {

	userId := data.ExtraParams[0]
	timeFrame := data.ExtraParams[1]
	rate := data.ExtraParams[2]

	key := core.MakeRlKey(data.ResourceId, userId)

	var timeFrameDuration int64
	if timeFrame == "MINUTE" {
		timeFrameDuration = int64(time.Minute)
	} else {
		timeFrameDuration = int64(time.Second)
	}

	rateInt, _ := strconv.Atoi(rate)
	now := time.Now()

	app.RateLimiters[key] = &core.RateLimiter{
		Rate:       int32(rateInt),
		TimeFrame:  timeFrameDuration,
		LastUsage:  &now,
		TokenCount: float64(rateInt),
	}

	ReleaseClient(data.Conn, message.GenerateSetupRlSuccessfulResponse(app))
}

func handleWaitForRateLimiter(app *core.App, data core.OperationData) {

	userId := data.ExtraParams[0]
	enqueue, _ := strconv.ParseBool(data.ExtraParams[1])

	key := core.MakeRlKey(data.ResourceId, userId)
	rlAvailable, queued, err := core.CheckAndMaintainRlAvailability(app, key, data, enqueue)

	if err != nil {
		ReleaseClient(data.Conn, message.GenerateWaitRlFailedResponse(app, err.ErrorCode, err.ErrorMessage))
		return
	}

	if rlAvailable {
		ReleaseClient(data.Conn, message.GenerateWaitRlSuccessfulResponse(app))
	} else if !queued {
		ReleaseClient(data.Conn, message.GenerateWaitRlFailedResponse(app, "LIMIT_EXCEEDED",
			"Rate limiter exceeded."))
	}
}
