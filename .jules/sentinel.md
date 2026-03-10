## 2024-03-10 - Hardcoded JWT Secret Remediation
**Vulnerability:** A hardcoded `jwtKey` secret (`your_secret_key`) was used globally in `internal/auth/auth.go` for signing and validating JWT tokens.
**Learning:** Hardcoded secrets present a major risk as anyone with access to the source code can forge tokens and bypass authentication.
**Prevention:** Rely on environment variables (like `JWT_SECRET`) loaded dynamically at runtime for symmetric keys to ensure secrets are managed securely outside of the source code repository. Ensure appropriate fallbacks or clear errors when the variable isn't set.
