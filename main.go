package main
import ("fmt"
 	"net/http"
 	//"html/template"
 	"groupie-tracker/Func"
// 	"os"
	// "io"
	// "encoding/json"
	// "log"
)

func main(){
	http.HandleFunc("/" ,Func.IndexHandler)
	http.HandleFunc("/band" ,Func.AllInfoHandler)
	fmt.Println("serving on port 8080.....")
	http.ListenAndServe(":8080" , nil)
}

