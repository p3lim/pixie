package tftp

import _ "embed"

// this will embed iPXE ROMs already present in the `ipxe` subdirectory relative to this file into
// variables during build time.

//go:embed ipxe/bin/undionly.kpxe
var undionly []byte

//go:embed ipxe/bin-x86_64-efi/ipxe.efi
var ipxe64 []byte

//go:embed ipxe/bin-i386-efi/ipxe.efi
var ipxe32 []byte

//go:embed ipxe/bin-x86_64-efi/snponly.efi
var snponly64 []byte

//go:embed ipxe/bin-i386-efi/snponly.efi
var snponly32 []byte

//go:embed ipxe/bin-arm64-efi/snponly.efi
var snponly64arm []byte

//go:embed ipxe/bin-arm32-efi/snponly.efi
var snponly32arm []byte

//go:embed chain.ipxe
var chainTemplate string
