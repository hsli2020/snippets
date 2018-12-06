    Domain:
        Customer entity
        Item entity
        Order entity
    Use Cases:
        User entity
        Use case: Add Item to Order
        Use case: Get Items in Order
        Use case: Admin adds Item to Order
    Interfaces:
        Web Services for Item/Order handling
        Repositories for Use Cases and Domain entities persistence
    Infrastructure:
        The Database
        Code that handles DB connections
        The HTTP server
        Go Standard Library

Implementing the architecture

The domain

We will first create the domain layer. As said, our application and its use cases will be fully
working, but it won’t be a complete shop. Therefore, the code that defines our domain will be
short enough to justify putting it into a single file:

$GOPATH/src/domain/domain.go

package domain

import (
    "errors"
)

type CustomerRepository interface {
    Store(customer Customer)
    FindById(id int) Customer
}

type ItemRepository interface {
    Store(item Item)
    FindById(id int) Item
}

type OrderRepository interface {
    Store(order Order)
    FindById(id int) Order
}

type Customer struct {
    Id   int
    Name string
}

type Item struct {
    Id        int
    Name      string
    Value     float64
    Available bool
}

type Order struct {
    Id       int
    Customer Customer
    Items    []Item
}

func (order *Order) Add(item Item) error {
    if !item.Available {
        return errors.New("Cannot add unavailable items to order")
    }
    if order.value()+item.Value > 250.00 {
        return errors.New(`An order may not exceed
            a total value of $250.00`)
    }
    order.Items = append(order.Items, item)
    return nil
}

func (order *Order) value() float64 {
    sum := 0.0
    for i := range order.Items {
        sum = sum + order.Items[i].Value
    }
    return sum
}

The use cases

Let’s now look at the code of the use cases layer – again, this perfectly fits into one file:

$GOPATH/src/usecases/usecases.go

package usecases

import (
    "domain"
    "fmt"
)

type UserRepository interface {
    Store(user User)
    FindById(id int) User
}

type User struct {
    Id       int
    IsAdmin  bool
    Customer domain.Customer
}

type Item struct {
    Id    int
    Name  string
    Value float64
}

type Logger interface {
    Log(message string) error
}

type OrderInteractor struct {
    UserRepository  UserRepository
    OrderRepository domain.OrderRepository
    ItemRepository  domain.ItemRepository
    Logger          Logger
}

func (interactor *OrderInteractor) Items(userId, orderId int) ([]Item, error) {
    var items []Item
    user := interactor.UserRepository.FindById(userId)
    order := interactor.OrderRepository.FindById(orderId)
    if user.Customer.Id != order.Customer.Id {
        message := "User #%i (customer #%i) "
        message += "is not allowed to see items "
        message += "in order #%i (of customer #%i)"
        err := fmt.Errorf(message,
            user.Id,
            user.Customer.Id,
            order.Id,
            order.Customer.Id)
        interactor.Logger.Log(err.Error())
        items = make([]Item, 0)
        return items, err
    }
    items = make([]Item, len(order.Items))
    for i, item := range order.Items {
        items[i] = Item{item.Id, item.Name, item.Value}
    }
    return items, nil
}

func (interactor *OrderInteractor) Add(userId, orderId, itemId int) error {
    var message string
    user := interactor.UserRepository.FindById(userId)
    order := interactor.OrderRepository.FindById(orderId)
    if user.Customer.Id != order.Customer.Id {
        message = "User #%i (customer #%i) "
        message += "is not allowed to add items "
        message += "to order #%i (of customer #%i)"
        err := fmt.Errorf(message,
            user.Id,
            user.Customer.Id,
            order.Id,
            order.Customer.Id)
        interactor.Logger.Log(err.Error())
        return err
    }
    item := interactor.ItemRepository.FindById(itemId)
    if domainErr := order.Add(item); domainErr != nil {
        message = "Could not add item #%i "
        message += "to order #%i (of customer #%i) "
        message += "as user #%i because a business "
        message += "rule was violated: '%s'"
        err := fmt.Errorf(message,
            item.Id,
            order.Id,
            order.Customer.Id,
            user.Id,
            domainErr.Error())
        interactor.Logger.Log(err.Error())
        return err
    }
    interactor.OrderRepository.Store(order)
    interactor.Logger.Log(fmt.Sprintf(
        "User added item '%s' (#%i) to order #%i",
        item.Name, item.Id, order.Id))
    return nil
}

