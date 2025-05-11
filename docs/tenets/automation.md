---
id: automation
last_modified: "2025-05-02"
---

# Tenet: Automation Is Mandatory

Manual, repetitive steps are bugs in the process. Automate every feasible repetitive task to ensure consistency, eliminate human error, and free up developers to focus on creative problem-solving.

## Core Belief

Automation reduces manual error, ensures consistency, frees up developer time, provides faster feedback, and makes processes repeatable and reliable. If a task is done more than twice, it should be automated.

## Practical Guidelines

1. **Identify Automation Opportunities**: Common candidates include:
   - Testing (unit, integration, end-to-end)
   - Code quality checks (linting, formatting, static analysis)
   - Build and deployment processes
   - Dependency management
   - Vulnerability scanning
   - Version management
   - Documentation generation
   - Release notes and changelog creation

2. **Invest in Developer Tooling**: Good tools pay for themselves quickly in time saved.

3. **Automate Checks Early**: Run automated checks as early in the development cycle as possible:
   - Pre-commit hooks
   - Continuous Integration pipelines
   - Pull request validations

4. **Enforce Automation Use**: Tools like pre-commit hooks should be mandatory, not optional.

5. **Keep Automation Maintained**: Regularly review and update automation scripts and workflows.

## Warning Signs

- "It's faster to just do it manually"
- "I'll automate it later"
- Bypassing quality gates (e.g., using `--no-verify` with Git hooks)
- Inconsistent build or release processes
- Manual steps documented in READMEs or wikis
- "Works on my machine" problems

## Related Tenets

- [No Secret Suppressions](/tenets/no-secret-suppression.md): Enforcing quality checks
- [Simplicity](/tenets/simplicity.md): Automation should simplify, not complicate