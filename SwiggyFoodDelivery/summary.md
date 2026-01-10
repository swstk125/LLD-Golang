# Swiggy Food Delivery System - Design Summary

## Overview
This is a low-level design implementation of a food delivery system similar to Swiggy/Zomato. The system handles the complete order lifecycle from restaurant selection to order delivery.

## Core Components

### 1. CustomerService
- **Purpose**: Manages customer information
- **Key Features**:
  - Stores customer data (ID, address)
  - Provides customer lookup functionality
- **Data Structure**: Map of customer ID to Customer object

### 2. RestaurantService
- **Purpose**: Manages restaurants and their menus
- **Key Features**:
  - Create restaurants
  - Add menu items to restaurants
  - Retrieve restaurant list
  - Get menu items for a specific restaurant
- **Data Structure**: List of Restaurant objects, each containing menu items

### 3. CartService
- **Purpose**: Manages shopping cart functionality
- **Key Features**:
  - Add items to cart
  - Auto-create cart if doesn't exist
  - Calculate total cart amount
  - Retrieve user's cart
- **Data Structure**: Map of customer ID to Cart object

### 4. OrderService
- **Purpose**: Manages order creation and lifecycle
- **Key Features**:
  - Create orders from cart
  - Place orders (moves to PAYMENT_PENDING state)
  - Assign delivery agents to orders
  - Track order status through various states
- **Order States**: CREATED → PAYMENT_PENDING → CONFIRMED → PREPARING → ON_THE_WAY → DELIVERED/CANCELLED
- **Data Structure**: Map of order ID to Order object

### 5. InventoryService
- **Purpose**: Manages restaurant inventory
- **Key Features**:
  - Track inventory items per restaurant
  - Reserve items when orders are placed
  - Maintain total quantity and reserved quantity
- **Data Structure**: Map of restaurant ID to Inventory, which contains map of item ID to InventoryItem

### 6. AgentService
- **Purpose**: Manages delivery agents
- **Key Features**:
  - Track agent status (AVAILABLE, UNAVAILABLE, ON_DELIVERY)
  - Agent assignment using strategy pattern
  - Nearest agent assignment strategy (implemented)
- **Design Pattern**: Strategy Pattern for agent assignment

### 7. PaymentService
- **Status**: Placeholder (not implemented)

## Design Patterns Used

1. **Strategy Pattern**: Used in AgentService for different agent assignment strategies
2. **Service Layer Pattern**: Each domain has its own service class
3. **State Pattern**: Order status transitions represent state changes

## Data Flow

1. Customer browses restaurants → RestaurantService
2. Customer adds items to cart → CartService
3. Customer places order → OrderService creates order
4. Order moves to PAYMENT_PENDING → Payment processing (not implemented)
5. Inventory reservation → InventoryService
6. Agent assignment → AgentService
7. Order tracking through various states → OrderService

## Key Design Decisions

- **Separation of Concerns**: Each service handles a specific domain
- **In-memory Storage**: Uses maps for simplicity (can be replaced with database)
- **Status-based Order Management**: Clear state transitions for order lifecycle
- **Extensible Agent Assignment**: Strategy pattern allows easy addition of new assignment algorithms

---

## Suggestions and Important Points for Low-Level Design Interview

### 1. Clarify Requirements First
- Ask about scale (number of users, orders per day)
- Understand functional requirements (order cancellation, refunds, ratings)
- Clarify non-functional requirements (consistency, availability, latency)

### 2. Design Patterns to Consider
- **State Pattern**: For order status transitions (already partially implemented)
- **Observer Pattern**: For order status updates to customers/restaurants
- **Factory Pattern**: For creating different types of orders (scheduled, instant)
- **Repository Pattern**: For data access abstraction
- **Singleton Pattern**: For service instances (if needed)

