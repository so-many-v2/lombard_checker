package app

import (
	"fmt"
	"lombardchecker/internal/clients"
	"lombardchecker/internal/config"
	fr "lombardchecker/internal/fileReader"
	"lombardchecker/pkg/logger"
	"sync"
	"sync/atomic"
)

func StartApp() {
	logger := logger.NewLogger()
	config := config.NewConfig()

	reader := fr.NewFileReader(logger)
	fetcher := clients.NewFetcher(config, logger)

	wallets := reader.ScanFile("data/wallets.txt")

	wg := sync.WaitGroup{}

	var totalAllocation atomic.Int64
	semaphore := make(chan struct{}, config.WorkerAmount)

	for wallets.Scan() {

		address := wallets.Text()
		semaphore <- struct{}{}
		wg.Add(1)
		go func(address string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			amount, err := fetcher.FetchWalletAllocation(address)
			if err != nil {
				fmt.Printf("Error | error fetching wallet data: %s\n", err.Error())
				return
			}

			fmt.Printf("Success | Wallet: %s | Allocation: %d\n", address, amount)
			totalAllocation.Add(amount)

		}(address)
	}

	wg.Wait()
	fmt.Printf("Total allocation: %d", totalAllocation.Load())
}
