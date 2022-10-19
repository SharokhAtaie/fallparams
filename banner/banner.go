package banner

import (
	"github.com/projectdiscovery/gologger"
)

func Banner() {
	gologger.Print().Msgf(`
___________        ____  ___  __________      
\_   _____/_____   |  |  |  | \______   \_____  _______ _____     _____    ______ 
 |    __)  \__  \  |  |  |  |  |     ___/\__  \ \_  __ \\__  \   /     \  /  ___/ 
 |     \    / __ \_|  |__|  |__|    |     / __ \_|  | \/ / __ \_|  Y Y  \ \___ \  
 \___  /   (____  /|____/|____/|____|    (____  /|__|   (____  /|__|_|  //____  > 
     \/         \/                            \/             \/       \/      \/
`)
	gologger.Print().Msgf("   Created by Sharo_k_h :)\n\n")
}
