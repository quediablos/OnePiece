package core

import "time"

type RateLimiter struct {
	TokenCount     float64
	Rate           int32
	TimeFrame      int64
	LastUsage      *time.Time
	WaitingClients []OperationData //Map of each resourceId:userId pair. Can reach which clients are queued.
}

// CheckAndMaintainRlAvailability Checks if there is availability in the rate. If there is availability, one token is used.
// ------------- THREAD-SAFE: This method needs to run thread-safe -------------
//
// Returns:
//   - bool: tokenAvailable — true if a token was available and consumed, false otherwise.
//   - bool: queued        — true if the client was added to WaitingClients (only possible when holdOption is true and no token was available).
//   - *Error: err         — non-nil if a fatal error occurred (e.g. rate limiter not found); nil on normal operation.
func CheckAndMaintainRlAvailability(app *App, key string, operationData OperationData, enqueue bool) (bool, bool, *Error) {

	rateLimiter := app.RateLimiters[key]

	if rateLimiter == nil {
		return false, false, &Error{
			ErrorCode:    "RATE_LIMITER_NOT_FOUND",
			ErrorMessage: "Rate limiter not found.",
		}
	}

	//First add the tokens that the bucket gained during cooldown.
	tokensToAddPerSecond := rateLimiter.TokensToAddPerSecond()

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
		if enqueue {
			rateLimiter.WaitingClients = append(rateLimiter.WaitingClients, operationData)

			prioritizedClient := PrioritizedClient{
				OperationData: operationData,
				Priority:      rateLimiter.CalculateProximity(),
			}
			app.WaitingClientsAll.Push(prioritizedClient)
		}
		return false, enqueue, nil
	}

	rateLimiter.TokenCount--
	return true, false, nil

}

func MakeRlKey(resourceId string, userId string) string {
	return resourceId + ":" + userId
}

// CalculateProximity calculates how many seconds until the rate limiter can accept the client.
// ------------- THREAD-SAFE: This method needs to run thread-safe -------------
func (rl *RateLimiter) CalculateProximity() float64 {

	tokensToAddPerSecond := rl.TokensToAddPerSecond()
	costOfSelf := rl.HowLongUntilNextAllocation()

	//-1 for self
	costOfRest := float64(len(rl.WaitingClients)-1) / tokensToAddPerSecond

	return costOfSelf + costOfRest
}

// TokensToAddPerSecond Calculates how many tokens are added per second.
// ------------- THREAD-SAFE: This method needs to run thread-safe -------------
func (rl *RateLimiter) TokensToAddPerSecond() float64 {
	return float64(rl.Rate) / (float64(rl.TimeFrame) / 1_000_000_000)
}

// HowLongUntilNextAllocation Calculates how long it will take until the next allocation for the next client,
// considering there are no other clients already queued.
// ------------- THREAD-SAFE: This method needs to run thread-safe -------------
func (rl *RateLimiter) HowLongUntilNextAllocation() float64 {
	tokensToAddPerSecond := rl.TokensToAddPerSecond()
	currentTokens := rl.TokenCount
	return (1 - currentTokens) / tokensToAddPerSecond
}
