# Go-Enjin Third-Party Features

This repository contains Go-Enjin features that rely upon modified 3rd party
libraries or libraries which require CGO. The features and examples are Go-Enjin
things and everything under the `pkg` path are strictly 3rd-party, external to
Go-Enjin, codebases.

The forks present are to be considered in an intermediate phase where changes
have been made to the upstream code and may or may not make it all the way back
upstream.

Currently these 3rd party features focus on supporting [Atlassian] [Jira Cloud]
plugins providing [Atlassian Connect] modules (general pages and dashboard
items) and supporting SCSS features (which requires CGO to compile).

This repository (and Go-Enjin in general) are not associated with [Atlassian].

[Atlassian]: https://atlassian.com
[Jira Cloud]: https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/
[Atlassian Connect]: https://developer.atlassian.com/cloud/jira/platform/about-connect-modules-for-jira/
