# ab - Apache Benchmark

Simple command-line tool used to load test HTTP/HTTPS servers.

Sends a large number of HTTP requests to a URL and measures:
- How many requests per second the server can handle
- How long responses take (latency)
- How many requests fail
- Overall throughput

## Basic Syntax
```bash
ab -n <total_requests> -c <concurrency> <URL>

# Example:
ab -n 1000 -c 50 http://example.com/
-n 1000: Sends 1000 total requests
-c 50: 50 requests at the same time (concurrent users)
```

