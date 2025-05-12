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

Used by: [Package.enums](#Package), and [PackageBuilder.Enums](#PackageBuilder).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
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

Used by: [Enum.values](#Enum).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
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
---
  config:
    class:
      hideEmptyMembersBox: true
---
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

Used by: [Markers.children](#Markers).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
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

Used by: [PackageMarkers.group](#PackageMarkers), and [PackageMarkers.version](#PackageMarkers).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
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

Used by: [Package.objects](#Package), and [PackageBuilder.Objects](#PackageBuilder).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
classDiagram
class Object["Object"] {
    description string[]
    embeds PropertyList
}


Object *-- TypeReference
class TypeReference["TypeReference"]
Object -- Package : pkg
Object -- Property : properties
Object -- PropertyReference : usage
class Package["Package"] 
class Property["Property"] 
class PropertyReference["PropertyReference"] 
```

| Property                        | Description | Type                                      |
|---------------------------------|-------------|-------------------------------------------|
| [TypeReference](#TypeReference) |             |                                           |
| description                     |             | string[]                                  |
| embeds                          |             | PropertyList                              |
| pkg                             |             | [Package](#Package)                       |
| properties                      |             | [map[string]Property](#Property)          |
| usage                           |             | [PropertyReference[]](#PropertyReference) |

<a id="Order"></a>Order
-----------------------

<a id="Package"></a>Package
---------------------------

Package is a struct containing all of the declarations found in a package directory.

Used by: [Enum.pkg](#Enum), and [Object.pkg](#Object).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
classDiagram
class Package["Package"] {
    cfg config.Config
    log logr.Logger
    ranks map[string]int
}


Package -- Enum : enums
Package -- PackageMarkers : metadata
Package -- Object : objects
Package -- Resource : resources
class Enum["Enum"] 
class PackageMarkers["PackageMarkers"] 
class Object["Object"] 
class Resource["Resource"] 
```

| Property  | Description | Type                              |
|-----------|-------------|-----------------------------------|
| cfg       |             | config.Config                     |
| enums     |             | [map[string]Enum](#Enum)          |
| log       |             | logr.Logger                       |
| metadata  |             | [PackageMarkers](#PackageMarkers) |
| objects   |             | [map[string]Object](#Object)      |
| ranks     |             | map[string]int                    |
| resources |             | [map[string]Resource](#Resource)  |

<a id="PackageBuilder"></a>PackageBuilder
-----------------------------------------

PackageBuilder is a builder for Package instances.

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
classDiagram
class PackageBuilder["PackageBuilder"] {
    Config config.Config
    Log logr.Logger
}


PackageBuilder -- Enum : Enums
PackageBuilder -- PackageMarkers : Metadata
PackageBuilder -- Object : Objects
PackageBuilder -- Resource : Resources
class Enum["Enum"] 
class PackageMarkers["PackageMarkers"] 
class Object["Object"] 
class Resource["Resource"] 
```

| Property  | Description | Type                              |
|-----------|-------------|-----------------------------------|
| Config    |             | config.Config                     |
| Enums     |             | [Enum[]](#Enum)                   |
| Log       |             | logr.Logger                       |
| Metadata  |             | [PackageMarkers](#PackageMarkers) |
| Objects   |             | [Object[]](#Object)               |
| Resources |             | [Resource[]](#Resource)           |

<a id="PackageMarkers"></a>PackageMarkers
-----------------------------------------

PackageMarkers captures specific package markers read from the source code.

Used by: [Package.metadata](#Package), and [PackageBuilder.Metadata](#PackageBuilder).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
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

Used by: [Enum.usage](#Enum), and [Object.usage](#Object).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
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

Used by: [Package.resources](#Package), and [PackageBuilder.Resources](#PackageBuilder).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
classDiagram
class Resource["Resource"]

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

Used by: [Object.properties](#Object), [Resource.Spec](#Resource), and [Resource.Status](#Resource).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
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

Used by: [Property.markers](#Property).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
classDiagram
class PropertyMarkers["PropertyMarkers"]

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

Used by: [Enum.base](#Enum), and [Property.Type](#Property).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
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

Used by: [PackageMarkers.optional](#PackageMarkers), [PackageMarkers.required](#PackageMarkers), [PropertyMarkers.optional](#PropertyMarkers), and [PropertyMarkers.required](#PropertyMarkers).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
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
