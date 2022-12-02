package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"token-supply-api/app"
	"token-supply-api/app/uncirculating"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "config file path")

	a, err := app.NewTokenSupply(configPath, &uncirculating.ExcludedAddresses{}, &uncirculating.VestedTokens{})
	if err != nil {
		log.Fatal(err)
	}
	if err := a.Start(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/total", totalSupplyHandler(a))
	http.HandleFunc("/circulating", circulatingSupplyHandler(a))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Starting Token Supply API on port %d\n", a.Config.Port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.Config.Port), nil))
	}()

	<-sigs
	a.Stop()
	log.Printf("Stopped Token Supply API")
}

func totalSupplyHandler(a *app.TokenSupply) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, a.TotalSupply().String())
	}
}

func circulatingSupplyHandler(a *app.TokenSupply) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = fmt.Fprintln(w, a.CirculatingSupply().String())
	}
}
