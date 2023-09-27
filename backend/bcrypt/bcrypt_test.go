package bcrypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBcrypt(t *testing.T) {
	t.Run("비밀번호 암호화를 진행한다.", func(t *testing.T) {

		password := "pwd1234"
		generate, err := Generate(password)
		assert.Nil(t, err)
		assert.NotEmpty(t, generate)
		assert.NotEqual(t, password, generate)
	})
}
