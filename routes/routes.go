package routes

import (
	"github.com/gorilla/mux"
	"inventory-app/controllers"
	"inventory-app/middleware"
	"net/http"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session, _ := middleware.Store.Get(r, "session")
		if session.Values["authenticated"] == true {
			if session.Values["role"] == "admin" {
				http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
				return
			} else {
				http.Redirect(w, r, "/inventory", http.StatusSeeOther)
				return
			}
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	r.HandleFunc("/login", controllers.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/register", controllers.RegisterHandler).Methods("GET", "POST")
	r.HandleFunc("/logout", controllers.LogoutHandler)

	//	r.HandleFunc("/login", controllers.Login).Methods("GET", "POST")
	//	r.HandleFunc("/register", controllers.Register).Methods("GET", "POST")
	//r.HandleFunc("/logout", controllers.Logout)

	// Gunakan middleware admin & user
	r.Handle("/admin/dashboard", middleware.RequireAdmin(http.HandlerFunc(controllers.AdminDashboard)))
	r.Handle("/inventory", middleware.RequireLogin(http.HandlerFunc(controllers.InventoryPage)))
	r.Handle("/inventory/create", middleware.RequireAdmin(http.HandlerFunc(controllers.CreateItem)))
	r.Handle("/inventory/update/{id}", middleware.RequireAdmin(http.HandlerFunc(controllers.UpdateItem)))
	r.Handle("/inventory/delete/{id}", middleware.RequireAdmin(http.HandlerFunc(controllers.DeleteItem)))
	r.Handle("/inventory/export/pdf", middleware.RequireLogin(http.HandlerFunc(controllers.ExportPDF)))

	return r
}
