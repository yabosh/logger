# logger
Simple wrapper around internal go logger that provides support for levels.

## Usage

This example shows how to configure and use logger.  logger is a simple wrapper around the go "log" module so any changes the logging configuration for "log" will be reflected in log output when using logger.

All logger functions accept a format and parameters that function in the same way that fmt.Println() operates.

```go
message := "this is my message"
logger.SetLevel("DEBUG")
logger.Debug("Sample message: %s", message)
logger.Trace("This message should not show up because of the current level")
```

Sample output from this example would be:
```
2022/07/25 10:13:25 DEBUG - Sample message: this is my message
```
