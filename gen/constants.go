package gen

// copy from io/swagger/codegen/CodegenConstants.java
const (
	APIS             = "apis"
	MODELS           = "models"
	SUPPORTING_FILES = "supportingFiles"
	MODEL_TESTS      = "modelTests"
	MODEL_DOCS       = "modelDocs"
	API_TESTS        = "apiTests"
	API_DOCS         = "apiDocs"
	WITH_XML         = "withXml"
	/* /end System Properties */

	API_PACKAGE      = "apiPackage"
	API_PACKAGE_DESC = "package for generated api classes"

	MODEL_PACKAGE      = "modelPackage"
	MODEL_PACKAGE_DESC = "package for generated models"

	TEMPLATE_DIR = "templateDir"

	ALLOW_UNICODE_IDENTIFIERS      = "allowUnicodeIdentifiers"
	ALLOW_UNICODE_IDENTIFIERS_DESC = "boolean, toggles whether unicode identifiers are allowed in names or not, default is false"

	INVOKER_PACKAGE      = "invokerPackage"
	INVOKER_PACKAGE_DESC = "root package for generated code"

	PHP_INVOKER_PACKAGE      = "phpInvokerPackage"
	PHP_INVOKER_PACKAGE_DESC = "root package for generated php code"

	PERL_MODULE_NAME      = "perlModuleName"
	PERL_MODULE_NAME_DESC = "root module name for generated perl code"

	PYTHON_PACKAGE_NAME      = "pythonPackageName"
	PYTHON_PACKAGE_NAME_DESC = "package name for generated python code"

	GROUP_ID      = "groupId"
	GROUP_ID_DESC = "groupId in generated pom.xml"

	ARTIFACT_ID      = "artifactId"
	ARTIFACT_ID_DESC = "artifactId in generated pom.xml"

	ARTIFACT_VERSION      = "artifactVersion"
	ARTIFACT_VERSION_DESC = "artifact version in generated pom.xml"

	ARTIFACT_URL      = "artifactUrl"
	ARTIFACT_URL_DESC = "artifact URL in generated pom.xml"

	ARTIFACT_DESCRIPTION      = "artifactDescription"
	ARTIFACT_DESCRIPTION_DESC = "artifact description in generated pom.xml"

	SCM_CONNECTION      = "scmConnection"
	SCM_CONNECTION_DESC = "SCM connection in generated pom.xml"

	SCM_DEVELOPER_CONNECTION      = "scmDeveloperConnection"
	SCM_DEVELOPER_CONNECTION_DESC = "SCM developer connection in generated pom.xml"

	SCM_URL      = "scmUrl"
	SCM_URL_DESC = "SCM URL in generated pom.xml"

	DEVELOPER_NAME      = "developerName"
	DEVELOPER_NAME_DESC = "developer name in generated pom.xml"

	DEVELOPER_EMAIL      = "developerEmail"
	DEVELOPER_EMAIL_DESC = "developer email in generated pom.xml"

	DEVELOPER_ORGANIZATION      = "developerOrganization"
	DEVELOPER_ORGANIZATION_DESC = "developer organization in generated pom.xml"

	DEVELOPER_ORGANIZATION_URL      = "developerOrganizationUrl"
	DEVELOPER_ORGANIZATION_URL_DESC = "developer organization URL in generated pom.xml"

	LICENSE_NAME      = "licenseName"
	LICENSE_NAME_DESC = "The name of the license"

	LICENSE_URL      = "licenseUrl"
	LICENSE_URL_DESC = "The URL of the license"

	SOURCE_FOLDER      = "sourceFolder"
	SOURCE_FOLDER_DESC = "source folder for generated code"

	IMPL_FOLDER      = "implFolder"
	IMPL_FOLDER_DESC = "folder for generated implementation code"

	LOCAL_VARIABLE_PREFIX      = "localVariablePrefix"
	LOCAL_VARIABLE_PREFIX_DESC = "prefix for generated code members and local variables"

	SERIALIZABLE_MODEL      = "serializableModel"
	SERIALIZABLE_MODEL_DESC = "boolean - toggle \"implements Serializable\" for generated models"

	SERIALIZE_BIG_DECIMAL_AS_STRING      = "bigDecimalAsString"
	SERIALIZE_BIG_DECIMAL_AS_STRING_DESC = "Treat BigDecimal values as Strings to avoid precision loss."

	LIBRARY      = "library"
	LIBRARY_DESC = "library template (sub-template)"

	SORT_PARAMS_BY_REQUIRED_FLAG      = "sortParamsByRequiredFlag"
	SORT_PARAMS_BY_REQUIRED_FLAG_DESC = "Sort method arguments to place required parameters before optional parameters."

	USE_DATETIME_OFFSET      = "useDateTimeOffset"
	USE_DATETIME_OFFSET_DESC = "Use DateTimeOffset to model date-time properties"

	ENSURE_UNIQUE_PARAMS      = "ensureUniqueParams"
	ENSURE_UNIQUE_PARAMS_DESC = "Whether to ensure parameter names are unique in an operation (rename parameters that are not)."

	PROJECT_NAME    = "projectName"
	PACKAGE_NAME    = "packageName"
	PACKAGE_VERSION = "packageVersion"

	PACKAGE_TITLE            = "packageTitle"
	PACKAGE_TITLE_DESC       = "Specifies an AssemblyTitle for the .NET Framework global assembly attributes stored in the AssemblyInfo file."
	PACKAGE_PRODUCTNAME      = "packageProductName"
	PACKAGE_PRODUCTNAME_DESC = "Specifies an AssemblyProduct for the .NET Framework global assembly attributes stored in the AssemblyInfo file."
	PACKAGE_DESCRIPTION      = "packageDescription"
	PACKAGE_DESCRIPTION_DESC = "Specifies a AssemblyDescription for the .NET Framework global assembly attributes stored in the AssemblyInfo file."
	PACKAGE_COMPANY          = "packageCompany"
	PACKAGE_COMPANY_DESC     = "Specifies an AssemblyCompany for the .NET Framework global assembly attributes stored in the AssemblyInfo file."
	PACKAGE_AUTHORS          = "packageAuthors"
	PACKAGE_AUTHORS_DESC     = "Specifies Authors property in the .NET Core project file."
	PACKAGE_COPYRIGHT        = "packageCopyright"
	PACKAGE_COPYRIGHT_DESC   = "Specifies an AssemblyCopyright for the .NET Framework global assembly attributes stored in the AssemblyInfo file."

	POD_VERSION = "podVersion"

	OPTIONAL_METHOD_ARGUMENT      = "optionalMethodArgument"
	OPTIONAL_METHOD_ARGUMENT_DESC = "Optional method argument, e.g. void square(int x=10) (.net 4.0+ only)."

	OPTIONAL_ASSEMBLY_INFO      = "optionalAssemblyInfo"
	OPTIONAL_ASSEMBLY_INFO_DESC = "Generate AssemblyInfo.cs."

	NETCORE_PROJECT_FILE      = "netCoreProjectFile"
	NETCORE_PROJECT_FILE_DESC = "Use the new format (.NET Core) for .NET project files (.csproj)."

	USE_COLLECTION      = "useCollection"
	USE_COLLECTION_DESC = "Deserialize array types to Collection<T> instead of List<T>."

	INTERFACE_PREFIX      = "interfacePrefix"
	INTERFACE_PREFIX_DESC = "Prefix interfaces with a community standard or widely accepted prefix."

	RETURN_ICOLLECTION      = "returnICollection"
	RETURN_ICOLLECTION_DESC = "Return ICollection<T> instead of the concrete type."

	OPTIONAL_PROJECT_FILE      = "optionalProjectFile"
	OPTIONAL_PROJECT_FILE_DESC = "Generate {PackageName}.csproj."

	OPTIONAL_PROJECT_GUID      = "packageGuid"
	OPTIONAL_PROJECT_GUID_DESC = "The GUID that will be associated with the C# project"

	MODEL_PROPERTY_NAMING      = "modelPropertyNaming"
	MODEL_PROPERTY_NAMING_DESC = "Naming convention for the property: 'camelCase', 'PascalCase', 'snake_case' and 'original', which keeps the original name"

	DOTNET_FRAMEWORK      = "targetFramework"
	DOTNET_FRAMEWORK_DESC = "The target .NET framework version."

	//enum MODEL_PROPERTY_NAMING_TYPE {camelCase, PascalCase, snake_case, original}
	//enum ENUM_PROPERTY_NAMING_TYPE {camelCase, PascalCase, snake_case, original, UPPERCASE}

	ENUM_PROPERTY_NAMING      = "enumPropertyNaming"
	ENUM_PROPERTY_NAMING_DESC = "Naming convention for enum properties: 'camelCase', 'PascalCase', 'snake_case', 'UPPERCASE', and 'original'"

	MODEL_NAME_PREFIX      = "modelNamePrefix"
	MODEL_NAME_PREFIX_DESC = "Prefix that will be prepended to all model names. Default is the empty string."

	MODEL_NAME_SUFFIX      = "modelNameSuffix"
	MODEL_NAME_SUFFIX_DESC = "Suffix that will be appended to all model names. Default is the empty string."

	OPTIONAL_EMIT_DEFAULT_VALUES      = "optionalEmitDefaultValues"
	OPTIONAL_EMIT_DEFAULT_VALUES_DESC = "Set DataMember's EmitDefaultValue."

	GIT_USER_ID      = "gitUserId"
	GIT_USER_ID_DESC = "Git user ID, e.g. swagger-api."

	GIT_REPO_ID      = "gitRepoId"
	GIT_REPO_ID_DESC = "Git repo ID, e.g. swagger-codegen."

	RELEASE_NOTE      = "releaseNote"
	RELEASE_NOTE_DESC = "Release note, default to 'Minor update'."

	HTTP_USER_AGENT      = "httpUserAgent"
	HTTP_USER_AGENT_DESC = "HTTP user agent, e.g. codegen_csharp_api_client, default to 'Swagger-Codegen/{packageVersion}}/{language}'"

	SUPPORTS_ES6      = "supportsES6"
	SUPPORTS_ES6_DESC = "Generate code that conforms to ES6."

	SUPPORTS_ASYNC      = "supportsAsync"
	SUPPORTS_ASYNC_DESC = "Generate code that supports async operations."

	EXCLUDE_TESTS      = "excludeTests"
	EXCLUDE_TESTS_DESC = "Specifies that no tests are to be generated."

	// Not user-configurable. System provided for use in templates.

	GENERATE_APIS     = "generateApis"
	GENERATE_API_DOCS = "generateApiDocs"

	GENERATE_API_TESTS      = "generateApiTests"
	GENERATE_API_TESTS_DESC = "Specifies that api tests are to be generated."

	// Not user-configurable. System provided for use in templates.
	GENERATE_MODELS     = "generateModels"
	GENERATE_MODEL_DOCS = "generateModelDocs"

	GENERATE_MODEL_TESTS      = "generateModelTests"
	GENERATE_MODEL_TESTS_DESC = "Specifies that model tests are to be generated."

	HIDE_GENERATION_TIMESTAMP      = "hideGenerationTimestamp"
	HIDE_GENERATION_TIMESTAMP_DESC = "Hides the generation timestamp when files are generated."

	GENERATE_PROPERTY_CHANGED      = "generatePropertyChanged"
	GENERATE_PROPERTY_CHANGED_DESC = "Specifies that models support raising property changed events."

	NON_PUBLIC_API      = "nonPublicApi"
	NON_PUBLIC_API_DESC = "Generates code with reduced access modifiers allows embedding elsewhere without exposing non-public API calls to consumers."

	VALIDATABLE      = "validatable"
	VALIDATABLE_DESC = "Generates self-validatable models."

	IGNORE_FILE_OVERRIDE      = "ignoreFileOverride"
	IGNORE_FILE_OVERRIDE_DESC = "Specifies an override location for the .swagger-codegen-ignore file. Most useful on initial generation."

	REMOVE_OPERATION_ID_PREFIX      = "removeOperationIdPrefix"
	REMOVE_OPERATION_ID_PREFIX_DESC = "Remove prefix of operationId, e.g. config_getId => getId"

	STRIP_PACKAGE_NAME      = "stripPackageName"
	STRIP_PACKAGE_NAME_DESC = "Whether to strip leading dot-separated packages from generated model classes"
)
