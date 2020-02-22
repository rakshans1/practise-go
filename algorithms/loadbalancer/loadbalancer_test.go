package loadbalancer

import (
	"fmt"
	"testing"
)

func (b *Balancer) print() {
	sum := 0
	sumq := 0

	// Print pending stats for each worker
	for _, w := range b.pool {
		fmt.Printf("%d ", w.pending)
		sum += w.pending
		sumq += w.pending * w.pending
	}

	// Print avg for worker pool
	avg := float64(sum) / float64(len(b.pool))
	varaince := float64(sumq)/float64(len(b.pool)) - avg*avg
	fmt.Printf(" %.2f %.2f\n", avg, varaince)
}

func TestLoadbalancer(t *testing.T) {
	t.Run("Should return result of Request", func(t *testing.T) {

		const nRequester = 1000

		const nWorker = 2

		work := make(chan Request)
		for i := 0; i < nRequester; i++ {
			go requester(work)
		}
		InitBalancer(nWorker, nRequester).balance(work)
	})
}
