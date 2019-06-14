service.MonthlyReport.Generate()
service.DailyReport.Generate()

service.DataImport.DoImport()
service.DataSource.ImportFile()

service.User.ChangePassword()
service.User.Login/service.User.Logout()

service.Project.GetAll()

//----------------------------------------------------------
// services.go
package service

var MonthlyReport = &MonthlyReportService{ DB: db }
var DailyReport   = &DailyReportService{ DB: db }
var DataImport    = &DataImportService{ DB: db }
var User          = &UserService{ DB: db }

func Init(db *sql.DB) { ... }

//----------------------------------------------------------
// monthly-report.go
package service

type MonthlyReportService struct { DB *sql.DB }

func (z MonthlyReportService) Generate() {
    z.DB.Query(...)
}

//----------------------------------------------------------
// daily-report.go
package service

type DailyReportService struct { DB *sql.DB }

func (z DailyReportService) Generate() {
    z.DB.Query(...)
}

//----------------------------------------------------------
// data-import.go
package service

type DataImportService struct { DB *sql.DB }

func (z DataImportService) ImportFile() {
    z.DB.Query(...)
}

//----------------------------------------------------------
// user.go
package service

type UserService struct { DB *sql.DB }

func (z UserService) CreateNew() {
    z.DB.Query(...)
}
