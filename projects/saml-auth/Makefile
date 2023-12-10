.PHONY: certs
.DEFAULT_GOAL: certs

METAFILE= saml-test-$(USER)-$(shell date -Iseconds -r 1533415339).xml

certs:
	openssl req -x509 -newkey rsa:2048 -keyout myservice.key -out myservice.cert -days 365 -nodes -subj "/CN=myservice.example.com"

metadata:
	curl -s localhost:8000/saml/metadata > $(METAFILE)

clean:
	rm *.cert *.key saml-test-*.xml