### 3. Scalability Considerations
- **Database Design**: 
  - Use relational DB (PostgreSQL) for transactional data (orders, payments)
  - Use NoSQL (MongoDB/Cassandra) for high-volume reads (menu, restaurant data)
  - Consider read replicas for high read traffic
- **Caching Strategy**:
  - Cache restaurant menus (Redis)
  - Cache user carts (Redis with TTL)
  - Cache frequently accessed restaurant data
- **Message Queue**: 
  - Use Kafka/RabbitMQ for async processing (order placement, notifications)
  - Decouple services for better scalability

### 4. Consistency and Reliability
- **ACID Transactions**: For order placement (cart → order → inventory → payment)
- **Idempotency**: Ensure order creation is idempotent (use unique order IDs)
- **Eventual Consistency**: For non-critical updates (restaurant ratings, reviews)
- **Saga Pattern**: For distributed transactions across services

### 5. Concurrency Handling
- **Optimistic Locking**: For inventory updates (version numbers)
- **Pessimistic Locking**: For critical operations (payment processing)
- **Distributed Locks**: Use Redis/Zookeeper for cross-service locks
- **Race Conditions**: Handle concurrent cart updates, inventory reservations

### 6. API Design
- **RESTful APIs**: Follow REST principles
- **Idempotent Operations**: Use idempotency keys for POST requests
- **Pagination**: For listing restaurants, orders, menu items
- **Rate Limiting**: Protect APIs from abuse
- **Versioning**: API versioning strategy (v1, v2)

### 7. Error Handling
- **Retry Logic**: For transient failures (network issues)
- **Circuit Breaker**: Prevent cascade failures
- **Dead Letter Queue**: For failed messages
- **Graceful Degradation**: Fallback mechanisms (e.g., if payment service is down)

### 8. Monitoring and Observability
- **Logging**: Structured logging for all operations
- **Metrics**: Track order success rate, latency, error rates
- **Distributed Tracing**: Track requests across services
- **Alerts**: Set up alerts for critical failures

### 9. Security Considerations
- **Authentication**: JWT tokens, OAuth
- **Authorization**: Role-based access control (customer, restaurant, admin)
- **Data Encryption**: Encrypt sensitive data (payment info, addresses)
- **Input Validation**: Sanitize all inputs
- **Rate Limiting**: Prevent abuse and DDoS attacks

### 10. Additional Features to Discuss
- **Order Scheduling**: Allow customers to schedule orders
- **Split Orders**: Multiple restaurants in one order
- **Loyalty Program**: Points and rewards system
- **Recommendations**: ML-based restaurant/item recommendations
- **Real-time Tracking**: WebSocket for live order tracking
- **Multi-language Support**: Internationalization
- **A/B Testing**: For feature rollouts

### 11. Database Schema Considerations
- **Normalization**: Balance between normalization and performance
- **Indexing**: Index frequently queried fields (user_id, restaurant_id, order_status)
- **Partitioning**: Partition orders table by date/region
- **Archiving**: Archive old orders to cold storage

### 12. Testing Strategy
- **Unit Tests**: Test individual service methods
- **Integration Tests**: Test service interactions
- **Load Testing**: Test system under load
- **Chaos Engineering**: Test failure scenarios

### 13. Interview Tips
- **Think Aloud**: Explain your thought process
- **Start Simple**: Begin with basic design, then add complexity
- **Ask Questions**: Clarify ambiguities before designing
- **Consider Trade-offs**: Discuss pros/cons of design decisions
- **Draw Diagrams**: Use UML diagrams (class, sequence, state diagrams)
- **Code Quality**: Write clean, maintainable code with proper naming
- **Handle Edge Cases**: Discuss error scenarios and edge cases

### 14. Common Questions to Prepare For
- How would you handle a surge in orders (e.g., during festivals)?
- How would you ensure inventory consistency across multiple orders?
- How would you implement real-time order tracking?
- How would you handle partial order cancellations?
- How would you design the payment system to be fault-tolerant?
- How would you scale this system to handle 10x traffic?

