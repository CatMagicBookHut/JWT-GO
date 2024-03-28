# JWT-GO
Golang's JWT middleware.

You can use redis instead of `usertoken.go`, I use it save and manage token.

## How to use it?

1. Create your jwt secreat as a replacement for `utils.JwtKey`.
2. Use `GenerateJWT(user *UserClaims, t time.Duration)` create jwt, the "t" is the token duration time.
3. Use `ValidateJWT(tokenString string, userip string)` vertify token, you should input token and login IP.
4. `JwtToken()` is a middleware from gin. Sure, you can design your own middleware. It will get token from Header's Authorization and vertify it.
