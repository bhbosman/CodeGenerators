package corbaFiles

//Code generated by  DO NOT EDIT.
//line CORBA_DomainManager.idl:4
//Interface Decl: DomainManager
type DomainManager interface {
	//line CORBA_DomainManager.idl:5
	Get_domain_policy(policy_type PolicyType) (Policy, error)
}
//line CORBA_DomainManager.idl:12
//Interface Decl: ConstructionPolicy
type ConstructionPolicy interface {
	//line CORBA_DomainManager.idl:13
	Make_domain_manager(object_type InterfaceDef,constr_policy bool) (void, error)
}
