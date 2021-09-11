package errors

const (
	RepositoryDownErr = Kind("Repository connection problem")
	RepositoryQueryErr = Kind("Failed to perform query")
	RepositoryNoRows = Kind("No rows were found")
	UnexpectedErr = Kind("Unexpected error occurred")
)
