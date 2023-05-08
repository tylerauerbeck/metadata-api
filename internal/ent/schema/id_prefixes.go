package schema

const (
	// ApplicationPrefix is the prefix for all application IDs owned by metadata-api
	ApplicationPrefix string = "meta"
	// MetadataPrefix is the prefix for metadata
	MetadataPrefix string = ApplicationPrefix + "dat"
	// AnnotationPrefix is the prefix for all annotations
	AnnotationPrefix string = ApplicationPrefix + "ano"
	// AnnotationNamespacePrefix is the prefix for all annotation namespaces
	AnnotationNamespacePrefix string = ApplicationPrefix + "mns"
	// StatusPrefix is the prefix for all statuses
	StatusPrefix string = ApplicationPrefix + "sts"
	// StatusNamespacePrefix is the prefix for all status namespaces
	StatusNamespacePrefix string = ApplicationPrefix + "sns"
)
