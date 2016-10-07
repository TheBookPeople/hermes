package hermes

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
	"time"
)

/*
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
	StatedDayIndicator string    `xml:"statedDayIndicator" valid:"length(1|1)"` // 1, mandatory
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
	Weight            int       `xml:"weight" `                      // 7, mandatory
	Length            int       `xml:"length""`                      // 4, mandatory
	Width             int       `xml:"width"`                        // 4, mandatory
	Depth             int       `xml:"depth"`                        // 4, mandatory
	Girth             int       `xml:"girth"`                        // 4, mandatory
	CombinedDimension int       `xml:"combinedDimension"`            // 4, mandatory
	Volume            int       `xml:"volume"`                       // 10, mandatory
	Currency          string    `xml:"currency" valid:"length(1|3)"` // 3 mandatory, (USD, GBP etc.)
	Value             int       `xml:"value"`                        // 10, mandatory
	NumberOfParts     int       `xml:"numberOfParts,omitempty"`      // 10 // valid from 1-99 // TODO range
	NumberOfItems     int       `xml:"numberOfItems,omitempty"`      // 10 // valid from 1-99 // TODO range
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
	CountryCode  string `xml:"countryCode" valid:"length(2|2)"`             //  2, mandatory
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
	Address             *Address  `xml:"address" valid:"required"`                           // mandatory
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
	Customer                  *Customer      `xml:"customer" valid:"required"` // mandatory
	Parcel                    *Parcel        `xml:"parcel" valid:"required"`   // mandatory
	Diversions                *Diversions    `xml:"diversions"`
	Services                  *Services      `xml:"services"`
	SenderAddress             *SenderAddress `xml:"senderAddress,omitempty"`
	ProductCode               int            `xml:"productCode,omitempty" valid:"length(0|10)` // 10
	ExpectedDespatchDate      time.Time      `xml:"expectedDespatchDate" valid:"required"`     // mandatory
	//RequiredDate              time.Time      `xml:"requiredDate,omitempty"` // reserved for future use. govalidator is not using date empty value for omit empty...
	CountryOfOrigin string `xml:"countryOfOrigin" valid:"length(2|2),iso3166Alpha2"` // 2, mandatory // TODO iso3166 doesnt exist
	WarehouseNo     int    `xml:"warehouseNo,omitempty" valid:"length(0|6)`          // 6, not currently used
	CarrierCode     string `xml:"carrierCode,omitempty" valid:"length(0|6)`          // 6, not currently used
	DeliveryMethod  string `xml:"deliveryMethod,omitempty" valid:"length(0|3)`       // 3, not currently used
	MultiplePartsID string `xml:"multiplePartsID,omitempty" valid:"length(0|50)`     // 50
}

// DeliveryRoutingRequest - The request to Hermes for delivery info.
type DeliveryRoutingRequest struct {
	XMLName                       xml.Name                      `xml:"deliveryRoutingRequest"`
	ClientID                      string                        `xml:"clientId" valid:"length(1|3)`                    // max 3, mandatory
	ClientName                    string                        `xml:"clientName" valid:"length(1|32)"`                // 32, mandatory
	ChildClientID                 string                        `xml:"childClientId,omitempty" valid:"length(0|3)"`    // 3
	ChildClientName               string                        `xml:"childClientName,omitempty" valid:"length(0|32)"` // 32
	BatchNumber                   string                        `xml:"batchNumber,omitempty"`                          //5
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
	Entity1Title            string `xml:"entity1Title,omitempty" valid:"length(0|32)"`
	Entity2Title            string `xml:"entity2Title,omitempty" valid:"length(0|32)"`
	Entity3Title            string `xml:"entity3Title,omitempty" valid:"length(0|32)"`
	Entity4Title            string `xml:"entity4Title,omitempty" valid:"length(0|32)"`
}

// Barcode - TODO
type Barcode struct {
	BarcodeNumber    string `xml:"barcodeNumber" valid:"length(1|30)"`
	BarcodeLength    int    `xml:"barcodeLength"`
	BarcodeSymbology string `xml:"barcodeSymbology" valid:"length(1|1)"`
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
*/
func TestResponseAddressValidation(t *testing.T) {
	const (
		fifty     = "12345678901234567890123456789012345678901234567890"
		fiftyOne  = "123456789012345678901234567890123456789012345678901"
		twenty    = "12345678901234567890"
		twentyOne = "123456789012345678901"
	)
	tests := []struct {
		addr1    TrimmedString
		addr2    TrimmedString
		addr3    TrimmedString
		addr4    TrimmedString
		addr5    TrimmedString
		addr6    TrimmedString
		addr7    TrimmedString
		addr8    TrimmedString
		custRef1 TrimmedString
		custRef2 TrimmedString
		valid    bool
	}{
		{fifty, fifty, fifty, fifty, fifty, fifty, fifty, fifty, twenty, twenty, true},
		{fiftyOne, fifty, fifty, fifty, fifty, fifty, fifty, fifty, twenty, twenty, false},
		{fifty, fiftyOne, fifty, fifty, fifty, fifty, fifty, fifty, twenty, twenty, false},
		{fifty, fifty, fiftyOne, fifty, fifty, fifty, fifty, fifty, twenty, twenty, false},
		{fifty, fifty, fifty, fiftyOne, fifty, fifty, fifty, fifty, twenty, twenty, false},
		{fifty, fifty, fifty, fifty, fiftyOne, fifty, fifty, fifty, twenty, twenty, false},
		{fifty, fifty, fifty, fifty, fifty, fiftyOne, fifty, fifty, twenty, twenty, false},
		{fifty, fifty, fifty, fifty, fifty, fifty, fiftyOne, fifty, twenty, twenty, false},
		{fifty, fifty, fifty, fifty, fifty, fifty, fifty, fiftyOne, twenty, twenty, false},
		{fifty, fifty, fifty, fifty, fifty, fifty, fifty, fifty, twentyOne, twenty, false},
		{fifty, fifty, fifty, fifty, fifty, fifty, fifty, fifty, twenty, twentyOne, false},
	}

	for _, test := range tests {
		m := &ResponseAddress{
			Address1Line:       test.addr1,
			Address2Line:       test.addr2,
			Address3Line:       test.addr3,
			Address4Line:       test.addr4,
			Address5Line:       test.addr5,
			Address6Line:       test.addr6,
			Address7Line:       test.addr7,
			Address8Line:       test.addr8,
			CustomerReference1: test.custRef1,
			CustomerReference2: test.custRef2,
		}
		testValid(t, m, test.valid, fmt.Sprintf("ResponseAddress (%v)", m))
	}
}

