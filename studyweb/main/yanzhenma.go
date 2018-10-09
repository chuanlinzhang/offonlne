package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/haisum/recaptcha"
)

func main() {
	sitekey := "6LfiVjcUAAAAAJyjCEGVyTpmFqlpOMGVIZpZPy6p"
	re := recaptcha.R{
		Secret: "6LfiVjcUAAAAAJ7wALWYNew2yx0qbT0WxRR-kYu9",
	}

	form := fmt.Sprintf(`
        <html>
            <head>
                <script src='https://www.google.com/recaptcha/api.js'></script>
            </head>
            <body>
                <form action="/submit" method="post">
                    <div class="g-recaptcha" data-sitekey="%s"></div>
                    <input type="submit">
                </form>
            </body>
        </html>
    `, sitekey)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, form)
	})
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		isValid := re.Verify(*r)
		if isValid {
			fmt.Fprintf(w, "Valid")
		} else {
			fmt.Fprintf(w, "Invalid! These errors ocurred: %v", re.LastError())
		}
	})

	log.Printf("\n Starting server on http://localhost:8080 .")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Could not start server. %s", err)
	}
}
