package Func 

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
