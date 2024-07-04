package errorMessage

const (
	// Other
	InternalServer    = "internal_server_error"
	NoAuthorization   = "no_authorization_header_found"
	NotHaveAccess     = "not_have_access"
	InvalidSortFormat = "invalid_sort_format"
	InvalidSortOrder  = "invalid_sort_order"

	// Authentication
	TokenInvalid         = "token_invalid"
	GetPayloadFailed     = "get_payload_from_token_failed"
	EmailPasswordInvalid = "email_password_invalid"

	// User
	UserNotFound = "user_not_found"

	// Medication
	MedicationAlreadyExists = "medication_already_exists"
	MedicationNotFound      = "medication_not_found"
	DuplicateMedicationId   = "duplicate_medication_id"
)
