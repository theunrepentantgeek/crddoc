config
======

<a id="config"></a>Config
-------------------------

.

| Property    | Description                                                                                                                                                  | Type                |
|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------|
| editors     | Editors allow you to make precision changes to the documentation output. Editors are applied in the order specified.                                         | [Editor[]](#editor) |
| prettyPrint | PrettyPrint controls whether the Markdown output is pretty-printed or not. Defaults to true.                                                                 | bool                |
| typeFilters | TypeFilters allow you to filter out types from the output. Filters are applied in the order specified, with earlier filters taking priority over later ones. | [Filter[]](#filter) |

<a id="editor"></a>Editor
-------------------------

Editor represents a point modification to make to exported documentation

Used by: [Config](#config).

| Property      | Description                                                                                                                                                     | Type          |
|---------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------|
| context       | Context is a regex identifying a substring to modify, limiting the scope of the search and replace. If omitted, the entire string is eligible for modification. | string        |
| contextRegexp |                                                                                                                                                                 | regexp.Regexp |
| replace       | Replace is the string to substitute for the search regex.                                                                                                       | string        |
| search        | Search is a regex identifying a substring to replace.                                                                                                           | string        |
| searchRegexp  |                                                                                                                                                                 | regexp.Regexp |

<a id="filter"></a>Filter
-------------------------

Used by: [Config](#config).

| Property | Description                                                                     | Type   |
|----------|---------------------------------------------------------------------------------|--------|
| because  | Because is an explanation of why this filter is being applied and what it does. | string |
| exclude  | Exclude is a glob identifying types to exclude from the output.                 | string |
| include  | Include is a glob identifying types to include in the output.                   | string |
