# scheduler-topic

Now, we have the following limited sequence of positive integers, which can be stored in the single machine
memory.
```text
  2,4,100,6,10,90,1,1,10,2,15,30,1,5,9,10
```

This sequence represents the time slice that needs to be consumed for the task execution to complete. The
smallest unit of time is seconds, and the system’s bandwidth is assumed to be 5 (the maximum time slice
consumption is 5).
```text
task_index 0 1 2 ……
task_time_slice 2 sec 4 sec 100sec ……

Task 0 needs 2 seconds to complete, and task 1 needs 4 seconds to complete
```

## Questions
Design a task scheduling algorithm that outputs a ternary table of time series, task indexes, and remaining
execution times of tasks and describes your algorithm’s thought

| time | index | remain_time_slice |
|------|-------|-------------------|
| 0    | 0,1   | 0, 1 (2-2,4-3) |
| 1    | 1,2   | 0,96 (1-1, 100-4) |
| 2    | 2     | 91 (96-5) |
| ...  | ...   | ...               |


## Advanced Questions

Tasks longer than 5 seconds (including 5 seconds) are long. If we assume that the shorter the task’s running
time is, the higher the priority, what are the algorithm and the schedukubling sequence at this time?

##### Please use any one of Golang/Java/C++/Rust to complete the program. It is required to run and have the correct output.