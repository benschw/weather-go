package location

import (
	"github.com/benschw/weather-go/location/api"
	wclient "github.com/benschw/weather-go/openweather/client"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var _ = log.Printf

type LocationService struct {
	Bind          string
	Db            gorm.DB
	WeatherClient WeatherClient
}

func NewLocationService(bind string, dbStr string) (*LocationService, error) {
	s := &LocationService{}

	db, err := DbOpen(dbStr)
	if err != nil {
		return s, err
	}

	s.Db = db
	s.Bind = bind
	s.WeatherClient = &wclient.WeatherClient{}

	return s, nil
}

func (s *LocationService) MigrateDb() error {

	s.Db.AutoMigrate(&api.Location{})
	return nil
}

func (s *LocationService) Run() error {

	// route handlers
	resource := &LocationResource{
		Db:            s.Db,
		WeatherClient: s.WeatherClient,
	}

	// Configure Routes
	r := mux.NewRouter()

	r.HandleFunc("/location", resource.Add).Methods("POST")
	r.HandleFunc("/location", resource.FindAll).Methods("GET")
	r.HandleFunc("/location/{id}", resource.Find).Methods("GET")
	r.HandleFunc("/location/{id}", resource.Save).Methods("PUT")
	r.HandleFunc("/location/{id}", resource.Delete).Methods("DELETE")

	http.Handle("/", r)

	// Start HTTP Server
	return http.ListenAndServe(s.Bind, nil)
}

func DbOpen(dbStr string) (gorm.DB, error) {
	db, err := gorm.Open("mysql", dbStr)
	if err != nil {
		return db, err
	}
	db.SingularTable(true)
	return db, nil
}
