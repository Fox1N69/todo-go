package controller

type AuthController struct{}

var secretKey = []byte("secret")

func hashPassword(password string) (string, error) {
	return "", nil
}

func (c *AuthController) Login() error {
	return nil
}

func (c *AuthController) Register() error {
	return nil
}
