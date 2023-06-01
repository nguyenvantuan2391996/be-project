package constant

const (
	ParseRequestFail             = "Parse request fail"
	CreateNewUserFail            = "Create new user fail"
	CreateNewStandardFail        = "Create new standard fail"
	DeleteStandardFail           = "Delete standard fail"
	GetStandardsFail             = "Get list standard fail"
	BulkCreateNewScoreRatingFail = "Bulk create new score rating fail"
	DeleteScoreRatingFail        = "Delete score rating fail"
	GetListScoreRatingFail       = "Get list score rating fail"
	UpdateScoreRatingFail        = "Update score rating fail"
	ConsultResultFail            = "Consult result fail"
)

var MapErrorValidateRequest = map[string]string{
	"Name":         "User name is invalid (suggest: 3-20 characters)",
	"StandardName": "Standard name is invalid (suggest: 3-30 characters)",
	"Weight":       "Weight is required",
	"UserID":       "User id is required",
	"Metadata":     "Metadata is required",
}
