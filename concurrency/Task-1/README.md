You are working on a system that fetches data from multiple remote APIs in parallel. The data from each API is processed and combined into a final result. To speed up the process, you need to implement a concurrent solution in Go using goroutines and channels.

The system needs to support the following APIs:

API A: Returns a list of numbers.
API B: Returns a list of words.
API C: Returns a list of boolean values.
Your task is to implement a function called `FetchData` that concurrently calls three other functions, each simulating an API request: `FetchNumbers`, `FetchWords`, and `FetchBooleans`. These functions should simulate the latency of an API request using time.Sleep.

The FetchData function should return a struct containing the combined results from all three APIs:

```go
type CombinedResult struct {
	Numbers  []int
	Words    []string
	Booleans []bool
}
```

Additionally, implement a timeout feature in the `FetchData` function. If any of the API requests take longer than the specified timeout, the function should return an error indicating which API request timed out.

Remember to test your solution with different inputs and provide test cases for your implementation. At the end of the interview, I will ask for your solution.