package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/satori/go.uuid"

	"database/sql"

	_ "github.com/lib/pq"

	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gin-gonic/gin"
	"fmt"

	"context"
	"time"
	_ "io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/credentials"

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

	// 		"host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")

	//var args string

	//args += "host=" + os.Getenv("myhost") + " "
	//args += "user=" + os.Getenv("user") + " "
	//args += "dbname=" + os.Getenv("dbname") + " "
	//args += "sslmode=disable "
	//args += "password=" + os.Getenv("password")
	//log.Println("args is: ", args)

	var URI = os.Getenv("URI")

	db, errDB := sql.Open("postgres", URI)
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


	router.GET("/books/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "books.create.tmpl.html", nil)
	})

	router.GET("/books/select/:id", func(c *gin.Context) {


		id := c.Param("id")
		//value, _ := strconv.Atoi(id)

		rows, err := db.Query("SELECT * FROM books where ISBN = $1", id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		bk := new(Book)

		for rows.Next() {
			err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
			if err != nil {
				log.Fatal(err)
			}
			// need to trim it
			//bk.Isbn = strings.TrimSpace(bk.Isbn)
		}


		c.HTML(http.StatusOK, "books.select.tmpl.html", bk)
	})

	router.POST("/books/create", func(c *gin.Context) {
		Isbn := c.PostForm("Isbn")
		Author := c.PostForm("Author")
		Title := c.PostForm("Title")
		Price := c.PostForm("Price")


		val, _ := strconv.ParseFloat(Price, 32)

		_, errInsert := db.
		Exec("INSERT INTO books(isbn, title, author, price) VALUES($1, $2, $3, $4)",
			Isbn, Title, Author, val)

		if errInsert != nil {
			log.Println("DB Insertion is in error.")
			c.HTML(http.StatusOK, "books.create_error.tmpl.html", errInsert)
		} else {
			log.Println("DB Insertion successful.")
			rows, err := db.Query("SELECT * FROM books order by isbn")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			bks := make([]*Book, 0)
			for rows.Next() {
				bk := new(Book)
				err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
				if err != nil {
					log.Fatal(err)
				}
				bks = append(bks, bk)
			}

			c.HTML(http.StatusOK, "books.create_ok.tmpl.html", nil)

		}



		// go back to the main page
		// c.HTML(http.StatusOK, "index.tmpl.html", bks)

	})




	router.GET("/books/update/:id", func(c *gin.Context) {
		id := c.Param("id")
		//value, _ := strconv.Atoi(id)

		rows, err := db.Query("SELECT * FROM books where ISBN = $1", id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		bk := new(Book)

		for rows.Next() {
			err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
			if err != nil {
				log.Fatal(err)
			}
			// need to trim it
			//bk.Isbn = strings.TrimSpace(bk.Isbn)
		}
		c.HTML(http.StatusOK, "books.update.tmpl.html", bk)

	})

	router.POST("/books/update/:id", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "index.tmpl.html", data)
		id := c.PostForm("id")

		Isbn := c.PostForm("Isbn")
		Author := c.PostForm("Author")
		Title := c.PostForm("Title")
		Price := c.PostForm("Price")

		// Update
		stmt, err := db.Prepare(
			"update BOOKs set Author = $1, Title = $2, Price = $3 where ISBN=$4")
		checkErr(err)
		fmt.Println("statement is: ", stmt)

		val, err := strconv.ParseFloat(Price, 32)

		fmt.Println("Author, Title, val, Isbn are: ", Author, Title, val, Isbn)

		res, err2 := stmt.Exec(Author, Title, val, Isbn)
		checkErr(err2)
		defer stmt.Close()

		rowsAffected, err3 := res.RowsAffected()
		checkErr(err3)
		fmt.Println("rowsAffected is: ", rowsAffected)



		c.HTML(http.StatusOK, "books.update_post.tmpl.html", id)

	})

	router.GET("/books/delete/:id", func(c *gin.Context) {
		id := c.Param("id")

		bk := new(Book)
		bk.Isbn = id

		c.HTML(http.StatusOK, "books.delete.tmpl.html", bk)

	})

	router.POST("/books/delete/:id", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "index.tmpl.html", data)
		id := c.Param("id")

		Isbn := id

		// Update
		stmt, err := db.Prepare(
			"delete from BOOKs where ISBN=$1")
		checkErr(err)
		fmt.Println("statement is: ", stmt)

		fmt.Println("ISBN is: ", Isbn)

		res, err2 := stmt.Exec(Isbn)
		checkErr(err2)
		defer stmt.Close()

		rowsAffected, err3 := res.RowsAffected()
		checkErr(err3)
		fmt.Println("rowsAffected is: ", rowsAffected)

		c.HTML(http.StatusOK, "books.delete_post.tmpl.html", id)

	})


	router.GET("/onlinetraces", func(c *gin.Context) {
		c.HTML(http.StatusOK, "onlinetraces.tmpl.html", nil)
	})

	router.GET("/traces", func(c *gin.Context) {
		c.HTML(http.StatusOK, "onlinetraces.tmpl.html", nil)
	})

	router.GET("/fileupload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "fileupload.tmpl.html", nil)
	})

	router.GET("/online", func(c *gin.Context) {
		c.HTML(http.StatusOK, "onlinetraces.tmpl.html", nil)
	})

	//router.GET("/repeat", repeatHandler)
	// the above one causes problem. Max Li

	router.POST("/fileupload", func(c *gin.Context) {

		// single file
		file, _ := c.FormFile("file")
		log.Println("The file name is: ", file.Filename)

		emailAddress := c.PostForm("email_address")
		log.Println("The email address is: ", emailAddress)
		cellPhoneNumber := c.PostForm("cell_phone_number")
		log.Println("The cell phone number is: ", cellPhoneNumber)

		var bucket, key string
		var timeout time.Duration

		timeout = 60 * time.Minute

		//bucket = os.Getenv("BUCKET-NAME")
		//key = os.Getenv("S3-KEY")
		AWS_ACCESS_KEY_ID :=
			os.Getenv("AWS-ACCESS-KEY-ID")
		AWS_SECRET_ACCESS_KEY :=
			os.Getenv("AWS-SECRET-ACCESS-KEY")

		// If you're working with temporary security credentials,
		// you can also keep the session token in AWS_SESSION_TOKEN.
		token := ""
		creds := credentials.NewStaticCredentials(
			AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, token)

		_, errCred := creds.Get()
		if errCred != nil {
			log.Fatal(errCred)
		}

		bucket = "ithreeman"
		// key is the same as file name to be stored
		key = ""

		sess := session.Must(session.NewSession(
			&aws.Config{
				Region:      aws.String(endpoints.UsEast2RegionID),
				Credentials: creds,
			}))
		// Create a new instance of the service's client with a Session.
		// Optional aws.Config values can also be provided as variadic arguments
		// to the New function. This option allows you to provide service
		// specific configuration.
		svc := s3.New(sess)

		// Create a context with a timeout that will abort the upload if it takes
		// more than the passed in timeout.
		ctx := context.Background()
		var cancelFn func()
		if timeout > 0 {
			ctx, cancelFn = context.WithTimeout(ctx, timeout)
		}
		// Ensure the context is canceled to prevent leaking.
		// See context package for more information, https://golang.org/pkg/context/
		defer cancelFn()

		//f, errOpen  := os.Open(file.Filename)
		f, errOpen := file.Open()
		if errOpen != nil {
			log.Fatalf("failed to open file %q, %v",
				file.Filename, errOpen)
		}

		key = file.Filename

		// Uploads the object to S3. The Context will interrupt the request if the
		// timeout expires.
		_, err := svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
			Body:   f,
		})
		if err != nil {
			aerr, ok := err.(awserr.Error);
			if ok && aerr.Code() ==
				request.CanceledErrorCode {
				// If the SDK can determine the request or retry delay was canceled
				// by a context the CanceledErrorCode error code will be returned.
				log.Fatalf("upload canceled due to timeout, %v\n", err)
			} else {
				log.Fatalf("failed to upload object, %v\n", err)
			}
		}

		log.Printf("successfully uploaded file to %s/%s\n", bucket, key)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})


	router.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "index.tmpl.html", data)
		rows, err := db.Query("SELECT * FROM books order by isbn")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		bks := make([]*Book, 0)
		for rows.Next() {
			bk := new(Book)
			err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
			if err != nil {
				log.Fatal(err)
			}
			bks = append(bks, bk)
		}
		c.HTML(http.StatusOK, "index.tmpl.html", bks)

	})

	router.Run(":" + port)
}
