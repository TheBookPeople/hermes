package hermes

import (
	"encoding/base64"
	"encoding/xml"
	"strings"
	"time"
)

// TrimmedString - trims whitespace when marshalling.
type TrimmedString string

// MarshalXML - trims whitespace when marshalling.
func (ts *TrimmedString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(strings.TrimSpace(string(*ts)), start)
}

// UnmarshalXML - trims whitespace when Unmarshalling.
func (ts *TrimmedString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	*ts = TrimmedString(strings.TrimSpace(v))
	return nil
}

const (
	timeFormat     = "2006-01-02T15:04:05Z" // Was -07:00 instead of Z
	longTimeFormat = "2006-01-02T15:04:05-07:00"
)

// Time - Wraps time but marshalls to expected format.
type Time time.Time

func (t Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(time.Time(t).Format(timeFormat), start)
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string

	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	format := longTimeFormat
	if strings.HasSuffix(v, "Z") {
		format = timeFormat
	}

	parsedTime, err := time.Parse(format, v)
	if err != nil {
		return err
	}
	*t = Time(parsedTime)
	return nil
}

func Now() Time {
	return Time(time.Now())
}

// IdentityService - TODO
type IdentityService struct {
	IDCardNo      TrimmedString `xml:"idCardNo,omitempty" valid:"length(0|20)"`     // 20
	IdcardType    TrimmedString `xml:"idCardType,omitempty" valid:"length(0|3)"`    // 3
	AgeValidation int           `xml:"ageValidation,omitempty" valid:"length(0|3)"` // 3 (0-100 range) // TODO range
	DateOfBirth   Time
	Pin           TrimmedString `xml:"pin,omitempty" valid:"length(0|35)"` // 35
	Module        TrimmedString `xml:"Module" valid:"length(0|3)"`         // 3
}

//RetailStoreService - TODO
type RetailStoreService struct {
	RetailStoreID TrimmedString `xml:"retailStoreId" valid:"length(1|20)"` // 20, mandatory
	Address       Address       `xml:"address,omitempty"`
}

//ParcelShopService  - TODO
type ParcelShopService struct {
	ParcelShopID TrimmedString `xml:"parcelShopId" valid:"length(1|20)"` // 20, mandatory
	Address      *Address      `xml:"address,omitempty"`
}

// CashOnDelivery - TODO
type CashOnDelivery struct {
	CashValue            int           `xml:"cashValue,omitempty" valid:"length(0|10)"`           // 10 (in pence etc.)
	CashCurrency         TrimmedString `xml:"cashCurrency,omitempty valid:length(0|3)"`           // 3 (e.g. EUR, GBP)
	BankTransferValue    int           `xml:"bankTransferValue,omitempty" valid:"length(0|10)"`   // 10 (in pence etc.)
	BankTransferCurrency TrimmedString `xml:"bankTransferCurrency,omitempty" valid:"length(0|3)"` // 3
}

// StatedDay - TODO
type StatedDay struct {
	StatedDayIndicator TrimmedString `xml:"statedDayIndicator" valid:"length(1|1)"` // 1, mandatory
	StatedDate         Time          `xml:"statedDate,omitempty"`
}

// StatedTime - TODO
type StatedTime TrimmedString

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
	SkuCode        TrimmedString `xml:"skuCode" valid:"length(1|30)"`          // 30, mandatory
	SkuDescription TrimmedString `xml:"skuDescription" valid:"length(1|2000)"` // 2000, mandatory
	HsCode         TrimmedString `xml:"hsCode" valid:"length(1|10)"`           // 10, mandatory
	Value          int           `xml:"value"`                                 // 10, mandatory // TODO int length
}

type dutyPaid string

const (
	Paid   dutyPaid = "P"
	Unpaid dutyPaid = "U"
)

