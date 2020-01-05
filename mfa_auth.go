package main

/*
#include <security/pam_appl.h>
*/
import "C"
import (
	"github.com/dgryski/dgoogauth"
	"strings"
)

func authenticateTOTP(pamh *C.pam_handle_t, totp_secret string) bool {
	totp_secret = strings.ToUpper(totp_secret)
	totp_config := dgoogauth.OTPConfig{
		Secret:     totp_secret,
		WindowSize: totpWindow,
		UTC:        true,
	}
	totp_code := strings.TrimSpace(requestPass(pamh, C.PAM_PROMPT_ECHO_OFF, "TOTP: "))
	ok, err := totp_config.Authenticate(totp_code)
	if err != nil {
		return false
	}
	return ok
}
