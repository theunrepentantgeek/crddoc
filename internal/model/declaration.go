package model

// Declaration is a common interface for declarations found in Go files
type Declaration interface {
	Kind() DeclarationType
	Name() string
	Usage() []PropertyReference
	SetUsage([]PropertyReference)
	Description() []string
}

type DeclarationType string

const (
	ObjectDeclaration = DeclarationType("Object")
	EnumDeclaration   = DeclarationType("Enum")
)
