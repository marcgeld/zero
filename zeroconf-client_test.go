package zero

import (
	"testing"

	"github.com/grandcat/zeroconf"
	"gotest.tools/assert"
)

// Test for zeroconf Discovery
func TestDiscovery(t *testing.T) {

	t.Log("Starting zeroconf test server")

	zeroInstance := "zerocfg"
	zeroService := "_workstation._tcp"
	zeroDomain := "local."

	server, err := zeroconf.Register(zeroInstance, zeroService, zeroDomain, 42424, []string{"txt=test"}, nil)
	if err != nil {
		t.Error("Expected Register ok, got error: ", err)
	}
	defer server.Shutdown()

	var gw *zeroconf.ServiceEntry
	gw, err = Discovery(zeroInstance, zeroService, zeroDomain)
	if err != nil {
		t.Error("Expected Discovery ok, got error: ", err)
	}

	assert.Equal(t, gw.Instance, zeroInstance)
	assert.Equal(t, gw.Service, zeroService)
	assert.Equal(t, gw.Domain, zeroDomain)

	t.Log("Shutting down zeroconf test server")
}
