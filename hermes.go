package hermes

import (
	"bytes"
	"encoding/xml"
	"fmt"
	//"io/ioutil"
	"net/http"
	"time"
)

// comments: max length, mandatory?

// IdentityService - TODO
type IdentityService struct {
	IDCardNo      string `xml:"idCardNo,omitempty" valid:"length(0|20)"`     // 20
	IdcardType    string `xml:"idCardType,omitempty" valid:"length(0|3)"`    // 3
	AgeValidation int    `xml:"ageValidation,omitempty" valid:"length(0|3)"` // 3 (0-100 range) // TODO range
	DateOfBirth   time.Time
	Pin           string `xml:"pin,omitempty" valid:"length(0|35)"` // 35
	Module        string `xml:"Module" valid:"length(0|3)"`         // 3
}

//RetailStoreService - TODO
type RetailStoreService struct {
	RetailStoreID string  `xml:"retailStoreId" valid:"length(1|20)"` // 20, mandatory
	Address       Address `xml:"address,omitempty"`
}

//ParcelShopService  - TODO
type ParcelShopService struct {
	ParcelShopID string   `xml:"parcelShopId" valid:"length(1|20)"` // 20, mandatory
	Address      *Address `xml:"address,omitempty"`
}

// CashOnDelivery - TODO
type CashOnDelivery struct {
	CashValue            int    `xml:"cashValue,omitempty" valid:"length(0|10)"`           // 10 (in pence etc.)
	CashCurrency         string `xml:"cashCurrency,omitempty valid:length(0|3)"`           // 3 (e.g. EUR, GBP)
	BankTransferValue    int    `xml:"bankTransferValue,omitempty" valid:"length(0|10)"`   // 10 (in pence etc.)
	BankTransferCurrency string `xml:"bankTransferCurrency,omitempty" valid:"length(0|3)"` // 3
}

// StatedDay - TODO
type StatedDay struct {
	StatedDayIndicator string    `xml:"statedDayIndicator" valid:"length(1)"` // 1, mandatory
	StatedDate         time.Time `xml:"statedDate,omitempty"`
}

// StatedTime - TODO
type StatedTime string

const (
	// AM - Morning
	AM StatedTime = "1"
	// PM - Afternoon
	PM StatedTime = "2"
	//Evening - Evening. (PM is afternoon).
	Evening StatedTime = "3"
	//Midday - self explanitory.
	Midday StatedTime = "4"
)

// Services - TODO
type Services struct {
	StatedDay             StatedDay           `xml:"statedDay,omitempty"`
	StatedTime            StatedTime          `xml:"statedTime,omitempty"` // 1, 1=AM 2=PM,
	NextDay               bool                `xml:"nextDay,omitempty"`
	HouseholdSignature    bool                `xml:"householdSignature,omitempty"`
	Signature             bool                `xml:"signature,omitempty"`
	RedirectionProhibited bool                `xml:"redirectionProhibited,omitempty"`
	LimitedQuantity       bool                `xml:"limitedQuantity,omitempty"`
	CashOnDelivery        *CashOnDelivery     `xml:"cashOnDelivery,omitempty"`
	ParcelShopService     *ParcelShopService  `xml:"parcelShopService,omitempty"`
	RetailStoreService    *RetailStoreService `xml:"retailStoreService,omitempty"`
	IdentityService       *IdentityService    `xml:"identityService,omitempty"`
	TrackedServiceOptOut  bool                `xml:"tracedServiceOptOut,omitempty"`
}

// Content - TODO
type Content struct {
	SkuCode        string `xml:"skuCode" valid:"length(1|30)"`          // 30, mandatory
	SkuDescription string `xml:"skuDescription" valid:"length(1|2000)"` // 2000, mandatory
	HsCode         string `xml:"hsCode" valid:"length(1|10)"`           // 10, mandatory
	Value          int    `xml:"value"`                                 // 10, mandatory // TODO int length
}

