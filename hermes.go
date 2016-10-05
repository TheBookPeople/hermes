package hermes

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

// Valid - Check that a response looks to be correct.
func (r *RoutingResponse) Valid() error {
	return valid(r)
}
