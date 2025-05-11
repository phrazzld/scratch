---
id: document-decisions
last_modified: "2025-05-02"
---

# Tenet: Document Decisions, Not Mechanics

Strive for self-documenting code through clear naming, structure, and types for the "how." Reserve comments and external documentation primarily for the "why": rationale, context, constraints, and trade-offs.

## Core Belief

Code mechanics change; comments detailing "how" quickly become outdated or redundant. The reasoning behind decisions provides enduring value. Self-documenting code reduces the documentation synchronization burden while meaningful documentation of decisions preserves critical context.

## Practical Guidelines

1. **Focus on the Why**: Document the rationale behind non-obvious choices, trade-offs considered, constraints addressed, and business rules implemented.

2. **Self-Documenting Code**: Use clear naming, structure, and types to make the "how" obvious without comments.

3. **Document at the Right Level**: Different documentation serves different purposes:
   - Inline comments: Explain the "why" for specific code decisions
   - Function/method docs: Describe contracts, side effects, and usage expectations
   - Module/package docs: Explain purpose, responsibility, and relationship to other components
   - Project docs: Provide architecture overview, development workflows, and contribution guidelines

4. **Keep Documentation Current**: Outdated documentation is worse than no documentation. Update docs alongside code changes.

5. **Use Living Documents**: Consider structured formats like Architecture Decision Records (ADRs) to document significant choices.

## Warning Signs

- Comments that simply restate what the code does
- Outdated or inconsistent documentation
- Missing rationale for complex or unusual approaches
- No architecture overview for new team members
- Documentation that contradicts the actual code

## Related Tenets

- [Explicit is Better than Implicit](/tenets/explicit-over-implicit.md): Clear code reduces the need for mechanical documentation
- [Maintainability](/tenets/maintainability.md): Good documentation improves long-term maintenance