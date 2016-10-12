hermes
======

A Go / Golang client for the Hermes Distribution Interface / Routing Web Service.

Installation
------------

  go get github.com/asaskevich/govalidator
  go get github.com/TheBookPeople/hermes

Usage
-----

Import with:
```go
import "github.com/TheBookPeople/hermes"
```

Perform request with:
```go
// Create a client
client := hermes.NewClient("username", "123 (client ID)", "client name", "password")

// create a routing request
request := client.NewDeliveryRoutingRequest()

// Note: that the string type used is hermes.TrimmedString, ensuring that whitespace will trimmed.
// Hermes can return values with excessive whitespace, often far longer than their own documented field lengths.

// create an entry for the request
entry := &hermes.DeliveryRoutingRequestEntry{
  Customer: &hermes.Customer{
    Address: &hermes.Address{
      Title:  "Mrs",
      FirstName:  "Prime",
      LastName:   "Minister",
      HouseName:  ""
      HouseNo:    "10",
      StreetName: "Downing Street",
      City:       "London",
      Region:     "Greater London"
      PostCode:    "SW1A 2AA"
      CountryCode: "GB",
    },
    HomePhoneNo:    "0123456789",
    MobilePhoneNo:  "077123456789",
    Email:          "pm@gov.uk",
    //  CustomerReference1:
    //  CustomerReference2:
    CustomerAlertType: hermes.Email,
    //  DeliveryMessage: "Please leave at No.11 if not in."
    //  SpecialInstruction1:
  },
  Parcel: &hermes.Parcel{
    Weight: 1000, // grams
    Length: 30,
    Width:  10,
    Depth:  5,
    //  Girth:  2*(10+5),
    //  CombinedDimension:
    Volume:            1500,
    Currency:          "GBP",
    Value:             999, // pence
    NumberOfItems:     1,
    },
  SenderAddress: &hermes.SenderAddress{
    AddressLine1: "My Company",
    AddressLine2: "My Road",
    AddressLine3: "My Town",
    AddressLine4: "LL57 4FB",
  },
  ExpectedDespatchDate: time.Now(),
  CountryOfOrigin:      "GB",
}

// assign the entry to the request
request.AddEntry(*entry)

// Validate the address, returns *hermes.RoutingResponse
routingResponse, err := client.ValidateDeliveryAddress(request)
```

Copyright
---------

Copyright The Book People 2016.

License
-------

This library is distributed under under the GPLv3 license found in the LICENSE file.

