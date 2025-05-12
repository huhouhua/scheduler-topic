package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// Task represents a task with its index, remaining time, and priority.
type Task struct {
	index     int
	remaining int
	priority  int
}

// TaskQueue implements a priority queue for tasks.
type TaskQueue []*Task

func (q TaskQueue) Len() int { return len(q) }

func (q TaskQueue) Less(i, j int) bool {
	if q[i].priority == q[j].priority {
		return q[i].index < q[j].index
	}
	return q[i].priority < q[j].priority
}

func (q TaskQueue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

func (q *TaskQueue) Push(x interface{}) {
	*q = append(*q, x.(*Task))
}

func (q *TaskQueue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func runScheduleTasks(taskTimes []int, bandwidth int) {
	queue := &TaskQueue{}
	heap.Init(queue)
	for i, t := range taskTimes {
		heap.Push(queue, &Task{
			index:     i,
			remaining: t,
			priority:  t,
		})
	}

	time := 0
	fmt.Printf("time\t\tindex\t\tremain_time_slice\n")
	for queue.Len() > 0 {
		var processedTasks []string
		var remainingTasks []string

		timeSlice := 0
		for timeSlice < bandwidth && queue.Len() > 0 {
			task := heap.Pop(queue).(*Task)
			consumed := min(task.remaining, bandwidth-timeSlice)
			task.remaining -= consumed
			timeSlice += consumed

			// record processed task
			processedTasks = append(processedTasks, fmt.Sprintf("%d", task.index))
			// record remaining task
			remainingTasks = append(remainingTasks, fmt.Sprintf("%d-%d", task.index, task.remaining))

			// add task to the queue if not finished
			if task.remaining > 0 {
				heap.Push(queue, task)
			}
		}
		fmt.Printf("%d\t\t%s\t\t%s\n", time, strings.Join(processedTasks, ","), strings.Join(remainingTasks, ","))
		time++
	}
}
func printTaskList(taskTimes []int) {
	fmt.Print("task_index\t")
	for i := range taskTimes {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Print("task_time_slice\t")
	for _, t := range taskTimes {
		fmt.Printf("%d sec ", t)
	}
	fmt.Println()

	for i, t := range taskTimes {
		fmt.Printf("Task %d needs %d seconds to complete.\n", i, t)
	}
	fmt.Println()
}

func main() {
	taskTimes := []int{2, 4, 100, 6, 10, 90, 1, 1, 10, 2, 15, 30, 1, 5, 9, 10}
	bandwidth := 5

	printTaskList(taskTimes)

	// Run task scheduler
	runScheduleTasks(taskTimes, bandwidth)
}
