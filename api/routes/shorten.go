package routes

type request struct {
	URL
	CustomShort
	Expiry
}

type response struct {
	URL
	CustomShort
	Expiry
	XRateRemaining
	XRateLimitReset
}
