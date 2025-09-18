package schemas

type OperationIdEnumPairItem struct {
	OperationId string `json:"operationId" doc:"Operation ID"`
	Enum        int    `json:"enum" doc:"Enum Value"`
}
