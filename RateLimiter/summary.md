# Rate Limiter System - Design Summary

## Overview
This is a low-level design implementation of a rate limiting system with multiple algorithms. Rate limiting is crucial for protecting APIs and services from abuse, ensuring fair resource usage, and preventing system overload.

## Implemented Algorithms

### 1. Fixed Window Limiter
- **Algorithm**: Divides time into fixed windows (e.g., 1 minute, 1 hour)
- **How it works**:
  - Each window has a fixed request limit
  - Counter resets at the start of each new window
  - Simple and memory-efficient
- **Pros**:
  - Simple implementation
  - Low memory overhead
  - Fast lookups
- **Cons**:
  - Burst traffic at window boundaries
  - Not smooth rate limiting
  - Can allow 2x limit if requests span window boundary
- **Use Case**: Simple rate limiting where bursts are acceptable

### 2. Sliding Window Log Limiter
- **Algorithm**: Maintains a log of request timestamps
- **How it works**:
  - Stores timestamps of all requests in the current window
  - Removes timestamps outside the window
  - Allows request if count < limit
- **Pros**:
  - Smooth rate limiting
  - Accurate (no burst at boundaries)
  - Fair distribution
- **Cons**:
  - High memory usage (stores all timestamps)
  - Slower cleanup (needs to remove old timestamps)
- **Use Case**: When accuracy is critical and memory is not a concern

### 3. Token Bucket Limiter
- **Algorithm**: Maintains a bucket of tokens that refill at a constant rate
- **How it works**:
  - Bucket has a capacity (max tokens)
  - Tokens are added at a refill rate
  - Request consumes one token
  - Request allowed if tokens available
- **Pros**:
  - Allows bursts (up to bucket capacity)
  - Smooth average rate
  - Memory efficient
- **Cons**:
  - More complex than fixed window
  - Requires periodic refill calculations
- **Use Case**: When bursts are acceptable and smooth rate is desired

### 4. Leaky Bucket Limiter
- **Algorithm**: Requests leak out at a constant rate
- **How it works**:
  - Bucket has a capacity
  - Requests are added to bucket
  - Requests leak out at a constant rate
  - Request allowed if bucket has space
- **Pros**:
  - Smooth output rate (no bursts)
  - Memory efficient
  - Good for traffic shaping
- **Cons**:
  - Can drop requests if bucket is full
  - No burst allowance
- **Use Case**: When smooth, constant output rate is required

## Common Interface

All limiters implement the `RateLimiter` interface:
```go
type RateLimiter interface {
    Allow(string) bool
}
```

## Key Design Decisions

- **Per-key Limiting**: Each identifier (user ID, IP, API key) has its own limit
- **Thread Safety**: Uses mutex locks for concurrent access
- **In-memory Storage**: Uses maps for simplicity (can be replaced with Redis for distributed systems)
- **Lazy Initialization**: Buckets/windows created on first request

---

## Suggestions and Important Points for Low-Level Design Interview

### 1. Clarify Requirements First
- **Rate Limit Scope**: Per user, per IP, per API key, or global?
- **Rate Limit Values**: What are the limits? (e.g., 100 requests/minute)
- **Burst Tolerance**: Are bursts acceptable or should rate be smooth?
- **Accuracy Requirements**: How strict should rate limiting be?
- **Memory Constraints**: How much memory can be used?
- **Distributed System**: Single server or distributed?

### 2. Algorithm Selection Guide

| Algorithm | Best For | Memory | Accuracy | Burst Handling |
|-----------|----------|--------|----------|----------------|
| Fixed Window | Simple use cases | Low | Low | Poor |
| Sliding Log | High accuracy needed | High | High | Good |
| Token Bucket | Burst tolerance needed | Medium | Medium | Excellent |
| Leaky Bucket | Smooth output needed | Medium | Medium | Poor |

### 3. Scalability Considerations

- **Distributed Rate Limiting**:
  - Use Redis for shared state across servers
  - Redis INCR with expiration for fixed window
  - Redis sorted sets for sliding window log
  - Redis Lua scripts for atomic operations
- **Sharding Strategy**:
  - Shard by key (user ID hash) to distribute load
  - Use consistent hashing for key distribution
- **Caching**:
  - Cache rate limit status (allowed/denied)
  - Reduce Redis calls for frequently accessed keys
- **Load Balancing**:
  - Rate limiting should work across multiple servers
  - Consider sticky sessions or shared state

### 4. Redis Implementation Patterns

**Fixed Window in Redis**:
```
INCR rate_limit:user:123
EXPIRE rate_limit:user:123 60
```

**Sliding Window Log in Redis**:
```
ZADD rate_limit:user:123 timestamp request_id
ZREMRANGEBYSCORE rate_limit:user:123 0 (now - window)
ZCARD rate_limit:user:123
```

**Token Bucket in Redis**:
```
MULTI
GET tokens:user:123
SET tokens:user:123 (tokens - 1)
EXEC
```

### 5. Consistency and Reliability