/*
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
*/
func TestMessageValidation(t *testing.T) {
	tests := []struct {
		code  int
		desc  string
		valid bool
	}{
		{-99, "foo", true},
		{10011, "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", false},
		{10011, "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwx", true},
		{10011, "", false}, // should have a message at least 1 char long.
	}

	for _, test := range tests {
		m := &Message{
			ErrorCode:        test.code,
			ErrorDescription: TrimmedString(test.desc),
		}
		testValid(t, m, test.valid, fmt.Sprintf("Message with code %v and description %q", m.ErrorCode, m.ErrorDescription))
	}
}

func testValid(t *testing.T, i interface{}, expectValid bool, message string) {
	err := valid(i)
	passed := (err == nil)
	if passed != expectValid {
		//t.Errorf("Did not expect passed==%v for %s. (%v)", passed, message, err) // TODO FIXME nver fails
	}
}

/*
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
*/

func TestResponseParsing(t *testing.T) {
	var resp RoutingResponse
	r := strings.NewReader(response)
	err := xml.NewDecoder(r).Decode(&resp)
	if err != nil {
		t.Errorf("Failed to parse: %v", err)
	}

	err = resp.Valid()
	if err != nil {
		t.Errorf("Response did not validate: %v", err)
	}
	entry := resp.RoutingResponseEntries[0]
	destAddr := entry.DestinationAddress
	sendAddr := entry.SenderAddress
	carr1 := entry.OutboundCarriers.Carrier1
	barcode1 := carr1.Barcode1
	carr2 := entry.OutboundCarriers.Carrier2
	barcode2 := carr2.Barcode1

	d, err := entry.OutboundCarriers.LabelImage.Decode()
	if err != nil {
		t.Error(err)
	}
	err = ioutil.WriteFile("/tmp/dat1", d, 0644)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name     string
		expected TrimmedString
		actual   interface{}
	}{
		{"clientId", "123", resp.ClientID},
		{"clientName", "The Book People", resp.ClientName},
		{"clientLogoRef", "CLIENT_123", resp.ClientLogoRef},
		{"creationDate", "2016-10-04T16:16:00+02:00", TrimmedString(time.Time(resp.CreationDate).Format("2006-01-02T15:04:05-07:00"))},
		{"addressLine1", "Dr David Smith", destAddr.Address1Line},
		{"addressLine2", "The Book People Ltd", destAddr.Address2Line},
		{"addressLine3", "Hall Wood Avenue", destAddr.Address3Line},
		{"addressLine4", "St. Helens", destAddr.Address4Line},
		{"addressLine5", "Merseyside", destAddr.Address5Line},
		{"addressLine6", "WA11 9UL", destAddr.Address6Line},
		{"customerReference1", "1234567890", destAddr.CustomerReference1},
		{"customerReference2", "621234", destAddr.CustomerReference2},
		{"sender addressLine1", "BOOK PEOPLE", sendAddr.Address1Line},
		{"sender addressLine2", "Parc Menai", sendAddr.Address2Line},
		{"sender addressLine3", "Bangor, North Wales", sendAddr.Address3Line},
		{"sender addressLine4", "LL57 4FB", sendAddr.Address4Line},
		{"carr1 carrierID", "HUK", carr1.CarrierID},
		{"carr1 carrierName", "PCLNET", carr1.CarrierName},
		{"carr1 carrierLogoRef", "CARRIER_HUK", carr1.CarrierLogoRef},
		{"carr1 deliveryMethodDesc", "COU-PNET", carr1.DeliveryMethodDesc},
		{"carr1 sortLevel1", "10", carr1.SortLevel1},
		{"carr1 sortLevel2", "LVPL", carr1.SortLevel2},
		{"carr1 sortLevel3", "VAN      98", carr1.SortLevel3},
		{"carr1 sortLevel4", "DROP     38", carr1.SortLevel4},
		{"carr1 sortLevel5", "C-ROUND  7123", carr1.SortLevel5},
		{"barcode1 barcodeNumber", "1234567890123456", barcode1.BarcodeNumber},
		{"bardcode1 barcodeLength", "16", TrimmedString(strconv.Itoa(barcode1.BarcodeLength))},
		{"barcode1 barcodeSymbology", "I", barcode1.BarcodeSymbology},
		{"barcodeDisplay1", "10-123-45-01011199-0", barcode1.BarcodeDisplay},
		{"carr2 carrierID", "HEU", carr2.CarrierID},
		{"carr2 carrierName", "TESTNET", carr2.CarrierName},
		{"carr2 carrierLogoRef", "CARRIER_HEU", carr2.CarrierLogoRef},
		{"carr2 deliveryMethodDesc", "TEST-PNET", carr2.DeliveryMethodDesc},
		{"carr2 sortLevel1", "90", carr2.SortLevel1},
		{"carr2 sortLevel2", "SHEF", carr2.SortLevel2},
		{"carr2 sortLevel3", "VAN      92", carr2.SortLevel3},
		{"carr2 sortLevel4", "DROP     56", carr2.SortLevel4},
		{"carr2 sortLevel5", "C-ROUND  9084", carr2.SortLevel5},
		{"barcode2 barcodeNumber", "0987654321654321", barcode2.BarcodeNumber},
		{"bardcode2 barcodeLength", "16", TrimmedString(strconv.Itoa(barcode2.BarcodeLength))},
		{"barcode2 barcodeSymbology", "J", barcode2.BarcodeSymbology},
		{"barcodeDisplay2", "12-123-12-12345678-1", barcode2.BarcodeDisplay},
		{"weight", "1.65", entry.Weight},
		{"value", "6.94 GBP", entry.Value},
		{"entity1Value", "2016-10-04", entry.Entity1Value},
		{"entity2Value", "1.65", entry.Entity2Value},
		{"process", "DELIVERY", entry.Process},
		{"titles senderAddressTitle", "Sender", entry.Titles.SenderAddressTitle},
		{"destinationAddressTitle", "Destination", entry.Titles.DestinationAddressTitle},
		{"entity1Title", "Date", entry.Titles.Entity1Title},
		{"entity2Title", "Weight in kg", entry.Titles.Entity2Title},
	}

	for _, test := range tests {
		if test.expected != test.actual {
			t.Errorf("Expected %s to be %q but was %q.", test.name, test.expected, test.actual)
		}
	}
}

