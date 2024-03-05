package apis

type KubeConfig struct {
	Clusters []KubeCluster `yaml:"clusters"`
}

type KubeCluster struct {
	Cluster KubeClusterSettings `yaml:"cluster"`
}

type KubeClusterSettings struct {
	Server string `yaml:"server"`
}
