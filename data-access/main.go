package main

import (
	"database/sql"
    "fmt"
    "log"
    "os"
	"github.com/go-sql-driver/mysql"

    "net/http"
    "github.com/gin-gonic/gin"
)


var db *sql.DB

type Album struct {
    ID     int64
    Title  string
    Style string
    Price  float32
    ArtistId int64
}

// type album struct {
//     ID     string  `json:"id"`
//     Title  string  `json:"title"`
//     Artist string  `json:"artist"`
//     Price  float64 `json:"price"`
// }


func main() {
    fmt.Printf("os getEnv: %v\n", os.Getenv("DBUSER"))

    // Capture connection properties.
    cfg := mysql.Config{
        // User:   'ali'
        // Passwd: 'ali@1234'
        User:   os.Getenv("DBUSER"), //ali
        Passwd: os.Getenv("DBPASS"), //ali@1234
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

	//added after writing albumsByArtists
	// albums, err := albumsByArtist(3)
	// if err != nil {
    // log.Fatal(err)
	// }
	// fmt.Printf("Albums found: %v\n", albums)

	// //added after writing albumByID
	// alb, err := albumByID(2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Album found: %v\n", alb)

	// //add an album
	// albID, err := addAlbum(Album{
	// 	Title:  "The Modern Sound of Betty Carter",
	// 	Artist: "Betty Carter",
	// 	Price:  49.99,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ID of added album: %v\n", albID)

    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/post-albums", postAlbum)
    router.GET("/artist-albums/:name", getArtistAlbums)
    router.Run("localhost:8080")
    

}

func getAlbums(c *gin.Context) {
	// myAlbumsFromDB := database.albumsByArtist("Zahra Golmohammadi")
    albums, err := albumsByArtist(2)
    if err != nil {
        log.Fatal(err)
    }
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {

    var alb Album

    if err := c.BindJSON(&alb); err != nil {
        return
    }

    result, err := db.Exec("INSERT INTO album (title, style, price, artist_id) VALUES (?, ?, ?, ?)", alb.Title, alb.Style, alb.Price, alb.ArtistId)
    if err != nil {
        fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        fmt.Errorf("addAlbum: %v id: %v", err, id)
    }
    // return id, nil
    c.IndentedJSON(http.StatusCreated, alb) //result
}

func getArtistAlbums(c *gin.Context) {
    //todo
    name := c.Param("name")
    fmt.Printf(name)

    result, err := db.Query(`SELECT * FROM artist art JOIN album alb ON art.id = alb.artist_id WHERE art.name = "?"`, name)
    // result, err := db.Query("SELECT * FROM artist art JOIN album alb ON art.id = alb.artist_id WHERE art.name = ?", name)

    if err != nil {
        fmt.Errorf("Errorrrr in getArtistAlbum")
    }
    // c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

    c.IndentedJSON(http.StatusOK, result)

}


// *   *   *   *   *   *   *   *   *   *   *   *   *   *   *   *   *   *   *   *   *


// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(artistID int64) ([]Album, error) {
// func albumsByArtist(name string) ([]Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []Album

    // rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name) // ***** foreignKey Query
    rows, err := db.Query("SELECT * FROM album WHERE artist_id = ?", artistID) // ***** foreignKey Query
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", artistID, err)
    }

	// Defer closing rows so that any resources it holds will be released when the function exits.
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title,&alb.Style, &alb.Price, &alb.ArtistId); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", artistID, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", artistID, err)
    }
    return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
    // An album to hold data from the returned row.
    var alb Album

    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&alb.ID, &alb.Title,&alb.Style, &alb.Price, &alb.ArtistId); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("albumsById %d: no such album", id)
        }
        return alb, fmt.Errorf("albumsById %d: %v", id, err)
    }
    return alb, nil
}


// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, style, price, artist_id) VALUES (?, ?, ?, ?)", alb.Title, alb.Style, alb.Price, alb.ArtistId)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
    // c.IndentedJSON(http.StatusCreated, alb) //result
}




// Scan takes a list of pointers to Go values, where the column values will be written. Here,
//	 you pass pointers to fields in the alb variable, created using the & operator. Scan writes through the pointers to update the struct fields.


// Use DB.Exec to execute an INSERT statement.
