package legacy

import (
	"github.com/noxworld-dev/opennox/v1/client"
	"github.com/noxworld-dev/opennox/v1/client/noxrender"
)

func nox_thing_base_draw(vp *noxrender.Viewport, dr *client.Drawable) int {
	nox_thing_weapon_draw(vp, dr)
	return 1
}
