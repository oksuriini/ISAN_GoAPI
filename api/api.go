package api

import (
	"ISAN-api/isan"
	"ISAN-api/queries"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Api struct {
	Server http.Server
	mux    http.ServeMux
	Addr   string
	Port   string
	db     *queries.Database
	log    *log.Logger
	errlog *log.Logger
}

func CreateApi(addr, port string, db *queries.Database) *Api {

	newApi := &Api{
		Server: http.Server{
			Addr: port,
		},
		mux:    *http.NewServeMux(),
		Addr:   addr,
		Port:   port,
		db:     db,
		log:    log.New(os.Stdout, "/INFO: \t", log.Ldate|log.Ltime),
		errlog: log.New(os.Stdout, "/ERROR: \t", log.Llongfile|log.Ldate|log.Ltime),
	}

	newApi.registerHandlers()

	return newApi
}

func (a *Api) registerHandlers() {

	a.mux.HandleFunc("/test", a.testHandler)
	a.mux.HandleFunc("/api/v1/all", a.getAll)
	a.mux.HandleFunc("/api/v1/insertTest", a.insertTest)
	a.Server.Handler = &a.mux
}

func (a *Api) testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Testing works"))
	fmt.Println("Testing handler works")
}

func (a *Api) getAll(w http.ResponseWriter, r *http.Request) {
	isanRecords, err := a.db.GetAll()
	if err != nil {
		a.errlog.Println("Fatal error on reading database")
		a.errlog.Fatalf("Error: %s\n", err)
	}
	w.Write([]byte(fmt.Sprintln(isanRecords)))
}

func (a *Api) insertTest(w http.ResponseWriter, r *http.Request) {
	isantest := isan.CreateNewISAN("12345Insert", "workInsert", "idunnoInsert", "releaseInsert", "2019-05-05", "dunnoInsert", "100", "langInsert")

	res, err := a.db.InsertISAN(isantest)
	if err != nil {
		a.errlog.Println("Error inserting into database")
		a.errlog.Fatalf("Error: %s\n", err)
	}
	w.Write([]byte(fmt.Sprintf("Inserted value %v\n Response: %d", isantest, res)))
}
