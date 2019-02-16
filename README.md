# Journal

A logging library with standard output with JSON format written in Go.

## How to use

You can print info log by doing:

```go
journal.Info("Starting application at port: 8080").Log()
```

You can use any other log level like warning or error:

```go
journal.Error("Failed to unmarshal JSON.", err).Log()
journal.Warning("Cannot connect to MYSQL").Log()
```

Set tags for your log:

```go
journal.Info("Starting application at port: 8080").
  SetTags("server", "boot").
  Log()
```

or add custom field:

```go
journal.Info("Starting application at port: 8080").
  SetTags("server", "boot").
  AddField("env", "production")
  Log()
```
