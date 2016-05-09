package iu

import (
	"strconv"

	"github.com/maxence-charriere/iu-log"
)

// DriverToken is an identifier for a driver.
type DriverToken string

// ComponentToken is an unique identifier for a component.
type ComponentToken uint

// ComponentTokenFromString converts a string to a ComponentToken.
func ComponentTokenFromString(s string) ComponentToken {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		iulog.Panic(err)
	}

	return ComponentToken(id)
}

func nextComponentToken() ComponentToken {
	componentMtx.Lock()
	defer componentMtx.Unlock()

	currentComponentID++
	return currentComponentID
}
