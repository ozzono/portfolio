package rest

// MapPortRoutes
func (h *portHandlers) MapPortRoutes() {
	h.group.GET("/port/{id}", h.get())
	h.group.GET("/ports/", h.query())
	h.group.PUT("/port", h.upsert())
	h.group.DELETE("/port/{id}", h.del())
	h.group.GET("/parse-json", h.parseJson())
}
