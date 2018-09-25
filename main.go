package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

type MyResponse struct {
	Message string `json:"Answer:"`
}

type PullRequestEvent struct {
	Action      string `json:"action"`
	Number      int    `json:"number"`
	PullRequest struct {
		URL      string `json:"url"`
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		HTMLURL  string `json:"html_url"`
		DiffURL  string `json:"diff_url"`
		PatchURL string `json:"patch_url"`
		IssueURL string `json:"issue_url"`
		Number   int    `json:"number"`
		State    string `json:"state"`
		Locked   bool   `json:"locked"`
		Title    string `json:"title"`
		User     struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"user"`
		Body               string        `json:"body"`
		CreatedAt          time.Time     `json:"created_at"`
		UpdatedAt          time.Time     `json:"updated_at"`
		ClosedAt           time.Time     `json:"closed_at"`
		MergedAt           time.Time     `json:"merged_at"`
		MergeCommitSha     string        `json:"merge_commit_sha"`
		Assignee           interface{}   `json:"assignee"`
		Assignees          []interface{} `json:"assignees"`
		RequestedReviewers []interface{} `json:"requested_reviewers"`
		RequestedTeams     []interface{} `json:"requested_teams"`
		Labels             []interface{} `json:"labels"`
		Milestone          interface{}   `json:"milestone"`
		CommitsURL         string        `json:"commits_url"`
		ReviewCommentsURL  string        `json:"review_comments_url"`
		ReviewCommentURL   string        `json:"review_comment_url"`
		CommentsURL        string        `json:"comments_url"`
		StatusesURL        string        `json:"statuses_url"`
		Head               struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  struct {
				Login             string `json:"login"`
				ID                int    `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"user"`
			Repo struct {
				ID       int    `json:"id"`
				NodeID   string `json:"node_id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Private  bool   `json:"private"`
				Owner    struct {
					Login             string `json:"login"`
					ID                int    `json:"id"`
					NodeID            string `json:"node_id"`
					AvatarURL         string `json:"avatar_url"`
					GravatarID        string `json:"gravatar_id"`
					URL               string `json:"url"`
					HTMLURL           string `json:"html_url"`
					FollowersURL      string `json:"followers_url"`
					FollowingURL      string `json:"following_url"`
					GistsURL          string `json:"gists_url"`
					StarredURL        string `json:"starred_url"`
					SubscriptionsURL  string `json:"subscriptions_url"`
					OrganizationsURL  string `json:"organizations_url"`
					ReposURL          string `json:"repos_url"`
					EventsURL         string `json:"events_url"`
					ReceivedEventsURL string `json:"received_events_url"`
					Type              string `json:"type"`
					SiteAdmin         bool   `json:"site_admin"`
				} `json:"owner"`
				HTMLURL          string      `json:"html_url"`
				Description      interface{} `json:"description"`
				Fork             bool        `json:"fork"`
				URL              string      `json:"url"`
				ForksURL         string      `json:"forks_url"`
				KeysURL          string      `json:"keys_url"`
				CollaboratorsURL string      `json:"collaborators_url"`
				TeamsURL         string      `json:"teams_url"`
				HooksURL         string      `json:"hooks_url"`
				IssueEventsURL   string      `json:"issue_events_url"`
				EventsURL        string      `json:"events_url"`
				AssigneesURL     string      `json:"assignees_url"`
				BranchesURL      string      `json:"branches_url"`
				TagsURL          string      `json:"tags_url"`
				BlobsURL         string      `json:"blobs_url"`
				GitTagsURL       string      `json:"git_tags_url"`
				GitRefsURL       string      `json:"git_refs_url"`
				TreesURL         string      `json:"trees_url"`
				StatusesURL      string      `json:"statuses_url"`
				LanguagesURL     string      `json:"languages_url"`
				StargazersURL    string      `json:"stargazers_url"`
				ContributorsURL  string      `json:"contributors_url"`
				SubscribersURL   string      `json:"subscribers_url"`
				SubscriptionURL  string      `json:"subscription_url"`
				CommitsURL       string      `json:"commits_url"`
				GitCommitsURL    string      `json:"git_commits_url"`
				CommentsURL      string      `json:"comments_url"`
				IssueCommentURL  string      `json:"issue_comment_url"`
				ContentsURL      string      `json:"contents_url"`
				CompareURL       string      `json:"compare_url"`
				MergesURL        string      `json:"merges_url"`
				ArchiveURL       string      `json:"archive_url"`
				DownloadsURL     string      `json:"downloads_url"`
				IssuesURL        string      `json:"issues_url"`
				PullsURL         string      `json:"pulls_url"`
				MilestonesURL    string      `json:"milestones_url"`
				NotificationsURL string      `json:"notifications_url"`
				LabelsURL        string      `json:"labels_url"`
				ReleasesURL      string      `json:"releases_url"`
				DeploymentsURL   string      `json:"deployments_url"`
				CreatedAt        time.Time   `json:"created_at"`
				UpdatedAt        time.Time   `json:"updated_at"`
				PushedAt         time.Time   `json:"pushed_at"`
				GitURL           string      `json:"git_url"`
				SSHURL           string      `json:"ssh_url"`
				CloneURL         string      `json:"clone_url"`
				SvnURL           string      `json:"svn_url"`
				Homepage         interface{} `json:"homepage"`
				Size             int         `json:"size"`
				StargazersCount  int         `json:"stargazers_count"`
				WatchersCount    int         `json:"watchers_count"`
				Language         interface{} `json:"language"`
				HasIssues        bool        `json:"has_issues"`
				HasProjects      bool        `json:"has_projects"`
				HasDownloads     bool        `json:"has_downloads"`
				HasWiki          bool        `json:"has_wiki"`
				HasPages         bool        `json:"has_pages"`
				ForksCount       int         `json:"forks_count"`
				MirrorURL        interface{} `json:"mirror_url"`
				Archived         bool        `json:"archived"`
				OpenIssuesCount  int         `json:"open_issues_count"`
				License          interface{} `json:"license"`
				Forks            int         `json:"forks"`
				OpenIssues       int         `json:"open_issues"`
				Watchers         int         `json:"watchers"`
				DefaultBranch    string      `json:"default_branch"`
			} `json:"repo"`
		} `json:"head"`
		Base struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  struct {
				Login             string `json:"login"`
				ID                int    `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"user"`
			Repo struct {
				ID       int    `json:"id"`
				NodeID   string `json:"node_id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Private  bool   `json:"private"`
				Owner    struct {
					Login             string `json:"login"`
					ID                int    `json:"id"`
					NodeID            string `json:"node_id"`
					AvatarURL         string `json:"avatar_url"`
					GravatarID        string `json:"gravatar_id"`
					URL               string `json:"url"`
					HTMLURL           string `json:"html_url"`
					FollowersURL      string `json:"followers_url"`
					FollowingURL      string `json:"following_url"`
					GistsURL          string `json:"gists_url"`
					StarredURL        string `json:"starred_url"`
					SubscriptionsURL  string `json:"subscriptions_url"`
					OrganizationsURL  string `json:"organizations_url"`
					ReposURL          string `json:"repos_url"`
					EventsURL         string `json:"events_url"`
					ReceivedEventsURL string `json:"received_events_url"`
					Type              string `json:"type"`
					SiteAdmin         bool   `json:"site_admin"`
				} `json:"owner"`
				HTMLURL          string      `json:"html_url"`
				Description      interface{} `json:"description"`
				Fork             bool        `json:"fork"`
				URL              string      `json:"url"`
				ForksURL         string      `json:"forks_url"`
				KeysURL          string      `json:"keys_url"`
				CollaboratorsURL string      `json:"collaborators_url"`
				TeamsURL         string      `json:"teams_url"`
				HooksURL         string      `json:"hooks_url"`
				IssueEventsURL   string      `json:"issue_events_url"`
				EventsURL        string      `json:"events_url"`
				AssigneesURL     string      `json:"assignees_url"`
				BranchesURL      string      `json:"branches_url"`
				TagsURL          string      `json:"tags_url"`
				BlobsURL         string      `json:"blobs_url"`
				GitTagsURL       string      `json:"git_tags_url"`
				GitRefsURL       string      `json:"git_refs_url"`
				TreesURL         string      `json:"trees_url"`
				StatusesURL      string      `json:"statuses_url"`
				LanguagesURL     string      `json:"languages_url"`
				StargazersURL    string      `json:"stargazers_url"`
				ContributorsURL  string      `json:"contributors_url"`
				SubscribersURL   string      `json:"subscribers_url"`
				SubscriptionURL  string      `json:"subscription_url"`
				CommitsURL       string      `json:"commits_url"`
				GitCommitsURL    string      `json:"git_commits_url"`
				CommentsURL      string      `json:"comments_url"`
				IssueCommentURL  string      `json:"issue_comment_url"`
				ContentsURL      string      `json:"contents_url"`
				CompareURL       string      `json:"compare_url"`
				MergesURL        string      `json:"merges_url"`
				ArchiveURL       string      `json:"archive_url"`
				DownloadsURL     string      `json:"downloads_url"`
				IssuesURL        string      `json:"issues_url"`
				PullsURL         string      `json:"pulls_url"`
				MilestonesURL    string      `json:"milestones_url"`
				NotificationsURL string      `json:"notifications_url"`
				LabelsURL        string      `json:"labels_url"`
				ReleasesURL      string      `json:"releases_url"`
				DeploymentsURL   string      `json:"deployments_url"`
				CreatedAt        time.Time   `json:"created_at"`
				UpdatedAt        time.Time   `json:"updated_at"`
				PushedAt         time.Time   `json:"pushed_at"`
				GitURL           string      `json:"git_url"`
				SSHURL           string      `json:"ssh_url"`
				CloneURL         string      `json:"clone_url"`
				SvnURL           string      `json:"svn_url"`
				Homepage         interface{} `json:"homepage"`
				Size             int         `json:"size"`
				StargazersCount  int         `json:"stargazers_count"`
				WatchersCount    int         `json:"watchers_count"`
				Language         interface{} `json:"language"`
				HasIssues        bool        `json:"has_issues"`
				HasProjects      bool        `json:"has_projects"`
				HasDownloads     bool        `json:"has_downloads"`
				HasWiki          bool        `json:"has_wiki"`
				HasPages         bool        `json:"has_pages"`
				ForksCount       int         `json:"forks_count"`
				MirrorURL        interface{} `json:"mirror_url"`
				Archived         bool        `json:"archived"`
				OpenIssuesCount  int         `json:"open_issues_count"`
				License          interface{} `json:"license"`
				Forks            int         `json:"forks"`
				OpenIssues       int         `json:"open_issues"`
				Watchers         int         `json:"watchers"`
				DefaultBranch    string      `json:"default_branch"`
			} `json:"repo"`
		} `json:"base"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Issue struct {
				Href string `json:"href"`
			} `json:"issue"`
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			ReviewComments struct {
				Href string `json:"href"`
			} `json:"review_comments"`
			ReviewComment struct {
				Href string `json:"href"`
			} `json:"review_comment"`
			Commits struct {
				Href string `json:"href"`
			} `json:"commits"`
			Statuses struct {
				Href string `json:"href"`
			} `json:"statuses"`
		} `json:"_links"`
		AuthorAssociation string      `json:"author_association"`
		Merged            bool        `json:"merged"`
		Mergeable         interface{} `json:"mergeable"`
		Rebaseable        interface{} `json:"rebaseable"`
		MergeableState    string      `json:"mergeable_state"`
		MergedBy          struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"merged_by"`
		Comments            int  `json:"comments"`
		ReviewComments      int  `json:"review_comments"`
		MaintainerCanModify bool `json:"maintainer_can_modify"`
		Commits             int  `json:"commits"`
		Additions           int  `json:"additions"`
		Deletions           int  `json:"deletions"`
		ChangedFiles        int  `json:"changed_files"`
	} `json:"pull_request"`
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		HTMLURL          string      `json:"html_url"`
		Description      interface{} `json:"description"`
		Fork             bool        `json:"fork"`
		URL              string      `json:"url"`
		ForksURL         string      `json:"forks_url"`
		KeysURL          string      `json:"keys_url"`
		CollaboratorsURL string      `json:"collaborators_url"`
		TeamsURL         string      `json:"teams_url"`
		HooksURL         string      `json:"hooks_url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		EventsURL        string      `json:"events_url"`
		AssigneesURL     string      `json:"assignees_url"`
		BranchesURL      string      `json:"branches_url"`
		TagsURL          string      `json:"tags_url"`
		BlobsURL         string      `json:"blobs_url"`
		GitTagsURL       string      `json:"git_tags_url"`
		GitRefsURL       string      `json:"git_refs_url"`
		TreesURL         string      `json:"trees_url"`
		StatusesURL      string      `json:"statuses_url"`
		LanguagesURL     string      `json:"languages_url"`
		StargazersURL    string      `json:"stargazers_url"`
		ContributorsURL  string      `json:"contributors_url"`
		SubscribersURL   string      `json:"subscribers_url"`
		SubscriptionURL  string      `json:"subscription_url"`
		CommitsURL       string      `json:"commits_url"`
		GitCommitsURL    string      `json:"git_commits_url"`
		CommentsURL      string      `json:"comments_url"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		ContentsURL      string      `json:"contents_url"`
		CompareURL       string      `json:"compare_url"`
		MergesURL        string      `json:"merges_url"`
		ArchiveURL       string      `json:"archive_url"`
		DownloadsURL     string      `json:"downloads_url"`
		IssuesURL        string      `json:"issues_url"`
		PullsURL         string      `json:"pulls_url"`
		MilestonesURL    string      `json:"milestones_url"`
		NotificationsURL string      `json:"notifications_url"`
		LabelsURL        string      `json:"labels_url"`
		ReleasesURL      string      `json:"releases_url"`
		DeploymentsURL   string      `json:"deployments_url"`
		CreatedAt        time.Time   `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
		PushedAt         time.Time   `json:"pushed_at"`
		GitURL           string      `json:"git_url"`
		SSHURL           string      `json:"ssh_url"`
		CloneURL         string      `json:"clone_url"`
		SvnURL           string      `json:"svn_url"`
		Homepage         interface{} `json:"homepage"`
		Size             int         `json:"size"`
		StargazersCount  int         `json:"stargazers_count"`
		WatchersCount    int         `json:"watchers_count"`
		Language         interface{} `json:"language"`
		HasIssues        bool        `json:"has_issues"`
		HasProjects      bool        `json:"has_projects"`
		HasDownloads     bool        `json:"has_downloads"`
		HasWiki          bool        `json:"has_wiki"`
		HasPages         bool        `json:"has_pages"`
		ForksCount       int         `json:"forks_count"`
		MirrorURL        interface{} `json:"mirror_url"`
		Archived         bool        `json:"archived"`
		OpenIssuesCount  int         `json:"open_issues_count"`
		License          interface{} `json:"license"`
		Forks            int         `json:"forks"`
		OpenIssues       int         `json:"open_issues"`
		Watchers         int         `json:"watchers"`
		DefaultBranch    string      `json:"default_branch"`
	} `json:"repository"`
	Sender struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"sender"`
	Installation struct {
		ID     int    `json:"id"`
		NodeID string `json:"node_id"`
	} `json:"installation"`
}

type GithubKeys struct {
	AccessToken string `json:"GITHUB_ACCESS_TOKEN"`
}

func delete_branch(event PullRequestEvent) (MyResponse, error) {
	fmt.Println("log start")
	fmt.Println(event)
	fmt.Println("log end")

	var message MyResponse
	var errorType error

	// hi @kamontia! I checked this lambda function works from GitHub-App via API-Gateway.
	// I leave the hint to implement we realy want to do!
	//
	// currently, I already cache PullRequestEvent from GitHub,
	// and Got values we need.
	//
	// action: event.Action (i.e closed or merged)
	// repo  : event.PullRequest.Head.Repo.FullName (i.e chaspy/delete-branch)
	// branch: event.PullRequest.Head.Ref (i.e test-branch)
	// ref https://github.com/chaspy/test-delete-branch/pull/1
	// (I used this PR to test)
	//
	// So what you should implement is very simple!
	//
	// if action is "closed" or "merged" then
	//  delete branch(repo,branch)
	// end
	// return nothing(204 No Content)
	//
	// If there is concern, it will be part of certification.
	// ref: https://github.com/github/platform-samples/blob/b047a807dd43a3f76c2cbf0e85af3ffadeb2b880/api/ruby/building-your-first-github-app/advanced_server.rb#L24-L44
	if event.Action == "closed" || event.Action == "merged" {
		fmt.Println(event.Action)

		// Create client via go-github library
		tr := http.DefaultTransport
		itr, err := ghinstallation.NewKeyFromFile(tr, 17630, 348213, "private-key.pem")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("itr:%v", itr)
		// Use installation transport with github.com/google/go-github
		client := github.NewClient(&http.Client{Transport: itr})

		//Create request
		Owner := event.PullRequest.User.Login
		RepoName := event.PullRequest.Head.Repo.Name
		HeadRefs := "heads/" + event.PullRequest.Head.Ref

		ctx := context.Background()
		resp, err := client.Git.DeleteRef(ctx, Owner, RepoName, HeadRefs)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error Occured")
			return message, errorType
		}
		fmt.Println("response Status : ", resp.Status)
		fmt.Println("response Headers : ", resp.Header)
		fmt.Println("response Body : ", resp.Body)

	} else {
		fmt.Println("no match action")
	}

	return MyResponse{Message: fmt.Sprintf("PullRequest action is %s!! repo is %s!! branch name is %s!!", event.Action, event.PullRequest.Head.Ref, event.PullRequest.Head.Repo.FullName)}, nil
}

func main() {
	lambda.Start(delete_branch)
}
