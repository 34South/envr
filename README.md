# envr :earth_asia:

A Go package to check / set environment vars, optionally from a local *.env* file.

## Purpose
----------

Created for using a *.env* file in a local development set up, and deferring to env vars in testing/production - eg Heroku settings.


## Installation

```bash
$ go get github.com/34South/envr
```


## Usage

Chain the method calls in your init(), thus:

```go
func init() {
	env := envr.New("myEnv", []string{
  		"SOME_VAR1",
  		"SOME_VAR2",
  		"SOME_APP_ID",
  		"SOME_API_KEY",
	}).Passive().Fatal()
}
```

**Envr type**

The following struct is used: 

```Go
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
* Initialises a new Envr value with arbitrary name "envName"
* Second argument is a []string containing names of all expected / required env vars
* Third agument is an optional list of one or more files containing env var names and values. If not present it will look for a single file named *.env*

The *.env* file is formatted like this:

```
# Comments are allowed
VAR1="valueOne"
VAR2=1234
```

**.Auto()**
Quick way to do **.Clean.Fatal()** - sets all the expected vars that can
be found in *.env*, and do log.Fatal() if any are missing.

**.Passive()**
Set expected vars ONLY if they are NOT already set.

**.Clean()**
Set all expected vars without checking.

**.Fatal()**
Do log.Fatal() if the environment is not all ok.

**.JSON()**
Print a JSON version of the Envr value.



## Todo
* Make it better
* Check vars that exist in config but are MISSING in the expectatons list
* Learn how to write tests
* Write tests
* Go sailing

... it's early days... I will keep tweaking it.


## Credit
This is just a wrapper for [godotenv](https://github.com/joho/godotenv) package to load *.env*  
- skills, @joho


## License
The [MIT](https://mit-license.org/) License (MIT)
Copyright © 2017 Mike Donnici

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
