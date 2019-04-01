package zero

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/grandcat/zeroconf"
)

// Discovery func to find the gateway via zeroconfig (mDns)
func Discovery(instance string, service string, domain string) (*zeroconf.ServiceEntry, error) {

	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Channel to receive discovered service entries
	lookupChannel := make(chan *zeroconf.ServiceEntry)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = resolver.Lookup(ctx, instance, service, domain, lookupChannel)
	if err != nil {
		log.Fatalln("Failed to lookup device:", err.Error())
		return nil, err
	}

	var deviceEntry *zeroconf.ServiceEntry
	select {
	case deviceEntry = <-lookupChannel:
		// Stop lookup
		cancel()
	case <-ctx.Done():
		log.Println(ctx.Err())
	}

	if deviceEntry == nil {
		return nil, errors.New("Device not found")
	}

	//log.Println("Found service:", deviceEntry)
	return deviceEntry, nil
}