// Parcel - TODO
type Parcel struct {
	Weight            int           `xml:"weight"`                       // 7, mandatory
	Length            int           `xml:"length"`                       // 4, mandatory
	Width             int           `xml:"width"`                        // 4, mandatory
	Depth             int           `xml:"depth"`                        // 4, mandatory
	Girth             int           `xml:"girth"`                        // 4, mandatory
	CombinedDimension int           `xml:"combinedDimension"`            // 4, mandatory
	Volume            int           `xml:"volume"`                       // 10, mandatory
	Currency          TrimmedString `xml:"currency" valid:"length(1|3)"` // 3 mandatory, (USD, GBP etc.)
	Value             int           `xml:"value"`                        // 10, mandatory
	NumberOfParts     int           `xml:"numberOfParts,omitempty"`      // 10 // valid from 1-99 // TODO range
	NumberOfItems     int           `xml:"numberOfItems,omitempty"`      // 10 // valid from 1-99 // TODO range
	HangingGarment    bool          `xml:"hangingGarment,omitempty"`
	TheftRisk         bool          `xml:"theftRisk,omitempty"`     // Not currently used.
	MultipleParts     bool          `xml:"multipleParts,omitempty"` // Not currently used.
	Catalogue         int           `xml:"catalogue,omitempty"`
	Description       int           `xml:"description,omitempty" valid:"length(0|32)"`    // 32
	OriginOfParcel    int           `xml:"originOfParcel,omitempty" valid:"length(0|32)"` // 32
	DutyPaid          dutyPaid      `xml:"dutyPaid,omitempty" valid:"length(0|1)"`        // 1, mandatory if non EU U = unpaid, P = paid
	Contents          []Content     `xml:"contents>content"`
}

// SenderAddress - TODO
type SenderAddress struct {
	AddressLine1 TrimmedString `xml:"addressLine1,omitempty" valid:"length(0|50)"` // 50
	AddressLine2 TrimmedString `xml:"addressLine2,omitempty" valid:"length(0|50)"` // 50
	AddressLine3 TrimmedString `xml:"addressLine3,omitempty" valid:"length(0|50)"` // 50
	AddressLine4 TrimmedString `xml:"addressLine4,omitempty" valid:"length(0|50)"` // 50
}

