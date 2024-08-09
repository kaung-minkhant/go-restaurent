package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSigningAndVerify(t *testing.T) {
	sampleData := []byte("{refresh_token:sHh6AVagtYRvqtM5isWdkewrg,role:admin}")
	signature, err := SignData(sampleData, ToPointer[string]("../keys/privKey.pem"))
	fmt.Println("signature", signature)
	assert.Nil(t, err)
	err = VerifySign(sampleData, signature, ToPointer[string]("../keys/pubKey.pem"))
	assert.Nil(t, err)
}

func TestVerify(t *testing.T) {
	dataFromCookie := "{refresh_token:bKjhhUfkXT7d09TTxOrmDNsEa,role:admin}::aa9bd91901f29bccfb7e56570e5c0c0b5e2c3cba2d890d7b24a24a4f797605f2690162c4c52a38c4bca1d09f46c50bccefde2c6fb48f3293023140c5d67b68d9b33ac909b1371c8bc20505a60c83726ff9712ad76deb72cb1f08307d815604146a4336a91aaf0419226f4bfd3c048aa45448d508de521e8f4d03d5086b22f60300a749ff2a2fab6c9dce46e3bd703b45edc03e05e87d3b0b95dc8d05618bcc1e7a93c2e1b9a59bf5c3701b359e930587bea0ff7f672d223bc66e8ed95a73b818999feca2f08c34d8efd82a33dc606a850d97d78b810d9c0e7fd18ab71550b0bfd2954d3dabce3df5032a1ae63140088437bc18610ede8196d6773db8dc0a7b8ccb7bc055eb7854410be9d9c6d6ce98fcc9e56e4cf0163e5441f1a86d81fdba3dc439f977387022ec9b724a2122c51cb69640b740d9ac87a242e047346cb5343eb4f929d13e47ddd35f1839341756a5e28a514a7823712f78f7945c9c7721eb825d3face3107828a9907816786038dc0b746f745be2ff1b01601c7c0abcf8f28388bf46fa724fea34a75b817ef589eae36df68c63f048184c583806e9e2e19001c69aa10ef67d10b5b9c1294af6030ff58a9c445333f12c6ed741d545c19dd40cdfdbd5bec13f7a9ffd4dc6c06ad548097c2b6a929e95dc08ef9b5650102650e9b2b2c14cc3a66ef7eb9f795a7ceb7efb31cedddf21599bc730f59e29d3e519c4"
	cookieChunks := strings.Split(dataFromCookie, "::")
	var cookieData any
	err := json.Unmarshal([]byte(cookieChunks[0]), &cookieData)
	assert.Nil(t, err)
	sampleData := []byte(cookieChunks[0])
	signature := cookieChunks[1]
	err = VerifySign(sampleData, signature, ToPointer[string]("../keys/pubKey.pem"))
	assert.Nil(t, err)

}
