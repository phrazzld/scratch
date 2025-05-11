---
id: testability
last_modified: "2025-05-02"
---

# Tenet: Design for Testability

Testability is a fundamental, non-negotiable design constraint that must be considered from the start. Difficulty testing code is a strong signal to refactor the design before proceeding.

## Core Belief

Automated tests build confidence, prevent regressions, enable safe refactoring, and act as executable documentation. Code difficult to test often indicates poor design (high coupling, mixed concerns). Tests should verify behavior, not implementation.

## Practical Guidelines

1. **Test-Driven Development**: Consider writing tests before implementation where appropriate, as this naturally leads to more testable designs.

2. **Structure for Testability**: Organize code with clear interfaces, dependency inversion, separation of concerns, and pure functions to enable easy verification.

3. **Test Behavior Not Implementation**: Focus tests on *what* (public API, behavior), not *how* (internal implementation).

4. **Refactor First**: If code is difficult to test, this is a signal to refactor the code under test rather than creating complex test setups.

5. **No Mocking Internal Components**: Mocking should only be used at true external system boundaries. The need to mock internal collaborators indicates a design problem.

## Warning Signs

- Test setup requires complex mocking
- Tests are brittle, breaking with minor implementation changes
- Tests require knowledge of internal implementation details
- Tests are difficult to write or understand
- Low test coverage in core business logic

## Related Tenets

- [Modularity](/tenets/modularity.md): Small, focused components are easier to test
- [Explicitness](/tenets/explicit-over-implicit.md): Clear dependencies improve testability