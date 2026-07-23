package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type BenchmarkMetrics struct {
	Success       int64
	SoldOut       int64
	AlreadyBooked int64
	Errors        int64

	Start time.Time
	End   time.Time

	mu        sync.Mutex
	Latencies []time.Duration
}

var Metrics BenchmarkMetrics

func BenchmarkStart() {
	Metrics = BenchmarkMetrics{}
	Metrics.Start = time.Now()
}

func BenchmarkFinish() {
	Metrics.End = time.Now()
}

func AddLatency(d time.Duration) {
	Metrics.mu.Lock()
	Metrics.Latencies = append(Metrics.Latencies, d)
	Metrics.mu.Unlock()
}

func IncSuccess() {
	atomic.AddInt64(&Metrics.Success, 1)
}

func IncSoldOut() {
	atomic.AddInt64(&Metrics.SoldOut, 1)
}

func IncAlreadyBooked() {
	atomic.AddInt64(&Metrics.AlreadyBooked, 1)
}

func IncError() {
	atomic.AddInt64(&Metrics.Errors, 1)
}

func PrintBenchmark() {

	success := atomic.LoadInt64(&Metrics.Success)
	soldOut := atomic.LoadInt64(&Metrics.SoldOut)
	alreadyBooked := atomic.LoadInt64(&Metrics.AlreadyBooked)
	errCount := atomic.LoadInt64(&Metrics.Errors)

	total := success + soldOut + alreadyBooked + errCount

	duration := Metrics.End.Sub(Metrics.Start)

	var avgLatency time.Duration

	Metrics.mu.Lock()
	if len(Metrics.Latencies) > 0 {
		var sum time.Duration
		for _, d := range Metrics.Latencies {
			sum += d
		}
		avgLatency = sum / time.Duration(len(Metrics.Latencies))
	}
	Metrics.mu.Unlock()

	var throughput float64

	if duration.Seconds() > 0 {
		throughput = float64(total) / duration.Seconds()
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("BENCHMARK REPORT")
	fmt.Println("========================================")
	fmt.Printf("Total Requests   : %d\n", total)
	fmt.Printf("Success          : %d\n", success)
	fmt.Printf("Sold Out         : %d\n", soldOut)
	fmt.Printf("Already Booked   : %d\n", alreadyBooked)
	fmt.Printf("Errors           : %d\n", errCount)
	fmt.Printf("Duration         : %v\n", duration)
	fmt.Printf("Average Latency  : %v\n", avgLatency)
	fmt.Printf("Throughput       : %.2f req/sec\n", throughput)
	fmt.Println("========================================")
}
