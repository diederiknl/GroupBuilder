## 2024-10-18 - [Hardcoded JWT Secret]
**Vulnerability:** A hardcoded "your_secret_key" was found in `internal/auth/auth.go`, used for signing JWT tokens.
**Learning:** Hardcoded secrets in source code are a critical risk as they can be easily extracted if the code is leaked or accessed.
**Prevention:** Use environment variables (e.g., `JWT_SECRET`) to inject secrets at runtime. Implement a "fail secure" or warning mechanism if the secret is missing.