type AdminOrderInteractor struct {
    OrderInteractor
}

func (interactor *AdminOrderInteractor) Add(userId, orderId, itemId int) error {
    var message string
    user := interactor.UserRepository.FindById(userId)
    order := interactor.OrderRepository.FindById(orderId)
    if !user.IsAdmin {
        message = "User #%i (customer #%i) "
        message += "is not allowed to add items "
        message += "to order #%i (of customer #%i), "
        message += "because he is not an administrator"
        err := fmt.Errorf(message,
            user.Id,
            user.Customer.Id,
            order.Id,
            order.Customer.Id)
        interactor.Logger.Log(err.Error())
        return err
    }
    item := interactor.ItemRepository.FindById(itemId)
    if domainErr := order.Add(item); domainErr != nil {
        message = "Could not add item #%i "
        message += "to order #%i (of customer #%i) "
        message += "as user #%i because a business "
        message += "rule was violated: '%s'"
        err := fmt.Errorf(message,
            item.Id,
            order.Id,
            order.Customer.Id,
            user.Id,
            domainErr.Error())
        interactor.Logger.Log(err.Error())
        return err
    }
    interactor.OrderRepository.Store(order)
    interactor.Logger.Log(fmt.Sprintf(
        "Admin added item '%s' (#%i) to order #%i",
        item.Name, item.Id, order.Id))
    return nil
}

The interfaces

At this point, everything that has to be said, code wise, about our actual business and our
application use cases, is said. Let’s see what that means for the interfaces layer’s code. While
all code in the respective inner layers logically belongs together, the interfaces layer consists
of several parts that exist separately – therefore, we will split the code in this layer into
several files.

As our shop has to be accessible through the web, let’s start with the web service:

$GOPATH/src/interfaces/webservice.go

package interfaces

import (
    "fmt"
    "io"
    "net/http"
    "strconv"
    "usecases"
)

type OrderInteractor interface {
    Items(userId, orderId int) ([]usecases.Item, error)
    Add(userId, orderId, itemId int) error
}

type WebserviceHandler struct {
    OrderInteractor OrderInteractor
}

func (handler WebserviceHandler) ShowOrder(res http.ResponseWriter, req *http.Request) {
    userId, _ := strconv.Atoi(req.FormValue("userId"))
    orderId, _ := strconv.Atoi(req.FormValue("orderId"))
    items, _ := handler.OrderInteractor.Items(userId, orderId)
    for _, item := range items {
        io.WriteString(res, fmt.Sprintf("item id: %d\n", item.Id))
        io.WriteString(res, fmt.Sprintf("item name: %v\n", item.Name))
        io.WriteString(res, fmt.Sprintf("item value: %f\n", item.Value))
    }
}

Let’s create such an interface in src/interfaces/repositories.go:

type DbHandler interface {
  Execute(statement string)
  Query(statement string) Row 
}

type Row interface {
  Scan(dest ...interface{})
  Next() bool
}

That’s really a very limited interface, but it allows for all the operations the repositories need
to perform: reading, inserting, updating and deleting rows.

In the infrastructure layer, we will implement some glue code that uses a sqlite3 library to
actually talk to the database, while satisfying this interface – but first, let’s fully implement
the repositories:

$GOPATH/src/interfaces/repositories.go

package interfaces

import (
    "domain"
    "fmt"
    "usecases"
)

type DbHandler interface {
    Execute(statement string)
    Query(statement string) Row
}

type Row interface {
    Scan(dest ...interface{})
    Next() bool
}

type DbRepo struct {
    dbHandlers map[string]DbHandler
    dbHandler  DbHandler
}

type DbUserRepo DbRepo
type DbCustomerRepo DbRepo
type DbOrderRepo DbRepo
type DbItemRepo DbRepo

func NewDbUserRepo(dbHandlers map[string]DbHandler) *DbUserRepo {
    dbUserRepo := new(DbUserRepo)
    dbUserRepo.dbHandlers = dbHandlers
    dbUserRepo.dbHandler = dbHandlers["DbUserRepo"]
    return dbUserRepo
}

