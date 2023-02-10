package github

type workflow struct {
	// Concurrency ensures that only a single job or workflow using the same concurrency group
	// will run at a time. A concurrency group can be any string or expression. The expression
	// can use any context except for the secrets context.
	// You can also specify concurrency at the workflow level.
	// When a concurrent job or workflow is queued, if another job or workflow using the same
	// concurrency group in the repository is in progress, the queued job or workflow will be
	// pending. Any previously pending job or workflow in the concurrency group will be
	// canceled. To also cancel any currently running job or workflow in the same concurrency
	// group, specify cancel-in-progress: true.
	Concurrency *ConcurrencyUnion `json:"concurrency"`
	// A map of default settings that will apply to all jobs in the workflow.
	Defaults *Defaults `json:"defaults,omitempty"`
	// A map of environment variables that are available to all jobs and steps in the workflow.
	Env *Env `json:"env"`
	// A workflow run is made up of one or more jobs. Jobs run in parallel by default. To run
	// jobs sequentially, you can define dependencies on other jobs using the
	// jobs.<job_id>.needs keyword.
	// Each job runs in a fresh instance of the virtual environment specified by runs-on.
	// You can run an unlimited number of jobs as long as you are within the workflow usage
	// limits. For more information, see
	// https://help.github.com/en/github/automating-your-workflow-with-github-actions/workflow-syntax-for-github-actions#usage-limits.
	Jobs Jobs `json:"jobs"`
	// The name of your workflow. GitHub displays the names of your workflows on your
	// repository's actions page. If you omit this field, GitHub sets the name to the workflow's
	// filename.
	Name *string `json:"name,omitempty"`
	// The name of the GitHub event that triggers the workflow. You can provide a single event
	// string, array of events, array of event types, or an event configuration map that
	// schedules a workflow or restricts the execution of a workflow to specific files, tags, or
	// branch changes. For a list of available events, see
	// https://help.github.com/en/github/automating-your-workflow-with-github-actions/events-that-trigger-workflows.
	On          *OnUnion    `json:"on"`
	Permissions interface{} `json:"permissions"`
	// The name for workflow runs generated from the workflow. GitHub displays the workflow run
	// name in the list of workflow runs on your repository's 'Actions' tab.
	RunName *string `json:"run-name,omitempty"`
}

type Concurrency struct {
	// To cancel any currently running job or workflow in the same concurrency group, specify
	// cancel-in-progress: true.
	CancelInProgress *CancelInProgress `json:"cancel-in-progress"`
	// When a concurrent job or workflow is queued, if another job or workflow using the same
	// concurrency group in the repository is in progress, the queued job or workflow will be
	// pending. Any previously pending job or workflow in the concurrency group will be canceled.
	Group string `json:"group"`
}

// A map of default settings that will apply to all jobs in the workflow.
type Defaults struct {
	Run *Run `json:"run,omitempty"`
}

type Run struct {
	Shell            *string `json:"shell,omitempty"`
	WorkingDirectory *string `json:"working-directory,omitempty"`
}

// A workflow run is made up of one or more jobs. Jobs run in parallel by default. To run
// jobs sequentially, you can define dependencies on other jobs using the
// jobs.<job_id>.needs keyword.
// Each job runs in a fresh instance of the virtual environment specified by runs-on.
// You can run an unlimited number of jobs as long as you are within the workflow usage
// limits. For more information, see
// https://help.github.com/en/github/automating-your-workflow-with-github-actions/workflow-syntax-for-github-actions#usage-limits.
type Jobs struct {
}

