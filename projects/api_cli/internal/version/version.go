package version

// Main version number being run right now.
var Version = "dev"

// User Agent name set in requests.
const UserAgent = "Medtronic-CLI"

// String prints the version of the Spin CLI.
func String() string {
	return Version
}
