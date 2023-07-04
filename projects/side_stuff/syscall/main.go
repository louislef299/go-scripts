package main

import (
	"github.com/louislef299/go-scripts/projects/side_stuff/syscall/nl"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

// To see how a single interface is gathered, view:
// https://github.com/vishvananda/netlink/blob/v1.1.0/link_linux.go#L1398

// Empty handle used by the netlink package methods
var pkgHandle = &Handle{}

func main() {
	link, err := netlink.LinkByName("enp0s20f0u7")
	if err != nil {
		panic(err)
	}

	base := link.Attrs()
	pkgHandle.ensureIndex(base)
	req := h.newNetlinkRequest(unix.RTM_NEWLINK, unix.NLM_F_ACK)

	msg := nl.NewIfInfomsg(unix.AF_UNSPEC)
	msg.Change = unix.IFF_UP
	msg.Index = int32(base.Index)
	req.AddData(msg)

	_, err := req.Execute(unix.NETLINK_ROUTE, 0)
	return err
}

func (h *Handle) ensureIndex(link *netlink.LinkAttrs) {
	if link != nil && link.Index == 0 {
		newlink, _ := h.LinkByName(link.Name)
		if newlink != nil {
			link.Index = newlink.Attrs().Index
		}
	}
}
