package deploy

import (
	"synck8s/pkgs/dstk8s"
	"synck8s/pkgs/srck8s"
)

type Deployment struct {
	Deployment   string
	SrcNamespace string
	DstNamespace string
	image        string
}

func (d *Deployment) getimage() {
	d.image = srck8s.GetImages(d.Deployment, d.SrcNamespace)

}
func (d *Deployment) Sync() {
	d.getimage()
	dstk8s.SetImages(d.Deployment, d.DstNamespace, d.image)

}
