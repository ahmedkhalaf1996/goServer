/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Pricing
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// PricingV2VoiceNumber struct for PricingV2VoiceNumber
type PricingV2VoiceNumber struct {
	// The destination phone number, in E.164 format
	DestinationNumber *string `json:"destination_number,omitempty"`
	// The origination phone number, in E.164 format
	OriginationNumber *string `json:"origination_number,omitempty"`
	// The name of the country
	Country *string `json:"country,omitempty"`
	// The ISO country code
	IsoCountry *string `json:"iso_country,omitempty"`
	// The list of OutboundCallPriceWithOrigin records
	OutboundCallPrices *[]PricingV2VoiceVoiceNumberOutboundCallPrices `json:"outbound_call_prices,omitempty"`
	InboundCallPrice   *PricingV2VoiceVoiceNumberInboundCallPrice     `json:"inbound_call_price,omitempty"`
	// The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy)
	PriceUnit *string `json:"price_unit,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}