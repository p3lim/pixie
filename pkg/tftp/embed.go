//go:generate make -j8 -C ../../ipxe/src bin/undionly.kpxe bin-x86_64-efi/ipxe.efi bin-i386-efi/ipxe.efi bin-x86_64-efi/snponly.efi bin-i386-efi/snponly.efi
//go:generate make -j8 -C ../../ipxe/src CROSS=arm-none-eabi- ARCH=arm32 bin-arm32-efi/snponly.efi
//go:generate make -j8 -C ../../ipxe/src CROSS=aarch64-linux-gnu- ARCH=arm64 bin-arm64-efi/snponly.efi
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

//go:generate cp ../../ipxe/src/bin-x86_64-efi/snponly.efi ./snponly-x64.efi
//go:embed snponly-x64.efi
var snponly64 []byte

//go:generate cp ../../ipxe/src/bin-i386-efi/snponly.efi ./snponly-x86.efi
//go:embed snponly-x86.efi
var snponly32 []byte

//go:generate cp ../../ipxe/src/bin-arm64-efi/snponly.efi ./snponly-arm64.efi
//go:embed snponly-arm64.efi
var snponly64arm []byte

//go:generate cp ../../ipxe/src/bin-arm32-efi/snponly.efi ./snponly-arm32.efi
//go:embed snponly-arm32.efi
var snponly32arm []byte

//go:embed chain.ipxe
var chainTemplate string
