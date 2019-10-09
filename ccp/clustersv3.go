package ccp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ClusterV3 struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Provider string `json:"provider"`
	Status   string `json:"status"`
	Spec     struct {
		Name                 string        `json:"name"`
		Type                 string        `json:"type"`
		KubernetesVersion    string        `json:"kubernetes_version"`
		IPAllocationMethod   string        `json:"ip_allocation_method"`
		MasterVip            string        `json:"master_vip"`
		LoadBalancerNum      int           `json:"load_balancer_num"`
		SubnetID             string        `json:"subnet_id"`
		NtpPools             []interface{} `json:"ntp_pools"`
		NtpServers           []interface{} `json:"ntp_servers"`
		RootCaRegistries     []interface{} `json:"root_ca_registries"`
		SelfSignedRegistries struct {
		} `json:"self_signed_registries"`
		VsphereInfra struct {
			Cluster      string   `json:"cluster"`
			Datacenter   string   `json:"datacenter"`
			Datastore    string   `json:"datastore"`
			Folder       string   `json:"folder"`
			GuestOS      string   `json:"guestOS"`
			HostSystem   string   `json:"hostSystem"`
			Networks     []string `json:"networks"`
			ResourcePool string   `json:"resource_pool"`
		} `json:"vsphere_infra"`
		MasterGroup struct {
			Gpus     []interface{} `json:"gpus"`
			Labels   interface{}   `json:"labels"`
			Name     string        `json:"name"`
			Size     int           `json:"size"`
			Taints   interface{}   `json:"taints"`
			Template string        `json:"template"`
			Vcpus    int           `json:"vcpus"`
			MemoryMb int           `json:"memory_mb"`
			SSHKey   string        `json:"ssh_key"`
			SSHUser  string        `json:"ssh_user"`
			Nodes    []struct {
				Name      string `json:"name"`
				Status    string `json:"status"`
				Phase     string `json:"phase"`
				PrivateIP string `json:"private_ip"`
				PublicIP  string `json:"public_ip"`
			} `json:"nodes"`
		} `json:"master_group"`
		NodeGroups []struct {
			Gpus     []interface{} `json:"gpus"`
			Labels   interface{}   `json:"labels"`
			Name     string        `json:"name"`
			Size     int           `json:"size"`
			Taints   interface{}   `json:"taints"`
			Template string        `json:"template"`
			Vcpus    int           `json:"vcpus"`
			MemoryMb int           `json:"memory_mb"`
			SSHKey   string        `json:"ssh_key"`
			SSHUser  string        `json:"ssh_user"`
			Nodes    []struct {
				Name      string `json:"name"`
				Status    string `json:"status"`
				Phase     string `json:"phase"`
				PrivateIP string `json:"private_ip"`
				PublicIP  string `json:"public_ip"`
			} `json:"nodes"`
		} `json:"node_groups"`
		NetworkPluginProfile struct {
			Details struct {
				TyphaReplicas string `json:"typhaReplicas"`
				PodCidr       string `json:"pod_cidr"`
				SSHUser       string `json:"ssh_user"`
			} `json:"details"`
			Name string `json:"name"`
		} `json:"network_plugin_profile"`
		KubernetesConfigSecret string        `json:"kubernetes_config_secret"`
		IngressAsLb            bool          `json:"ingress_as_lb"`
		NginxIngressClass      string        `json:"nginx_ingress_class"`
		EtcdEncrypted          bool          `json:"etcd_encrypted"`
		SkipManagement         interface{}   `json:"skip_management"`
		DockerNoProxy          []interface{} `json:"docker_no_proxy"`
	} `json:"spec"`
	Kubeconfig             string        `json:"kubeconfig"`
	KubernetesVersion      string        `json:"kubernetes_version"`
	KubernetesConfigSecret interface{}   `json:"kubernetes_config_secret"`
	IPAllocationMethod     string        `json:"ip_allocation_method"`
	MasterVip              string        `json:"master_vip"`
	LoadBalancerNum        int           `json:"load_balancer_num"`
	SubnetID               string        `json:"subnet_id"`
	NtpPools               []interface{} `json:"ntp_pools"`
	NtpServers             []interface{} `json:"ntp_servers"`
	RootCaRegistries       []interface{} `json:"root_ca_registries"`
	SelfSignedRegistries   struct {
	} `json:"self_signed_registries"`
	InsecureRegistries []interface{} `json:"insecure_registries"`
	DockerHTTPProxy    string        `json:"docker_http_proxy"`
	DockerHTTPSProxy   string        `json:"docker_https_proxy"`
	VsphereInfra       struct {
		Datacenter   string   `json:"datacenter"`
		Datastore    string   `json:"datastore"`
		Networks     []string `json:"networks"`
		Cluster      string   `json:"cluster"`
		ResourcePool string   `json:"resource_pool"`
		Folder       string   `json:"folder"`
	} `json:"vsphere_infra"`
	MasterGroup struct {
		Name     string        `json:"name"`
		Size     int           `json:"size"`
		Template string        `json:"template"`
		Vcpus    int           `json:"vcpus"`
		MemoryMb int           `json:"memory_mb"`
		Gpus     []interface{} `json:"gpus"`
		SSHUser  string        `json:"ssh_user"`
		SSHKey   string        `json:"ssh_key"`
		Nodes    []struct {
			Name      string `json:"name"`
			Status    string `json:"status"`
			Phase     string `json:"phase"`
			PrivateIP string `json:"private_ip"`
			PublicIP  string `json:"public_ip"`
		} `json:"nodes"`
	} `json:"master_group"`
	NodeGroups []struct {
		Name     string        `json:"name"`
		Size     int           `json:"size"`
		Template string        `json:"template"`
		Vcpus    int           `json:"vcpus"`
		MemoryMb int           `json:"memory_mb"`
		Gpus     []interface{} `json:"gpus"`
		SSHUser  string        `json:"ssh_user"`
		SSHKey   string        `json:"ssh_key"`
		Nodes    []struct {
			Name      string `json:"name"`
			Status    string `json:"status"`
			Phase     string `json:"phase"`
			PrivateIP string `json:"private_ip"`
			PublicIP  string `json:"public_ip"`
		} `json:"nodes"`
	} `json:"node_groups"`
	NetworkPluginProfile struct {
		Details struct {
			TyphaReplicas string `json:"typhaReplicas"`
			PodCidr       string `json:"pod_cidr"`
			SSHUser       string `json:"ssh_user"`
		} `json:"details"`
		Name string `json:"name"`
	} `json:"network_plugin_profile"`
	IngressAsLb       bool          `json:"ingress_as_lb"`
	NginxIngressClass string        `json:"nginx_ingress_class"`
	EtcdEncrypted     bool          `json:"etcd_encrypted"`
	SkipManagement    bool          `json:"skip_management"`
	DockerNoProxy     []interface{} `json:"docker_no_proxy"`
	RoutableCidr      interface{}   `json:"routable_cidr"`
	ImagePrefix       interface{}   `json:"image_prefix"`
	AciProfile        interface{}   `json:"aci_profile"`
}

func (s *Client) GetClusterV3(clusterName string) (*ClusterV3, error) {

	url := fmt.Sprintf(s.BaseURL + "/v3/clusters/" + clusterName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *ClusterV3

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Client) GetClustersV3() ([]ClusterV3, error) {

	url := fmt.Sprintf(s.BaseURL + "/v3/clusters")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data []ClusterV3

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}