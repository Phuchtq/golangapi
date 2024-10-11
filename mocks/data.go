package mocks

const (
	Existed           string = "1"
	NotExisted        string = "-1"
	EmptyString       string = ""
	UpdatedName       string = "Update"
	Token             string = "abcxyz"
	PositiveStatus    bool   = true
	RawPositiveStatus string = "true"
	NegativeStatus    bool   = false
	RawNegativeStatus string = "false"
	SecurePassword    string = "SecurePASSWORD@1234"
	InvalidPassword   string = "abcd"
	ValidEmail        string = "emailMockSample@gmail.com"
)

var Roles = map[string]string{
	"Admin":    "R001",
	"Staff":    "R002",
	"Customer": "R003",
}

var SignUpModelKeyCases = []string{
	"Empty email",
	"Empty password",
	"Email registered",
	"Password not secure",
	"Staff Provides Admin",
	"Invalid role",
	"Valid",
	"Valid Provide",
}

func GetSignUpModelCases() map[string]int {
	res := make(map[string]int)
	//---------------------------------
	for i, v := range SignUpModelKeyCases {
		res[v] = i + 1
	}
	//---------------------------------
	return res
}

var AccountStates = []string{
	"Banned",
	"Staff locked",
	"Customer locked",
	"Self locked",
	"Not activate",
	"Reset password",
	"Active",
}
