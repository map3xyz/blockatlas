package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getEvents(customer-apikey)
// apikey -> keychainID
// customer-apikey -> wallet
func GetKeychainEvents(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"events": []string{"keychain_address"},
	})
}

// getAddress(user, customer-apikey, network, asset, addressType (user, memo))
// -> { address, memo }
// apikey -> keychainID
// customer-apikey -> wallet
// network/asset -> address
// generate new memo
// call store api -> (watchAddress/userAddress, memo)
func GetKeychainAddress(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"address": "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
		"memo":    "UUID",
	})
}
