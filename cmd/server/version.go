package main

var (
	buildVersion = "N/A" // Default version placeholder; updated at compile time via ldflags.
	buildDate    = "N/A" // Build date; populated automatically using go build flags.
	buildCommit  = "N/A" // Commit hash; retrieved from git repository at build time.
)
