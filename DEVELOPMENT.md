Check out [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) for
additional information about the OpenAPI project, including additional
libraries with support for other languages and more.

## How do I use this?

At this point, you've likely generated a client setup.  It will include something along these lines:

```
.
|- README.md    // this file
|- pom.xml      // build script
|-- src
|--- main
|---- java
|----- com.phogolabs.codegen.GolangServerGenerator.java // generator file
|---- resources
|----- golang-server // template files
|----- META-INF
|------ services
|------- org.openapitools.codegen.CodegenConfig
```

You _will_ need to make changes in at least the following:

`GolangServerGenerator.java`

Templates in this folder:

`src/main/resources/golang-server`

Once modified, you can run this:

```
mvn package
```

In your generator project. A single jar file will be produced in `target`. You
can now use that with [OpenAPI Generator](https://openapi-generator.tech):

For mac/linux:

```
java -cp /path/to/openapi-generator-cli.jar:/path/to/your.jar org.openapitools.codegen.OpenAPIGenerator generate -g golang-server -i /path/to/openapi.yaml -o ./test
```

(Do not forget to replace the values `/path/to/openapi-generator-cli.jar`, `/path/to/your.jar` and `/path/to/openapi.yaml` in the previous command)

For Windows users, you will need to use `;` instead of `:` in the classpath, e.g.

```
java -cp /path/to/openapi-generator-cli.jar;/path/to/your.jar org.openapitools.codegen.OpenAPIGenerator generate -g golang-server -i /path/to/openapi.yaml -o ./test
```

Now your templates are available to the client generator and you can write output values

## But how do I modify this?

The `GolangServerGenerator.java` has comments in it--lots of comments.  There
is no good substitute for reading the code more, though.  See how the
`GolangServerGenerator` implements `CodegenConfig`.
That class has the signature of all values that can be overridden.

For the templates themselves, you have a number of values available to you for
generation.  You can execute the `java` command from above while passing
different debug flags to show
the object you have available during client generation:

```
# The following additional debug options are available for all codegen targets:
# -DdebugOpenAPI prints the OpenAPI Specification as interpreted by the codegen
# -DdebugModels prints models passed to the template engine
# -DdebugOperations prints operations passed to the template engine
# -DdebugSupportingFiles prints additional data passed to the template engine

java -DdebugOperations -cp /path/to/openapi-generator-cli.jar:/path/to/your.jar org.openapitools.codegen.OpenAPIGenerator generate -g golang-server -i /path/to/openapi.yaml -o ./test
```
