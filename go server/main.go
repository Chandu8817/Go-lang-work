package main 
import "fmt"
import	"log"
import	"net/http"



func formHandler(w http.ResponseWriter, r *http.Request){

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm failed", err)
		return

	}

	fmt.Fprintln(w, "POst Request successfully");

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintln(w, "Name: %s", name)
	fmt.Fprintln(w, "Address: %s", address)



}
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "Invalid URL", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Invalid method", http.StatusNotFound)
		return

	}

	fmt.Fprintf(w,"hello")
}

func main() {

	fileserver := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileserver)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("statring server at port 8000")	
	if err := http.ListenAndServe(":8080", nil); err !=nil{
		log.Fatal(err)
	}
}