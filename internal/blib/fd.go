package blib

type FDI struct{ InitTargetDomain string } // "weja.us"

type FDC struct { // fd_origin metadata for future use
	FdBuild            bool   `required:"optional" fd_origin:"1"`                    // true
	FdClean            bool   `required:"optional" fd_origin:"1"`                    // false
	FdDebug            bool   `required:"optional" fd_origin:"1"`                    // false
	FdLocal            bool   `required:"optional" fd_origin:"1"`                    // true
	FdQuiet            bool   `required:"optional" fd_origin:"1"`                    // false
	FdRemote           bool   `required:"optional" fd_origin:"1"`                    // false
	FdTest             bool   `required:"optional" fd_origin:"1"`                    // false
	FdVerbose          bool   `required:"optional" fd_origin:"1"`                    // false
	FdAdc              string `required:"optional" fd_origin:"1" split_words:"true"` // "/.secrets/credentials.json",
	FdBuildContext     string `required:"optional" fd_origin:"1" split_words:"true"` // "ng"
	FdInit             string `required:"optional" fd_origin:"1" split_words:"true"` // "private-element"
	FdService          string `required:"optional" fd_origin:"1" split_words:"true"` // "private-element"
	FdRouteBase        string `required:"optional" fd_origin:"1" split_words:"true"` // "private-element"
	FdRepo             string `required:"optional" fd_origin:"1" split_words:"true"` // "micro-private-element",
	FdTitle            string `required:"optional" fd_origin:"1" split_words:"true"` // "Weja Too"
	FdTargetAlias      string `required:"optional" fd_origin:"1" split_words:"true"` // "wes"
	FdTargetDomain     string `required:"optional" fd_origin:"1" split_words:"true"` // "weja.us"
	FdTargetImageTag   string `required:"optional" fd_origin:"1" split_words:"true"` // "latest"
	FdTargetLocalPort  string `required:"optional" fd_origin:"1" split_words:"true"` // 4021
	FdTargetLogLevel   string `required:"optional" fd_origin:"1" split_words:"true"` // "error"
	FdTargetProjectId  string `required:"optional" fd_origin:"1" split_words:"true"` // "weja-us"
	FdTargetRealm      string `required:"optional" fd_origin:"1" split_words:"true"` // "too."
	FdTargetRemotePort string `required:"optional" fd_origin:"1" split_words:"true"` // 8080
}

type FDA struct {
	BuildPtr            *bool
	CleanPtr            *bool
	DebugPtr            *bool
	LocalPtr            *bool
	QuietPtr            *bool
	RemotePtr           *bool
	TestPtr             *bool
	VerbosePtr          *bool
	AdcPtr              *string
	BuildContextPtr     *string
	InitPtr             *string
	ServicePtr          *string
	RouteBasePtr        *string
	RepoPtr             *string
	TitlePtr            *string
	TargetAliasPtr      *string
	TargetDomainPtr     *string
	TargetImageTagPtr   *string
	TargetLocalPortPtr  *string
	TargetLogLevelPtr   *string
	TargetProjectIdPtr  *string
	TargetRealmPtr      *string
	TargetRemotePortPtr *string
}
