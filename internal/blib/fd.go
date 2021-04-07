

package blib

type FDI struct	{ InitTargetDomain	string } 								// "weja.us"

type FDC struct	{
	FdDebug				bool	`required:"optional"`						// false
	FdLocal				bool	`required:"optional"`						// true
	FdQuiet				bool	`required:"optional"`						// false
	FdRemote			bool	`required:"optional"`						// false
	FdVerbose			bool	`required:"optional"`						// 8080
	FdBuildContext		string	`required:"optional" split_words:"true"`	// "ng"
	FdNickname			string	`required:"optional" split_words:"true"`	// "private-element"
	FdServiceName		string	`required:"optional" split_words:"true"`	// "micro-private-element",
	FdTargetAlias		string	`required:"optional" split_words:"true"`	// "wes"
	FdTargetDomain		string	`required:"optional" split_words:"true"`	// "weja.us"
	FdTargetImageTag	string	`required:"optional" split_words:"true"`	// "latest"
	FdTargetLocalPort	int64	`required:"optional" split_words:"true"`	// 4021
	FdTargetLogLevel	string	`required:"optional" split_words:"true"`	// "error"
	FdTargetProjectId	string	`required:"optional" split_words:"true"`	// "weja-us"
	FdTargetRealm		string	`required:"optional" split_words:"true"`	// "too.fb."
	FdTargetRemotePort	int64	`required:"optional" split_words:"true"`	// 8080
	Success				bool	`required:"optional"`
}

type FDA struct {
	DebugPtr			*bool
	LocalPtr			*bool
	QuietPtr			*bool
	RemotePtr			*bool
	VerbosePtr			*bool
	BuildContextPtr		*string
	NicknamePtr			*string
	ServiceNamePtr		*string
	TargetAliasPtr		*string
	TargetDomainPtr		*string
	TargetImageTagPtr	*string
	TargetLocalPortPtr	*int64
	TargetLogLevelPtr	*string
	TargetProjectIdPtr	*string
	TargetRealmPtr		*string
	TargetRemotePortPtr	*int64
}
