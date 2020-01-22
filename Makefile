

generate:
	controller-gen "crd:trivialVersions=true" object paths="./..." rbac:roleName=manager-role webhook paths="./..."