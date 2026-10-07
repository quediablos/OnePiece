package core

import "time"

type RateLimiter struct {
	TokenCount float64
	Rate       int32
	TimeFrame  int64
	LastUsage  *time.Time
}

// CheckRlAvailability Checks if there is availability in the rate. If there is availability, one token is used.
// ------------- THREAD-SAFE: This method needs to run thread-safe -------------
func CheckRlAvailability(rateLimiter *RateLimiter) bool {

	//TODO:check if rate limiter exists

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
		return false
	}

	rateLimiter.TokenCount--
	return true

}
