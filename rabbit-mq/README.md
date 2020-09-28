## Install

```
brew update
```

```
brew install rabbitmq
```

## Run rabbitmq

```
export PATH=$PATH:/usr/local/sbin
```

```
brew services start rabbitmq
```

## Check queue
```
rabbitmqctl list_queues 
```

```
rabbitmqctl list_queues name messages_ready messages_unacknowledged
```

```
rabbitmqctl list_bindings
```

# Routing
If you want to save only 'warning' and 'error' (and not 'info') log messages to a file, just open a console and type:
```go run receive_logs_direct.go warning error > logs_from_rabbit.log```

If you'd like to see all the log messages on your screen, open a new terminal and do:
```
go run receive_logs_direct.go info warning error
# => [*] Waiting for logs. To exit press CTRL+C
```

And, for example, to emit an error log message just type:
```
go run emit_log_direct.go error "Run. Run. Or it will explode."
# => [x] Sent 'error':'Run. Run. Or it will explode.'
```

# Topic
To receive all the logs:
```go run receive_logs_topic.go "#"```

To receive all logs from the facility "kern":
```go run receive_logs_topic.go "kern.*"```

Or if you want to hear only about "critical" logs:
```go run receive_logs_topic.go "*.critical"```

You can create multiple bindings:
```go run receive_logs_topic.go "kern.*" "*.critical"```

And to emit a log with a routing key "kern.critical" type:
```go run emit_log_topic.go "kern.critical" "A critical kernel error"```
