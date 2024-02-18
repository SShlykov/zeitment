package config

type HTTPError struct {
	Message string `json:"message"`
}

type repoError struct {
	Error string `json:"error"`
}

var (
	ErrorNotFound   = HTTPError{"record not found"}
	ErrorBadInput   = HTTPError{"bad input"}
	ErrorForbidden  = HTTPError{"forbidden"}
	ErrorNotCreated = HTTPError{"not created"}
	ErrorNotUpdated = HTTPError{"not updated"}
)
