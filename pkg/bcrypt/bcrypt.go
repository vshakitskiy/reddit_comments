package bcrypt

import "golang.org/x/crypto/bcrypt"

func Hash(s string) string {
	hashed, _ := bcrypt.GenerateFromPassword(
		[]byte(s),
		bcrypt.DefaultCost,
	)
	return string(hashed)
}

func Compare(hashed string, normal string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(normal),
	)
}
