package scratch

import (
	"fmt"
	"net"
	"net/mail"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
)

const (
	Redacted     = "redacted"
	InvalidEmail = "invalidEmail"

	Email string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
)

var (
	userRegexp    = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")
	hostRegexp    = regexp.MustCompile("^[^\\s]+\\.[^\\s]+$")
	userDotRegexp = regexp.MustCompile("(^[.]{1})|([.]{1}$)|([.]{2,})")
	rxEmail       = regexp.MustCompile(Email)
)

func isValidEmail(email string) bool {
	return govalidator.IsExistingEmail(email)
}
func isValidEmail2(email string) (*mail.Address, error) {
	return mail.ParseAddress(email)
}

func isValidEmail3(email string) (bool, error) {
	// Define a more comprehensive regular expression for validating email addresses
	var re = regexp.MustCompile(`^(?i:[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?i:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)])$`)
	return re.MatchString(email), nil
}

func CheckEmail() {
	emails := []string{
		"example@example.com",
		"user.name+tag+sorting@example.com",
		"user@example.co.uk",
		"user@localhost",
		"user@.com",
		"user@com",
		"user@.com.",
		"user@-example.com",
		"user@example..com",
		"evertonkopec@hotmail.com",
		"test@test.xyz",
	}

	for _, email := range emails {
		_, err := IsAValidEmail(email)
		if err != nil {
			fmt.Printf("%s is a invalid email\n", email)
		} else {
			fmt.Printf("%s is an valid email\n", email)
		}

		//if isValidEmail(email) {
		//	fmt.Printf("%s is a valid email\n", email)
		//} else {
		//	fmt.Printf("%s is an invalid email\n", email)
		//}
	}
}

func IsAValidEmail(email string) (bool, error) {
	redactedEmail := strings.Split(email, "@")
	if redactedEmail[0] == Redacted {
		return false, fmt.Errorf("%s", Redacted)
	}

	if len(email) < 6 || len(email) > 254 {
		return false, fmt.Errorf("%s", InvalidEmail)
	}
	at := strings.LastIndex(email, "@")
	if at <= 0 || at > len(email)-3 {
		return false, fmt.Errorf("%s", InvalidEmail)
	}
	user := email[:at]
	host := email[at+1:]
	if len(user) > 64 {
		return false, fmt.Errorf("%s", InvalidEmail)
	}
	switch host {
	case "localhost", "example.com", "test.xyz":
		return true, nil
	}
	if userDotRegexp.MatchString(user) || !userRegexp.MatchString(user) || !hostRegexp.MatchString(host) {
		return false, fmt.Errorf("%s", InvalidEmail)
	}
	if _, err := net.LookupMX(host); err != nil {
		if _, err := net.LookupIP(host); err != nil {
			return false, fmt.Errorf("%s", InvalidEmail)
		}
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, fmt.Errorf("%s", InvalidEmail)
	}

	if !rxEmail.MatchString(email) {
		return false, fmt.Errorf("%s", InvalidEmail)
	}

	return true, nil
}
