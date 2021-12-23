package errno

// Model errors
var (
	OK = NewError(1, "success")
	ErrInvalidArgs  = NewError(1000, "Invalid Args")
	ErrKeyConflict  = NewError(1001, "Key Conflict")
	ErrDataNotFound = NewError(1002, "Record Not Found")
	ErrUserExists   = NewError(1003, "User already exists")
	ErrUnknown      = NewError(1004, "Unknown Error")
	ErrFailed       = NewError(1005, "Failed")
)
