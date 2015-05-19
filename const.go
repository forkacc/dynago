package dynago

type CapacityDetail string

const (
	CapacityIndexes CapacityDetail = "INDEXES"
	CapacityTotal   CapacityDetail = "TOTAL"
	CapacityNone    CapacityDetail = "NONE"
)

type Select string

const (
	SelectAllAttributes      Select = "ALL_ATTRIBUTES"
	SelectAllProjected       Select = "ALL_PROJECTED_ATTRIBUTES"
	SelectCount              Select = "COUNT"
	SelectSpecificAttributes Select = "SPECIFIC_ATTRIBUTES"
)

type ReturnValues string

const (
	ReturnNone       ReturnValues = "NONE"
	ReturnAllOld     ReturnValues = "ALL_OLD"
	ReturnUpdatedOld ReturnValues = "UPDATED_OLD"
	ReturnAllNew     ReturnValues = "ALL_NEW"
	ReturnUpdatedNew ReturnValues = "UPDATED_NEW"
)

const dynamoTargetPrefix = "DynamoDB_20120810." // This is the Dynamo API version we support
