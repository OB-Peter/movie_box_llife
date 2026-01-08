package main

import (
	"flag"
	"log"
	"net/http"
)

// Define a string constant containing the HTML for the webpage. This consists of a
// header tag, and some JavaScript which calls our POST /v1/tokens/authentication
// endpoint and writes the response body to inside the <div id="output"> tag.
const html = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
</head>
<body>
	<h1>Preflight CORS</h1>
	<div id="output"></div>
	<script>
		fetch("/v1/tokens/authentication", {
			method: "POST",
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				email: 'alice@example.com',
				password: 'pa55word'
			})
		})
		.then(function(response) {
			return response.json();
		})
		.then(function(data) {
			document.getElementById("output").innerHTML = JSON.stringify(data);
		})
		.catch(function(err) {
			document.getElementById("output").innerHTML = err;
		});
	</script>
</body>
</html>
`

func main() {
	addr := flag.String("addr", ":9000", "Server address")
	flag.Parse()

	log.Printf("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	}))
	log.Fatal(err)
}
