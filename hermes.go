package hermes

import (
	"bytes"
	"encoding/xml"
	"fmt"
	//"io/ioutil"
	"net/http"
)

// AddEntry - Add a DeliveryRoutingRequestEntry.
func (r *DeliveryRoutingRequest) AddEntry(entry DeliveryRoutingRequestEntry) {
	r.DeliveryRoutingRequestEntries = append(r.DeliveryRoutingRequestEntries, entry)
}

// Valid - Check that a request has all mandatory fields filled in.
func (r *DeliveryRoutingRequest) Valid() error {
	return valid(r)
}

// HasWarnings - returns true as first arg if warning are present, and the warning messages as the second.
func (r *RoutingResponse) HasWarnings() (bool, []Message) {
	var warnings []Message
	for _, re := range r.RoutingResponseEntries {
		for _, w := range re.WarningMessages {
			warnings = append(warnings, w)
		}
	}
	return len(warnings) > 0, warnings
}

// HasErrors - returns true as first arg if errors are present, and the error messages as the second.
func (r *RoutingResponse) HasErrors() (bool, []Message) {
	var errors []Message
	for _, re := range r.RoutingResponseEntries {
		for _, e := range re.ErrorMessages {
			errors = append(errors, e)
		}
	}
	return len(errors) > 0, errors
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
	req.SetBasicAuth("USER", "PASSWORD")
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

	hasErrors, errors := routingResp.HasErrors()
	if hasErrors {
		return nil, fmt.Errorf("Hermes Distribution Interface error: %v:%v", errors[0].ErrorCode, errors[0].ErrorDescription)
	}
	return &routingResp, nil
}
