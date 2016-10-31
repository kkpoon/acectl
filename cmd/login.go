package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"

	"github.com/howeyc/gopass"
	aceproject "github.com/kkpoon/go-aceproject"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login aceproject",
	Long: `This command help you to login to aceproject and get the token for
other command to use.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Account ID: ")
		account, _ := reader.ReadString('\n')
		fmt.Printf("Username: ")
		username, _ := reader.ReadString('\n')
		fmt.Printf("Password: ")
		password, _ := gopass.GetPasswd()

		authInfo := aceproject.AuthInfo{
			AccountID: strings.TrimSpace(account),
			Username:  strings.TrimSpace(username),
			Password:  string(password),
		}
		loginSvc := aceproject.NewLoginService(&http.Client{})

		guidInfo, _, err := loginSvc.Login(&authInfo)

		if err != nil || guidInfo == nil || len(guidInfo.GUID) != 36 {
			fmt.Printf("Login fail:: %v\n", err)
			os.Exit(1)
		}
		viper.Set("GUID", guidInfo.GUID)
		fmt.Printf("Login success!\n")

		user, err := user.Current()
		var filename string
		if cfgFile != "" {
			filename = cfgFile
		} else {
			filename = user.HomeDir + "/.acectl.json"
		}
		cfgContent, err := json.Marshal(guidInfo)
		ioutil.WriteFile(filename, cfgContent, 0600)
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
}
