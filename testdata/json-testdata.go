package testdata

// JSONObj ...
type JSONObj struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var (
	// JSONDataStruct ...
	JSONDataStruct = JSONObj{
		Status:  "OK",
		Message: "Hello world",
	}
	// JSONDataMap ...
	JSONDataMap = map[string]string{
		"status":  "OK",
		"message": "Hello world",
	}
	// JSONStr ...
	JSONStr = `{
		"status": "OK",
		"message": "Hello world"
	}`
)
