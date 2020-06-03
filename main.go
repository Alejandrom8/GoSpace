package gospace

import (
	"gospace/src/entities"
	"gospace/src/types"
	"net/http"
)

/*GlobalHandler - contiene todas las rutas designadas por gospace
sirve como manejador global de las rutas*/
var GlobalHandler *entities.Handler = &entities.Handler{
	Router: make(map[string]map[string]types.MiddleWare),
}

/*CreateServer - crea y configura una instancia de Indian
@param {string} port - el puerto donde escuchará el servidor, ejemplo: ":8080"
@return {*Indian} una instancia de Indian
*/
func CreateServer(port string) *entities.Indian {
	return &entities.Indian{
		Port: port,
		Server: &http.Server{
			Addr:    port,
			Handler: GlobalHandler,
		},
		Handler: GlobalHandler,
	}
}
