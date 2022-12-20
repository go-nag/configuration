![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pavleprica/configuration)
![test build](https://github.com/pavleprica/configuration/actions/workflows/test.yml/badge.svg)
[![codecov](https://codecov.io/gh/go-nag/configuration/branch/master/graph/badge.svg?token=C2YYJG5U3C)](https://codecov.io/gh/go-nag/configuration)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/pavleprica/configuration)


# configuration
Configuration is used to effectively manage environment configurations within Golang. It offers an array of methods to fetch the environment in addition to load it from the .env file.

## Usage



### Bootstraping .env files

In the event that you want to use a .env file to store variables, you can do so be using the `conf_loader` package.


Example:
```go
// This will load all variables from .env.example into environment variables 
cfge.LoadEnvFile(".env.example")
```
or you can rely on the default one `.env`
```go
// This will load all variables from .env into environment variables
cfge.LoadDefaultEnvFile()
```

### GetEnv functions

The package `cfge` offers a set of methods for fetching environment variables. They are 
generally split into two 'types'. The first one is to fetch the method or return an error if the variable
is not there or the parsing was unsuccessful. The second one offers the default value that will be returned
in case of an error.

#### GetEnv

```go
// Fetches the environment variable named TEST_STR
variable, err := GetEnv("TEST_STR")
if err != nil {
	log.Fatalln(err)
}
// Prints the variable value
fmt.Println(variable)
```

#### GetEnvOrDefault

```go
// Fetches the environment variable named TEST_STR
// there is no error returned here, in case it fails it will provide the default value
variable := GetEnvOrDefault("TEST_STR", "default value")
// Prints the variable value
fmt.Println(variable)
```

#### GetEnvInt

```go
// Fetches the environment variable named TEST_INT
// in addition to that it parses it to int value
variable, err := GetEnvInt("TEST_INT")
if err != nil {
	log.Fatalln(err)
}
// Prints the variable value
fmt.Println(variable)
```

#### GetEnvIntOrDefault

```go
// Fetches the environment variable named TEST_INT
// There is no error returned here, in case it fails it will provide the default value
// in addition to that it parses it to int value
variable := GetEnvIntOrDefault("TEST_INT", 5)
// Prints the variable value
fmt.Println(variable)
```

#### GetEnvBool

```go
// Fetches the environment variable named TEST_BOOL
// in addition to that it parses it to bool value
variable, err := GetEnvBool("TEST_BOOL")
if err != nil {
	log.Fatalln(err)
}
// Prints the variable value
fmt.Println(variable)
```

#### GetEnvBoolOrDefault

```go
// Fetches the environment variable named TEST_BOOL
// There is no error returned here, in case it fails it will provide the default value
// in addition to that it parses it to bool value
variable := GetEnvBoolOrDefault("TEST_BOOL", false)
// Prints the variable value
fmt.Println(variable)
```

## Loading configuration from environment files

_Currently only supporting yaml files_

### yaml

The configuration for yaml files is divided between environments that your app runs in. 
When loading up the configuration, the method expects a `environment string` parameter which
indicates which file to choose. Let's say you have `local`, `dev`, `prod` as running environments.
Following that you would have `config-local.yaml`, `config-dev.yaml` and `config-prod.yaml` files with
different value configurations. See this [issue](https://github.com/go-nag/configuration/issues/1) for more information.

In addition to that, because of sensitive values, and values changing per environment you can supply a template holder
that will load up the value from the environment. For example `database_pw: ${DATABASE_PASSWORD}` will take the environment
value of `DATABASE_PASSWORD` from your system. Making it convenient for deployments.

#### Example files

Example file for **local** config:
```yaml
database:
  host: http://localhost:5042
  username: user
  password: my-secret-pw

kafka:
  url: http://localhost:5555
  clientId: localApp

something: wow

number: 7000

boolean: true
```

Example file for **dev** config:
```yaml
database:
  host: http://remote-database:5042
  username: ${DATABASE_USERNAME}
  password: ${DATABASE_PASSWORD}

kafka:
  url: http://remote-kafka:5555
  clientId: ${KAFKA_CLIENT_ID}

something: wow

number: 7000

boolean: true
```
_In this example, the `${}` template values will be loaded from system environment._

#### Using the loader

`cfgm.LoadConfigFile()`

To use the loader, just invoke `cfgm.LoadConfigFile("local")` or in dev case `cfgm.LoadConfigFile("dev")`.
It in turn will return the `cfgm.Manager` [type](cfgm/manager.go). Which offers functions to get values. 
In a bigger context an example would be:

```go
func main() {
	manager, err := cfgm.LoadConfigFile("local")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	// Single value
	loggerEnabled, err := manager.Get("server.logging")
	if err != nil {
		log.Fatal(err)
	}

	if loggerEnabled == "enabled" {
		log.Println("Using logger")
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	manager.GetOrDefault("port", "9000")

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", manager.GetOrDefault("port", "9000"))))
}
```
With the config file being `config-local.yaml`
```yaml
port: 8080
server:
  logging: enabled
```

**Arrays** can be loaded as well.
```go
func main() {
    manager, err := cfgm.LoadConfigFile("local")
    if err != nil {
    log.Fatal(err)
    }
    
    // Loads array, if encounters error, returns empty array.
    arrayConfig := manager.GetArr("array.value")
    
}
```

`cfgm.LoadConfigFileWithPath()`

This loader is doing the same job as `cfgm.LoadConfigFile()` but it doesn't look via the `environment` variable, rather just takes
the custom path provided.
It in turn will return the `cfgm.Manager` [type](cfgm/manager.go). Which offers functions to get values.
In a bigger context an example would be:

```go
func main() {
	manager, err := cfgm.LoadConfigFileWithPath("/Users/someuser/yourproject/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	// Single value
	loggerEnabled, err := manager.Get("server.logging")
	if err != nil {
		log.Fatal(err)
	}

	if loggerEnabled == "enabled" {
		log.Println("Using logger")
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	manager.GetOrDefault("port", "9000")

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", manager.GetOrDefault("port", "9000"))))
}
```

Example project code can be found [here](https://github.com/go-nag/configuration-example).