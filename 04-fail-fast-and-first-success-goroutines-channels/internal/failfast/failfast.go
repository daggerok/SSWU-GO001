package failfast

import (
	"context"
	"fmt"
	"time"

	"github.com/daggerok/SSWU-GO001/04-fail-fast-and-first-success-goroutines-channels/internal/fetch"
	"golang.org/x/sync/errgroup"
)

// FailFast fails with error or timeout if at least one non-success result
func FailFast(urls []string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	for _, url := range urls {
		g.Go(func() error {
			res, err := fetch.Fetch("FailFast", ctx, url)
			if err != nil {
				return err
			}
			fmt.Println("Done:", res)
			return nil
		})
	}
	return "OK", g.Wait()
}
