package rpc

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/emehrkay/rpc/web"
	"github.com/spf13/cobra"
)

func init() {
	var (
		port string
	)
	var startServer = &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			router := http.NewServeMux()
			if !strings.HasPrefix(port, ":") {
				port = ":" + port
			}

			server, err := web.New(port, rpcService, router)
			if err != nil {
				panic(fmt.Sprintf(`unable to start web server: %v`, err))
			}

			server.Run()
		},
	}

	startServer.PersistentFlags().StringVar(&port, "port", ":8090", "set the port with the flag --port=:8090")

	RootCmd.AddCommand(startServer)
}
