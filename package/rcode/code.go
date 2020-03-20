package rcode

const (
	SUCCESS                 = 200
	ERROR                   = 500
	INVALID_PARAMS          = 400
	UNPASS                  = 403
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_EXIST_USER     = 30001
	ERROR_NOT_EXIST_USER = 30002

	ERROR_BIND_DATA   = 40001
	ERROR_UNFIND_DATA = 40004

	ERROR_IMAGE_BAD_EXT    = 50001
	ERROR_IMAGE_TOO_LARGE  = 50002
	ERROR_FILE_CREATE_FAIL = 50003
	ERROR_FILE_SAVE_FAIL   = 50004

	ERROR_SQL_INSERT_FAIL = 60001
	ERROR_SQL_DELETE_FAIL = 60002
	ERROR_SQL_UPDATE_FAIL = 60003
)
