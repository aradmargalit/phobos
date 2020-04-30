package transport

func BuildRouter(svc service.PhobosAPI) gin.Engine {
	someEndpoint := makeSomeEndpoint(svc)
} 