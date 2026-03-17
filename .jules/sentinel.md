## 2024-05-15 - Hardcoded JWT Secret
**Vulnerability:** A hardcoded `jwtKey` was found in `internal/auth/auth.go`. This means that if an attacker were to get hold of the source code, they could easily generate valid JWTs for any user role.
**Learning:** Hardcoded secrets present a huge security vulnerability when exposed, allowing unauthorized access or generation of tokens.
**Prevention:** Always use environment variables, secure secret managers, or external config files that are not committed to source control for sensitive keys such as those used for signing JWTs.
