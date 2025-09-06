// keys package contains the public key for verifying digital signature of templates
package keys

import _ "embed"

const PDVerifier = "g1i1u/nuclei-edited-templates"

//go:embed nuclei.crt
var NucleiCert []byte // public key for verifying digital signature of templates
