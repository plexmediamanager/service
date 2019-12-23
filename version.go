package service

import (
    format "fmt"
    "reflect"
    "strconv"
    "strings"
)

type VersionDescriptor struct {
    Major       uint64
    Minor       uint64
    Patch       uint64
    Build       string
    BuiltOn     string
    Branch      string
}

var DefaultVersion = VersionDescriptor {
    Major:      1,
    Minor:      0,
    Patch:      0,
    Build:      "",
    BuiltOn:    "",
}

// Convert application version to string
func (version VersionDescriptor) String() string {
    return format.Sprintf("%d.%d.%d-%s",
        version.Major,
        version.Minor,
        version.Patch,
        version.Build,
    )
}

// Just version number string
func (version VersionDescriptor) StringVersion() string {
    return format.Sprintf("%d.%d.%d",
        version.Major,
        version.Minor,
        version.Patch,
    )
}

// Convert application version to string with branch name
func (version VersionDescriptor) StringWithBranch() string {
    return format.Sprintf("%s-%s", version.String(), version.Branch)
}

// Convert application version to string with build date
func (version VersionDescriptor) StringWithBuildDate() string {
    return format.Sprintf("%s Built On %s", version.String(), version.BuiltOn)
}

// Convert application version to string branch name and build date
func (version VersionDescriptor) StringWithBranchAndBuildDate() string {
    return format.Sprintf("%s Built On %s", version.StringWithBranch(), version.BuiltOn)
}

// Pretty application version
func (version VersionDescriptor) StringPretty() string {
    return format.Sprintf("ver. %s rev. %s %s (%s branch)",
        version.StringVersion(),
        version.Build,
        version.BuiltOn,
        version.Branch,
    )
}

// "Initialize" application version
func (version *VersionDescriptor) Initialize() {
    var majorVersion uint64
    var minorVersion uint64
    var patchVersion uint64
    versionPart := strings.SplitN(GitVersion, ".", strings.Count(GitVersion, ".") + 1)
    if len(versionPart) < 3 {
        majorVersion = 0
        minorVersion = 0
        patchVersion = 0
    } else {
        var majorError error
        var minorError error
        var patchError error

        majorVersion, majorError = strconv.ParseUint(versionPart[0], 10, 64)
        if majorError != nil {
            majorVersion = 1
        }

        minorVersion, minorError = strconv.ParseUint(versionPart[1], 10, 64)
        if minorError != nil {
            minorVersion = 0
        }
        patchVersion, patchError = strconv.ParseUint(versionPart[2], 10, 64)
        if patchError != nil {
            patchVersion = 0
        }
    }

    if GitCommitShort == "" {
        GitCommitShort = "unknown"
    }

    if BuildDateUTC == "" {
        BuildDateUTC = "0000-00-00 00:00:00 +0000"
    }

    if GitBranch == "" {
        GitBranch = "archived"
    }

    version.Major = majorVersion
    version.Minor = minorVersion
    version.Patch = patchVersion
    version.Build    = GitCommitShort
    version.BuiltOn  = BuildDateUTC
    version.Branch   = GitBranch
}

// Compare two versions for being equal to each other
func (version VersionDescriptor) Equal(otherVersion VersionDescriptor) bool {
    return reflect.DeepEqual(version, otherVersion)
}

// Check if application version is newer than one that is provided
func (version VersionDescriptor) Newer(otherVersion VersionDescriptor) bool {
    if version.Major > otherVersion.Major {
        return true
    } else if version.Minor > otherVersion.Minor {
        return true
    } else if version.Patch > otherVersion.Patch {
        return true
    } else if version.Build != otherVersion.Build {
        return true
    }
    return false
}

// Check if current application version is older than one that is provided
func (version VersionDescriptor) Older(otherVersion VersionDescriptor) bool {
    return !version.Equal(otherVersion) && !version.Newer(otherVersion)
}
