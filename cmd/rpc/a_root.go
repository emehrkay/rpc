package rpc

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/emehrkay/rpc/service"
	"github.com/emehrkay/rpc/storage"
)

var (
	err        error
	rpcService *service.Service
	store      storage.Storage
	logger     *slog.Logger
)

func init() {
	store = storage.NewMemory()
	logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
	rpcService, err = service.New(store, logger)
	if err != nil {
		msg := fmt.Sprintf(`unable to create rpc service -- %v`, err)
		panic(msg)
	}
}

var RootCmd = &cobra.Command{
	Use:   "rpc",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
