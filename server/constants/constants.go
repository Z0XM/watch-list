package constants

const (
	RequestID          = "X-Request-Id"
	JWTHeader          = "x-jwt-header"
	PrivateUserDetails = "private_user_details"
	Authorization      = "Authorization"
)

// Error constants
const (
	InvalidRequest = "Invalid Request"
	BadRequest     = "Bad Request"
)

// User constants
const (
	UserID   = "user_id"
	IsActive = "is_active"
)

// Pagination constants
const (
	PageNumber        = "page_number"
	PageSize          = "page_size"
	Offset            = "offset"
	Sort              = "sort"
	Ascending         = "asc"
	Descending        = "desc"
	DefaultPageNumber = 1
	DefaultPageSize   = 25
)
