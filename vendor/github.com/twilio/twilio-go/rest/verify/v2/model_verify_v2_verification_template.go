/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// VerifyV2VerificationTemplate struct for VerifyV2VerificationTemplate
type VerifyV2VerificationTemplate struct {
	// A string that uniquely identifies this Template
	Sid *string `json:"sid,omitempty"`
	// Account Sid
	AccountSid *string `json:"account_sid,omitempty"`
	// A string to describe the verification template
	FriendlyName *string `json:"friendly_name,omitempty"`
	// A list of channels that support the Template
	Channels *[]string `json:"channels,omitempty"`
	// Object with the template translations.
	Translations *interface{} `json:"translations,omitempty"`
}