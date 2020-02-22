package loadbalancer

import (
	"container/heap"
	"math"
	"math/rand"
	"time"
)

type Request struct {
	data int          // The operation to perform
	c    chan float64 // The channel to return the result
}

type Worker struct {
	requests chan Request // work to do (buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in the heap
}

type Pool []*Worker

type Balancer struct {
	pool Pool
	done chan *Worker
}

func requester(work chan Request) {
	c := make(chan float64)

	for {
		// Kill some time (fake load).
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		r := Request{int(rand.Int31n(90)), c} // send request
		work <- r
		<-c // wait for answer
	}
}

func (w *Worker) work(done chan *Worker) {
	for req := range w.requests {
		// req := <-w.requests                  // get Request from balancer
		req.c <- math.Sin(float64(req.data)) // call fn and send result
		done <- w                            // we've finished this request
	}
}

func InitBalancer(nWorker int, nRequester int) *Balancer {
	done := make(chan *Worker, nWorker)
	p := make(Pool, 0, nWorker)
	// create nWorker
	b := &Balancer{p, done}
	for i := 0; i < nWorker; i++ {
		w := &Worker{requests: make(chan Request, nRequester)}
		heap.Push(&b.pool, w)
		go w.work(b.done)
	}
	return b
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work: // received a Request...
			b.dispatch(req) // ...so send it to a Worker
		case w := <-b.done: // a worker has finished...
			b.complete(w) // ...so update its info
		}
		b.print()
	}
}

// Send Request to worker
func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker...
	w := heap.Pop(&b.pool).(*Worker)
	// ...send it the task.
	w.requests <- req
	// One more in its work queue.
	w.pending++
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

// Job is complete; update heap
func (b *Balancer) complete(w *Worker) {
	// One fewer in the queue
	w.pending--
	// Remove it from the heap.
	heap.Remove(&b.pool, w.index)
	// Put it into its place  on the heap.
	heap.Push(&b.pool, w)
}

func (p Pool) Len() int { return len(p) }

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p *Pool) Swap(i, j int) {
	a := *p
	a[i], a[j] = a[j], a[i]
	a[i].index = i
	a[j].index = j
}

func (p *Pool) Push(x interface{}) {
	n := len(*p)
	item := x.(*Worker)
	item.index = n
	*p = append(*p, item)
}

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*p = old[0 : n-1]
	return item
}
