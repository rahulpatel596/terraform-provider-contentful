.PHONY: build
build:
	sudo -S docker build -t terraform-provider-contentful -f Dockerfile-test .

.PHONY: test-unit
test-unit: build
	sudo docker run \
		-e CONTENTFUL_MANAGEMENT_TOKEN=${CONTENTFUL_MANAGEMENT_TOKEN} \
		-e CONTENTFUL_ORGANIZATION_ID=${CONTENTFUL_ORGANIZATION_ID} \
		-e SPACE_ID=${SPACE_ID} \
		-e "TF_ACC=true" \
		terraform-provider-contentful \
		go test ./... -v

.PHONY: interactive
interactive:
	sudo -S docker run -it \
		-v $(shell pwd):/go/src/github.com/labd/terraform-contentful \
		-e CONTENTFUL_MANAGEMENT_TOKEN=${CONTENTFUL_MANAGEMENT_TOKEN} \
        -e CONTENTFUL_ORGANIZATION_ID=${CONTENTFUL_ORGANIZATION_ID} \
        -e SPACE_ID=${SPACE_ID} \
		terraform-provider-contentful \
		bash
