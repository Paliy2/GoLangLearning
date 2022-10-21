package main

import (
	"fmt"
	"time"
)

type Retriable interface {
	GetMaxAttempts() uint
	Retry(action func() bool)
}

type DefaultRetriable struct{}

func (dr DefaultRetriable) GetMaxAttempts() uint {
	return 5
}

func (dr DefaultRetriable) Retry(action func() bool) {
	var counter uint = 0
	for counter < dr.GetMaxAttempts() {
		counter += 1
		if action() == true {
			fmt.Printf("Action Comlpeted from the attempt %d", counter)
			break
		}
		fmt.Printf("Action Failed. Retry #%d\n", counter)
	}
}

type RetriableWithDelay struct {
	DelayInSec time.Duration
}

func (rwd RetriableWithDelay) GetMaxAttempts() uint {
	return 5
}

func (rwd RetriableWithDelay) Retry(action func() bool) {
	var counter uint = 0
	for counter < rwd.GetMaxAttempts() {
		counter += 1
		if action() == true {
			fmt.Printf("Action Comlpeted from the attempt %d", counter)
			break
		}
		fmt.Printf("Action Failed. Retry #%d\n", counter)
		time.Sleep(rwd.DelayInSec)
	}
}

func someFalseAction() bool {
	return false
}

func someTrueAction() bool {
	return true
}

func main() {
	var defaultRetriable = DefaultRetriable{}
	fmt.Printf("DefaultRetriable Max attempts: %d\n", defaultRetriable.GetMaxAttempts())
	defaultRetriable.Retry(someFalseAction)
	defaultRetriable.Retry(someTrueAction)

	var retriableWithDelay = RetriableWithDelay{
		DelayInSec: 10,
	}
	fmt.Printf("RetriableWithDelay Max attempts: %d\n", retriableWithDelay.GetMaxAttempts())
	fmt.Printf("RetriableWithDelay Delay in second: %d\n", retriableWithDelay.DelayInSec)
	retriableWithDelay.Retry(someFalseAction)
	retriableWithDelay.Retry(someTrueAction)
}
