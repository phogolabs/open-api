package com.phogolabs.codegen;

import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.responses.ApiResponse;
import io.swagger.v3.oas.models.Operation;
import io.swagger.v3.oas.models.Paths;
import org.apache.http.util.TextUtils;
import org.openapitools.codegen.*;
import org.openapitools.codegen.utils.StringUtils;
import java.io.File;
import java.util.*;

public class GolangServerGenerator extends DefaultCodegen {
    /**
     * Configures the type of generator.
     *
     * @return the CodegenType for this generator
     * @see org.openapitools.codegen.CodegenType
     */
    public CodegenType getTag() {
        return CodegenType.SERVER;
    }

    /**
     * Configures a friendly name for the generator.  This will be used by the generator
     * to select the library with the -g flag.
     *
     * @return the friendly name for the generator
     */
    public String getName() {
        return "golang-server";
    }

    /**
     * Returns human-friendly help for the generator.  Provide the consumer with help
     * tips, parameters here
     *
     * @return A string value for the help message
     */
    public String getHelp() {
        return "Generates an HTTP Server";
    }

    public GolangServerGenerator() {
        super();

        // set defaults packages
        apiPackage = "handler";
        modelPackage = "handler";
        testPackage = "handler_test";

        // templates
        embeddedTemplateDir = templateDir = "golang-server";

        // model templates
        modelTemplateFiles.clear();
        // model document templates
        modelDocTemplateFiles.put("doc_model.mustache", ".md");

        // handler templates
        apiTemplateFiles.put("handler.mustache", ".go");
        apiTemplateFiles.put("handler_model.mustache", "_model.go");
        // handler test templates
        apiTestTemplateFiles.put("handler_test.mustache", ".go");
        // handler document templates
        apiDocTemplateFiles.put("doc_handler.mustache", ".md");

        supportingFiles.add(new SupportingFile("README.mustache", "", "README.md"));
        supportingFiles.add(new SupportingFile("mod.mustache", "", "go.mod"));
        supportingFiles.add(new SupportingFile("doc.mustache", "documentation", "README.md"));
        supportingFiles.add(new SupportingFile("model.mustache", "service/handler", "model.go"));
        supportingFiles.add(new SupportingFile("gitkeep.mustache", "service/middleware", ".gitkeep"));
        supportingFiles.add(new SupportingFile("suite_test.mustache", "service/handler", "suite_test.go"));

        // reserved words
        setReservedWordsLowerCase(
                Arrays.asList(
                        // data type
                        "string", "bool", "uint", "uint8", "uint16", "uint32", "uint64",
                        "int", "int8", "int16", "int32", "int64", "float32", "float64",
                        "complex64", "complex128", "rune", "byte", "uintptr",

                        "break", "default", "func", "interface", "select",
                        "case", "defer", "go", "map", "struct",
                        "chan", "else", "goto", "package", "switch",
                        "const", "fallthrough", "if", "range", "type",
                        "continue", "for", "import", "return", "var", "error", "nil")
        );


        defaultIncludes = new HashSet<String>(
                Arrays.asList(
                        "DateTime",
                        "map",
                        "array",
                        "double",
                        "long",
                        "short",
                        "char",
                        "float",
                        "String",
                        "boolean",
                        "Boolean",
                        "Double",
                        "Void",
                        "int",
                        "integer",
                        "Integer",
                        "Long",
                        "Float")
        );

        typeMapping.put("integer", "int32");
        typeMapping.put("long", "int64");
        typeMapping.put("number", "float32");
        typeMapping.put("float", "float32");
        typeMapping.put("double", "float64");
        typeMapping.put("boolean", "bool");
        typeMapping.put("string", "string");
        typeMapping.put("UUID", "string");
        typeMapping.put("date", "string");
        typeMapping.put("DateTime", "time.Time");
        typeMapping.put("password", "string");
        typeMapping.put("object", "map[string]interface{}");
        typeMapping.put("File", "io.ReadCloser");
        typeMapping.put("file", "io.ReadCloser");
        typeMapping.put("binary", "io.ReadCloser");
        typeMapping.put("ByteArray", "io.ReadCloser");

        languageSpecificPrimitives = new HashSet<String>(
                Arrays.asList(
                        "string",
                        "bool",
                        "uint",
                        "uint32",
                        "uint64",
                        "int",
                        "int32",
                        "int64",
                        "float32",
                        "float64",
                        "complex64",
                        "complex128",
                        "rune",
                        "byte")
        );

        // type mappings
        typeMapping.clear();
        typeMapping.put("integer", "int32");
        typeMapping.put("long", "int64");
        typeMapping.put("number", "float32");
        typeMapping.put("float", "float32");
        typeMapping.put("double", "float64");
        typeMapping.put("boolean", "bool");
        typeMapping.put("string", "string");
        typeMapping.put("UUID", "string");
        typeMapping.put("date", "string");
        typeMapping.put("DateTime", "time.Time");
        typeMapping.put("password", "string");
        typeMapping.put("object", "map[string]interface{}");
        typeMapping.put("File", "io.ReadCloser");
        typeMapping.put("file", "io.ReadCloser");
        typeMapping.put("binary", "io.ReadCloser");
        typeMapping.put("ByteArray", "io.ReadCloser");

        // import mappings
        importMapping = new HashMap<String, String>();
        importMapping.put("DateTime", "time");
        importMapping.put("time.Time", "time");
        importMapping.put("File", "io");
        importMapping.put("UUID", "github.com/phogolabs/schema");
        importMapping.put("io.ReadCloser", "io");
    }

