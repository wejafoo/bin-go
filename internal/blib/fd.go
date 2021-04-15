package blib

type FDI struct{ InitTargetDomain string }											 // "weja.us"

type FDC struct {
	FdBuild            bool   `required:"optional" fd_origin:"1"`					 // true
	FdDebug            bool   `required:"optional" fd_origin:"1"`					 // false
	FdLocal            bool   `required:"optional" fd_origin:"1"`					 // true
	FdQuiet            bool   `required:"optional" fd_origin:"1"`					 // false
	FdRemote           bool   `required:"optional" fd_origin:"1"`					 // false
	FdVerbose          bool   `required:"optional" fd_origin:"1"`					 // 8080
	FdBuildContext     string `required:"optional" fd_origin:"1" split_words:"true"` // "ng"
	FdNickname         string `required:"optional" fd_origin:"1" split_words:"true"` // "private-element"
	FdServiceName      string `required:"optional" fd_origin:"1" split_words:"true"` // "micro-private-element",
	FdSiteNickname     string `required:"optional" fd_origin:"1" split_words:"true"` // "Weja Too"
	FdTargetAlias      string `required:"optional" fd_origin:"1" split_words:"true"` // "wes"
	FdTargetDomain     string `required:"optional" fd_origin:"1" split_words:"true"` // "weja.us"
	FdTargetImageTag   string `required:"optional" fd_origin:"1" split_words:"true"` // "latest"
	FdTargetLocalPort  string `required:"optional" fd_origin:"1" split_words:"true"` // 4021
	FdTargetLogLevel   string `required:"optional" fd_origin:"1" split_words:"true"` // "error"
	FdTargetProjectId  string `required:"optional" fd_origin:"1" split_words:"true"` // "weja-us"
	FdTargetRealm      string `required:"optional" fd_origin:"1" split_words:"true"` // "too.fb."
	FdTargetRemotePort string `required:"optional" fd_origin:"1" split_words:"true"` // 8080
	Success            bool   `required:"optional"`
}

type FDA struct {
	BuildPtr            *bool
	DebugPtr            *bool
	LocalPtr            *bool
	QuietPtr            *bool
	RemotePtr           *bool
	VerbosePtr          *bool
	BuildContextPtr     *string
	NicknamePtr         *string
	ServiceNamePtr      *string
	SiteNicknamePtr     *string
	TargetAliasPtr      *string
	TargetDomainPtr     *string
	TargetImageTagPtr   *string
	TargetLocalPortPtr  *string
	TargetLogLevelPtr   *string
	TargetProjectIdPtr  *string
	TargetRealmPtr      *string
	TargetRemotePortPtr *string
}
