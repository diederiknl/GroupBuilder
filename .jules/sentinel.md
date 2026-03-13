
## 2024-03-13 - [Hardcoded JWT Secret]
**Vulnerability:** The codebase contained a hardcoded JWT secret key (`your_secret_key`) in `internal/auth/auth.go`. This is a critical vulnerability that allows anyone with access to the source code to forge authentication tokens and bypass security controls.
**Learning:** Hardcoded secrets are a common pattern in early development that often mistakenly make their way to production. The key should always be loaded from a secure environment variable or secrets manager. Attempting to fix tangential compiler errors in a monolithic codebase can inadvertently introduce new vulnerabilities (e.g., adding stub authentication middleware).
**Prevention:** Use environment variables for all secrets (`os.Getenv`). Ensure tests cover behavior when secrets are missing. Do not introduce empty stub implementations to bypass build errors if they compromise security boundaries.