func (repo *DbUserRepo) Store(user usecases.User) {
    isAdmin := "no"
    if user.IsAdmin {
        isAdmin = "yes"
    }
    repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO users (id, customer_id, is_admin)
                                        VALUES ('%d', '%d', '%v')`,
                                        user.Id, user.Customer.Id, isAdmin))
    customerRepo := NewDbCustomerRepo(repo.dbHandlers)
    customerRepo.Store(user.Customer)
}

func (repo *DbUserRepo) FindById(id int) usecases.User {
    row := repo.dbHandler.Query(fmt.Sprintf(`SELECT is_admin, customer_id
                                             FROM users WHERE id = '%d' LIMIT 1`,
                                             id))
    var isAdmin string
    var customerId int
    row.Next()
    row.Scan(&isAdmin, &customerId)
    customerRepo := NewDbCustomerRepo(repo.dbHandlers)
    u := usecases.User{Id: id, Customer: customerRepo.FindById(customerId)}
    u.IsAdmin = false
    if isAdmin == "yes" {
        u.IsAdmin = true
    }
    return u
}

func NewDbCustomerRepo(dbHandlers map[string]DbHandler) *DbCustomerRepo {
    dbCustomerRepo := new(DbCustomerRepo)
    dbCustomerRepo.dbHandlers = dbHandlers
    dbCustomerRepo.dbHandler = dbHandlers["DbCustomerRepo"]
    return dbCustomerRepo
}

func (repo *DbCustomerRepo) Store(customer domain.Customer) {
    repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO customers (id, name)
                                        VALUES ('%d', '%v')`,
                                        customer.Id, customer.Name))
}

func (repo *DbCustomerRepo) FindById(id int) domain.Customer {
    row := repo.dbHandler.Query(fmt.Sprintf(`SELECT name FROM customers
                                             WHERE id = '%d' LIMIT 1`,
                                             id))
    var name string
    row.Next()
    row.Scan(&name)
    return domain.Customer{Id: id, Name: name}
}

func NewDbOrderRepo(dbHandlers map[string]DbHandler) *DbOrderRepo {
    dbOrderRepo := new(DbOrderRepo)
    dbOrderRepo.dbHandlers = dbHandlers
    dbOrderRepo.dbHandler = dbHandlers["DbOrderRepo"]
    return dbOrderRepo
}

func (repo *DbOrderRepo) Store(order domain.Order) {
    repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO orders (id, customer_id)
                                        VALUES ('%d', '%v')`,
                                        order.Id, order.Customer.Id))
    for _, item := range order.Items {
        repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO items2orders (item_id, order_id)
                                            VALUES ('%d', '%d')`,
                                            item.Id, order.Id))
    }
}

func (repo *DbOrderRepo) FindById(id int) domain.Order {
    row := repo.dbHandler.Query(fmt.Sprintf(`SELECT customer_id FROM orders
                                             WHERE id = '%d' LIMIT 1`,
                                             id))
    var customerId int
    row.Next()
    row.Scan(&customerId)
    customerRepo := NewDbCustomerRepo(repo.dbHandlers)
    order := domain.Order{Id: id, Customer: customerRepo.FindById(customerId)}
    var itemId int
    itemRepo := NewDbItemRepo(repo.dbHandlers)
    row = repo.dbHandler.Query(fmt.Sprintf(`SELECT item_id FROM items2orders
                                            WHERE order_id = '%d'`,
                                            order.Id))
    for row.Next() {
        row.Scan(&itemId)
        order.Add(itemRepo.FindById(itemId))
    }
    return order
}

func NewDbItemRepo(dbHandlers map[string]DbHandler) *DbItemRepo {
    dbItemRepo := new(DbItemRepo)
    dbItemRepo.dbHandlers = dbHandlers
    dbItemRepo.dbHandler = dbHandlers["DbItemRepo"]
    return dbItemRepo
}

