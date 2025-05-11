---
id: immutable-by-default
last_modified: "2025-05-03"
derived_from: simplicity
enforced_by: linters & code review
applies_to:
  - all
  - typescript
  - javascript
---

# Binding: Immutable by Default

Data structures must be treated as immutable by default. Creating new instances instead of modifying existing data in place is required. Direct mutation of objects, arrays, or other data structures is only permitted with explicit justification (e.g., measured critical performance need).

## Rationale

Immutable data structures simplify reasoning about state, eliminate entire classes of bugs from shared mutable state, and make changes predictable and traceable. Immutability supports referential transparency, which is essential for building pure functions and makes code easier to test and debug. By making immutability the default, we reduce cognitive load and make code behavior more predictable.

## Enforcement

This binding is enforced by:

1. Language-specific linters (e.g., ESLint rules like `prefer-const`, `no-param-reassign`)
2. Code review requiring immutable update patterns
3. Use of immutability libraries or language features where available

## Implementation by Language

### TypeScript/JavaScript

- Use `const` for variable declarations by default
- Mark properties as `readonly` where appropriate
- Use spread syntax for object/array updates
- Use functional methods (`map`, `filter`, `reduce`) over mutation methods
- Consider libraries like Immer for complex state updates

### Go

- Pass values instead of pointers for immutable data
- Create new instances rather than updating in place
- Use constructor functions to create complete instances

### Rust

- Leverage Rust's ownership system and immutable by default semantics
- Use `mut` keyword only when necessary

## Examples

```typescript
// ❌ BAD: Mutating objects
function addItem(cart, item) {
  cart.items.push(item);       // Mutates the cart
  cart.total += item.price;    // Mutates the cart
  return cart;
}

// ✅ GOOD: Creating new instances
function addItem(cart, item) {
  return {
    ...cart,
    items: [...cart.items, item],
    total: cart.total + item.price
  };
}

// ❌ BAD: Mutating parameters
function sortItems(items) {
  return items.sort((a, b) => a.id - b.id);  // Mutates the original array
}

// ✅ GOOD: Creating new sorted array
function sortItems(items) {
  return [...items].sort((a, b) => a.id - b.id);  // Returns new sorted array
}
```

```go
// ❌ BAD: Mutating the receiver
func (u *User) UpdateName(name string) {
  u.Name = name  // Mutates the original User
}

// ✅ GOOD: Creating a new instance
func (u User) WithName(name string) User {
  newUser := u    // Create a copy
  newUser.Name = name
  return newUser  // Return new instance
}
```

## Exceptions

In rare cases where performance is critical and has been measured to be a bottleneck, controlled mutation may be acceptable. Such cases must be:

1. Clearly documented with comments explaining the justification
2. Localized to the smallest possible scope
3. Protected from external mutation where possible
4. Verified with performance benchmarks

## Related Bindings

- [function-purity.md](./function-purity.md) - Pure functions work well with immutable data
- [no-side-effects.md](./no-side-effects.md) - Avoiding side effects complements immutability