package domain

var (
	// OnOK = 200
	OnOK = map[string]string{"message": "Resource loaded successfully."}
	// OnAccepted = 202
	OnAccepted = map[string]string{"message": "The record has been successfully action"}
	// OnCreated when store = 201
	OnCreated = map[string]string{"message": "The record has been successfully created."}
	// OnNoContent = 204
	OnNoContent = map[string]string{"message": "No Content, record not found."}
)
