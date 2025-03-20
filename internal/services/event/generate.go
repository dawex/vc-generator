package event

//go:generate oapi-codegen --old-config-style -o ports/handler_types.gen.go -package=ports -include-tags=Dataset-Execution-Event -generate types ../../../documentation/oas/dist/vc-generator.yaml
//go:generate oapi-codegen --old-config-style -o ports/handler_api.gen.go -package=ports -include-tags=Dataset-Execution-Event -generate chi-server ../../../documentation/oas/dist/vc-generator.yaml
