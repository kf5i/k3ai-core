package cloud

import (
	"fmt"

	"github.com/kf5i/k3ai-core/internal/shared"
)

//Init call the specific function for cloud. The function will be used later
func Init(data shared.TargetCustoms) error {
	fmt.Print(data)
	return nil
}
