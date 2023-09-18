/*
 * Project: Application Return Code Registry
 * Filename: /err-1XXX.go
 * Created Date: Tuesday September 5th 2023 04:55:59 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Tuesday September 5th 2023 05:10:36 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package returncode

// Bad Request
const (
	// General Error
	ERR1000 int = iota + 1000
	// Header: Trace ID Missing
	ERR1001
	// Header: Partner ID Missing
	ERR1002
	// Header: System ID Missing
	ERR1003
	// Header: User ID Missing
	ERR1004
	// Payload: Binding/Processing Failure
	ERR1005
)

// Security Error
const (
	// General Request Security Security Error
	ERR1100 int = iota + 1100
	// Signature Missing
	ERR1101
	// Signature Not Matched
	ERR1102
	// Requesting Username Security Error
	ERR1103
	// Signing Key Fetching Error
	ERR1104
	// CRC Check Error
	ERR1105
)

// Input Validation Error
const (
	// General Input Validation Error
	ERR1200 int = iota + 1200
	// Invalid input format
	ERR1201
	// Input validation error
	ERR1202
	// Missing input
	ERR1203
	// Input not allowed
	ERR1204
	// Input duplicates
	ERR1205
)

// Response Error
const (
	// General Response Error
	ERR1300 int = iota + 1300
	// Record not found
	ERR1301
	// Record already exists
	ERR1302
	// Record out of range
	ERR1303
	// Record incomplete
	ERR1304
)
