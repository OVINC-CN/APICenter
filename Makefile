swagger:
	scripts/swagger.sh

tidy:
	scripts/tidy.sh

lint:
	scripts/tidy.sh
	scripts/swagger.sh