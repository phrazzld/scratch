______________________________________________________________________

id: rust-error-handling last_modified: '2025-05-06' derived_from: simplicity
enforced_by: Code review, Clippy static analysis applies_to:

- rust

______________________________________________________________________

# Binding: Rust Error Handling

Handle errors explicitly in Rust using the type system. Return `Result` for all
operations that can fail, use custom error types to convey precise error information,
and leverage the `?` operator for concise propagation.

## Rationale

This binding directly implements our simplicity and explicit-over-implicit tenets by
making errors a first-class concern in the type system instead of hidden control flow.
Rust's approach to error handling represents a significant advancement over traditional
exceptions by making the possibility of failure explicit in function signatures. When
you see a function returning `Result<T, E>`, you immediately know it might fail and in
what ways it could fail through the `E` type parameter.

The traditional approach to error handling in many languages relies on exceptions that
can propagate invisibly through multiple layers of code. This creates an implicit
contract where any function could potentially throw an exception that callers might not
be aware of or prepared to handle. Rust's approach makes these error paths explicit and
ensures they're handled deliberately. Think of Rust's `Result` type as a contract that
says, "I promise to give you either the value you want or a clear explanation of why I
couldn't."

By enforcing explicit error handling through the type system, we eliminate entire
categories of bugs related to unhandled exceptions, unexpected program termination, and
unclear error propagation paths. This significantly improves both code robustness and
maintainability, as every potential point of failure is clearly marked and must be
explicitly addressed. The compiler becomes your ally in ensuring complete error
handling, preventing problems from reaching production.

## Rule Definition

This binding establishes the following rules for error handling in Rust:

- **Use `Result<T, E>` for all fallible operations**: Any function that can fail in
  expected ways must return a `Result` type that clearly communicates both the success
  value and the potential error types.

- **Never use `panic!` for expected failures**: The `panic!` mechanism is reserved
  exclusively for truly unrecoverable situations that indicate actual programming
  errors, invariant violations, or impossible states. Never use it for expected error
  conditions that should be handled by normal program flow.

- **Create domain-specific error types**: Define custom error types that precisely
  represent the kinds of failures that can occur in your domain, rather than using
  generic error messages or strings.

- **Use the `?` operator for error propagation**: Leverage Rust's `?` operator for
  concise and clear error propagation rather than verbose match or if-let expressions.

- **Provide context with wrapped errors**: When propagating errors from lower-level
  components, add context to help diagnose issues by wrapping the original error with
  additional information about what operation was being attempted.

Exceptions to these rules are extremely limited:

- In application `main()` functions or similar top-level contexts where errors can only
  be reported, not handled
- In test code where panics are part of the testing mechanism
- In prototyping or exploratory code that will be refactored before production use

## Practical Implementation

1. **Create Custom Error Types**: Define error types that are specific to your domain
   using enums to represent different failure modes. Leverage the `thiserror` crate to
   reduce boilerplate.

```rust
use thiserror::Error;

#[derive(Error, Debug)]
pub enum UserServiceError {
    #[error("user not found with id {0}")]
    UserNotFound(String),

    #[error("database error: {0}")]
    DatabaseError(#[from] DatabaseError),

    #[error("validation error: {0}")]
    ValidationError(String),
}
```

2. **Design Function Signatures with `Result`**: Make the possibility of failure
   explicit in your function signatures by returning `Result` types. This makes the
   contract clear to callers.

```rust
pub fn get_user(id: &str) -> Result<User, UserServiceError> {
    // Implementation...
}
```

3. **Leverage the `?` Operator**: Use the `?` operator for concise error propagation
   within functions that return `Result`. This operator extracts the value from a
   `Result` if it's `Ok`, or returns early with the error if it's `Err`.

```rust
pub fn process_user_data(id: &str) -> Result<ProcessedData, UserServiceError> {
    let user = get_user(id)?; // Early return if error
    let settings = get_user_settings(id)?; // Early return if error
    let processed = transform_data(user, settings)?; // Early return if error

    Ok(processed)
}
```

4. **Add Context When Propagating Errors**: When wrapping errors from other modules or
   libraries, add context to make debugging easier. Use the `context` method from the
   `anyhow` crate in applications or implement similar functionality in your custom
   errors.

