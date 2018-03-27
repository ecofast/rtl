# rtl
Package rtl implements Cross-Platform Runtime Library in the way of which Delphi(2007) has done.</br></br>

# How to use</br>
```Go
import (
	"fmt"
	
	"github.com/ecofast/rtl/inifiles"
	"github.com/ecofast/rtl/sysutils"
)

var (
	clientListenPort int = 7788
	acceptTimeout        = 2 // sec
)

iniName := sysutils.ChangeFileExt(os.Args[0], ".ini")
ini := inifiles.New(iniName, true)
clientListenPort = ini.ReadInt("setup", "clientlistenport", clientListenPort)
acceptTimeout = ini.ReadInt("setup", "clientaccepttimeout", acceptTimeout)
fmt.Printf("client listenport: %d\n", clientListenPort)
fmt.Printf("client accept timeout: %ds\n", acceptTimeout)
```
