package jwt

func AuthToken() (bool){
	// // Parse the token string
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// Check signing method
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	// Return the secret key
	// 	return []byte("your-secret-key"), nil
	// })

	// // Check if there's an error parsing the token
	// if err != nil {
	// 	// Handle error
	// }

	// // Check if the token is valid
	// if !token.Valid {
	// 	// Handle invalid token
	// }

	// // Get claims from the token
	// claims, ok := token.Claims.(jwt.MapClaims)
	// if !ok {
	// 	// Handle invalid claims
	// }

	// // Get username claim from the claims
	// username := claims["username"].(string)
	// fmt.Println("Username:", username)

}