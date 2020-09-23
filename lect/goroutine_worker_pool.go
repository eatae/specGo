package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

type Job struct {
	id           int
	randomNumber int
}

type Result struct {
	job Job
	sum int
}

/* действия грузчика (взять коробку, переложить) */
func Digits(number int) int {
	sum := 0
	num := number
	for num != 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}
	time.Sleep(3 * time.Second)
	return sum
}

/* воркер это фура с коробками (сколько/чего там лежит) */
func Worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, Digits(job.randomNumber)}
		results <- output
	}
	wg.Done()
}

/* сколько грузчиков за раз будет заходить в фуру брать ящики и уходить */
func CreateWorkerPool(NumberOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < NumberOfWorkers; i++ {
		wg.Add(1)
		go Worker(&wg)
	}
	wg.Wait()
	close(results)
}

/* какой грузчик какой ящик берет */
func Allocator(NumberOfJobs int) {
	for i := 0; i < NumberOfJobs; i++ {
		randNum := rand.Intn(999)
		job := Job{i, randNum}
		jobs <- job
	}
	close(jobs)
}

/* отмечаем, что такой-то грузчик отнес такой-то ящик */
func ResultFunc(done chan bool) {
	for res := range results {
		fmt.Printf("Job id is %d, value is %d, result is %d\n", res.job.id, res.job.randomNumber, res.sum)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	JobsNumber := 1000
	go Allocator(JobsNumber)
	done := make(chan bool)
	go ResultFunc(done)
	NumOfWorkers := 10
	CreateWorkerPool(NumOfWorkers)
	<-done
	endTime := time.Now()
	deltaT := endTime.Sub(startTime)
	fmt.Println("Total time:", deltaT.Seconds(), " sec.")
}
