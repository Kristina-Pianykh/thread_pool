# Thread Pool in Go

## Usage

```go
pool := CreateThreadPool(runtime.NumCPU())
go pool.Execute(Task(func() {
    myFunc(param1, param2, ...)
}))


// [Optional] wait for tasks to be done
pool.WaitTasks()

// Shut down the pool and clean up workers
pool.Done()

```

## Example

```sh
go run .

# Shut down by sending an interrupt signal (Ctrl+C)
```
