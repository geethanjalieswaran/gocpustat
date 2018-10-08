package example
import (
	"fmt"
	cpustat "github.com/geethanjalieswaran/gocpustat"
)

func main () {
	fmt.Printf("%+v", cpustat.GetCPUStat())
}