// Parcel - TODO
type Parcel struct {
	Weight            int       `xml:"weight" valid:"length(0|7)"`            // 7, mandatory
	Length            int       `xml:"length" valid:"length(0|4)"`            // 4, mandatory
	Width             int       `xml:"width" valid:"length(0|4)"`             // 4, mandatory
	Depth             int       `xml:"depth" valid:"length(0|4)"`             // 4, mandatory
	Girth             int       `xml:"girth" valid:"length(0|4)"`             // 4, mandatory
	CombinedDimension int       `xml:"combinedDimension" valid:"length(1|4)"` // 4, mandatory
	Volume            int       `xml:"volume"`                                // 10, mandatory
	Currency          string    `xml:"currency" valid:"length(1|3)"`          // 3 mandatory, (USD, GBP etc.)
	Value             int       `xml:"value"`                                 // 10, mandatory
	NumberOfParts     int       `xml:"numberOfParts,omitempty"`               // 10 // valid from 1-99 // TODO range
	NumberOfItems     int       `xml:"numberOfItems,omitempty"`               // 10 // valid from 1-99 // TODO range
	HangingGarment    bool      `xml:"hangingGarment,omitempty"`
	TheftRisk         bool      `xml:"theftRisk,omitempty"`     // Not currently used.
	MultipleParts     bool      `xml:"multipleParts,omitempty"` // Not currently used.
	Catalogue         int       `xml:"catalogue,omitempty"`
	Description       int       `xml:"description,omitempty" valid:"length(0|32)"`    // 32
	OriginOfParcel    int       `xml:"originOfParcel,omitempty" valid:"length(0|32)"` // 32
	DutyPaid          int       `xml:"dutyPaid,omitempty" valid:"length(0|1)"`        // 1, mandatory if non EU
	Contents          []Content `xml:"contents"`
}

// SenderAddress - TODO
type SenderAddress struct {
	AddressLine1 string `xml:"addressLine1,omitempty" valid:"length(0|50)"` // 50
	AddressLine2 string `xml:"addressLine2,omitempty" valid:"length(0|50)"` // 50
	AddressLine3 string `xml:"addressLine3,omitempty" valid:"length(0|50)"` // 50
	AddressLine4 string `xml:"addressLine4,omitempty" valid:"length(0|50)"` // 50
}

// Address - TODO
type Address struct {
	Title        string `xml:"title,omitempty" valid:"length(0|20)"`        // 20
	FirstName    string `xml:"firstName,omitempty" valid:"length(0|20)"`    // 20
	LastName     string `xml:"lastName,omitempty" valid:"length(0|20)"`     // 20
	HouseNo      string `xml:"houseNo,omitempty" valid:"length(0|20)"`      // 20
	HouseName    string `xml:"houseName,omitempty" valid:"length(0|20)"`    // 20
	StreetName   string `xml:"streetName,omitempty" valid:"length(0|20)"`   // 20
	AddressLine1 string `xml:"addressLine1,omitempty" valid:"length(0|20)"` // 20
	AddressLine2 string `xml:"addressLine2,omitempty" valid:"length(0|20)"` // 20
	AddressLine3 string `xml:"addressLine3,omitempty" valid:"length(0|20)"` // 20
	City         string `xml:"city,omitempty" valid:"length(0|20)"`         // 20
	Region       string `xml:"region,omitempty" valid:"length(0|20)"`       // 20
	PostCode     string `xml:"postCode,omitempty" valid:"length(0|10)"`     // 10,
	CountryCode  string `xml:"countryCode" valid:"length(1|2)"`             //  2, mandatory
}

// AlertType - TODO
type AlertType int

// AlertType Constants
const (
	Telephone   AlertType = 1
	SMS         AlertType = 2
	Email       AlertType = 3
	EmailAndSMS AlertType = 4
)

// Customer - TODO
type Customer struct {
	Address             *Address  `xml:"address"`                                            // mandatory
	HomePhoneNo         string    `xml:"homePhoneNo,omitempty" valid:"length(0|15)"`         // 15
	WorkPhoneNo         string    `xml:"workPhoneNo,omitempty" valid:"length(0|15)"`         // 15
	MobilePhoneNo       string    `xml:"mobilePhoneNo,omitempty" valid:"length(0|15)"`       // 15
	FaxNo               string    `xml:"faxNo,omitempty" valid:"length(0|15)"`               // 15
	Email               string    `xml:"email,omitempty" valid:"email,length(0|80)"`         // 80
	CustomerReference1  string    `xml:"customerReference1" valid:"length(1|20)"`            // 20, mandatory
	CustomerReference2  string    `xml:"customerReference2,omitempty" valid:"length(1|20)"`  // 20
	CustomerAlertType   AlertType `xml:"customerAlertType,omitempty"`                        // 1
	CustomerAlertGroup  string    `xml:"customerAlertGroup,omitempty" valid:"length(0|4)"`   // 4
	DeliveryMessage     string    `xml:"deliveryMessage,omitempty" valid:"length(0|32)"`     // 32
	SpecialInstruction1 string    `xml:"specialInstruction1,omitempty" valid:"length(0|32)"` // 32
	SpecialInstruction2 string    `xml:"specialInstruction2,omitempty" valid:"length(0|32)"` // 32
}

