package utils

import (
	"time"
)

const (
	tokenTimeCurrency    time.Duration = time.Hour
	normalActionDuration time.Duration = 3 * tokenTimeCurrency
	refreshTokenDuration time.Duration = 24 * 7 * tokenTimeCurrency
	adminLockDuration    time.Duration = time.Minute * 15
	staffLockDuration    time.Duration = time.Minute * 30
)

func isActionExpired(period time.Time, duration time.Duration) bool {
	return time.Now().After(period.Add(duration)) // The amount of token expiration for normal action
}

func GetPrimitiveTime() time.Time {
	return time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
}

func IsAuthenticationLevelExpired(period time.Time) bool {
	return isActionExpired(period, 0)
}
