# Backend authentication in Go

## High level flow:
- Login: User sends email, password through a POST request
- Verify: Server hashes password and checks email & hash in database
- Sign: If credentials are valid, server generated a JWT signed with a secret key
- Authorize: Server sends JWT as response and for future requests, the client sends this token in the Authorization header

## Core external libraries
- golang.org/x/crypto/bcrypt: For hashing passwords
- github.com/golang-jwt/jwt: For creating and parsing tokens

```bash
// Password Hashing
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// JWT generation
var jwtKey = []byte("your_secret_key")

func GenerateJWT(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    })

    return token.SignedString(jwtKey)
}
```

## Authentication middleware
```bash
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 1. Grab the token from the "Authorization" header
        tokenString := r.Header.Get("Authorization")
        
        // 2. Parse and validate
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            w.WriteHeader(http.StatusUnauthorized)
            fmt.Fprintf(w, "Not Authorized")
            return
        }

        // 3. Proceed to the actual endpoint
        handler(w, r)
    }
}
```