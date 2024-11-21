# this policy checks if the repo collaborators site admin are in the allowed list

package policies

import rego.v1

default pass := false

default allowed_admin := "allowed_admin_username"

pass if {
    not any_not_allowed_admin(input.collaborators)
}

any_not_allowed_admin(collaborators) if {
    collaborator := collaborators[_]
    collaborator.site_admin == true
    collaborator.login != "Maytar"
}