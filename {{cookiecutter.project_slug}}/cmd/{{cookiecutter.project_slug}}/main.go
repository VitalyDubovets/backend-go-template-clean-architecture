package main

import (
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	defer sentry.Flush(3 * time.Second)
}
