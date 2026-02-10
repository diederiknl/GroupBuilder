## 2024-05-22 - [Hardcoded JWT Secret in Go]
**Vulnerability:** Found hardcoded secrets `var jwtKey = []byte("neinneinnein")` in `main.go` and `var jwtKey = []byte("your_secret_key")` in `auth.go`.
**Learning:** Hardcoded secrets often appear in boilerplate or tutorial code and get carried over to production. Separate `jwtKey` variables in different files can lead to confusion and security gaps.
**Prevention:** Use `os.Getenv` strictly for secrets. Enforce environment variable checks at startup or at usage point (like `getJwtKey` helper).
