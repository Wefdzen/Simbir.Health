package middleware

import "github.com/gin-gonic/gin"

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if http.Answear != StatusOk abort return
		c.Next()
	}
}

//will be after middleware Authe
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// незнаю либо просто брать из accessToken либо брать(вроде сверка с бд ролей не надо)
		c.Next()
	}
}
