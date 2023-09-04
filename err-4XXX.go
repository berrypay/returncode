/*
 * Project: Application Return Code Registry
 * Filename: /err-4XXX.go
 * Created Date: Tuesday September 5th 2023 04:55:59 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Tuesday September 5th 2023 04:56:53 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package returncode

// Provider Service Error
const (
	// General partner service operation Error
	ERR4000 AppReturnCode = iota + 4000
	// API invalid credential
	ERR4001
	// API call timeout
	ERR4002
	// API call failure
	ERR4003
	// API call response payload processing error
	ERR4004
	// API call response payload signature verification error
	ERR4005
)
