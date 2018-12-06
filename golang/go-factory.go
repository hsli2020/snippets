// The factory method pattern in Go.  Jan 23, 2016 7 minute read
// 
// I have been writing web services in Go for approaching 2 years now. The factory pattern is
// something I have found particularly useful, especially for writing clean, concise, maintainable
// and testable code. Originally I used it just for the data access layer so I could swap in and
// out upon MySQL, PostgreSQL, etc without changing the application layer code but it has proved
// useful in many other places.
// 
// Starting at the very beginning we define our package and import the required libraries.

// Package definition and import the required stdlib packages.
package main

import (
    "fmt"
    "log"
    "errors"
    "strings"
    "database/sql"
    "sync"
)

// For the purposes of illustration we will create a DataStore interface that is implemented by
// PostgreSQLDataStore and MemoryDataStore. The DataStore interface will specify 2 methods: Name
// and FindUserNameById.

var UserNotFoundError = errors.New("User not found")

type DataStore interface {
    Name() string
    FindUserNameById(id int64) (string, error)
}

// Then create 2 structs that implement the the interface.

// The first implementation.
type PostgreSQLDataStore struct {
    DSN string
    DB sql.DB
}

func (pds *PostgreSQLDataStore) Name() string {
    return "PostgreSQLDataStore"
}

func (pds *PostgreSQLDataStore) FindUserNameById(id int64) (string, error) {
    var username string
    err := pds.DB.Query("SELECT username FROM users WHERE id=$1", id).Scan(&username)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", UserNotFoundError
        }
        return "", err
    }
    return username, nil
}

//The second implementation.
type MemoryDataStore struct {
    sync.RWMutex
    Users map[int64]string
}

func (mds *MemoryDataStore) Name() string {
    return "MemoryDataStore"
}

func (mds *MemoryDataStore) FindUserNameById(id int64) (string, error) {
    mds.RWMutex.RLock()
    defer mds.RWMutex.RUnlock()
    username, ok := mds.Users[id];
    if !ok {
        return "", UserNotFoundError
    }
    return username, nil
}

// Now we can start to explore the power of the Factory method pattern.
// 
// ...create objects without having to specify the exact 'type' of the object that will be created
// 
// First we must create factory methods for all our implementations that return the common
// interface. These are essentially constructors that accept a common argument. In our case this
// common argument is a map[string]string.

type DataStoreFactory func(conf map[string]string) (DataStore, error)

func NewPostgreSQLDataStore(conf map[string]string) (DataStore, error) {
    dsn, ok := conf.Get("DATASTORE_POSTGRES_DSN", "")
    if !ok {
        return nil, errors.New(fmt.Sprintf("%s is required for the postgres datastore",
            "DATASTORE_POSTGRES_DSN"))
    }

    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Panicf("Failed to connect to datastore: %s", err.Error())
        return nil, datastore.FailedToConnect
    }

    return &PostgresDataStore{
        DSN: dsn,
        DB:  db,
    }, nil
}

func NewMemoryDataStore(conf map[string]string) (DataStore, error) {
     return &MemoryDataStore{
        Users: &map[int64]string{
            1: "mnbbrown",
            0: "root",
        },
        RWMutex: &sync.RWMutex{},
     }, nil
}

// Now we must store these factory methods somewhere to be called upon as needed. I'll create a
// Register helper method to add factories to datastoreFactories. The init function registers both
// the factories we created above using easy to remember names.

var datastoreFactories = make(map[string]DataStoreFactory)

func Register(name string, factory DataStoreFactory) {
    if factory == nil {
        log.Panicf("Datastore factory %s does not exist.", name)
    }
    _, registered := datastoreFactories[name]
    if registered {
        log.Errorf("Datastore factory %s already registered. Ignoring.", name)
    }
    datastoreFactories[name] = factory
}

func init() {
    Register("postgres", NewPostgreSQLDataStore)
    Register("memory", NewMemoryDataStore)
}

// Now the magic happens. Using the Create function below the appropriate factory method will be
// called using the the conf argument to create an instance of the DataStore interface.

func CreateDatastore(conf map[string]string) (DataStore, error) {

    // Query configuration for datastore defaulting to "memory".
    engineName := conf.Get("DATASTORE", "memory")

    engineFactory, ok := datastoreFactories[engineName]
    if !ok {
        // Factory has not been registered.
        // Make a list of all available datastore factories for logging.
        availableDatastores := make([]string, len(datastoreFactories))
        for k, _ := range datastoreFactories {
            availableDatastores = append(availableDatastores, k)
        }
        return nil, errors.New(fmt.Sprintf("Invalid Datastore name. Must be one of: %s",
            strings.Join(availableDatastores, ", ")))
    }

    // Run the factory with the configuration.
    return engineFactory(conf)
}

// You can use the above CreateDatastore method in your application code as follows:

datastore, err := CreateDataStore(&map[string]string{
    "DATASTORE": "postgres",
    "DATASTORE_POSTGRES_DSN": "dbname=factoriesareamazing",
})

// of use another datastore:

datastore, err := CreateDataStore(&map[string]string{
    "DATASTORE": "memory",
})

// This interface driven approach allows you to do injection and makes mocking simple for unit
// tests. i.e you can test against MockDataStore which implements the DataStore interface rather
// than having to spin up an instances of PostgreSQL for your unit tests. That said, you should
// still be writing integration tests against a real version of PostgreSQL to account for SQL
// errors and other quirks you could have missed.
