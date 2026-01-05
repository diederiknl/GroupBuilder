package auth

import (
	"os"
	"os/exec"
	"testing"
)

// TestInitEnvVar verifies that the init function correctly respects the JWT_SECRET environment variable.
// Since init() runs at package initialization, we run this test in a subprocess.
func TestInitEnvVar(t *testing.T) {
	if os.Getenv("TEST_subprocess") == "1" {
		// Inside the subprocess
		// Just accessing the package is enough to trigger init()
		// But we need to verify the jwtKey variable.
		// Since jwtKey is private (unexported), we cannot access it directly from outside the package.
		// However, this test file is in package `auth`, so we CAN access `jwtKey`.

		expected := os.Getenv("EXPECTED_KEY")
		if string(jwtKey) != expected {
			t.Fatalf("Expected key %q, got %q", expected, string(jwtKey))
		}
		return
	}

	// Case 1: JWT_SECRET is set
	cmd := exec.Command(os.Args[0], "-test.run=TestInitEnvVar")
	cmd.Env = append(os.Environ(), "TEST_subprocess=1", "JWT_SECRET=my_custom_secret", "EXPECTED_KEY=my_custom_secret")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Test failed with custom secret: %v\nOutput: %s", err, out)
	}

	// Case 2: JWT_SECRET is not set (development default)
	cmd = exec.Command(os.Args[0], "-test.run=TestInitEnvVar")
	cmd.Env = append(os.Environ(), "TEST_subprocess=1", "JWT_SECRET=", "EXPECTED_KEY=dev_secret_do_not_use_in_prod")
	// Ensure GO_ENV is not production
	cmd.Env = append(cmd.Env, "GO_ENV=development")

	out, err = cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Test failed with default secret: %v\nOutput: %s", err, out)
	}
}
