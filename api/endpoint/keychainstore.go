package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/prometheus/common/log"
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/services/keychainstore"
)

func GetKeychainStoreEvents(c *gin.Context, instance keychainstore.Instance) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"build":  internal.Build,
		"date":   internal.Date,
	})
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func KeychainStoreWSHandler(c *gin.Context, instance keychainstore.Instance) {
	w := c.Writer
	r := c.Request

	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn("Failed to set websocket upgrade: ", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			log.Warn("Failed to read WS message: ", err)
			break
		}
		err = conn.WriteMessage(t, msg)
		if err != nil {
			log.Warn("Failed to write WS message: ", err)
			break
		}
	}
}
