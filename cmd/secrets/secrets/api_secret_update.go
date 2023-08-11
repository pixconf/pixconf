package secrets

// func (s *Secrets) apiSecretUpdate(c *gin.Context) {
// 	var secURI SecretURI

// 	if err := c.ShouldBindUri(&secURI); err != nil {
// 		resp := xerror.ErrorSingle(http.StatusBadRequest, err.Error())
// 		c.JSON(resp.Code, resp)
// 		return
// 	}

// 	if !xid.IsValidSecretID(secURI.ID) {
// 		resp := xerror.ErrorSingle(http.StatusBadRequest, "wrong secret id format")
// 		c.JSON(resp.Code, resp)
// 		return
// 	}

// }