// Address - TODO
type Address struct {
	Title        TrimmedString `xml:"title,omitempty" valid:"length(0|20)"`
	FirstName    TrimmedString `xml:"firstName,omitempty" valid:"length(0|50)"`
	LastName     TrimmedString `xml:"lastName" valid:"length(1|50)"`
	HouseNo      TrimmedString `xml:"houseNo,omitempty" valid:"length(0|10)"`
	HouseName    TrimmedString `xml:"houseName,omitempty" valid:"length(0|32)"`
	StreetName   TrimmedString `xml:"streetName" valid:"length(1|50)"`
	AddressLine1 TrimmedString `xml:"addressLine1,omitempty" valid:"length(0|50)"`
	AddressLine2 TrimmedString `xml:"addressLine2,omitempty" valid:"length(0|50)"`
	AddressLine3 TrimmedString `xml:"addressLine3,omitempty" valid:"length(0|50)"`
	City         TrimmedString `xml:"city" valid:"length(1|50)"`
	Region       TrimmedString `xml:"region,omitempty" valid:"length(0|50)"`
	PostCode     TrimmedString `xml:"postCode,omitempty" valid:"length(0|10)"`
	CountryCode  TrimmedString `xml:"countryCode" valid:"length(2|2)"`
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
	Address             *Address      `xml:"address" valid:"required"`                           // mandatory
	HomePhoneNo         TrimmedString `xml:"homePhoneNo,omitempty" valid:"length(0|15)"`         // 15
	WorkPhoneNo         TrimmedString `xml:"workPhoneNo,omitempty" valid:"length(0|15)"`         // 15
	MobilePhoneNo       TrimmedString `xml:"mobilePhoneNo,omitempty" valid:"length(0|15)"`       // 15
	FaxNo               TrimmedString `xml:"faxNo,omitempty" valid:"length(0|15)"`               // 15
	Email               TrimmedString `xml:"email,omitempty" valid:"email,length(0|80)"`         // 80
	CustomerReference1  TrimmedString `xml:"customerReference1" valid:"length(1|20)"`            // 20, mandatory
	CustomerReference2  TrimmedString `xml:"customerReference2,omitempty" valid:"length(1|20)"`  // 20
	CustomerAlertType   AlertType     `xml:"customerAlertType,omitempty"`                        // 1
	CustomerAlertGroup  TrimmedString `xml:"customerAlertGroup,omitempty" valid:"length(0|4)"`   // 4
	DeliveryMessage     TrimmedString `xml:"deliveryMessage,omitempty" valid:"length(0|32)"`     // 32
	SpecialInstruction1 TrimmedString `xml:"specialInstruction1,omitempty" valid:"length(0|32)"` // 32
	SpecialInstruction2 TrimmedString `xml:"specialInstruction2,omitempty" valid:"length(0|32)"` // 32
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
	AddressValidationRequired bool           `xml:"addressValidationRequired,omitempty"`
	Customer                  *Customer      `xml:"customer" valid:"required"` // mandatory
	Parcel                    *Parcel        `xml:"parcel" valid:"required"`   // mandatory
	Diversions                *Diversions    `xml:"diversions"`
	Services                  *Services      `xml:"services"`
	SenderAddress             *SenderAddress `xml:"senderAddress,omitempty"`
	ProductCode               int            `xml:"productCode,omitempty" valid:"length(0|10)"` // 10
	ExpectedDespatchDate      Time           `xml:"expectedDespatchDate" valid:"required"`      // mandatory
	//RequiredDate              Time      `xml:"requiredDate,omitempty"` // reserved for future use. govalidator is not using date empty value for omit empty...
	CountryOfOrigin TrimmedString `xml:"countryOfOrigin" valid:"length(2|2)"`            // 2, mandatory
	WarehouseNo     int           `xml:"warehouseNo,omitempty" valid:"length(0|6)"`      // 6, not currently used
	CarrierCode     TrimmedString `xml:"carrierCode,omitempty" valid:"length(0|6)"`      // 6, not currently used
	DeliveryMethod  TrimmedString `xml:"deliveryMethod,omitempty" valid:"length(0|3)"`   // 3, not currently used
	MultiplePartsID TrimmedString `xml:"multiplePartsID,omitempty" valid:"length(0|50)"` // 50
}

// DeliveryRoutingRequest - The request to Hermes for delivery info.
type DeliveryRoutingRequest struct {
	XMLName                       xml.Name                      `xml:"deliveryRoutingRequest"`
	ClientID                      TrimmedString                 `xml:"clientId" valid:"length(1|3)"`                   // max 3, mandatory
	ClientName                    TrimmedString                 `xml:"clientName" valid:"length(1|32)"`                // 32, mandatory
	ChildClientID                 TrimmedString                 `xml:"childClientId,omitempty" valid:"length(0|3)"`    // 3
	ChildClientName               TrimmedString                 `xml:"childClientName,omitempty" valid:"length(0|32)"` // 32
	BatchNumber                   TrimmedString                 `xml:"batchNumber,omitempty"`                          //5
	CreationDate                  Time                          `xml:"creationDate"`
	RoutingStartDate              Time                          `xml:"routingStartDate"`
	UserID                        TrimmedString                 `xml:"userId" valid:"length(0|32)"`               // 32
	SourceOfRequest               TrimmedString                 `xml:"sourceOfRequest" valid:"matches(CLIENTWS)"` // 8, mandatory
	DeliveryRoutingRequestEntries []DeliveryRoutingRequestEntry `xml:"deliveryRoutingRequestEntries>deliveryRoutingRequestEntry"`
}

