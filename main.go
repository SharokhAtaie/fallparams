package main

import (
	"fmt"
	"github.com/SharokhAtaie/fallparams/banner"
	"github.com/SharokhAtaie/fallparams/commands"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"io/ioutil"
	"log"
)

type options struct {
	url    string
	output string
	silent bool
}

func main() {
	opt := &options{}
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`
___________        ____  ___  __________      
\_   _____/_____   |  |  |  | \______   \_____  _______ _____     _____    ______ 
 |    __)  \__  \  |  |  |  |  |     ___/\__  \ \_  __ \\__  \   /     \  /  ___/ 
 |     \    / __ \_|  |__|  |__|    |     / __ \_|  | \/ / __ \_|  Y Y  \ \___ \  
 \___  /   (____  /|____/|____/|____|    (____  /|__|   (____  /|__|_|  //____  > 
     \/         \/                            \/             \/       \/      \/

	A tool for get all parameters from html source`)
	flagSet.StringVarP(&opt.url, "url", "u", "", "url for get parameters")
	flagSet.StringVarP(&opt.output, "output", "o", "", "filename for save output")
	flagSet.BoolVarP(&opt.silent, "silent", "s", false, "Show silent output")

	if err := flagSet.Parse(); err != nil {
		log.Fatalf("Could not parse flags: %s\n", err)
	}

	if opt.url == "" {
		banner.Banner()
		gologger.Print().Msgf("Flags:\n")
		gologger.Print().Msgf("\t-url,	 -u 		Url for Get All parameters")
		gologger.Print().Msgf("\t-output, -o 		Filename for save output")
		gologger.Print().Msgf("\t-silent, -s 		Show silent output")
		return
	}

	if opt.silent != true {
		banner.Banner()
	}

	data := commands.Regex(opt.url)

	// save output
	if opt.output != "" {
		b := []byte(data)
		err1 := ioutil.WriteFile(opt.output, b, 0644)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
	fmt.Println(data)
}
