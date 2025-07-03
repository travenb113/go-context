# go-context ✨

This repository contains concise and focused examples demonstrating the use of the `context` package in Go.  
Each example is placed in a separate folder and illustrates a specific concept.

---

### 📁 01. HTTP Context Cancel  
Path: http-context-cancel/main.go

This example demonstrates:
- How to detect when the client cancels the request (for example, closes the tab or disconnects).
- Using `req.Context()` and listening to `<-ctx.Done()` to stop processing when the client is gone.
- Returning an error response instead of finishing unnecessary work.

Key Concepts:
- Context cancellation on client disconnect
- `select` statement with `<-ctx.Done()`
- Avoiding wasted work on cancelled requests

---

### 📁 02. With Timeout  
Path: with-timeout/main.go

This example shows:
- Limiting the execution time of a request handler.
- Using `context.WithTimeout` to automatically cancel a long-running task.
- Returning a timeout message to the client if the task takes too long.

Key Concepts:
- `context.WithTimeout`
- Deadline exceeded errors
- Time-limited request handling

---

### 📁 03. With Value  
Path: with-value/main.go

This example illustrates:
- Passing data (like userID or traceID) between functions using `context.WithValue`.
- Extracting the value in another function via `ctx.Value(...)`.
- Keeping data request-scoped instead of using global variables.

Key Concepts:
- `context.WithValue`
- Passing data safely through request lifecycle
- Clean design for request-scoped data

---

### 📁 04. Graceful Shutdown  
Path: graceful-shutdown/main.go

This example demonstrates:
- Running the HTTP server in a separate goroutine so the main thread can listen for OS signals.
- Catching Ctrl+C (`os.Interrupt`) using `signal.Notify`.
- Gracefully shutting down the server using `server.Shutdown(ctx)` with a timeout.

Key Concepts:
- Graceful shutdown with context
- Using `os.Signal` and `signal.Notify`
- Avoiding abrupt termination of active requests

---

## 🧠 Why These Examples Matter

The `context` package is one of the most practical and powerful tools in Go.  
It helps you:
- Cancel or time-limit operations cleanly.
- Pass request-scoped data without globals.
- Build robust, production-ready services that handle signals gracefully.

---

## 🚀 Getting Started

Clone the repository and run each example separately:

```bash
go run path/to/example/main.go
```
## 📚 Resources & References

This project is part of my learning path to master Go’s context package.
Examples are inspired by official docs and real-world experimentation.

[Go.dev Blog](https://pkg.go.dev/context) — in-depth guide.

[Go by Example: Context](https://gobyexample.com/context) — quick intro.

[YouTube tutorials](https://youtu.be/LSzR0VEraWw?si=y_C6XEvmioYX2eWr) — clear and simple video explanations.