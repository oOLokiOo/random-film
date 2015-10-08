package main

import (
	"net/http"
	"html/template"
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	//"encoding/json"
)


type Movies struct {
	XMLName 	xml.Name 	`xml:"movies"`
	FilmsList 	[]Film 		`xml:"film"`
}

type Film struct {
	Ru 			string 	`xml:"ru"`
	En 			string 	`xml:"en"`
	Year 		string 	`xml:"year"`
	Kinopoisk 	string 	`xml:"kinopoisk"`
	Imdb 		string 	`xml:"imdb"`
}

var (
	GOOGLE_IMAGES_URL 	string 	= "http://ajax.googleapis.com/ajax/services/search/images?v=1.0&imgsz=large&rsz=8&q="
	EN_SEARCH_PREFIX 	string 	= " film poster"
	RU_SEARCH_PREFIX 	string 	= " фильм постер"
	BLOCKED_RESOURCES = [...]string {
		"www.impawards.com",
		"en.wikipedia.org",
	}
	XML_PATH 			string 	= "../../users/1/films.xml"
)

type Page struct {
	Error 			string
	Title 			string
	TitleForSearch 	string
	Kinopoisk 		string
	Imdb 			string
	ImageUrl 		string
}


func catch(err error) {
	if err != nil {
		fmt.Println("\n\n--- --- ---\nERROR: ", err)
		fmt.Println("--- --- ---\n")
		panic(err)

		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	m := Movies{}
	p := Page{}

	xmlFile, _ := ioutil.ReadFile(XML_PATH)
	err := xml.Unmarshal(xmlFile, &m)
	catch(err)

	//for _, film := range m.FilmsList {
	//	fmt.Printf("\t%s\n", film)
	//}

	rand := rand.Intn(len(m.FilmsList));
	random_movie := m.FilmsList[rand]

	search_movie_title := ""
	if random_movie.En != "" {
		search_movie_title = random_movie.En + " " + random_movie.Year + " " + EN_SEARCH_PREFIX
	} else if random_movie.Ru != "" {
		search_movie_title = random_movie.Ru + " " + random_movie.Year + " " + EN_SEARCH_PREFIX // RU_SEARCH_PREFIX - bad result with "ru"
	}

	if search_movie_title != "" {
		//fmt.Println(search_movie_title)
		//res, err := http.Get(GOOGLE_IMAGES_URL + search_movie_title)
		//perror(err)
		//body, err := ioutil.ReadAll(res.Body)
		//perror(err)

		title := ""
		if random_movie.Ru != "" 	{ title += random_movie.Ru + " | " }
		if random_movie.En != "" 	{ title += random_movie.En + " | " }
		if random_movie.Year != "" 	{ title += random_movie.Year }

		if (title[len(title)-2:] == "| ") { title = title[:len(title)-2] }
		p.Title = title
		p.TitleForSearch = strings.Replace(title, " | ", " ", -1)

		p.Kinopoisk = random_movie.Kinopoisk
		p.Imdb = random_movie.Imdb
	} else {
		p.Error = "Movie [" + strconv.Itoa(rand) + "] - not found!"
	}


	w.Header().Set("Content-type", "text/html")

	t, err := template.ParseFiles("index.html")
	catch(err)
	t.Execute(w, p)
}


func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":80", nil)
}