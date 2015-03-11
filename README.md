# paysimple-go [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/Galvanize-IT/paysimple-go)
PaySimple Go Library

### Quickstart

Set the environmental variables `PAYSIMPLE_USER` and `PAYSIMPLE_SECRET`

Create either a live or sandboxed API:

```go
api := API()
// Or
api := Sandbox()
```

Create customers, accounts, and payments as needed:

```go
customer := Customer{
    FirstName:             "Test",
    LastName:              "Customer",
    ShippingSameAsBilling: true,
}
created, err := api.Customers.Create(customer)
if err != nil {
    log.Panic(err)
}
```

For more information, see the PaySimple API documentation: http://developer.paysimple.com/documentation/

- aodin, 2015