// Titles - TODO
type Titles struct {
	SenderAddressTitle      TrimmedString `xml:"senderAddressTitle,omitempty" valid:"length(0|32)"`
	DestinationAddressTitle TrimmedString `xml:"destinationAddressTitle,omitempty" valid:"length(0|32)"`
	Entity1Title            TrimmedString `xml:"entity1Title,omitempty" valid:"length(0|32)"`
	Entity2Title            TrimmedString `xml:"entity2Title,omitempty" valid:"length(0|32)"`
	Entity3Title            TrimmedString `xml:"entity3Title,omitempty" valid:"length(0|32)"`
	Entity4Title            TrimmedString `xml:"entity4Title,omitempty" valid:"length(0|32)"`
}

// Barcode - TODO
type Barcode struct {
	BarcodeNumber    TrimmedString `xml:"barcodeNumber" valid:"length(1|30)"`
	BarcodeLength    int           `xml:"barcodeLength"`
	BarcodeSymbology TrimmedString `xml:"barcodeSymbology" valid:"length(1|4)"` // Documented incorectly as 1. Can be C128 etc.
	BarcodeDisplay   TrimmedString `xml:"barcodeDisplay" valid:"length(1|35)"`
}

// ServiceDescription - TODO
type ServiceDescription struct {
	ServiceDescriptionText TrimmedString `xml:"serviceDescriptionText" valid:"length(1|50)"`
	ServiceLogoRef         TrimmedString `xml:"serviceLogoRef,omitempty" valid:"length(0|50)"`
	ServicePosition        int           `xml:"servicePosition"` // mandatory
}

// Carrier - TODO
type Carrier struct {
	CarrierID           TrimmedString        `xml:"carrierId,omitempty" valid:"length(0|6)"`
	CarrierName         TrimmedString        `xml:"carrierName,omitempty" valid:"length(0|32)"`
	CarrierLogoRef      TrimmedString        `xml:"carrierLogoRef,omitempty" valid:"length(0|50)"`
	DeliveryMethodDesc  TrimmedString        `xml:"deliveryMethodDesc,omitempty" valid:"length(0|50)"` // Documented incorrectly as 32.
	Barcode1            Barcode              `xml:"barcode1,omitempty"`
	Barcode2            Barcode              `xml:"barcode2,omitempty"`
	SortLevel1          TrimmedString        `xml:"sortLevel1,omitempty" valid:"length(0|32)"`
	SortLevel2          TrimmedString        `xml:"sortLevel2,omitempty" valid:"length(0|32)"`
	SortLevel3          TrimmedString        `xml:"sortLevel3,omitempty" valid:"length(0|32)"`
	SortLevel4          TrimmedString        `xml:"sortLevel4,omitempty" valid:"length(0|32)"`
	SortLevel5          TrimmedString        `xml:"sortLevel5,omitempty" valid:"length(0|32)"`
	SortLevel6          TrimmedString        `xml:"sortLevel6,omitempty" valid:"length(0|32)"`
	SortLevel7          TrimmedString        `xml:"sortLevel7,omitempty" valid:"length(0|32)"`
	SortLevel8          TrimmedString        `xml:"sortLevel8,omitempty" valid:"length(0|32)"`
	SortLevel9          TrimmedString        `xml:"sortLevel9,omitempty" valid:"length(0|32)"`
	SortLevel10         TrimmedString        `xml:"sortLevel10,omitempty" valid:"length(0|32)"`
	NodeName            TrimmedString        `xml:"nodeName,omitempty" valid:"length(0|50)"`
	Address             ResponseAddress      `xml:"address,omitempty" valid:"length(0|32)"`
	ServiceDescriptions []ServiceDescription `xml:"serviceDescriptions,omitempty" valid:"length(0|32)"`
}

type LabelImage []byte

func (li *LabelImage) Decode() ([]byte, error) {
	d := make([]byte, base64.StdEncoding.DecodedLen(len(*li)))
	_, err := base64.StdEncoding.Decode(d, *li)

	if err != nil {
		return nil, err
	}
	return d, nil
}

