# pixie

TFTP and HTTP server specifically designed to serve [iPXE](https://ipxe.org).

pixie comes embedded with the following ROMs provided by the iPXE project:

- `undionly.kpxe` - for legacy (BIOS) machines
- `snponly64.efi` - for 64-bit EFI machines
- `snponly32.efi` - for 32-bit EFI machines
- `ipxe64.efi` - for 64-bit EFI machines (contains embedded networking stack)
- `ipxe32.efi` - for 32-bit EFI machines (contains embedded networking stack)
- `snponly-arm64.efi` - for 64-bit ARM machines
- `snponly-arm32.efi` - for 32-bit ARM machines

It comes with an embedded [iPXE script](https://ipxe.org/scripting), which is used to chainload into per-machine iPXE scripts or access an iPXE shell.

![iPXE menu](https://user-images.githubusercontent.com/26496/136672695-f7598db1-d067-4ac7-93d3-6a640726d471.png)

It has no dependencies, everything is in one binary.

## Usage

1. Run the `pixie` binary using default ports, specifying path to store iPXE scripts:

```
pixie -d /srv/ipxe
```

> See `--help` for more options

2. Define iPXE script for a machine with the MAC-address `00:11:22:33:44:55`:

```
cat >/srv/ipxe/00-11-22-33-44-55.ipxe <<EOF
#!ipxe
initrd http://dl.rockylinux.org/pub/rocky/8/BaseOS/x86_64/os/images/pxeboot/initrd.img
kernel http://dl.rockylinux.org/pub/rocky/8/BaseOS/x86_64/os/images/pxeboot/vmlinuz initrd=initrd.magic
boot
EOF
```

3. Configure pixie as the next-server on the DHCP server:

Example for [ISC dhcpd](https://www.isc.org/dhcp/):

```
option client-architecture code 93 = unsiged integer 16;
if exists user-class and option user-class = "iPXE" {
	filename "chain.ipxe";
} elsif option client-architecture = 00:00 {
	filename "undionly.kpxe";
} else {
	filename "snponly.efi";
}
next-server 192.168.0.100;
```

Example for [dnsmasq](https://thekelleys.org.uk/dnsmasq/doc.html):

```
dhcp-match=set:ipxe,175
dhcp-vendorclass=BIOS,PXEClient:Arch:00000
dhcp-boot=tag:!ipxe,tag:BIOS,undionly.kpxe,pixie,192.168.0.100
dhcp-boot=tag:!ipxe,tag:!BIOS,snponly.efi,pixie,192.168.0.100
dhcp-boot=tag:ipxe,chain.ipxe,pixie,192.168.0.100
```

> In both these examples `pixie` runs on `192.168.0.100`

### Walkthrough

When a machine with the MAC-address of `00:11:22:33:44:55` now attempts to PXE boot, the following will happen:

1. The machine queries DHCP (standard procedure from PXE)
2. The DHCP server will instruct the machine to download a new ROM from pixie (`next-server`) using TFTP
	- if the machine is running BIOS it will ask for `undionly.kpxe`
	- if the machine is running EFI it will ask for `snponly.efi`
3. The machine downloads and runs the iPXE ROM from pixie
4. The machine queries DHCP again (this time from iPXE)
5. The DHCP server will instruct the machine to download the `chain.ipxe` file from pixie using TFTP
6. The machine downloads and runs the `chain.ipxe` script
	- the screenshot above shows this stage
7. Unless manually intervened, the machine chainloads into its iPXE script from pixie using HTTP
8. The machine boots the operating system
