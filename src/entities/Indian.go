package entities

import (
	"fmt"
	typedef "gospace/src/types"
	"net/http"
)

/*Indian - estructura que funge como el principal manejador del server.
Sirve como una colección de rutas y servidor. Se debe referenciar a esta
estructura siempre que se quiera interactuar con el servidor.
@prop {string} Port - el puerto en el que iniciara el servidor.
@prop {*http.Server} Server - el servidor de la aplicación.
@prop {*Handler} Handler - el manejador de peticiones y rutas del servidor.
*/
type Indian struct {
	Port    string
	Server  *http.Server
	Handler *Handler
}

//HTTP METHODS

/*Get - adds an HTTP get method to the router*/
func (i *Indian) Get(path string, handler typedef.MiddleWare) {
	i.Handler.AddRoute(path, "GET", handler)
}

/*Post - adds an HTTP post method to the router*/
func (i *Indian) Post(path string, handler typedef.MiddleWare) {
	i.Handler.AddRoute(path, "POST", handler)
}

//Controllers and other funcions

/*Use - register a middleware*/
// func (i *Indian) Use(middleware func(http.ResponseWriter, *http.Request, ...typedef.MiddleWare)) {

// }

/*Start - inicia el servidor
 */
func (i *Indian) Start() {
	fmt.Printf("Server ready and listening at port: %s\n", i.Port)
	i.Server.ListenAndServe()
}
