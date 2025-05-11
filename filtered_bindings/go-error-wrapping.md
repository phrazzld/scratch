---
id: go-error-wrapping
last_modified: "2025-05-02"
derived_from: automation
enforced_by: golangci-lint("wrapcheck") & code review
applies_to:
  - go
---

# Binding: Wrap Errors with Context

All errors propagated across package boundaries must be wrapped with contextual information using `fmt.Errorf("context: %w", err)` or a custom error type. This provides critical context for debugging and error handling.

## Rationale

Raw unwrapped errors lack context about what operation failed, making debugging difficult. By systematically wrapping errors with context about the operation that failed, we create more actionable error messages that include the chain of operations leading to the failure, while preserving the ability to check error types.

## Enforcement

This binding is enforced by:

1. The `wrapcheck` linter in golangci-lint
2. Code review requiring proper error wrapping

## Guidelines

1. **Add meaningful context**: Include what operation was attempted and relevant identifiers.
2. **Use the `%w` verb**: Always use `fmt.Errorf("context: %w", err)` to preserve the original error for unwrapping.
3. **Wrap at package boundaries**: Always wrap errors when returning them from exported functions.
4. **Don't wrap internal errors**: Generally avoid wrapping within private functions of the same package.
5. **Use custom error types** when additional structured data is needed.

## Examples

```go
// ❌ BAD: Returning raw error with no context
func ProcessOrder(id string) error {
    order, err := db.GetOrder(id)
    if err != nil {
        return err // No context about what failed
    }
    // ...
}

// ✅ GOOD: Wrapping with context 
func ProcessOrder(id string) error {
    order, err := db.GetOrder(id)
    if err != nil {
        return fmt.Errorf("fetching order %s: %w", id, err)
    }
    // ...
}

// ✅ ALSO GOOD: Custom error type with context
type OrderError struct {
    OrderID string
    Err     error
}

func (e *OrderError) Error() string {
    return fmt.Sprintf("order %s: %v", e.OrderID, e.Err)
}

func (e *OrderError) Unwrap() error {
    return e.Err
}

func ProcessOrder(id string) error {
    order, err := db.GetOrder(id)
    if err != nil {
        return &OrderError{OrderID: id, Err: err}
    }
    // ...
}
```

## Related Bindings

- [go-error-handling.md](./go-error-handling.md) - General error handling patterns
- [go-error-sentinel.md](./go-error-sentinel.md) - When to use sentinel errors