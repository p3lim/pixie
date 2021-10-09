package tftp

import _ "embed"

// we cannot embed from parent directories, so we'll have to copy the files closer to embed them
// https://github.com/golang/go/issues/46056

//go:generate cp ../../ipxe/src/bin/undionly.kpxe ./
//go:embed undionly.kpxe
var undionly []byte

//go:generate cp ../../ipxe/src/bin-x86_64-efi/ipxe.efi ./
//go:embed ipxe.efi
var ipxe64 []byte

//go:generate cp ../../ipxe/src/bin-i386-efi/ipxe.efi ./ipxe32.efi
//go:embed ipxe32.efi
var ipxe32 []byte

//go:embed chain.ipxe
var chainTemplate string
