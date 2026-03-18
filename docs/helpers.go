package analytics_worker

import (
	"context"
	"math/rand"
	"time"
	"unicode/utf8"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/lib/pq"
)

func validateUUID(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}
	if !utf8.ValidString(uuid) {
		return false
	}
	return true
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func extractError(err error, logger log.Logger) {
	if err != nil {
		// Log any database errors.
		if pq.ErrorCode(err) == pq.ErrorCodeUnknown {
			level.Error(logger).Log("err", err)
		} else if pq.ErrorCode(err) == pq.ErrorCodeUniqueViolation {
			level.Error(logger).Log("err", err)
		} else {
			level.Error(logger).Log("err", err)
		}
	}
}

func extractErrorWithCode(err error, code string, logger log.Logger) {
	if err != nil {
		// Log any database errors.
		if code == "23505" {
			level.Error(logger).Log("err", err)
		} else {
			extractError(err, logger)
		}
	}
}

func isNil(interface{}) bool {
	if v, ok := interface{}(nil); ok {
		return ok
	}
	return false
}

func contextTimeout(ctx context.Context, duration time.Duration) context.Context {
	return context.WithTimeout(ctx, duration)
}

func contextWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}