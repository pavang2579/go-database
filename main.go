package main()

import (
	"fmt"
	"os"
	"encoding/json"
	"sync"
	"path/filepath"
	"github.com/blend/go-sdk/stringutil"
	"github.com/jcelliott/lumber"

)

const Version = "1.0.0"

type (
	Logger interface{
		Fatal(string, ...interface{})
		Error(string, ....interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct{
		mutex sync.Mutex
		mutexes map[string]*sync.Mutex
		dir string
		log Logger
	}
)

type Options struct{
	Logger
}

func New(dir string, options *Options)(*Driver, error){
dir = filepath.Clean(dir)

opts := Options{}

if options != nil{
	opts = *options
}

if opts.Logger == nil {
	opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
}

driver := Driver{
	dir: dir
	mutexes: make(map[string]*sync.Mutex),
	log: opts.Logger,
}

if _, err := os.stat(dir); err == nil {
	opts.Logger.Debug("Using '%s' (database already exists)\n", dir)
	return &driver, nil
}

}

func (d *Driver) Write() error{

}

func (d *Driver) Read() error{

}

func (d *Driver) ReadAll()(){

}

func (d *Driver) Delete() error{

}

func (d *Driver) getorCreateMutex() *sync.Mutex{

}


type Address struct {
	City string
	State string
	Country string
	Pincode json.Number
}

type User struct {
	Name string
	Age json.Number
	Contact string
	Company string
	Address Address
}

func main(){
	dir := "./"

	db = New(dir, nil)
	if err != nil{
		fmt.Println("Error", err)
	}

	employeess := []User {
		{"John","23","8123827270","Mytech",Address{"bengaluru","karnataka","India","560078"}},
		{"Jade","24","8123827220","Myntra",Address{"New yrk city","New york","USA","654438"}},
		{"kasi","29","8123827279","Mytech",Address{"bengaluru","karnataka","India","560078"}},
		{"Ravi","34","8123827170","Mytech",Address{"bengaluru","karnataka","India","560040"}},
		{"Raj","32","8123827240","Analyec",Address{"bengaluru","karnataka","India","560044"}},
		{"Jay","22","8123827243","Mytech",Address{"bengaluru","karnataka","India","560078"}},

	}

	for _, value := range employeess{
		db.Write("users", value.Name, User{
			Name: value.Name,
			Age: value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error",err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshall([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append (allusers, employeeFound
		)
	}
	fmt.Println((allusers))


}
