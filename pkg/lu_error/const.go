package lue

const (
	InternalServerError = iota
	BadRequest
	ErrorMigrate
	ErrorConnectDb
	ErrorReader
	ErrorWriter
	ErrorDecoder
	ErrorEncode
	ErrorCreate
	ErrorUpdate
	ErrorDelete
	ErrorDatabaseQuery
	ErrorDatabaseArgument
	Unauthorized

	InternalServerErrorText   = "Internal Server Error"
	BadRequestText            = "Bad Request"
	ErrorMigrateText          = "Error Migrate"
	ErrorConnectDbText        = "Error Connect Database"
	ErrorReaderText           = "Error Reader"
	ErrorWriterText           = "Error Writer"
	ErrorDecoderText          = "Error Decoder"
	ErrorEncodeText           = "Error Encode"
	ErrorCreateText           = "Error Create"
	ErrorUpdateText           = "Error Update"
	ErrorDeleteText           = "Error Delete"
	ErrorDatabaseQueryText    = "Error Database Query"
	ErrorDatabaseArgumentText = "Error Database Argument"
	UnauthorizedText          = "Error Unauthorized access"
)

var (
	CodeInText = map[int]string{
		InternalServerError:   InternalServerErrorText,
		BadRequest:            BadRequestText,
		ErrorMigrate:          ErrorMigrateText,
		ErrorConnectDb:        ErrorConnectDbText,
		ErrorReader:           ErrorReaderText,
		ErrorWriter:           ErrorWriterText,
		ErrorDecoder:          ErrorDecoderText,
		ErrorEncode:           ErrorEncodeText,
		ErrorCreate:           ErrorCreateText,
		ErrorUpdate:           ErrorUpdateText,
		ErrorDelete:           ErrorDeleteText,
		ErrorDatabaseQuery:    ErrorDatabaseQueryText,
		ErrorDatabaseArgument: ErrorDatabaseArgumentText,
		Unauthorized:          UnauthorizedText,
	}
)