// Diversions - TODO
type Diversions struct {
	ExcludeCancelDelivery bool `xml:"excludeCancelDelivery,omitempty"`
	ExcludeLaterDate      bool `xml:"excludeLaterDate,omitempty"`
	ExcludeNeighbours     bool `xml:"excludeNeighbours,omitempty"`
	ExcludeSafePlace      bool `xml:"excludeSafePlace,omitempty"`
	ExcludeParcelshop     bool `xml:"excludeParcelShop,omitempty"`
	ExcludeRetailStore    bool `xml:"excludeRetailStore,omitempty"`
}

// DeliveryRoutingRequestEntry - TODO
type DeliveryRoutingRequestEntry struct {
	addressValidationRequired bool           `xml:"omitempty"`
	Customer                  *Customer      `xml:"customer"` // mandatory
	Parcel                    *Parcel        `xml:"parcel"`   // mandatory
	Diversions                *Diversions    `xml:"diversions"`
	Services                  *Services      `xml:"services"`
	SenderAddress             *SenderAddress `xml:"senderAddress,omitempty"`
	ProductCode               int            `xml:"productCode,omitempty" valid:"length(0|10)` // 10
	ExpectedDespatchDate      time.Time      `xml:"expectedDespatchDate"`                      // mandatory
	RequiredDate              time.Time      `xml:"requiredDate,omitempty"`
	CountryOfOrigin           string         `xml:"countryOfOrigin" valid:"iso3166Alpha2"`         // 2, mandatory // TODO iso3166 doesnt exist
	WarehouseNo               int            `xml:"warehouseNo,omitempty" valid:"length(0|6)`      // 6, not currently used
	CarrierCode               string         `xml:"carrierCode,omitempty" valid:"length(0|6)`      // 6, not currently used
	DeliveryMethod            string         `xml:"deliveryMethod,omitempty" valid:"length(0|3)`   // 3, not currently used
	MultiplePartsID           string         `xml:"multiplePartsID,omitempty" valid:"length(0|50)` // 50
}

// DeliveryRoutingRequest - The request to Hermes for delivery info.
type DeliveryRoutingRequest struct {
	XMLName                       xml.Name                      `xml:"deliveryRoutingRequest"`
	ClientID                      string                        `xml:"clientId" valid:"length(1|3)`          // max 3, mandatory
	ClientName                    string                        `xml:"clientName" valid:"length(1|32)"`      // 32, mandatory
	ChildClientID                 string                        `xml:"childClientId" valid:"length(0|3)"`    // 3
	ChildClientName               string                        `xml:"childClientName" valid:"length(0|32)"` // 32
	BatchNumber                   string                        `xml:"batchNumber"`                          //5
	CreationDate                  time.Time                     `xml:"creationDate"`
	RoutingStartDate              time.Time                     `xml:"routingStartDate"`
	UserID                        string                        `xml:"userId" valid:"length(0|32)"`               // 32
	SourceOfRequest               string                        `xml:"sourceOfRequest" valid:"matches(CLIENTWS)"` // 8, mandatory
	DeliveryRoutingRequestEntries []DeliveryRoutingRequestEntry `xml:"deliveryRoutingRequestEntries>deliveryRoutingRequestEntry"`
}

// Titles - TODO
type Titles struct {
	SenderAddressTitle      string `xml:"senderAddressTitle,omitempty" valid:"length(0|32)"`
	DestinationAddressTitle string `xml:"destinationAddressTitle,omitempty" valid:"length(0|32)"`
	Entity1Title            string `xml:"entity2Title,omitempty" valid:"length(0|32)"`
	Entity2Title            string `xml:"entity4Title,omitempty" valid:"length(0|32)"`
	Entity3Title            string `xml:"entity3Title,omitempty" valid:"length(0|32)"`
	Entity4Title            string `xml:"entity4Title,omitempty" valid:"length(0|32)"`
}

