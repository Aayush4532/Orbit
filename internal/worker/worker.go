package worker

import (
	"Orbit/internal/db"
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProductPayload struct {
	ID        bson.ObjectID `bson:"_id"`
	sellerId  bson.ObjectID `bson:"sellerId"`
	eventId   bson.ObjectID `bson:"eventId"`
	Price     float64       `bson:"price"`
	Frequency int           `bson:"frequency"`
}

type WorkerClient struct {
	numWorkers     int
	productChannel chan ProductPayload
	batchSize      int
	workerWg       sync.WaitGroup

	firstErr error
	errMu    sync.Mutex
}

func InitWorkerPool() *WorkerClient {
	wc := &WorkerClient{
		numWorkers:     11,
		productChannel: make(chan ProductPayload, 200),
		batchSize:      100,
	}
	for i := 0; i < wc.numWorkers; i++ {
		wc.workerWg.Add(1)
		go wc.runWorker()
	}
	return wc
}

func (wc *WorkerClient) Send(p ProductPayload) {
	wc.productChannel <- p
}

func (wc *WorkerClient) Close() {
	close(wc.productChannel)
}

func (wc *WorkerClient) Wait() error {
	wc.workerWg.Wait()
	return wc.firstErr
}

func (wc *WorkerClient) runWorker() {
	defer wc.workerWg.Done()

	localBuf := make([]ProductPayload, 0, wc.batchSize)

	for product := range wc.productChannel {
		localBuf = append(localBuf, product)

		if len(localBuf) >= wc.batchSize {
			wc.flush(localBuf)
			localBuf = localBuf[:0]
		}
	}

	if len(localBuf) > 0 {
		wc.flush(localBuf)
	}
}

func (wc *WorkerClient) flush(batch []ProductPayload) {
	toFlush := make([]ProductPayload, len(batch))
	copy(toFlush, batch)

	if err := wc.flushToRedis(toFlush); err != nil {
		log.Printf("worker: flush failed: %v", err)
		wc.errMu.Lock()
		if wc.firstErr == nil {
			wc.firstErr = err
		}
		wc.errMu.Unlock()
	}
}

func (wc *WorkerClient) flushToRedis(batch []ProductPayload) error {
	const maxAttempts = 3
	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		if err := wc.flushOnce(batch); err != nil {
			lastErr = err
			time.Sleep(time.Duration(attempt) * 100 * time.Millisecond)
			continue
		}
		return nil
	}
	return fmt.Errorf("flush failed after %d attempts: %w", maxAttempts, lastErr)
}

func (wc *WorkerClient) flushOnce(batch []ProductPayload) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rdb := db.GetRedisClient()
	pipe := rdb.Pipeline()
	for _, p := range batch {
		priceKey := strconv.FormatFloat(p.Price, 'f', -1, 64)
		key := fmt.Sprintf("product:%s:%s:%s:%s", p.ID.Hex(), p.sellerId.Hex(), p.eventId.Hex(), priceKey)
		pipe.Set(ctx, key, p.Frequency, 0)
	}
	_, err := pipe.Exec(ctx)
	return err
}
