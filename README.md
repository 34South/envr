## Environment Loader

Envr is a Go package for setting up environment vars.

It works, however I am a pretty average programmer at best, and new to Go.

The idea is to declare the environment variables required and then Envr will
check and load them from a *.env* file.

Like this:

```go
env := envr.New("myEnv", []string{
  "MONGO_URL",
  "MONGO_DB",
  "MONGO_LINKS_COLLECTION",
  "MONGO_STATS_COLLECTION",
}).Passive().Fatal()
```

Envr struct:

```Go
type Envr struct {
	Ready        bool              `json:"ready"`           // Flag for the goodness
	Name         string            `json:"environmentName"` // name of environment
	Files        []string          `json:"configFiles"`     // files to read from, default .env
	RequiredVars []string          `json:"requiredVars"`    // the env vars we need
	ExistingVars []string          `json:"existingVars"`    // the env vars that are set
	MissingVars  []string          `json:"missingVars"`     // the env vars not set
	V            map[string]string `json:"values"`          // map of existing vars and values
	Status       string            `json:"status"`          // a message about current statuss
	Error        error             `json:"error"`           // error field, for easier method chaining
}
```

##### New("someName", []string{"var1", "var2",}, "configFile1, configFile2...")
Set up a new environment, expectated vars and one or more config files. If omitted
config files will default to *.env*.

##### .Auto()
Quick way to do .Clean.Fatal(), that is, will set all the expected vars that can
be found in *.env*, and do log.Fatal() if any are missing.

##### .Passive()
Set expected vars, but only if they are NOT already set.

##### .Clean()
Set all expected vars without checking.

##### .Fatal()
Do a log.Fatal() if the environment is not all ok.

##### .JSON()
Print a JSON version of the Envr value.


This example could use an .env to run things locally, but when deploy to Heroku
(for example), it will load the vars from the Heroku environment.


#### Todo
* Make it better
* Check vars that exist in config but are MISSING in the expectatons list
* Learn how to write tests
* Write tests
* Proper docs
* Examples
* Go sailing

... it's early days... I will keep tweaking it.

#### Credit

It uses [godotenv](https://github.com/joho/godotenv) package to load *.env*  
- skills, @joho
