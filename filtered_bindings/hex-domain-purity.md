---
id: hex-domain-purity
last_modified: "2025-05-02"
derived_from: simplicity
enforced_by: import graph analysis & code review
applies_to:
  - all
---

# Binding: Domain Core Must Remain Pure

Core business logic/domain code must be free from infrastructure dependencies. It should not directly import database libraries, HTTP frameworks, file I/O, or third-party services. Domain code should be pure and unaware of specific I/O mechanisms.

## Rationale

Separating core business logic from infrastructure concerns creates a clean design that's easier to test, maintain, and evolve. It allows business logic to be verified in isolation and infrastructure to be changed with minimal impact on core functionality. This architectural principle prevents accidental coupling and ensures that domain logic remains focused on modeling the problem domain, not technical implementation details.

## Enforcement

This binding is enforced by:

1. Code review processes that verify proper separation
2. Import graph analysis to ensure core packages don't import infrastructure
3. Directory/package structure that physically separates domains from infrastructure

## Implementation Pattern

Follow the Hexagonal Architecture (Ports & Adapters) or Clean Architecture pattern:

1. **Core Domain:** Contains business logic and defines interfaces (ports) that it needs to interact with the outside world
2. **Adapters:** Implement the interfaces required by the core domain to connect to databases, APIs, etc.
3. **Dependency Flow:** Always inward - adapters depend on domain, never the reverse

## Examples

```go
// ❌ BAD: Domain directly using infrastructure
package domain

import (
    "database/sql"
    "net/http"
)

type OrderService struct {
    db *sql.DB
}

func (s *OrderService) PlaceOrder(r *http.Request) error {
    // Direct use of infrastructure
}

// ✅ GOOD: Domain defines interfaces, infrastructure implements them
package domain

type OrderRepository interface {
    Save(order Order) error
    FindByID(id string) (Order, error)
}

type OrderService struct {
    repo OrderRepository
}

func (s *OrderService) PlaceOrder(order Order) error {
    // Pure business logic
    if !order.IsValid() {
        return ErrInvalidOrder
    }
    return s.repo.Save(order)
}

// In infrastructure layer:
package postgres

import (
    "database/sql"
    "myapp/domain"
)

type OrderRepositoryImpl struct {
    db *sql.DB
}

func (r *OrderRepositoryImpl) Save(order domain.Order) error {
    // Infrastructure implementation
}
```

## Related Bindings

- [dependency-inversion.md](./dependency-inversion.md) - Use dependency inversion to manage dependencies
- [feature-folders.md](./feature-folders.md) - Organize code by business feature