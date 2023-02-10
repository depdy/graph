package github

type Xaction action

type action struct {
	// The name of the action's author.
	Author *string `json:"author,omitempty"`
	// You can use a color and Feather icon to create a badge to personalize and distinguish
	// your action. Badges are shown next to your action name in GitHub Marketplace.
	Branding *Branding `json:"branding,omitempty"`
	// A short description of the action.
	Description string `json:"description"`
	// Input parameters allow you to specify data that the action expects to use during runtime.
	// GitHub stores input parameters as environment variables. Input ids with uppercase letters
	// are converted to lowercase during runtime. We recommended using lowercase input ids.
	Inputs *Inputs `json:"inputs,omitempty"`
	// The name of your action. GitHub displays the `name` in the Actions tab to help visually
	// identify actions in each job.
	Name    string      `json:"name"`
	Outputs interface{} `json:"outputs"`
	Runs    Runs        `json:"runs"`
}

// You can use a color and Feather icon to create a badge to personalize and distinguish
// your action. Badges are shown next to your action name in GitHub Marketplace.
type Branding struct {
	// The background color of the badge.
	Color *Color `json:"color,omitempty"`
	// The name of the Feather icon to use.
	Icon *Icon `json:"icon,omitempty"`
}

// Input parameters allow you to specify data that the action expects to use during runtime.
// GitHub stores input parameters as environment variables. Input ids with uppercase letters
// are converted to lowercase during runtime. We recommended using lowercase input ids.
type Inputs struct {
}

// Configures the path to the action's code and the application used to execute the code.
//
// Configures the path to the composite action, and the application used to execute the
// code.
//
// Configures the image used for the Docker action.
type Runs struct {
	// The file that contains your action code. The application specified in `using` executes
	// this file.
	Main *string `json:"main,omitempty"`
	// Allows you to run a script at the end of a job, once the `main:` action has completed.
	// For example, you can use `post:` to terminate certain processes or remove unneeded files.
	// The application specified with the `using` syntax will execute this file. The `post:`
	// action always runs by default but you can override this using `post-if`.
	Post   *string `json:"post,omitempty"`
	PostIf *string `json:"post-if,omitempty"`
	// Allows you to run a script at the start of a job, before the `main:` action begins. For
	// example, you can use `pre:` to run a prerequisite setup script. The application specified
	// with the `using` syntax will execute this file. The `pre:` action always runs by default
	// but you can override this using `pre-if`.
	Pre   *string `json:"pre,omitempty"`
	PreIf *string `json:"pre-if,omitempty"`
	// The application used to execute the code specified in `main`.
	//
	// To use a composite run steps action, set this to 'composite'.
	//
	// You must set this value to 'docker'.
	Using interface{} `json:"using"`
	// The run steps that you plan to run in this action.
	Steps []Step `json:"steps,omitempty"`
	// An array of strings that define the inputs for a Docker container. Inputs can include
	// hardcoded strings. GitHub passes the `args` to the container's `ENTRYPOINT` when the
	// container starts up.
	// The `args` are used in place of the `CMD` instruction in a `Dockerfile`. If you use `CMD`
	// in your `Dockerfile`, use the guidelines ordered by preference:
	// - Document required arguments in the action's README and omit them from the `CMD`
	// instruction.
	// - Use defaults that allow using the action without specifying any `args`.
	// - If the action exposes a `--help` flag, or something similar, use that to make your
	// action self-documenting.
	Args []string `json:"args,omitempty"`
	// Overrides the Docker `ENTRYPOINT` in the `Dockerfile`, or sets it if one wasn't already
	// specified. Use `entrypoint` when the `Dockerfile` does not specify an `ENTRYPOINT` or you
	// want to override the `ENTRYPOINT` instruction. If you omit `entrypoint`, the commands you
	// specify in the Docker `ENTRYPOINT` instruction will execute. The Docker `ENTRYPOINT
	// instruction has a *shell* form and *exec* form. The Docker `ENTRYPOINT` documentation
	// recommends using the *exec* form of the `ENTRYPOINT` instruction.
	Entrypoint *string `json:"entrypoint,omitempty"`
	// Specifies a key/value map of environment variables to set in the container environment.
	Env map[string]string `json:"env,omitempty"`
	// The Docker image to use as the container to run the action. The value can be the Docker
	// base image name, a local `Dockerfile` in your repository, or a public image in Docker Hub
	// or another registry. To reference a `Dockerfile` local to your repository, use a path
	// relative to your action metadata file. The `docker` application will execute this file.
	Image *string `json:"image,omitempty"`
	// Allows you to run a cleanup script once the `runs.entrypoint` action has completed.
	// GitHub Actions uses `docker run` to launch this action. Because GitHub Actions runs the
	// script inside a new container using the same base image, the runtime state is different
	// from the main `entrypoint` container. You can access any state you need in either the
	// workspace, `HOME`, or as a `STATE_` variable. The `post-entrypoint:` action always runs
	// by default but you can override this using `post-if`.
	PostEntrypoint *string `json:"post-entrypoint,omitempty"`
	// Allows you to run a script before the `entrypoint` action begins. For example, you can
	// use `pre-entrypoint:` to run a prerequisite setup script. GitHub Actions uses `docker
	// run` to launch this action, and runs the script inside a new container that uses the same
	// base image. This means that the runtime state is different from the main `entrypoint`
	// container, and any states you require must be accessed in either the workspace, `HOME`,
	// or as a `STATE_` variable. The `pre-entrypoint:` action always runs by default but you
	// can override this using `pre-if`.
	PreEntrypoint *string `json:"pre-entrypoint,omitempty"`
}

