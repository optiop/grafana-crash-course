package errors

import "fmt"

// internal errors
type Error struct {
	Code    string
	Message string
}

var (
	// http errors
	ErrHttpParsBody = &Error{
		Code:    "ErrHttpParsBody",
		Message: "error to parse body",
	}
	ErrHttpInvalidUrl = &Error{
		Code:    "ErrHttpInvalidUrl",
		Message: "error parse url",
	}
	ErrHttpCreateNewReq = &Error{
		Code:    "ErrHttpCreateNewReq",
		Message: "error to create http req",
	}
	ErrHttpNotValidStatus = &Error{
		Code:    "ErrHttpNotValidStatus",
		Message: "error not valid status code",
	}

	// docker errors
	ErrDockerGetAllContainer = &Error{
		Code:    "ErrDockerGetAllContainer",
		Message: "error get all container docker",
	}
	ErrDockerClientCreation = &Error{
		Code:    "ErrDockerClientCreation",
		Message: "error creating Docker client",
	}

	// Prometheus errors
	ErrPrometheusCreateClient = &Error{
		Code:    "ErrPrometheusCreateClient",
		Message: "error to create new client prometheus",
	}
	ErrPrometheusGetAllTargets = &Error{
		Code:    "ErrPrometheusGetAllTargets",
		Message: "error to get all target prometheus",
	}
)

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// fail test error message
type FailError struct {
	Code    string
	Message string
}

var (
	// http fail test
	FailHttpUnreachable = &FailError{
		Code:    "FailHttpUnreachable",
		Message: "fail unreachable address",
	}

	// docker fail test
	FailDockerContainerNotFound = &FailError{
		Code:    "FailDockerContainerNotFound",
		Message: "fail docker container not found",
	}

	// Prometheus fail test
	FailPrometheusNotFoundTarget = &FailError{
		Code:    "FailPrometheusNotFoundTarget",
		Message: "fail to find prometheus target",
	}

	// loki fail errors test
	FailLokiConfigurationNotFound = &FailError{
		Code:    "FailLokiConfigurationNotFound",
		Message: "fail to get loki Configuration",
	}

	FailLokiDatasourceNotFound = &FailError{
		Code:    "FailLokiDatasourceNotFound",
		Message: "fail to find loki datasource",
	}

	FailLokiLogsNotFound = &FailError{
		Code:    "LokiLogsNotFound",
		Message: "fail to find loki log entries",
	}
)

func (fe *FailError) Error() string {
	return fmt.Sprintf("%s: %s", fe.Code, fe.Message)
}

func isError(err error) bool {
	switch err.(type) {
	case *Error:
		return true
	default:
		return false
	}
}

func isFailError(err error) bool {
	switch err.(type) {
	case *FailError:
		return true
	default:
		return false
	}
}

func TypeErrorDetection(err error) string {
	if ok := isError(err); ok {
		return "error"
	}

	if ok := isFailError(err); ok {
		return "fail"
	}

	return ""
}
