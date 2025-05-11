---
id: require-conventional-commits
last_modified: "2025-05-02"
derived_from: automation
enforced_by: commit hooks & CI checks
applies_to:
  - all
---

# Binding: Use Conventional Commits

All commit messages MUST follow the Conventional Commits specification to enable automated versioning, changelog generation, and semantic release management.

## Rationale

Standardized commit messages make the project history more readable and enable automation of release processes. With Conventional Commits, we can automatically determine the next semantic version, generate changelogs, and trigger appropriate CI/CD workflows based on the types of changes made.

## Enforcement

This binding is enforced by:

1. Pre-commit hooks that validate commit message format
2. CI checks that verify commit messages adhere to the convention
3. Rejecting pull requests with non-conforming commit messages

## Commit Message Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

Where:

- `type` must be one of:
  - `feat`: A new feature (minor version bump)
  - `fix`: A bug fix (patch version bump)
  - `docs`: Documentation only changes
  - `style`: Changes that don't affect code meaning (whitespace, formatting)
  - `refactor`: Code change that neither fixes a bug nor adds a feature
  - `perf`: Performance improvement
  - `test`: Adding or correcting tests
  - `build`: Changes to build system or external dependencies
  - `ci`: Changes to CI configuration files and scripts
  - `chore`: Other changes that don't modify src or test files

- `scope` is optional and indicates the section of the codebase affected

- `description` is a short description of the change

- Add a `!` after the type/scope to indicate a breaking change (major version bump), e.g., `refactor!: remove deprecated API`

## Examples

```
# Feature with scope
feat(api): add user profile endpoint

# Bug fix
fix: prevent race condition in connection pool

# Documentation change
docs: update installation instructions

# Breaking change
refactor!: rename core service interfaces
```

## Related Bindings

- [semantic-versioning.md](./semantic-versioning.md) - Versioning strategy using SemVer
- [automate-changelog.md](./automate-changelog.md) - Automated changelog generation