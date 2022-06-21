package workflow

type ConditionReason string

// TODO move 'ConditionReason' to 'api' package?

// General reasons
const (
	AtlasCredentialsNotProvided ConditionReason = "AtlasCredentialsNotProvided"
	Internal                    ConditionReason = "InternalError"
)

// Atlas Project reasons
const (
	ProjectNotCreatedInAtlas                ConditionReason = "ProjectNotCreatedInAtlas"
	ProjectIPAccessInvalid                  ConditionReason = "ProjectIPAccessListInvalid"
	ProjectIPNotCreatedInAtlas              ConditionReason = "ProjectIPAccessListNotCreatedInAtlas"
	ProjectWindowInvalid                    ConditionReason = "ProjectWindowInvalid"
	ProjectWindowNotCreatedInAtlas          ConditionReason = "ProjectWindowNotCreatedInAtlas"
	ProjectWindowNotDeletedInAtlas          ConditionReason = "projectWindowNotDeletedInAtlas"
	ProjectWindowNotDeferredInAtlas         ConditionReason = "ProjectWindowNotDeferredInAtlas"
	ProjectWindowNotAutoDeferredInAtlas     ConditionReason = "ProjectWindowNotAutoDeferredInAtlas"
	ProjectPEServiceIsNotReadyInAtlas       ConditionReason = "ProjectPrivateEndpointServiceIsNotReadyInAtlas"
	ProjectPrivateEndpointIsNotReadyInAtlas ConditionReason = "ProjectPrivateEndpointIsNotReadyInAtlas"
	ProjectIPAccessListNotActive            ConditionReason = "ProjectIPAccessListNotActive"
	ProjectIntegrationInternal              ConditionReason = "ProjectIntegrationInternalError"
	ProjectIntegrationRequest               ConditionReason = "ProjectIntegrationRequestError"
	ProjectIntegrationReady                 ConditionReason = "ProjectIntegrationReady"
)

// Atlas Cluster reasons
const (
	ClusterNotCreatedInAtlas           ConditionReason = "ClusterNotCreatedInAtlas"
	ClusterNotUpdatedInAtlas           ConditionReason = "ClusterNotUpdatedInAtlas"
	ClusterCreating                    ConditionReason = "ClusterCreating"
	ClusterUpdating                    ConditionReason = "ClusterUpdating"
	ClusterConnectionSecretsNotCreated ConditionReason = "ClusterConnectionSecretsNotCreated"
	ClusterAdvancedOptionsAreNotReady  ConditionReason = "ClusterAdvancedOptionsAreNotReady"
)

// Atlas Database User reasons
const (
	DatabaseUserNotCreatedInAtlas           ConditionReason = "DatabaseUserNotCreatedInAtlas"
	DatabaseUserNotUpdatedInAtlas           ConditionReason = "DatabaseUserNotUpdatedInAtlas"
	DatabaseUserConnectionSecretsNotCreated ConditionReason = "DatabaseUserConnectionSecretsNotCreated"
	DatabaseUserStaleConnectionSecrets      ConditionReason = "DatabaseUserStaleConnectionSecrets"
	DatabaseUserClustersAppliedChanges      ConditionReason = "ClustersAppliedDatabaseUsersChanges"
	DatabaseUserInvalidSpec                 ConditionReason = "DatabaseUserInvalidSpec"
	DatabaseUserExpired                     ConditionReason = "DatabaseUserExpired"
)
