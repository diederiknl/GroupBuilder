## 2025-05-23 - Hardcoded JWT Secret
**Vulnerability:** Hardcoded JWT signing key found in source code.
**Learning:** Legacy development keys were left in production code.
**Prevention:** Use environment variables for secrets and fail or warn if defaults are used.
