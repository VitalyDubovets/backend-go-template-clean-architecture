package infraerrors

var (
	AppConfigError = &InfrastructureError{
		detail: "couldn't create an app config",
	}
	LogInitError = &InfrastructureError{
		detail: "couldn't create a logger",
	}
)
