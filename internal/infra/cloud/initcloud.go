package cloud

import (
	"fmt"

	"github.com/kf5i/k3ai-core/internal/shared"
)

func CloudInit(data shared.TargetCustoms) error {
	fmt.Print(data)
	return nil
}
