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

		userId := data.ExtraParams[0]
		now := time.Now()

		if data.Operation == core.OpSetupRateLimiter {

			key := makeKey(data.ResourceId, userId)

			timeFrame := data.ExtraParams[1]
			rate := data.ExtraParams[2]

			var timeFrameDuration int64
			if timeFrame == "MINUTE" {
				timeFrameDuration = int64(time.Minute)
			} else {
				timeFrameDuration = int64(time.Second)
			}

			rateInt, _ := strconv.Atoi(rate)

			app.RateLimiters[key] = &core.RateLimiter{
				Rate:       int32(rateInt),
				TimeFrame:  timeFrameDuration,
				LastUsage:  &now,
				TokenCount: float64(rateInt),
			}

			ReleaseClient(data.Conn, message.GenerateSetupRlSuccessfulResponse(app))

		} else if data.Operation == core.OpWaitForRateLimiter {

			key := makeKey(data.ResourceId, userId)
			rlAvailable := core.CheckRlAvailability(app.RateLimiters[key])

			if rlAvailable {
				ReleaseClient(data.Conn, message.GenerateWaitRlSuccessfulResponse(app))
			} else {
				ReleaseClient(data.Conn, message.GenerateWaitRlFailedResponse(app, "LIMIT_EXCEEDED",
					"Rate limiter exceeded."))
			}
		}

		app.MutexForRl.Unlock()
	}
}

func makeKey(resourceId string, userId string) string {
	return resourceId + ":" + userId
}
