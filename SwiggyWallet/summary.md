# Swiggy Wallet System - Design Summary

## Overview
This is a low-level design implementation of a digital wallet system similar to Swiggy Wallet/Paytm. The system handles wallet operations including adding money, making payments, transaction history, and refunds.

## Core Components

### 1. WalletManager
- **Purpose**: Manages user wallet balances
- **Key Features**:
  - Create wallet for new users (lazy initialization)
  - Add money to wallet
  - Deduct money from wallet
  - Get wallet balance
- **Data Structure**: Map of user ID to Wallet object
- **Wallet Properties**: ID, balance, user ID

### 2. TransactionManager
- **Purpose**: Manages all financial transactions
- **Key Features**:
  - Create transactions (credit, debit, refund)
  - Store transaction history per user
  - Retrieve all transactions for a user
  - Get specific transaction by ID
- **Transaction Types**: CREDIT, DEBIT, REFUND
- **Transaction Status**: Success, Failed
- **Data Structure**: Map of user ID to list of transactions
- **Transaction Properties**: ID, user ID, amount, status, type, timestamp

### 3. RefundManager
- **Purpose**: Handles refund processing
- **Key Features**:
  - Validate refund requests
  - Process refunds by creating credit transactions
  - Interacts with TransactionManager to retrieve original transaction
- **Design Pattern**: Uses dependency injection (receives TransactionManager)

### 4. Swiggy Service (Main Orchestrator)
- **Purpose**: Main service that coordinates wallet operations
- **Key Features**:
  - Provides high-level API for wallet operations
  - Coordinates between WalletManager, TransactionManager, and RefundManager
- **Operations**:
  - `addMoneyToWallet`: Adds money and creates credit transaction
  - `makePayment`: Deducts money and creates debit transaction
  - `getTransactionHistory`: Retrieves user's transaction history
  - `requestRefund`: Processes refund requests

## Design Patterns Used

1. **Service Layer Pattern**: Separation of concerns with dedicated managers
2. **Dependency Injection**: RefundManager receives TransactionManager as dependency
3. **Lazy Initialization**: Wallets are created on-demand when first accessed

## Data Flow

1. **Add Money Flow**:
   - User requests to add money → WalletManager.Add() → Update balance → TransactionManager.Add() (CREDIT)

2. **Payment Flow**:
   - User makes payment → WalletManager.Deduct() → Update balance → TransactionManager.Add() (DEBIT)

3. **Refund Flow**:
   - User requests refund → RefundManager.ValidateRefund() → Retrieve original transaction → Create credit transaction → Update wallet balance

## Key Design Decisions

- **Separation of Concerns**: Wallet operations, transactions, and refunds are separate concerns
- **Transaction Logging**: All operations are logged as transactions for audit trail
- **In-memory Storage**: Uses maps for simplicity (can be replaced with database)
- **Atomic Operations**: Wallet updates and transaction creation should be atomic (not fully implemented in current design)

---

## Suggestions and Important Points for Low-Level Design Interview

### 1. Clarify Requirements First
- Ask about transaction limits (min/max amounts)
- Understand refund policies (time limits, conditions)
- Clarify wallet balance limits
- Ask about multi-currency support
- Understand KYC requirements

### 2. Design Patterns to Consider
- **Repository Pattern**: Abstract data access layer
- **Factory Pattern**: For creating different transaction types
- **Observer Pattern**: For transaction notifications (SMS, email, push)
- **Strategy Pattern**: For different refund policies
- **Command Pattern**: For transaction operations (undo/redo capability)

### 3. Scalability Considerations
- **Database Design**:
  - Use ACID-compliant database (PostgreSQL) for financial data
  - Separate read and write databases (CQRS pattern)
  - Partition transaction table by user_id or date
  - Use time-series database for analytics
- **Caching Strategy**:
  - Cache wallet balances (Redis) with TTL
  - Cache recent transactions
  - Invalidate cache on balance updates
- **Message Queue**:
  - Use Kafka for transaction events
  - Async processing for notifications
  - Event sourcing for audit trail

### 4. Consistency and Reliability (CRITICAL)
- **ACID Transactions**: 
  - Use database transactions for wallet update + transaction log
  - Ensure atomicity: balance update and transaction creation must succeed/fail together
- **Double-Entry Bookkeeping**: 
  - Consider implementing double-entry accounting
  - Every debit has corresponding credit
- **Idempotency**:
  - Use idempotency keys for all transactions
  - Prevent duplicate transactions
- **Distributed Transactions**:
  - Use Saga pattern for cross-service transactions
  - Two-phase commit for critical operations
- **Eventual Consistency**:
  - For non-critical updates (notifications, analytics)

### 5. Concurrency Handling (CRITICAL)
- **Optimistic Locking**: 
  - Use version numbers for wallet balance
  - Retry on version conflicts
