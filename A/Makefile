
LOCAL_PREFIX := local

define build
$(eval lower=$(shell echo $(1) | tr '[:upper:]' '[:lower:]'))
docker build \
	--build-arg SERVICE_DIR=$(1) \
	--build-context repo_b=../B \
	--tag $(LOCAL_PREFIX)/$(lower):latest \
	-f Dockerfile \
	.
endef

SERVICES := A/SA A/SB B/SC B/SD

.PHONY: all $(SERVICES)


$(SERVICES): 
	@echo "Building $@..."
	$(call build,$@)

all: $(SERVICES)


clean:
	@echo "Cleaning up..."
	docker images --format "{{.Repository}}:{{.Tag}}" \
		| grep "^${LOCAL_PREFIX}[/:]" \
		| xargs -r -n1 docker rmi -f || true
	docker system prune -f || true
	@echo "Cleanup complete."