package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Avg(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	url := r.URL.Query().Get("url")
	var totalRate = 0
	var totalComment = 0
	c := colly.NewCollector()
	//pour les restaurants
	c.OnHTML(".review-container", func(e *colly.HTMLElement) {
		e.ForEach(".bubble_50",func(_ int, elem *colly.HTMLElement){
			totalRate+=5
			totalComment+=1
		})
		e.ForEach(".bubble_40",func(_ int, elem *colly.HTMLElement){
			totalRate+=4
			totalComment+=1
		})
		e.ForEach(".bubble_30",func(_ int, elem *colly.HTMLElement){
			totalRate+=3
			totalComment+=1
		})
		e.ForEach(".bubble_20",func(_ int, elem *colly.HTMLElement){
			totalRate+=2
			totalComment+=1
		})
		e.ForEach(".bubble_10",func(_ int, elem *colly.HTMLElement){
			totalRate+=1
			totalComment+=1
		})

	})
	//pour les hotels
	c.OnHTML(".hotels-review-list-parts-RatingLine__bubbles--1oCI4", func(e *colly.HTMLElement) {
		e.ForEach(".bubble_50",func(_ int, elem *colly.HTMLElement){
			totalRate+=5
			totalComment+=1
		})
		e.ForEach(".bubble_40",func(_ int, elem *colly.HTMLElement){
			totalRate+=4
			totalComment+=1
		})
		e.ForEach(".bubble_30",func(_ int, elem *colly.HTMLElement){
			totalRate+=3
			totalComment+=1
		})
		e.ForEach(".bubble_20",func(_ int, elem *colly.HTMLElement){
			totalRate+=2
			totalComment+=1
		})
		e.ForEach(".bubble_10",func(_ int, elem *colly.HTMLElement){
			totalRate+=1
			totalComment+=1
		})

	})
	c.Visit(url)
	if totalComment == 0{
		totalComment = 1
	}
	fmt.Print("Rate ",totalRate)
	fmt.Print("Coms ",totalComment)
	data := float32(totalRate)/float32(totalComment)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}


func Words(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	url := r.URL.Query().Get("url")
	var totalWords = 0
	c := colly.NewCollector()
	//pour les retaurants
	c.OnHTML("p.partial_entry", func(e *colly.HTMLElement) {
		var numberWords = 0
		for pos,char := range e.Text {
			_ = pos
			if char == ' '{
				numberWords+=1
			}
		}
		numberWords+=1
		totalWords+=numberWords
	})
	//pour les hotels
	c.OnHTML("q", func(e *colly.HTMLElement) {
		var numberWords = 0
		fmt.Print(e.Text);
		for pos,char := range e.Text {
			_ = pos
			if char == ' '{
				numberWords+=1
			}
		}
		numberWords+=1
		totalWords+=numberWords
	})

	c.Visit(url)
	data := totalWords
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func main() {
	router := httprouter.New()
	router.GET("/words", Words)
	router.GET("/avg", Avg)
	log.Fatal(http.ListenAndServe(":8083", router))
}