type OnClass struct {
	// Runs your workflow anytime the branch_protection_rule event occurs. More than one
	// activity type triggers this event.
	BranchProtectionRule *PurpleEventObject `json:"branch_protection_rule"`
	// Runs your workflow anytime the check_run event occurs. More than one activity type
	// triggers this event. For information about the REST API, see
	// https://developer.github.com/v3/checks/runs.
	CheckRun *FluffyEventObject `json:"check_run"`
	// Runs your workflow anytime the check_suite event occurs. More than one activity type
	// triggers this event. For information about the REST API, see
	// https://developer.github.com/v3/checks/suites/.
	CheckSuite *TentacledEventObject `json:"check_suite"`
	// Runs your workflow anytime someone creates a branch or tag, which triggers the create
	// event. For information about the REST API, see
	// https://developer.github.com/v3/git/refs/#create-a-reference.
	Create map[string]interface{} `json:"create"`
	// Runs your workflow anytime someone deletes a branch or tag, which triggers the delete
	// event. For information about the REST API, see
	// https://developer.github.com/v3/git/refs/#delete-a-reference.
	Delete map[string]interface{} `json:"delete"`
	// Runs your workflow anytime someone creates a deployment, which triggers the deployment
	// event. Deployments created with a commit SHA may not have a Git ref. For information
	// about the REST API, see https://developer.github.com/v3/repos/deployments/.
	Deployment map[string]interface{} `json:"deployment"`
	// Runs your workflow anytime a third party provides a deployment status, which triggers the
	// deployment_status event. Deployments created with a commit SHA may not have a Git ref.
	// For information about the REST API, see
	// https://developer.github.com/v3/repos/deployments/#create-a-deployment-status.
	DeploymentStatus map[string]interface{} `json:"deployment_status"`
	// Runs your workflow anytime the discussion event occurs. More than one activity type
	// triggers this event. For information about the GraphQL API, see
	// https://docs.github.com/en/graphql/guides/using-the-graphql-api-for-discussions
	Discussion *StickyEventObject `json:"discussion"`
	// Runs your workflow anytime the discussion_comment event occurs. More than one activity
	// type triggers this event. For information about the GraphQL API, see
	// https://docs.github.com/en/graphql/guides/using-the-graphql-api-for-discussions
	DiscussionComment *IndigoEventObject `json:"discussion_comment"`
	// Runs your workflow anytime when someone forks a repository, which triggers the fork
	// event. For information about the REST API, see
	// https://developer.github.com/v3/repos/forks/#create-a-fork.
	Fork map[string]interface{} `json:"fork"`
	// Runs your workflow when someone creates or updates a Wiki page, which triggers the gollum
	// event.
	Gollum map[string]interface{} `json:"gollum"`
	// Runs your workflow anytime the issue_comment event occurs. More than one activity type
	// triggers this event. For information about the REST API, see
	// https://developer.github.com/v3/issues/comments/.
	IssueComment *IndecentEventObject `json:"issue_comment"`
	// Runs your workflow anytime the issues event occurs. More than one activity type triggers
	// this event. For information about the REST API, see
	// https://developer.github.com/v3/issues.
	Issues *HilariousEventObject `json:"issues"`
	// Runs your workflow anytime the label event occurs. More than one activity type triggers
	// this event. For information about the REST API, see
	// https://developer.github.com/v3/issues/labels/.
	Label *AmbitiousEventObject `json:"label"`
	// Runs your workflow anytime the member event occurs. More than one activity type triggers
	// this event. For information about the REST API, see
	// https://developer.github.com/v3/repos/collaborators/.
	Member *CunningEventObject `json:"member"`
	// Runs your workflow when a pull request is added to a merge queue, which adds the pull
	// request to a merge group. For information about the merge queue, see
	// https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/incorporating-changes-from-a-pull-request/merging-a-pull-request-with-a-merge-queue
	// .
	MergeGroup *MagentaEventObject `json:"merge_group"`
	// Runs your workflow anytime the milestone event occurs. More than one activity type
	// triggers this event. For information about the REST API, see
	// https://developer.github.com/v3/issues/milestones/.
	Milestone *FriskyEventObject `json:"milestone"`
	// Runs your workflow anytime someone pushes to a GitHub Pages-enabled branch, which
	// triggers the page_build event. For information about the REST API, see
	// https://developer.github.com/v3/repos/pages/.
	PageBuild map[string]interface{} `json:"page_build"`
	// Runs your workflow anytime the project event occurs. More than one activity type triggers
	// this event. For information about the REST API, see
	// https://developer.github.com/v3/projects/.
	Project *MischievousEventObject `json:"project"`
	// Runs your workflow anytime the project_card event occurs. More than one activity type
	// triggers this event. For information about the REST API, see
	// https://developer.github.com/v3/projects/cards.
	ProjectCard *BraggadociousEventObject `json:"project_card"`
	// Runs your workflow anytime the project_column event occurs. More than one activity type
	// triggers this event. For information about the REST API, see
	// https://developer.github.com/v3/projects/columns.
	ProjectColumn *EventObject1 `json:"project_column"`
	// Runs your workflow anytime someone makes a private repository public, which triggers the
	// public event. For information about the REST API, see
	// https://developer.github.com/v3/repos/#edit.
	Public map[string]interface{} `json:"public"`
	// Runs your workflow anytime the pull_request event occurs. More than one activity type
	// triggers this event. For information about the REST API, see
	// https://developer.github.com/v3/pulls.
	// Note: Workflows do not run on private base repositories when you open a pull request from
	// a forked repository.
	// When you create a pull request from a forked repository to the base repository, GitHub
	// sends the pull_request event to the base repository and no pull request events occur on
	// the forked repository.
	// Workflows don't run on forked repositories by default. You must enable GitHub Actions in
	// the Actions tab of the forked repository.
	// The permissions for the GITHUB_TOKEN in forked repositories is read-only. For more
	// information about the GITHUB_TOKEN, see
	// https://help.github.com/en/articles/virtual-environments-for-github-actions.
	PullRequest *PurpleRef `json:"pull_request"`
	// Runs your workflow anytime the pull_request_review event occurs. More than one activity
	// type triggers this event. For information about the REST API, see
	// https://developer.github.com/v3/pulls/reviews.
	// Note: Workflows do not run on private base repositories when you open a pull request from
	// a forked repository.
	// When you create a pull request from a forked repository to the base repository, GitHub
	// sends the pull_request event to the base repository and no pull request events occur on
	// the forked repository.
	// Workflows don't run on forked repositories by default. You must enable GitHub Actions in
	// the Actions tab of the forked repository.
	// The permissions for the GITHUB_TOKEN in forked repositories is read-only. For more
	// information about the GITHUB_TOKEN, see
	// https://help.github.com/en/articles/virtual-environments-for-github-actions.
	PullRequestReview *EventObject2 `json:"pull_request_review"`
	// Runs your workflow anytime a comment on a pull request's unified diff is modified, which
	// triggers the pull_request_review_comment event. More than one activity type triggers this
	// event. For information about the REST API, see
	// https://developer.github.com/v3/pulls/comments.
	// Note: Workflows do not run on private base repositories when you open a pull request from
	// a forked repository.
	// When you create a pull request from a forked repository to the base repository, GitHub
	// sends the pull_request event to the base repository and no pull request events occur on
	// the forked repository.
	// Workflows don't run on forked repositories by default. You must enable GitHub Actions in
	// the Actions tab of the forked repository.
	// The permissions for the GITHUB_TOKEN in forked repositories is read-only. For more
	// information about the GITHUB_TOKEN, see
	// https://help.github.com/en/articles/virtual-environments-for-github-actions.
	PullRequestReviewComment *EventObject3 `json:"pull_request_review_comment"`
	// This event is similar to pull_request, except that it runs in the context of the base
	// repository of the pull request, rather than in the merge commit. This means that you can
	// more safely make your secrets available to the workflows triggered by the pull request,
	// because only workflows defined in the commit on the base repository are run. For example,
	// this event allows you to create workflows that label and comment on pull requests, based
	// on the contents of the event payload.
	PullRequestTarget *FluffyRef `json:"pull_request_target"`
	// Runs your workflow when someone pushes to a repository branch, which triggers the push
	// event.
	// Note: The webhook payload available to GitHub Actions does not include the added,
	// removed, and modified attributes in the commit object. You can retrieve the full commit
	// object using the REST API. For more information, see
	// https://developer.github.com/v3/repos/commits/#get-a-single-commit.
	Push *TentacledRef `json:"push"`
	// Runs your workflow anytime a package is published or updated. For more information, see
	// https://help.github.com/en/github/managing-packages-with-github-packages.
	RegistryPackage *EventObject4 `json:"registry_package"`
	// Runs your workflow anytime the release event occurs. More than one activity type triggers
	// this event. For information about the REST API, see
	// https://developer.github.com/v3/repos/releases/ in the GitHub Developer documentation.
	Release *EventObject5 `json:"release"`
	// You can use the GitHub API to trigger a webhook event called repository_dispatch when you
	// want to trigger a workflow for activity that happens outside of GitHub. For more
	// information, see
	// https://developer.github.com/v3/repos/#create-a-repository-dispatch-event.
	// To trigger the custom repository_dispatch webhook event, you must send a POST request to
	// a GitHub API endpoint and provide an event_type name to describe the activity type. To
	// trigger a workflow run, you must also configure your workflow to use the
	// repository_dispatch event.
	RepositoryDispatch map[string]interface{} `json:"repository_dispatch"`
	// You can schedule a workflow to run at specific UTC times using POSIX cron syntax
	// (https://pubs.opengroup.org/onlinepubs/9699919799/utilities/crontab.html#tag_20_25_07).
	// Scheduled workflows run on the latest commit on the default or base branch. The shortest
	// interval you can run scheduled workflows is once every 5 minutes.
	// Note: GitHub Actions does not support the non-standard syntax @yearly, @monthly, @weekly,
	// @daily, @hourly, and @reboot.
	// You can use crontab guru (https://crontab.guru/). to help generate your cron syntax and
	// confirm what time it will run. To help you get started, there is also a list of crontab
	// guru examples (https://crontab.guru/examples.html).
	Schedule []ScheduleElement `json:"schedule,omitempty"`
	// Runs your workflow anytime the status of a Git commit changes, which triggers the status
	// event. For information about the REST API, see
	// https://developer.github.com/v3/repos/statuses/.
	Status map[string]interface{} `json:"status"`
	// Runs your workflow anytime the watch event occurs. More than one activity type triggers
	// this event. For information about the REST API, see
	// https://developer.github.com/v3/activity/starring/.
	Watch map[string]interface{} `json:"watch"`
	// Allows workflows to be reused by other workflows.
	WorkflowCall *WorkflowCallUnion `json:"workflow_call"`
	// You can now create workflows that are manually triggered with the new workflow_dispatch
	// event. You will then see a 'Run workflow' button on the Actions tab, enabling you to
	// easily trigger a run.
	WorkflowDispatch *WorkflowDispatchUnion `json:"workflow_dispatch"`
	// This event occurs when a workflow run is requested or completed, and allows you to
	// execute a workflow based on the finished result of another workflow. For example, if your
	// pull_request workflow generates build artifacts, you can create a new workflow that uses
	// workflow_run to analyze the results and add a comment to the original pull request.
	WorkflowRun *EventObject6 `json:"workflow_run"`
}

type PurpleEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type FluffyEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type TentacledEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type StickyEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type IndigoEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type IndecentEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type HilariousEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type AmbitiousEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type CunningEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type MagentaEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type FriskyEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type MischievousEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type BraggadociousEventObject struct {
	Types []interface{} `json:"types,omitempty"`
}

type EventObject1 struct {
	Types []interface{} `json:"types,omitempty"`
}

type PurpleRef struct {
	Types []interface{} `json:"types,omitempty"`
}

type EventObject2 struct {
	Types []interface{} `json:"types,omitempty"`
}

type EventObject3 struct {
	Types []interface{} `json:"types,omitempty"`
}

type FluffyRef struct {
	Types []interface{} `json:"types,omitempty"`
}

type TentacledRef struct {
}

type EventObject4 struct {
	Types []interface{} `json:"types,omitempty"`
}

type EventObject5 struct {
	Types []interface{} `json:"types,omitempty"`
}

type ScheduleClass struct {
	Cron *string `json:"cron,omitempty"`
}

type WorkflowCallClass struct {
	// When using the workflow_call keyword, you can optionally specify inputs that are passed
	// to the called workflow from the caller workflow.
	Inputs *WorkflowCallInputs `json:"inputs,omitempty"`
	// A map of the secrets that can be used in the called workflow. Within the called workflow,
	// you can use the secrets context to refer to a secret.
	Secrets *SecretsUnion `json:"secrets"`
}

