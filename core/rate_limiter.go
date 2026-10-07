package core

import "time"

type RateLimiter struct {
	TokenCount float64
	Rate       int32
	TimeFrame  int64
	LastUsage  *time.Time
	Queue      []OperationData // Waiting clients when no tokens are available (FIFO)
}

// CheckRlAvailability Checks if there is availability in the rate. If there is availability, one token is used.
// ------------- THREAD-SAFE: This method needs to run thread-safe -------------
func CheckRlAvailability(app *App, key string) (bool, *Error) {

	rateLimiter := app.RateLimiters[key]

	if rateLimiter == nil {
		return false, &Error{
			ErrorCode:    "RATE_LIMITER_NOT_FOUND",
			ErrorMessage: "Rate limiter not found.",
		}
	}

	//First add the tokens that the bucket gained during cooldown.
	tokensToAddPerSecond := float64(rateLimiter.Rate) / (float64(rateLimiter.TimeFrame) / 1_000_000_000)

	if rateLimiter.LastUsage != nil {
		elapsed := time.Since(*rateLimiter.LastUsage)
		tokensToAdd := tokensToAddPerSecond * elapsed.Seconds()
		rateLimiter.TokenCount += tokensToAdd
		// Cap at max bucket capacity (Rate tokens per TimeFrame)
		if maxTokens := float64(rateLimiter.Rate); rateLimiter.TokenCount > maxTokens {
			rateLimiter.TokenCount = maxTokens
		}
	}

	now := time.Now()
	rateLimiter.LastUsage = &now

	if rateLimiter.TokenCount < 1 {
		return false, nil
	}

	rateLimiter.TokenCount--
	return true, nil

}

func MakeRlKey(resourceId string, userId string) string {
	return resourceId + ":" + userId
}
