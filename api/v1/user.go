package v1

//func UserRegister(c *gin.Context) {
//	var userRegister service.UserService
//	if err := c.ShouldBind(&userRegister); err == nil {
//		res := userRegister.Register(c.Request.Context())
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"code":    "1",
//			"message": "error " + err.Error(),
//		})
//	}
//}
//
//func UserLogin(c *gin.Context) {
//	var userLogin service.UserService
//	if err := c.BindJSON(&userLogin); err == nil {
//		res := userLogin.Login(c.Request.Context())
//		c.JSON(http.StatusOK, res)
//	} else {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"code":    "1",
//			"message": "error " + err.Error(),
//		})
//	}
//
//}