// Barcode - TODO
type Barcode struct {
	BarcodeNumber    string `xml:"barcodeNumber" valid:"length(1|30)"`
	BarcodeLength    int    `xml:"barcodeLength"`
	BarcodeSymbology string `xml:"barcodeSymbology" valid:"length(1)"`
	BarcodeDisplay   string `xml:"barcodeDisplay" valid:"length(1|35)"`
}

// ServiceDescription - TODO
type ServiceDescription struct {
	ServiceDescriptionText string `xml:"serviceDescriptionText" valid:"length(1|50)"`
	ServiceLogoRef         string `xml:"serviceLogoRef,omitempty" valid:"length(0|50)"`
	ServicePosition        int    `xml:"servicePosition"` // mandatory
}

// Carrier - TODO
type Carrier struct {
	CarrierID           string               `xml:"carrierId,omitempty" valid:"length(0|6)"`
	CarrierName         string               `xml:"carrierName,omitempty" valid:"length(0|32)"`
	CarrierLogoRef      string               `xml:"carrierLogoRef,omitempty" valid:"length(0|50)"`
	DeliveryMethodDesc  string               `xml:"deliveryMethodDesc,omitempty" valid:"length(0|32)"`
	Barcode1            Barcode              `xml:"barcode1,omitempty"`
	Barcode2            Barcode              `xml:"barcode2,omitempty"`
	SortLevel1          string               `xml:"sortLevel1,omitempty" valid:"length(0|32)"`
	SortLevel2          string               `xml:"sortLevel2,omitempty" valid:"length(0|32)"`
	SortLevel3          string               `xml:"sortLevel3,omitempty" valid:"length(0|32)"`
	SortLevel4          string               `xml:"sortLevel4,omitempty" valid:"length(0|32)"`
	SortLevel5          string               `xml:"sortLevel5,omitempty" valid:"length(0|32)"`
	SortLevel6          string               `xml:"sortLevel6,omitempty" valid:"length(0|32)"`
	SortLevel7          string               `xml:"sortLevel7,omitempty" valid:"length(0|32)"`
	SortLevel8          string               `xml:"sortLevel8,omitempty" valid:"length(0|32)"`
	SortLevel9          string               `xml:"sortLevel9,omitempty" valid:"length(0|32)"`
	SortLevel10         string               `xml:"sortLevel10,omitempty" valid:"length(0|32)"`
	NodeName            string               `xml:"nodeName,omitempty" valid:"length(0|50)"`
	Address             ResponseAddress      `xml:"address,omitempty" valid:"length(0|32)"`
	ServiceDescriptions []ServiceDescription `xml:"serviceDescriptions,omitempty" valid:"length(0|32)"`
}

// Carriers - TODO
type Carriers struct {
	Carrier1     Carrier `xml:"carrier1"`
	Carrier2     Carrier `xml:"carrier1"`
	LabelImage   []byte  `xml:"labelImage"`
	Entity1Value string  `xml:"entity1Value" valid:"length(0|32)"`
	Entity2Value string  `xml:"entity2Value" valid:"length(0|32)"`
	Entity3Value string  `xml:"entity3Value" valid:"length(0|32)"`
	Entity4Value string  `xml:"entity4Value" valid:"length(0|32)"`
	Titles       Titles  `xml:"titles"`
}

// ResponseAddress - TODO
type ResponseAddress struct {
	Address1Line       string `xml:"addressLine1" valid:"length(0|50)"`
	Address2Line       string `xml:"addressLine2" valid:"length(0|50)"`
	Address3Line       string `xml:"addressLine3" valid:"length(0|50)"`
	Address4Line       string `xml:"addressLine4" valid:"length(0|50)"`
	Address5Line       string `xml:"addressLine5" valid:"length(0|50)"`
	Address6Line       string `xml:"addressLine6" valid:"length(0|50)"`
	Address7Line       string `xml:"addressLine7" valid:"length(0|50)"`
	Address8Line       string `xml:"addressLine8" valid:"length(0|50)"`
	CustomerReference1 string `xml:"customerReference1" valid:"length(0|20)"`
	CustomerReference2 string `xml:"customerReference2" valid:"length(0|20)"`
}

