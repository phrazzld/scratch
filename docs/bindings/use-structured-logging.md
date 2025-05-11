---
id: use-structured-logging
last_modified: "2025-05-02"
derived_from: automation
enforced_by: linters & code review
applies_to:
  - all
  - typescript
  - javascript
  - go
---

# Binding: Use Structured Logging Only

All operational logging must use structured logging in JSON format. Direct use of `fmt.Println`, `console.log`, or similar print-style logging is forbidden in production code. Logs must include standard context fields.

## Rationale

Structured logging enables efficient parsing, filtering, and analysis by log aggregation systems. Consistent log formats with mandatory contextual fields make troubleshooting easier, especially in distributed systems. By enforcing a standard approach, we ensure logs are machine-readable while remaining human-friendly.

## Enforcement

This binding is enforced by:

1. Language-specific linters (e.g., ESLint rules against `console.log`)
2. Code review requiring proper logging practices
3. CI checks that fail on detection of prohibited logging patterns

## Mandatory Context Fields

All log entries must include at minimum:

- `timestamp` (ISO 8601 format, UTC)
- `level` (e.g., "info", "error")
- `message` (clear description of the event)
- `service_name` / `application_id`
- `correlation_id` (Request ID, Trace ID)
- `function_name` / `module_name` (Where the log originated)
- `error_details` (For ERROR level: type, message, stack trace)

## Implementation by Language

### JavaScript/TypeScript

Use structured logging libraries like `pino` or `winston` configured for JSON output:

```typescript
// ❌ BAD: Unstructured logging
console.log("User logged in:", userId);

// ✅ GOOD: Structured logging
logger.info({ userId, action: "login" }, "User successfully authenticated");
```

### Go

Use structured logging packages like `zap`, `zerolog`, or `slog`:

```go
// ❌ BAD: Unstructured logging
fmt.Printf("Processing order %s\n", orderID)

// ✅ GOOD: Structured logging
logger.Info("Processing order", 
    zap.String("order_id", orderID),
    zap.String("correlation_id", ctx.Value("correlation_id").(string)))
```

## Related Bindings

- [context-propagation.md](./context-propagation.md) - Propagating correlation IDs across service boundaries
- [logging-levels.md](./logging-levels.md) - Standard log levels and their appropriate usage