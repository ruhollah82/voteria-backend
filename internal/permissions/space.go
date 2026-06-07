package permissions

import (
	"github.com/yaghoubi-mn/voter/internal/config"
	"github.com/yaghoubi-mn/voter/internal/enums"
	"github.com/yaghoubi-mn/voter/internal/models"
)

type SubPermission interface {
	HasCreationPermission(userRole enums.Permissions) bool
	HasClosePermission(user models.User, sub models.Space) bool
	HasDeletePermission(user models.User, sub models.Space) bool
	HasEditPermission(user models.User, sub models.Space) bool
}

type subPermission struct {
	settings *config.Settings
}

func NewSubPermission(s *config.Settings) SubPermission {
	return &subPermission{
		settings: s,
	}
}

func (s *subPermission) HasCreationPermission(userRole enums.Permissions) bool {
	return userRole >= s.settings.SubCreationPermission
}

func (s *subPermission) HasClosePermission(user models.User, sub models.Space) bool {
	return user.Role >= int(s.settings.SubClosePermission) || (sub.OwnerID == user.ID && (sub.ClosedByRole == "owner" || sub.ClosedByRole == ""))
}

func (s *subPermission) HasDeletePermission(user models.User, sub models.Space) bool {
	return user.Role >= int(s.settings.SubDeletePermission) || sub.OwnerID == user.ID
}

func (s *subPermission) HasEditPermission(user models.User, sub models.Space) bool {
	return user.Role >= int(s.settings.SubDeletePermission) || sub.OwnerID == user.ID
}
