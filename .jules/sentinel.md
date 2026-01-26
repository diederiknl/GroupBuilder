## 2026-01-26 - Hardcoded JWT Secret in Auth Package
**Vulnerability:** The `internal/auth` package defined `jwtKey` as a hardcoded byte slice, making it impossible to rotate secrets without recompiling and exposing the key in the source code.
**Learning:** Hardcoded secrets often appear in initial prototypes or "stub" implementations (like the `TODO` stubs found elsewhere) and persist if not explicitly replaced by configuration logic.
**Prevention:** Use `os.Getenv` in an `init()` function or a configuration loader to populate sensitive keys, providing a fallback only for local development with a clear warning.
