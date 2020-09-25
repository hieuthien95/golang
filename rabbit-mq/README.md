# Install

```
brew update
```

```
brew install rabbitmq
```

# Run rabbitmq

```
export PATH=$PATH:/usr/local/sbin
```

```
brew services start rabbitmq
```

# Check queue
```
rabbitmqctl list_queues 
```

```
rabbitmqctl list_queues name messages_ready messages_unacknowledged
```