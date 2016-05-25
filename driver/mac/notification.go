package mac

import "github.com/maxence-charriere/iu"

func init() {
	iu.SetBadgeHandler = setAppBadge
}
