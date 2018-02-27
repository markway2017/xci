// The XCI is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

package ca

import (
	"crypto/elliptic"
)

var (
	defaultCurve elliptic.Curve
)

// GetDefaultCurve returns the default elliptic curve used by the crypto layer
func GetDefaultCurve() elliptic.Curve {
	return defaultCurve
}
