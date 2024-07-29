package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/selvamtech08/equation-compile/compile"
)

// help function for encode given data and set status code
func returnResponse(w http.ResponseWriter, kind string, code int, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	data := map[string]any{kind: message}
	json.NewEncoder(w).Encode(data)
}

// display html page for get request
func TeXGetForm(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/texpage.html")
	if err != nil {
		log.Println(err.Error())
		returnResponse(w, "error", http.StatusInternalServerError, "failed to parse rendering templates")
		return
	}
	temp.Execute(w, nil)
}

// recevied post request from user
func TeXPostForm(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("equ")
	preamble := r.FormValue("preamble")
	mode := r.FormValue("equmode")
	fileName, err := compile.SaveAsFile(code, preamble, mode)
	if err != nil {
		log.Println(err.Error())
		returnResponse(w, "error", http.StatusInternalServerError, "failed to process the equation")
		return
	}
	outFile, err := compile.Run(fileName)
	if err != nil {
		if err.Error() == "compile error" {
			// return compilation log as result and it will be shown in html page
			w.WriteHeader(http.StatusAccepted)
			log := fmt.Sprintf(`<pre>%s</pre>`, outFile)
			w.Write([]byte(log))
		} else {
			returnResponse(w, "error", http.StatusInternalServerError, "failed to generate image")
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	img := fmt.Sprintf(`<img src="data:image/png;base64,%s" alt="Blank" />`, outFile)
	w.Write([]byte(img))
}
