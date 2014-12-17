package weather

import (
	"github.com/benschw/weather-go/weather/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var _ = log.Printf

type Server struct {
	Database string
	Bind     string
}

func (s *Server) getDb() (gorm.DB, error) {
	db, err := gorm.Open("mysql", s.Database)
	if err != nil {
		return db, err
	}
	db.SingularTable(true)
	return db, nil
}

func (s *Server) Migrate() error {
	db, err := s.getDb()
	if err != nil {
		return err
	}

	db.AutoMigrate(api.Location{})
	return nil
}

func (s *Server) Run() error {
	db, err := s.getDb()
	if err != nil {
		return err
	}

	// route handlers
	resource := &LocationResource{
		Db: db,
	}

	// Configure Routes
	r := mux.NewRouter()

	r.HandleFunc("/location", resource.Add).Methods("POST")
	r.HandleFunc("/location", resource.findAll).Methods("GET")
	r.HandleFunc("/location/{id}", resource.Find).Methods("GET")
	r.HandleFunc("/location/{id}", resource.Save).Methods("PUT")
	r.HandleFunc("/location/{id}", resource.delete).Methods("DELETE")

	http.Handle("/", r)

	// Start HTTP Server
	return http.ListenAndServe(s.Bind, nil)
}
