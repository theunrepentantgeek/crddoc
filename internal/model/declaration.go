package model

// Declaration is a common interface for declarations found in Go files
type Declaration interface {
	Name() string
	Id() string
	Kind() DeclarationType
	Usage() []PropertyReference
	SetUsage([]PropertyReference)
	Description() []string
}

type DeclarationType string

const (
	ResourceDeclaration = DeclarationType("Resource")
	ObjectDeclaration   = DeclarationType("Object")
	EnumDeclaration     = DeclarationType("Enum")
)
