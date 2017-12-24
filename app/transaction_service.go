package app

import m "github.com/agustin-sarasua/rs-model"

var RentFlowMap = map[string]func(tx *m.Transaction) (string, error){
	"START": StartTransaction,
	"TENANT_DOCUMENTATION_UPLOAD": ValidateTenantDocumentation,
	"DOCUMENTATION_VALIDATION":    DocumentationValidation,
	"SIGN_CONTRACT":               SignContract}

func StartTransaction(tx *m.Transaction) (string, error) { return "OK", nil }

func ValidateTenantDocumentation(tx *m.Transaction) (string, error) { return "OK", nil }

func DocumentationValidation(tx *m.Transaction) (string, error) { return "OK", nil }

func SignContract(tx *m.Transaction) (string, error) { return "OK", nil }

func LoadTransaction(id int64) *m.Transaction {
	return nil
}
