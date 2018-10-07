gocpustat
===================

Get CPU stats for a Linux machine

Reads and parses /proc/cpuinfo and returns useful stats

Usage
---------------

```go
import (
	"fmt"
	cpustat "github.com/geethanjalieswaran/gocpustat"
)

func main () {
	fmt.Printf("%+v", cpustat.GetCPUStat())
}
```
License
-------

gocpustat is distributed under the GNU GENERAL PUBLIC LICENSE.