// When using the workflow_call keyword, you can optionally specify inputs that are passed
// to the called workflow from the caller workflow.
type WorkflowCallInputs struct {
}

type SecretsClass struct {
}

type WorkflowDispatchClass struct {
	// Input parameters allow you to specify data that the action expects to use during runtime.
	// GitHub stores input parameters as environment variables. Input ids with uppercase letters
	// are converted to lowercase during runtime. We recommended using lowercase input ids.
	Inputs *WorkflowDispatchInputs `json:"inputs,omitempty"`
}

// Input parameters allow you to specify data that the action expects to use during runtime.
// GitHub stores input parameters as environment variables. Input ids with uppercase letters
// are converted to lowercase during runtime. We recommended using lowercase input ids.
type WorkflowDispatchInputs struct {
}

type EventObject6 struct {
	Types     []interface{} `json:"types,omitempty"`
	Workflows []string      `json:"workflows,omitempty"`
}

type PermissionsEvent struct {
	Actions            *PermissionsLevel `json:"actions,omitempty"`
	Checks             *PermissionsLevel `json:"checks,omitempty"`
	Contents           *PermissionsLevel `json:"contents,omitempty"`
	Deployments        *PermissionsLevel `json:"deployments,omitempty"`
	Discussions        *PermissionsLevel `json:"discussions,omitempty"`
	IDToken            *PermissionsLevel `json:"id-token,omitempty"`
	Issues             *PermissionsLevel `json:"issues,omitempty"`
	Packages           *PermissionsLevel `json:"packages,omitempty"`
	Pages              *PermissionsLevel `json:"pages,omitempty"`
	PullRequests       *PermissionsLevel `json:"pull-requests,omitempty"`
	RepositoryProjects *PermissionsLevel `json:"repository-projects,omitempty"`
	SecurityEvents     *PermissionsLevel `json:"security-events,omitempty"`
	Statuses           *PermissionsLevel `json:"statuses,omitempty"`
}

