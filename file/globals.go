package file

const (
	MetadataFilename = ".barber.yaml"
	ConfigFilename   = ".barber.toml"
)

// Represents a file purpose (ie. type) in relation to the application logic.
// This is not the file extension, but rather the purpose of the file.
type purpose string

const (
	purposeMetadata purpose = "metadata"
	purposeConfig   purpose = "config"
)

// Represents a file operation.
// It is a verb that describes what is being done to the file.
type Operation string

const (
	OperationDelete Operation = "delete"
	OperationReset  Operation = "reset"
)

type context string

const (
	contextCollection context = "collection"
	contextApp        context = "app"
)
