package platlogger

import (
	"cloud.google.com/go/logging"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ShareChat/service-template/config"
	"time"
)

const (
	logName = "plat-logger"
	productionEnv = "PRODUCTION"
	serviceNameLabel = "serviceName"
	deploymentIdLabel = "deploymentId"
	ConcurrentWriteLimit = 1 // TODO: Check the effect of increasing the parameter value
	DefaultDelayThreshold = 3 * time.Second // This can be changed as required
)

type Client struct {
	activeEnv string
	enableStackTrace bool
	standardLabel map[string]string
	nop bool
	outputPath struct{
		Console bool
		StackDriver bool
	}
	sdLogger *logging.Logger
	sdClient *logging.Client
}

func NewNop() *Client {
	return &Client{nop: true}
}

type LoggerOptions interface {
	set(*Client)
}

// set console output parameters that output logs entries to console as well
type consoleOutput bool

func ConsoleOutput(b bool) LoggerOptions {return consoleOutput(b)}

func (co consoleOutput) set(c *Client) {
	c.outputPath.Console = bool(co)
}

// set stack driver output parameters
type stackDriverOutput bool

func StackDriverOutput(b bool) LoggerOptions {return stackDriverOutput(b)}

func (sdo stackDriverOutput) set(c *Client) {
	c.outputPath.StackDriver = bool(sdo)
}

// set label parameters that will be attached to evey log entry
type standardLabel map[string]string

func StandardLabels(m map[string]string) LoggerOptions {return standardLabel(m)}

func (sl standardLabel) set(c *Client) {
	for key, val := range sl {
		c.standardLabel[key] = val
	}
}

// set parameter to enable stack trace
type enableStackTrace bool

func EnableStackTrace(b bool) LoggerOptions {return enableStackTrace(b)}

func (est enableStackTrace) set(c *Client) {
	c.enableStackTrace = bool(est)
}

func NewStagingClient(serviceName string) *Client {
	labels := map[string]string{serviceNameLabel: serviceName}
	return &Client{
		enableStackTrace: true,
		nop:              false,
		standardLabel: labels,
		outputPath: struct {
			Console     bool
			StackDriver bool
		}{Console:true, StackDriver:true},
	}
}

func NewProductionClient(serviceName string, config *config.Store) *Client {
	labels := map[string]string{
		serviceNameLabel: serviceName,
		deploymentIdLabel: config.Env.DeploymentId,
	}
	return &Client{
		enableStackTrace: false,
		nop:              false,
		standardLabel: labels,
		outputPath: struct {
			Console     bool
			StackDriver bool
		}{Console:false, StackDriver:true},
	}
}

func NewLogger(serviceName string, config *config.Store, opts ...LoggerOptions) (*Client, error){
	var client *Client
	if config.Env.ActiveEnv == productionEnv {
		client = NewProductionClient(serviceName, config)
	} else {
		client = NewStagingClient(serviceName)
	}
	client.activeEnv = config.Env.ActiveEnv
	sdClient, err := logging.NewClient(context.Background(), config.ProjectId)
	if err != nil {
		nopImpl := NewNop()
		return nopImpl, err
	}
	client.sdClient = sdClient

	// sets the optional config parameters
	for _, opt := range opts {
		opt.set(client)
	}
	sdLogger := sdClient.Logger(
		logName,
		logging.CommonLabels(client.standardLabel),
		logging.DelayThreshold(DefaultDelayThreshold),
		logging.ConcurrentWriteLimit(ConcurrentWriteLimit),
		)
	client.sdLogger = sdLogger
	return client, nil
}

func getErrorMsg(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%+v\n", err)
}

type JsonPayload struct {
	Message string `json:"message,omitempty"`
	Err string `json:"err,omitempty"`
	Parameters interface{} `json:"parameters,omitempty"`
}

func (c *Client) standardLog(logLevel logging.Severity, msg string, err error, parameters interface{}) {
	// return if no-op implementation
	if c.nop {
		return
	}
	logEntry := logging.Entry{
		Severity:       logLevel,
		Payload:        JsonPayload{
			Message:    msg,
			Err:        getErrorMsg(err),
			Parameters: parameters,
		},
	}
	if c.outputPath.StackDriver{
		c.sdLogger.Log(logEntry)
	}
	if c.outputPath.Console{
		params, _ := json.MarshalIndent(parameters, "", " ")
		fmt.Printf("msg: %s\n\nerr: %+v\n\nparams: %s\n", msg, err, params)
	}
}

func (c *Client) Debug(msg string, err  error, parameters interface{}) {
	c.standardLog(logging.Debug, msg, err, parameters)
}

func (c *Client) Info(msg string, err  error, parameters interface{}) {
	c.standardLog(logging.Info, msg, err, parameters)
}

func (c *Client) Notice(msg string, err  error, parameters interface{}) {
	c.standardLog(logging.Notice, msg, err, parameters)
}

func (c *Client) Warning(msg string, err  error, parameters interface{}) {
	c.standardLog(logging.Warning, msg, err, parameters)
}

func (c *Client) Error(msg string, err  error, parameters interface{}) {
	c.standardLog(logging.Error, msg, err, parameters)
}

func (c *Client) Critical(msg string, err  error, parameters interface{}) {
	c.standardLog(logging.Critical, msg, err, parameters)
}

func (c *Client) Alert(msg string, err  error, parameters interface{}) {
	c.standardLog(logging.Alert, msg, err, parameters)
}

func (c *Client) Emergency(msg string, err  error, parameters interface{}) {
	c.standardLog(logging.Emergency, msg, err, parameters)
}

func (c *Client) Close() {
	c.sdClient.Close()
}
