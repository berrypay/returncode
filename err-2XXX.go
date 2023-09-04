/*
 * Project: Application Return Code Registry
 * Filename: /err-2XXX.go
 * Created Date: Tuesday September 5th 2023 04:55:59 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Tuesday September 5th 2023 05:10:43 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package returncode

// Internal Server Error
const (
	// General server error
	ERR2000 int = iota + 2000
	// General Data Processing Error
	ERR2001
)

// Internal Server Error: Web Server Engine
const (
	// HTTP Server Error
	ERR2100 int = iota + 2100
	// Server Connectivity Error
	ERR2101
)

// Internal Server Error: MQ Server Engine
const (
	// MQ Server Error
	ERR2200 int = iota + 2200
	// MQ Connectivity Error
	ERR2201
)
