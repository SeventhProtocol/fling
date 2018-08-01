package actions

import "github.com/gobuffalo/buffalo"

/* RequestsHandler is responsible for serving the requests frontend*/

func RequestsHandler(c buffalo.Context) error {
	c.Set("title", "Requests")
	return c.Render(200, r.HTML("requests.html"))
}
