# Playground for Network Interfaces

Just messing around with network interfaces in an attempt to understand
some of the low-level OS fundamentals. I was led in this direction when
looking at Wireguard's packages and getting a fun,system-level example of
how to run system calls and file descriptors to configure network interfaces
in Go.

## Wireguard's TUN interface

The TUN interface is created at `/dev/net/tun` and represents the tunnel
used to communicate with the external service. When running through
Cloudflare's [Understanding TUN/TAP][], it made more sense that this TUN
device represented a local point-to-point connection with the Wireguard
service. Found this fun rhyme to remember the difference between the 
device types:

> Tap is like a switch,\
Ethernet headers it'll hitch.\
Tun is like a tunnel,\
VPN connections it'll funnel.\
Ethernet headers it won't hold,\
Tap uses, tun does not, we're told.

To see Wireguard create the TUN interface, a lot of the code in 
`FirstExample()` is pulled directly from the [Wireguard Mirror][]'s
`main.go` function. See the example with the following:

```bash
# In one terminal, run FirstExampe() to create the TUN interface:
sudo go run main.go

# Then in another terminal, view the TUN interface using iproute2:
# (Should output louis0: tun vnet_hdr)
ip tuntap list
```

## References

- [Point-to-Point topology][]
- [Linux TUN/TAP Networking][]
- [Wireguard Mirror][]
- [Understanding TUN/TAP][]

[Linux TUN/TAP Networking]: https://docs.kernel.org/networking/tuntap.html
[Wireguard Mirror]: https://github.com/tailscale/wireguard-go
[Point-to-Point topology]: https://lightyear.ai/blogs/point-to-point-leased-lines-p2p-vs-wavelength-circuits
[Understanding TUN/TAP]: https://blog.cloudflare.com/virtual-networking-101-understanding-tap/
