package banner

import (
	"InceptionDeployer/tooling"
	"fmt"
)

func Show() {
	bannerText := Color.Cyan + Color.Bold + `
   ____                  __  _           _____       
  /  _/__  _______ ___  / /_(_)__  ___  / ___/__ ___ 
 _/ // _ \/ __/ -_) _ \/ __/ / _ \/ _ \/ (_ / -_) _ \
/___/_//_/\__/\__/ .__/\__/_/\___/_//_/\___/\__/_//_/
                /_/                                  
` + Color.Reset

	fmt.Println(bannerText)
	fmt.Println(Color.Dim + "  > shitty Project generator v.0" + Color.Reset)
	fmt.Println()
}