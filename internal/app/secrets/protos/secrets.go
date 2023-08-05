package protos

type TransitEncryptRequest struct {
	Data string `json:"data" form:"data" binding:"required"`
}

type TransitEncryptResponse struct {
	Data string `json:"data"`
}

type TransitDecryptRequest struct {
	Data string `json:"data" form:"data" binding:"required"`
}

type TransitDecryptResponse struct {
	Data string `json:"data"`
}
