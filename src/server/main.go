package main

import (
	"encoding/json"
	"net/http"
	"database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

var kengTableName = "keng"

type glossSchema struct {
	ID int
	Surface string
	Hanja string
	Gloss string
	Level string
}

type Glossary struct {
	Original string `json:"original"`
	Translated string `json:"translated"`
}

type GlossaryCollection struct {
	Glossarys []Glossary `json:"glossarys"`
}

type DB struct {
	db *sql.DB
}

func (db *DB)randomWord(w http.ResponseWriter, _ *http.Request) {
	sqlStmt := "SELECT id, surface, gloss FROM keng ORDER BY RANDOM() LIMIT 1"
	rows, err := db.db.Query(sqlStmt)

    if err != nil {
        fmt.Println(err)
        return
    }
	defer rows.Close()
	rows.Next()
	var d glossSchema
	if err := rows.Scan(&d.ID, &d.Surface, &d.Gloss); err != nil {
        fmt.Println(err)
        return
	}
	var g Glossary
	g.Original = d.Gloss
	g.Translated = d.Surface
	fmt.Fprintf(w, "%s", g)
}

type RandomWordsReq struct {
	NumWords int `json:"num_words"`
}

func (db *DB)randomWords(w http.ResponseWriter, r *http.Request) {
	var req RandomWordsReq
	err := json.NewDecoder(r.Body).Decode(&req)

	sqlStmt := fmt.Sprintf("SELECT id, surface, gloss FROM keng ORDER BY RANDOM() LIMIT %d", req.NumWords)
	rows, err := db.db.Query(sqlStmt)

    if err != nil {
        fmt.Println(err)
        return
    }
	defer rows.Close()
	gc := GlossaryCollection {
		Glossarys: make([]Glossary, req.NumWords),
	}
	i := 0
	for rows.Next() {
		var d glossSchema
		if err := rows.Scan(&d.ID, &d.Surface, &d.Gloss); err != nil {
    	    fmt.Println(err)
    	    return
		}
		gc.Glossarys[i] = Glossary{
			Original: d.Gloss,
			Translated: d.Surface,
		}
		i++
	}
	fmt.Fprintf(w, "%s", gc)
}

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./word.db")
    if err != nil {
        fmt.Println(err)
        return
    }

    defer db.Close()
    fmt.Println("Connected to the SQLite database successfully.")
	wordDB := DB{db: db}
	http.HandleFunc("/random_word", wordDB.randomWord)
	http.HandleFunc("/random_words", wordDB.randomWords)
	http.ListenAndServe(":8090", nil)
}