type Step struct {
	// Prevents a job from failing when a step fails. Set to true to allow a job to pass when
	// this step fails.
	ContinueOnError *ContinueOnError `json:"continue-on-error"`
	// Sets a map of environment variables for only that step.
	Env map[string]string `json:"env,omitempty"`
	// A unique identifier for the step. You can use the `id` to reference the step in contexts.
	ID *string `json:"id,omitempty"`
	// You can use the if conditional to prevent a step from running unless a condition is met.
	// You can use any supported context and expression to create a conditional.
	// Expressions in an if conditional do not require the ${{ }} syntax. For more information,
	// see https://help.github.com/en/articles/contexts-and-expression-syntax-for-github-actions.
	If *string `json:"if,omitempty"`
	// The name of the composite run step.
	Name *string `json:"name,omitempty"`
	// The command you want to run. This can be inline or a script in your action repository.
	Run *string `json:"run,omitempty"`
	// The shell where you want to run the command.
	Shell *string `json:"shell,omitempty"`
	// Selects an action to run as part of a step in your job.
	Uses *string `json:"uses,omitempty"`
	// A map of the input parameters defined by the action. Each input parameter is a key/value
	// pair. Input parameters are set as environment variables. The variable is prefixed with
	// INPUT_ and converted to upper case.
	With map[string]interface{} `json:"with,omitempty"`
	// Specifies the working directory where the command is run.
	WorkingDirectory *string `json:"working-directory,omitempty"`
}

// The background color of the badge.
type Color string

const (
	Blue     Color = "blue"
	GrayDark Color = "gray-dark"
	Green    Color = "green"
	Orange   Color = "orange"
	Purple   Color = "purple"
	Red      Color = "red"
	White    Color = "white"
	Yellow   Color = "yellow"
)

// The name of the Feather icon to use.
type Icon string

