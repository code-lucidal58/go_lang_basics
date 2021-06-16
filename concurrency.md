# Concurrency and Go routines

## Concurrency
Concurrency is about executing a multiple code snippets in a context switching manner. Go, by default, executes
code sequentially. Concurrency can be handy when there are multiple conde snippets that are independent (or sort of)
from each other, calls that are blocking. Executing them concurrently can save up time. Usually threads or processes are used to achieve this behaviour.
In Go, goroutines are used to achieve this.

## Go routines
Go routines are green threads that internally take care of resource management with the kernel of the system.
Placing `go` keyword before the function call, starts the function execution is a new go routine.
It is important to understand that concurrency is different from parallelism. At a time one code line is executed. 
There can be a delay in getting the output of that line (blocking call). At that point, some other line is also triggered. 
**Go Scheduler** by default run on only one CPU. It monitors the code running each routine.  
Multiple cores can be used by explicitly mentioning it. In this case, parallel execution of code takes place.
This triggers a huge debate about concurrency and parallelism.  
```go
for k, v := range my_map {
	go call_my_function(key, value)
}
```

When a program is executed, the first go routine is triggered. This is called the **main routine**. 
The routines spawned by user are called **child routines**.
Main routines controls the time when the program exits. It does NOT take into account the half executed child routines.
This means that, once the code inside `main` function is complete, the program exits.

## Channels
Channels are used to establish communication between routines, main and child, both. It can be used to share messages
of only one datatype. This data type is declared while declaring the channel. 
No other data type values can be placed inside the channel. 
Below is the code for declaring a channel for string message.

```go
c := make(chan string)
```

At a time, only one string can be placed into the channel. 
To put a message into a channel, the syntax is `my_channel <- "Hello`
To receive a message from channel , the syntax is `my_msg := <- my_channel`

