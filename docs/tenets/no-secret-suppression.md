---
id: no-secret-suppression
last_modified: "2025-05-02"
---

# Tenet: No Secret Suppressions

No disabling lint/errors without public rationale. Code quality tools and type systems exist to prevent defects - suppressing their warnings without justification hides problems rather than addressing them.

## Core Belief

Quality gates and static analysis tools are critical safeguards that help prevent defects and maintain consistent code quality. Suppressing these safeguards without clear justification represents hidden technical debt and potential future issues.

## Practical Guidelines

1. **Address The Root Cause**: When facing a lint error or type issue, fix the underlying problem, not the warning.

2. **Justify Any Exception**: In the rare case where a suppression is genuinely needed:
   - Document the specific reason with a detailed comment
   - Explain why it's safe in this particular case
   - Include who approved the exception
   - Consider adding a ticket reference to revisit later

3. **Require Review**: Any suppressions must be explicitly called out and justified during code review.

4. **Re-evaluate Suppressions**: Regularly audit existing suppressions to see if they can be removed.

5. **Adapt Rules When Necessary**: If a rule consistently causes problems, consider if it should be modified globally rather than suppressed repeatedly.

## Warning Signs

- Comments like `// eslint-disable-line` or `// nolint` without explanation
- Type assertions (`as any` in TypeScript) to bypass type system
- Many suppressions in a single file or module
- Using override flags to bypass checks (`--no-verify`, `--force`)
- Lack of static analysis tools in the development pipeline

## Related Tenets

- [Automation](/tenets/automation.md): Tooling and CI for consistent enforcement
- [Explicitness](/tenets/explicit-over-implicit.md): Making intentions and decisions clear