package game

import (
	"net/http"

	"github.com/Synaxis/battlefieldHeroesBackend/magma/tpl"
)

// GET /en/game/store?personaId=2&lvl=-1&eqp=8+10+56+69+70+112+113+251+416+417+418+420+979+981+2145+3002+5000+5001+5002+5003+5004+5005+5006+5007+5008 HTTP/1.1
// Host: 127.0.0.1
// Accept: application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5
// Accept-Charset: iso-8859-1,*,utf-8
// Accept-Encoding: gzip,deflate
// Accept-Language: en-us,en
// Connection: keep-alive
// Cookie: magma=topsecret
func (c *Controller) store(w http.ResponseWriter, r *http.Request) {
	c.rdr.RenderXML(w, r, tpl.XmlStore, nil)
}
