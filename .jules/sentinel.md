## 2025-05-15 - [Database Testability Pattern]
**Vulnerability:** Difficulty in verifying security fixes with isolated tests due to hardcoded database paths.
**Learning:** `InitDB` was hardcoding `./groupbuilder.db`, preventing the use of in-memory databases for secure authentication testing.
**Prevention:** Refactored `InitDB` to accept a connection string (e.g., `":memory:"`), enabling isolated, safe unit tests for security handlers.
