## 2024-05-23 - Hardcoded JWT Secret
**Vulnerability:** Found hardcoded secrets (`var jwtKey = []byte(...)`) in `internal/auth/auth.go` and `cmd/api/main.go`. This exposes the application to token forgery if the source code is leaked.
**Learning:** Hardcoded secrets often appear in initial development phases and are forgotten. Even unused secrets (like in `main.go`) can be confusing or dangerous if mistakenly used.
**Prevention:** Always use environment variables for secrets from the start. Use tools like `gosec` to scan for hardcoded credentials.