// Carriers - TODO
type Carriers struct {
	Carrier1     Carrier       `xml:"carrier1"`
	Carrier2     Carrier       `xml:"carrier2"`
	LabelImage   LabelImage    `xml:"labelImage"`
	Entity1Value TrimmedString `xml:"entity1Value" valid:"length(0|32)"`
	Entity2Value TrimmedString `xml:"entity2Value" valid:"length(0|32)"`
	Entity3Value TrimmedString `xml:"entity3Value" valid:"length(0|32)"`
	Entity4Value TrimmedString `xml:"entity4Value" valid:"length(0|32)"`
	Titles       Titles        `xml:"titles"`
}

// ResponseAddress - TODO
type ResponseAddress struct {
	Address1Line       TrimmedString `xml:"addressLine1" valid:"length(0|50)"`
	Address2Line       TrimmedString `xml:"addressLine2" valid:"length(0|50)"`
	Address3Line       TrimmedString `xml:"addressLine3" valid:"length(0|50)"`
	Address4Line       TrimmedString `xml:"addressLine4" valid:"length(0|50)"`
	Address5Line       TrimmedString `xml:"addressLine5" valid:"length(0|50)"`
	Address6Line       TrimmedString `xml:"addressLine6" valid:"length(0|50)"`
	Address7Line       TrimmedString `xml:"addressLine7" valid:"length(0|50)"`
	Address8Line       TrimmedString `xml:"addressLine8" valid:"length(0|50)"`
	CustomerReference1 TrimmedString `xml:"customerReference1" valid:"length(0|20)"`
	CustomerReference2 TrimmedString `xml:"customerReference2" valid:"length(0|20)"`
}

// RoutingResponseEntry - TODO
type RoutingResponseEntry struct {
	SenderAddress       ResponseAddress `xml:"senderAddress"`
	DestinationAddress  ResponseAddress `xml:"destinationAddress"`
	OutboundCarriers    Carriers        `xml:"outboundCarriers"`
	InboundCarriers     Carriers        `xml:"inboundCarriers"`
	ServiceDescriptions []ServiceDescription
	Weight              TrimmedString `xml:"weight" valid:"length(0|10)"`
	Value               TrimmedString `xml:"value" valid:"length(0|12)"`
	Entity1Value        TrimmedString `xml:"entity1Value" valid:"length(0|32)"`
	Entity2Value        TrimmedString `xml:"entity2Value" valid:"length(0|32)"`
	Entity3Value        TrimmedString `xml:"entity3Value" valid:"length(0|32)"`
	Entity4Value        TrimmedString `xml:"entity4Value" valid:"length(0|32)"`
	ErrorMessages       []Message     `xml:"errorMessages"`
	WarningMessages     []Message     `xml:"warningMessages"`
	Titles              Titles        `xml:"titles"`
	Process             TrimmedString `xml:"process" valid:"length(0|32)"`
}

// Message - TODO
type Message struct {
	ErrorCode        int           `xml:"errorCode"`
	ErrorDescription TrimmedString `xml:"errorDescription" valid:"length(1|50)"`
}

// RoutingResponse - TODO
type RoutingResponse struct {
	ClientID               TrimmedString          `xml:"clientId" valid:"length(1|3)"`
	ClientName             TrimmedString          `xml:"clientName" valid:"length(1|32)"`
	ChildClientID          TrimmedString          `xml:"childClientId" valid:"length(0|3)"`
	ChildClientName        TrimmedString          `xml:"childClientName" valid:"length(0|32)"`
	ClientLogoRef          TrimmedString          `xml:"clientLogoRef" valid:"length(0|50)"`
	BatchNumber            TrimmedString          `xml:"batchNumber"` // should be "number" - not currently used though.
	CreationDate           Time                   `xml:"creationDate"`
	RoutingResponseEntries []RoutingResponseEntry `xml:"routingResponseEntries>routingResponseEntry"`
}
