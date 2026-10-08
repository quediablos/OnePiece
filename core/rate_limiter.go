package core

import "time"

type RateLimiter struct {
	TokenCount     float64
	Rate           int32
	TimeFrame      int64
	LastUsage      *time.Time
	WaitingClients []OperationData
}

// CheckAndMaintainRlAvailability Checks if there is availability in the rate. If there is availability, one token is used.
// ------------- THREAD-SAFE: This method needs to run thread-safe -------------
//
// Returns:
//   - bool: tokenAvailable — true if a token was available and consumed, false otherwise.
//   - bool: queued        — true if the client was added to WaitingClients (only possible when holdOption is true and no token was available).
//   - *Error: err         — non-nil if a fatal error occurred (e.g. rate limiter not found); nil on normal operation.
func CheckAndMaintainRlAvailability(app *App, key string, operationData OperationData, holdOption bool) (bool, bool, *Error) {

	rateLimiter := app.RateLimiters[key]

	if rateLimiter == nil {
		return false, false, &Error{
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
		if holdOption {
			rateLimiter.WaitingClients = append(rateLimiter.WaitingClients, operationData)
		}
		return false, holdOption, nil
	}

	rateLimiter.TokenCount--
	return true, false, nil

}

func MakeRlKey(resourceId string, userId string) string {
	return resourceId + ":" + userId
}
