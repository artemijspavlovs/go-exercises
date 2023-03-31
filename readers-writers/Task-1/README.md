**Task Description:**

Imagine you are working on a log monitoring system for a large application. The application generates logs in a specific format and saves them to a file. Your task is to implement a log parser in Go that reads the log file, processes the log entries, and writes the processed logs to a new file in a different format.

The original log file format is as follows:

```yaml
[timestamp] [log level] [module]: [message]
```

Example:

```yaml
2023-03-30T10:00:00Z INFO main: Application started
2023-03-30T10:01:01Z ERROR database: Connection failed
```

The desired output format is:

```yaml
[log level] [module] [timestamp]: [message]
```

Example:

```yaml
INFO main 2023-03-30T10:00:00Z: Application started
ERROR database 2023-03-30T10:01:01Z: Connection failed
```

Your program should implement a LogParser that reads the original log file line by line, processes each line, and writes the processed lines to a new file. The program should utilize io.Reader and io.Writer interfaces for reading and writing files. The LogParser should also handle edge cases, such as invalid log entries or incorrect file paths.

Remember to test your solution with different inputs and provide test cases for your implementation. At the end of the interview, I will ask for your solution.
