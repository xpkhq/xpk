package xpkg

// PackageMeta defines the metadata
// structure of a package
type PackageMeta struct {
	edition int                  `yaml:"edition"`
	Name    string               `yaml:"name"`
	Author  []*PackageMetaAuthor `yaml:"author"`
}

// PackageMetaAuthor defines the
// package author structure
type PackageMetaAuthor struct {
	Name    string `yaml:"name"`
	Email   string `yaml:"email"`
	Website string `yaml:"website"`
}