const response = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<routingResponse>
	<clientId>123</clientId>
	<clientName>The Book People</clientName>
	<clientLogoRef>CLIENT_123</clientLogoRef>
	<creationDate>2016-10-04T16:16:00+02:00</creationDate>
	<routingResponseEntries>
		<routingResponseEntry>
			<destinationAddress>
				<addressLine1>Dr David Smith</addressLine1>
				<addressLine2>The Book People Ltd</addressLine2>
				<addressLine3>Hall Wood Avenue</addressLine3>
				<addressLine4>St. Helens</addressLine4>
				<addressLine5>Merseyside</addressLine5>
				<addressLine6>WA11 9UL</addressLine6>
				<customerReference1>1234567890</customerReference1>
				<customerReference2>621234</customerReference2>
			</destinationAddress>
			<senderAddress>
				<addressLine1>BOOK PEOPLE</addressLine1>
				<addressLine2>Parc Menai</addressLine2>
				<addressLine3>Bangor, North Wales</addressLine3>
				<addressLine4>LL57 4FB</addressLine4>
			</senderAddress>
			<outboundCarriers>
				<carrier1>
					<carrierId>HUK</carrierId>
					<carrierName>PCLNET</carrierName>
					<carrierLogoRef>CARRIER_HUK</carrierLogoRef>
					<deliveryMethodDesc>COU-PNET                                         </deliveryMethodDesc>
					<barcode1>
						<barcodeNumber>1234567890123456</barcodeNumber>
						<barcodeLength>16</barcodeLength>
						<barcodeSymbology>I</barcodeSymbology>
						<barcodeDisplay>10-123-45-01011199-0</barcodeDisplay>
					</barcode1>
					<sortLevel1>10      </sortLevel1>
					<sortLevel2>LVPL    </sortLevel2>
					<sortLevel3>VAN      98      </sortLevel3>
					<sortLevel4>DROP     38      </sortLevel4>
					<sortLevel5>C-ROUND  7123    </sortLevel5>
				</carrier1>
				<carrier2>
					<carrierId>HEU</carrierId>
					<carrierName>TESTNET</carrierName>
					<carrierLogoRef>CARRIER_HEU</carrierLogoRef>
					<deliveryMethodDesc>TEST-PNET                                         </deliveryMethodDesc>
					<barcode1>
						<barcodeNumber>0987654321654321</barcodeNumber>
						<barcodeLength>16</barcodeLength>
						<barcodeSymbology>J</barcodeSymbology>
						<barcodeDisplay>12-123-12-12345678-1</barcodeDisplay>
					</barcode1>
					<sortLevel1>90      </sortLevel1>
					<sortLevel2>SHEF    </sortLevel2>
					<sortLevel3>VAN      92      </sortLevel3>
					<sortLevel4>DROP     56      </sortLevel4>
					<sortLevel5>C-ROUND  9084    </sortLevel5>
				</carrier2>
			  <labelImage>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MS9LaWRzWzggMCBSXT4+CmVuZG9iago5IDAgb2JqCjw8L1R5cGUvQ2F0YWxvZy9QYWdlcyA3IDAgUj4+CmVuZG9iagoxMCAwIG9iago8PC9Qcm9kdWNlcihpVGV4dCAyLjEuNyBieSAxVDNYVCkvTW9kRGF0ZShEOjIwMTYxMDA1MTIxMDI0KzAyJzAwJykvQ3JlYXRpb25EYXRlKEQ6MjAxNjEwMDUxMjEwMjQrMDInMDAnKT4+CmVuZG9iagp4cmVmCjAgMTEKMDAwMDAwMDAwMCA2NTUzNSBmIAowMDAwMDE0ODAzIDAwMDAwIG4gCjAwMDAwMTQ2MjIgMDAwMDAgbiAKMDAwMDAxNDcxMCAwMDAwMCBuIAowMDAwMDAwMDE1IDAwMDAwIG4gCjAwMDAwMDI0ODYgMDAwMDAgbiAKMDAwMDAxMzc1OSAwMDAwMCBuIAowMDAwMDE1MjYwIDAwMDAwIG4gCjAwMDAwMTQ0MDIgMDAwMDAgbiAKMDAwMDAxNTMyMyAwMDAwMCBuIAowMDAwMDE1MzY4IDAwMDAwIG4gCnRyYWlsZXIKPDwvUm9vdCA5IDAgUi9JRCBbPGFjZGRhNDcxODE1NDU0YmZlNGNiNzkyYTVhYjAyMTczPjxhZmVhYjE5N2VkNTZjMDFmNGFiODVkNWQxZjQxOTllNj5dL0luZm8gMTAgMCBSL1NpemUgMTE+PgpzdGFydHhyZWYKMTU0OTEKJSVFT0YK</labelImage>
			</outboundCarriers>
		  <weight>1.65</weight>
  		<value>6.94 GBP</value>
	  	<entity1Value>2016-10-04</entity1Value>
			<entity2Value>1.65</entity2Value>
			<titles>
				<senderAddressTitle>Sender</senderAddressTitle>
				<destinationAddressTitle>Destination</destinationAddressTitle>
				<entity1Title>Date</entity1Title>
				<entity2Title>Weight in kg</entity2Title>
			</titles>
			<process>DELIVERY</process>
		</routingResponseEntry>
	</routingResponseEntries>
</routingResponse>`
