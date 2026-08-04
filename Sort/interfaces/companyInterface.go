package interfaces

type CompanyInterface interface{
    AddWorkerInfo(name, position string, salary, experience uint) error
    SortWorkers() ([]string, error) 
}