    @Override
    public void processOpts() {
        super.processOpts();

        additionalProperties.put(CodegenConstants.PROJECT_NAME, this.projectName());
        additionalProperties.put(CodegenConstants.PACKAGE_NAME, this.projectName());
        additionalProperties.put(CodegenConstants.PACKAGE_VERSION, "1.0.0");

        supportingFiles.add(new SupportingFile("main.mustache", this.mainFilePath()));
    }

     @Override
     public Map<String, Object> postProcessOperationsWithModels(Map<String, Object> objs, List<Object> allModels) {
         objs = super.postProcessOperationsWithModels(objs, allModels);

         @SuppressWarnings("unchecked")
         Map<String, Object> objectMap = (Map<String, Object>) objs.get("operations");

         @SuppressWarnings("unchecked")
         List<CodegenOperation> operations = (List<CodegenOperation>) objectMap.get("operation");

         Set<String> imports = new TreeSet<String>();

         for (CodegenOperation operation : operations) {
             operation.httpMethod = StringUtils.camelize(operation.httpMethod.toLowerCase(Locale.ROOT));

             List<Map<String, Object>> input = new ArrayList<Map<String, Object>>();

             for (CodegenParameter param: operation.allParams) {
                 if (importMapping.containsKey(param.dataType)) {
                     imports.add(importMapping.get(param.dataType));
                 }

                 if (param.isFile) {
                     imports.add("mime/multipart");
                 }

                 if (param.isModel || param.isFreeFormObject) {
                     imports.add("encoding/json");
                 }

                 if(!param.isBodyParam) {
                     input.add(createMapping("param", param));
                 }

                 if (param.isBodyParam) {
                     operation.vendorExtensions.put("hasBodyParams", true);
                 }

                 if(param.isPathParam) {
                   operation.vendorExtensions.put("hasPathParams", true);
                 }

                 if(param.isQueryParam) {
                   operation.vendorExtensions.put("hasQueryParams", true);
                 }

                 if(param.isHeaderParam) {
                   operation.vendorExtensions.put("hasHeaderParams", true);
                 }
             }

             for(CodegenProperty header: operation.responseHeaders) {
                 if (importMapping.containsKey(header.dataType)) {
                     imports.add(importMapping.get(header.dataType));
                 }

                 imports.add("github.com/phogolabs/rest");
             }

             for (CodegenResponse response: operation.responses) {
                 if (response.code.startsWith("2") || response.code.equals("default")) {
                   operation.vendorExtensions.put("defaultResponse", response);
                   break;
                 }

                 if (response.isModel || response.isFreeFormObject) {
                     imports.add("encoding/json");
                 }

                 if (importMapping.containsKey(response.dataType)) {
                     imports.add(importMapping.get(response.dataType));
                 }
             }

             operation.vendorExtensions.put("hasInputParams", input.size() > 0);
             operation.vendorExtensions.put("inputParams", input);
             operation.vendorExtensions.put("hello", "world");

             if (input.size() > 0) {
                 imports.add("github.com/phogolabs/rest");
             }
         }

         objs.put("imports", this.buildImports(imports));
         return objs;
     }

