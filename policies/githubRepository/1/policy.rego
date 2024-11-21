# this policy checks if the repo github collaborators are not in the blocked list

package policies

import rego.v1

blocked_users = {"blocked_user_1", "blocked_user_2", "blocked_user_3"}

default pass := false

pass if {
    not any_blocked_collaborators(input.collaborators)
}

any_blocked_collaborators(collaborators) if {
    collaborator := collaborators[_]
    blocked_users[collaborator.login]
}
