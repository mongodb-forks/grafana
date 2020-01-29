package api

import (
	"testing"

	m "github.com/grafana/grafana/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestIsRoleAssignable(t *testing.T) {
	// table test to  validate isRoleAssignable(currentRole, incomingRole)
	assert.True(t, isRoleAssignable("", m.ROLE_VIEWER))
	assert.True(t, isRoleAssignable(m.ROLE_VIEWER, m.ROLE_EDITOR))
	assert.True(t, isRoleAssignable(m.ROLE_VIEWER, m.ROLE_ADMIN))
	assert.True(t, isRoleAssignable(m.ROLE_EDITOR, m.ROLE_ADMIN))
	assert.False(t, isRoleAssignable(m.ROLE_ADMIN, m.ROLE_EDITOR))
	assert.False(t, isRoleAssignable(m.ROLE_ADMIN, m.ROLE_VIEWER))
	assert.False(t, isRoleAssignable(m.ROLE_EDITOR, m.ROLE_VIEWER))
	assert.True(t, isRoleAssignable(m.ROLE_VIEWER, m.ROLE_VIEWER))

	roles := map[int64]m.RoleType{}
	assert.True(t, isRoleAssignable(roles[0], m.ROLE_VIEWER))

}
