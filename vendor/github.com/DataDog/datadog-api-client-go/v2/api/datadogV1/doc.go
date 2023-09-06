// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

// List of APIs:
//   - [AWSIntegrationApi.CreateAWSAccount]
//   - [AWSIntegrationApi.CreateAWSTagFilter]
//   - [AWSIntegrationApi.CreateNewAWSExternalID]
//   - [AWSIntegrationApi.DeleteAWSAccount]
//   - [AWSIntegrationApi.DeleteAWSTagFilter]
//   - [AWSIntegrationApi.ListAWSAccounts]
//   - [AWSIntegrationApi.ListAWSTagFilters]
//   - [AWSIntegrationApi.ListAvailableAWSNamespaces]
//   - [AWSIntegrationApi.UpdateAWSAccount]
//   - [AWSLogsIntegrationApi.CheckAWSLogsLambdaAsync]
//   - [AWSLogsIntegrationApi.CheckAWSLogsServicesAsync]
//   - [AWSLogsIntegrationApi.CreateAWSLambdaARN]
//   - [AWSLogsIntegrationApi.DeleteAWSLambdaARN]
//   - [AWSLogsIntegrationApi.EnableAWSLogServices]
//   - [AWSLogsIntegrationApi.ListAWSLogsIntegrations]
//   - [AWSLogsIntegrationApi.ListAWSLogsServices]
//   - [AuthenticationApi.Validate]
//   - [AzureIntegrationApi.CreateAzureIntegration]
//   - [AzureIntegrationApi.DeleteAzureIntegration]
//   - [AzureIntegrationApi.ListAzureIntegration]
//   - [AzureIntegrationApi.UpdateAzureHostFilters]
//   - [AzureIntegrationApi.UpdateAzureIntegration]
//   - [DashboardListsApi.CreateDashboardList]
//   - [DashboardListsApi.DeleteDashboardList]
//   - [DashboardListsApi.GetDashboardList]
//   - [DashboardListsApi.ListDashboardLists]
//   - [DashboardListsApi.UpdateDashboardList]
//   - [DashboardsApi.CreateDashboard]
//   - [DashboardsApi.CreatePublicDashboard]
//   - [DashboardsApi.DeleteDashboard]
//   - [DashboardsApi.DeleteDashboards]
//   - [DashboardsApi.DeletePublicDashboard]
//   - [DashboardsApi.DeletePublicDashboardInvitation]
//   - [DashboardsApi.GetDashboard]
//   - [DashboardsApi.GetPublicDashboard]
//   - [DashboardsApi.GetPublicDashboardInvitations]
//   - [DashboardsApi.ListDashboards]
//   - [DashboardsApi.RestoreDashboards]
//   - [DashboardsApi.SendPublicDashboardInvitation]
//   - [DashboardsApi.UpdateDashboard]
//   - [DashboardsApi.UpdatePublicDashboard]
//   - [DowntimesApi.CancelDowntime]
//   - [DowntimesApi.CancelDowntimesByScope]
//   - [DowntimesApi.CreateDowntime]
//   - [DowntimesApi.GetDowntime]
//   - [DowntimesApi.ListDowntimes]
//   - [DowntimesApi.ListMonitorDowntimes]
//   - [DowntimesApi.UpdateDowntime]
//   - [EventsApi.CreateEvent]
//   - [EventsApi.GetEvent]
//   - [EventsApi.ListEvents]
//   - [GCPIntegrationApi.CreateGCPIntegration]
//   - [GCPIntegrationApi.DeleteGCPIntegration]
//   - [GCPIntegrationApi.ListGCPIntegration]
//   - [GCPIntegrationApi.UpdateGCPIntegration]
//   - [HostsApi.GetHostTotals]
//   - [HostsApi.ListHosts]
//   - [HostsApi.MuteHost]
//   - [HostsApi.UnmuteHost]
//   - [IPRangesApi.GetIPRanges]
//   - [KeyManagementApi.CreateAPIKey]
//   - [KeyManagementApi.CreateApplicationKey]
//   - [KeyManagementApi.DeleteAPIKey]
//   - [KeyManagementApi.DeleteApplicationKey]
//   - [KeyManagementApi.GetAPIKey]
//   - [KeyManagementApi.GetApplicationKey]
//   - [KeyManagementApi.ListAPIKeys]
//   - [KeyManagementApi.ListApplicationKeys]
//   - [KeyManagementApi.UpdateAPIKey]
//   - [KeyManagementApi.UpdateApplicationKey]
//   - [LogsApi.ListLogs]
//   - [LogsApi.SubmitLog]
//   - [LogsIndexesApi.CreateLogsIndex]
//   - [LogsIndexesApi.GetLogsIndex]
//   - [LogsIndexesApi.GetLogsIndexOrder]
//   - [LogsIndexesApi.ListLogIndexes]
//   - [LogsIndexesApi.UpdateLogsIndex]
//   - [LogsIndexesApi.UpdateLogsIndexOrder]
//   - [LogsPipelinesApi.CreateLogsPipeline]
//   - [LogsPipelinesApi.DeleteLogsPipeline]
//   - [LogsPipelinesApi.GetLogsPipeline]
//   - [LogsPipelinesApi.GetLogsPipelineOrder]
//   - [LogsPipelinesApi.ListLogsPipelines]
//   - [LogsPipelinesApi.UpdateLogsPipeline]
//   - [LogsPipelinesApi.UpdateLogsPipelineOrder]
//   - [MetricsApi.GetMetricMetadata]
//   - [MetricsApi.ListActiveMetrics]
//   - [MetricsApi.ListMetrics]
//   - [MetricsApi.QueryMetrics]
//   - [MetricsApi.SubmitDistributionPoints]
//   - [MetricsApi.SubmitMetrics]
//   - [MetricsApi.UpdateMetricMetadata]
//   - [MonitorsApi.CheckCanDeleteMonitor]
//   - [MonitorsApi.CreateMonitor]
//   - [MonitorsApi.DeleteMonitor]
//   - [MonitorsApi.GetMonitor]
//   - [MonitorsApi.ListMonitors]
//   - [MonitorsApi.SearchMonitorGroups]
//   - [MonitorsApi.SearchMonitors]
//   - [MonitorsApi.UpdateMonitor]
//   - [MonitorsApi.ValidateExistingMonitor]
//   - [MonitorsApi.ValidateMonitor]
//   - [NotebooksApi.CreateNotebook]
//   - [NotebooksApi.DeleteNotebook]
//   - [NotebooksApi.GetNotebook]
//   - [NotebooksApi.ListNotebooks]
//   - [NotebooksApi.UpdateNotebook]
//   - [OrganizationsApi.CreateChildOrg]
//   - [OrganizationsApi.DowngradeOrg]
//   - [OrganizationsApi.GetOrg]
//   - [OrganizationsApi.ListOrgs]
//   - [OrganizationsApi.UpdateOrg]
//   - [OrganizationsApi.UploadIdPForOrg]
//   - [PagerDutyIntegrationApi.CreatePagerDutyIntegrationService]
//   - [PagerDutyIntegrationApi.DeletePagerDutyIntegrationService]
//   - [PagerDutyIntegrationApi.GetPagerDutyIntegrationService]
//   - [PagerDutyIntegrationApi.UpdatePagerDutyIntegrationService]
//   - [SecurityMonitoringApi.AddSecurityMonitoringSignalToIncident]
//   - [SecurityMonitoringApi.EditSecurityMonitoringSignalAssignee]
//   - [SecurityMonitoringApi.EditSecurityMonitoringSignalState]
//   - [ServiceChecksApi.SubmitServiceCheck]
//   - [ServiceLevelObjectiveCorrectionsApi.CreateSLOCorrection]
//   - [ServiceLevelObjectiveCorrectionsApi.DeleteSLOCorrection]
//   - [ServiceLevelObjectiveCorrectionsApi.GetSLOCorrection]
//   - [ServiceLevelObjectiveCorrectionsApi.ListSLOCorrection]
//   - [ServiceLevelObjectiveCorrectionsApi.UpdateSLOCorrection]
//   - [ServiceLevelObjectivesApi.CheckCanDeleteSLO]
//   - [ServiceLevelObjectivesApi.CreateSLO]
//   - [ServiceLevelObjectivesApi.DeleteSLO]
//   - [ServiceLevelObjectivesApi.DeleteSLOTimeframeInBulk]
//   - [ServiceLevelObjectivesApi.GetSLO]
//   - [ServiceLevelObjectivesApi.GetSLOCorrections]
//   - [ServiceLevelObjectivesApi.GetSLOHistory]
//   - [ServiceLevelObjectivesApi.ListSLOs]
//   - [ServiceLevelObjectivesApi.SearchSLO]
//   - [ServiceLevelObjectivesApi.UpdateSLO]
//   - [SlackIntegrationApi.CreateSlackIntegrationChannel]
//   - [SlackIntegrationApi.GetSlackIntegrationChannel]
//   - [SlackIntegrationApi.GetSlackIntegrationChannels]
//   - [SlackIntegrationApi.RemoveSlackIntegrationChannel]
//   - [SlackIntegrationApi.UpdateSlackIntegrationChannel]
//   - [SnapshotsApi.GetGraphSnapshot]
//   - [SyntheticsApi.CreateGlobalVariable]
//   - [SyntheticsApi.CreatePrivateLocation]
//   - [SyntheticsApi.CreateSyntheticsAPITest]
//   - [SyntheticsApi.CreateSyntheticsBrowserTest]
//   - [SyntheticsApi.DeleteGlobalVariable]
//   - [SyntheticsApi.DeletePrivateLocation]
//   - [SyntheticsApi.DeleteTests]
//   - [SyntheticsApi.EditGlobalVariable]
//   - [SyntheticsApi.GetAPITest]
//   - [SyntheticsApi.GetAPITestLatestResults]
//   - [SyntheticsApi.GetAPITestResult]
//   - [SyntheticsApi.GetBrowserTest]
//   - [SyntheticsApi.GetBrowserTestLatestResults]
//   - [SyntheticsApi.GetBrowserTestResult]
//   - [SyntheticsApi.GetGlobalVariable]
//   - [SyntheticsApi.GetPrivateLocation]
//   - [SyntheticsApi.GetSyntheticsCIBatch]
//   - [SyntheticsApi.GetSyntheticsDefaultLocations]
//   - [SyntheticsApi.GetTest]
//   - [SyntheticsApi.ListGlobalVariables]
//   - [SyntheticsApi.ListLocations]
//   - [SyntheticsApi.ListTests]
//   - [SyntheticsApi.TriggerCITests]
//   - [SyntheticsApi.TriggerTests]
//   - [SyntheticsApi.UpdateAPITest]
//   - [SyntheticsApi.UpdateBrowserTest]
//   - [SyntheticsApi.UpdatePrivateLocation]
//   - [SyntheticsApi.UpdateTestPauseStatus]
//   - [TagsApi.CreateHostTags]
//   - [TagsApi.DeleteHostTags]
//   - [TagsApi.GetHostTags]
//   - [TagsApi.ListHostTags]
//   - [TagsApi.UpdateHostTags]
//   - [UsageMeteringApi.GetDailyCustomReports]
//   - [UsageMeteringApi.GetHourlyUsageAttribution]
//   - [UsageMeteringApi.GetIncidentManagement]
//   - [UsageMeteringApi.GetIngestedSpans]
//   - [UsageMeteringApi.GetMonthlyCustomReports]
//   - [UsageMeteringApi.GetMonthlyUsageAttribution]
//   - [UsageMeteringApi.GetSpecifiedDailyCustomReports]
//   - [UsageMeteringApi.GetSpecifiedMonthlyCustomReports]
//   - [UsageMeteringApi.GetUsageAnalyzedLogs]
//   - [UsageMeteringApi.GetUsageAttribution]
//   - [UsageMeteringApi.GetUsageAuditLogs]
//   - [UsageMeteringApi.GetUsageBillableSummary]
//   - [UsageMeteringApi.GetUsageCIApp]
//   - [UsageMeteringApi.GetUsageCWS]
//   - [UsageMeteringApi.GetUsageCloudSecurityPostureManagement]
//   - [UsageMeteringApi.GetUsageDBM]
//   - [UsageMeteringApi.GetUsageFargate]
//   - [UsageMeteringApi.GetUsageHosts]
//   - [UsageMeteringApi.GetUsageIndexedSpans]
//   - [UsageMeteringApi.GetUsageInternetOfThings]
//   - [UsageMeteringApi.GetUsageLambda]
//   - [UsageMeteringApi.GetUsageLogs]
//   - [UsageMeteringApi.GetUsageLogsByIndex]
//   - [UsageMeteringApi.GetUsageLogsByRetention]
//   - [UsageMeteringApi.GetUsageNetworkFlows]
//   - [UsageMeteringApi.GetUsageNetworkHosts]
//   - [UsageMeteringApi.GetUsageOnlineArchive]
//   - [UsageMeteringApi.GetUsageProfiling]
//   - [UsageMeteringApi.GetUsageRumSessions]
//   - [UsageMeteringApi.GetUsageRumUnits]
//   - [UsageMeteringApi.GetUsageSDS]
//   - [UsageMeteringApi.GetUsageSNMP]
//   - [UsageMeteringApi.GetUsageSummary]
//   - [UsageMeteringApi.GetUsageSynthetics]
//   - [UsageMeteringApi.GetUsageSyntheticsAPI]
//   - [UsageMeteringApi.GetUsageSyntheticsBrowser]
//   - [UsageMeteringApi.GetUsageTimeseries]
//   - [UsageMeteringApi.GetUsageTopAvgMetrics]
//   - [UsersApi.CreateUser]
//   - [UsersApi.DisableUser]
//   - [UsersApi.GetUser]
//   - [UsersApi.ListUsers]
//   - [UsersApi.UpdateUser]
//   - [WebhooksIntegrationApi.CreateWebhooksIntegration]
//   - [WebhooksIntegrationApi.CreateWebhooksIntegrationCustomVariable]
//   - [WebhooksIntegrationApi.DeleteWebhooksIntegration]
//   - [WebhooksIntegrationApi.DeleteWebhooksIntegrationCustomVariable]
//   - [WebhooksIntegrationApi.GetWebhooksIntegration]
//   - [WebhooksIntegrationApi.GetWebhooksIntegrationCustomVariable]
//   - [WebhooksIntegrationApi.UpdateWebhooksIntegration]
//   - [WebhooksIntegrationApi.UpdateWebhooksIntegrationCustomVariable]
package datadogV1
