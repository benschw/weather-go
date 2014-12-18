package weather

import (
	"github.com/benschw/weather-go/weather/api"
	"github.com/benschw/weather-go/weather/openweather"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var _ = log.Printf

type WeatherService struct {
	Db   gorm.DB
	Bind string
}

func NewWeatherService(dbStr string, bind string) *WeatherService {
	db, err := dbOpen(dbStr)
	if err != nil {
		panic(err)
	}

	return &WeatherService{
		Db:   db,
		Bind: bind,
	}
}

func dbOpen(dbStr string) (gorm.DB, error) {
	db, err := gorm.Open("mysql", dbStr)
	if err != nil {
		return db, err
	}
	db.SingularTable(true)
	return db, nil
}

func (s *WeatherService) MigrateDb() error {

	s.Db.AutoMigrate(api.Location{})
	return nil
}

func (s *WeatherService) Run() error {
	c := openweather.WeatherClient{}

	// route handlers
	resource := &LocationResource{
		Db:         s.Db,
		CondClient: c,
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
