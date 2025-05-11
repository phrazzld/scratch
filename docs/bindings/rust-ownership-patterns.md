______________________________________________________________________

id: rust-ownership-patterns last_modified: '2025-05-04' derived_from: simplicity
enforced_by: rust compiler & code review applies_to:

- rust

______________________________________________________________________

# Binding: Embrace Rust's Ownership System, Don't Fight It

Design Rust code to work with the ownership system, not against it. Use ownership,
borrowing, and lifetimes as core design elements that guide your APIs and data
structures. Embrace the constraints of the ownership model to create safer, more
maintainable code without resorting to excessive cloning, unsafe code, or overly complex
lifetime annotations.

## Rationale

This binding implements our simplicity tenet by leveraging Rust's ownership system to
prevent entire categories of bugs that would otherwise add tremendous complexity to your
codebase.

Rust's ownership model fundamentally changes the nature of complexity in your software.
Most languages require you to manually track who owns what data, when it's safe to
modify values, and when resources need to be cleaned up. Even with disciplined
programmers, these mental bookkeeping tasks are error-prone and create a form of
"complexity debt" that grows exponentially with codebase size. Rust's ownership system
moves this complexity from runtime (where it causes subtle, hard-to-fix bugs) to compile
time (where it can be systematically addressed).

Think of Rust's ownership system like a diligent legal clerk who reviews every transfer
of property to ensure it's legitimate and properly recorded. While this might initially
seem like added bureaucracy, it actually eliminates endless property disputes that would
otherwise plague your codebase. Just as clear ownership records prevent real-world legal
battles over who owns what, Rust's ownership system prevents runtime battles over who
can access and modify memory. The upfront cost in design consideration pays enormous
dividends in reduced debugging time, fewer crashes, and more maintainable code.

When developers first encounter Rust, they often struggle against the borrow checker,
trying to make patterns from other languages fit into Rust's model. This fighting
creates artificial complexity. By designing your code to work with ownership constraints
rather than against them, you create naturally simpler, more robust software. The
patterns in this binding help you shift your thinking to align with Rust's
model—treating ownership as a feature to be leveraged, not an obstacle to be overcome.

## Rule Definition

This binding establishes clear rules for effective use of Rust's ownership, borrowing,
and lifetime systems:

- **Ownership as API Design:** Structure APIs around ownership patterns, not despite
  them. Consider who owns each value and for how long as a foundational aspect of API
  design.

  - Use ownership transfer (moving values) when the function takes responsibility for
    the value
  - Use shared references (`&T`) for read-only access to data
  - Use mutable references (`&mut T`) when temporary write access is needed
  - Return newly created or transformed values rather than mutating inputs when possible

- **Borrow Checking Compliance:** Never try to circumvent the borrow checker with
  complex schemes, unsafe code, or excessive use of interior mutability.

  - If the borrow checker is fighting your design, rethink the design
  - Prefer ownership patterns that naturally satisfy the borrow checker
  - Use reference counting (`Rc`, `Arc`) and interior mutability (`RefCell`, `Mutex`)
    judiciously, only when truly needed

- **Borrowing Over Copying:** Prefer borrowing (`&T`) over cloning/copying where
  possible.

  - Only clone data when the clone will be modified independently
  - Use references when you only need to inspect data
  - Consider Copy traits only for small, stack-based types where copying is cheaper than
    borrowing

- **Lifetime Management:** Keep lifetime annotations as simple as possible.

  - Let Rust's lifetime elision rules work for you when possible
  - Use explicit lifetimes only when necessary to express constraints
  - Avoid creating complex lifetime relationships that are difficult to understand
  - Structure your types to minimize the need for complex lifetime annotations

- **RAII Resource Management:** Use Rust's RAII (Resource Acquisition Is Initialization)
  pattern for all resource management.

  - Every resource should have a clear owner responsible for cleanup
  - Use the `Drop` trait to ensure proper cleanup of resources
  - Avoid manual resource management patterns from other languages

