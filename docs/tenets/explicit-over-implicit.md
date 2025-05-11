---
id: explicit-over-implicit
last_modified: "2025-05-02"
---

# Tenet: Explicit is Better than Implicit

Make dependencies, data flow, control flow, contracts, and side effects clear and obvious. Clarity trumps magic - prefer explicitness even when it requires slightly more code or effort.

## Core Belief

Explicit code is easier to understand, reason about, debug, and refactor safely. Implicit behavior obscures dependencies, hinders tracing, and leads to unexpected side effects. When code is explicit, its intent and behavior are transparent.

## Practical Guidelines

1. **Make Dependencies Visible**: Use explicit dependency injection rather than global state, singletons, or implicit context.

2. **Document Contracts**: Clearly define what functions accept, return, and the preconditions they require.

3. **Signal Side Effects**: Make it obvious when functions mutate state or have effects beyond their return value.

4. **Name Meaningfully**: Use descriptive names that reveal intent and behavior.

5. **Use Strong Typing**: Leverage type systems to make constraints and valid operations clear.

## Warning Signs

- Magic strings or numbers
- Global state or hidden singletons
- Implicit coupling between components
- Undocumented assumptions about usage
- Complex inheritances or mixins that make behavior difficult to trace
- "Clever" code that hides what's actually happening

## Related Tenets

- [Simplicity](/tenets/simplicity.md): Explicitness leads to more understandable code
- [Modularity](/tenets/modularity.md): Clear boundaries and interfaces