func (repo *DbItemRepo) Store(item domain.Item) {
    available := "no"
    if item.Available {
        available = "yes"
    }
    repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO items (id, name, value, available)
                                        VALUES ('%d', '%v', '%f', '%v')`,
                                        item.Id, item.Name, item.Value, available))
}

func (repo *DbItemRepo) FindById(id int) domain.Item {
    row := repo.dbHandler.Query(fmt.Sprintf(`SELECT name, value, available
                                             FROM items WHERE id = '%d' LIMIT 1`,
                                             id))
    var name string
    var value float64
    var available string
    row.Next()
    row.Scan(&name, &value, &available)
    item := domain.Item{Id: id, Name: name, Value: value}
    item.Available = false
    if available == "yes" {
        item.Available = true
    }
    return item
}

The infrastructure

As stated above, our repositories view “The Database” as an abstract being where SQL queries can
be send to and rows can be retrieved from. They don’t care about infrastructural issues like
connecting to the database or even figuring out which database to use. This is done in
src/infrastructure/sqlitehandler.go, where the high level DbHandler interface is implemented using
low level means:

$GOPATH/src/infrastructure/sqlitehandler.go

package infrastructure

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "interfaces"
)

type SqliteHandler struct {
    Conn *sql.DB
}

func (handler *SqliteHandler) Execute(statement string) {
    handler.Conn.Exec(statement)
}

func (handler *SqliteHandler) Query(statement string) interfaces.Row {
    rows, err := handler.Conn.Query(statement)
    if err != nil {
        fmt.Println(err)
        return new(SqliteRow)
    }
    row := new(SqliteRow)
    row.Rows = rows
    return row
}

type SqliteRow struct {
    Rows *sql.Rows
}

func (r SqliteRow) Scan(dest ...interface{}) {
    r.Rows.Scan(dest...)
}

func (r SqliteRow) Next() bool {
    return r.Rows.Next()
}

func NewSqliteHandler(dbfileName string) *SqliteHandler {
    conn, _ := sql.Open("sqlite3", dbfileName)
    sqliteHandler := new(SqliteHandler)
    sqliteHandler.Conn = conn
    return sqliteHandler
}

(Again, zero error handling, among other things, in order to keep out code that doesn’t contribute
to the architectural ideas).

Using Yasuhiro Matsumoto’s sqlite3 library, this infrastructure code implements the DbHandler
interface that allows the repositories to talk to the database without the need to fiddle with low
level details.

Putting it all together

That’s it, all our architectural building blocks are now in place – let’s put them together in
main.go:

$GOPATH/main.go

package main

import (
    "usecases"
    "interfaces"
    "infrastructure"
    "net/http"
)

func main() {
    dbHandler := infrastructure.NewSqliteHandler("/var/tmp/production.sqlite")

    handlers := make(map[string] interfaces.DbHandler)
    handlers["DbUserRepo"] = dbHandler
    handlers["DbCustomerRepo"] = dbHandler
    handlers["DbItemRepo"] = dbHandler
    handlers["DbOrderRepo"] = dbHandler

    orderInteractor := new(usecases.OrderInteractor)
    orderInteractor.UserRepository = interfaces.NewDbUserRepo(handlers)
    orderInteractor.ItemRepository = interfaces.NewDbItemRepo(handlers)
    orderInteractor.OrderRepository = interfaces.NewDbOrderRepo(handlers)

    webserviceHandler := interfaces.WebserviceHandler{}
    webserviceHandler.OrderInteractor = orderInteractor

    http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
        webserviceHandler.ShowOrder(res, req)
    })
    http.ListenAndServe(":8080", nil)
}

We can use the following SQL to create a minimal data set in /var/tmp/production.sqlite:

CREATE TABLE users (id INTEGER, customer_id INTEGER, is_admin VARCHAR(3));
CREATE TABLE customers (id INTEGER, name VARCHAR(42));
CREATE TABLE orders (id INTEGER, customer_id INTEGER);
CREATE TABLE items (id INTEGER, name VARCHAR(42), value FLOAT, available VARCHAR(3));
CREATE TABLE items2orders (item_id INTEGER, order_id INTEGER);

INSERT INTO users (id, customer_id, is_admin) VALUES (40, 50, "yes");
INSERT INTO customers (id, name) VALUES (50, "John Doe");
INSERT INTO orders (id, customer_id) VALUES (60, 50);
INSERT INTO items (id, name, value, available) VALUES (101, "Soap", 4.99, "yes");
INSERT INTO items (id, name, value, available) VALUES (102, "Fork", 2.99, "yes");
INSERT INTO items (id, name, value, available) VALUES (103, "Bottle", 6.99, "no");
INSERT INTO items (id, name, value, available) VALUES (104, "Chair", 43.00, "yes");

INSERT INTO items2orders (item_id, order_id) VALUES (101, 60);
INSERT INTO items2orders (item_id, order_id) VALUES (104, 60);

Now, we can start the application, and point our browser at
http://localhost:8080/orders?userId=40&orderId=60. The result should be:

item id: 101
item name: Soap
item value: 4.990000
item id: 104
item name: Chair
item value: 43.000000

