package envr

import (
	"fmt"
	"os"
	"testing"
)

func TestNewEnv(t *testing.T) {

	// pre set VAR1
	// os.Setenv("ENVR_VAR1", "true")

	name := "testenv"
	var1 := "ENVR_VAR1"
	var2 := "ENVR_VAR2"
	var3 := "ENVR_VAR3"
	e := New(name, []string{var1, var2, var3})

	fmt.Println(e.JSON())

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
}

func TestCheck(t *testing.T) {

	e := Envr{
		Name: "testenv",
	}

	os.Setenv("ENVR_TEST_POSITIVE", "true")

	if e.IsSet("ENVR_TEST_POSITIVE") == false {
		t.Error("Expected ENVR_TEST_POSITIVE = true, got false")
	}

	if e.IsSet("ENVR_TEST_NEGATIVE") == true {
		t.Error("Expected ENVR_TEST_NEGATIVE = false, got true")
	}

}

func TestAuto(t *testing.T) {

	fmt.Println("TestAuto() -----------------------------------------")
	// Start fresh
	os.Unsetenv("ENV_VAR1")
	os.Unsetenv("ENV_VAR2")
	os.Unsetenv("ENV_VAR3")

	name := "testenv"
	var1 := "ENVR_VAR1"
	var2 := "ENVR_VAR2"
	var3 := "ENVR_VAR3"
	e := New(name, []string{var1, var2, var3})

	j, _ := e.JSON()
	fmt.Println("Before Auto():", j)
	e.Auto()
	j, _ = e.JSON()
	fmt.Println("After Auto():", j)

}
