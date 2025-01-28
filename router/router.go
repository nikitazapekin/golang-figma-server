
package router

import (
    "github.com/gorilla/mux"
 
)

func InitRoutes(r *mux.Router) {
	SetAuthRoutes(r);
	SetDraftsRoutes(r)
}
 