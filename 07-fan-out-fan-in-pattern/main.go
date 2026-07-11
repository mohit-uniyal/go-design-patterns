package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Log struct {
	Message string
}

// Stage 1 --> receive the logs
func ReceiveLogs() <-chan Log {
	out := make(chan Log)

	// go routine for simulating receiving the log and sending it to the channel
	go func() {
		defer close(out)

		logs := []Log{
			{"INFO User login"},
			{"ERROR Payment failed"},
			{"INFO Profile updated"},
			{"ERROR Database timeout"},
		}

		for _, log := range logs {
			out <- log
		}
	}()

	return out
}

type ParsedLog struct {
	Text string
}

func parseLogHeavy(log Log) ParsedLog {
	time.Sleep(8 * time.Second)
	return ParsedLog{
		Text: strings.ToUpper(log.Message),
	}
}

const (
	WorkerCount int = 4
)

// Stage 2 --> Parse the logs(change to uppercase)
// ***********Implementation: 1***********
func ParseLogs(logs <-chan Log) <-chan ParsedLog {
	out := make(chan ParsedLog)
	var wg sync.WaitGroup

	//go routine to parse the logs and sending to the channel
	for workerNo := 1; workerNo <= WorkerCount; workerNo++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for log := range logs {
				parsedLog := parseLogHeavy(log)
				out <- parsedLog
			}

		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

type Result struct {
	Text    string
	IsError bool
}

// Stage 3 --> Detect Errors
func DetectErrors(parsedLogs <-chan ParsedLog) <-chan Result {
	results := make(chan Result)

	// go routine to detect errors and sending to the channel
	go func() {
		defer close(results)
		for parsedLog := range parsedLogs {
			result := Result{
				Text:    parsedLog.Text,
				IsError: strings.Contains(parsedLog.Text, "ERROR"),
			}

			results <- result
		}
	}()

	return results
}

// Consume / Store

func Store(results <-chan Result) {
	for result := range results {
		fmt.Printf("%s Error=%v\n", result.Text, result.IsError)
	}
}

func main() {
	// **********Pipeline**********
	// A sequence of independent stages where each stage receives data, performs one job, and passes the result to the next stage.

	logsChan := ReceiveLogs()

	parsedLogsChan := ParseLogs(logsChan)

	resultsChan := DetectErrors(parsedLogsChan)

	Store(resultsChan)

}
