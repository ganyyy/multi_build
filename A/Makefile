
REGISTRY := local


SERVICE_TARGET      := A B C
SERVICE_TARGET_SUB  := $(addsuffix /XXX,$(SERVICE_TARGET))

.PHONY: $(SERVICE_TARGET) $(SERVICE_TARGET_SUB)

# 原始目标
$(SERVICE_TARGET):
	@echo ">> build $@"

# 二级目标 A/XXX 依赖 A
$(SERVICE_TARGET_SUB): %/XXX :
	@echo "P=$*"
	@echo "S=$(notdir $@)"

SERVICES := A/SA \
			A/SB \
			B/SC \
			B/SD
SERVICES_WITH_DATA := A/SA \
			B/SC

define build
$(eval NEED_DATA:=$(if $(filter $(1),$(SERVICES_WITH_DATA)),true,false))
$(eval LOWER:=$(shell echo $(1) | tr '[:upper:]' '[:lower:]'))
$(eval IMAGE_TAG:=$(REGISTRY)/$(LOWER):latest)
$(eval OLD_IMAGE_ID:=$(shell docker images -q $(IMAGE_TAG) 2>/dev/null | head -1))
DOCKER_BUILDKIT=1 docker build \
	--platform linux/amd64 \
	--build-arg SERVICE_DIR=$(1) \
	--build-arg NEED_DATA=$(NEED_DATA) \
	--build-context repo_b=../B \
	--tag $(IMAGE_TAG) \
	-f Dockerfile \
	.
@if [ -n "$(OLD_IMAGE_ID)" ]; then \
	NEW_IMAGE_ID=$$(docker images -q $(IMAGE_TAG)); \
	if [ "$(OLD_IMAGE_ID)" != "$$NEW_IMAGE_ID" ]; then \
		docker rmi $(OLD_IMAGE_ID) 2>/dev/null || true; \
	fi; \
fi
endef

.PHONY: all $(SERVICES)


$(SERVICES): 
	@echo "Building $@..."
	$(call build,$@)

all: $(SERVICES)


clean:
	@echo "Cleaning up..."
	docker images --format "{{.Repository}}:{{.Tag}}" \
		| grep "^${REGISTRY}[/:]" \
		| xargs -r -n1 docker rmi -f || true
	docker system prune -f || true
	@echo "Cleanup complete."