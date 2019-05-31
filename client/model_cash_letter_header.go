/*
 * X9 API
 *
 * Moov X9 () implements an HTTP API for creating, parsing and validating X9 (Check21) files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

type CashLetterHeader struct {
	// CashLetterHeader ID
	Id string `json:"id,omitempty"`
	// CollectionTypeIndicator is a code that identifies the type of cash letter.  * `00` - Preliminary Forward Information * `01` - Forward Presentment * `02` - Forward Presentment - Same-Day Settlement * `03` - Return * `04` - Return Notification * `05` - Preliminary Return Notification * `06` - Final Return Notification * `20` - No Detail * `99` - Bundles not the same collection type. Use of the value is only allowed by clearing arrangement. 
	CollectionTypeIndicator string `json:"collectionTypeIndicator,omitempty"`
	// DestinationRoutingNumber is the routing and transit number of the institution that receives and processes the cash letter or the bundle.
	DestinationRoutingNumber string `json:"destinationRoutingNumber,omitempty"`
	// ECEInstitutionRoutingNumber is the routing and transit number of the institution that creates the Cash Letter Header record
	ECEInstitutionRoutingNumber string `json:"eCEInstitutionRoutingNumber,omitempty"`
	// cashLetterBusinessDate is the business date of the cash letter. (Format YYYYMMDD, where - YYYY year, MM month, DD day)
	CashLetterBusinessDate string `json:"cashLetterBusinessDate,omitempty"`
	// cashLetterCreationDate is the date that the cash letter is created (Format YYYYMMDD, where - YYYY year, MM month, DD day)
	CashLetterCreationDate time.Time `json:"cashLetterCreationDate,omitempty"`
	// CashLetterCreationTime is the time that the cash letter is created. (Format - hhmm, where - hh hour, mm minute)
	CashLetterCreationTime string `json:"cashLetterCreationTime,omitempty"`
	// RecordTypeIndicator is a code that indicates the presence of records or the type of records contained in the cash letter. If an image is associated with any CheckDetail or Return, the cash letter must have a RecordTypeIndicator of I or F.  * `N` - No electronic check records or image records (Type 2x’s, 3x’s, 5x’s); e.g., an empty cash letter. * `E` - Cash letter contains electronic check records with no images (Type 2x’s and 3x’s only). * `I` - Cash letter contains electronic check records (Type 2x’s, 3x’s) and image records (Type 5x’s). * `F` - Cash letter contains electronic check records (Type 2x’s and 3x’s) and image records (Type 5x’s) that correspond to a previously sent cash letter (i.e., E file). 
	RecordTypeIndicator string `json:"recordTypeIndicator,omitempty"`
	// DocumentationTypeIndicator is a code that indicates the type of documentation that supports all check records in the cash letter.  * `A` - No image provided, paper provided separately * `B` - No image provided, paper provided separately, image upon request * `C` - Image provided separately, no paper provided * `D` - Image provided separately, no paper provided, image upon request * `E` - Image and paper provided separately * `F` - Image and paper provided separately, image upon request * `G` - Image included, no paper provided * `H` - Image included, no paper provided, image upon request * `I` - Image included, paper provided separately * `J` - Image included, paper provided separately, image upon request * `K` - No image provided, no paper provided * `L` - No image provided, no paper provided, image upon request * `M` - No image provided, Electronic Check provided separately * `Z` - Not Same Type–Documentation associated with each item in Cash Letter will be different. The Check Detail 
	DocumentationTypeIndicator string `json:"documentationTypeIndicator,omitempty"`
	// CashLetterID uniquely identifies the cash letter. It is assigned by the institution that creates the cash letter and must be unique within a Cash Letter Business Date.
	CashLetterID string `json:"cashLetterID,omitempty"`
	// OriginatorContactName is the name of contact at the institution that creates the cash letter.
	OriginatorContactName string `json:"originatorContactName,omitempty"`
	// OriginatorContactPhoneNumber is the phone number of the contact at the institution that creates the cash letter.
	OriginatorContactPhoneNumber string `json:"originatorContactPhoneNumber,omitempty"`
	// fedWorkType is any valid codes specified by the Federal Reserve Bank
	FedWorkType string `json:"fedWorkType,omitempty"`
	// ReturnsIndicator ddentifies type of returns  * ` ` - Original Message * `E` - Administrative - items being returned that are handled by the bank and usually do not directly affect the customer or its account. * `R` - Customer–items being returned that directly affect a customer’s account. * `J` - Reject Return 
	ReturnsIndicator string `json:"returnsIndicator,omitempty"`
	// UserField is a field used at the discretion of users of the standard
	UserField string `json:"userField,omitempty"`
}
