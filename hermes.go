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