type Event string

const (
	BranchProtectionRule     Event = "branch_protection_rule"
	CheckRun                 Event = "check_run"
	CheckSuite               Event = "check_suite"
	Create                   Event = "create"
	Delete                   Event = "delete"
	Deployment               Event = "deployment"
	DeploymentStatus         Event = "deployment_status"
	Discussion               Event = "discussion"
	DiscussionComment        Event = "discussion_comment"
	Fork                     Event = "fork"
	Gollum                   Event = "gollum"
	IssueComment             Event = "issue_comment"
	Issues                   Event = "issues"
	Label                    Event = "label"
	Member                   Event = "member"
	Milestone                Event = "milestone"
	PageBuild                Event = "page_build"
	Project                  Event = "project"
	ProjectCard              Event = "project_card"
	ProjectColumn            Event = "project_column"
	Public                   Event = "public"
	PullRequest              Event = "pull_request"
	PullRequestReview        Event = "pull_request_review"
	PullRequestReviewComment Event = "pull_request_review_comment"
	PullRequestTarget        Event = "pull_request_target"
	Push                     Event = "push"
	RegistryPackage          Event = "registry_package"
	Release                  Event = "release"
	RepositoryDispatch       Event = "repository_dispatch"
	Status                   Event = "status"
	Watch                    Event = "watch"
	WorkflowCall             Event = "workflow_call"
	WorkflowDispatch         Event = "workflow_dispatch"
	WorkflowRun              Event = "workflow_run"
)

