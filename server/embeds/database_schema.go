package embeds

import "embed"

//go:embed sqls/version_*.sql
var DatabaseSchema embed.FS