    @Override
    public Map<String, Object> postProcessSupportingFileData(Map<String, Object> bundle) {
        bundle = super.postProcessSupportingFileData(bundle);
        this.addModelImports(bundle);
        return bundle;
    }

    private void addModelImports(Map<String, Object> bundle) {
        @SuppressWarnings("unchecked")
        List<Map<String, CodegenModel>> models = (List<Map<String, CodegenModel>>) bundle.get("models");

        Set<String> unique = new TreeSet<String>();

        for (Map<String, CodegenModel> model : models) {
            unique.addAll(model.get("model").imports);
        }

        bundle.put("imports", this.buildImports(unique));
    }

    private List<Map<String, Object>> buildImports(Set<String> unique) {
        List<Map<String, Object>> imports = new ArrayList<Map<String, Object>>();

        for (String value : unique) {
            imports.add(createMapping("import", value));
        }

        return imports;
    }

    @Override
    public void postProcessModelProperty(CodegenModel model, CodegenProperty property) {
        super.postProcessModelProperty(model, property);

        String baseType = "";

        if (property.isModel) {
            baseType = property.dataType;
        } else if (property.isListContainer) {
            baseType = property.items.dataType;
        }

        model.imports.remove(baseType);

        if (importMapping.containsKey(baseType)) {
            model.imports.add(importMapping.get(baseType));
        }

        this.buildPropertyTag(property);
    }

    private void buildPropertyTag(CodegenProperty property) {
        List<String> tags = new ArrayList<String>();

        List<String> jsonTag = new ArrayList<String>();
        jsonTag.add(property.baseName);

        if (!property.required) {
            jsonTag.add("omitempty");
        }

        tags.add(String.format("json:\"%s\"", String.join(",", jsonTag)));

        if (this.withXml()) {
            List<String> xmlTag = new ArrayList<String>();
            xmlTag.add(property.baseName);

            if (property.isXmlAttribute) {
                xmlTag.add("attr");
            }

            tags.add(String.format("xml:\"%s\"", String.join(",", xmlTag)));
        }

        if (!isNull(property.defaultValue)) {
            tags.add(String.format("default:\"%s\"", property.defaultValue));
        }

        this.buildPropertyValidationTag(property, tags);
        property.vendorExtensions.put("tag", String.format("`%s`", String.join(" ", tags)));
    }


    private void buildPropertyValidationTag(CodegenProperty property, List<String> tags) {
        List<String> validation = new ArrayList<String>();

        if (property.required) {
            validation.add("required");
        }

        if (property.hasValidation) {
            Object min = null;
            Object max = null;

            if (property.baseType == "string") {
                min = property.minLength;
                max = property.maxLength;

                if (!isNull(property.pattern)) {
                    validation.add(String.format("regex=%s", property.pattern));
                }
            }

            if (property.isNumeric) {
                min = property.minimum;
                max = property.maximum;
            }

            if (property.isListContainer) {
                min = property.minItems;
                max = property.maxItems;
            }

            if (min != null) {
                if (property.exclusiveMinimum) {
                    validation.add(String.format("gt=%s", min));
                } else {
                    validation.add(String.format("gte=%s", min));
                }
            }

            if (max != null) {
                if (property.exclusiveMaximum) {
                    validation.add(String.format("lt=%s", max));
                } else {
                    validation.add(String.format("lte=%s", max));
                }
            }
        }

        if (validation.size() > 0) {
            tags.add(String.format("validate:\"%s\"", String.join(",", validation)));
        } else {
            tags.add("validate:\"-\"");
        }
    }

