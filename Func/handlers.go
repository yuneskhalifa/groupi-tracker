package Func

import ("fmt"
 	"net/http"
 	"html/template"
// 	"groupie-tracker/Func"
// 	"os"
	 "io"
	 "encoding/json"
	// "log"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		if r.URL.Path != "/" {
		errorHandler(w ,"page not found", http.StatusNotFound)
			return 	
		}
		tmp1, _ := template.ParseFiles("templates/index.html")
		w.WriteHeader(200)
		path := "https://groupietrackers.herokuapp.com/api/artists"
		response, err := http.Get(path)
		if err != nil {
			fmt.Print(err.Error())
			errorHandler(w,"Internal Server Error", http.StatusInternalServerError )
			return
		}
		defer response.Body.Close()
		body, Err := io.ReadAll(response.Body)
		if Err != nil {
			fmt.Print(err.Error())
			errorHandler(w,"Internal Server Error", http.StatusInternalServerError)
			return
		}

		var artists []Artist
		json.Unmarshal(body, &artists)

		tmp1.Execute(w, artists)


}
}

func errorHandler(w http.ResponseWriter, status string , code int ){
	w.WriteHeader(code)
	var error1 Error
	error1.Code= code
	error1.Status = status
	template, err := template.ParseFiles("templates/error.html")
	if err != nil {
		fmt.Println("error parsing error.html page")
		return 
	}

	err = template.Execute(w, error1)
	if err != nil {
		fmt.Println("error executing error.html page!")
		return
	}
}

func AllInfoHandler(w http.ResponseWriter, r *http.Request){
	artistName:= r.URL.Query().Get("name")
	temp1, err1:= template.ParseFiles("templates/band.html")
	if err1 != nil {
		fmt.Print("there is a error with parsing the band.html page")
		return 
	}

	response, err2 := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err2 != nil {
		fmt.Print("there is an error with getting the api of artist")
		return 
	}

	defer response.Body.Close()

	Body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		fmt.Print("there is an error with reading the response")
		errorHandler(w,"internal server error", http.StatusInternalServerError)
		return 
	}

	var artists []Artist
	json.Unmarshal(Body, &artists)

	var exactArtist *Artist
	for _, artist:= range artists {
		if artist.Name == artistName {
			exactArtist = &artist
			break
		}
	}

	if exactArtist == nil {
		errorHandler(w,"artist not found", http.StatusNotFound)
		return 
	}

	location, errLo := GetLocation(exactArtist.Locations)
	if errLo != nil {
		errorHandler(w,"internal server error1", http.StatusInternalServerError)
	}

	date, errDa := GetDate(exactArtist.ConcertDates)
	if errDa != nil {
		errorHandler(w,"internal server error2", http.StatusInternalServerError)
	}

	relation, errRe := GetRelation(exactArtist.Relations)
	if errRe != nil {
		errorHandler(w,"internal server error3", http.StatusInternalServerError)
	}

	err := temp1.Execute(w, WebTemp{Artist: *exactArtist , Location:location, Date:date, Relation:relation})
	if err != nil {
		errorHandler(w,"internal server error4", http.StatusInternalServerError)
	}
	

}

