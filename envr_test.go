package envr

import (
	"log"
	"os"
	"testing"
)

const envName = "testenv"

var testVars = []string{"ENVR_VAR1", "ENVR_VAR2", "ENVR_VAR3"}

// newTestEnvr() unsets the env vars and returns a fresh *Envr
func newTestEnvr() *Envr {
	for _, v := range testVars {
		if err := os.Unsetenv(v); err != nil {
			log.Fatalf("os.Unsetenv() err = %s", err)
		}
	}
	return New(envName, testVars)
}

// TestNew checks Envr field values
func TestNew(t *testing.T) {
	e := newTestEnvr()
	if e.Name != envName {
		t.Errorf(".Name = %s, want %s", e.Name, envName)
	}
	for i, v := range testVars {
		if e.RequiredVars[i] != v {
			t.Errorf("RequiredVars[%d] = %s, want %s", i, e.RequiredVars[i], v)
		}
	}
}

// TestIsSet runs .Clean() and checks that env vars have been set correctly
func TestIsSet(t *testing.T) {
	e := newTestEnvr().Clean()
	for _, v := range testVars {
		if !e.IsSet(v) {
			t.Errorf("IsSet(%q) = false, want true", v)
		}
	}
}

// TestClean verifies that .Clean() method will overwrite env vars with new values
func TestClean(t *testing.T) {
	e := newTestEnvr()
	// set env vars to new values
	newValues := []string{"CAT", "DOG", "BIRD"}
	for i, v := range testVars {
		if err := os.Setenv(v, newValues[i]); err != nil {
			t.Fatalf("os.Setenv(%q, %q) err = %s", v, newValues[i], err)
		}
	}
	// Verify values
	for i, v := range testVars {
		want := newValues[i]
		got := os.Getenv(v)
		if got != want {
			t.Fatalf("os.GetEnv(%q) = %q, want %q", v, got, want)
		}
	}
	// Set with .Clean() which should overwrite the values
	e.Clean()
	// Check them again, should be changed
	for i, v := range testVars {
		if os.Getenv(v) == newValues[i] {
			t.Errorf("os.GetEnv(%q) = %q, should have changed after .Clean() to old value in .env", v, os.Getenv(v))
		}
	}
}

// TestPassive checks the .Passive() method which should NOT overwrite env values that are already set.
func TestPassive(t *testing.T) {
	e := newTestEnvr()
	// set env vars to new values
	newValues := []string{"DONKEY", "CAMEL", "SLOTH"}
	for i, v := range testVars {
		if err := os.Setenv(v, newValues[i]); err != nil {
			t.Fatalf("os.Setenv(%q, %q) err = %s", v, newValues[i], err)
		}
	}
	// Verify values
	for i, v := range testVars {
		want := newValues[i]
		got := os.Getenv(v)
		if got != want {
			t.Fatalf("os.GetEnv(%q) = %q, want %q", v, got, want)
		}
	}
	// Set with .Passive() which should NOT alter any values
	t.Log(os.Getenv("ENVR_VAR1"))
	e.Passive()
	t.Log(os.Getenv("ENVR_VAR1"))
	for i, v := range testVars {
		want := newValues[i]
		got := os.Getenv(v)
		if got != want {
			t.Fatalf("os.GetEnv(%q) = %q, want %q", v, got, want)
		}
	}
}