- **Unsafe Usage Restrictions:** Unsafe code must be:

  - Strictly minimized and isolated in private implementation details
  - Thoroughly documented with `// SAFETY:` comments
  - Abstracted behind safe interfaces with appropriate invariants
  - Carefully reviewed to ensure all language safety requirements are met

## Practical Implementation

1. **Design Ownership-Friendly APIs**: Structure function signatures to clearly
   communicate ownership intent:

   ```rust
   // Consuming APIs - take ownership when the function needs to store or transform the value
   fn process_message(message: Message) -> Result<Response, Error> {
       // Function takes ownership of message, can store it or transform it freely
   }

   // Non-consuming APIs - borrow when the function only needs to read
   fn validate_message(message: &Message) -> bool {
       // Function only reads message, doesn't affect its ownership
   }

   // Mutating APIs - use mutable borrows for modification
   fn update_message(message: &mut Message, new_content: &str) {
       // Function temporarily modifies message in-place
   }

   // Factory pattern - return values instead of mutating
   fn enrich_message(message: Message, metadata: &Metadata) -> Message {
       // Returns a new, enhanced Message rather than mutating the input
   }
   ```

1. **Use Ownership Transfer for Clear Resource Management**: Design with ownership
   transfer for resources:

   ```rust
   // Connection clearly owns the socket and is responsible for cleanup
   struct Connection {
       socket: Socket,
       // other fields...
   }

   impl Connection {
       // Constructor takes ownership of the socket
       fn new(socket: Socket, config: &ConnectionConfig) -> Self {
           Connection {
               socket,
               // initialize other fields...
           }
       }

       // Methods use &self or &mut self to operate on the owned socket
       fn send_data(&mut self, data: &[u8]) -> Result<usize, Error> {
           self.socket.write(data)
       }
   }

   // When Connection is dropped, socket is automatically cleaned up
   ```

1. **Implement Borrowing Patterns for Shared Data**: Use references for data that's
   shared but not transferred:

   ```rust
   // Registry shares access to configurations without taking ownership
   struct Registry<'a> {
       configurations: Vec<&'a Configuration>,
   }

   impl<'a> Registry<'a> {
       fn new() -> Self {
           Registry {
               configurations: Vec::new(),
           }
       }

       // Borrows the configuration rather than taking ownership
       fn register(&mut self, config: &'a Configuration) {
           self.configurations.push(config);
       }

       fn find_by_name(&self, name: &str) -> Option<&'a Configuration> {
           self.configurations.iter()
               .find(|config| config.name == name)
               .copied()
       }
   }
   ```

1. **Use References and Slices Instead of Cloning**: Prefer borrowing over cloning when
   possible:

   ```rust
   // ✅ GOOD: Takes string slices instead of owned Strings
   fn format_name(first: &str, last: &str) -> String {
       format!("{}, {}", last, first)
   }

   // ✅ GOOD: Takes a slice of records instead of cloning a vector
   fn calculate_average(values: &[f64]) -> f64 {
       if values.is_empty() {
           return 0.0;
       }
       values.iter().sum::<f64>() / values.len() as f64
   }
   ```

1. **Implement Borrowing with Lifetimes**: Use lifetimes to express relationships
   between borrowed values:

   ```rust
   // Parser borrows the input text and returns slices of that same text
   struct Parser<'a> {
       input: &'a str,
       position: usize,
   }

   impl<'a> Parser<'a> {
       fn new(input: &'a str) -> Self {
           Parser {
               input,
               position: 0,
           }
       }

       // Return type shares the same lifetime as the input
       fn parse_identifier(&mut self) -> Option<&'a str> {
           // Implementation returns a slice of the original input
           let start = self.position;

           // Find the end of the identifier
           while self.position < self.input.len() &&
                 self.input.as_bytes()[self.position].is_ascii_alphanumeric() {
               self.position += 1;
           }

           if start == self.position {
               None
           } else {
               Some(&self.input[start..self.position])
           }
       }
   }
   ```

