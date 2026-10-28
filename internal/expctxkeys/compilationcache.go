package expctxkeys

// CompilationCache is a context.Context Value key. Its associated value is the
// wazero.CompilationCache that wazero.NewRuntime uses when it is given no
// RuntimeConfig.
type CompilationCache struct{}
