package envr

import (
	"os"
	"fmt"
	"testing"
)

const name = "testenv"
const var1 = "ENVR_VAR1"
const var2 = "ENVR_VAR2"
const var3 = "ENVR_VAR3"

// newTestEnvr() unsets the env vars and returns a fresh *Envr
func newTestEnvr() *Envr {
	os.Unsetenv("ENV_VAR1")
	os.Unsetenv("ENV_VAR2")
	os.Unsetenv("ENV_VAR3")
	return New(name, []string{var1, var2, var3})
}

// TestNew checks Envr field values
func TestNew(t *testing.T) {

	e := newTestEnvr()

	// Check field values
	if e.Name != name {
		msg := fmt.Sprintf("Expected %S, got %S", name, e.Name)
		t.Error(msg)
	}

	if e.RequiredVars[0] != var1 {
		msg := fmt.Sprintf("Expected %S, got %S", var1, e.RequiredVars[0])
		t.Error(msg)
	}

	if e.RequiredVars[1] != var2 {
		msg := fmt.Sprintf("Expected %S, got %S", var2, e.RequiredVars[1])
		t.Error(msg)
	}

	if e.RequiredVars[2] != var3 {
		msg := fmt.Sprintf("Expected %S, got %S", var2, e.RequiredVars[2])
		t.Error(msg)
	}
}

// TestIsSet checks that each env var is set
func TestIsSet(t *testing.T) {

	e := newTestEnvr().Auto()

	if e.IsSet("ENVR_VAR1") == false {
		t.Error("Expected ENVR_VAR1 to be set")
	}

	if e.IsSet("ENVR_VAR2") == false {
		t.Error("Expected ENVR_VAR2 to be set")
	}

	if e.IsSet("ENVR_VAR3") == false {
		t.Error("Expected ENVR_VAR3 to be set")
	}

	if e.IsSet("ENVR_VAR_FALSE") == true {
		t.Error("ENVR_VAR_FALSE to be NOT set")
	}
}

// TestClean checks the .Clean() method.
func TestClean(t *testing.T) {

	fmt.Println("TestClean() -----------------------------------------")

	// New blank envr
	e := newTestEnvr()

	// Set env vars prior to .Clean()
	if err := os.Setenv("ENVR_VAR1", "ABCD"); err != nil {
		t.Error("Could not set env var")
	}
	if err := os.Setenv("ENVR_VAR2", "EFGH"); err != nil {
		t.Error("Could not set env var")
	}
	if err := os.Setenv("ENVR_VAR3", "IJKL"); err != nil {
		t.Error("Could not set env var")
	}

	// Verify values
	v1 := os.Getenv("ENVR_VAR1")
	if v1 != "ABCD" {
		msg := fmt.Sprintf("Expected ENVR_VAR1 to be ABCD before call to .Clean(), got %s", v1)
		t.Error(msg)
	}
	v2 := os.Getenv("ENVR_VAR2")
	if v2 != "EFGH" {
		msg := fmt.Sprintf("Expected ENVR_VAR2 to be EFGH before call to .Clean(), got %s", v2)
		t.Error(msg)
	}
	v3 := os.Getenv("ENVR_VAR3")
	if v3 != "IJKL" {
		msg := fmt.Sprintf("Expected ENVR_VAR3 to be IJKL before call to .Clean(), got %s", v3)
		t.Error(msg)
	}

	// Set with .Clean() which should overwrite the values
	e.Clean()

	// Check them again, should be new
	v4 := os.Getenv("ENVR_VAR1")
	if v4 != var1 {
		msg := fmt.Sprintf("Expected ENVR_VAR1 to be ENVR_VAR1 before call to .Clean(), got %s", v4)
		t.Error(msg)
	}
	v5 := os.Getenv("ENVR_VAR2")
	if v5 != var2 {
		msg := fmt.Sprintf("Expected ENVR_VAR2 to be ENVR_VAR2 before call to .Clean(), got %s", v5)
		t.Error(msg)
	}
	v6 := os.Getenv("ENVR_VAR3")
	if v6 != var3 {
		msg := fmt.Sprintf("Expected ENVR_VAR3 to be ENVR_VAR3 before call to .Clean(), got %s", v6)
		t.Error(msg)
	}
}

// TestPassive checks the .Passive() method which should NOT overwrite env values that
// are already set.
func TestPassive(t *testing.T) {

	fmt.Println("TestPassive() -----------------------------------------")

	// New blank envr
	e := newTestEnvr()

	// Set env vars prior to .Passive()
	if err := os.Setenv("ENVR_VAR1", "ABCD"); err != nil {
		t.Error("Could not set env var")
	}
	if err := os.Setenv("ENVR_VAR2", "EFGH"); err != nil {
		t.Error("Could not set env var")
	}
	if err := os.Setenv("ENVR_VAR3", "IJKL"); err != nil {
		t.Error("Could not set env var")
	}

	// Verify values
	v1 := os.Getenv("ENVR_VAR1")
	if v1 != "ABCD" {
		msg := fmt.Sprintf("Expected ENVR_VAR1 to be ABCD before call to .Clean(), got %s", v1)
		t.Error(msg)
	}
	v2 := os.Getenv("ENVR_VAR2")
	if v2 != "EFGH" {
		msg := fmt.Sprintf("Expected ENVR_VAR2 to be EFGH before call to .Clean(), got %s", v2)
		t.Error(msg)
	}
	v3 := os.Getenv("ENVR_VAR3")
	if v3 != "IJKL" {
		msg := fmt.Sprintf("Expected ENVR_VAR3 to be IJKL before call to .Clean(), got %s", v3)
		t.Error(msg)
	}

	// Set with .Clean() which should overwrite the values
	e.Passive()

	// Check them again, should be NO CHANGE
	v4 := os.Getenv("ENVR_VAR1")
	if v4 != "ABCD" {
		msg := fmt.Sprintf("Expected ENVR_VAR1 to be unchanged after call to .Passive(), got %s", v4)
		t.Error(msg)
	}
	v5 := os.Getenv("ENVR_VAR2")
	if v5 != "EFGH" {
		msg := fmt.Sprintf("Expected ENVR_VAR2 to be unchanged after call to .Passive(), got %s", v5)
		t.Error(msg)
	}
	v6 := os.Getenv("ENVR_VAR3")
	if v6 != "IJKL" {
		msg := fmt.Sprintf("Expected ENVR_VAR3 to be unchanged after call to .Passive(), got %s", v6)
		t.Error(msg)
	}
}

