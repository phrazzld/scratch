---
id: no-lint-suppression
last_modified: "2025-05-02"
derived_from: no-secret-suppression
enforced_by: code review & custom linters
applies_to:
  - all
---

# Binding: No Lint Suppression Without Justification

Directives that disable linter rules (such as `// eslint-disable-line`, `// nolint`, `#[allow(...)]`) must be accompanied by a detailed comment explaining why the suppression is necessary and safe. Unjustified suppressions are forbidden.

Lint suppressions represent a deliberate circumvention of quality checks. While sometimes necessary due to tool limitations or specific use cases, they must be explicitly justified rather than quietly bypassing quality measures. Undocumented suppressions can hide potential bugs, technical debt, or design problems.

## Rationale

Lint suppressions represent a deliberate circumvention of quality checks. While sometimes necessary due to tool limitations or specific use cases, they must be explicitly justified rather than quietly bypassing quality measures. Undocumented suppressions can hide potential bugs, technical debt, or design problems.

## Enforcement

This binding is enforced by:

1. Code review processes that reject unjustified suppressions
2. Custom linters that verify all suppressions have accompanying comments
3. Git hooks that prevent commits with unexplained suppressions
4. Regular audits of suppression usage

## Guidelines

When a suppression is genuinely necessary:

1. Add a detailed comment explaining:
   - Why the rule is triggering in this specific case
   - Why the code is actually correct/safe despite the lint warning
   - Why refactoring to avoid the suppression isn't feasible

2. Scope the suppression as narrowly as possible:
   - Target specific lines rather than entire files
   - Disable only the specific rule needed, not all linting
   - Apply suppression only to the minimum scope required

3. Consider adding a ticket reference to revisit the suppression in the future

## Examples

```typescript
// ❌ BAD: Unexplained suppression
// eslint-disable-next-line @typescript-eslint/no-non-null-assertion
const value = data.optionalField!;

// ✅ GOOD: Justified suppression
// eslint-disable-next-line @typescript-eslint/no-non-null-assertion
// This field is populated in validateData() which is always called before this function
// The assertion is safer than adding a runtime check that would never fail
const value = data.optionalField!;
```

## Related Bindings

- [ts-no-type-assertion.md](./ts-no-type-assertion.md) - Similar rules for TypeScript type assertions
- [pre-commit-hooks.md](./pre-commit-hooks.md) - Configuration of pre-commit hooks that enforce this rule