    @Override
    public void postProcessParameter(CodegenParameter param) {
        buildParameterTag(param);
    }

    private void buildParameterTag(CodegenParameter property) {
        List<String> tags = new ArrayList<String>();
        tags.add("json:\"-\"");

        if (!withXml()) {
            tags.add("xml:\"-\"");
        }

        if (property.isFormParam) {
            tags.add(String.format("form:\"%s\"", property.baseName));
        }

        if (property.isHeaderParam) {
            tags.add(String.format("header:\"%s\"", property.baseName));
        }

        if (property.isPathParam) {
            tags.add(String.format("path:\"%s\"", property.baseName));
        }

        if (property.isQueryParam) {
            tags.add(String.format("query:\"%s\"", property.baseName));
        }

        if (!isNull(property.defaultValue)) {
            tags.add(String.format("default:\"%s\"", property.defaultValue));
        }

        List<String> validation = new ArrayList<String>();

        if (property.required) {
            validation.add("required");
        }

        if (property.hasValidation) {
            Object min = null;
            Object max = null;

            if (property.baseType == "string") {
                min = property.minLength;
                max = property.maxLength;

                if (!isNull(property.pattern)) {
                    validation.add(String.format("regex=%s", property.pattern));
                }
            }

            if (property.isNumeric) {
                min = property.minimum;
                max = property.maximum;
            }

            if (property.isListContainer) {
                min = property.minItems;
                max = property.maxItems;

                if (property.uniqueItems) {
                    validation.add("unique");
                }
            }

            if (min != null) {
                if (property.exclusiveMinimum) {
                    validation.add(String.format("gt=%s", min));
                } else {
                    validation.add(String.format("gte=%s", min));
                }
            }

            if (max != null) {
                if (property.exclusiveMaximum) {
                    validation.add(String.format("lt=%s", max));
                } else {
                    validation.add(String.format("lte=%s", max));
                }
            }
        }

        if (validation.size() > 0) {
            tags.add(String.format("validate:\"%s\"", String.join(",", validation)));
        } else{
            tags.add("validate:\"-\"");
        }

        property.vendorExtensions.put("tag", String.format("`%s`", String.join(" ", tags)));
    }

    @Override
    public CodegenResponse fromResponse(OpenAPI openAPI, java.lang.String responseCode, ApiResponse response) {
        CodegenResponse result = super.fromResponse(openAPI, responseCode, response);

        List<String> tags = new ArrayList<String>();
        tags.add("json:\"-\"");

        if (withXml()) {
            tags.add("xml:\"-\"");
        }

        for (CodegenProperty property: result.headers) {
            List<String> propertyTags = new ArrayList<String>();

            propertyTags.add(String.format("header:\"%s\"", property.baseName));
            this.buildPropertyValidationTag(property, propertyTags);

            property.vendorExtensions.put("tag", String.format("`%s`", String.join(" ", propertyTags)));
        }

        result.vendorExtensions.put("tag", String.format("`%s`", String.join(" ", tags)));
        return result;
    }

    @Override
    public String toModelImport(String name) {
        if (this.apiPackage().equals(this.modelPackage())) {
            return null;
        } else {
            //TODO: pass the full qualified golang package
            return modelPackage() + "." + name;
        }
    }

    private String projectName() {
        String name = new File(this.outputFolder()).getName();
        return name;
    }

    private String mainFilePath() {
        String name = this.projectName();
        return "cmd" + File.separator + name + File.separator + "main.go";
    }

    @Override
    public String modelFileFolder() {
        return super.outputFolder() + File.separator + "separator" + File.separator + this.modelPackage();
    }

    @Override
    public String modelTestFileFolder() {
        return this.modelFileFolder();
    }

    @Override
    public String modelDocFileFolder() {
        return super.outputFolder() + File.separator + "documentation";
    }

    @Override
    public String toModelDocFilename(String name) {
        return this.toModelName(name);
    }