1. **Apply the Builder Pattern for Complex Construction**: Use the builder pattern to
   construct complex objects:

   ```rust
   struct ServerConfig {
       address: String,
       port: u16,
       max_connections: usize,
       timeout_seconds: u64,
       tls_enabled: bool,
       // More fields...
   }

   // Builder owns the partially constructed config
   struct ServerConfigBuilder {
       config: ServerConfig,
   }

   impl ServerConfigBuilder {
       fn new() -> Self {
           ServerConfigBuilder {
               config: ServerConfig {
                   address: String::from("127.0.0.1"),
                   port: 8080,
                   max_connections: 100,
                   timeout_seconds: 30,
                   tls_enabled: false,
               },
           }
       }

       // Each method takes and returns ownership of the builder
       fn address(mut self, address: impl Into<String>) -> Self {
           self.config.address = address.into();
           self
       }

       fn port(mut self, port: u16) -> Self {
           self.config.port = port;
           self
       }

       // More builder methods...

       // Finalize by transferring ownership of the config
       fn build(self) -> ServerConfig {
           self.config
       }
   }
   ```

1. **Use Smart Pointers Judiciously**: Apply appropriate smart pointers for specific
   ownership needs:

   ```rust
   use std::rc::Rc;
   use std::cell::RefCell;

   // Shared ownership with interior mutability - use when truly needed
   struct Document {
       content: String,
       // Other fields...
   }

   // Editor shares ownership of the document with other components
   struct Editor {
       document: Rc<RefCell<Document>>,
       // Editor-specific fields...
   }

   // UndoStack also needs shared access to the document
   struct UndoStack {
       document: Rc<RefCell<Document>>,
       history: Vec<String>,
   }

   impl Editor {
       fn new(document: Rc<RefCell<Document>>) -> Self {
           Editor {
               document,
               // Initialize other fields...
           }
       }

       fn insert_text(&mut self, position: usize, text: &str) {
           // Borrow mutably, but only for the scope of this block
           let mut doc = self.document.borrow_mut();
           // Modify the document...
           doc.content.insert_str(position, text);
       }
   }
   ```

## Examples

```rust
// ❌ BAD: Fighting the borrow checker with excessive cloning
fn process_data(data: &mut Vec<String>, query: &str) -> Vec<String> {
    let results = data.iter()
        .filter(|item| item.contains(query))
        .cloned()  // Unnecessary clone of each matching item
        .collect::<Vec<_>>();

    for item in &results {
        data.push(item.clone());  // Another clone to add back to data
    }

    results  // Return cloned results
}
```

```rust
// ✅ GOOD: Working with the ownership system
fn process_data<'a>(data: &'a mut Vec<String>, query: &str) -> Vec<&'a String> {
    // Collect references to matching items without cloning
    let results: Vec<&String> = data.iter()
        .filter(|item| item.contains(query))
        .collect();

    // No need to clone when adding new items
    for &item in &results {
        // Since we're borrowing from data, we can't modify it while those
        // borrows are active. This wouldn't compile as written.
        // We'd need to restructure, which is good - the borrow checker
        // is highlighting a potential problem in our logic.
    }

    results  // Return references to the original data
}

// Better restructured version with clear ownership semantics
fn process_data(data: &[String], query: &str) -> Vec<String> {
    // First collect matching items
    let matches: Vec<String> = data.iter()
        .filter(|item| item.contains(query))
        .cloned()
        .collect();

    matches  // Return owned matches
}
```

```rust
// ❌ BAD: Excessive use of Rc<RefCell<T>> for simple operations
struct UserManager {
    users: Rc<RefCell<Vec<User>>>,
}

impl UserManager {
    fn add_user(&self, user: User) {
        self.users.borrow_mut().push(user);
    }

    fn get_user(&self, id: &str) -> Option<User> {
        self.users.borrow()
            .iter()
            .find(|user| user.id == id)
            .cloned()
    }
}

// Usage
let manager = UserManager { users: Rc::new(RefCell::new(Vec::new())) };
manager.add_user(User::new("alice", "Alice"));
let user = manager.get_user("alice");  // Gets a cloned User
```

