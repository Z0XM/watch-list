package database

func Migrate() {
	err := GetCluster().Cluster.master.AutoMigrate()
	if err != nil {
		//panic("Failed to migrate user table") // TODO
	}
}