    @java.lang.Override
    public java.lang.String toModelFilename(java.lang.String name) {
        return StringUtils.underscore(this.toModelName(name));
    }

    @java.lang.Override
    protected java.lang.String getOrGenerateOperationId(Operation operation, java.lang.String path, java.lang.String httpMethod) {
        return super.getOrGenerateOperationId(operation, path, httpMethod);
    }

    @Override
    public String toModelName(String name) {
//        if (!StringUtils.isEmpty(modelNamePrefix)) {
//            name = modelNamePrefix + "_" + name;
//        }
//
//        if (!StringUtils.isEmpty(modelNameSuffix)) {
//            name = name + "_" + modelNameSuffix;
//        }

        name = sanitizeName(name);

        // model name starts with number
        if (name.matches("^\\d.*")) {
            name = "model_" + name;
        }

        name = StringUtils.camelize(name, false);
        name = this.format(name);
        return name;
    }

    @Override
    public String toVarName(String name) {
        name = name.replaceAll("X-", "");
        name = name.replaceAll("x-", "");
        name = sanitizeName(name);

        if (name.matches("^[A-Z_]*$"))
            return name;

        // for reserved word or word starting with number, append _
        if (name.matches("^\\d.*"))
            name = "Var" + name;

        name = StringUtils.camelize(name, false);
        name = this.format(name);
        return name;
    }


    @Override
    public String toParamName(String name) {
        name = sanitizeName(name);
        name = StringUtils.camelize(name, false);
        name = this.format(name);
        return name;
    }

    @Override
    public String apiFileFolder() {
        return super.outputFolder() + File.separator + "service" + File.separator + this.apiPackage;
    }

    @Override
    public String apiTestFileFolder() {
        return this.apiFileFolder();
    }

    @Override
    public String apiDocFileFolder() {
        return super.outputFolder() + File.separator + "documentation";
    }

    @Override
    public String toApiFilename(String name) {
        name = StringUtils.underscore(name) + "_api";
        return name;
    }

    @Override
    public String toApiTestFilename(String name) {
        name = StringUtils.underscore(toApiName(name)) + "_test";
        return name;
    }

    @Override
    public String toApiDocFilename(String name) {
        name = toApiName(name);
        return name;
    }

    @Override
    public String toApiName(String name) {
        name = super.toApiName(name);
        name = this.format(name);
        return name;
    }

    @Override
    public String toApiVarName(String name) {
        name = super.toApiVarName(name);
        name = this.format(name);
        return name;
    }

    @Override
    public String toOperationId(String operationId) {
        operationId = super.toOperationId(operationId);
        operationId = this.format(operationId);
        return operationId;
    }

    @Override
    public String escapeQuotationMark(String input) {
        return input.replace("\"", "");
    }

    @Override
    public String escapeUnsafeCharacters(String input) {
        return input.replace("*/", "*_/").replace("/*", "/_*");
    }

    @Override
    public String escapeReservedWord(String name) {
        if (this.reservedWordsMappings().containsKey(name)) {
            return this.reservedWordsMappings().get(name);
        }

        return StringUtils.camelize(name) + '_';
    }

    @Override
    protected boolean needToImport(String type) {
        return !defaultIncludes.contains(type) && !languageSpecificPrimitives.contains(type);
    }

    private String format(String name) {
        name = name.replaceAll("Id", "ID");
        name = name.replaceAll("Api", "API");
        return name;
    }

    private String titleCase(final String input) {
        return input.substring(0, 1).toUpperCase(Locale.ROOT) + input.substring(1);
    }

    private boolean isNull(String value) {
        return value == null || value.equals("null");
    }

    private Map<String, Object> createMapping(String key, Object value) {
        Map<String, Object> data = new HashMap<String, Object>();
        data.put(key, value);
        return data;
    }

    private boolean withXml() {
        if (additionalProperties.containsKey(CodegenConstants.WITH_XML)) {
            return Boolean.parseBoolean(additionalProperties.get(CodegenConstants.WITH_XML).toString());
        }

        return false;
    }

}
