### Description

The logger is a basic implementation that creates structured logs in the format required by StackDriver and pushes
these logs using standard cloud logging library. The aim is to generate logs in the proper format which can be easily
queried on StackDriver. The aim of the package is to give a decent logging framework to the services starting with Go.

More functionality will be added to this and it will be separated out as a separate module later.

#### Definition

```go
NewLogger(serviceName string, config *config.Store, opts ...LoggerOptions) (*Client, error)
```

#### Default Config
 - Uses default config based on the runtime environment
 - Default config is based on `environment(STAG/PROD)` and can be found in code
 - Contains two output paths by default `console` & `stackdriver`
 - Labels like `serviceName`, `deploymentId` are attached to log entry by default
 - If there is an error the logger returned is a no-op implementation. This is added so that in case of an error,
   application has the options to either return or continue with no-op implementation
```go
logger, err := platlogger.NewLogger("post-counter-service", config)
if err != nil {
    log.Println(err)
}
```


#### With Config
  - Allows to add config parameters to modify logger behaviour
  - These include
    - Enable/Disable console output behaviour
    - Enable/Disable stack driver output behaviour
    - Add some custom labels that are attached to every log entry
    ```go
    logger, err := platlogger.NewLogger(
        "post-counter-service",
        config, 
        platlogger.ConsoleOutput(false),
        platlogger.StackDriverOutput(true),
        platlogger.StandardLabels(map[string]string{"someLabelKey": "someLabelValue"}
    ))
    ```

#### Implementation

##### Methods:
```go
logger.<level>(msg string, err  error, parameters interface{})

level: info, error ....
```

```go
if err != nil {
    // passing parameters as map
    someMap := map[string]interface{}{
        "key1": "value1",
        "key2": map[string]int64{
            "nestedKey3": 6,
        },
    }
    logger.Error("PackageMessage: some error", err, someMap)
    
    // passing parameters as map
    someStruct := struct{
        Key1 string
        Key2 string
    }{Key1: "value1", key2: "value2"}
    logger.Error("PackageMessage: some error", err, someStruct)
}   
```
