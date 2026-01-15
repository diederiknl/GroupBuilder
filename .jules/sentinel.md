## 2024-05-22 - [Hardcoded JWT Secret]
**Vulnerability:** A hardcoded secret key (`"your_secret_key"`) was found in `internal/auth/auth.go` for signing JWTs.
**Learning:** Hardcoded secrets often appear in initial prototypes or development phases and are forgotten. They are critical risks because source code exposure leads to full system compromise.
**Prevention:** Always initialize secrets from environment variables (e.g., `os.Getenv`) and never commit them to the repository. Use a fallback with a warning for local development only if necessary.
