package main
import ("fmt"
 	"net/http"
 	"html/template"
// 	"groupie-tracker/Func"
// 	"os"
	 "io"
	 "encoding/json"
	// "log"
)

type Artist struct {
	ID           string      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type Locations struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
}

type Date struct {
	ID  int `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
		ID  int `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`	
}

type Error struct{
	Code int 
	Status string 
}

type WebTemp struct{
	Artist Artist
	Location Locations
	Date Date 
	Relation Relation
} 

func main(){
	http.HandleFunc("/" ,indexHandler)
	http.HandleFunc("/band" ,allInfoHandler)
	
	fmt.Println("serving on port 8080.....")
	http.ListenAndServe(":8080" , nil)
	
	
}

func indexHandler(w http.ResponseWriter, r *http.Request){
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

func allInfoHandler(w http.ResponseWriter, r *http.Request){
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

	location, errLo := getLocation(exactArtist.Locations)
	if errLo != nil {
		errorHandler(w,"internal server error1", http.StatusInternalServerError)
	}

	date, errDa := getDate(exactArtist.ConcertDates)
	if errDa != nil {
		errorHandler(w,"internal server error2", http.StatusInternalServerError)
	}

	relation, errRe := getRelation(exactArtist.Relations)
	if errRe != nil {
		errorHandler(w,"internal server error3", http.StatusInternalServerError)
	}

	err := temp1.Execute(w, WebTemp{Artist: *exactArtist , Location:location, Date:date, Relation:relation})
	if err != nil {
		errorHandler(w,"internal server error4", http.StatusInternalServerError)
	}
	

}

func getLocation(url string ) (Locations , error){
	response, err := http.Get(url)
	if err != nil {
		fmt.Print("error with getting the url of the location ")
		return Locations{} , err
	}

	defer response.Body.Close()

	location , _ := io.ReadAll(response.Body)

	var location1 Locations
	json.Unmarshal(location, &location1)

	return location1, nil 
}

func getDate(url string ) (Date , error){
	response, err := http.Get(url)
	if err != nil {
		fmt.Print("error with getting the url of the location ")
		return Date{} , err
	}

	defer response.Body.Close()

	date , _ := io.ReadAll(response.Body)

	var date1 Date
	json.Unmarshal(date, &date1)

	return date1, nil 
}

func getRelation(url string ) (Relation  , error){
	response, err := http.Get(url)
	if err != nil {
		fmt.Print("error with getting the url of the location ")
		return Relation{} , err
	}
	defer response.Body.Close()
	relation , _ := io.ReadAll(response.Body)
	var relation1 Relation 
	json.Unmarshal(relation, &relation1)
	return relation1, nil 
}
