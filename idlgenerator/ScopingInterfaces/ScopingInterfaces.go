package ScopingInterfaces

type ScopingContextFlags uint

type IScopingContext interface {
	Find(name string) (bool, IDeclaredType)
	Add(name string, structType IDeclaredType) error
	Replace(name string, structType IDeclaredType) error
	Iterate(cb func(key string, value IDeclaredType) error) error
	AddUnresolved(name string, information IFileInformation) error
	FindTypeSpec(fileInformation IFileInformation, s string) (IDeclaredType, error)
	Previous() IScopingContext
	//IterateUnresolved(func(name string, information []IFileInformation) error) error
}

type IDefaultTypeService interface {
	Iterate(cb func(key string, declaredType IDeclaredType))
	FindOk(s string) (dt IDeclaredType, ok bool)
	Find(s string) (dt IDeclaredType)
}
