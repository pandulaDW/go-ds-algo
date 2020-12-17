package queues

import "testing"

// adding number of elements and removing them
func checkPerformance(q Queue, numElements int) {
	for i := 0; i < numElements; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < numElements; i++ {
		q.Dequeue()
	}
}

func BenchmarkPerformanceQueueArray(t *testing.B) {
	q := CreateQueueUsingArray(100)
	for i := 0; i < t.N; i++ {
		checkPerformance(q, 100)
	}
}

func BenchmarkPerformanceQueueList(t *testing.B) {
	q := CreateQueueUsingList()
	for i := 0; i < t.N; i++ {
		checkPerformance(q, 100)
	}
}
