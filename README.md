# Loggan [![Build Status](https://travis-ci.org/naoina/loggan.svg?branch=master)](https://travis-ci.org/naoina/loggan) [![GoDoc](https://godoc.org/github.com/naoina/loggan?status.svg)](https://godoc.org/github.com/naoina/loggan)

Leveled structured logger for [Go](https://golang.org/).

## Installation

    go get -u github.com/naoina/loggan

## Usage

```go
package main

import (
	"os"

	"github.com/naoina/loggan"
)

func main() {
	log := loggan.New(os.Stdout, &loggan.LTSVFormatter{}, loggan.INFO)
	log.Debug("This is a debug log")
	log.Info("This is an info log")
	log.Warn("This is a warning log")
	log.Error("This is an error log")

	log.With(loggan.Fields{
		"code":   500,
		"status": "Internal Server Error",
	}).Error("500 Internal Server Error")
	log.With(loggan.Fields{
		"code":   404,
		"status": "Not Found",
	}).Info()
}
```

```text
level:INFO	time:2016-06-14T08:52:43.712404034+09:00	message:This is an info log
level:WARN	time:2016-06-14T08:52:43.712449902+09:00	message:This is a warning log
level:ERROR	time:2016-06-14T08:52:43.712455164+09:00	message:This is an error log
level:ERROR	time:2016-06-14T08:52:43.712460346+09:00	message:500 Internal Server Error	code:500	status:Internal Server Error
level:INFO	time:2016-06-14T08:52:43.712466459+09:00	code:404	status:Not Found
```

## Logging levels

```go
log.Print("This is a none log")
log.Debug("This is a debug log")
log.Info("This is an info log")
log.Warn("This is a warning log")
log.Error("This is an error log")
log.Fatal("This is a fatal log") // Calls os.Exit(1) after output.
log.Panic("This is a panic log") // Panics after output.
```

You can set the logging level by the two ways.

```go
log := loggan.New(os.Stdout, &loggan.LTSVFormatter{}, loggan.INFO)
```

or

```go
log.SetLevel(loggan.WARN)
```

Also you can specified the logging level for each logging output.

```go
log.Output(loggan.ERROR, "This is an error log")
```

## Formatters

Loggan includes some formatter.

- `loggan.LTSVFormatter`

    Format to [Labeled Tab-separated Values](http://ltsv.org/).

    ```text
    level:INFO	time:2016-06-14T10:09:57.529623894+09:00	message:This is an info log
    ```

- `loggan.JSONFormatter`

    Format to JSON

    ```json
    {"level":"INFO","time":"2016-06-14T21:05:53.497413624+09:00","message":"This is an info log"}
    ```

- `loggan.RawFormatter`

    Output the given message only.

    ```text
    This is an info log
    ```

### Create a custom formatter

Implement the [Formatter](https://godoc.org/github.com/naoina/loggan#Formatter) interface.

```go
type CustomFormatter struct{}

func (f *CustomFormatter) Format(w io.Writer, entry *loggan.Entry) error {
	_, err := fmt.Fprintf(w, "[%s] [%s] %s", entry.Level, entry.Time.Format(time.RFC3339), entry.Message)
	return err
}
```

Then pass it to `loggan.New`.

```go
log := loggan.New(os.Stdout, &CustomFormatter{}, loggan.INFO)
log.Info("This is an info log")
```

See also https://github.com/naoina/loggan/blob/master/formatter.go to get some examples.

## License

Loggan is licensed under the MIT License.
