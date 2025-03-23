KUBE_CONFIG="${HOME}/.kube/config"

run:
	go run ./main.go --kubeconfig="${KUBE_CONFIG}"

generate:
	bash hack/update-codegen.sh

verify:
	bash hack/verify-codegen.sh