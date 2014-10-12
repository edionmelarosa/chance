package chance

import (
	"github.com/nu7hatch/gouuid"
)

func UUIDString() string {
	u, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return u.String()
}
