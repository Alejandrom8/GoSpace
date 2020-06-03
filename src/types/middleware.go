package types

import "net/http"

/*MiddleWare - la estructura de un middleware*/
type MiddleWare func(http.ResponseWriter, *http.Request)
