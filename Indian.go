package gospace

import (
	"fmt"
	"net/http"
)

/*Handler - sirve como el manejador de middlewares y rutas del servidor
@prop {map[string]map[string]func(http.ResponseWriter, *http.Request)} Router -
una tabla hash con las rutas y manejadores de las rutas del servidor. Esta tabla
se debe componer de los siguientes datos:

	[ruta][metodo] funcion manejadora
*/
type Handler struct {
	Router map[string]map[string]func(http.ResponseWriter, *http.Request)
}

/*ServeHTTP - principal manejador de las peticiones
A través de esta función pasarán todas las peticiones de los usuarios.
@param {http.ResponseWriter} w - el writer de la petición.
@param {*http.Request} r - la petición.
*/
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if routePath, ok := h.Router[r.URL.String()]; ok { //si existe la ruta
		if hf, ok2 := routePath[r.Method]; ok2 { // y existe el metdodo para esa ruta
			hf(w, r) // se ejecuta su manejador
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - nothing here"))
}

/*AddRoute - añade una ruta a la tabla de Router.
@param {string} path - dirección a la que responderá el hadler
@param {string} method - metodo HTTP al que responderá el handler
@param {func(http.ResponseWriter, *http.Request)} handler - manejador de la ruta
@return {bool} - retorna verdadero si se añadio la ruta al Router.
*/
func (h *Handler) AddRoute(path string, method string, handler func(http.ResponseWriter, *http.Request)) bool {
	if _, pathExists := h.Router[path]; !pathExists { // si no existe la ruta
		// se crea el espacio en memoria para esta
		h.Router[path] = make(map[string]func(http.ResponseWriter, *http.Request))
	}

	if _, methodExists := h.Router[path][method]; !methodExists { //si no existe el método
		// se añade a la ruta
		h.Router[path][method] = handler
		return true
	}

	fmt.Printf("\nThere is a handler for path: %s and for method: GET", path)
	return false
}

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

/*Start - inicia el servidor
 */
func (i *Indian) Start() {
	fmt.Printf("Server ready and listening at port: %s", i.Port)
	i.Server.ListenAndServe()
}

/*CreateServer - crea y configura una instancia de Indian
@param {string} port - el puerto donde escuchará el servidor, ejemplo: ":8080"
@return {*Indian} una instancia de Indian
*/
func CreateServer(port string) *Indian {
	handler := &Handler{
		Router: make(map[string]map[string]func(http.ResponseWriter, *http.Request)),
	}

	return &Indian{
		Port: port,
		Server: &http.Server{
			Addr:    port,
			Handler: handler,
		},
		Handler: handler,
	}
}

// func main() {
// 	server := CreateServer(":8080")
// 	server.Handler.AddRoute("/", "GET", func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "Hello world")
// 	})
// 	server.Handler.AddRoute("/hola", "GET", func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "Hola x2")
// 	})
// 	server.Start()
// }
