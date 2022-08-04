## Example

```go
// Create limiter with a rate of 100 per second.
limiter := rate.NewLimiter(100, time.Second)
for _ = range <-someChannel {
	limiter.Take()
}
limiter.Stop()
```