package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/satori/go.uuid"

	"database/sql"

	_ "github.com/lib/pq"


	"github.com/gin-gonic/gin"
	"fmt"

	_ "io/ioutil"


	_ "github.com/jinzhu/gorm"
	_ "github.com/gin-gonic/gin"
	_ "github.com/aws/aws-sdk-go/private/protocol"
	_ "strings"
	_ "strconv"
	"strconv"
)


type Event struct {
	EventNum   sql.NullInt64
	EventName  sql.NullString
	EventOrganization sql.NullString
	OrganizationStreet sql.NullString
	OrganizationCity sql.NullString
	OrganizationState sql.NullString
	OrganizationZip sql.NullString
	ContactName1 sql.NullString
	ContactCellphone1 sql.NullString
	ContactName2 sql.NullString
	ContactCellphone2 sql.NullString
	StartDt sql.NullString
	EndDt sql.NullString
	Speaker sql.NullString
	Title sql.NullString
	NumOfAttendees sql.NullInt64
	ArrangedBy sql.NullString
}

type Book struct {

	Isbn string
	Title string
	Author string
	Price float32
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	var DATABASE_URL = os.Getenv("DATABASE_URL")

	db, errDB := sql.Open("postgres", DATABASE_URL)
	defer db.Close()

	if errDB != nil {
		log.Fatalf("Error connecting to the DB")
	} else {
		log.Println("Connection is good!")
	}

	rows, err := db.Query("SELECT * FROM events")
	if err != nil {
		log.Println("point 1")
		log.Fatal(err)
	}
	defer rows.Close()

	events := make([]*Event, 0)
	for rows.Next() {
		log.Println("point 2")
		event := new(Event)
		err := rows.Scan(
			&event.EventNum,
			&event.EventName,
		    &event.EventOrganization,
		    &event.OrganizationStreet,
			&event.OrganizationCity,
			&event.OrganizationState,
			&event.OrganizationZip,
			&event.ContactName1,
			&event.ContactCellphone1,
			&event.ContactName2,
			&event.ContactCellphone2,
			&event.StartDt,
			&event.EndDt,
			&event.Speaker,
			&event.Title,
			&event.NumOfAttendees,
			&event.ArrangedBy)

		if err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		if event.EventNum.Valid {
			log.Print(event.EventNum.Int64)
		} else {
			log.Print("event.EventNum is NULL")
		}


		if event.EventName.Valid {
			log.Print(event.EventName.String)
		} else {
			log.Print("event.EventName is NULL")
		}

		if event.EventOrganization.Valid {
			log.Print(event.EventOrganization.String)
		} else {
			log.Print("event.EventOrganization is NULL")
		}

		if event.OrganizationStreet.Valid {
			log.Print(event.OrganizationStreet.String)
		} else {
			log.Print("event.OrganizationStreet is NULL")
		}

		if event.OrganizationCity.Valid {
			log.Print(event.OrganizationCity.String)
		} else {
			log.Print("event.OrganizationCity is NULL")
		}

		if event.OrganizationState.Valid {
			log.Print(event.OrganizationState.String)
		} else {
			log.Print("event.OrganizationState is NULL")
		}

		if event.OrganizationZip.Valid {
			log.Print(event.OrganizationZip.String)
		} else {
			log.Print("event.OrganizationZip is NULL")
		}

		if event.ContactName1.Valid {
			log.Print(event.ContactName1.String)
		} else {
			log.Print("event.ContactName1 is NULL")
		}

		if event.ContactCellphone1.Valid {
			log.Print(event.ContactCellphone1.String)
		} else {
			log.Print("event.ContactCellphone1 is NULL")
		}


		if event.ContactName2.Valid {
			log.Print(event.ContactName2.String)
		} else {
			log.Print("event.ContactName2 is NULL")
		}

		if event.ContactCellphone2.Valid {
			log.Print(event.ContactCellphone2.String)
		} else {
			log.Print("event.ContactCellphone2 is NULL")
		}

		if event.StartDt.Valid {
			log.Print(event.StartDt.String)
		} else {
			log.Print("event.StartDt is NULL")
		}

		if event.EndDt.Valid {
			log.Print(event.EndDt.String)
		} else {
			log.Print("event.EndDt is NULL")
		}

		if event.Speaker.Valid {
			log.Print(event.Speaker.String)
		} else {
			log.Print("event.Speaker is NULL")
		}

		if event.Title.Valid {
			log.Print(event.Title.String)
		} else {
			log.Print("event.Title is NULL")
		}


		if event.NumOfAttendees.Valid {
			log.Print(event.NumOfAttendees.Int64)
		} else {
			log.Print("event.NumOfAttendees is NULL")
		}

		if event.ArrangedBy.Valid {
			log.Print(event.ArrangedBy.String)
		} else {
			log.Print("event.ArrangedBy is NULL")
		}





	}



	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")


	router.GET("/events/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "events.create.tmpl.html", nil)
	})

	router.GET("/events/select/:id", func(c *gin.Context) {


		id := c.Param("id")

		rows, err := db.Query("SELECT * FROM events where event_num = $1", id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		event := new(Event)

		for rows.Next() {
			err := rows.Scan(
				&event.EventNum,
				&event.EventName,
				&event.EventOrganization,
				&event.OrganizationStreet,
				&event.OrganizationCity,
				&event.OrganizationState,
				&event.OrganizationZip,
				&event.ContactName1,
				&event.ContactCellphone1,
				&event.ContactName2,
				&event.ContactCellphone2,
				&event.StartDt,
				&event.EndDt,
				&event.Speaker,
				&event.Title,
				&event.NumOfAttendees,
				&event.ArrangedBy)

			if err != nil {
				log.Fatal(err)
			}
		}


		c.HTML(http.StatusOK, "events.select.tmpl.html", event)
	})

	router.POST("/events/create", func(c *gin.Context) {
		Event_Name := c.PostForm("Event_Name")
		Event_Organization := c.PostForm("Event_Organization")
		Organization_Street := c.PostForm("Organization_Street")
		Start_DT := c.PostForm("Start_DT")


		_, errInsert := db.
		Exec("INSERT INTO events(event_name," +
			" event_organization, organization_street," +
				" start_dt) VALUES($1, $2, $3, $4)",
			Event_Name,
			Event_Organization,
		    Organization_Street,
		    Start_DT)

		if errInsert != nil {
			log.Println("DB Insertion is in error.")
			c.HTML(http.StatusOK, "events.create_error.tmpl.html", errInsert)
		} else {
			log.Println("DB Insertion successful.")

			c.HTML(http.StatusOK, "events.create_ok.tmpl.html", nil)

		}



		// go back to the main page
		// c.HTML(http.StatusOK, "index.tmpl.html", bks)

	})




	router.GET("/events/update/:id", func(c *gin.Context) {
		id := c.Param("id")

		rows, err := db.Query("SELECT * FROM events where event_num = $1", id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		event := new(Event)

		for rows.Next() {
			err := rows.Scan(
				&event.EventNum,
				&event.EventName,
				&event.EventOrganization,
				&event.OrganizationStreet,
				&event.OrganizationCity,
				&event.OrganizationState,
				&event.OrganizationZip,
				&event.ContactName1,
				&event.ContactCellphone1,
				&event.ContactName2,
				&event.ContactCellphone2,
				&event.StartDt,
				&event.EndDt,
				&event.Speaker,
				&event.Title,
				&event.NumOfAttendees,
				&event.ArrangedBy)

			if err != nil {
				log.Fatal(err)
			}
		}
		c.HTML(http.StatusOK, "events.update.tmpl.html", event)

	})

	router.POST("/events/update/:id", func(c *gin.Context) {

		id := c.PostForm("id")
		idInteger, err := strconv.Atoi(id)


		Event_Name := c.PostForm("Event_Name")
		Start_Dt := c.PostForm("Start_Dt")


		// Update
		stmt, err := db.Prepare(
			"update events set event_name = $1, start_dt = $2 where event_num=$3")
		checkErr(err)
		fmt.Println("statement is: ", stmt)

		res, err2 := stmt.Exec(Event_Name, Start_Dt, idInteger)
		checkErr(err2)
		defer stmt.Close()

		rowsAffected, err3 := res.RowsAffected()
		checkErr(err3)
		fmt.Println("rowsAffected is: ", rowsAffected)

		c.HTML(http.StatusOK, "events.update_post.tmpl.html", id)

	})

	router.GET("/events/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInteger, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}

		event := new(Event)
		event.EventNum = sql.NullInt64{
			Int64: int64(idInteger), Valid: true}

		c.HTML(http.StatusOK, "events.delete.tmpl.html", event)

	})

	router.POST("/events/delete/:id", func(c *gin.Context) {

		id := c.Param("id")
		idInteger, err := strconv.Atoi(id)
		checkErr(err)

		stmt, err := db.Prepare(
			"delete from events where event_num=$1")
		checkErr(err)

		res, err2 := stmt.Exec(idInteger)
		checkErr(err2)
		defer stmt.Close()

		rowsAffected, err3 := res.RowsAffected()
		checkErr(err3)
		fmt.Println("rowsAffected is: ", rowsAffected)

		c.HTML(http.StatusOK, "events.delete_post.tmpl.html", id)

	})



	router.GET("/traces", func(c *gin.Context) {
		c.HTML(http.StatusOK, "onlinetraces.tmpl.html", nil)
	})

	router.GET("/fileupload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "fileupload.tmpl.html", nil)
	})


	//router.GET("/repeat", repeatHandler)
	// the above one causes problem. Max Li


	router.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "index.tmpl.html", data)
		rows, err := db.Query("SELECT * FROM events order by start_dt desc, event_num desc")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		events := make([]*Event, 0)
		for rows.Next() {
			log.Println("point 3")
			event := new(Event)
			err := rows.Scan(
				&event.EventNum,
				&event.EventName,
				&event.EventOrganization,
				&event.OrganizationStreet,
				&event.OrganizationCity,
				&event.OrganizationState,
				&event.OrganizationZip,
				&event.ContactName1,
				&event.ContactCellphone1,
				&event.ContactName2,
				&event.ContactCellphone2,
				&event.StartDt,
				&event.EndDt,
				&event.Speaker,
				&event.Title,
				&event.NumOfAttendees,
				&event.ArrangedBy)

			if err != nil {
				log.Fatal(err)
			}
			events = append(events, event)
		}
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}

		c.HTML(http.StatusOK, "index.tmpl.html", events)

	})

	router.Run(":" + port)
}
