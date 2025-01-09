package tserrors

// These are transaction-service specific errorCode
// This has nothing to do with http status codes
var (
	DBError             = Error{Code: 450, Message: "Database error"}
	InvalidRequestError = Error{Code: 451, Message: "Invalid request"}
	NotFoundError       = Error{Code: 452, Message: "Not found"}
	UnauthorizedError   = Error{Code: 453, Message: "Unauthorized"}
	ForbiddenError      = Error{Code: 454, Message: "Forbidden"}
)
