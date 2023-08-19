package main

import (
	"context"
	"fmt"
	"time"

	auth "github.com/de-bugsBunny/go-ndm/pkg/infrastructure/authentication"
	device "github.com/de-bugsBunny/go-ndm/pkg/infrastructure/device"
	inter "github.com/de-bugsBunny/go-ndm/pkg/infrastructure/interface"
	store "github.com/de-bugsBunny/go-ndm/pkg/infrastructure/store"
	tconfig "github.com/de-bugsBunny/go-ndm/tool/config"
	telnet "github.com/de-bugsBunny/go-ndm/tool/telnet"
)

var appConfig tconfig.AppConfig
var assetConfig tconfig.AssetConfig
var userConfig tconfig.UserConfig

func init() {
	appConfig.ReadConfig()
	assetConfig.ReadConfig()
	userConfig.ReadConfig()
}

func main() {

	if tconfig.GetAppConfigInstance().Debug {
		fmt.Printf("%+v\n", assetConfig)
		fmt.Printf("%+v\n", userConfig)
	}
	var telConn telnet.Telnet
	sr7750 := device.NewNokiaSr7750(&telConn)
	user := auth.User{
		UserName: userConfig.User.UserName,
		Password: userConfig.User.Password,
	}
	asset := auth.Asset{
		Host: assetConfig.Asset.Host,
		Port: assetConfig.Asset.Port,
	}

	store := store.NewFileStorage(appConfig.App.FilePath, appConfig.App.FilePrefix+time.Now().Format("2006_01_02_15_04_05")+appConfig.App.FileEndfix)
	err := FullProc(context.Background(), asset, user, sr7750, store)
	if err != nil {
		panic(err)
	}
}

func FullProc(ctx context.Context, asset auth.Asset, user auth.User, device inter.Router, store inter.Storer) error {
	err := device.Connect(ctx, asset)
	if err != nil {
		return err
	}
	err = device.Login(ctx, user)
	if err != nil {
		return err
	}
	err = device.ShowRouter(ctx)
	if err != nil {
		return err
	}
	var backupData string
	backupData, err = device.Backup(ctx)
	if err != nil {
		return err
	}
	err = store.Store(backupData)
	if err != nil {
		return err
	}
	err = store.Clean(3)
	if err != nil {
		return err
	}
	return nil
}
