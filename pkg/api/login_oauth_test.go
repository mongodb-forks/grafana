package api

import (
	"testing"

	m "github.com/grafana/grafana/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestIsRoleAssignable(t *testing.T) {

	viewer := m.RoleType("Viewer")
	editor := m.RoleType("Editor")
	admin := m.RoleType("Admin")

	// table test to  validate isRoleAssignable(currentRole, incomingRole)
	assert.True(t, isRoleAssignable("", viewer))
	assert.True(t, isRoleAssignable(viewer, editor))
	assert.True(t, isRoleAssignable(viewer, admin))
	assert.True(t, isRoleAssignable(editor, admin))
	assert.False(t, isRoleAssignable(admin, editor))
	assert.False(t, isRoleAssignable(admin, viewer))
	assert.False(t, isRoleAssignable(editor, viewer))
	assert.True(t, isRoleAssignable(viewer, viewer))
}
