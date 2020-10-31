# Go Concurrent Retry

Go Concurrent Retry is a project that implements goroutines to execute a large list of tasks quickly. 

Spawns a designated amount of worker goroutines that distribute the list of tasks based on their availability, if the task failed the worker will retry upto a designated amount.

## Requirements

Go version 1.15.x or higher

For more information about installing Go, visit [golang.org/doc/install](https://golang.org/doc/install)

>Note: May work with older versions but is not guaranteed to function properly.

## Installation
```bash
git clone ...
```

## Usage
An example using the ConcurrentRetry function
```go
  // Create a list of tasks
  var tasks []func() (string, error)
  for i := 0; i < 20; i++ {
    tasks = append(tasks, coding)
  }

  workers := 3
  retries := 5

  // Returns a channel where the
  // results can be collected from
  results := ConcurrentRetry(tasks, workers, retries)

  for r := range results {
    fmt.Println("Result received from goroutine:", r)
  }
```

## Contributing
Pull requests are welcome.

## License
[MIT]()