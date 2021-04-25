/*
 * ImageCashLetter API
 *
 * Moov Image Cash Letter (ICL) implements an HTTP API for creating, parsing and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CheckDetailAddendumB struct for CheckDetailAddendumB
type CheckDetailAddendumB struct {
	// CheckDetailAddendumB ID
	ID string `json:"ID,omitempty"`
	// ImageReferenceKeyIndicator identifies whether ImageReferenceKeyLength contains a variable value within the allowable range, or contains a defined value and the content is ItemReferenceKey.  * `0` - ImageReferenceKeyIndicator has Defined Value of 0034 and ImageReferenceKey contains the Image Reference Key. * `1`- ImageReferenceKeyIndicator contains a value other than Value 0034; or ImageReferenceKeyIndicator contains Value 0034, which is not a Defined Value, and the content of ImageReferenceKey has no special significance with regards to an Image Reference Key; or ImageReferenceKeyIndicator is 0000, meaning the ImageReferenceKey is not present.
	ImageReferenceKeyIndicator int32 `json:"imageReferenceKeyIndicator,omitempty"`
	// microfilmArchiveSequenceNumber is a number that identifies the item in the microfilm archive system; it may be different than the Check.ECEInstitutionItemSequenceNumber and from the ImageReferenceKey.
	MicrofilmArchiveSequenceNumber string `json:"microfilmArchiveSequenceNumber"`
	// MicrofilmArchiveSequenceNumber A number that identifies the item in the microfilm archive system; it may be different than the Check.ECEInstitutionItemSequenceNumber and from the ImageReferenceKey.  * `0034` - ImageReferenceKey contains the ImageReferenceKey (ImageReferenceKeyIndicator is 0). * `0000` - ImageReferenceKey not present (ImageReferenceKeyIndicator is 1). * `0001` - 9999: May include Value 0034, and ImageReferenceKey has no special significance to Image Reference Key (ImageReferenceKey is 1).
	LengthImageReferenceKey string `json:"lengthImageReferenceKey,omitempty"`
	// ImageReferenceKey is used to find the image of the item in the image data system.  Size is variable based on lengthImageReferenceKey. The position within the file is variable based on the lengthImageReferenceKey.
	ImageReferenceKey string `json:"imageReferenceKey,omitempty"`
	// Descript describes the transaction.  The position within the file is variable based on the lengthImageReferenceKey.
	Descript string `json:"descript,omitempty"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField,omitempty"`
}
