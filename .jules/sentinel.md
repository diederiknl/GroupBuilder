## 2026-01-25 - [Hardcoded Secrets Pattern]
**Vulnerability:** Found multiple hardcoded secrets in `cmd/api/main.go` and `internal/auth/auth.go`. The one in `main.go` was unused.
**Learning:** Secrets were hardcoded presumably for ease of development. `init()` function with fallback allows fixing this while maintaining dev convenience.
**Prevention:** Always use environment variables for secrets. Use `os.Getenv` and log warnings if using defaults in dev.
