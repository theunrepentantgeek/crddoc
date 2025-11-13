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

class Enum["Enum"]{  
      description string[]  
      AddValue(value EnumValue)  
      Description()string[]  
      Kind()DeclarationType  
      Name()string  
      Package()Package  
      SetPackage(pkg Package)  
      SetUsage(usage PropertyReference[])  
      Usage()PropertyReference[]  
      Values()EnumValue[]


} 

Package -- Enum : enums 
PackageBuilder -- Enum : Enums 

Enum -- TypeReference : base 
Enum -- Package : pkg 
Enum -- PropertyReference : usage 
Enum -- EnumValue : values 

class Package["Package"]
class PackageBuilder["PackageBuilder"]
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

### Functions

| Function    | Description | Parameters                | Returns             |
|-------------|-------------|---------------------------|---------------------|
| AddValue    |             | value EnumValue           |                     |
| Description |             |                           | string[]            |
| Kind        |             |                           | DeclarationType     |
| Name        |             |                           | string              |
| Package     |             |                           | Package             |
| SetPackage  |             | pkg Package               |                     |
| SetUsage    |             | usage PropertyReference[] |                     |
| Usage       |             |                           | PropertyReference[] |
| Values      |             |                           | EnumValue[]         |

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

class EnumValue["EnumValue"]{  
      description string[]  
      kind string  
      name string  
      value dst.BasicLit  
      Description()string[]  
      Kind()string  
      Name()string  
      Value()string


} 

Enum -- EnumValue : values 


class Enum["Enum"]

```

| Property    | Description | Type         |
|-------------|-------------|--------------|
| description |             | string[]     |
| kind        |             | string       |
| name        |             | string       |
| value       |             | dst.BasicLit |

### Functions

| Function    | Description | Parameters | Returns  |
|-------------|-------------|------------|----------|
| Description |             |            | string[] |
| Kind        |             |            | string   |
| Name        |             |            | string   |
| Value       |             |            | string   |

<a id="Function"></a>Function
-----------------------------

Function represents a method declared on a struct type.

Used by: [Object.functions](#Object).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
classDiagram

class Function["Function"]{  
      description string[]  
      IsPointerReceiver bool  
      Name string  
      DeclaredOn()Object  
      Description()string[]  
      setDeclaredOn(obj Object)


} 

Object -- Function : functions 

Function -- Object : declaredOn 
Function -- Parameter : Parameters 
Function -- TypeReference : Receiver 
Function -- Parameter : Results 

class Object["Object"]
class Object["Object"]
class Parameter["Parameter"]
class TypeReference["TypeReference"]
class Parameter["Parameter"]

```

