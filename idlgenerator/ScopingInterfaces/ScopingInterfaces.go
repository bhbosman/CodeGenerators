package ScopingInterfaces

type ScopingContextFlags uint

type IScopingContext interface {
	Find(name string) (bool, IBaseDeclaredType)
	Add(name string, structType IBaseDeclaredType) error
	Replace(name string, structType IBaseDeclaredType) error
	Iterate(cb func(key string, value IBaseDeclaredType) error) error
	AddUnresolved(name string, information IFileInformation) error
	FindTypeSpec(fileInformation IFileInformation, s string) (IBaseDeclaredType, error)
	Previous() IScopingContext
	//IterateUnresolved(func(name string, information []IFileInformation) error) error
}

type IDefaultTypeService interface {
	Iterate(cb func(key string, declaredType IBaseDeclaredType))
	FindOk(s string) (dt IBaseDeclaredType, ok bool)
	Find(s string) (dt IBaseDeclaredType)
}
