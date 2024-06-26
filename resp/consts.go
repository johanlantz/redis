package resp

const (
	DT_SIMPLE_STRING    = '+'
	DT_SIMPLE_ERROR     = '-'
	DT_INTEGER          = ':'
	DT_BULK_STRINGS     = '$'
	DT_ARRAYS           = '*'
	DT_NULLS            = '_'
	DT_BOOLEANS         = '#'
	DT_DOUBLES          = ','
	DT_BIG_NUMBERS      = '('
	DT_BULK_ERRORS      = '|'
	DT_VERBATIM_STRINGS = '='
	DT_MAPS             = '%'
	DT_SETS             = '~'
	DT_PUSHES           = '>'
)

const (
	RESP_OK  = "OK"
	RESP_ERR = "ERR"
)

const (
	RESP_GET  RespCommand = "GET"
	RESP_SET  RespCommand = "SET"
	RESP_INCR RespCommand = "INCR"
	RESP_DEL  RespCommand = "DEL"
)