- **Pessimistic Locking**:
  - Use row-level locks for critical operations
  - Lock user's wallet during balance updates
- **Distributed Locks**:
  - Use Redis/Zookeeper for cross-service locks
  - Prevent concurrent balance updates
- **Race Conditions**:
  - Handle concurrent payment requests
  - Prevent negative balance (check before deduct)

### 6. Security Considerations (CRITICAL)
- **Encryption**:
  - Encrypt sensitive data at rest
  - Use TLS for data in transit
- **Authentication & Authorization**:
  - Strong authentication (2FA, biometric)
  - Verify user owns wallet before operations
- **Fraud Detection**:
  - Rate limiting on transactions
  - Anomaly detection (unusual patterns)
  - Transaction limits per day/hour
- **Audit Trail**:
  - Immutable transaction logs
  - Log all operations with timestamps
  - Compliance with financial regulations

### 7. API Design
- **RESTful APIs**: Follow REST principles
- **Idempotent Operations**: All POST requests should be idempotent
- **Idempotency Keys**: Required for payment/refund operations
- **Pagination**: For transaction history
- **Rate Limiting**: Strict limits on financial operations
- **Versioning**: API versioning for backward compatibility

### 8. Error Handling
- **Retry Logic**: 
  - Exponential backoff for transient failures
  - Idempotent retries
- **Circuit Breaker**: 
  - Prevent cascade failures
  - Fallback mechanisms
- **Transaction Rollback**:
  - Rollback on failures
  - Compensating transactions for refunds
- **Dead Letter Queue**: For failed transaction processing

### 9. Monitoring and Observability
- **Logging**:
  - Structured logging for all transactions
  - Log balance changes with before/after values
- **Metrics**:
  - Transaction success/failure rates
  - Average transaction amount
  - Wallet balance distribution
  - Refund rate
- **Alerts**:
  - Failed transactions
  - Unusual transaction patterns
  - System downtime
- **Distributed Tracing**: Track transactions across services

### 10. Additional Features to Discuss
- **Wallet to Wallet Transfer**: P2P transfers
- **Scheduled Payments**: Recurring payments
- **Payment Methods**: Integration with UPI, cards, net banking
- **Loyalty Points**: Convert points to wallet balance
- **Cashback**: Cashback on transactions
- **Multi-wallet Support**: Separate wallets for different purposes
- **Wallet Freeze**: Freeze wallet for security
- **Transaction Disputes**: Handle transaction disputes

### 11. Database Schema Considerations
- **Normalization**: 
  - Separate tables for wallets, transactions
  - Foreign key relationships
- **Indexing**:
  - Index on user_id, transaction_id, timestamp
  - Composite indexes for common queries
- **Partitioning**:
  - Partition transaction table by date
  - Archive old transactions
- **Constraints**:
  - Check constraints for balance >= 0
  - Unique constraints on transaction IDs

### 12. Testing Strategy
- **Unit Tests**: 
  - Test wallet operations
  - Test transaction creation
  - Test refund logic
- **Integration Tests**:
  - Test wallet + transaction operations together
  - Test concurrent operations
- **Load Testing**: 
  - Test under high transaction volume
  - Test concurrent balance updates
- **Chaos Engineering**: 
  - Test failure scenarios
  - Test partial failures

### 13. Compliance and Regulations
- **PCI DSS**: If handling card data
- **KYC/AML**: Know Your Customer, Anti-Money Laundering
- **Data Privacy**: GDPR, data retention policies
- **Financial Regulations**: Compliance with local financial laws
- **Audit Requirements**: Maintain audit logs for regulatory compliance

### 14. Interview Tips
- **Emphasize Consistency**: Financial systems require strong consistency
- **Discuss Edge Cases**: 
  - Negative balance prevention
  - Concurrent transactions
  - System failures mid-transaction
- **Think About Scale**: 
  - How to handle millions of transactions per day
  - Database sharding strategies
- **Security First**: Always prioritize security in financial systems
- **Draw Diagrams**: 
  - Sequence diagrams for transaction flows
  - Database schema diagrams
  - System architecture diagrams

### 15. Common Questions to Prepare For
- How would you prevent double-spending?
- How would you handle a transaction failure after deducting balance?
- How would you ensure transaction consistency across services?
- How would you implement wallet freeze/unfreeze?
- How would you handle refunds for failed transactions?
- How would you scale to handle 1 million transactions per day?
- How would you detect fraudulent transactions?
- How would you implement wallet-to-wallet transfers?

### 16. Code Improvements Needed
- **Atomic Operations**: Implement database transactions
- **Balance Validation**: Check balance before deducting
- **Error Handling**: Proper error handling and rollback
- **Idempotency**: Add idempotency keys
- **Validation**: Validate amounts (positive, within limits)
- **Logging**: Add comprehensive logging
- **Thread Safety**: Ensure thread-safe operations

