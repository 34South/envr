[![Go Report Card](https://goreportcard.com/badge/34South/envr)](https://goreportcard.com/report/34South/envr) [![Build Status](https://travis-ci.org/34South/envr.svg?branch=master)](https://travis-ci.org/34South/envr)

# envr

Check / set required env vars, optionally from a local `.env` file. Uses the [godotenv](https://github.com/joho/godotenv) by [John Barton](https://github.com/joho/).

## Purpose

A convenient way to set env vars with a local file, or defer to env vars already set by some other means, eg Heroku config.


## Installation

```bash
$ go get github.com/34South/envr
```


## Usage

Set up a new Envr and pass in the env vars your app expects to be set. 

```go
func init() {
	env := envr.New("myEnv", []string{
  		"SOME_VAR1",
  		"SOME_VAR2",
  		"SOME_APP_ID",
  		"SOME_API_KEY",
	}).Auto()
}
```

```Go
// Envr contains fields related to the environment vars 
type Envr struct {
	Ready        bool              `json:"ready"`           // Flag for the goodness
	Name         string            `json:"environmentName"` // name of environment
	Files        []string          `json:"configFiles"`     // files to read from, default .env
	RequiredVars []string          `json:"requiredVars"`    // the env vars we need
	ExistingVars []string          `json:"existingVars"`    // the env vars that are set
	MissingVars  []string          `json:"missingVars"`     // the env vars not set
	V            map[string]string `json:"values"`          // map of existing vars and values
	Status       string            `json:"status"`          // a message about current status
	Error        error             `json:"error"`           // error field, for easier method chaining
}
```

**Envr Methods**

**.New("envName", []string{"var1", "var2"}, "configFile1, configFile2...")**
* Initialises a new Envr value with (arbitrary) name "envName"
* Second argument is a []string containing names of all expected / required env vars
* Third argument is an optional list of files containing env var names and values. If not present defaults to `.env`.

The `.env` file can be formatted like this:

```
# Comments are allowed
VAR1="value one"
VAR2=1234
```

See [godotenv](https://github.com/joho/godotenv) for details.

## Methods

**.Auto()**
Quick way to do **.Clean.Fatal()** - sets all the expected vars found in `.env`, and exists if any are missing.

**.Passive()**
Set expected vars *ONLY* if they are *NOT* already set.

**.Clean()**
Set all expected vars without checking first.

**.Fatal()**
Exits if the environment is *NOT* `.Ready = true`

## Todo
* Check vars that exist in config but are MISSING in the expectations list
* Learn how to write tests
* Write tests
* Go sailing


## License
The [MIT](https://mit-license.org/) License (MIT)
Copyright © 2017 Mike Donnici

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