// RoutingResponseEntry - TODO
type RoutingResponseEntry struct {
	SenderAddress       ResponseAddress `xml:"senderAddress"`
	DestinationAddress  ResponseAddress
	OutboundCarriers    Carriers
	InboundCarriers     Carriers
	ServiceDescriptions []ServiceDescription
	Weight              string    `xml:"weight" valid:"length(0|10)"`
	Value               string    `xml:"value" valid:"length(0|10)"`
	Entity1Value        string    `xml:"entity1Value" valid:"length(0|32)"`
	Entity2Value        string    `xml:"entity2Value" valid:"length(0|32)"`
	Entity3Value        string    `xml:"entity3Value" valid:"length(0|32)"`
	Entity4Value        string    `xml:"entity4Value" valid:"length(0|32)"`
	ErrorMessages       []Message `xml:"errorMessages"`
	WarningMessages     []Message `xml:"warningMessages"`
	Titles              Titles    `xml:"titles"`
	Process             string    `xml:"process" valid:"length(0|32)"`
}

// Message - TODO
type Message struct {
	ErrorCode        int    `xml:"errorCode"`
	ErrorDescription string `xml:"errorDescription" valid:"length(1|50)"`
}

// RoutingResponse - TODO
type RoutingResponse struct {
	ClientID               string                 `xml:"clientId" valid:"length(1|3)"`
	ClientName             string                 `xml:"clientName" valid:"length(1|32)"`
	ChildClientID          string                 `xml:"childClientId" valid:"length(0|3)"`
	ChildClientName        string                 `xml:"childClientName" valid:"length(0|32)"`
	ClientLogoRef          string                 `xml:"clientLogoRef" valid:"length(0|50)"`
	BatchNumber            string                 `xml:"batchNumber"` // should be "number" - not currently used though.
	CreationDate           time.Time              `xml:"creatingDate"`
	RoutingResponseEntries []RoutingResponseEntry `xml:"routingResponseEntries>routingResponseEntry"`
}

// Client - Hermes Distribution Interface Client
type Client struct {
	id     string
	name   string
	userID string
}

// NewClient - create a new client.
func NewClient(ID string, name string, userID string) *Client {
	return &Client{
		id:     ID,
		name:   name,
		userID: userID,
	}
}

// NewDeliveryRoutingRequest - create a new deliveryRoutingRequest
func (c *Client) NewDeliveryRoutingRequest() DeliveryRoutingRequest {
	return DeliveryRoutingRequest{
		ClientID:         c.id,
		UserID:           c.userID,
		ClientName:       c.name,
		CreationDate:     time.Now(),
		RoutingStartDate: time.Now(),
		SourceOfRequest:  "CLIENTWS",
	}
}

// NewEntry - create & add a new entry to the request.
//func NewEntry() *DeliveryRoutingRequestEntry {
//	return new(DeliveryRoutingRequestEntry)
//}

// AddEntry - Add a DeliveryRoutingRequestEntry.
func (r *DeliveryRoutingRequest) AddEntry(entry DeliveryRoutingRequestEntry) {
	r.DeliveryRoutingRequestEntries = append(r.DeliveryRoutingRequestEntries, entry)
}

// Valid - Check that a request has all mandatory fields filled in.
func (r *DeliveryRoutingRequest) Valid() error {
	return valid(r)
}

// Call - Perform the actual request.
func (r *DeliveryRoutingRequest) Call() (*RoutingResponse, error) {
	err := r.Valid()
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	enc.Indent("  ", "    ")
	err = enc.Encode(r)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://sit.hermes-europe.co.uk/routing/service/rest/v3/validateDeliveryAddress", &buf)
	req.SetBasicAuth("***REMOVED***", "***REMOVED***")
	resp, err := client.Do(req)
	fmt.Println(resp)
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	var routingResp RoutingResponse
	err = xml.NewDecoder(resp.Body).Decode(&routingResp)
	if err != nil {
		return nil, err
	}
	err = valid(routingResp)
	if err != nil {
		return nil, err
	}
	errors := routingResp.RoutingResponseEntries[0].ErrorMessages
	if len(errors) > 0 {
		e := errors[0]
		return nil, fmt.Errorf("Hermes Distribution Interface error: %v:%v", e.ErrorCode, e.ErrorDescription)
	}
	return &routingResp, nil
}
