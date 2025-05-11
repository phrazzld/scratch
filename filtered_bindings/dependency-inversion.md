---
id: dependency-inversion
last_modified: "2025-05-03"
derived_from: testability
enforced_by: code review & architecture analysis
applies_to:
  - all
---

# Binding: Enforce Dependency Inversion Principle

High-level modules (business logic) must not depend on low-level modules (infrastructure). Both should depend on abstractions. Abstractions should not depend on details; details should depend on abstractions. Source code dependencies must point inward toward the core domain.

## Rationale

The Dependency Inversion Principle (DIP) is fundamental to creating maintainable, testable systems. By ensuring that business logic depends only on abstractions—not concrete implementations—we achieve loose coupling and the ability to replace infrastructure components without modifying core code. This allows the core domain to remain pure and focused on business problems while infrastructure details can evolve independently.

## Enforcement

This binding is enforced by:

1. Code reviews that verify dependency direction
2. Architecture analysis tools that detect violations of the dependency hierarchy
3. Package/module organization that physically separates domain from infrastructure

## Implementation

1. **Define Abstractions in Core**: Interfaces and abstract types should be defined in the core/domain layer.

2. **Implement Adapters in Infrastructure**: Concrete implementations of these interfaces should live in the infrastructure layer.

3. **Use Dependency Injection**: Provide implementations to the core layer at runtime, typically via constructor injection.

4. **Check Import Direction**: Ensure that infrastructure code imports domain code, not vice versa.

## Examples

```typescript
// ❌ BAD: Domain depends on infrastructure
// Domain code importing infrastructure
import { MongoDBUserRepository } from '../infrastructure/mongodb/MongoDBUserRepository';

class UserService {
  private repository: MongoDBUserRepository;
  
  constructor() {
    this.repository = new MongoDBUserRepository();  // Direct dependency on concrete implementation
  }
  
  async getUser(id: string) {
    return this.repository.findById(id);
  }
}

// ✅ GOOD: Infrastructure depends on domain
// Domain code defines the interfaces it needs
export interface UserRepository {
  findById(id: string): Promise<User | null>;
  save(user: User): Promise<void>;
}

export class UserService {
  constructor(private repository: UserRepository) {}  // Depends on abstraction
  
  async getUser(id: string) {
    return this.repository.findById(id);
  }
}

// Infrastructure code implements the domain interfaces
import { UserRepository, User } from '../../domain/UserService';

export class MongoDBUserRepository implements UserRepository {
  async findById(id: string): Promise<User | null> {
    // MongoDB specific implementation
  }
  
  async save(user: User): Promise<void> {
    // MongoDB specific implementation
  }
}
```

## Related Bindings

- [hex-domain-purity.md](./hex-domain-purity.md) - Keep domain logic free from infrastructure details
- [no-internal-mocking.md](./no-internal-mocking.md) - Proper dependency inversion facilitates testing