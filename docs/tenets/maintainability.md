---
id: maintainability
last_modified: "2025-05-02"
---

# Tenet: Maintainability Over Premature Optimization

Write code primarily for human understanding and ease of future modification. Clarity, readability, and consistency are paramount - optimize only after identifying actual, measured performance bottlenecks.

## Core Belief

Most development time is spent reading/maintaining existing code, not writing new code. Premature optimization often adds complexity, obscures intent, hinders debugging, and targets non-critical paths, yielding negligible benefit at high maintenance cost.

## Practical Guidelines

1. **Optimize for Readability**: Code should clearly communicate intent to other developers (including your future self).

2. **Establish Consistent Patterns**: Follow consistent conventions and patterns throughout the codebase.

3. **Prioritize Clear Naming**: Invest time in naming variables, functions, and components to accurately reflect their purpose.

4. **Comment the Why, Not the How**: Self-documenting code explains how it works; comments should explain why decisions were made.

5. **Measure Before Optimizing**: Use profiling tools to identify actual performance bottlenecks before optimizing. Focus efforts only on proven hot spots.

## Warning Signs

- Clever tricks to save a few CPU cycles at the expense of readability
- Inconsistent coding styles or patterns within a codebase
- Cryptic variable or function names
- Comments explaining what the code does rather than why
- "Optimizations" without performance measurements
- Complex caching mechanisms without evidence they're needed

## Related Tenets

- [Simplicity](/tenets/simplicity.md): Simple code is easier to maintain
- [Explicit is Better than Implicit](/tenets/explicit-over-implicit.md): Clarity improves maintainability