package service

import format "fmt"

const (
    RedisServiceName            =   "micro.redis"
    DatabaseServiceName         =   "micro.database"
    TMDBServiceName             =   "micro.tmdb"
    JackettServiceName          =   "micro.jackett"
    TorrentServiceName          =   "micro.torrent"
)

var (
    BuildDateLocal      string
    BuildDateUTC        string
    GitCommitLong       string
    GitCommitShort      string
    GitBranch           string
    GitState            string
    GitAuthor           string
    GitVersion          string
    GitSummary          string
)

func GetServiceName(serviceName string) string {
    if application, done := FromContext(); done {
        return format.Sprintf("%s.%s-%s", application.Vendor(), serviceName, application.Environment())
    }
    panic("[Net::Service::getServiceName] This was not supposed to happen")
}
