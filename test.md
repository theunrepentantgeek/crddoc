model
=====

| Metadata             | Value                                               |
|----------------------|-----------------------------------------------------|
| Group                |                                                     |
| Version              |                                                     |
| Module               | github.com/theunrepentantgeek/crddoc/internal/model |
| Property Optionality |                                                     |

<a id="DeclarationType"></a>DeclarationType
-------------------------------------------

| Value      | Description |
|------------|-------------|
| "Resource" |             |
| "Object"   |             |
| "Enum"     |             |

<a id="Enum"></a>Enum
---------------------

```mermaid
classDiagram
class Enum["Enum"] {
    description string[]
}
Enum *-- TypeReference
class TypeReference["TypeReference"]
Enum -- TypeReference : base
Enum -- Package : pkg
Enum -- PropertyReference : usage
Enum -- EnumValue : values
class TypeReference["TypeReference"] 
class Package["Package"] 
class PropertyReference["PropertyReference"] 
class EnumValue["EnumValue"] 

```

| Property                        | Description | Type                                      |
|---------------------------------|-------------|-------------------------------------------|
| [TypeReference](#TypeReference) |             |                                           |
| base                            |             | [TypeReference](#TypeReference)           |
| description                     |             | string[]                                  |
| pkg                             |             | [Package](#Package)                       |
| usage                           |             | [PropertyReference[]](#PropertyReference) |
| values                          |             | [EnumValue[]](#EnumValue)                 |

<a id="EnumValue"></a>EnumValue
-------------------------------

Used by: [Enum](#Enum).

```mermaid
classDiagram
class EnumValue["EnumValue"] {
    description string[]
    kind string
    name string
    value dst.BasicLit
}

```

| Property    | Description | Type         |
|-------------|-------------|--------------|
| description |             | string[]     |
| kind        |             | string       |
| name        |             | string       |
| value       |             | dst.BasicLit |

<a id="ImportReference"></a>ImportReference
-------------------------------------------

```mermaid
classDiagram
class ImportReference["ImportReference"] {
    Alias string
    ImportPath string
}

```

| Property   | Description | Type   |
|------------|-------------|--------|
| Alias      |             | string |
| ImportPath |             | string |

<a id="Markers"></a>Markers
---------------------------

Used by: [Markers](#Markers).

```mermaid
classDiagram
class Markers["Markers"] {
    name string
    value string
}
Markers -- Markers : children
class Markers["Markers"] 

```

| Property | Description | Type                           |
|----------|-------------|--------------------------------|
| children |             | [map[string]Markers](#Markers) |
| name     |             | string                         |
| value    |             | string                         |

<a id="MarkerValue"></a>MarkerValue
-----------------------------------

MarkerValue captures the value of a specific marker read from the source code.

Used by: [PackageMarkers](#PackageMarkers), and [PackageMarkers](#PackageMarkers).

```mermaid
classDiagram
class MarkerValue["MarkerValue"] {
    path string[]
    value string
}

```

| Property | Description | Type     |
|----------|-------------|----------|
| path     |             | string[] |
| value    |             | string   |

<a id="Object"></a>Object
-------------------------

```mermaid
classDiagram
class Object["Object"] {
    description string[]
}
Object *-- TypeReference
class TypeReference["TypeReference"]
Object -- Property : embeds
Object -- Package : pkg
Object -- Property : properties
Object -- PropertyReference : usage
class Property["Property"] 
class Package["Package"] 
class Property["Property"] 
class PropertyReference["PropertyReference"] 

```

| Property                        | Description | Type                                      |
|---------------------------------|-------------|-------------------------------------------|
| [TypeReference](#TypeReference) |             |                                           |
| description                     |             | string[]                                  |
| embeds                          |             | [Property[]](#Property)                   |
| pkg                             |             | [Package](#Package)                       |
| properties                      |             | [map[string]Property](#Property)          |
| usage                           |             | [PropertyReference[]](#PropertyReference) |

<a id="Order"></a>Order
-----------------------

<a id="Package"></a>Package
---------------------------

Package is a struct containing all of the declarations found in a package directory.

Used by: [Enum](#Enum), and [Object](#Object).

```mermaid
classDiagram
class Package["Package"] {
    cfg config.Config
    declarations map[string]Declaration
    log logr.Logger
    ranks map[string]int
}
Package -- PackageMarkers : metadata
class PackageMarkers["PackageMarkers"] 

```

| Property     | Description | Type                              |
|--------------|-------------|-----------------------------------|
| cfg          |             | config.Config                     |
| declarations |             | map[string]Declaration            |
| log          |             | logr.Logger                       |
| metadata     |             | [PackageMarkers](#PackageMarkers) |
| ranks        |             | map[string]int                    |

<a id="PackageMarkers"></a>PackageMarkers
-----------------------------------------

PackageMarkers captures specific package markers read from the source code.

Used by: [Package](#Package).

```mermaid
classDiagram
class PackageMarkers["PackageMarkers"] {
    DefaultGroup string
    DefaultVersion string
    Module string
    Name string
}
PackageMarkers -- MarkerValue : group
PackageMarkers -- MarkerSwitch : optional
PackageMarkers -- MarkerSwitch : required
PackageMarkers -- MarkerValue : version
class MarkerValue["MarkerValue"] 
class MarkerSwitch["MarkerSwitch"] 
class MarkerSwitch["MarkerSwitch"] 
class MarkerValue["MarkerValue"] 

```

| Property       | Description | Type                          |
|----------------|-------------|-------------------------------|
| DefaultGroup   |             | string                        |
| DefaultVersion |             | string                        |
| group          |             | [MarkerValue](#MarkerValue)   |
| Module         |             | string                        |
| Name           |             | string                        |
| optional       |             | [MarkerSwitch](#MarkerSwitch) |
| required       |             | [MarkerSwitch](#MarkerSwitch) |
| version        |             | [MarkerValue](#MarkerValue)   |

<a id="PropertyReference"></a>PropertyReference
-----------------------------------------------

Used by: [Enum](#Enum), and [Object](#Object).

```mermaid
classDiagram
class PropertyReference["PropertyReference"] {
    HostID string
    HostName string
    Property string
}

```

| Property | Description | Type   |
|----------|-------------|--------|
| HostID   |             | string |
| HostName |             | string |
| Property |             | string |

<a id="Resource"></a>Resource
-----------------------------

```mermaid
classDiagram
class Resource["Resource"] {
}
Resource *-- Object
class Object["Object"]
Resource -- Property : Spec
Resource -- Property : Status
class Property["Property"] 
class Property["Property"] 

```

| Property          | Description | Type                  |
|-------------------|-------------|-----------------------|
| [Object](#Object) |             |                       |
| Spec              |             | [Property](#Property) |
| Status            |             | [Property](#Property) |

### <a id="Property"></a>Property

| Property    | Description | Type                                |
|-------------|-------------|-------------------------------------|
| DeclaredOn  |             | PropertyContainer                   |
| description |             | string[]                            |
| Field       |             | string                              |
| markers     |             | [PropertyMarkers](#PropertyMarkers) |
| Name        |             | string                              |
| Type        |             | [TypeReference](#TypeReference)     |

### <a id="Property"></a>Property

| Property    | Description | Type                                |
|-------------|-------------|-------------------------------------|
| DeclaredOn  |             | PropertyContainer                   |
| description |             | string[]                            |
| Field       |             | string                              |
| markers     |             | [PropertyMarkers](#PropertyMarkers) |
| Name        |             | string                              |
| Type        |             | [TypeReference](#TypeReference)     |

<a id="Property"></a>Property
-----------------------------

Used by: [Object](#Object), [Object](#Object), [Resource](#Resource), and [Resource](#Resource).

```mermaid
classDiagram
class Property["Property"] {
    DeclaredOn PropertyContainer
    description string[]
    Field string
    Name string
}
Property -- PropertyMarkers : markers
Property -- TypeReference : Type
class PropertyMarkers["PropertyMarkers"] 
class TypeReference["TypeReference"] 

```

| Property    | Description | Type                                |
|-------------|-------------|-------------------------------------|
| DeclaredOn  |             | PropertyContainer                   |
| description |             | string[]                            |
| Field       |             | string                              |
| markers     |             | [PropertyMarkers](#PropertyMarkers) |
| Name        |             | string                              |
| Type        |             | [TypeReference](#TypeReference)     |

<a id="PropertyMarkers"></a>PropertyMarkers
-------------------------------------------

Used by: [Property](#Property).

```mermaid
classDiagram
class PropertyMarkers["PropertyMarkers"] {
}
PropertyMarkers -- MarkerSwitch : optional
PropertyMarkers -- MarkerSwitch : required
class MarkerSwitch["MarkerSwitch"] 
class MarkerSwitch["MarkerSwitch"] 

```

| Property | Description | Type                          |
|----------|-------------|-------------------------------|
| optional |             | [MarkerSwitch](#MarkerSwitch) |
| required |             | [MarkerSwitch](#MarkerSwitch) |

<a id="TypeReference"></a>TypeReference
---------------------------------------

Used by: [Enum](#Enum), and [Property](#Property).

```mermaid
classDiagram
class TypeReference["TypeReference"] {
    display string
    id string
    impPath string
    name string
    pkg string
}

```

| Property | Description | Type   |
|----------|-------------|--------|
| display  |             | string |
| id       |             | string |
| impPath  |             | string |
| name     |             | string |
| pkg      |             | string |

<a id="MarkerSwitch"></a>MarkerSwitch
-------------------------------------

MarkerSwitch captures the presence of a specific marker read from the source code.

Used by: [PackageMarkers](#PackageMarkers), [PackageMarkers](#PackageMarkers), [PropertyMarkers](#PropertyMarkers), and [PropertyMarkers](#PropertyMarkers).

```mermaid
classDiagram
class MarkerSwitch["MarkerSwitch"] {
    path string[]
    seen bool
}

```

| Property | Description | Type     |
|----------|-------------|----------|
| path     |             | string[] |
| seen     |             | bool     |
