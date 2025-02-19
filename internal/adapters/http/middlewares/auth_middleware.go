package middlewares

/*func AuthMiddleware(ss *services.SessionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionID := c.Cookies("session_id")
		if sessionID == "" {
			return c.Redirect("/login", http.StatusFound)
		}

		userID, err := ss.ValidateSession(c.Context(), sessionID)
		if err != nil {
			return c.Redirect("/login", http.StatusFound)
		}

		// Stocker l'userID dans le contexte
		c.Locals("userID", userID)

		// Passer au handler suivant
		return c.Next()
	}
}*/
