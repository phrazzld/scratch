---
id: modularity
last_modified: "2025-05-02"
---

# Tenet: Modularity is Mandatory

Construct software from small, well-defined, independent components with clear responsibilities and explicit interfaces. Systems should comprise focused units that do one thing well and compose together.

## Core Belief

Modularity tames complexity, enables parallel development, independent testing/deployment, reuse, fault isolation, and easier evolution. Properly modularized systems are more robust, maintainable, and adaptable to change.

## Practical Guidelines

1. **Do One Thing Well**: Each module, package, service, or function should have a single, clear responsibility.

2. **Define Clear Boundaries**: Modules should have well-defined interfaces and hide their implementation details.

3. **Minimize Coupling**: Reduce dependencies between modules; when dependencies exist, they should be through abstract interfaces rather than concrete implementations.

4. **Maximize Cohesion**: Related functionality should be grouped together within a module.

5. **Design for Composition**: Smaller modules should combine easily to build more complex functionality.

## Warning Signs

- Monolithic components that handle multiple concerns
- "God objects" or classes with too many responsibilities
- Tangled dependencies between modules
- Changes in one area frequently breaking others
- Testing requiring complex setup or mocking
- Difficulty understanding how components fit together

## Related Tenets

- [Simplicity](/tenets/simplicity.md): Smaller, focused modules are easier to understand
- [Testability](/tenets/testability.md): Modular code is inherently more testable