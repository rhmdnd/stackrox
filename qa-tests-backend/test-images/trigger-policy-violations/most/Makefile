
# Bump this when you make changes
VERSION=0.19

.PHONY: image
image:
	docker build -t trigger-policy-violations/most:${VERSION} .
	docker tag trigger-policy-violations/most:${VERSION} us.gcr.io/acs-san-stackroxci/qa/trigger-policy-violations/most:${VERSION}
	docker push us.gcr.io/acs-san-stackroxci/qa/trigger-policy-violations/most:${VERSION}
