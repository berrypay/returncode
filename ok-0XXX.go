/*
 * Project: Application Return Code Registry
 * Filename: /ok-0XXX.go
 * Created Date: Tuesday September 5th 2023 04:55:59 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Tuesday September 5th 2023 04:57:10 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package returncode

// Generic OK Status
const (
	// success OK
	OK0000 AppReturnCode = iota
	// OK, but with warning
	OK0001
)
