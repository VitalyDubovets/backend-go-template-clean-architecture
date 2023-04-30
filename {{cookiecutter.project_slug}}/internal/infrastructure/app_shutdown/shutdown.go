package appshutdown

import (
	"context"
	"os"
	"sync"
	"time"
	"{{ cookiecutter.project_slug }}/internal/infrastructure/logger"
)

type Handler func(ctx context.Context) error

func ShutDownApplication(ctx context.Context, timeout time.Duration, handlers map[string]Handler) <-chan struct{} {
	wait := make(chan struct{})
	log := logger.GetLog()

	go func() {
		<-ctx.Done()

		log.Info("shitting down")

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Info("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})
		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, handler := range handlers {
			innerKey := key
			innerHandler := handler

			wg.Add(1)

			go func() {
				defer wg.Done()

				log.Infof("Shutting down: %s", innerKey)

				if err := innerHandler(ctx); err != nil {
					log.Infof("%s: shutdown was failed: %s", innerKey, err)
					return
				}

				log.Infof("%s was shut down", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
