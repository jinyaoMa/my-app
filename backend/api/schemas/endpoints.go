package schemas

type EndpointsItem struct {
	Method      string   `json:"method" doc:"Method"`
	Path        string   `json:"path" doc:"Path"`
	Summary     string   `json:"summary" doc:"Summary"`
	OperationID string   `json:"operationId" doc:"Operation ID"`
	Tags        []string `json:"tags" doc:"Tags"`
	Required    bool     `json:"required" doc:"Security Scope Required"`
}

func (e *EndpointsItem) Security(s []map[string][]string) {
	for _, schemeMap := range s {
    for _, scopes := range schemeMap {
      if len(scopes) > 0 {
        e.Required = true
        return
      }
    }
	}
}
