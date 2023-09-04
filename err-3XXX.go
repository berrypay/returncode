/*
 * Project: Application Return Code Registry
 * Filename: /err-3XXX.go
 * Created Date: Tuesday September 5th 2023 04:55:59 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Tuesday September 5th 2023 05:00:07 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package returncode

// Internal Server Error: DB
const (
	// Repository client (PostgreSQL) error
	ERR3000 AppReturnCode = iota + 3000
	// DB operation error
	ERR3001
)

const (
	// Repository client (MySQL) error
	ERR3100 AppReturnCode = iota + 3100
	// DB operation error
	ERR3101
)