const (
	Activity         Icon = "activity"
	Airplay          Icon = "airplay"
	AlertCircle      Icon = "alert-circle"
	AlertOctagon     Icon = "alert-octagon"
	AlertTriangle    Icon = "alert-triangle"
	AlignCenter      Icon = "align-center"
	AlignJustify     Icon = "align-justify"
	AlignLeft        Icon = "align-left"
	AlignRight       Icon = "align-right"
	Anchor           Icon = "anchor"
	Aperture         Icon = "aperture"
	Archive          Icon = "archive"
	ArrowDown        Icon = "arrow-down"
	ArrowDownCircle  Icon = "arrow-down-circle"
	ArrowDownLeft    Icon = "arrow-down-left"
	ArrowDownRight   Icon = "arrow-down-right"
	ArrowLeft        Icon = "arrow-left"
	ArrowLeftCircle  Icon = "arrow-left-circle"
	ArrowRight       Icon = "arrow-right"
	ArrowRightCircle Icon = "arrow-right-circle"
	ArrowUp          Icon = "arrow-up"
	ArrowUpCircle    Icon = "arrow-up-circle"
	ArrowUpLeft      Icon = "arrow-up-left"
	ArrowUpRight     Icon = "arrow-up-right"
	AtSign           Icon = "at-sign"
	Award            Icon = "award"
	BarChart         Icon = "bar-chart"
	BarChart2        Icon = "bar-chart-2"
	Battery          Icon = "battery"
	BatteryCharging  Icon = "battery-charging"
	Bell             Icon = "bell"
	BellOff          Icon = "bell-off"
	Bluetooth        Icon = "bluetooth"
	Bold             Icon = "bold"
	Book             Icon = "book"
	BookOpen         Icon = "book-open"
	Bookmark         Icon = "bookmark"
	Box              Icon = "box"
	Briefcase        Icon = "briefcase"
	CPU              Icon = "cpu"
	Calendar         Icon = "calendar"
	Camera           Icon = "camera"
	CameraOff        Icon = "camera-off"
	Cast             Icon = "cast"
	Check            Icon = "check"
	CheckCircle      Icon = "check-circle"
	CheckSquare      Icon = "check-square"
	ChevronDown      Icon = "chevron-down"
	ChevronLeft      Icon = "chevron-left"
	ChevronRight     Icon = "chevron-right"
	ChevronUp        Icon = "chevron-up"
	ChevronsDown     Icon = "chevrons-down"
	ChevronsLeft     Icon = "chevrons-left"
	ChevronsRight    Icon = "chevrons-right"
	ChevronsUp       Icon = "chevrons-up"
	Circle           Icon = "circle"
	Clipboard        Icon = "clipboard"
	Clock            Icon = "clock"
	Cloud            Icon = "cloud"
	CloudDrizzle     Icon = "cloud-drizzle"
	CloudLightning   Icon = "cloud-lightning"
	CloudOff         Icon = "cloud-off"
	CloudRain        Icon = "cloud-rain"
	CloudSnow        Icon = "cloud-snow"
	Code             Icon = "code"
	Command          Icon = "command"
	Compass          Icon = "compass"
	Copy             Icon = "copy"
	CornerDownLeft   Icon = "corner-down-left"
	CornerDownRight  Icon = "corner-down-right"
	CornerLeftDown   Icon = "corner-left-down"
	CornerLeftUp     Icon = "corner-left-up"
	CornerRightDown  Icon = "corner-right-down"
	CornerRightUp    Icon = "corner-right-up"
	CornerUpLeft     Icon = "corner-up-left"
	CornerUpRight    Icon = "corner-up-right"
	CreditCard       Icon = "credit-card"
	Crop             Icon = "crop"
	Crosshair        Icon = "crosshair"
	Database         Icon = "database"
	// Delete           Icon = "delete"
	Disc           Icon = "disc"
	DollarSign     Icon = "dollar-sign"
	Download       Icon = "download"
	DownloadCloud  Icon = "download-cloud"
	Droplet        Icon = "droplet"
	Edit           Icon = "edit"
	Edit2          Icon = "edit-2"
	Edit3          Icon = "edit-3"
	ExternalLink   Icon = "external-link"
	Eye            Icon = "eye"
	EyeOff         Icon = "eye-off"
	Facebook       Icon = "facebook"
	FastForward    Icon = "fast-forward"
	Feather        Icon = "feather"
	File           Icon = "file"
	FileMinus      Icon = "file-minus"
	FilePlus       Icon = "file-plus"
	FileText       Icon = "file-text"
	Film           Icon = "film"
	Filter         Icon = "filter"
	Flag           Icon = "flag"
	Folder         Icon = "folder"
	FolderMinus    Icon = "folder-minus"
	FolderPlus     Icon = "folder-plus"
	Gift           Icon = "gift"
	GitBranch      Icon = "git-branch"
	GitCommit      Icon = "git-commit"
	GitMerge       Icon = "git-merge"
	GitPullRequest Icon = "git-pull-request"
	Globe          Icon = "globe"
	Grid           Icon = "grid"
	HardDrive      Icon = "hard-drive"
	Hash           Icon = "hash"
	Headphones     Icon = "headphones"
	Heart          Icon = "heart"
	HelpCircle     Icon = "help-circle"
	Home           Icon = "home"
	Image          Icon = "image"
	Inbox          Icon = "inbox"
	Info           Icon = "info"
	Italic         Icon = "italic"
	Layers         Icon = "layers"
	Layout         Icon = "layout"
	LifeBuoy       Icon = "life-buoy"
	Link           Icon = "link"
	Link2          Icon = "link-2"
	List           Icon = "list"
	Loader         Icon = "loader"
	Lock           Icon = "lock"
	LogIn          Icon = "log-in"
	LogOut         Icon = "log-out"
	Mail           Icon = "mail"
	Map            Icon = "map"
	MapPin         Icon = "map-pin"
	Maximize       Icon = "maximize"
	Maximize2      Icon = "maximize-2"
	Menu           Icon = "menu"
	MessageCircle  Icon = "message-circle"
	MessageSquare  Icon = "message-square"
	Mic            Icon = "mic"
	MicOff         Icon = "mic-off"
	Minimize       Icon = "minimize"
	Minimize2      Icon = "minimize-2"
	Minus          Icon = "minus"
	MinusCircle    Icon = "minus-circle"
	MinusSquare    Icon = "minus-square"
	Monitor        Icon = "monitor"
	Moon           Icon = "moon"
	MoreHorizontal Icon = "more-horizontal"
	MoreVertical   Icon = "more-vertical"
	Move           Icon = "move"
	Music          Icon = "music"
	Navigation     Icon = "navigation"
	Navigation2    Icon = "navigation-2"
	Octagon        Icon = "octagon"
	Package        Icon = "package"
	Paperclip      Icon = "paperclip"
	Pause          Icon = "pause"
	PauseCircle    Icon = "pause-circle"
	Percent        Icon = "percent"
	Phone          Icon = "phone"
	PhoneCall      Icon = "phone-call"
	PhoneForwarded Icon = "phone-forwarded"
	PhoneIncoming  Icon = "phone-incoming"
	PhoneMissed    Icon = "phone-missed"
	PhoneOff       Icon = "phone-off"
	PhoneOutgoing  Icon = "phone-outgoing"
	PieChart       Icon = "pie-chart"
	Play           Icon = "play"
	PlayCircle     Icon = "play-circle"
	Plus           Icon = "plus"
	PlusCircle     Icon = "plus-circle"
	PlusSquare     Icon = "plus-square"
	Pocket         Icon = "pocket"
	Power          Icon = "power"
	Printer        Icon = "printer"
	RSS            Icon = "rss"
	Radio          Icon = "radio"
	RefreshCcw     Icon = "refresh-ccw"
	RefreshCw      Icon = "refresh-cw"
	Repeat         Icon = "repeat"
	Rewind         Icon = "rewind"
	RotateCcw      Icon = "rotate-ccw"
	RotateCw       Icon = "rotate-cw"
	Save           Icon = "save"
	Scissors       Icon = "scissors"
	Search         Icon = "search"
	Send           Icon = "send"
	Server         Icon = "server"
	Settings       Icon = "settings"
	Share          Icon = "share"
	Share2         Icon = "share-2"
	Shield         Icon = "shield"
	ShieldOff      Icon = "shield-off"
	ShoppingBag    Icon = "shopping-bag"
	ShoppingCart   Icon = "shopping-cart"
	Shuffle        Icon = "shuffle"
	Sidebar        Icon = "sidebar"
	SkipBack       Icon = "skip-back"
	SkipForward    Icon = "skip-forward"
	Slash          Icon = "slash"
	Sliders        Icon = "sliders"
	Smartphone     Icon = "smartphone"
	Speaker        Icon = "speaker"
	Square         Icon = "square"
	Star           Icon = "star"
	StopCircle     Icon = "stop-circle"
	Sun            Icon = "sun"
	Sunrise        Icon = "sunrise"
	Sunset         Icon = "sunset"
	Tablet         Icon = "tablet"
	Tag            Icon = "tag"
	Target         Icon = "target"
	Terminal       Icon = "terminal"
	Thermometer    Icon = "thermometer"
	ThumbsDown     Icon = "thumbs-down"
	ThumbsUp       Icon = "thumbs-up"
	ToggleLeft     Icon = "toggle-left"
	ToggleRight    Icon = "toggle-right"
	Trash          Icon = "trash"
	Trash2         Icon = "trash-2"
	TrendingDown   Icon = "trending-down"
	TrendingUp     Icon = "trending-up"
	Triangle       Icon = "triangle"
	Truck          Icon = "truck"
	Tv             Icon = "tv"
	Type           Icon = "type"
	Umbrella       Icon = "umbrella"
	Underline      Icon = "underline"
	Unlock         Icon = "unlock"
	Upload         Icon = "upload"
	UploadCloud    Icon = "upload-cloud"
	User           Icon = "user"
	UserCheck      Icon = "user-check"
	UserMinus      Icon = "user-minus"
	UserPlus       Icon = "user-plus"
	UserX          Icon = "user-x"
	Users          Icon = "users"
	Video          Icon = "video"
	VideoOff       Icon = "video-off"
	Voicemail      Icon = "voicemail"
	Volume         Icon = "volume"
	Volume1        Icon = "volume-1"
	Volume2        Icon = "volume-2"
	VolumeX        Icon = "volume-x"
	// Watch            Icon = "watch"
	Wifi    Icon = "wifi"
	WifiOff Icon = "wifi-off"
	Wind    Icon = "wind"
	X       Icon = "x"
	XCircle Icon = "x-circle"
	XSquare Icon = "x-square"
	Zap     Icon = "zap"
	ZapOff  Icon = "zap-off"
	ZoomIn  Icon = "zoom-in"
	ZoomOut Icon = "zoom-out"
)

// Prevents a job from failing when a step fails. Set to true to allow a job to pass when
// this step fails.
type ContinueOnError struct {
	Bool   *bool
	String *string
}
