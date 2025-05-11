---
id: no-internal-mocking
last_modified: "2025-05-02"
derived_from: testability
enforced_by: code review & linters
applies_to:
  - all
---

# Binding: No Mocking Internal Components

When writing tests, mocking or stubbing internal classes, structs, functions, or interfaces defined within the same application/service is STRICTLY FORBIDDEN. Only mock true external system boundaries.

## Rationale

Mocking internal collaborators leads to brittle tests that break with refactoring, even when the actual behavior remains unchanged. It creates tight coupling between tests and implementation details rather than testing public behavior. The need to mock internal components is a strong signal of design issues (high coupling, poor separation of concerns) that should be addressed through refactoring.

## Enforcement

This binding is enforced by:

1. Code review processes that reject tests with internal mocking
2. Test-specific linters where available
3. Regular test quality audits

## Acceptable Mocking Boundaries

Mocking is permissible ONLY for interfaces/abstractions representing genuinely external dependencies:

- Network I/O (HTTP clients, gRPC clients)
- Databases
- Filesystem
- System Clock
- External message brokers/queues
- Hardware interfaces

## Guidelines

1. **Abstract External Dependencies**: Always access external dependencies via interfaces defined within your codebase (applying the Ports & Adapters pattern). Mock these local abstractions.

2. **Refactor If Needed**: If you feel the need to mock internal collaborators, this is a signal to refactor the code under test. Consider:
   - Extracting pure functions
   - Introducing proper interfaces
   - Using dependency injection
   - Breaking down overly complex components

3. **Use Real Implementations**: For internal components, always use real implementations in tests instead of mocks.

## Examples

```typescript
// ❌ BAD: Mocking an internal service
// OrderProcessor and InventoryService are both internal components
it("should process an order", () => {
  // Mocking an internal service
  const mockInventoryService = { checkAvailability: jest.fn().mockReturnValue(true) };
  const processor = new OrderProcessor(mockInventoryService);
  
  processor.process(sampleOrder);
  
  expect(mockInventoryService.checkAvailability).toHaveBeenCalled();
});

// ✅ GOOD: Only mocking external dependencies
it("should process an order", () => {
  // Using real internal components
  const inventoryService = new InventoryService();
  const processor = new OrderProcessor(inventoryService);
  
  // Only mock the database (external dependency)
  const dbClient = new TestDatabaseClient(); // Test double implementing DatabaseClient interface
  
  processor.process(sampleOrder);
  
  // Verify behavior through observable results
  expect(dbClient.getSavedOrders()).toContain(sampleOrder);
});
```

## Related Bindings

- [hex-domain-purity.md](./hex-domain-purity.md) - Properly structuring code improves testability
- [dependency-inversion.md](./dependency-inversion.md) - Using DI to facilitate testing