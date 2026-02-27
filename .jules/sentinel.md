## 2024-05-22 - [Defense in Depth: Validating External Data Structure]
**Vulnerability:** A panic (DoS) vulnerability existed in the CSV import feature where the code assumed all rows had at least 3 columns without validation, leading to an index out of range error on malformed input.
**Learning:** Assumptions about the structure of external data (like user-uploaded files) are dangerous. Even if the file type is correct (CSV), the internal structure (column count) must be validated before access.
**Prevention:** Always validate the length or structure of data structures (arrays, slices, maps) derived from external input before accessing them by index or key. Fail gracefully with an error message rather than allowing a runtime panic.
