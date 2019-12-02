package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/joshuaherrera/snippetbox/pkg/models"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Make it so this url path only renders for '/',
	// otherwise, return a 404 response to the client
	// Must return or else rest of fcn executes.
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}
	// make a slice referencing templates, home must be first
	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }
	// // use template.ParseFiles() fcn to read the template file
	// // into a template set. Log any errors with status code 500
	// // pass slice as variadic param
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	// use serverError helper from helpers
	// 	app.serverError(w, err)
	// 	return
	// }

	// // Use Execute() mehtod on template set to write template content
	// // as res body. Last param to Execute() reps dynamic data we want
	// // to pass in.
	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
}

// Add a showSnippet handler function.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// extract val of id param and convert to int. If fails
	// or less than 1, render 404 error
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)

	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "%v", s)
}

// Add a createSnippet handler function.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// NOTE: can only call w.WriteHeader once per res.
	//       Also, if we don't use the method to send a status code
	//		 w.Write will automatically send a 200 OK status code.
	if r.Method != "POST" {
		// Go automatically sets 3 sys-gen'd headers: Date, Content-Length,
		// and Content-Type... if can't detect Content-Type defaults to
		// application/octet-stream.
		// MUST set content-type for JSON b/c Go detecs it as plaintext

		// can also Get, Add, Del with .Header()
		// Del() wont rm sys-gen'd headers, must access map and set to nil
		// eg w.Header()["Date"] = nil
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mt. Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