type PermissionsLevel string

const (
	None  PermissionsLevel = "none"
	Read  PermissionsLevel = "read"
	Write PermissionsLevel = "write"
)

type PermissionsEnum string

const (
	ReadAll  PermissionsEnum = "read-all"
	WriteAll PermissionsEnum = "write-all"
)

// Concurrency ensures that only a single job or workflow using the same concurrency group
// will run at a time. A concurrency group can be any string or expression. The expression
// can use any context except for the secrets context.
// You can also specify concurrency at the workflow level.
// When a concurrent job or workflow is queued, if another job or workflow using the same
// concurrency group in the repository is in progress, the queued job or workflow will be
// pending. Any previously pending job or workflow in the concurrency group will be
// canceled. To also cancel any currently running job or workflow in the same concurrency
// group, specify cancel-in-progress: true.
type ConcurrencyUnion struct {
	Concurrency *Concurrency
	String      *string
}

// To cancel any currently running job or workflow in the same concurrency group, specify
// cancel-in-progress: true.
type CancelInProgress struct {
	Bool   *bool
	String *string
}

// A map of environment variables that are available to all jobs and steps in the workflow.
//
// To set custom environment variables, you need to specify the variables in the workflow
// file. You can define environment variables for a step, job, or entire workflow using the
// jobs.<job_id>.steps[*].env, jobs.<job_id>.env, and env keywords. For more information,
// see
// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idstepsenv
type Env struct {
	String   *string
	UnionMap map[string]*EnvValue
}

type EnvValue struct {
	Bool   *bool
	Double *float64
	String *string
}

// The name of the GitHub event that triggers the workflow. You can provide a single event
// string, array of events, array of event types, or an event configuration map that
// schedules a workflow or restricts the execution of a workflow to specific files, tags, or
// branch changes. For a list of available events, see
// https://help.github.com/en/github/automating-your-workflow-with-github-actions/events-that-trigger-workflows.
type OnUnion struct {
	Enum      *Event
	EnumArray []Event
	OnClass   *OnClass
}

type ScheduleElement struct {
	AnythingArray []interface{}
	Bool          *bool
	Double        *float64
	Integer       *int64
	ScheduleClass *ScheduleClass
	String        *string
}

// Allows workflows to be reused by other workflows.
type WorkflowCallUnion struct {
	AnythingArray     []interface{}
	Bool              *bool
	Double            *float64
	Integer           *int64
	String            *string
	WorkflowCallClass *WorkflowCallClass
}

// A map of the secrets that can be used in the called workflow. Within the called workflow,
// you can use the secrets context to refer to a secret.
type SecretsUnion struct {
	AnythingArray []interface{}
	Bool          *bool
	Double        *float64
	Integer       *int64
	SecretsClass  *SecretsClass
	String        *string
}

// You can now create workflows that are manually triggered with the new workflow_dispatch
// event. You will then see a 'Run workflow' button on the Actions tab, enabling you to
// easily trigger a run.
type WorkflowDispatchUnion struct {
	AnythingArray         []interface{}
	Bool                  *bool
	Double                *float64
	Integer               *int64
	String                *string
	WorkflowDispatchClass *WorkflowDispatchClass
}

// You can modify the default permissions granted to the GITHUB_TOKEN, adding or removing
// access as required, so that you only allow the minimum required access.
type Permissions struct {
	Enum             *PermissionsEnum
	PermissionsEvent *PermissionsEvent
}
