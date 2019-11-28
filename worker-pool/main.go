package main

import (
	"fmt"
	"sync"
	"time"
)

var NUMBER_WORKER = 4

// Job ...
type Job interface {
	Process()
}

func worker(q chan Job, ks chan bool, wg *sync.WaitGroup) {
	for true {
		select {
		case job := <-q:
			func() {
				defer func() {
					if err := recover(); err != nil {
						wg.Done()
					}
				}()

				job.Process()
				wg.Done()
			}()
		case <-ks:
			fmt.Println("Worker is killed ")
			return
		}
	}
}

// ================================================================

// EmailSender ...
type EmailSender struct {
	Email string
}

// Process ...
func (job EmailSender) Process() {
	fmt.Println("Wellcome", job.Email)
	time.Sleep(time.Second)
}

// ================================================================

func main() {
	emails := []string{
		"1@gmail.com", "2@gmail.com", "3@gmail.com", "4@gmail.com", "5@gmail.com", "6@gmail.com", "7@gmail.com", "8@gmail.com", "9@gmail.com", "10@gmail.com",
		"11@gmail.com", "12@gmail.com", "13@gmail.com", "14@gmail.com", "15@gmail.com", "16@gmail.com", "17@gmail.com", "18@gmail.com", "19@gmail.com", "20@gmail.com",
		"21@gmail.com", "22@gmail.com", "23@gmail.com", "24@gmail.com", "25@gmail.com", "26@gmail.com", "27@gmail.com", "28@gmail.com", "29@gmail.com", "30@gmail.com",
		"31@gmail.com", "32@gmail.com", "33@gmail.com", "34@gmail.com", "35@gmail.com", "36@gmail.com", "37@gmail.com", "38@gmail.com", "39@gmail.com", "40@gmail.com",
	}

	// edit here
	total := len(emails)

	queue := make(chan Job)
	killsignal := make(chan bool)
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(int(total))

	defer func() {
		waitGroup.Wait()
		close(killsignal)

		fmt.Println("Done")
	}()

	numberOfWorkers := NUMBER_WORKER
	for i := 0; i < numberOfWorkers; i++ {
		go worker(queue, killsignal, waitGroup)
	}

	for _, s := range emails {
		// edit here
		queue <- EmailSender{Email: s}
	}
}
