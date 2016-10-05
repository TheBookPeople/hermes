package hermes

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
)

const host = "sit.hermes-europe.co.uk"

// Client - Hermes Distribution Interface Client
type Client struct {
	id       string
	name     string
	userID   string
	password string
}

// NewClient - Create a new Client
func NewClient(userID, ID, name, password string) *Client {
	return &Client{
		id:       ID,
		name:     name,
		userID:   userID,
		password: password,
	}
}

// NewDeliveryRoutingRequest - create a new deliveryRoutingRequest
func (c *Client) NewDeliveryRoutingRequest() *DeliveryRoutingRequest {
	return &DeliveryRoutingRequest{
		ClientID:         c.id,
		UserID:           c.userID,
		ClientName:       c.name,
		CreationDate:     Now(),
		RoutingStartDate: Now(),
		SourceOfRequest:  "CLIENTWS",
	}
}

// ValidateDeliveryAddress - TODO
func (c *Client) ValidateDeliveryAddress(drr *DeliveryRoutingRequest) (*RoutingResponse, error) {
	return c.call(drr, "validateDeliveryAddress")
}

// DetermineDeliveryRouting - TODO
func (c *Client) DetermineDeliveryRouting(drr *DeliveryRoutingRequest) (*RoutingResponse, error) {
	return c.call(drr, "determineDeliveryRouting")
}

// RouteDeliveryCreatePreadvice - TODO
func (c *Client) RouteDeliveryCreatePreadvice(drr *DeliveryRoutingRequest, label bool, barcode bool) (*RoutingResponse, error) {
	command := "routeDeliveryCreatePreadvice"
	if label && barcode {
		command += "ReturnBarcodeAndLabel"
	} else if label {
		command += "AndLabel"
	} else if barcode {
		command += "ReturnBarcode"
	}
	return c.call(drr, command)
}

// Call - Perform the actual request.
func (c *Client) call(r *DeliveryRoutingRequest, command string) (*RoutingResponse, error) {
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
	fmt.Println(buf.String())
	httpClient := &http.Client{}

	url := fmt.Sprintf("https://%s/routing/service/rest/v3/%s", host, command)
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, &buf)
	req.SetBasicAuth(c.userID, c.password)
	fmt.Println(c.userID, c.password)
	resp, err := httpClient.Do(req)
	fmt.Println(resp)
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	var routingResp RoutingResponse
	//	s, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("moo", string(s))
	err = xml.NewDecoder(resp.Body).Decode(&routingResp)
	if err != nil {
		return nil, err
	}
	err = valid(routingResp)
	if err != nil {
		return nil, err
	}

	if hasErrors, errors := routingResp.HasErrors(); hasErrors {
		return nil, fmt.Errorf("Hermes Distribution Interface error: %v:%v", errors[0].ErrorCode, errors[0].ErrorDescription)
	}
	if hasWarnings, warnings := routingResp.HasWarnings(); hasWarnings {
		for _, w := range warnings {
			fmt.Println("WARNING: ", w)
		}
	}
	return &routingResp, nil
}
