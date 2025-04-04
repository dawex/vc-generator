package compliancelogs

//go:generate oapi-codegen --old-config-style -o ports/handler_types.gen.go -package=ports -include-tags=Compliance-Logs -generate types ../../../documentation/oas/dist/vc-generator.yaml
//go:generate oapi-codegen --old-config-style -o ports/handler_api.gen.go -package=ports -include-tags=Compliance-Logs -generate chi-server ../../../documentation/oas/dist/vc-generator.yaml
