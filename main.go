package main

import (
	"log"
	"os"
	"unisun/api/course-listenner/src"
	"unisun/api/course-listenner/src/config"
	"unisun/api/course-listenner/src/constants"

	"github.com/getsentry/sentry-go"
)

func main() {
	if os.Getenv(constants.NODE) != constants.PRODUCTION {
		config.ConfigENV()
	} else {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: "https://83b609f06aaa4f0cb47558de66e98a2e@o1238561.ingest.sentry.io/6389373",
			// Specify a fixed sample rate:
			TracesSampleRate: 0.2,
			// Or provide a custom sampler:
			TracesSampler: sentry.TracesSamplerFunc(func(ctx sentry.SamplingContext) sentry.Sampled {
				return sentry.SampledTrue
			}),
		})
		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		}
	}
	r := src.App()
	port := os.Getenv(constants.PORT)
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
