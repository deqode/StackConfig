# What is StackConfig

## About
We have defined a universal backend application configuration standard. An application might consist of a single or multiple services, an array of such services is stored in the application configuration. Each service contains details such as service runtime, databse, CPU, memory, network, scaling and other general settings including environment variables.
Using this Golang library users can Validate, Store, Get and maintain multiple versions of an application configuration in their own databse. User can use this library with any databse supportd by Gokv library. Interface for Stack defines all these implemented methods.
## Where can be used?
- to generate cloud infrastructure template for an app

## Usage
go get https://github.com/deqode/StackConfig

### Example with leveldb
```go
import (
    "github.com/deqode/StackConfig/stackconfig"
    "github.com/philippgille/gokv/leveldb"
    "go.uber.org/zap"
)

function main() {
    options := leveldb.DefaultOptions
    store, err := leveldb.NewStore(options)
    if err != nil {
        panic(err)
    }
    logger := zap.NewExample()
    stackService := stackconfig.StackService{
        Store:  store,
        Logger: logger,
    }
}
```

### Example with postgresql
```go
import (
    "github.com/deqode/StackConfig/stackconfig"
    "github.com/philippgille/gokv/postgresql"
    "go.uber.org/zap"
)

function main() {
    options := postgresql.DefaultOptions
    store, err := postgresql.NewClient(options)
    if err != nil {
        panic(err)
    }
    logger := zap.NewExample()
    stackService := stackconfig.StackService{
        Store:  store,
        Logger: logger,
    }
}
```
### Example with s3
```go
import (
    "github.com/deqode/StackConfig/stackconfig"
    "github.com/philippgille/gokv/s3"
    "go.uber.org/zap"
)

function main() {
    options := s3.DefaultOptions
    store, err := s3.NewClient(options)
    if err != nil {
        panic(err)
    }
    logger := zap.NewExample()
    stackService := stackconfig.StackService{
        Store:  store,
        Logger: logger,
    }
}
```
### App service Interface
```go
type Stack Interface {
    // used to validate an app config
    ValidateAppConfig(app *pb.StackConfig) error
    // save latest version of app config to key-value store
    Save(app *pb.StackConfig) (*pb.StackConfig, error)
    // get latest version of app config from key-value store
    GetAppConfig(id string) (*pb.StackConfig, error)
    // get app config corresponding to any available version from key-value store
    GetAppConfigForVersion(id string, version int32) (*pb.StackConfig, error)
}
```



