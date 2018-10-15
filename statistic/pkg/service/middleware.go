package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	io "github.com/jscastelblancoh/statistic_service/statistic/pkg/io"
)

// Middleware describes a service middleware.
type Middleware func(StatisticService) StatisticService

type loggingMiddleware struct {
	logger log.Logger
	next   StatisticService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a StatisticService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next StatisticService) StatisticService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context) (t []io.Statistic, err error) {
	defer func() {
		l.logger.Log("method", "Get", "t", t, "err", err)
	}()
	return l.next.Get(ctx)
}
func (l loggingMiddleware) GetbyId(ctx context.Context, id string) (t []io.Statistic, err error) {
	defer func() {
		l.logger.Log("method", "GetbyId", "id", id, "t", t, "err", err)
	}()
	return l.next.GetbyId(ctx, id)
}
func (l loggingMiddleware) Put(ctx context.Context, id string) (t io.Statistic, err error) {
	defer func() {
		l.logger.Log("method", "Put", "id", id, "t", t, "err", err)
	}()
	return l.next.Put(ctx, id)
}
func (l loggingMiddleware) Delete(ctx context.Context, id string) (err error) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "err", err)
	}()
	return l.next.Delete(ctx, id)
}
func (l loggingMiddleware) Post(ctx context.Context, statistic io.Statistic) (t io.Statistic, err error) {
	defer func() {
		l.logger.Log("method", "Post", "statistic", statistic, "t", t, "err", err)
	}()
	return l.next.Post(ctx, statistic)
}
