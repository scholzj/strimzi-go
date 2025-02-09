/*
Copyright 2025 Jakub Scholz.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta2

import (
    corev1 "k8s.io/api/core/v1"
    networkingv1 "k8s.io/api/networking/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    
    Items []Kafka `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Kafka struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    
    Spec *KafkaSpec `json:"spec,omitempty"`
    Status *KafkaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaNodePoolList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    
    Items []KafkaNodePool `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaNodePool struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    
    Spec *KafkaNodePoolSpec `json:"spec,omitempty"`
    Status *KafkaNodePoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaConnectList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    
    Items []KafkaConnect `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaConnect struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    
    Spec *KafkaConnectSpec `json:"spec,omitempty"`
    Status *KafkaConnectStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaMirrorMaker2List struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    
    Items []KafkaMirrorMaker2 `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaMirrorMaker2 struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    
    Spec *KafkaMirrorMaker2Spec `json:"spec,omitempty"`
    Status *KafkaMirrorMaker2Status `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaBridgeList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    
    Items []KafkaBridge `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaBridge struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    
    Spec *KafkaBridgeSpec `json:"spec,omitempty"`
    Status *KafkaBridgeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaRebalanceList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    
    Items []KafkaRebalance `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaRebalance struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    
    Spec *KafkaRebalanceSpec `json:"spec,omitempty"`
    Status *KafkaRebalanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaTopicList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    
    Items []KafkaTopic `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaTopic struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    
    Spec *KafkaTopicSpec `json:"spec,omitempty"`
    Status *KafkaTopicStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaConnectorList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    
    Items []KafkaConnector `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaConnector struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    
    Spec *KafkaConnectorSpec `json:"spec,omitempty"`
    Status *KafkaConnectorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaUserList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    
    Items []KafkaUser `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KafkaUser struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
    
    Spec *KafkaUserSpec `json:"spec,omitempty"`
    Status *KafkaUserStatus `json:"status,omitempty"`
}

type KafkaUserStatus struct {
    Conditions []Condition `json:"conditions,omitempty"`
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`
    Username string `json:"username,omitempty"`
    Secret string `json:"secret,omitempty"`
}

type Condition struct {
    Type string `json:"type,omitempty"`
    Status string `json:"status,omitempty"`
    LastTransitionTime string `json:"lastTransitionTime,omitempty"`
    Reason string `json:"reason,omitempty"`
    Message string `json:"message,omitempty"`
}

type KafkaUserSpec struct {
    Authentication *KafkaUserAuthentication `json:"authentication,omitempty"`
    Authorization *KafkaUserAuthorization `json:"authorization,omitempty"`
    Quotas *KafkaUserQuotas `json:"quotas,omitempty"`
    Template *KafkaUserTemplate `json:"template,omitempty"`
}

type KafkaUserTemplate struct {
    Secret *ResourceTemplate `json:"secret,omitempty"`
}

type ResourceTemplate struct {
    Metadata *MetadataTemplate `json:"metadata,omitempty"`
}

type MetadataTemplate struct {
    Labels map[string]string `json:"labels,omitempty"`
    Annotations map[string]string `json:"annotations,omitempty"`
}

type KafkaUserQuotas struct {
    ProducerByteRate int32 `json:"producerByteRate,omitempty"`
    ConsumerByteRate int32 `json:"consumerByteRate,omitempty"`
    RequestPercentage int32 `json:"requestPercentage,omitempty"`
    ControllerMutationRate float64 `json:"controllerMutationRate,omitempty"`
}

type KafkaUserAuthorizationType string
const (
    SIMPLE_KAFKAUSERAUTHORIZATIONTYPE KafkaUserAuthorizationType = "simple"
)

type KafkaUserAuthorization struct {
    Type KafkaUserAuthorizationType `json:"type,omitempty"`
    Acls []AclRule `json:"acls,omitempty"`
}

type AclRule struct {
    Type AclRuleType `json:"type,omitempty"`
    Resource *AclRuleResource `json:"resource,omitempty"`
    Host string `json:"host,omitempty"`
    Operation AclOperation `json:"operation,omitempty"`
    Operations []AclOperation `json:"operations,omitempty"`
}

type AclOperation string
const (
    READ_ACLOPERATION AclOperation = "Read"
    WRITE_ACLOPERATION AclOperation = "Write"
    CREATE_ACLOPERATION AclOperation = "Create"
    DELETE_ACLOPERATION AclOperation = "Delete"
    ALTER_ACLOPERATION AclOperation = "Alter"
    DESCRIBE_ACLOPERATION AclOperation = "Describe"
    CLUSTERACTION_ACLOPERATION AclOperation = "ClusterAction"
    ALTERCONFIGS_ACLOPERATION AclOperation = "AlterConfigs"
    DESCRIBECONFIGS_ACLOPERATION AclOperation = "DescribeConfigs"
    IDEMPOTENTWRITE_ACLOPERATION AclOperation = "IdempotentWrite"
    ALL_ACLOPERATION AclOperation = "All"
)

type AclRuleResourceType string
const (
    TOPIC_ACLRULERESOURCETYPE AclRuleResourceType = "topic"
    GROUP_ACLRULERESOURCETYPE AclRuleResourceType = "group"
    CLUSTER_ACLRULERESOURCETYPE AclRuleResourceType = "cluster"
    TRANSACTIONALID_ACLRULERESOURCETYPE AclRuleResourceType = "transactionalId"
)

type AclRuleResource struct {
    Name string `json:"name,omitempty"`
    PatternType AclResourcePatternType `json:"patternType,omitempty"`
    Type AclRuleResourceType `json:"type,omitempty"`
}

type AclResourcePatternType string
const (
    LITERAL_ACLRESOURCEPATTERNTYPE AclResourcePatternType = "literal"
    PREFIX_ACLRESOURCEPATTERNTYPE AclResourcePatternType = "prefix"
)

type AclRuleType string
const (
    ALLOW_ACLRULETYPE AclRuleType = "allow"
    DENY_ACLRULETYPE AclRuleType = "deny"
)

type KafkaUserAuthenticationType string
const (
    TLS_KAFKAUSERAUTHENTICATIONTYPE KafkaUserAuthenticationType = "tls"
    TLS_EXTERNAL_KAFKAUSERAUTHENTICATIONTYPE KafkaUserAuthenticationType = "tls-external"
    SCRAM_SHA_512_KAFKAUSERAUTHENTICATIONTYPE KafkaUserAuthenticationType = "scram-sha-512"
)

type KafkaUserAuthentication struct {
    Type KafkaUserAuthenticationType `json:"type,omitempty"`
    Password *Password `json:"password,omitempty"`
}

type Password struct {
    ValueFrom *PasswordSource `json:"valueFrom,omitempty"`
}

type PasswordSource struct {
    SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

type KafkaConnectorStatus struct {
    Conditions []Condition `json:"conditions,omitempty"`
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`
    AutoRestart *AutoRestartStatus `json:"autoRestart,omitempty"`
    ConnectorStatus JSONValue `json:"connectorStatus,omitempty"`
    TasksMax int32 `json:"tasksMax,omitempty"`
    Topics []string `json:"topics,omitempty"`
}

type AutoRestartStatus struct {
    Count int32 `json:"count,omitempty"`
    ConnectorName string `json:"connectorName,omitempty"`
    LastRestartTimestamp string `json:"lastRestartTimestamp,omitempty"`
}

type KafkaConnectorSpec struct {
    Class string `json:"class,omitempty"`
    TasksMax int32 `json:"tasksMax,omitempty"`
    AutoRestart *AutoRestart `json:"autoRestart,omitempty"`
    Config JSONValue `json:"config,omitempty"`
    Pause bool `json:"pause,omitempty"`
    State ConnectorState `json:"state,omitempty"`
    ListOffsets *ListOffsets `json:"listOffsets,omitempty"`
    AlterOffsets *AlterOffsets `json:"alterOffsets,omitempty"`
}

type AlterOffsets struct {
    FromConfigMap *corev1.LocalObjectReference `json:"fromConfigMap,omitempty"`
}

type ListOffsets struct {
    ToConfigMap *corev1.LocalObjectReference `json:"toConfigMap,omitempty"`
}

type ConnectorState string
const (
    PAUSED_CONNECTORSTATE ConnectorState = "paused"
    STOPPED_CONNECTORSTATE ConnectorState = "stopped"
    RUNNING_CONNECTORSTATE ConnectorState = "running"
)

type AutoRestart struct {
    Enabled bool `json:"enabled,omitempty"`
    MaxRestarts int32 `json:"maxRestarts,omitempty"`
}

type KafkaTopicStatus struct {
    Conditions []Condition `json:"conditions,omitempty"`
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`
    TopicName string `json:"topicName,omitempty"`
    TopicId string `json:"topicId,omitempty"`
    ReplicasChange *ReplicasChangeStatus `json:"replicasChange,omitempty"`
}

type ReplicasChangeStatus struct {
    TargetReplicas int32 `json:"targetReplicas,omitempty"`
    State ReplicasChangeState `json:"state,omitempty"`
    Message string `json:"message,omitempty"`
    SessionId string `json:"sessionId,omitempty"`
}

type ReplicasChangeState string
const (
    PENDING_REPLICASCHANGESTATE ReplicasChangeState = "pending"
    ONGOING_REPLICASCHANGESTATE ReplicasChangeState = "ongoing"
)

type KafkaTopicSpec struct {
    TopicName string `json:"topicName,omitempty"`
    Partitions int32 `json:"partitions,omitempty"`
    Replicas int32 `json:"replicas,omitempty"`
    Config JSONValue `json:"config,omitempty"`
}

type KafkaRebalanceStatus struct {
    Conditions []Condition `json:"conditions,omitempty"`
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`
    SessionId string `json:"sessionId,omitempty"`
    OptimizationResult JSONValue `json:"optimizationResult,omitempty"`
}

type KafkaRebalanceSpec struct {
    Mode KafkaRebalanceMode `json:"mode,omitempty"`
    Brokers []int32 `json:"brokers,omitempty"`
    Goals []string `json:"goals,omitempty"`
    SkipHardGoalCheck bool `json:"skipHardGoalCheck,omitempty"`
    RebalanceDisk bool `json:"rebalanceDisk,omitempty"`
    ExcludedTopics string `json:"excludedTopics,omitempty"`
    ConcurrentPartitionMovementsPerBroker int32 `json:"concurrentPartitionMovementsPerBroker,omitempty"`
    ConcurrentIntraBrokerPartitionMovements int32 `json:"concurrentIntraBrokerPartitionMovements,omitempty"`
    ConcurrentLeaderMovements int32 `json:"concurrentLeaderMovements,omitempty"`
    ReplicationThrottle int64 `json:"replicationThrottle,omitempty"`
    ReplicaMovementStrategies []string `json:"replicaMovementStrategies,omitempty"`
    MoveReplicasOffVolumes []BrokerAndVolumeIds `json:"moveReplicasOffVolumes,omitempty"`
}

type BrokerAndVolumeIds struct {
    BrokerId int32 `json:"brokerId,omitempty"`
    VolumeIds []int32 `json:"volumeIds,omitempty"`
}

type KafkaRebalanceMode string
const (
    FULL_KAFKAREBALANCEMODE KafkaRebalanceMode = "full"
    ADD_BROKERS_KAFKAREBALANCEMODE KafkaRebalanceMode = "add-brokers"
    REMOVE_BROKERS_KAFKAREBALANCEMODE KafkaRebalanceMode = "remove-brokers"
    REMOVE_DISKS_KAFKAREBALANCEMODE KafkaRebalanceMode = "remove-disks"
)

type KafkaBridgeStatus struct {
    Conditions []Condition `json:"conditions,omitempty"`
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`
    Url string `json:"url,omitempty"`
    Replicas int32 `json:"replicas,omitempty"`
    LabelSelector string `json:"labelSelector,omitempty"`
}

type KafkaBridgeSpec struct {
    Replicas int32 `json:"replicas,omitempty"`
    Image string `json:"image,omitempty"`
    BootstrapServers string `json:"bootstrapServers,omitempty"`
    Tls *ClientTls `json:"tls,omitempty"`
    Authentication *KafkaClientAuthentication `json:"authentication,omitempty"`
    Http *KafkaBridgeHttpConfig `json:"http,omitempty"`
    AdminClient *KafkaBridgeAdminClientSpec `json:"adminClient,omitempty"`
    Consumer *KafkaBridgeConsumerSpec `json:"consumer,omitempty"`
    Producer *KafkaBridgeProducerSpec `json:"producer,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
    JvmOptions *JvmOptions `json:"jvmOptions,omitempty"`
    Logging *Logging `json:"logging,omitempty"`
    ClientRackInitImage string `json:"clientRackInitImage,omitempty"`
    Rack *Rack `json:"rack,omitempty"`
    EnableMetrics bool `json:"enableMetrics,omitempty"`
    LivenessProbe *Probe `json:"livenessProbe,omitempty"`
    ReadinessProbe *Probe `json:"readinessProbe,omitempty"`
    Template *KafkaBridgeTemplate `json:"template,omitempty"`
    Tracing *Tracing `json:"tracing,omitempty"`
}

type TracingType string
const (
    JAEGER_TRACINGTYPE TracingType = "jaeger"
    OPENTELEMETRY_TRACINGTYPE TracingType = "opentelemetry"
)

type Tracing struct {
    Type TracingType `json:"type,omitempty"`
}

type KafkaBridgeTemplate struct {
    Deployment *DeploymentTemplate `json:"deployment,omitempty"`
    Pod *PodTemplate `json:"pod,omitempty"`
    ApiService *InternalServiceTemplate `json:"apiService,omitempty"`
    PodDisruptionBudget *PodDisruptionBudgetTemplate `json:"podDisruptionBudget,omitempty"`
    BridgeContainer *ContainerTemplate `json:"bridgeContainer,omitempty"`
    ClusterRoleBinding *ResourceTemplate `json:"clusterRoleBinding,omitempty"`
    ServiceAccount *ResourceTemplate `json:"serviceAccount,omitempty"`
    InitContainer *ContainerTemplate `json:"initContainer,omitempty"`
}

type ContainerTemplate struct {
    Env []ContainerEnvVar `json:"env,omitempty"`
    SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty"`
    VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`
}

type ContainerEnvVar struct {
    Name string `json:"name,omitempty"`
    Value string `json:"value,omitempty"`
    ValueFrom *ContainerEnvVarSource `json:"valueFrom,omitempty"`
}

type ContainerEnvVarSource struct {
    SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty"`
    ConfigMapKeyRef *corev1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
}

type PodDisruptionBudgetTemplate struct {
    Metadata *MetadataTemplate `json:"metadata,omitempty"`
    MaxUnavailable int32 `json:"maxUnavailable,omitempty"`
}

type InternalServiceTemplate struct {
    Metadata *MetadataTemplate `json:"metadata,omitempty"`
    IpFamilyPolicy IpFamilyPolicy `json:"ipFamilyPolicy,omitempty"`
    IpFamilies []IpFamily `json:"ipFamilies,omitempty"`
}

type IpFamily string
const (
    IPV4_IPFAMILY IpFamily = "IPv4"
    IPV6_IPFAMILY IpFamily = "IPv6"
)

type IpFamilyPolicy string
const (
    SINGLE_STACK_IPFAMILYPOLICY IpFamilyPolicy = "SingleStack"
    PREFER_DUAL_STACK_IPFAMILYPOLICY IpFamilyPolicy = "PreferDualStack"
    REQUIRE_DUAL_STACK_IPFAMILYPOLICY IpFamilyPolicy = "RequireDualStack"
)

type PodTemplate struct {
    Metadata *MetadataTemplate `json:"metadata,omitempty"`
    ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
    SecurityContext *corev1.PodSecurityContext `json:"securityContext,omitempty"`
    TerminationGracePeriodSeconds int32 `json:"terminationGracePeriodSeconds,omitempty"`
    Affinity *corev1.Affinity `json:"affinity,omitempty"`
    Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
    TopologySpreadConstraints []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
    PriorityClassName string `json:"priorityClassName,omitempty"`
    SchedulerName string `json:"schedulerName,omitempty"`
    HostAliases []corev1.HostAlias `json:"hostAliases,omitempty"`
    EnableServiceLinks bool `json:"enableServiceLinks,omitempty"`
    TmpDirSizeLimit string `json:"tmpDirSizeLimit,omitempty"`
    Volumes []AdditionalVolume `json:"volumes,omitempty"`
}

type AdditionalVolume struct {
    Name string `json:"name,omitempty"`
    Secret *corev1.SecretVolumeSource `json:"secret,omitempty"`
    ConfigMap *corev1.ConfigMapVolumeSource `json:"configMap,omitempty"`
    EmptyDir *corev1.EmptyDirVolumeSource `json:"emptyDir,omitempty"`
    PersistentVolumeClaim *corev1.PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"`
    Csi *corev1.CSIVolumeSource `json:"csi,omitempty"`
}

type DeploymentTemplate struct {
    Metadata *MetadataTemplate `json:"metadata,omitempty"`
    DeploymentStrategy DeploymentStrategy `json:"deploymentStrategy,omitempty"`
}

type DeploymentStrategy string
const (
    ROLLING_UPDATE_DEPLOYMENTSTRATEGY DeploymentStrategy = "RollingUpdate"
    RECREATE_DEPLOYMENTSTRATEGY DeploymentStrategy = "Recreate"
)

type Probe struct {
    InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty"`
    TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`
    PeriodSeconds int32 `json:"periodSeconds,omitempty"`
    SuccessThreshold int32 `json:"successThreshold,omitempty"`
    FailureThreshold int32 `json:"failureThreshold,omitempty"`
}

type Rack struct {
    TopologyKey string `json:"topologyKey,omitempty"`
}

type LoggingType string
const (
    INLINE_LOGGINGTYPE LoggingType = "inline"
    EXTERNAL_LOGGINGTYPE LoggingType = "external"
)

type Logging struct {
    Type LoggingType `json:"type,omitempty"`
    Loggers map[string]string `json:"loggers,omitempty"`
    ValueFrom *ExternalConfigurationReference `json:"valueFrom,omitempty"`
}

type ExternalConfigurationReference struct {
    ConfigMapKeyRef *corev1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
}

type JvmOptions struct {
    XX map[string]string `json:"-XX,omitempty"`
    Xmx string `json:"-Xmx,omitempty"`
    Xms string `json:"-Xms,omitempty"`
    GcLoggingEnabled bool `json:"gcLoggingEnabled,omitempty"`
    JavaSystemProperties []SystemProperty `json:"javaSystemProperties,omitempty"`
}

type SystemProperty struct {
    Name string `json:"name,omitempty"`
    Value string `json:"value,omitempty"`
}

type KafkaBridgeProducerSpec struct {
    Enabled bool `json:"enabled,omitempty"`
    Config JSONValue `json:"config,omitempty"`
}

type KafkaBridgeConsumerSpec struct {
    Enabled bool `json:"enabled,omitempty"`
    TimeoutSeconds int64 `json:"timeoutSeconds,omitempty"`
    Config JSONValue `json:"config,omitempty"`
}

type KafkaBridgeAdminClientSpec struct {
    Config JSONValue `json:"config,omitempty"`
}

type KafkaBridgeHttpConfig struct {
    Port int32 `json:"port,omitempty"`
    Cors *KafkaBridgeHttpCors `json:"cors,omitempty"`
}

type KafkaBridgeHttpCors struct {
    AllowedOrigins []string `json:"allowedOrigins,omitempty"`
    AllowedMethods []string `json:"allowedMethods,omitempty"`
}

type KafkaClientAuthenticationType string
const (
    TLS_KAFKACLIENTAUTHENTICATIONTYPE KafkaClientAuthenticationType = "tls"
    SCRAM_SHA_256_KAFKACLIENTAUTHENTICATIONTYPE KafkaClientAuthenticationType = "scram-sha-256"
    SCRAM_SHA_512_KAFKACLIENTAUTHENTICATIONTYPE KafkaClientAuthenticationType = "scram-sha-512"
    PLAIN_KAFKACLIENTAUTHENTICATIONTYPE KafkaClientAuthenticationType = "plain"
    OAUTH_KAFKACLIENTAUTHENTICATIONTYPE KafkaClientAuthenticationType = "oauth"
)

type KafkaClientAuthentication struct {
    ClientAssertionLocation string `json:"clientAssertionLocation,omitempty"`
    Type KafkaClientAuthenticationType `json:"type,omitempty"`
    AccessTokenIsJwt bool `json:"accessTokenIsJwt,omitempty"`
    TlsTrustedCertificates []CertSecretSource `json:"tlsTrustedCertificates,omitempty"`
    SaslExtensions map[string]string `json:"saslExtensions,omitempty"`
    DisableTlsHostnameVerification bool `json:"disableTlsHostnameVerification,omitempty"`
    Scope string `json:"scope,omitempty"`
    ClientAssertionType string `json:"clientAssertionType,omitempty"`
    ClientSecret *GenericSecretSource `json:"clientSecret,omitempty"`
    AccessTokenLocation string `json:"accessTokenLocation,omitempty"`
    CertificateAndKey *CertAndKeySecretSource `json:"certificateAndKey,omitempty"`
    Audience string `json:"audience,omitempty"`
    ClientAssertion *GenericSecretSource `json:"clientAssertion,omitempty"`
    ClientId string `json:"clientId,omitempty"`
    ConnectTimeoutSeconds int32 `json:"connectTimeoutSeconds,omitempty"`
    MaxTokenExpirySeconds int32 `json:"maxTokenExpirySeconds,omitempty"`
    HttpRetryPauseMs int32 `json:"httpRetryPauseMs,omitempty"`
    AccessToken *GenericSecretSource `json:"accessToken,omitempty"`
    ReadTimeoutSeconds int32 `json:"readTimeoutSeconds,omitempty"`
    EnableMetrics bool `json:"enableMetrics,omitempty"`
    TokenEndpointUri string `json:"tokenEndpointUri,omitempty"`
    PasswordSecret *PasswordSecretSource `json:"passwordSecret,omitempty"`
    IncludeAcceptHeader bool `json:"includeAcceptHeader,omitempty"`
    HttpRetries int32 `json:"httpRetries,omitempty"`
    Username string `json:"username,omitempty"`
    RefreshToken *GenericSecretSource `json:"refreshToken,omitempty"`
}

type PasswordSecretSource struct {
    SecretName string `json:"secretName,omitempty"`
    Password string `json:"password,omitempty"`
}

type CertAndKeySecretSource struct {
    SecretName string `json:"secretName,omitempty"`
    Certificate string `json:"certificate,omitempty"`
    Key string `json:"key,omitempty"`
}

type GenericSecretSource struct {
    Key string `json:"key,omitempty"`
    SecretName string `json:"secretName,omitempty"`
}

type CertSecretSource struct {
    SecretName string `json:"secretName,omitempty"`
    Certificate string `json:"certificate,omitempty"`
    Pattern string `json:"pattern,omitempty"`
}

type ClientTls struct {
    TrustedCertificates []CertSecretSource `json:"trustedCertificates,omitempty"`
}

type KafkaMirrorMaker2Status struct {
    Conditions []Condition `json:"conditions,omitempty"`
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`
    Url string `json:"url,omitempty"`
    Connectors []Map `json:"connectors,omitempty"`
    AutoRestartStatuses []AutoRestartStatus `json:"autoRestartStatuses,omitempty"`
    ConnectorPlugins []ConnectorPlugin `json:"connectorPlugins,omitempty"`
    LabelSelector string `json:"labelSelector,omitempty"`
    Replicas int32 `json:"replicas,omitempty"`
}

type ConnectorPlugin struct {
    Class string `json:"class,omitempty"`
    Type string `json:"type,omitempty"`
    Version string `json:"version,omitempty"`
}

type Map struct {
    Empty bool `json:"empty,omitempty"`
}

type KafkaMirrorMaker2Spec struct {
    Version string `json:"version,omitempty"`
    Replicas int32 `json:"replicas,omitempty"`
    Image string `json:"image,omitempty"`
    ConnectCluster string `json:"connectCluster,omitempty"`
    Clusters []KafkaMirrorMaker2ClusterSpec `json:"clusters,omitempty"`
    Mirrors []KafkaMirrorMaker2MirrorSpec `json:"mirrors,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
    LivenessProbe *Probe `json:"livenessProbe,omitempty"`
    ReadinessProbe *Probe `json:"readinessProbe,omitempty"`
    JvmOptions *JvmOptions `json:"jvmOptions,omitempty"`
    JmxOptions *KafkaJmxOptions `json:"jmxOptions,omitempty"`
    Logging *Logging `json:"logging,omitempty"`
    ClientRackInitImage string `json:"clientRackInitImage,omitempty"`
    Rack *Rack `json:"rack,omitempty"`
    MetricsConfig *MetricsConfig `json:"metricsConfig,omitempty"`
    Tracing *Tracing `json:"tracing,omitempty"`
    Template *KafkaConnectTemplate `json:"template,omitempty"`
    ExternalConfiguration *ExternalConfiguration `json:"externalConfiguration,omitempty"`
}

type ExternalConfiguration struct {
    Env []ExternalConfigurationEnv `json:"env,omitempty"`
    Volumes []ExternalConfigurationVolumeSource `json:"volumes,omitempty"`
}

type ExternalConfigurationVolumeSource struct {
    Name string `json:"name,omitempty"`
    Secret *corev1.SecretVolumeSource `json:"secret,omitempty"`
    ConfigMap *corev1.ConfigMapVolumeSource `json:"configMap,omitempty"`
}

type ExternalConfigurationEnv struct {
    Name string `json:"name,omitempty"`
    ValueFrom *ExternalConfigurationEnvVarSource `json:"valueFrom,omitempty"`
}

type ExternalConfigurationEnvVarSource struct {
    SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty"`
    ConfigMapKeyRef *corev1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
}

type KafkaConnectTemplate struct {
    Deployment *DeploymentTemplate `json:"deployment,omitempty"`
    PodSet *ResourceTemplate `json:"podSet,omitempty"`
    Pod *PodTemplate `json:"pod,omitempty"`
    ApiService *InternalServiceTemplate `json:"apiService,omitempty"`
    HeadlessService *InternalServiceTemplate `json:"headlessService,omitempty"`
    ConnectContainer *ContainerTemplate `json:"connectContainer,omitempty"`
    InitContainer *ContainerTemplate `json:"initContainer,omitempty"`
    PodDisruptionBudget *PodDisruptionBudgetTemplate `json:"podDisruptionBudget,omitempty"`
    ServiceAccount *ResourceTemplate `json:"serviceAccount,omitempty"`
    ClusterRoleBinding *ResourceTemplate `json:"clusterRoleBinding,omitempty"`
    BuildPod *PodTemplate `json:"buildPod,omitempty"`
    BuildContainer *ContainerTemplate `json:"buildContainer,omitempty"`
    BuildConfig *BuildConfigTemplate `json:"buildConfig,omitempty"`
    BuildServiceAccount *ResourceTemplate `json:"buildServiceAccount,omitempty"`
    JmxSecret *ResourceTemplate `json:"jmxSecret,omitempty"`
}

type BuildConfigTemplate struct {
    Metadata *MetadataTemplate `json:"metadata,omitempty"`
    PullSecret string `json:"pullSecret,omitempty"`
}

type MetricsConfigType string
const (
    JMXPROMETHEUSEXPORTER_METRICSCONFIGTYPE MetricsConfigType = "jmxPrometheusExporter"
)

type MetricsConfig struct {
    Type MetricsConfigType `json:"type,omitempty"`
    ValueFrom *ExternalConfigurationReference `json:"valueFrom,omitempty"`
}

type KafkaJmxOptions struct {
    Authentication *KafkaJmxAuthentication `json:"authentication,omitempty"`
}

type KafkaJmxAuthenticationType string
const (
    PASSWORD_KAFKAJMXAUTHENTICATIONTYPE KafkaJmxAuthenticationType = "password"
)

type KafkaJmxAuthentication struct {
    Type KafkaJmxAuthenticationType `json:"type,omitempty"`
}

type KafkaMirrorMaker2MirrorSpec struct {
    SourceCluster string `json:"sourceCluster,omitempty"`
    TargetCluster string `json:"targetCluster,omitempty"`
    SourceConnector *KafkaMirrorMaker2ConnectorSpec `json:"sourceConnector,omitempty"`
    HeartbeatConnector *KafkaMirrorMaker2ConnectorSpec `json:"heartbeatConnector,omitempty"`
    CheckpointConnector *KafkaMirrorMaker2ConnectorSpec `json:"checkpointConnector,omitempty"`
    TopicsPattern string `json:"topicsPattern,omitempty"`
    TopicsBlacklistPattern string `json:"topicsBlacklistPattern,omitempty"`
    TopicsExcludePattern string `json:"topicsExcludePattern,omitempty"`
    GroupsPattern string `json:"groupsPattern,omitempty"`
    GroupsBlacklistPattern string `json:"groupsBlacklistPattern,omitempty"`
    GroupsExcludePattern string `json:"groupsExcludePattern,omitempty"`
}

type KafkaMirrorMaker2ConnectorSpec struct {
    TasksMax int32 `json:"tasksMax,omitempty"`
    Pause bool `json:"pause,omitempty"`
    Config JSONValue `json:"config,omitempty"`
    State ConnectorState `json:"state,omitempty"`
    AutoRestart *AutoRestart `json:"autoRestart,omitempty"`
    ListOffsets *ListOffsets `json:"listOffsets,omitempty"`
    AlterOffsets *AlterOffsets `json:"alterOffsets,omitempty"`
}

type KafkaMirrorMaker2ClusterSpec struct {
    Alias string `json:"alias,omitempty"`
    BootstrapServers string `json:"bootstrapServers,omitempty"`
    Tls *ClientTls `json:"tls,omitempty"`
    Authentication *KafkaClientAuthentication `json:"authentication,omitempty"`
    Config JSONValue `json:"config,omitempty"`
}

type KafkaConnectStatus struct {
    Conditions []Condition `json:"conditions,omitempty"`
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`
    Url string `json:"url,omitempty"`
    ConnectorPlugins []ConnectorPlugin `json:"connectorPlugins,omitempty"`
    Replicas int32 `json:"replicas,omitempty"`
    LabelSelector string `json:"labelSelector,omitempty"`
}

type KafkaConnectSpec struct {
    Version string `json:"version,omitempty"`
    Replicas int32 `json:"replicas,omitempty"`
    Image string `json:"image,omitempty"`
    BootstrapServers string `json:"bootstrapServers,omitempty"`
    Tls *ClientTls `json:"tls,omitempty"`
    Authentication *KafkaClientAuthentication `json:"authentication,omitempty"`
    Config JSONValue `json:"config,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
    LivenessProbe *Probe `json:"livenessProbe,omitempty"`
    ReadinessProbe *Probe `json:"readinessProbe,omitempty"`
    JvmOptions *JvmOptions `json:"jvmOptions,omitempty"`
    JmxOptions *KafkaJmxOptions `json:"jmxOptions,omitempty"`
    Logging *Logging `json:"logging,omitempty"`
    ClientRackInitImage string `json:"clientRackInitImage,omitempty"`
    Rack *Rack `json:"rack,omitempty"`
    MetricsConfig *MetricsConfig `json:"metricsConfig,omitempty"`
    Tracing *Tracing `json:"tracing,omitempty"`
    Template *KafkaConnectTemplate `json:"template,omitempty"`
    ExternalConfiguration *ExternalConfiguration `json:"externalConfiguration,omitempty"`
    Build *Build `json:"build,omitempty"`
}

type Build struct {
    Output *Output `json:"output,omitempty"`
    Plugins []Plugin `json:"plugins,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

type Plugin struct {
    Name string `json:"name,omitempty"`
    Artifacts []Artifact `json:"artifacts,omitempty"`
}

type ArtifactType string
const (
    JAR_ARTIFACTTYPE ArtifactType = "jar"
    TGZ_ARTIFACTTYPE ArtifactType = "tgz"
    ZIP_ARTIFACTTYPE ArtifactType = "zip"
    MAVEN_ARTIFACTTYPE ArtifactType = "maven"
    OTHER_ARTIFACTTYPE ArtifactType = "other"
)

type Artifact struct {
    Artifact string `json:"artifact,omitempty"`
    Sha512sum string `json:"sha512sum,omitempty"`
    FileName string `json:"fileName,omitempty"`
    Insecure bool `json:"insecure,omitempty"`
    Type ArtifactType `json:"type,omitempty"`
    Repository string `json:"repository,omitempty"`
    Version string `json:"version,omitempty"`
    Url string `json:"url,omitempty"`
    Group string `json:"group,omitempty"`
}

type OutputType string
const (
    DOCKER_OUTPUTTYPE OutputType = "docker"
    IMAGESTREAM_OUTPUTTYPE OutputType = "imagestream"
)

type Output struct {
    Image string `json:"image,omitempty"`
    AdditionalKanikoOptions []string `json:"additionalKanikoOptions,omitempty"`
    Type OutputType `json:"type,omitempty"`
    PushSecret string `json:"pushSecret,omitempty"`
}

type KafkaNodePoolStatus struct {
    Conditions []Condition `json:"conditions,omitempty"`
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`
    NodeIds []int32 `json:"nodeIds,omitempty"`
    ClusterId string `json:"clusterId,omitempty"`
    Roles []ProcessRoles `json:"roles,omitempty"`
    Replicas int32 `json:"replicas,omitempty"`
    LabelSelector string `json:"labelSelector,omitempty"`
}

type ProcessRoles string
const (
    CONTROLLER_PROCESSROLES ProcessRoles = "controller"
    BROKER_PROCESSROLES ProcessRoles = "broker"
)

type KafkaNodePoolSpec struct {
    Replicas int32 `json:"replicas,omitempty"`
    Storage *Storage `json:"storage,omitempty"`
    Roles []ProcessRoles `json:"roles,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
    JvmOptions *JvmOptions `json:"jvmOptions,omitempty"`
    Template *KafkaNodePoolTemplate `json:"template,omitempty"`
}

type KafkaNodePoolTemplate struct {
    PodSet *ResourceTemplate `json:"podSet,omitempty"`
    Pod *PodTemplate `json:"pod,omitempty"`
    PerPodService *ResourceTemplate `json:"perPodService,omitempty"`
    PerPodRoute *ResourceTemplate `json:"perPodRoute,omitempty"`
    PerPodIngress *ResourceTemplate `json:"perPodIngress,omitempty"`
    PersistentVolumeClaim *ResourceTemplate `json:"persistentVolumeClaim,omitempty"`
    KafkaContainer *ContainerTemplate `json:"kafkaContainer,omitempty"`
    InitContainer *ContainerTemplate `json:"initContainer,omitempty"`
}

type StorageType string
const (
    EPHEMERAL_STORAGETYPE StorageType = "ephemeral"
    PERSISTENT_CLAIM_STORAGETYPE StorageType = "persistent-claim"
    JBOD_STORAGETYPE StorageType = "jbod"
)

type Storage struct {
    SizeLimit string `json:"sizeLimit,omitempty"`
    KraftMetadata KRaftMetadataStorage `json:"kraftMetadata,omitempty"`
    Size string `json:"size,omitempty"`
    DeleteClaim bool `json:"deleteClaim,omitempty"`
    Volumes []SingleVolumeStorage `json:"volumes,omitempty"`
    Selector map[string]string `json:"selector,omitempty"`
    Id int32 `json:"id,omitempty"`
    Overrides []PersistentClaimStorageOverride `json:"overrides,omitempty"`
    Type StorageType `json:"type,omitempty"`
    Class string `json:"class,omitempty"`
}

type PersistentClaimStorageOverride struct {
    Class string `json:"class,omitempty"`
    Broker int32 `json:"broker,omitempty"`
}

type SingleVolumeStorageType string
const (
    EPHEMERAL_SINGLEVOLUMESTORAGETYPE SingleVolumeStorageType = "ephemeral"
    PERSISTENT_CLAIM_SINGLEVOLUMESTORAGETYPE SingleVolumeStorageType = "persistent-claim"
)

type SingleVolumeStorage struct {
    SizeLimit string `json:"sizeLimit,omitempty"`
    KraftMetadata KRaftMetadataStorage `json:"kraftMetadata,omitempty"`
    Size string `json:"size,omitempty"`
    DeleteClaim bool `json:"deleteClaim,omitempty"`
    Selector map[string]string `json:"selector,omitempty"`
    Id int32 `json:"id,omitempty"`
    Overrides []PersistentClaimStorageOverride `json:"overrides,omitempty"`
    Type SingleVolumeStorageType `json:"type,omitempty"`
    Class string `json:"class,omitempty"`
}

type KRaftMetadataStorage string
const (
    SHARED_KRAFTMETADATASTORAGE KRaftMetadataStorage = "shared"
)

type KafkaStatus struct {
    Conditions []Condition `json:"conditions,omitempty"`
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`
    Listeners []ListenerStatus `json:"listeners,omitempty"`
    KafkaNodePools []UsedNodePoolStatus `json:"kafkaNodePools,omitempty"`
    RegisteredNodeIds []int32 `json:"registeredNodeIds,omitempty"`
    ClusterId string `json:"clusterId,omitempty"`
    OperatorLastSuccessfulVersion string `json:"operatorLastSuccessfulVersion,omitempty"`
    KafkaVersion string `json:"kafkaVersion,omitempty"`
    KafkaMetadataVersion string `json:"kafkaMetadataVersion,omitempty"`
    KafkaMetadataState KafkaMetadataState `json:"kafkaMetadataState,omitempty"`
    AutoRebalance *KafkaAutoRebalanceStatus `json:"autoRebalance,omitempty"`
}

type KafkaAutoRebalanceStatus struct {
    State KafkaAutoRebalanceState `json:"state,omitempty"`
    LastTransitionTime string `json:"lastTransitionTime,omitempty"`
    Modes []KafkaAutoRebalanceStatusBrokers `json:"modes,omitempty"`
}

type KafkaAutoRebalanceStatusBrokers struct {
    Mode KafkaAutoRebalanceMode `json:"mode,omitempty"`
    Brokers []int32 `json:"brokers,omitempty"`
}

type KafkaAutoRebalanceMode string
const (
    ADD_BROKERS_KAFKAAUTOREBALANCEMODE KafkaAutoRebalanceMode = "add-brokers"
    REMOVE_BROKERS_KAFKAAUTOREBALANCEMODE KafkaAutoRebalanceMode = "remove-brokers"
)

type KafkaAutoRebalanceState string
const (
    IDLE_KAFKAAUTOREBALANCESTATE KafkaAutoRebalanceState = "Idle"
    REBALANCEONSCALEDOWN_KAFKAAUTOREBALANCESTATE KafkaAutoRebalanceState = "RebalanceOnScaleDown"
    REBALANCEONSCALEUP_KAFKAAUTOREBALANCESTATE KafkaAutoRebalanceState = "RebalanceOnScaleUp"
)

type KafkaMetadataState string
const (
    ZOOKEEPER_KAFKAMETADATASTATE KafkaMetadataState = "ZooKeeper"
    KRAFTMIGRATION_KAFKAMETADATASTATE KafkaMetadataState = "KRaftMigration"
    KRAFTDUALWRITING_KAFKAMETADATASTATE KafkaMetadataState = "KRaftDualWriting"
    KRAFTPOSTMIGRATION_KAFKAMETADATASTATE KafkaMetadataState = "KRaftPostMigration"
    PREKRAFT_KAFKAMETADATASTATE KafkaMetadataState = "PreKRaft"
    KRAFT_KAFKAMETADATASTATE KafkaMetadataState = "KRaft"
)

type UsedNodePoolStatus struct {
    Name string `json:"name,omitempty"`
}

type ListenerStatus struct {
    Type string `json:"type,omitempty"`
    Name string `json:"name,omitempty"`
    Addresses []ListenerAddress `json:"addresses,omitempty"`
    BootstrapServers string `json:"bootstrapServers,omitempty"`
    Certificates []string `json:"certificates,omitempty"`
}

type ListenerAddress struct {
    Host string `json:"host,omitempty"`
    Port int32 `json:"port,omitempty"`
}

type KafkaSpec struct {
    Kafka *KafkaClusterSpec `json:"kafka,omitempty"`
    Zookeeper *ZookeeperClusterSpec `json:"zookeeper,omitempty"`
    EntityOperator *EntityOperatorSpec `json:"entityOperator,omitempty"`
    ClusterCa *CertificateAuthority `json:"clusterCa,omitempty"`
    ClientsCa *CertificateAuthority `json:"clientsCa,omitempty"`
    CruiseControl *CruiseControlSpec `json:"cruiseControl,omitempty"`
    JmxTrans *JmxTransSpec `json:"jmxTrans,omitempty"`
    KafkaExporter *KafkaExporterSpec `json:"kafkaExporter,omitempty"`
    MaintenanceTimeWindows []string `json:"maintenanceTimeWindows,omitempty"`
}

type KafkaExporterSpec struct {
    Image string `json:"image,omitempty"`
    GroupRegex string `json:"groupRegex,omitempty"`
    TopicRegex string `json:"topicRegex,omitempty"`
    GroupExcludeRegex string `json:"groupExcludeRegex,omitempty"`
    TopicExcludeRegex string `json:"topicExcludeRegex,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
    Logging string `json:"logging,omitempty"`
    LivenessProbe *Probe `json:"livenessProbe,omitempty"`
    ReadinessProbe *Probe `json:"readinessProbe,omitempty"`
    EnableSaramaLogging bool `json:"enableSaramaLogging,omitempty"`
    ShowAllOffsets bool `json:"showAllOffsets,omitempty"`
    Template *KafkaExporterTemplate `json:"template,omitempty"`
}

type KafkaExporterTemplate struct {
    Deployment *DeploymentTemplate `json:"deployment,omitempty"`
    Pod *PodTemplate `json:"pod,omitempty"`
    Service *ResourceTemplate `json:"service,omitempty"`
    Container *ContainerTemplate `json:"container,omitempty"`
    ServiceAccount *ResourceTemplate `json:"serviceAccount,omitempty"`
}

type JmxTransSpec struct {
    Image string `json:"image,omitempty"`
    OutputDefinitions []JmxTransOutputDefinitionTemplate `json:"outputDefinitions,omitempty"`
    LogLevel string `json:"logLevel,omitempty"`
    KafkaQueries []JmxTransQueryTemplate `json:"kafkaQueries,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
    Template *JmxTransTemplate `json:"template,omitempty"`
}

type JmxTransTemplate struct {
    Deployment *DeploymentTemplate `json:"deployment,omitempty"`
    Pod *PodTemplate `json:"pod,omitempty"`
    Container *ContainerTemplate `json:"container,omitempty"`
    ServiceAccount *ResourceTemplate `json:"serviceAccount,omitempty"`
}

type JmxTransQueryTemplate struct {
    TargetMBean string `json:"targetMBean,omitempty"`
    Attributes []string `json:"attributes,omitempty"`
    Outputs []string `json:"outputs,omitempty"`
}

type JmxTransOutputDefinitionTemplate struct {
    OutputType string `json:"outputType,omitempty"`
    Host string `json:"host,omitempty"`
    Port int32 `json:"port,omitempty"`
    FlushDelayInSeconds int32 `json:"flushDelayInSeconds,omitempty"`
    TypeNames []string `json:"typeNames,omitempty"`
    Name string `json:"name,omitempty"`
}

type CruiseControlSpec struct {
    Image string `json:"image,omitempty"`
    TlsSidecar *TlsSidecar `json:"tlsSidecar,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
    LivenessProbe *Probe `json:"livenessProbe,omitempty"`
    ReadinessProbe *Probe `json:"readinessProbe,omitempty"`
    JvmOptions *JvmOptions `json:"jvmOptions,omitempty"`
    Logging *Logging `json:"logging,omitempty"`
    Template *CruiseControlTemplate `json:"template,omitempty"`
    BrokerCapacity *BrokerCapacity `json:"brokerCapacity,omitempty"`
    Config JSONValue `json:"config,omitempty"`
    MetricsConfig *MetricsConfig `json:"metricsConfig,omitempty"`
    ApiUsers *CruiseControlApiUsers `json:"apiUsers,omitempty"`
    AutoRebalance []KafkaAutoRebalanceConfiguration `json:"autoRebalance,omitempty"`
}

type KafkaAutoRebalanceConfiguration struct {
    Mode KafkaAutoRebalanceMode `json:"mode,omitempty"`
    Template *corev1.LocalObjectReference `json:"template,omitempty"`
}

type CruiseControlApiUsersType string
const (
    HASHLOGINSERVICE_CRUISECONTROLAPIUSERSTYPE CruiseControlApiUsersType = "hashLoginService"
)

type CruiseControlApiUsers struct {
    Type CruiseControlApiUsersType `json:"type,omitempty"`
    ValueFrom *PasswordSource `json:"valueFrom,omitempty"`
}

type BrokerCapacity struct {
    Disk string `json:"disk,omitempty"`
    CpuUtilization int32 `json:"cpuUtilization,omitempty"`
    Cpu string `json:"cpu,omitempty"`
    InboundNetwork string `json:"inboundNetwork,omitempty"`
    OutboundNetwork string `json:"outboundNetwork,omitempty"`
    Overrides []BrokerCapacityOverride `json:"overrides,omitempty"`
}

type BrokerCapacityOverride struct {
    Brokers []int32 `json:"brokers,omitempty"`
    Cpu string `json:"cpu,omitempty"`
    InboundNetwork string `json:"inboundNetwork,omitempty"`
    OutboundNetwork string `json:"outboundNetwork,omitempty"`
}

type CruiseControlTemplate struct {
    Deployment *DeploymentTemplate `json:"deployment,omitempty"`
    Pod *PodTemplate `json:"pod,omitempty"`
    ApiService *InternalServiceTemplate `json:"apiService,omitempty"`
    PodDisruptionBudget *PodDisruptionBudgetTemplate `json:"podDisruptionBudget,omitempty"`
    CruiseControlContainer *ContainerTemplate `json:"cruiseControlContainer,omitempty"`
    TlsSidecarContainer *ContainerTemplate `json:"tlsSidecarContainer,omitempty"`
    ServiceAccount *ResourceTemplate `json:"serviceAccount,omitempty"`
}

type TlsSidecar struct {
    Image string `json:"image,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
    LivenessProbe *Probe `json:"livenessProbe,omitempty"`
    ReadinessProbe *Probe `json:"readinessProbe,omitempty"`
    LogLevel TlsSidecarLogLevel `json:"logLevel,omitempty"`
}

type TlsSidecarLogLevel string
const (
    EMERG_TLSSIDECARLOGLEVEL TlsSidecarLogLevel = "emerg"
    ALERT_TLSSIDECARLOGLEVEL TlsSidecarLogLevel = "alert"
    CRIT_TLSSIDECARLOGLEVEL TlsSidecarLogLevel = "crit"
    ERR_TLSSIDECARLOGLEVEL TlsSidecarLogLevel = "err"
    WARNING_TLSSIDECARLOGLEVEL TlsSidecarLogLevel = "warning"
    NOTICE_TLSSIDECARLOGLEVEL TlsSidecarLogLevel = "notice"
    INFO_TLSSIDECARLOGLEVEL TlsSidecarLogLevel = "info"
    DEBUG_TLSSIDECARLOGLEVEL TlsSidecarLogLevel = "debug"
)

type CertificateAuthority struct {
    GenerateCertificateAuthority bool `json:"generateCertificateAuthority,omitempty"`
    GenerateSecretOwnerReference bool `json:"generateSecretOwnerReference,omitempty"`
    ValidityDays int32 `json:"validityDays,omitempty"`
    RenewalDays int32 `json:"renewalDays,omitempty"`
    CertificateExpirationPolicy CertificateExpirationPolicy `json:"certificateExpirationPolicy,omitempty"`
}

type CertificateExpirationPolicy string
const (
    RENEW_CERTIFICATE_CERTIFICATEEXPIRATIONPOLICY CertificateExpirationPolicy = "renew-certificate"
    REPLACE_KEY_CERTIFICATEEXPIRATIONPOLICY CertificateExpirationPolicy = "replace-key"
)

type EntityOperatorSpec struct {
    TopicOperator *EntityTopicOperatorSpec `json:"topicOperator,omitempty"`
    UserOperator *EntityUserOperatorSpec `json:"userOperator,omitempty"`
    TlsSidecar *TlsSidecar `json:"tlsSidecar,omitempty"`
    Template *EntityOperatorTemplate `json:"template,omitempty"`
}

type EntityOperatorTemplate struct {
    Deployment *DeploymentTemplate `json:"deployment,omitempty"`
    Pod *PodTemplate `json:"pod,omitempty"`
    TopicOperatorContainer *ContainerTemplate `json:"topicOperatorContainer,omitempty"`
    UserOperatorContainer *ContainerTemplate `json:"userOperatorContainer,omitempty"`
    TlsSidecarContainer *ContainerTemplate `json:"tlsSidecarContainer,omitempty"`
    ServiceAccount *ResourceTemplate `json:"serviceAccount,omitempty"`
    EntityOperatorRole *ResourceTemplate `json:"entityOperatorRole,omitempty"`
    TopicOperatorRoleBinding *ResourceTemplate `json:"topicOperatorRoleBinding,omitempty"`
    UserOperatorRoleBinding *ResourceTemplate `json:"userOperatorRoleBinding,omitempty"`
}

type EntityUserOperatorSpec struct {
    WatchedNamespace string `json:"watchedNamespace,omitempty"`
    Image string `json:"image,omitempty"`
    ReconciliationIntervalSeconds int64 `json:"reconciliationIntervalSeconds,omitempty"`
    ReconciliationIntervalMs int64 `json:"reconciliationIntervalMs,omitempty"`
    ZookeeperSessionTimeoutSeconds int64 `json:"zookeeperSessionTimeoutSeconds,omitempty"`
    SecretPrefix string `json:"secretPrefix,omitempty"`
    LivenessProbe *Probe `json:"livenessProbe,omitempty"`
    ReadinessProbe *Probe `json:"readinessProbe,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
    Logging *Logging `json:"logging,omitempty"`
    JvmOptions *JvmOptions `json:"jvmOptions,omitempty"`
}

type EntityTopicOperatorSpec struct {
    WatchedNamespace string `json:"watchedNamespace,omitempty"`
    Image string `json:"image,omitempty"`
    ReconciliationIntervalSeconds int32 `json:"reconciliationIntervalSeconds,omitempty"`
    ReconciliationIntervalMs int64 `json:"reconciliationIntervalMs,omitempty"`
    ZookeeperSessionTimeoutSeconds int32 `json:"zookeeperSessionTimeoutSeconds,omitempty"`
    StartupProbe *Probe `json:"startupProbe,omitempty"`
    LivenessProbe *Probe `json:"livenessProbe,omitempty"`
    ReadinessProbe *Probe `json:"readinessProbe,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
    TopicMetadataMaxAttempts int32 `json:"topicMetadataMaxAttempts,omitempty"`
    Logging *Logging `json:"logging,omitempty"`
    JvmOptions *JvmOptions `json:"jvmOptions,omitempty"`
}

type ZookeeperClusterSpec struct {
    Replicas int32 `json:"replicas,omitempty"`
    Image string `json:"image,omitempty"`
    Storage *SingleVolumeStorage `json:"storage,omitempty"`
    Config JSONValue `json:"config,omitempty"`
    LivenessProbe *Probe `json:"livenessProbe,omitempty"`
    ReadinessProbe *Probe `json:"readinessProbe,omitempty"`
    JvmOptions *JvmOptions `json:"jvmOptions,omitempty"`
    JmxOptions *KafkaJmxOptions `json:"jmxOptions,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
    MetricsConfig *MetricsConfig `json:"metricsConfig,omitempty"`
    Logging *Logging `json:"logging,omitempty"`
    Template *ZookeeperClusterTemplate `json:"template,omitempty"`
}

type ZookeeperClusterTemplate struct {
    Statefulset *StatefulSetTemplate `json:"statefulset,omitempty"`
    PodSet *ResourceTemplate `json:"podSet,omitempty"`
    Pod *PodTemplate `json:"pod,omitempty"`
    ClientService *InternalServiceTemplate `json:"clientService,omitempty"`
    NodesService *InternalServiceTemplate `json:"nodesService,omitempty"`
    PersistentVolumeClaim *ResourceTemplate `json:"persistentVolumeClaim,omitempty"`
    PodDisruptionBudget *PodDisruptionBudgetTemplate `json:"podDisruptionBudget,omitempty"`
    ZookeeperContainer *ContainerTemplate `json:"zookeeperContainer,omitempty"`
    ServiceAccount *ResourceTemplate `json:"serviceAccount,omitempty"`
    JmxSecret *ResourceTemplate `json:"jmxSecret,omitempty"`
}

type StatefulSetTemplate struct {
    Metadata *MetadataTemplate `json:"metadata,omitempty"`
    PodManagementPolicy PodManagementPolicy `json:"podManagementPolicy,omitempty"`
}

type PodManagementPolicy string
const (
    ORDERED_READY_PODMANAGEMENTPOLICY PodManagementPolicy = "OrderedReady"
    PARALLEL_PODMANAGEMENTPOLICY PodManagementPolicy = "Parallel"
)

type KafkaClusterSpec struct {
    Version string `json:"version,omitempty"`
    MetadataVersion string `json:"metadataVersion,omitempty"`
    Replicas int32 `json:"replicas,omitempty"`
    Image string `json:"image,omitempty"`
    Listeners []GenericKafkaListener `json:"listeners,omitempty"`
    Config JSONValue `json:"config,omitempty"`
    Storage *Storage `json:"storage,omitempty"`
    Authorization *KafkaAuthorization `json:"authorization,omitempty"`
    Rack *Rack `json:"rack,omitempty"`
    BrokerRackInitImage string `json:"brokerRackInitImage,omitempty"`
    LivenessProbe *Probe `json:"livenessProbe,omitempty"`
    ReadinessProbe *Probe `json:"readinessProbe,omitempty"`
    JvmOptions *JvmOptions `json:"jvmOptions,omitempty"`
    JmxOptions *KafkaJmxOptions `json:"jmxOptions,omitempty"`
    Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
    MetricsConfig *MetricsConfig `json:"metricsConfig,omitempty"`
    Logging *Logging `json:"logging,omitempty"`
    Template *KafkaClusterTemplate `json:"template,omitempty"`
    TieredStorage *TieredStorage `json:"tieredStorage,omitempty"`
    Quotas *QuotasPlugin `json:"quotas,omitempty"`
}

type QuotasPluginType string
const (
    KAFKA_QUOTASPLUGINTYPE QuotasPluginType = "kafka"
    STRIMZI_QUOTASPLUGINTYPE QuotasPluginType = "strimzi"
)

type QuotasPlugin struct {
    ProducerByteRate int64 `json:"producerByteRate,omitempty"`
    ConsumerByteRate int64 `json:"consumerByteRate,omitempty"`
    RequestPercentage int32 `json:"requestPercentage,omitempty"`
    MinAvailableRatioPerVolume float64 `json:"minAvailableRatioPerVolume,omitempty"`
    MinAvailableBytesPerVolume int64 `json:"minAvailableBytesPerVolume,omitempty"`
    Type QuotasPluginType `json:"type,omitempty"`
    ControllerMutationRate float64 `json:"controllerMutationRate,omitempty"`
    ExcludedPrincipals []string `json:"excludedPrincipals,omitempty"`
}

type TieredStorageType string
const (
    CUSTOM_TIEREDSTORAGETYPE TieredStorageType = "custom"
)

type TieredStorage struct {
    Type TieredStorageType `json:"type,omitempty"`
    RemoteStorageManager *RemoteStorageManager `json:"remoteStorageManager,omitempty"`
}

type RemoteStorageManager struct {
    ClassName string `json:"className,omitempty"`
    ClassPath string `json:"classPath,omitempty"`
    Config map[string]string `json:"config,omitempty"`
}

type KafkaClusterTemplate struct {
    Statefulset *StatefulSetTemplate `json:"statefulset,omitempty"`
    Pod *PodTemplate `json:"pod,omitempty"`
    BootstrapService *InternalServiceTemplate `json:"bootstrapService,omitempty"`
    BrokersService *InternalServiceTemplate `json:"brokersService,omitempty"`
    ExternalBootstrapService *ResourceTemplate `json:"externalBootstrapService,omitempty"`
    PerPodService *ResourceTemplate `json:"perPodService,omitempty"`
    ExternalBootstrapRoute *ResourceTemplate `json:"externalBootstrapRoute,omitempty"`
    PerPodRoute *ResourceTemplate `json:"perPodRoute,omitempty"`
    ExternalBootstrapIngress *ResourceTemplate `json:"externalBootstrapIngress,omitempty"`
    PerPodIngress *ResourceTemplate `json:"perPodIngress,omitempty"`
    PersistentVolumeClaim *ResourceTemplate `json:"persistentVolumeClaim,omitempty"`
    PodDisruptionBudget *PodDisruptionBudgetTemplate `json:"podDisruptionBudget,omitempty"`
    KafkaContainer *ContainerTemplate `json:"kafkaContainer,omitempty"`
    InitContainer *ContainerTemplate `json:"initContainer,omitempty"`
    ClusterCaCert *ResourceTemplate `json:"clusterCaCert,omitempty"`
    ServiceAccount *ResourceTemplate `json:"serviceAccount,omitempty"`
    JmxSecret *ResourceTemplate `json:"jmxSecret,omitempty"`
    ClusterRoleBinding *ResourceTemplate `json:"clusterRoleBinding,omitempty"`
    PodSet *ResourceTemplate `json:"podSet,omitempty"`
}

type KafkaAuthorizationType string
const (
    SIMPLE_KAFKAAUTHORIZATIONTYPE KafkaAuthorizationType = "simple"
    OPA_KAFKAAUTHORIZATIONTYPE KafkaAuthorizationType = "opa"
    KEYCLOAK_KAFKAAUTHORIZATIONTYPE KafkaAuthorizationType = "keycloak"
    CUSTOM_KAFKAAUTHORIZATIONTYPE KafkaAuthorizationType = "custom"
)

type KafkaAuthorization struct {
    GrantsRefreshPeriodSeconds int32 `json:"grantsRefreshPeriodSeconds,omitempty"`
    AllowOnError bool `json:"allowOnError,omitempty"`
    ClientId string `json:"clientId,omitempty"`
    SuperUsers []string `json:"superUsers,omitempty"`
    GrantsMaxIdleTimeSeconds int32 `json:"grantsMaxIdleTimeSeconds,omitempty"`
    InitialCacheCapacity int32 `json:"initialCacheCapacity,omitempty"`
    AuthorizerClass string `json:"authorizerClass,omitempty"`
    ConnectTimeoutSeconds int32 `json:"connectTimeoutSeconds,omitempty"`
    ExpireAfterMs int64 `json:"expireAfterMs,omitempty"`
    Type KafkaAuthorizationType `json:"type,omitempty"`
    SupportsAdminApi bool `json:"supportsAdminApi,omitempty"`
    DelegateToKafkaAcls bool `json:"delegateToKafkaAcls,omitempty"`
    GrantsRefreshPoolSize int32 `json:"grantsRefreshPoolSize,omitempty"`
    Url string `json:"url,omitempty"`
    MaximumCacheSize int32 `json:"maximumCacheSize,omitempty"`
    ReadTimeoutSeconds int32 `json:"readTimeoutSeconds,omitempty"`
    TlsTrustedCertificates []CertSecretSource `json:"tlsTrustedCertificates,omitempty"`
    EnableMetrics bool `json:"enableMetrics,omitempty"`
    GrantsAlwaysLatest bool `json:"grantsAlwaysLatest,omitempty"`
    GrantsGcPeriodSeconds int32 `json:"grantsGcPeriodSeconds,omitempty"`
    TokenEndpointUri string `json:"tokenEndpointUri,omitempty"`
    DisableTlsHostnameVerification bool `json:"disableTlsHostnameVerification,omitempty"`
    IncludeAcceptHeader bool `json:"includeAcceptHeader,omitempty"`
    HttpRetries int32 `json:"httpRetries,omitempty"`
}

type GenericKafkaListener struct {
    Name string `json:"name,omitempty"`
    Port int32 `json:"port,omitempty"`
    Type KafkaListenerType `json:"type,omitempty"`
    Tls bool `json:"tls,omitempty"`
    Authentication *KafkaListenerAuthentication `json:"authentication,omitempty"`
    Configuration *GenericKafkaListenerConfiguration `json:"configuration,omitempty"`
    NetworkPolicyPeers []networkingv1.NetworkPolicyPeer `json:"networkPolicyPeers,omitempty"`
}

type GenericKafkaListenerConfiguration struct {
    BrokerCertChainAndKey *CertAndKeySecretSource `json:"brokerCertChainAndKey,omitempty"`
    Class string `json:"class,omitempty"`
    ExternalTrafficPolicy ExternalTrafficPolicy `json:"externalTrafficPolicy,omitempty"`
    LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges,omitempty"`
    Bootstrap *GenericKafkaListenerConfigurationBootstrap `json:"bootstrap,omitempty"`
    Brokers []GenericKafkaListenerConfigurationBroker `json:"brokers,omitempty"`
    IpFamilyPolicy IpFamilyPolicy `json:"ipFamilyPolicy,omitempty"`
    IpFamilies []IpFamily `json:"ipFamilies,omitempty"`
    CreateBootstrapService bool `json:"createBootstrapService,omitempty"`
    Finalizers []string `json:"finalizers,omitempty"`
    UseServiceDnsDomain bool `json:"useServiceDnsDomain,omitempty"`
    MaxConnections int32 `json:"maxConnections,omitempty"`
    MaxConnectionCreationRate int32 `json:"maxConnectionCreationRate,omitempty"`
    PreferredNodePortAddressType NodeAddressType `json:"preferredNodePortAddressType,omitempty"`
    PublishNotReadyAddresses bool `json:"publishNotReadyAddresses,omitempty"`
    HostTemplate string `json:"hostTemplate,omitempty"`
    AdvertisedHostTemplate string `json:"advertisedHostTemplate,omitempty"`
    AllocateLoadBalancerNodePorts bool `json:"allocateLoadBalancerNodePorts,omitempty"`
}

type NodeAddressType string
const (
    EXTERNAL_IP_NODEADDRESSTYPE NodeAddressType = "ExternalIP"
    EXTERNAL_DNS_NODEADDRESSTYPE NodeAddressType = "ExternalDNS"
    INTERNAL_IP_NODEADDRESSTYPE NodeAddressType = "InternalIP"
    INTERNAL_DNS_NODEADDRESSTYPE NodeAddressType = "InternalDNS"
    HOSTNAME_NODEADDRESSTYPE NodeAddressType = "Hostname"
)

type GenericKafkaListenerConfigurationBroker struct {
    Broker int32 `json:"broker,omitempty"`
    AdvertisedHost string `json:"advertisedHost,omitempty"`
    AdvertisedPort int32 `json:"advertisedPort,omitempty"`
    Host string `json:"host,omitempty"`
    NodePort int32 `json:"nodePort,omitempty"`
    LoadBalancerIP string `json:"loadBalancerIP,omitempty"`
    Annotations map[string]string `json:"annotations,omitempty"`
    Labels map[string]string `json:"labels,omitempty"`
    ExternalIPs []string `json:"externalIPs,omitempty"`
}

type GenericKafkaListenerConfigurationBootstrap struct {
    AlternativeNames []string `json:"alternativeNames,omitempty"`
    Host string `json:"host,omitempty"`
    NodePort int32 `json:"nodePort,omitempty"`
    LoadBalancerIP string `json:"loadBalancerIP,omitempty"`
    Annotations map[string]string `json:"annotations,omitempty"`
    Labels map[string]string `json:"labels,omitempty"`
    ExternalIPs []string `json:"externalIPs,omitempty"`
}

type ExternalTrafficPolicy string
const (
    LOCAL_EXTERNALTRAFFICPOLICY ExternalTrafficPolicy = "Local"
    CLUSTER_EXTERNALTRAFFICPOLICY ExternalTrafficPolicy = "Cluster"
)

type KafkaListenerAuthenticationType string
const (
    TLS_KAFKALISTENERAUTHENTICATIONTYPE KafkaListenerAuthenticationType = "tls"
    SCRAM_SHA_512_KAFKALISTENERAUTHENTICATIONTYPE KafkaListenerAuthenticationType = "scram-sha-512"
    OAUTH_KAFKALISTENERAUTHENTICATIONTYPE KafkaListenerAuthenticationType = "oauth"
    CUSTOM_KAFKALISTENERAUTHENTICATIONTYPE KafkaListenerAuthenticationType = "custom"
)

type KafkaListenerAuthentication struct {
    JwksMinRefreshPauseSeconds int32 `json:"jwksMinRefreshPauseSeconds,omitempty"`
    EnableECDSA bool `json:"enableECDSA,omitempty"`
    IntrospectionEndpointUri string `json:"introspectionEndpointUri,omitempty"`
    ValidIssuerUri string `json:"validIssuerUri,omitempty"`
    ValidTokenType string `json:"validTokenType,omitempty"`
    ListenerConfig JSONValue `json:"listenerConfig,omitempty"`
    Type KafkaListenerAuthenticationType `json:"type,omitempty"`
    UserNamePrefix string `json:"userNamePrefix,omitempty"`
    FallbackUserNamePrefix string `json:"fallbackUserNamePrefix,omitempty"`
    UserInfoEndpointUri string `json:"userInfoEndpointUri,omitempty"`
    AccessTokenIsJwt bool `json:"accessTokenIsJwt,omitempty"`
    FallbackUserNameClaim string `json:"fallbackUserNameClaim,omitempty"`
    TlsTrustedCertificates []CertSecretSource `json:"tlsTrustedCertificates,omitempty"`
    Sasl bool `json:"sasl,omitempty"`
    CheckIssuer bool `json:"checkIssuer,omitempty"`
    DisableTlsHostnameVerification bool `json:"disableTlsHostnameVerification,omitempty"`
    JwksIgnoreKeyUse bool `json:"jwksIgnoreKeyUse,omitempty"`
    ServerBearerTokenLocation string `json:"serverBearerTokenLocation,omitempty"`
    JwksEndpointUri string `json:"jwksEndpointUri,omitempty"`
    ClientSecret *GenericSecretSource `json:"clientSecret,omitempty"`
    ClientAudience string `json:"clientAudience,omitempty"`
    JwksExpirySeconds int32 `json:"jwksExpirySeconds,omitempty"`
    JwksRefreshSeconds int32 `json:"jwksRefreshSeconds,omitempty"`
    EnableOauthBearer bool `json:"enableOauthBearer,omitempty"`
    ClientId string `json:"clientId,omitempty"`
    GroupsClaimDelimiter string `json:"groupsClaimDelimiter,omitempty"`
    ConnectTimeoutSeconds int32 `json:"connectTimeoutSeconds,omitempty"`
    UserNameClaim string `json:"userNameClaim,omitempty"`
    HttpRetryPauseMs int32 `json:"httpRetryPauseMs,omitempty"`
    CustomClaimCheck string `json:"customClaimCheck,omitempty"`
    Secrets []GenericSecretSource `json:"secrets,omitempty"`
    FailFast bool `json:"failFast,omitempty"`
    GroupsClaim string `json:"groupsClaim,omitempty"`
    ReadTimeoutSeconds int32 `json:"readTimeoutSeconds,omitempty"`
    EnableMetrics bool `json:"enableMetrics,omitempty"`
    MaxSecondsWithoutReauthentication int32 `json:"maxSecondsWithoutReauthentication,omitempty"`
    TokenEndpointUri string `json:"tokenEndpointUri,omitempty"`
    IncludeAcceptHeader bool `json:"includeAcceptHeader,omitempty"`
    HttpRetries int32 `json:"httpRetries,omitempty"`
    EnablePlain bool `json:"enablePlain,omitempty"`
    CheckAccessTokenType bool `json:"checkAccessTokenType,omitempty"`
    ClientScope string `json:"clientScope,omitempty"`
    CheckAudience bool `json:"checkAudience,omitempty"`
}

type KafkaListenerType string
const (
    INTERNAL_KAFKALISTENERTYPE KafkaListenerType = "internal"
    ROUTE_KAFKALISTENERTYPE KafkaListenerType = "route"
    LOADBALANCER_KAFKALISTENERTYPE KafkaListenerType = "loadbalancer"
    NODEPORT_KAFKALISTENERTYPE KafkaListenerType = "nodeport"
    INGRESS_KAFKALISTENERTYPE KafkaListenerType = "ingress"
    CLUSTER_IP_KAFKALISTENERTYPE KafkaListenerType = "cluster-ip"
)
