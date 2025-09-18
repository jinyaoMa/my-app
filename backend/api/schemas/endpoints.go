package schemas

type EndpointsItem struct {
	Method      string   `json:"method" doc:"Method"`
	Path        string   `json:"path" doc:"Path"`
	Summary     string   `json:"summary" doc:"Summary"`
	OperationID string   `json:"operationId" doc:"Operation ID"`
	Tags        []string `json:"tags" doc:"Tags"`
}
