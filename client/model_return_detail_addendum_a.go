/*
 * ImageCashLetter API
 *
 * Moov Image Cash Letter (ICL) implements an HTTP API for creating, parsing and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// ReturnDetailAddendumA struct for ReturnDetailAddendumA
type ReturnDetailAddendumA struct {
	// ReturnDetailAddendumA ID
	ID string `json:"ID,omitempty"`
	// RecordNumber is a number representing the order in which each ReturnDetailAddendumA was created. ReturnDetailAddendumA shall be in sequential order starting with 1.
	RecordNumber int32 `json:"recordNumber"`
	// ReturnLocationRoutingNumber is a valid routing and transit number indicating where returns, final return notifications, and preliminary return notifications are sent, usually the BOFD.
	ReturnLocationRoutingNumber string `json:"returnLocationRoutingNumber"`
	// BOFDEndorsementDate is the date of endorsement
	BOFDEndorsementDate time.Time `json:"bOFDEndorsementDate,omitempty"`
	// BOFDItemSequenceNumber is a number that identifies the item in the CheckDetailAddendumA.
	BOFDItemSequenceNumber string `json:"bOFDItemSequenceNumber,omitempty"`
	// BOFDAccountNumber is a number that identifies the depository account at the Bank of First Deposit.
	BOFDAccountNumber string `json:"bOFDAccountNumber,omitempty"`
	// BOFDBranchCode is a code that identifies the branch at the Bank of First Deposit.
	BOFDBranchCode string `json:"bOFDBranchCode,omitempty"`
	// PayeeName is the name of the payee from the check.
	PayeeName string `json:"payeeName,omitempty"`
	// TruncationIndicator identifies if the institution truncated the original check item.
	TruncationIndicator string `json:"truncationIndicator"`
	// BOFDConversionIndicator is a code that indicates the conversion within the processing institution between original paper check, image and IRD. The indicator is specific to the action of institution that created this record.  * `0` - Did not convert physical document * `1` - Original paper converted to IRD * `2` - Original paper converted to image * `3` - IRD converted to another IRD * `4` - IRD converted to image of IRD * `5` - Image converted to an IRD * `6` - Image converted to another image (e.g., transcoded) * `7` - Did not convert image (e.g., same as source) * `8` - Undetermined
	BOFDConversionIndicator string `json:"bOFDConversionIndicator,omitempty"`
	// BOFDCorrectionIndicator identifies whether and how the MICR line of this item was repaired by the creator of this CheckDetailAddendumA Record for fields other than Payor Bank Routing Number and Amount. * `0` - No Repair * `1` - Repaired (form of repair unknown) * `2` - Repaired without Operator intervention * `3` - Repaired with Operator intervention * `4` - Undetermined if repair has been done or not
	BOFDCorrectionIndicator int32 `json:"BOFDCorrectionIndicator,omitempty"`
	// UserField identifies a field used at the discretion of users of the standard.
	UserField string `json:"userField,omitempty"`
}
