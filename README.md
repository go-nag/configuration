![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pavleprica/configuration)
![test build](https://github.com/pavleprica/configuration/actions/workflows/test.yml/badge.svg)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/pavleprica/configuration)


# configuration
Configuration is used to effectively manage environment configurations within Golang. It offers an array of methods to fetch the environment in addition to load it from the .env file.

## Usage

### Bootstraping .env files

In the event that you want to use a .env file to store variables, you can do so be using the `conf_loader` package.


Example:
```go
// This will load all variables from .env.example into environment variables 
conf_loader.LoadEnvFile(".env.example")
```
or you can rely on the default one `.env`
```go
// This will load all variables from .env into environment variables
conf_loader.LoadDefaultEnvFile()
```

### GetEnv functions

The package `conf` offers a set of methods for fetching environment variables. They are 
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