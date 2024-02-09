package Func

import ("fmt"
 	"net/http"
	 "encoding/json"
	 "io"
	
)

func GetLocation(url string ) (Locations , error){
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

func GetDate(url string ) (Date , error){
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

func GetRelation(url string ) (Relation  , error){
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
