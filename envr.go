package envr

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Envr contains info about the environment setup
type Envr struct {
	Ready        bool              `json:"ready"`           // Flag for the goodness
	Name         string            `json:"environmentName"` // envName of environment
	Files        []string          `json:"configFiles"`     // files to read from, default .env
	RequiredVars []string          `json:"requiredVars"`    // the env vars we need
	ExistingVars []string          `json:"existingVars"`    // the env vars that are set
	MissingVars  []string          `json:"missingVars"`     // the env vars not set
	V            map[string]string `json:"values"`          // map of existing vars and values
	Status       string            `json:"status"`          // a message about current statuss
	Error        error             `json:"error"`           // error field, for easier method chaining
}

// New sets up a new Environment. It takes an arbitrary envName (n), a list of required vars (vs) and
// zero or more file names from which to read the vars, eg ".env1,.env2". Defaults to .env
func New(n string, vs []string, f ...string) *Envr {

	e := Envr{}
	e.Ready = false
	e.Name = n
	// if no env file(s) are specified, default to a single .env
	if len(f) == 0 {
		e.Files = append(e.Files, ".env")
	}

	// Initialise the values map
	e.V = make(map[string]string)

	// Set the Required vars here only
	for _, v := range vs {
		e.RequiredVars = append(e.RequiredVars, v)
	}

	// Set values for what exists, and what is missing
	e.Update()

	return &e
}

// Update sets / updates fields in the Envr value
func (e *Envr) Update() {

	// Empty out first
	e.ExistingVars = []string{}
	e.MissingVars = []string{}

	// Set the required vars
	for _, v := range e.RequiredVars {
		if e.IsSet(v) {
			e.ExistingVars = append(e.ExistingVars, v)
			e.V[v] = os.Getenv(v)
		} else {
			e.MissingVars = append(e.MissingVars, v)
		}
	}

	// Ready?
	if len(e.MissingVars) == 0 {
		e.Ready = true
	}

	// Status message
	if e.Ready {
		e.Status = "All required environment vars are present"
	} else {
		mv := strings.Replace(strings.Trim(fmt.Sprintf("%v", e.MissingVars), "[]"), " ", ", ", -1)
		e.Status = fmt.Sprintf("no config found for %v", mv)
	}
}

// Auto does Clean().Fatal() so will force the setting of all the
// required vars from the config, and die if things didn't workout
func (e *Envr) Auto() *Envr {
	return e.Clean().Fatal()
}

// Passive checks if vars are already set, and only sets them if they are not.
func (e *Envr) Passive() *Envr {

	// Reads env vars into a map without setting to env
	ev, err := godotenv.Read(e.Files...)
	if err != nil {
		e.Error = err
		return e
	}

	for _, v := range e.RequiredVars {
		_, exists := os.LookupEnv(v)
		if !exists {
			e.SetVar(v, ev[v])
		}
	}

	return e
}

// Clean sets all vars present in the config to the env, overriding existing values
func (e *Envr) Clean() *Envr {
	ev, err := godotenv.Read(e.Files...)
	if err != nil {
		e.Error = err
		return e
	}
	e.SetList(e.RequiredVars, ev)
	return e
}

// Fatal  is chained on so we can log fatal in the event our environment is not set up properly
func (e *Envr) Fatal() *Envr {
	if e.Ready == false {
		log.Fatalf("Envr: Bailing out, as requested - %s\n", e.Status)
	}
	return e
}

// IsSet returns true if an env var is sey, even if it is blank
func (e *Envr) IsSet(v string) bool {
	_, exists := os.LookupEnv(v)
	return exists
}

// SetList just sets sets the required vars to the environment
func (e *Envr) SetList(lv []string, ev map[string]string) *Envr {
	for _, v := range e.RequiredVars {
		if val, ok := ev[v]; ok {
			err := e.SetVar(v, val)
			if err != nil {
				e.Error = err
				return e
			}
		}
	}
	return e
}

// SetVar sets env var 'v' to value 's', if successful it updates Envr
func (e *Envr) SetVar(v, s string) error {
	err := os.Setenv(v, s)
	if err != nil {
		return err
	}
	e.Update()
	return nil
}

// TODO: A method to check for vars that are INV the env file but NOT in the list of expected vars