```rust
// In binary crates using anyhow
use anyhow::{Context, Result};

fn load_config(path: &Path) -> Result<Config> {
    let content = std::fs::read_to_string(path)
        .with_context(|| format!("failed to read config from {}", path.display()))?;

    parse_config(&content)
        .with_context(|| format!("failed to parse config from {}", path.display()))
}

// In library crates with custom error types
fn load_config(path: &Path) -> Result<Config, ConfigError> {
    let content = std::fs::read_to_string(path)
        .map_err(|e| ConfigError::IoError {
            source: e,
            path: path.to_path_buf(),
        })?;

    parse_config(&content)
        .map_err(|e| ConfigError::ParseError {
            source: e,
            path: path.to_path_buf(),
        })
}
```

5. **Handle Different Error Cases Explicitly**: When consuming errors, handle different
   error cases explicitly using pattern matching where appropriate. This is especially
   important at API boundaries.

```rust
match get_user(user_id) {
    Ok(user) => {
        // Happy path processing...
    },
    Err(UserServiceError::UserNotFound(_)) => {
        // Handle missing user case...
    },
    Err(UserServiceError::ValidationError(msg)) => {
        // Handle validation error...
    },
    Err(e) => {
        // Handle other errors...
        log::error!("Unexpected error fetching user: {}", e);
    }
}
```

## Examples

```rust
// ❌ BAD: Using unwrap/expect for normal error handling
fn get_config() -> Config {
    let file = std::fs::File::open("config.json").expect("Failed to open config file");
    let reader = std::io::BufReader::new(file);
    serde_json::from_reader(reader).expect("Failed to parse config")
}

// ✅ GOOD: Using Result to make errors explicit and propagate them
fn get_config() -> Result<Config, ConfigError> {
    let file = std::fs::File::open("config.json")
        .map_err(|e| ConfigError::IoError { source: e, file: "config.json" })?;
    let reader = std::io::BufReader::new(file);
    let config = serde_json::from_reader(reader)
        .map_err(|e| ConfigError::ParseError { source: e, file: "config.json" })?;
    Ok(config)
}
```

```rust
// ❌ BAD: Using strings as errors loses type information and context
fn process_data(data: &str) -> Result<ProcessedData, String> {
    if data.is_empty() {
        return Err("Data cannot be empty".to_string());
    }

    let parsed = match parse_input(data) {
        Ok(p) => p,
        Err(e) => return Err(format!("Failed to parse: {}", e)),
    };

    // Rest of processing...
    Ok(ProcessedData::new(parsed))
}

// ✅ GOOD: Using typed errors with proper context
#[derive(Error, Debug)]
enum ProcessingError {
    #[error("empty input data")]
    EmptyData,

    #[error("parsing error: {0}")]
    ParseError(#[from] ParseError),

    #[error("validation error: {0}")]
    ValidationError(String),
}

fn process_data(data: &str) -> Result<ProcessedData, ProcessingError> {
    if data.is_empty() {
        return Err(ProcessingError::EmptyData);
    }

    let parsed = parse_input(data)?; // Automatically converts ParseError

    // Rest of processing...
    Ok(ProcessedData::new(parsed))
}
```

```rust
// ❌ BAD: Using panic for expected error conditions
fn divide(a: i32, b: i32) -> i32 {
    if b == 0 {
        panic!("Division by zero");
    }
    a / b
}

// ❌ BAD: Returning Option without error context
fn divide(a: i32, b: i32) -> Option<i32> {
    if b == 0 {
        None
    } else {
        Some(a / b)
    }
}

// ✅ GOOD: Using Result with a specific error type
#[derive(Error, Debug)]
pub enum MathError {
    #[error("division by zero")]
    DivisionByZero,
}

fn divide(a: i32, b: i32) -> Result<i32, MathError> {
    if b == 0 {
        Err(MathError::DivisionByZero)
    } else {
        Ok(a / b)
    }
}
```

## Related Bindings

- [go-error-wrapping.md](go-error-wrapping.md): Similar to this binding, but for Go.
  Both emphasize explicit error handling and proper context propagation, though Rust
  leverages its type system more extensively for compile-time guarantees.

- [rust-ownership-patterns.md](rust-ownership-patterns.md): Complements error handling
  by defining how resources should be managed. Proper ownership patterns create a
  foundation for robust error handling by ensuring resources are properly cleaned up
  even when errors occur.

- [dependency-inversion.md](dependency-inversion.md): Works with error handling to
  create cleanly separated components with well-defined error boundaries, making it
  easier to handle specific error types at the appropriate level of abstraction.

- [immutable-by-default.md](immutable-by-default.md): Supports robust error handling by
  reducing the complexity of state management, making it easier to reason about where
  and how errors might occur.