```rust
// ✅ GOOD: Clear ownership with appropriate borrowing
struct UserManager {
    users: Vec<User>,
}

impl UserManager {
    fn add_user(&mut self, user: User) {
        self.users.push(user);
    }

    fn get_user(&self, id: &str) -> Option<&User> {
        self.users.iter().find(|user| user.id == id)
    }
}

// Usage
let mut manager = UserManager { users: Vec::new() };
manager.add_user(User::new("alice", "Alice"));
let user = manager.get_user("alice");  // Gets a reference to the User
```

```rust
// ❌ BAD: Unnecessary complex lifetimes
struct Processor<'a, 'b, 'c> {
    input: &'a str,
    config: &'b Configuration,
    output_buffer: &'c mut String,
}

impl<'a, 'b, 'c> Processor<'a, 'b, 'c> {
    fn process(&mut self) -> Result<&'c str, Error> {
        // Complex processing with multiple lifetimes...
        self.output_buffer.push_str(self.input);
        Ok(&self.output_buffer[..])
    }
}
```

```rust
// ✅ GOOD: Simplified lifetime management
struct Processor<'a> {
    input: &'a str,
    config: &'a Configuration,
}

impl<'a> Processor<'a> {
    fn new(input: &'a str, config: &'a Configuration) -> Self {
        Processor { input, config }
    }

    fn process(&self, output_buffer: &mut String) -> Result<(), Error> {
        // Process and write to the provided buffer
        output_buffer.push_str(self.input);
        Ok(())
    }
}
```

## Real-World Example: Resource Management

Here's a more complete example showing ownership-based resource management for a file
processing system:

```rust
// Define a resource that requires cleanup
struct FileProcessor {
    file: std::fs::File,
    buffer: Vec<u8>,
}

impl FileProcessor {
    // Constructor takes ownership of the file
    fn new(file: std::fs::File) -> Self {
        FileProcessor {
            file,
            buffer: Vec::with_capacity(4096),
        }
    }

    // Process takes &mut self - temporary mutable access
    fn process(&mut self) -> Result<usize, std::io::Error> {
        use std::io::Read;
        self.buffer.clear();
        let bytes_read = self.file.read_to_end(&mut self.buffer)?;

        // Process the data...

        Ok(bytes_read)
    }

    // Extract results without taking ownership of the processor
    fn results(&self) -> &[u8] {
        &self.buffer
    }

    // Optional: explicit cleanup method if needed beyond Drop
    fn cleanup(self) -> std::fs::File {
        // Return ownership of the file for potential reuse
        self.file
    }
}

// The Drop trait ensures cleanup happens automatically
impl Drop for FileProcessor {
    fn drop(&mut self) {
        // Any cleanup beyond the automatic dropping of file and buffer
        println!("FileProcessor being cleaned up");
    }
}

// Usage example showing clear ownership flow
fn process_file(path: &str) -> Result<Vec<u8>, std::io::Error> {
    // Open takes &str (borrowed) and returns owned File
    let file = std::fs::File::open(path)?;

    // Create processor by transferring ownership of file
    let mut processor = FileProcessor::new(file);

    // Process borrows processor mutably
    processor.process()?;

    // Get results without taking ownership
    let results = processor.results().to_vec();

    // Processor is dropped here, automatically cleaning up resources
    Ok(results)
}
```

## Related Bindings

- [immutable-by-default](immutable-by-default.md): Rust's ownership system naturally
  encourages immutability through shared references (`&T`). When data is borrowed
  immutably, the compiler guarantees it won't change, creating the same benefits as
  explicit immutability patterns in other languages.

- [no-internal-mocking](no-internal-mocking.md): Rust's trait system combined with
  ownership makes it natural to define abstractions that can be easily tested without
  complex mocking. Clear ownership boundaries create natural seams for testing.

- [dependency-inversion](dependency-inversion.md): In Rust, dependency inversion is
  often implemented through traits, which pair naturally with the ownership system. The
  ownership constraints of trait objects add clarity about who owns what, making
  dependency relationships more explicit.

- [simplicity](../tenets/simplicity.md): Rust's ownership system might seem to add
  complexity at first, but it actually reduces accidental complexity by catching entire
  classes of bugs at compile time. What appears as additional syntax (borrowing,
  lifetimes) is actually a way to express constraints that would otherwise create hidden
  runtime complexity.
