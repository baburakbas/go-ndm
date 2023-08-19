package pkg

import (
	"context"
	"fmt"
	"strings"

	auth "github.com/de-bugsBunny/go-ndm/pkg/infrastructure/authentication"
	tconfig "github.com/de-bugsBunny/go-ndm/tool/config"
	telnet "github.com/de-bugsBunny/go-ndm/tool/telnet"
)

type NokiaSr7750 struct {
	telConn *telnet.Telnet
}

func NewNokiaSr7750(telConn *telnet.Telnet) NokiaSr7750 {
	return NokiaSr7750{
		telConn: telConn,
	}
}

func (n NokiaSr7750) Connect(ctx context.Context, asset auth.Asset) error {
	err := n.telConn.Connect(asset.Host, asset.Port)
	if err != nil {
		return err
	}
	return nil
}

func (n NokiaSr7750) Login(ctx context.Context, user auth.User) error {
	n.telConn.SenderTelnet("\n")
	n.telConn.SenderTelnet("\n")
	n.telConn.SenderTelnet("\n")
	n.telConn.SenderTelnet("logout\n")
	receivedData := n.telConn.ReaderTelnet("Login:")
	if tconfig.GetAppConfigInstance().Debug {
		fmt.Println(receivedData)
	}
	n.telConn.SenderTelnet("admin\n")
	receivedData = n.telConn.ReaderTelnet("Password:")
	if tconfig.GetAppConfigInstance().Debug {
		fmt.Println(receivedData)
	}
	n.telConn.SenderTelnet("admin\n")
	n.telConn.SenderTelnet("show router interface\n")

	return nil
}

func (n NokiaSr7750) ShowRouter(ctx context.Context) error {
	n.telConn.SenderTelnet("\n")
	n.telConn.SenderTelnet("show router interface\n")
	receivedData := n.telConn.ReaderTelnet("\r")
	if tconfig.GetAppConfigInstance().Debug {
		fmt.Println(receivedData)
	}
	return nil
}

func (n NokiaSr7750) Backup(ctx context.Context) (string, error) {
	n.telConn.SenderTelnet("\n")
	n.telConn.SenderTelnet("admin display-config\n")
	allReceivedData := ""
	for {
		receivedData := n.telConn.ReaderTelnet("\n")
		allReceivedData += receivedData
		if !strings.Contains(receivedData, "Finished") {
			n.telConn.SenderTelnet("\n")
		} else {
			break
		}
	}
	if tconfig.GetAppConfigInstance().Debug {
		fmt.Println(allReceivedData)
	}
	return allReceivedData, nil
}
