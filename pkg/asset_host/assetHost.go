package asset_host

import "poodle/pkg/utils"

type AssetHost struct {
	ip          string
	domain      utils.StDomain
	openedPorts []int
}