| Property          | Description | Type                            |
|-------------------|-------------|---------------------------------|
| declaredOn        |             | [Object](#Object)               |
| description       |             | string[]                        |
| IsPointerReceiver |             | bool                            |
| Name              |             | string                          |
| Parameters        |             | [Parameter[]](#Parameter)       |
| Receiver          |             | [TypeReference](#TypeReference) |
| Results           |             | [Parameter[]](#Parameter)       |

### Functions

| Function      | Description                                                 | Parameters | Returns  |
|---------------|-------------------------------------------------------------|------------|----------|
| DeclaredOn    | DeclaredOn returns the object this function is declared on. |            | Object   |
| Description   | Description returns the description of the function.        |            | string[] |
| setDeclaredOn | setDeclaredOn sets the object this function is declared on. | obj Object |          |

<a id="ImportReference"></a>ImportReference
-------------------------------------------

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
classDiagram

class ImportReference["ImportReference"]{  
      Alias string  
      ImportPath string  
      Name()string


} 




```

| Property   | Description | Type   |
|------------|-------------|--------|
| Alias      |             | string |
| ImportPath |             | string |

### Functions

| Function | Description | Parameters | Returns |
|----------|-------------|------------|---------|
| Name     |             |            | string  |

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

class Markers["Markers"]{  
      name string  
      value string  
      Add(marker string)  
      Any()bool  
      Exists(path ...string)bool  
      Lookup(path ...string)Markers, bool  
      requireChild(name string)Markers  
      Value()string


} 

Markers -- Markers : children 

Markers -- Markers : children 

class Markers["Markers"]
class Markers["Markers"]

```

| Property | Description | Type                           |
|----------|-------------|--------------------------------|
| children |             | [map[string]Markers](#Markers) |
| name     |             | string                         |
| value    |             | string                         |

### Functions

| Function     | Description                                                | Parameters     | Returns       |
|--------------|------------------------------------------------------------|----------------|---------------|
| Add          | Add a marker to the list.                                  | marker string  |               |
| Any          |                                                            |                | bool          |
| Exists       | Exists returns true if the marker exists.                  | path ...string | bool          |
| Lookup       | Lookup a marker value by path, returning the final marker. | path ...string | Markers, bool |
| requireChild |                                                            | name string    | Markers       |
| Value        |                                                            |                | string        |

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

class MarkerValue["MarkerValue"]{  
      path string[]  
      value string  
      Merge(other MarkerValue)error  
      SetValue(value string)error  
      Update(markers Markers)error  
      Value()string, bool


} 

PackageMarkers -- MarkerValue : group 
PackageMarkers -- MarkerValue : version 


class PackageMarkers["PackageMarkers"]
class PackageMarkers["PackageMarkers"]

```

| Property | Description | Type     |
|----------|-------------|----------|
| path     |             | string[] |
| value    |             | string   |

### Functions

| Function | Description                                                                                                                                                              | Parameters        | Returns      |
|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------|--------------|
| Merge    | Merge combines the supplied MetadataValue with the current one, returning an error if the values differ.                                                                 | other MarkerValue | error        |
| SetValue |                                                                                                                                                                          | value string      | error        |
| Update   | Update reads the value from the passed set of markers, updating the value if found. If a new value is found that's different from the current value, we return an error. | markers Markers   | error        |
| Value    | Value returns the current value of the metadata, if known.                                                                                                               |                   | string, bool |

<a id="Object"></a>Object
-------------------------

Used by: [Function.declaredOn](#Function), [Package.objects](#Package), and [PackageBuilder.Objects](#PackageBuilder).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
classDiagram

class Object["Object"]{  
      description string[]  
      embeds PropertyList  
      AddFunction(fn Function)  
      Description()string[]  
      Embed(name string)Property, bool  
      Embeds()PropertyList  
      findEmbeddedStructs(structType dst.StructType)PropertyList  
      findProperties(structType dst.StructType)map[string]Property  
      Function(name string)Function, bool  
      Functions()FunctionList  
      Kind()DeclarationType  
      linkImports(importReferences ImportReferenceSet)  
      linkImportsToType(typeRef TypeReference, importReferences ImportReferenceSet)  
      Package()Package  
      Properties()PropertyList  
      Property(name string)Property, bool  
      SetPackage(p Package)  
      SetUsage(uses PropertyReference[])  
      Usage()PropertyReference[]


} 

Function -- Object : declaredOn 
Package -- Object : objects 
PackageBuilder -- Object : Objects 

Object -- Function : functions 
Object -- Package : pkg 
Object -- Property : properties 
Object -- PropertyReference : usage 

class Function["Function"]
class Package["Package"]
class PackageBuilder["PackageBuilder"]
class Function["Function"]
class Package["Package"]
class Property["Property"]
class PropertyReference["PropertyReference"]

```

| Property                        | Description | Type                                      |
|---------------------------------|-------------|-------------------------------------------|
| [TypeReference](#TypeReference) |             |                                           |
| description                     |             | string[]                                  |
| embeds                          |             | PropertyList                              |
| functions                       |             | [map[string]Function](#Function)          |
| pkg                             |             | [Package](#Package)                       |
| properties                      |             | [map[string]Property](#Property)          |
| usage                           |             | [PropertyReference[]](#PropertyReference) |

### Functions

| Function            | Description                                                                                | Parameters                                                     | Returns             |
|---------------------|--------------------------------------------------------------------------------------------|----------------------------------------------------------------|---------------------|
| AddFunction         | AddFunction adds a function to the object.                                                 | fn Function                                                    |                     |
| Description         |                                                                                            |                                                                | string[]            |
| Embed               | Embed returns the embed with the given name and true, or nil and false if not found.       | name string                                                    | Property, bool      |
| Embeds              | Embeds returns all of the embeds of the object, in alphabetical order.                     |                                                                | PropertyList        |
| findEmbeddedStructs |                                                                                            | structType dst.StructType                                      | PropertyList        |
| findProperties      |                                                                                            | structType dst.StructType                                      | map[string]Property |
| Function            | Function returns the function with the given name and true, or nil and false if not found. | name string                                                    | Function, bool      |
| Functions           | Functions returns all the functions/methods of the object, in alphabetical order.          |                                                                | FunctionList        |
| Kind                |                                                                                            |                                                                | DeclarationType     |
| linkImports         |                                                                                            | importReferences ImportReferenceSet                            |                     |
| linkImportsToType   | linkImportsToType links a single TypeReference to its import path if available.            | typeRef TypeReference,<br/>importReferences ImportReferenceSet |                     |
| Package             |                                                                                            |                                                                | Package             |
| Properties          | Properties returns all the properties of the object, in alphabetical order.                |                                                                | PropertyList        |
| Property            | Property returns the property with the given name and true, or nil and false if not found. | name string                                                    | Property, bool      |
| SetPackage          |                                                                                            | p Package                                                      |                     |
| SetUsage            |                                                                                            | uses PropertyReference[]                                       |                     |
| Usage               |                                                                                            |                                                                | PropertyReference[] |

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

class Package["Package"]{  
      cfg config.Config  
      log logr.Logger  
      ranks map[string]int  
      addPropertyReferences(hostID string, name string, container PropertyContainer, refs map[string]PropertyReference[])  
      alphabeticalObjectComparison(left Declaration, right Declaration)int  
      calculateRanks()  
      calculateRanksFromRoot(name string, rank int)  
      catalogCrossReferences()  
      Declaration(name string)Declaration, bool  
      Declarations(order Order)Declaration[]  
      Group()string  
      indexUsage()map[string]PropertyReference[]  
      Module()string  
      Name()string  
      Object(name string)Object, bool  
      PropertiesRequiredByDefault()string  
      Rank(name string)int  
      rankedObjectComparison(left Declaration, right Declaration)int  
      Version()string


} 

Enum -- Package : pkg 
Object -- Package : pkg 

Package -- Enum : enums 
Package -- PackageMarkers : metadata 
Package -- Object : objects 
Package -- Resource : resources 

class Enum["Enum"]
class Object["Object"]
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

### Functions

| Function                     | Description                                                                                                                                              | Parameters                                                                                               | Returns                        |
|------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------|--------------------------------|
| addPropertyReferences        |                                                                                                                                                          | hostID string,<br/>name string,<br/>container PropertyContainer,<br/>refs map[string]PropertyReference[] |                                |
| alphabeticalObjectComparison |                                                                                                                                                          | left Declaration,<br/>right Declaration                                                                  | int                            |
| calculateRanks               | calculateRanks calculates the ranks of all declarations in the package. The rank is the depth from the root resource, with resources having a rank of 0. |                                                                                                          |                                |
| calculateRanksFromRoot       |                                                                                                                                                          | name string,<br/>rank int                                                                                |                                |
| catalogCrossReferences       |                                                                                                                                                          |                                                                                                          |                                |
| Declaration                  | Declaration returns the declaration with the given name, if found.                                                                                       | name string                                                                                              | Declaration, bool              |
| Declarations                 |                                                                                                                                                          | order Order                                                                                              | Declaration[]                  |
| Group                        | Group returns the group of the package, if known.                                                                                                        |                                                                                                          | string                         |
| indexUsage                   |                                                                                                                                                          |                                                                                                          | map[string]PropertyReference[] |
| Module                       | Module returns the module of the package.                                                                                                                |                                                                                                          | string                         |
| Name                         |                                                                                                                                                          |                                                                                                          | string                         |
| Object                       | Object returns the object with the given name, if there is one.                                                                                          | name string                                                                                              | Object, bool                   |
| PropertiesRequiredByDefault  |                                                                                                                                                          |                                                                                                          | string                         |
| Rank                         | Rank returns the usage rank (depth from the root resource) of the given declaration.                                                                     | name string                                                                                              | int                            |
| rankedObjectComparison       |                                                                                                                                                          | left Declaration,<br/>right Declaration                                                                  | int                            |
| Version                      | Version returns the version of the package, if known.                                                                                                    |                                                                                                          | string                         |

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

class PackageBuilder["PackageBuilder"]{  
      Config config.Config  
      Log logr.Logger  
      AddEnums(enums ...Enum)  
      AddObjects(objects ...Object)  
      AddResources(resources ...Resource)  
      Build()Package


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
| Objects   |             | [map[string]Object](#Object)      |
| Resources |             | [Resource[]](#Resource)           |

### Functions

| Function     | Description                                   | Parameters            | Returns |
|--------------|-----------------------------------------------|-----------------------|---------|
| AddEnums     | AddEnums adds enums to the builder.           | enums ...Enum         |         |
| AddObjects   | AddObjects adds objects to the builder.       | objects ...Object     |         |
| AddResources | AddResources adds resources to the builder.   | resources ...Resource |         |
| Build        | Build creates a new Package from the builder. |                       | Package |

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

class PackageMarkers["PackageMarkers"]{  
      DefaultGroup string  
      DefaultVersion string  
      Module string  
      Name string  
      Group()string  
      Merge(other PackageMarkers)error  
      PropertiesRequiredByDefault()string  
      Update(markers Markers)error  
      Version()string


} 

Package -- PackageMarkers : metadata 
PackageBuilder -- PackageMarkers : Metadata 

PackageMarkers -- MarkerValue : group 
PackageMarkers -- MarkerSwitch : optional 
PackageMarkers -- MarkerSwitch : required 
PackageMarkers -- MarkerValue : version 

class Package["Package"]
class PackageBuilder["PackageBuilder"]
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

### Functions

| Function                    | Description                                                                                                                      | Parameters           | Returns |
|-----------------------------|----------------------------------------------------------------------------------------------------------------------------------|----------------------|---------|
| Group                       | Group returns the group of the package, using the configured controller-runtime marker if set, or the directory name if not.     |                      | string  |
| Merge                       |                                                                                                                                  | other PackageMarkers | error   |
| PropertiesRequiredByDefault |                                                                                                                                  |                      | string  |
| Update                      |                                                                                                                                  | markers Markers      | error   |
| Version                     | Version returns the version of the package, using the configured controller-runtime marker if set, or the directory name if not. |                      | string  |

<a id="Parameter"></a>Parameter
-------------------------------

Parameter represents a function parameter or return value.

Used by: [Function.Parameters](#Function), and [Function.Results](#Function).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
classDiagram

class Parameter["Parameter"]{  
      Name string


} 

Function -- Parameter : Parameters 
Function -- Parameter : Results 

Parameter -- TypeReference : Type 

class Function["Function"]
class Function["Function"]
class TypeReference["TypeReference"]

```

| Property | Description | Type                            |
|----------|-------------|---------------------------------|
| Name     |             | string                          |
| Type     |             | [TypeReference](#TypeReference) |

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

class PropertyReference["PropertyReference"]{  
      HostID string  
      HostName string  
      Property string


} 

Enum -- PropertyReference : usage 
Object -- PropertyReference : usage 


class Enum["Enum"]
class Object["Object"]

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
Package -- Resource : resources 
PackageBuilder -- Resource : Resources 

Resource -- Property : Spec 
Resource -- Property : Status 

class Package["Package"]
class PackageBuilder["PackageBuilder"]
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

class Property["Property"]{  
      DeclaredOn PropertyContainer  
      description string[]  
      Field string  
      Name string  
      Description()string[]  
      Required()string  
      setContainer(container PropertyContainer)  
      tryParseName(field dst.Field)string, bool  
      tryParseNameFromTag(tag string, tagStruct reflect.StructTag)string, bool


} 

Object -- Property : properties 
Resource -- Property : Spec 
Resource -- Property : Status 

Property -- PropertyMarkers : markers 
Property -- TypeReference : Type 

class Object["Object"]
class Resource["Resource"]
class Resource["Resource"]
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

### Functions

| Function            | Description | Parameters                                  | Returns      |
|---------------------|-------------|---------------------------------------------|--------------|
| Description         |             |                                             | string[]     |
| Required            |             |                                             | string       |
| setContainer        |             | container PropertyContainer                 |              |
| tryParseName        |             | field dst.Field                             | string, bool |
| tryParseNameFromTag |             | tag string,<br/>tagStruct reflect.StructTag | string, bool |

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

class PropertyMarkers["PropertyMarkers"]{  
      Merge(other PropertyMarkers)error  
      Optional()bool  
      Parse(markers Markers)error  
      ParseDecorations(decs dst.NodeDecs)error  
      Required()bool


} 

Property -- PropertyMarkers : markers 

PropertyMarkers -- MarkerSwitch : optional 
PropertyMarkers -- MarkerSwitch : required 

class Property["Property"]
class MarkerSwitch["MarkerSwitch"]
class MarkerSwitch["MarkerSwitch"]

```

| Property | Description | Type                          |
|----------|-------------|-------------------------------|
| optional |             | [MarkerSwitch](#MarkerSwitch) |
| required |             | [MarkerSwitch](#MarkerSwitch) |

### Functions

| Function         | Description | Parameters            | Returns |
|------------------|-------------|-----------------------|---------|
| Merge            |             | other PropertyMarkers | error   |
| Optional         |             |                       | bool    |
| Parse            |             | markers Markers       | error   |
| ParseDecorations |             | decs dst.NodeDecs     | error   |
| Required         |             |                       | bool    |

<a id="TypeReference"></a>TypeReference
---------------------------------------

Used by: [Enum.base](#Enum), [Function.Receiver](#Function), [Parameter.Type](#Parameter), and [Property.Type](#Property).

```mermaid
---
  config:
    class:
      hideEmptyMembersBox: true
---
classDiagram

class TypeReference["TypeReference"]{  
      display string  
      id string  
      impPath string  
      name string  
      pkg string  
      Display()string  
      ID()string  
      ImportPath()string  
      Name()string  
      Package()string


} 

Enum -- TypeReference : base 
Function -- TypeReference : Receiver 
Parameter -- TypeReference : Type 
Property -- TypeReference : Type 


class Enum["Enum"]
class Function["Function"]
class Parameter["Parameter"]
class Property["Property"]

```

| Property | Description | Type   |
|----------|-------------|--------|
| display  |             | string |
| id       |             | string |
| impPath  |             | string |
| name     |             | string |
| pkg      |             | string |

### Functions

| Function   | Description | Parameters | Returns |
|------------|-------------|------------|---------|
| Display    |             |            | string  |
| ID         |             |            | string  |
| ImportPath |             |            | string  |
| Name       |             |            | string  |
| Package    |             |            | string  |

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

class MarkerSwitch["MarkerSwitch"]{  
      path string[]  
      seen bool  
      Merge(other MarkerSwitch)  
      Seen()bool  
      Update(markers Markers)


} 

PackageMarkers -- MarkerSwitch : optional 
PackageMarkers -- MarkerSwitch : required 
PropertyMarkers -- MarkerSwitch : optional 
PropertyMarkers -- MarkerSwitch : required 


class PackageMarkers["PackageMarkers"]
class PackageMarkers["PackageMarkers"]
class PropertyMarkers["PropertyMarkers"]
class PropertyMarkers["PropertyMarkers"]

```

| Property | Description | Type     |
|----------|-------------|----------|
| path     |             | string[] |
| seen     |             | bool     |

### Functions

| Function | Description                                                                                               | Parameters         | Returns |
|----------|-----------------------------------------------------------------------------------------------------------|--------------------|---------|
| Merge    | Merge combines the supplied MetadataSwitch with the current one, returning an error if the values differ. | other MarkerSwitch |         |
| Seen     | Seen returns true if this marker has been seen, false otherwise.                                          |                    | bool    |
| Update   | Update checks for the desired marker in passed set of markers, updating the value if found.               | markers Markers    |         |