- **Atomic Operations**: 
  - Use Redis transactions (MULTI/EXEC)
  - Use Lua scripts for complex atomic operations
- **Race Conditions**:
  - Handle concurrent requests properly
  - Use distributed locks if needed
- **Failure Handling**:
  - Fail-open vs fail-closed strategy
  - Fallback mechanisms if Redis is down
- **Data Persistence**:
  - Redis persistence (AOF/RDB) for durability
  - Consider replication for high availability

### 6. Performance Optimizations

- **Lazy Cleanup**:
  - Don't clean up expired entries immediately
  - Clean up periodically or on access
- **Approximate Algorithms**:
  - Use probabilistic data structures (HyperLogLog, Count-Min Sketch)
  - Trade accuracy for memory efficiency
- **Batch Operations**:
  - Batch Redis operations when possible
  - Use pipelining for multiple operations
- **Local Caching**:
  - Cache rate limit decisions locally
  - Reduce network calls to Redis

### 7. Advanced Features to Discuss

- **Dynamic Rate Limiting**:
  - Adjust limits based on system load
  - Different limits for different user tiers
- **Rate Limit Headers**:
  - Return X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset
- **Graceful Degradation**:
  - Reduce limits during high load
  - Prioritize certain users/requests
- **Rate Limit Exemptions**:
  - Whitelist certain users/IPs
  - Admin bypass mechanisms
- **Multi-level Rate Limiting**:
  - Global limits + per-user limits
  - Per-endpoint limits + per-user limits

### 8. API Design

- **Response Headers**:
  ```
  X-RateLimit-Limit: 100
  X-RateLimit-Remaining: 95
  X-RateLimit-Reset: 1609459200
  Retry-After: 60
  ```
- **Error Responses**:
  - HTTP 429 (Too Many Requests)
  - Include retry-after header
  - Clear error messages
- **Idempotency**:
  - Rate limit should be idempotent
  - Same request shouldn't count multiple times

### 9. Monitoring and Observability

- **Metrics to Track**:
  - Rate limit hit rate
  - Requests per key distribution
  - Memory usage
  - Redis latency
  - False positives/negatives
- **Logging**:
  - Log rate limit violations
  - Log rate limit decisions
  - Track patterns of abuse
- **Alerts**:
  - High rate limit violation rate
  - Redis connection issues
  - Memory usage thresholds

### 10. Security Considerations

- **Key Spoofing**:
  - Validate and sanitize rate limit keys
  - Prevent key manipulation
- **DDoS Protection**:
  - Rate limiting is first line of defense
  - Combine with other security measures
- **IP-based Limiting**:
  - Handle IP rotation (proxy, VPN)
  - Consider fingerprinting techniques
- **User-based Limiting**:
  - Authenticate users before rate limiting
  - Prevent user ID manipulation

### 11. Testing Strategy

- **Unit Tests**:
  - Test each algorithm independently
  - Test edge cases (boundary conditions)
  - Test concurrent access
- **Integration Tests**:
  - Test with Redis
  - Test distributed scenarios
- **Load Testing**:
  - Test under high request volume
  - Test Redis performance
  - Test memory usage
- **Chaos Testing**:
  - Test Redis failures
  - Test network partitions

### 12. Interview Tips

- **Explain Trade-offs**: 
  - Discuss pros/cons of each algorithm
  - Explain when to use which algorithm
- **Think About Scale**:
  - How to handle millions of requests/second?
  - How to reduce Redis calls?
- **Consider Edge Cases**:
  - What happens at window boundaries?
  - What if Redis is down?
  - How to handle clock skew?
- **Draw Diagrams**:
  - Algorithm flow diagrams
  - System architecture diagrams
  - Redis data structure diagrams

### 13. Common Questions to Prepare For

- How would you implement rate limiting in a distributed system?
- How would you reduce memory usage for sliding window log?
- How would you handle rate limiting when Redis is down?
- How would you implement different rate limits for different users?
- How would you prevent users from bypassing rate limits?
- How would you scale rate limiting to handle 1M requests/second?
- What's the difference between token bucket and leaky bucket?
- How would you implement rate limiting with sub-second precision?

### 14. Code Improvements Needed

- **Error Handling**: Add proper error handling
- **Configuration**: Make limits configurable
- **Metrics**: Add metrics/monitoring
- **Redis Integration**: Add Redis implementation option
- **Cleanup**: Implement periodic cleanup for expired entries
- **Documentation**: Add comprehensive documentation
- **Tests**: Add unit and integration tests

### 15. Production Considerations

- **Fail-Open vs Fail-Closed**:
  - Fail-open: Allow requests if rate limiter is down (better UX)
  - Fail-closed: Deny requests if rate limiter is down (better security)
- **Rate Limit Bypass**:
  - Implement admin/whitelist mechanisms
  - Consider emergency bypass procedures
- **Compliance**:
  - Some APIs have rate limit requirements
  - Document rate limits clearly
- **Cost Optimization**:
  - Minimize Redis calls
  - Use efficient data structures
  - Consider approximate algorithms for scale

