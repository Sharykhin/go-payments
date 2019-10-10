package deferrer

type DefferFunc func()

var queue = make([]DefferFunc, 0)

func QueueDeffer(fn DefferFunc) {
	queue = append(queue, fn)
}

func Flush() {
	for _, fn := range queue {
		fn()
	}
}
