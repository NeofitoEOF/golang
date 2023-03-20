package seguranca

import "golang.org/x/crypto/bcrypt"

// Gerar hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// Verificar senha com hash, e verifica se são iguais
func VerificarSenha(senhaComHash, senhasString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhasString))
}
