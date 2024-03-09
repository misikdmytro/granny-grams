package model

import "fmt"

type RESTClientError struct {
	Message  string
	Response any
}

func (e *RESTClientError) Error() string {
	return fmt.Sprintf("REST client error: %s. Response: %v", e.Message, e.Response)
}
