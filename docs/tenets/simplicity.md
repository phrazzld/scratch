---
id: simplicity
last_modified: "2025-05-02"
---

# Tenet: Simplicity Above All

Prefer the simplest design that works. Complexity is the enemy of reliability, maintainability, and understanding. Always seek the most straightforward solution that correctly meets requirements while actively resisting unnecessary complexity.

## Core Belief

Simplicity is not just a goal but a fundamental requirement for building high-quality, maintainable software. Simple code is easier to understand, debug, test, modify, and maintain. Complexity is the primary source of bugs, friction, and long-term costs.

## Practical Guidelines

1. **Apply YAGNI Rigorously**: "You Ain't Gonna Need It" - Don't add functionality or abstractions until they're demonstrably needed.

2. **Minimize Moving Parts**: Each component, abstraction, configuration option, or dependency introduces complexity. Question each one critically.

3. **Value Readability**: Code is read far more often than it's written. Optimize for human understanding over brevity or cleverness.

4. **Distinguish Complexity Types**: 
   - Essential complexity (inherent in the problem domain) must be managed.
   - Accidental complexity (introduced by implementation) must be ruthlessly eliminated.

5. **Refactor Towards Simplicity**: Regularly evaluate and simplify existing code. A growing codebase naturally tends toward complexity unless actively counteracted.

## Warning Signs

- Over-engineering
- Designing for imagined future requirements
- Premature abstraction
- Overly clever or obscure code
- Deep nesting (> 2-3 levels)
- Excessively long functions/methods
- Components violating the Single Responsibility Principle
- "I'll make it generic so we can reuse it later"

## Related Tenets

- [Modularity](/tenets/modularity.md): Breaking systems into focused, small components
- [Explicit is Better than Implicit](/tenets/explicit-over-implicit.md): Clarity through explicitness