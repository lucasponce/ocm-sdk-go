# Changes

This document describes the relevant changes between releases of the OCM API
SDK.

## 0.1.1 Sep 08 2026

- ROSAENG-13732 | feat: Implement BYO firewall rule endpoint verification
- chore: bump ocm-api-model to v0.0.465
- Update CONTRIBUTING.md
- chore: bump version to v0.1.509
- chore: bump ocm-api-model to v0.0.464
- ROSAENG-62396 | chore: bump version to v0.1.508
- ROSAENG-61162 | feat: changes to support BYO firewall for OSD-GCP deployments
- chore: bump ocm-api-model to v0.0.463
- ROSAENG-61032 | feat: changes to support SpotMarketOptions for ROSA HCP
- ROSAENG-62396 | fix: automated release github workflow
- Revert "Release 0.1.502"
- ROSAENG-61032 | task: Release 0.1.502
- ROSAENG-61032 | task: ocm-api/sdk changes to support SpotMarketOptions for ROSA HCP
- chore: bump ocm-api-model to v0.0.462
- ROSAENG-59819 | feat: bump ocm-api-model to v0.0.461 and regenerate severity constants
- chore: bump ocm-api-model to v0.0.459
- chore: bump ocm-api-model to v0.0.457
- ci: add automated release pipeline
- OCM-24592 | chore: release ocm-sdk-go v0.1.501
- OCM-24592 | chore: release preparation for v0.0.456
- chore: bump version to 0.1.500
- OCM-23443 | feat: add ZeroEgress into AWS
- Add one-command build/setup: README.md (Quick Start section)
- Add Dependabot configuration for dependency security scanning
- Improve README structure: add missing Installation, Usage, and Development sections
- Add agent-readiness improvements: AGENTS.md, CLAUDE.md, PR/issue templates, .gitignore, pre-commit hooks
- Release v0.1.499
- OCM-22967 | fix: removing pgx dependency
- Release v0.1.498
- chore: update ocm-api-model to 0.0.453
- Release v0.1.497
- OCM-22426 | chore: update ocm-api-model to 0.0.452
- chore: Release 0.1.496
- Bump api model to v0.0.451
- chore: Release 0.1.495
- Bump api model to v0.0.450
- chore: prepare next release 0.1.494
- chore: update ocm api model to v0.0.449
- chore: Release v0.1.493
- chore: update ocm api model to v0.0.448
- chore: Release v0.1.492
- chore: update ocm api model to v0.0.447
- chore: Release v0.1.491
- chore: update ocm api model to v0.0.446
- Release v0.1.490
- Add ExcludedNamespaceSelectors to ingress type
- Release 0.1.489
- OCM-20938 | chore: update ocm-api-model to v0.0.443
- release version 0.1.488
- chore: update ocm-api-model to v0.0.442
- OCM-21643 | chore: update golang-ci-lint version to latest stable v2.8.0
- release 0.1.487
- Revert "OCM-21643 | chore: update golang-ci-lint version to latest stable v2.8.0"
- OCM-21643 | chore: update golang ci lint version
- update API model to 0.0.441
- release v0.1.486
- update API model to 0.0.440
- release 1.485
- add patch operation for version gates to allow partial updates
- Release v0.1.484
- fix lint issues
- Enable the OCM linter
- Add a linter for the ocm logger Linting issues in using ocm logger were not detected by the linter because the logger methods are not ending with `f` and the first parameter is not the string format. This PR adds a custom linter to detect them.
- bump model to v0.0.438
- OCM-20466 | feat: Add new fields to environment
- update ocm-api-model to version 0.0.437
- chore: auto generated files
- update ocm-api-model to version 0.0.436
- Add ROSA and ARO approvers
- Release v0.1.480
- Update ocm api model to v0.0.435
- Release v0.1.479
- Update ocm api model to v0.0.434 ARO-HCP API models split
- feat: add OCM object response helpers for testing framework
- Release v0.1.478
- OCM-689: Propagate RBAC endpoints to OCM-SDK
- fix env var detection
- wif resource scoped permissions
- chore: bump version
- Update CHANGES.md for 0.1.476 Release
- OCM-17788 | feat: image mirrors for rosa hcp
- release v0.1.475
- update API model to v0.0.430
- OCM-16908 | chore: Update version of SDK
- OCM-16908 | chore: Update to v0.0.429 of API model
- Further enhancements to CONTRIBUTING.md reflecting changes in ocm-api-model
- Update version to 0.1.473
- Use ocm-api-model v0.0.426
- update api model to latest
- Reorganize the CONTRIBUTING.md file, remove unecessary instructions and add explenations
- Add a suggestion to run the git push from the root of the repo
- Add git hook to validate a pushed tag matches the Version constant
- Add bump version script and update the version file
- go mod tidy
- Use latest ocm-api-model
- chore: update ocm api metamodel to v0.0.422
- Add an optional commit SHA argument to the update-model script
- Update the release process of the ocm-api-model in the CONTRIBUTING file
- Add a script to simplify current update process
- chore: release v0.1.469
- bump(ocm-api-model): bump the ocm api model to latest commit https://github.com/openshift-online/ocm-api-model/commit/37dc3401307a53d0b4659d5d4f589ad5e15bef5a
- OCM-16570 | chore: update changes and version for release 0.1.468
- OCM-16581 | chore: upgrade golangci lint to v2 and fix lint errors
- leadership flag enhancement: precheck handler added
- generated
- add ocm-api-model/client-api as dep
- add metamodule submodule to have repeatable and externally detectable levels of metamodel
- add new generation and makefile scripts
- Release v0.1.467
- Add Logtype to Notification details request type
- add macgregor, cristianoveiga to owners/approvers
- ARO-17598: generate clients for latest model update
- Update gh setup-python to latest version
- add standard prow artifacts for integration
- chore: bump version to 0.1.465
- OCM-15122 | feat: add `RhRegionID` into `Subscription` and  `ClusterAuthorizationRequest`
- updated third-party dependencies
- Release v0.1.464
- Bump model to version 0.0.416
- Release v0.1.463
- bump ocm-api-model to v0.0.415
- New release v0.1.462
- ARO-15783: Add node pools and node pool status to aro-hcp
- ARO-15553 | Update model and metalmodel
- chore: release 0.1.460
- chore: update ocm-api-model version
- stop retrying generally on EOF and connection reset
- OCM-13888 | fix: Ensure labels for metrics correctly reference service name in path and not api version
- update version
- OCM-13846 | (chore) update model version
- chore: Release v.0.1.458
- ARO-14861: Update model and metamodel
- Release v0.1.457
- fix: initiate context through the initial supplied one
- Update version and release notes
- Update model to 0.0.409
- Bump version
- OCM-12069 | (chore): Update model version
- Bump version
- OCM-12069 | (feat): Add cluster migrations with examples
- bump version to 0.1.453
- chore: bump to ocm-api-model 0.0.406
- Add ARO HCP v1alpha1 root resource and changes required for derived structs
- bump version to 0.1.451
- chore: bump to ocm-api-model 0.0.404
- chore: release v0.1.450
- chore: bump to ocm-api-model 0.0.403
- bump version to 0.1.449
- chore: bump to ocm-api-model 0.0.402
- OCM-12240 | chore: adjust cluster example
- OCM-12240 | chore: Release v0.1.448
- OCM-12240 | feat: include hcp shared vpc support
- chore: release v0.1.447
- chore: synchronise to ocm-api-model 0.0.400
- ARO-7207 | chore: release v0.1.446
- chore: synchronise to ocm-api-model 0.0.399
- chore: update golangci-lint in GH action to 1.61.0
- gcp-psc update and version bump (#1002)
- Add FedRAMP example (#1000)
- OSDEV-1691: Bump version and added new changes
- Updated version in makefile and generate files
- OSDEV-1691: Bump version and added new changes
- chore: bump version to 0.1.442
- OCM-10734 | feat: bump api model to v0.0.395
- bump version to 0.1.441 (#996)
- chore: Release v0.1.440
- OCM-10883 | chore: bump api model to v0.0.393
- OCM-10883 | feat: include manifests example
- OCM-10939 | chore: bump metamodel to v0.0.61
- chore: release v0.1.439
- chore: bump to ocm-api-model 0.0.392
- chore: release v0.1.438
- chore: bump to ocm-api-model 0.0.391
- bump version to 0.1.437 (#988)
- chore: bump to ocm-api-model 0.0.387 (#987)
- chore: Release 0.1.436
- chore: Bump model into v0.0.389
- bump version to 0.1.435
- Update Makefile
- chore: bump ocm-api-model to v0.0.388
- bump version to 0.1.434 (#983)
- chore: bump to ocm-api-model 0.0.387 (#982)
- bump version to 0.1.433 (#981)
- chore: bump to ocm-api-model 0.0.385 (#980)
- Bump OCM API model version to 0.0.384
- bump version to 0.1.432 (#976)
- chore: bump to ocm-api-model 0.0.383 (#975)
- bump version to 0.1.431 (#974)
- chore: bump to ocm-api-model 0.0.382 (#973)
-  bump version to 0.1.430
- chore: bump to ocm-api-model 0.0.381
- bump version to 0.1.429
- chore: bump to ocm-api-model 0.0.380
- bump version to 0.1.428 Signed-off-by: marcolan018 <llan@redhat.com>
- add ec2MetadataHttpTokens to AWSNodePool Signed-off-by: marcolan018 <llan@redhat.com>
- chore: bump to ocm-api-model 0.0.378 (#967)
- chore: bump version and changes for 0.1.426
- Windows support has been included to systemcertpool
- OCM-8291 | chore: bump version and changes.md for 0.1.425
- OCM-8291 | chore: bump windows certs
- ARO-8404 | chore: release v0.1.424 (#963)
- chore: synchronise to ocm-api-model 0.0.377
- ARO-7434 | chore: release v0.1.423
- chore: synchronise to ocm-api-model 0.0.376
- OCM-7288 | chore: release v0.1.422
- OCM-7288 | feat: add managementUpgrade parameters to NodePool model OCM-7374 | feat: add AdditionalAllowedPrincipals to aws model
- OCM-5952: Server name inference for regionalized OCM redirects
- OCM-8260 | fix: revert change to connection default URL
- OCM-8260 | chore: Release v0.1.421
- OCM-8022 | feat: add timestamps for rpb
- OCM-8260 | feat: bump model to v0.0.374 and include access transparency examples
- ARO-7210 | chore: release v0.1.420
- chore: synchronise to ocm-api-model 0.0.373
- OCM-7308 | chore: Added release notes for 0.1.419 release
- OCM-7308 | feat: Generated a new version of the OCM-SDK to support KubeletConfigs on HCP MachinePools
- OCM-6715 | chore: Release v0.1.418
- OCM-6715 | feat: bump model to .371
- Update CHANGES.md
- OCM-6535 | feat: changes to support sts arbitrary policies
- OCM-7566 | chore: update changes and bump version
- OCM-7566 | chore: update windows cert for api.openshift.com:443
- chore: bump SDK version to 0.1.415
- chore: synchronise to ocm-api-model 0.0.369
- chore: update model to v0.0.367
- chore: remove explicit install of ginkgo in CI
- feat: install and use ginkgo at project level
- chore: set default Makefile target to `examples`
- OCM-7117 | chore: Release 0.1.414
- OCM-7117 | chore: update metamodel version in sdk to v0.060
- Bump sdk version to 0.1.413
- sync ocm-api-model version 366
- Bump sdk version to 0.0.412
- chore: synchronise to ocm-api-model 0.0.365
- Update CHANGES.md
- Update version of `pgx`
- OCM-5823 | chore: release v0.1.410
- OCM-5823 | chore: synchronise to ocm-api-model 0.0.364
- bump the sdk to 0.1.409
- bump the model to v0.0.363
- chore: release 0.1.408 (#924)
- chore: synchronise to OCM-API-MODEL 0.0.361
- OCM-6454 | chore: Release 0.1.407
- OCM-6454 | feat: bump model to .360 and example
- Release v0.1.406 (#921)
- Update CHANGES.md and version for v0.1.405
- Bump api-model to v0.0.359
- Bump github.com/dvsekhvalnov/jose2go from 1.5.0 to 1.6.0 in /examples
- Bump github.com/dvsekhvalnov/jose2go from 1.5.0 to 1.6.0
- Migrate keychain to use non-CGO libraries
- Update CHANGES.md and version for 0.1.404
- OCM-5822 | feat: Updating to 0.0.357
- OCM-5976 | fix: Generated files for 0.0.356
- OCM-5976 | feat: Updating Changes.md
- OCM-5976 | fix: reverting change for provison shard
- OCM-5976 | feat: Updating Changes.md
- OCM-5976 | feat: Updating to 0.0.355
- OCM-4965: Secure store tweaks (#905)
- feat: release v0.1.401
- feat: adding support for package_image in clusters_mgmt
- Removed version line.
- Fix the message format and update the version to 0.1.400.
- Updated model version.
- Generate the sdk based on the latest ocm-api-model modifications.
- Remove Fullname search parameter from status board product, application, and service.
- Release 0.1.399 (#902)
- OCM-4965 | feat: Secure Store additional error handling and removal of keys (#901)
- Update changed.md for Release 0.1.398 (#900)
- Release 0.1.398 (#899)
- OCM-5281 | Feat | Add regions support from ocm shards (#882)
- OCM-2780 | fix: Move fix to CheckContentType
- OCM-2780 | fix: Don't error when response is 204 and no content-type
- Prepare release 0.1.397
- Update CHANGES.md
- Bump to OCM API v0.0.350
- add test for connection with no authentication
- add get cluster with unauthenticated connection example
- feat: add unauthenticated connection builder
- OCM-4964: Refine allowed keystores (#894)
- Update CHANGES.md
- OCM-1599 | feat: display availability zone types and subnet outposts
- bump the sdk to 0.1.394
- bump the model to v0.0.347
- OCM-5759: Rework device auth as multi-step process
- OCM-5759: Update example logs
- OCM-5759: Remove references to access token
- Remove scope
- OCM-5759: Add device auth
- Release 0.1.393
- OCM-4964: Add example
- tidy
- OCM-4964: Error handling with no creds
- OCM-4964: Add compression
- OCM-4964: Remove OS specific files
- OCM4964: Revised secure store labels
- OCM-4946: Use single keyring impl for all OS
- OCM-4964: Generic Secure Store Service
- OCM-4946: use docker creds for linux
- Clean HEAD
- OCM-4946: Secure store darwin and linux
- Added consts for hard coded values
- Pointed to Red Hat prod SSO with designated clientId
- Fix linter error on client_selector.go
- Prepare v0.1.392 release
- Removed unnecessary log
- Shut down server gracefully
- Removed client secret from the method
- Added comments for better understanding
- Fix linter errors
- Remove unnecessary file
- Add authentication using OAuth2 and PCKE
- chore: Release v0.1.391
- OCM-1706 | chore: bump model to v0.0.345
- Update CHANGES.md
- OCM-5312 | Add Platform to subnet_network_verification_type
- Prepare v0.1.389 release
- prepare v0.1.388 release  -Update ocm api model version to 0.0.341
- update CHANGES.md
- Add Platform to network_verification_type
- chore: release 0.1.386
- Bump model to 0.0.339
- chore: release 0.1.385
- Bump model to 0.0.338
- OCM-4744 | chore: release notes for v0.1.384
- OCM-4744 | chore: Bump the SDK to use API model version 0.0.337
- chore: release 0.1.383
- Bump model to v0.0.336
- chore: release 0.1.382
- Bump model to v0.0.335
- Bump api-model to v0.0.334
- OCM-3852 | chore: Release notes for SDK version 0.1.381
- OCM-3852 | feat: Added SDK support for KubeletConfig and day 2 operations for setting podPidsLimits
- OCM-4601 | chore: Release v0.1.380
- OCM-4601 | chore: bump api-model to v0.0.332
- chore: release v.0.1.379
- OCM-3234 | chore(release): release 0.1.378
- OCM-3234 | bump model to 0.0.330
- OCM-4313: Update to go v1.21
- OCM-4313: Update golang in ocm-sdk-go v1.20
- Updated Example for Network Verifier
- Update CHANGES.md
- Update sdk to reflect new field in network_verification_type resource
- chore: Release v0.1.376
- OCM-4192 | feat: bump model to 0.0.328
- bump version to 0.1.375 Signed-off-by: marcolan018 <llan@redhat.com>
- add BestEffort in delete method for Cluster Signed-off-by: marcolan018 <llan@redhat.com>
- chore(release): release 0.1.374
- OCM-1534: Specify Backplane URL from CS API by returning the URL in the environment resource
- feat: install and use tools at project scope level
- Add OrganizationId to Feature review request type
- release 0.1.372
- Update the log entry resources from new model updates
- Release 0.1.371
- Add GCPMarketplaceEnabled to version type
- chore: Release v0.1.370
- OCM-3746 | feat: bump api-model to v0.0.322
- OCM-2437 | feat: Added support for /api/clusters_mgmt/v1/aws_inquiries/sts_account_roles to the SDK
- release 0.1.368
- SDK Bump to use model v0.0.318
- OCM-3277 | feat:Updated SDK version to 317
- Bump ocm-api-model version to 0.0.316
- Release 0.1.367 (#832)
- Add DisplayName and Description to BillingModelItem (#831)
- OCM-3548 | chore: release v0.1.367
- OCM-3548 | feat: update windows CAs
- Release 0.1.366 (#829)
- Added BillingModel resource and type (#828)
- chore: release v0.1.365
- feat: updating ocm-sdk with adding support for new endpoints in addons_mgmt
- chore: release v0.1.364
- chore: release v0.1.363
- OCM-250 | chore: Release v0.1.362 (#820)
- OCM-250 | redact aws access_key and secret_access_key
- chore: release v0.1.361
- bump version to 0.1.360 Signed-off-by: marcolan018 <llan@redhat.com>
- move PrivateHostedZoneID and PrivateHostedZoneRoleARN to AWS res Signed-off-by: marcolan018 <llan@redhat.com>
- Update CHANGES.md
- Bump ocm-api-model version to 0.0.306
- Updated CHANGES.md
- Bump ocm-api-model version to 0.0.304
- Update CHANGES.md
- Bump ocm-api-model
- Update CHANGES.md
- Update name of STS Support Jump Role
- bump version to 0.1.355 Signed-off-by: marcolan018 <llan@redhat.com>
- Add "UserDefined" in dns domain resource Signed-off-by: marcolan018 <llan@redhat.com>
- OCM-209 | chore: release v0.1.354
- Release v0.1.353 (#806)
- Prevent connection leak from retry transport wrapper
- chore: Release v0.1.352
- OCM-851,OCM-2150 | feat: bump api-model to add managed ingress attributes
- Release 0.1.351 (#800)
- Changed delete_associated_resources to deleteAssociatedResources for account client (#799)
- fix typo in CHANGES.md Signed-off-by: marcolan018 <llan@redhat.com>
- bump version to 0.1.350 Signed-off-by: marcolan018 <llan@redhat.com>
- user model version 0.0.295, change ReverveAt to ReverveAtTimestamp Signed-off-by: marcolan018 <llan@redhat.com>
- Add DeleteAssociatedResources parameter to accounts delete endpoint
- Update to metamodel 0.0.59
- update changes doc and version for release 0.1.349
- bump model to version 0.0.293
- chore: release v0.1.348
- chore: bump api-model to v0.0.291
- chore: release v0.1.347
- Bump model version to v0.0.290
- chore: release v0.1.346
- OCM-944 | chore: bump api model to v0.0.289
- Release v0.1.345
- Bump model version to v0.0.288
- bump version to 0.1.344 Signed-off-by: marcolan018 <llan@redhat.com>
- add HTPasswd to Cluster Signed-off-by: marcolan018 <llan@redhat.com>
- chore: release v0.1.343
- Bump model version to v0.0.286
- chore: release v0.1.342
- chore: bump model to version v0.0.285
- chore: release v0.1.341 (#773)
- chore: bump model to version v0.0.283 (#775)
- Update to metamodel 0.0.58
- Bump model version to v0.0.282
- chore: release v0.1.340
- Bump model version to v0.0.281
- Release v0.1.339 (#769)
- Bump model version to v0.0.280 (#768)
- chore: Release v0.1.338
- SDA-8990 | fix: api side windows CAs
- chore: Release v0.1.337
- SDA-8990 | fix: Replace an expired trusted CA on Windows with new ones
- Bump version and updates CHANGES.md
- Support audit log forwarding in cs
- Update model version to v0.0.278
- Release v0.1.334
- Add delete account example
- Bump model version to v0.0.276
- Bump OCM API model version
- chore: bump version and update change log for release v0.1.332
- chore: bump model to version v0.0.273
- chore: Release v0.1.331
- chore: Bump api-model to v0.0.272
- Release v0.1.330 (#752)
- chore: Release v0.1.329
- chore: bump ocm-api-model to v0.0.270
- chore: Release v0.1.328
- feat: bump api model and add delete protection example
- chore: release v0.1.327
- chore: bump api-model 0.0.268
- chore: Release v0.1.326
- chore: bump api-model to version 0.0.267
- Release v0.1.325
- fix: bump api model to v0.0.266
- Release v0.1.324
- Bump api model to 0.0.265
- chore: release v0.1.323
- chore: bump api model to v0.0.264
- Release 0.1.322
- Bump api model to 0.0.263
- Update api model to v0.0.263 and prepare for release
- [SDA-8369] Bump api model to 0.0.263 (#742)
- chore: release v0.1.321
- bump golang matrix and golangci-lint
- chore: bump api model to v0.0.262
- Bump golang.org/x/text from 0.3.7 to 0.3.8 in /examples
- release v0.1.320
- feat: [SDA-8223] updated addonsmgmt with annotations and labels
- Bump ocm-api-model to v0.0.260
- Bump ocm-api-model to v0.0.259
- Release 0.1.317
- Bump api model 0.0.257
- chore: update changes and bump version for release 0.1.256
- bump model to version 0.0.256
- Bump model version and create a new release 0.1.315
- Bump api model 0.0.254 and Release 0.1.314
- Bump api model 0.0.253 and Release 0.1.313
- updated version
- better formatted
- Update metamodel and release v0.1.312
- chore: bump api model 0.0.252 release 0.1.311
- chore: update change log for release 0.1.310, and bump version
- chore: update model to 0.0.251
- chore: update changes file for version v0.1.309
- chore: update model to v0.0.250
- chore: Release v0.1.308
- chore: bumping api model to 249
- Generated files
- Bump ocm-api-model version + create a new release
- Release 0.1.306 (#709)
- Bump model to v0.0.247 (#708)
- Release v0.1.305 (#707)
- Use model v0.0.246 (#706)
- chore: bump api model to v0.0.245 and release sdk v.0.1.304
- Release 0.1.303
- chore: release v0.1.302
- chore: bump api model to 243
- Release 0.1.301 (#700)
- Update model to v0.0.241 (#699)
- update changes for version 0.1.300
- bump model v0.0.240
- generated for go 1.19 and goimports 0.4.0
- bump pr check to include golang 1.19 and goimport 0.4.0
- Release 0.1.299
- Release 0.1.298 Update to model v0.0.238
- Release 0.1.297 Update to model v0.0.237
- release 0.1.296 (#692)
- updated model to v0.0.236 (#691)
- Release 0.1.295 (#690)
- Update to model 0.0.235
- Release 0.1.294
- Release 0.1.293
- Release 0.1.292
- bump sdk version to 0.1.230 and update change log
- bump model to 0.0.230
- Release 0.1.290
- Add Addons Management client in SDK
- Bump API model to 0.0.229
- update changes for version 0.1.289
- bump model to 0.0.228
- Release v0.1.288
- fix: Replace an expired trusted CA on Windows with new ones
- Add osdfleetmgmt entry for check-pull-reqest
- Fix OSD FM endpoint path
- Fix version after review
- Add OSD fleet management in SDK
- Release 0.1.287
- Replace deprecated ioutils calls
- Fix gosimple lint error
- skip osdfleetmgmt linting
- gitignore vscode
- Bump linter version
- make generate results
- Bump ocm-api-metamodel
- Revert "Disable linter until we fix it"
- release 0.1.286
- Added ManagedBy property in RoleBinding type
- bump model to 0.0.219 and update changes for version 0.1.285
- bump model to 0.0.219
- Bump OCM API model version: Add kubeconfig as well as server URL to provision shard
- Bump OCM API model version to v0.0.217
- Disable linter until we fix it
- Bump OCM API version
- release 0.1.281
- Added label locator to get label by key
- Bump ocm-api-model version: add update function to provision shard
- Bump ocm api model version
- Fix example
- Bump ocm-api-model and create a release
- release 0.1.277
- Added capabilities property in Account_Type
- Updated version to add delete registry credential method
- Changed model version to latest to include delete registry credentials method
- Release 0.1.275
- release 0.1.274
- release 0.1.274
- release 0.1.273
- updated model to v0.0.205
- update changes doc for version release 0.1.272
- bump model to v0.0.204
- release 0.1.271
- updated model to v0.0.203
- Update gopkg.in/yaml.v3
- release 0.1.270
- Updated model to v0.0.202
- Release v0.1.269
- Update model to v0.0.201
- authentication: Allow client credential grants with basic auth
- Release 0.1.268
- Update to model 0.0.200
- release v0.1.267
- update model v0.0.199
- Release v0.1.265
- Update model v0.0.132
- windows: Update certificates
- Fix example code.
- Bump ocm-api-model version to 0.0.197
- Update version of database used for tests
- Update dependencies
- release v0.1.264
- update model v0.0.196
- Release 0.1.262
- Release 0.1.262
- Update to model 0.0.194
- Updated to version 0.0193
- Release 0.1.260
- Update to model 0.0.191
- Release 0.1.259
- Update to model 0.0.190
- Update Release to v0.1.258
- Update model v0.0.189
- Update `golangci-lint` and disable cache
- Test with Go 1.18
- Skip speed sensitive tests in GitHub actions
- Release 0.1.257
- Update to model 0.0.188
- Update to metamodel 0.0.54
- Update to metamodel 0.0.53
- RCA vs Rca.
- Add some webrca examples.
- Release 0.1.256
- Update to model 0.0.187
- Updated model to v0.0.186
- upgraded version to 0.1.254
- fixed model version
- updated model to v0.0.185
- updated model 0.1.184
- Update to model 0.0.183
- Release 0.1.251
- Update to model 0.0.182
- Release v0.1.250
- Update model to v0.0.181
- Release 0.1.249
- Update to model 0.0.180
- Release 0.1.248
- Update to model 0.0.179
- Add impersonation test
- Release 0.1.247
- Update to metamodel 0.0.52
- Release 0.1.246
- Update to model 0.0.178
- Release 0.1.245
- Update to model 0.0.177
- updated api-model to v0.0.176
- Release v0.1.243
- Update model to v0.0.175
- Release v0.1.242
- Update model to v0.0.174
- Add an example program to list status updates for status-board.
- Release v0.1.241
- Update model to v0.0.173
- added list_quota_cost example
- Release 0.1.240
- Update to model 0.0.172
- Release 0.1.239
- Revert "Don't crash in debug mode deserializing an empty response"
- Update to metamodel 0.0.51
- Check that `204 No Content` without body is accepted
- Release 0.1.238
- Update to model 0.0.170
- Remove unnecessary environment variables
- Increase sleep time of leadership flag test
- Increase timeout of retry tests
- Update to Ginkgo 2.1.1
- Release 0.1.237
- Update to metamodel 0.0.50
- Update to model 0.0.169
- Update to Ginkgo 2.1.0
- Release v0.1.235
- Update model to v0.0.168
- Install metamodel with `go install`
- Release 0.1.234
- Use Go 1.17 for code generation tests
- Update to version 4 of JWT library
- Update to Ginkgo 2
- Bump ocm-api-model to v0.0.167
- Bump ocm-api-model version to v0.0.166
- Double period used by leadership flag tests
- Update changes
- Bump ocm-api-model to v0.0.165
- Update version of JSON interator
- Update SQL drivers
- Don't use `github.com/ghodss/yaml`
- Require Go 1.16
- Bump ocm api model to version 0.0.163
- Bump ocm-api-model to v0.0.162
- Remove unreliable leadership flag tests
- Bump ocm-api-model to version 0.0.161
- Update to metamodel 0.0.46
- Use metamodel released binary
- Bump ocm-api-model to version 0.0.160
- Bump ocm-api-model to version 0.0.159
- chore(version): bump to 0.1.225 [SDB-2509]
- MatchJQ should require at least one result
- Release 0.1.224
- Update to model 0.0.156
- Support pull-secret access token as a valid token.
- Release 0.1.223
- Update to model 0.0.155
- Add missing `requirements.txt` file for GitHub actions
- Release 0.1.222
- Add changes to release descriptions
- Rename `master` branch to `main`
- Release 0.1.221
- Fix bug - service log is not terminating properly
- Update README.md examples (again)
- Update README.md examples
- Minor README.md fix
- Use Markdown instead of AsciiDoc
- Release 0.1.220
- Add `RespondWithPatchedJSON`
- Add `jq` testing utilities
- Release v0.1.219
- Update model to v0.0.153
- Release 0.1.218
- Update to metamodel 0.0.44
- Release 0.1.217
- bump api model 0.0.152
- Release 0.1.216
- Update to metamodel 0.0.43
- Bump OCM API model to v0.0.151
- Update `github.com/microcosm-cc/bluemonday` to 1.0.16
- Release v0.1.214
- Update model to v0.0.150
- Release 0.1.213
- Retry when `REFUSED_STREAM`
- Update to metamodel 0.0.42
- Release 0.1.212
- Update model to 0.0.149
- Replace expired CA certs for windows
- Add support for digging inside maps
- Copy data package from CLI
- Fix typo in HTTP internal server error message
- Update the version const
- Update ocm-api-model version: remove archived clusters endpoint
- Release 0.1.210
- Update to model 0.1.147
- Release v0.1.209
- Bump API model to v0.0.146
- Avoid hard-coded private keys
- Release v0.1.208
- Update model to v0.0.145
- Release v0.1.207
- updated model to v0.0.144
- Update model version to include Add() limited support reason
- Update model version to v0.0.142
- Implement the `http.Flusher` interface
- Run tests with Go 1.17
- Run tests serially
- Increase wait times of leadership flag tests
- Fix retry wait
- Increase timeouts for retry tests
- Add Go submodule for examples
- Add publish release workflow
- Accept tokens without the `typ` claim
- Opaque refresh tokens
- Accept lower case token type
- Enable tests that need a database
- Use 'go install' instead of 'go get' to install tools
- Fix token close to expire test
- Remove `pr_check.sh`
- Empty `pr_check.sh`
- Add generate action
- Remove tests from `pr_check.sh`
- Build examples as part of the `test` job
- Add GitHub action to lint Go code
- Prepare to upgrade golangci-lint to 1.42.0
- Add GitHub actions
- Fix failing configuration file test
- Update Ginkgo
- Release 0.1.204
- Change level of token retry messages
- Remove incorrect comment
- Release 0.1.203
- Update to metamodel 0.0.39
- Retry requests with body
- Update ocm-api-model version to v0.0.141
- Write failed request details to the debug log
- Retry when EOF or connection reset by peer
- Release 0.1.201
- Reject impersonation
- Retry when server sends go away before settings
- Add 'check_optional_terms' and 'reduceClusterList' fields
- dump: Add tests for dumpJson
- Release 0.1.199
- dump: Skip when writing null
- Fix minor JSON debug issue
- Use `golang-jwt/jwt` instead of `dgrijalva/jwt-go`
- dump: Format output & redact sensitive fields using json helpers
- internal: Switch to bluemonday to strip tags from html content
- Release v0.1.198
- updated model to v0.0.139
- Release v0.1.197
- Update model to v0.0.138
- Release v0.1.196
- Update model version to v0.0.137
- Update to model version v0.0.136
- Rename `Random` to `Jitter` and introduce timeouts
- Add `MatchJSONTemplate`
- Move `EvaluateTemplate` to a separate file
- Fix typos
- Release 0.1.194
- Update to metamodel 0.0.38
- Release 0.1.193
- Update to model version v0.0.135
- Fix initialization of metrics registerer
- Release 0.1.191
- Add leadership flag
- Release 0.1.190
- Don't use refresh token if have client credentials
- Release 0.1.189
- Don't require refresh token for client credentials grant
- Release 0.1.188
- Release v0.1.187
- Update model to v0.0.133
- Release v0.1.186
- Update model to v0.0.132
- Release 0.1.185
- Add `All` and `Any` functions to the logging package
- Release v0.1.184
- Update ocm-api-model to v0.0.129
- Update model to v0.0.128 and Release v0.1.183
- Update model to v0.0.127
- Release v0.1.182
- Update model to v0.0.126
- Release 0.1.181
- Add support for authentication cookie
- Release 0.1.180
- Remove access token authentication
- Release 0.1.179
- Enable building with Go 1.13
- Release v0.1.178
- Update model to v0.0.125
- Update of parameters to JobQueue#Push
- Release 0.1.176
- Added Arguments to JobQueue#Pop
- Update `Logger` interface to include `Fatal` level
- Add Fatal logging level
- - JobQueue fixes and better example
- 1. We're unable to support 'list Job Queues' 2. Proper example of Push/Pop for 'Job Queues'
- Release v0.1.173
- Bump model to v0.0.119
- Bump ocm-api-model to 0.0.118
- Add JobQueue service
- Add JobQueue service
- Release 0.1.170
- Fix typo
- Add _tolerance_ parameter to authentication handler
- Delete example for get cluster metrics
- Update to model version 0.0.115
- Release v0.1.168
- Bump ocm-api-model to v0.0.114
- Release 0.1.167
- Update to model 0.0.113
- Use defaults from authentication package
- Release v0.1.166
- Updte to ocm-api-model v0.0.112
- Move token logic to separate transport wrapper
- Move client selection logic to separate type
- Release 0.1.165
- Fix wrong TLS server name
- Release 0.1.164
- Update to model 0.0.111
- typo error
- Change user agent to `OCM-SDK`
- Release 0.1.163
- Add support for HTTP/2 with TLS
- Add `h2c` support
- Add metrics handler wrapper
- Reorder execution of tests
- Remove logger from metrics transport wrapper
- Enable compression
- Release v0.1.162
- Update model to v0.0.110
- Bump ocm-api-model version
- Run tests in parallel
- Release 0.0.160
- Update to model 0.0.108
- Improve testing of metrics
- Release v0.1.159
- Update API model to v0.0.107
- Add metrics package
- Relase 0.1.158
- Downgrade to Go 1.14
- Release v0.1.157
- Update model to v0.0.106 and metamodel to v0.0.36
- Add trusted CA certificates for Windows
- Missplaced return after warning from send
- Accept Empty Reader as non-nil req body
- Release 0.1.156
- connection: Skip loading SystemCertPool on Windows
- Update to model 0.0.105
- Declare go 1.15 in go.mod
- Release 0.1.155
- Update to model 0.0.104
- Release 0.1.154
- Update to model 0.0.103
- Release 0.1.153
- Add support for customizing error responses
- Add support for Unix sockets
- Release v0.1.152
- Update model to v0.0.102
- bump version to 0.1.151
- bump model to version 0.0.101
- Add redirection tests
- Reject URLs without scheme or host name
- Load `os.Stdin` in dump configuration example
- Load string if it doesn't look like a file
- Load metrics subsystem configuration from file
- Add method to read metrics subsystem
- Rename `Metrics` to `MetricsSubsystem`
- Create version 0.1.150
- Bump to api-model v0.0.100
- Move logging code to `logging` package
- Release 0.1.149
- Fix method to get base URL
- Release 0.1.148
- Add `!yaml` configuration tag
- Fix minor issue in `MakeServer` test function
- Add `!shell` configuration tag
- Add test for binary configuration data
- Release 0.1.147
- Load connection configuration from file
- Add support for loading trusted CA files
- Release v0.1.146
- Update model v0.0.99 and metamodel v0.0.35
- Don't crash in debug mode deserializing an empty response
- Add support for alternative URLs
- Add `EvaluateTemplate` function for tests
- Create version 0.1.145
- Bump to api-model v0.0.98
- Fix incorrect use of `%w`
- Honour tokens context deadline
- Wrap errors
- Always use _Ginkgo_ writer from test servers
- Always use _Ginkgo_ writer for logs of tests
- Fix standard streams logger
- Add functions to simplify tests that require responses
- Release v0.1.144
- Update model to v0.0.96
- Bump SDK to version 0.1.143
- Bump SDK to ocm-api-model v0.0.95
- Release v0.1.142
- Update model to v0.0.94
- Allow disabling keep alive connections in the SDK connection transport
- Update version to 0.1.141 using ocm-api-model v0.0.92
- Release v0.1.140
- Update model to v0.0.91
- Release 0.1.139
- use new limit for connection errors
- Interpret HTML entities in logged summary of error response
- update model to 0.0.90
- Update api model to v0.0.89
- Release 0.1.138
- update model to 0.0.88
- Release 0.1.137
- update model to 0.0.87
- Release 0.1.136
- Bump sdk version
- Update doc and bump sdk version
- Bump ocm-api-model version
- Update to api model v0.0.84
- Release v0.1.133
- Update metamodel to v0.0.33
- increase the limit/size of content summary
- Release v0.1.132
- Update model to v0.0.83
- Release 0.1.131
- Support http/s proxy
- Release v0.1.130
- Update model to v0.0.82
- Release v0.1.129
- Request a token valid for longer than 1 min
- Update version
- go mod tidy
- Create new SDK version v0.1.127
- Release 0.1.126
- update ocm-api-model to 0.0.79
- Release v0.1.125
- model: Update to v0.0.78
- Release v0.1.124
- Allow auth header of type AccessToken
- Release v0.0.123
- Update to model v0.0.77
- Update to metamodel v0.0.32
- Set token expiry function public
- "got tokens on attempt 1" is not interesting
- Add tests for request metrics
- Create a new version
- Update changes
- Assume expiration 0 when no expiration included in the token
- Fix panic from unexpected prometheus label
- Add tests for token metrics
- mark 0.1.121 as broken
- Release 0.1.121
- Print golangci-lint version
- lint
- Better logging and metrics when retrying SSO
- [refactor] use explicit backoff.Permanent for non-500 errors
- Release 0.1.120
- Update to model 0.0.76
- Bump version
- Add unit tests
- Remove redundant mock hanlers
- Fix linter issues
- Perform access token retry only when receiving http 5xx code + fix UT
- Fix missing code assingment
- Support retry when fetching access token
- Release v.0.1.118
- Update to ocm-api-model 0.0.75
- Update version
- Fix example
- bump ocm api model to 0.0.74
- Release v0.1.116
- Update model to v0.0.73
- Added support for http PUT method
- Release v0.1.115
- Update metamodel to v0.0.31 and model to v0.0.72
- Update the version
- Upgrade ocm-api-model
- Release v0.1.113
- Update to model v0.0.70
- update sync set example
- improve error response from gateway
- Update to model v0.0.69 and release v0.1.112
- Release v0.1.111
- Update to model v0.0.67
- Update v0.1.110 changes
- Update SDK with latest version of API model
- Add syncsets example
- Release 0.0.109
-  Update to model v0.0.65
- Update to metamodel v0.0.30
- Update version to 0.1.108
- Update to model v0.0.64 and release v0.1.108
- Release v0.1.107
- Update to model v0.0.63
- Release v0.1.106
- Update to model v0.0.62 and metamodel v0.0.29
- add a new label to metrics to identify the service
- Use specific version of goimports as the latest gets error
- Use specific version of goimports as the latest gets error
- Use specific version of ginkgo as later version failes on fetching
- Tidy go modules
- Add doc to version 0.1.105
- Update to model version 0.0.59 - Feature Toggle support
- Release v0.1.104
- Release v0.1.103
- Update to model v0.0.58
- Release 0.1.102
- Update to model v0.0.57 and metamodel v0.0.28
- make: Allow passing local copy of model and metamodel
- Release 0.1.101
- Update to model 0.0.56
- Release 0.1.100
- Update to model 0.0.55
- Release 0.1.99
- Update to model 0.0.54 and metamodel 0.0.27
- Release 0.1.98
- Update to model 0.0.53
- Release 0.1.97
- Update to model 0.0.52
- Release 0.1.96
- Update to model 0.0.50
- Release 0.1.95
- Update to model 0.0.48
- Release 0.1.94
- Update to model 0.0.47
- Release 0.1.93
- Update to model 0.0.46
- Release 0.1.92
- Update to model 0.0.45
- Update to model 0.0.44
- Release 0.1.91
- Update to model 0.0.42
- Release 0.1.90
- Request new tokens when refresh grant retursn `invalid_grant`
- Add cookie tests
- Add tests for client credentials grant
- Check JSON response content type
- Add cookie support
- Release 0.1.89
- Update to metamodel 0.0.26
- Release 0.1.88
- Update to metamodel 0.0.25
- Remove service and version parametrs from authentication handler
- Release 0.1.87
- Update to metamodel 0.0.24
- Preserve order of debug JSON
- Release 0.1.86
- Update to model 0.0.41
- Release 0.1.85
- Update to metamodel 0.0.23
- ClusterOperatorMetricExample
- Release 0.1.84
- Update to model 0.0.40
- Fix comments in authentication handler
- Remove duplicated authentication tests
- Release 0.1.83
- Update to model 0.0.39
- check non-json response and return error
- Release 0.1.82
- Update to model 0.0.38
- Fix typos
- Add UT for connection transport wrapper
- Release 0.1.81
- Add ability to intercept request and response using a trasnport middleware of type http.RoundTripper
- Fix typos
- Release 0.1.80
- Add body details in case of error from token provider
- Release 0.1.79
- Update to metamodel 0.0.22
- Release 0.1.78
- Update to metamodel 0.0.21
- Fix URL prefix for the logs service
- Release 0.1.77
- Update to model 0.0.37
- Don't require Go 1.13
- Release 0.1.76
- Update to model 0.0.36
- Add support to CurrentAccess
- Release 0.1.75
- Release 0.1.74
- Update to model version 0.0.33
- Release 0.1.73
- Update to model 0.0.32
- Release 0.1.72
- Update to model 0.0.31
- Release 0.1.71
- Don't log token errors
- Send `WWW-Authenticate` response header
- Send 401 instead of 511
- Release 0.1.70
- Add authentication handler
- Update to metamodel 0.0.20
- Fix _godoc_ link
- Release 0.1.69
- Update to model 0.0.30
- Release 0.1.68
- Update to metamodel 0.0.19
- Release 0.1.67
- Don't check kinds of add-ons installations
- Release 0.1.66
- Update to model 0.0.29
- Release 0.1.65
- Update to model 0.0.28
- Release 0.1.64
- Update to model 0.0.27
- Release 0.1.63
- Update to model 0.0.26
- Fix metrics example
- Release 0.1.62
- Update to model 0.0.25
- Release 0.1.61
- Update to metamodel 0.0.18
- Release 0.0.60
- Release 0.1.59
- Update to model 0.0.23
- Release 0.1.58
- Update to metamodel 0.0.16
- Release 0.1.57
- Update to metamodel 0.0.15
- Release 0.1.56
- Update to model 0.0.22
- Release 0.1.55
- Add example showing how to sync add-ons
- Update to model 0.0.21
- Release 0.1.54
- Update to metamodel 0.0.14
- Drop support for _developers.redhat.com_
- Release 0.1.53
- Update to model 0.0.20
- Update to metamodel 0.0.13
- Release 0.1.52
- Update to model 0.0.19
- Release 0.1.51
- Update to model 0.0.18
- Release 0.1.50
- Update to metamodel 0.0.12
- Release 0.1.49
- Update to model 0.0.17
- Release 0.1.48
- Update to model 0.0.16
- Update to metamodel 0.0.11
- Release 0.1.47
- Update to metamodel 0.0.10
- Release 0.1.46
- Update to model 0.0.15
- Release 0.1.45
- Update to model 0.0.14
- Release 0.1.44
- Update to model 0.0.13
- Update to metamodel 0.0.9
- Update to metamodel 0.0.8
- Don't require POST body
- Update contributing doc to mention examples and link to new ocm-api-model docs
- Release 0.1.43
- Add a CONTRIBUTING document detailing release process
- Release 0.1.42
- Update to ocm-api-model v0.0.12 to add authorizations/v1/[self_]access_review
- Remove `authorizations` directory when generating code
- Move `golangci-lint` configuration to a separate file
- Use `GOPROXY`
- Release 0.1.41
- Update to model 0.0.11
- Release 0.1.40
- Update to model 0.0.10
- Release 0.1.39
- Update to model 0.0.9
- Release 0.1.38
- Update to model 0.0.8
- Release 0.1.37
- Update to model 0.0.7
- Release 0.1.36
- Update to model 0.0.6
- README: Change client to sdk
- Update to metamodel 0.0.7
- Release 0.1.35
- Release 0.1.34
- Update to model 0.0.3
- Use about to expire access token if needed
- Check that `make generate` doesn't introduce changes
- Release 0.1.33
- Update to model v0.0.2
- Release v0.1.32
- Use metamodel v0.0.4
- Fetch tags of model and metamodel
- Release v0.1.31
- Generated servers enforce trailing slashes as well send 'Content-Type' headers
- README: s/uhc-cli/ocm-cli/
- Release 0.1.30
- Move to `github.com/openshift-online/ocm-sdk-go`
- Release 0.1.29
- Generated servers should handle trailing slashes in routes
- Clone metamodel for code generation
- Clone model for code generation
- Rename main package
- Release v0.1.28
- Add Context parameter to Server methods
- Release 0.1.27
- Add servers to uhc-sdk-go
- changes ClusterRegistration response type from long to string
- Release 0.1.26
- Add support for compute nodes metrics
- Release 0.1.25
- Allows override of source files, fixes data type in cluster reg response
- Use Go modules
- Use `ocm-metamodel-tool`
- Add support for quota summary
- Release 0.1.24
- Select deprecated OpenID with user name and password
- Release 0.1.23
- Release 0.1.22
- Fix names of glog logger builder methods
- Don't warn about issuer URL when no token is provided
- Release 0.1.21
- Add getters for connection attributes
- Release 0.1.20
- Switch to _sso.redhat.com_
- Remove 'admin' field from debug output of ClusterCredentials
- Release 0.1.19
- Expose header of response
- Release 0.1.18
- Add cluster `expiration_timestamp`
- Release 0.1.17
- Add dashboard name
- Add get by index
- Release 0.1.16
- Add getters with presence flag to responses
- Release 0.1.15
- Add versions collection
- Add missing subscription and account attributes
- Release 0.1.14
- Don't panic if there is no response
- Expose request path and method. Can be used for logging.
- Fix typo
- Redact sensitive fields.
- Fix type
- Release 0.1.13
- Add support for building list of structs
- Release 0.1.12
- Add Prometheus metrics
- Add missing `os.Exit(1)`
- Add delete subscrition support
- Release 0.1.11
- Increase token slack
- Release 0.1.10
- Improve context support
- Release 0.1.9
- Add cluster credentials resource
- Release 0.1.8
- Move metrics to `metrics` attribute
- Fix version
- Release 0.1.7
- Prevent re-use of connections when fetching OpenID token
- Fix typo in comments
- Fix example in documentation
- Fix typos
- Improve documentation
- Add missing example
- Add logs collection
- Fix exit code of list clusters example
- Improve the list clusters example
- Fix unmarshalling of list of class instances
- Use typed interface in `README.adoc`
- Update model
- Release 0.1.6
- Add typed interface
- README: link to uhc-cli repo
- Add `godoc` and license badges
- Release 0.1.5
- Fix package paths
- init


## 0.1.510 Aug 19 2026

- ROSAENG-13732 | feat: Implement BYO firewall rule endpoint verification

## 0.1.509 Aug 10 2026

- feat: aro-hcp: allow providing CS Provision Shard ID during cluster creation
- feat: aro-hcp: remove provisionshard reference in cluster type

## 0.1.508 Aug 7 2026

- chore: version bump to v0.1.508 (skip v0.1.507 due to duplicate tag on proxy.golang.org)

## 0.1.507 Aug 5 2026

- ROSAENG-61162 | feat: changes to support BYO firewall for OSD-GCP deployments

## 0.1.506 Jul 23 2026

- ROSAENG-61032 | feat: changes to support SpotMarketOptions for ROSA HCP 

## 0.1.505 Jul 06 2026

- ROSAENG-59819 | feat: add HCC severity constants

## 0.1.504 Jun 23 2026

- chore: bump ocm-api-model to v0.0.459

## 0.1.503 Jun 11 2026

- chore: bump ocm-api-model to v0.0.457

## 0.1.502 Jun 01 2026

- Add two GitHub Actions workflows:
  - auto-tag.yaml, sync-from-model.yaml

## 0.1.501 May 19 2026

- Update model to 0.0.456
  - Add `access_shared_role_arn` field to AWS Disaster Recovery backup config
  - Add `node_count` field to `AutoNode` status

## 0.1.500 Apr 01 2026

- Update model to 0.0.454
  - Add the `ZeroEgress` type and include in the `AWS` as a field.

## 0.1.499 Mar 31 2026

- OCM-22967 | fix: removing pgx dependency

## 0.1.498 Mar 11 2026

- Update model to 0.0.453
  - Add `OidcIssuerUrl` field to ARO-HCP Azure type for workload identity federation
  - Add `Visibility` field and `AzureKmsEncryptionVisibility` type to `AzureKmsEncryption`

## 0.1.497 Feb 25 2026

- Update model to 0.0.452
  - Add the `AvailableChannels` method to the Version resource

## 0.1.496 Feb 24 2026
- Update model to 0.0.451
  - Add the `GcpDnsDomain` type and include in the `DNSDomain` as a field.

## 0.1.495 Feb 23 2026
- Update model to 0.0.450
  - Add `ImageDigestMirrors` field to `ClusterRegistryConfig` for ARO-HCP.

## 0.1.494 Feb 12 2026
- Update model to 0.0.449
  - Add ARO-HCP `ControlPlaneUpgradePolicies` resource to manage control plane upgrade policies
  - Add ARO-HCP `NodePoolUpgradePolicies` resource to manage node pool upgrade policies
  - Simplify ARO-HCP upgrade policy types by removing scheduling-related fields (`Schedule`, `ScheduleType`, `UpgradeType`, `NextRun`, `EnableMinorVersionUpgrades`)

## 0.1.493 Jan 25 2026
- Update model to 0.0.448
  - Add `VnetIntegrationSubnetResourceID` field to ARO-HCP Azure type for SWIFT
    networking (Azure Container Networking Interface) support
  - Expose ARO-HCP cluster's provision shard subresource GET endpoint
  - Expose ARO-HCP provision shard global endpoints

## 0.1.492
- Update model to 0.0.447
  - Make `Topology` attribute of ARO-HCP `ProvisionShard` required and restrict
    its allowed value to `shared`
  - Expand `EtcdEncryption` attribute description for ROSA-HCP Clusters

## 0.1.491
- Update model to 0.0.446
  - Add `AzureShard` and `ProvisionShardMaestroConfig` to ARO-HCP `ProvisionShard`
    type
  - Remove `HypershiftConfig`, `ManagementCluster`, `HiveConfig`,
    `AWSAccountOperatorConfig`, `GCPProjectOperator`, `AWSBaseDomain`, `AWSBaseDomain`
    from ARO-HCP `ProvisionShard` type

## 0.1.490
- Update model to 0.0.444
  - Add `ExcludedNamespaceSelectors` to ingress type

## 0.1.489
- Update model to 0.0.443
  - Add `Channel` field to `Cluster` type for Y-stream update channel management

## 0.1.488
- Update model to 0.0.442
  - Add DeletedClusters new resource to list deleted clusters and get a deleted cluster by its id

## 0.1.487
- Add `cloud_provider` field to CS `dns_domains` type

## 0.1.486
- Align log forwarding config endpoints with API

## 0.1.485
- Add patch operation on version gates

## 0.1.484
- This change exposes a new OCM linter to be used by other project depending on ocm-sdk and enables it into ocm-sdk pipeline

## 0.1.483

- Update model to 0.0.438
  - Add LastLimitedSupportOverrideCheck, LastHibernationCheck, and
    LastClusterImagesetSync fields to Environment type

## 0.1.482

- Update model to 0.0.437
  - Add LogForwarders cluster resource to configure log forwarding
  - Add endpoints to list log forwarding groups and versions

## 0.1.481

- Update model to 0.0.436
  - Add ImageType enum to the NodePool model
  - Add Preference enum field to the AwsCapacityReservation model
  - Add Features sub-object to the MachineType model

## 0.1.480

- Update to model 0.0.435:
  - Update deprecation message for DisableUserWorkloadMonitoring field in
    model/clusters_mgmt
  - update some code comments mentioning clusters_mgmt to aro_hcp in
    model/aro_hcp

## 0.1.479

- Update to model 0.0.434:
  - Split/Duplicate ARO-HCP models. From this point on, the ARO-HCP models
    become independent from the clusters_mgmt model

## 0.1.478

- Update to model 0.0.433:
  - Add `AccountGroup` and `AccountGroupAssignment` types for RBAC functionality.
  - Add `AccountGroups` and `AccountGroupAssignments` resources to `Organization`.
  - Add `AccountGroup` and `AccountGroupID` attributes to the `RoleBinding` type.
  - Add `AccountGroupManagedBy` and `AccountGroupAssignmentManagedBy` enums for system management tracking.

## 0.1.477

- Add resource-scoped permissions for WIF configurations

## 0.1.476

- Update model to version 0.0.431
  - Add support for ImageMirrors for ROSA HCP

## 0.1.475

- Update model to version 0.0.430
  - update MarketType enum to match the camelcase values expected by CS
  - deprecate os disk configuration old design
  - Add API endpoints for cluster autoscaler in ARO HCP

## 0.1.474

- Update mode to version 0.0.429
  - Add support for AutoNode for ROSA HCP

## 0.1.473

- Update model to version 0.0.426:
  - Restructure OS disk attributes for ARO-HCP NodePool
  - Add `AzureNodePoolOsDisk` type and related resources
  - Consolidate Azure node pool OS disk configuration into structured approach

## 0.1.472

- Add script to simplify model update process
- Add git hook to validate pushed tag matches the Version constant
- Add version bump script and update the version file
- Update ocm-api-metamodel to v0.0.422
- Update API model with Azure node pool OS disk types and endpoints

## 0.1.471

- Reorganize CONTRIBUTING.md file and remove unnecessary instructions
- Update release process documentation for ocm-api-model
- Add optional commit SHA argument to update-model script
- Update model with latest changes

## 0.1.470

- Update model to version 0.0.422:
  - Add AWS backup configuration support with `AWSBackupConfig` type
  - Add Azure encryption enhancements including `AzureEtcdDataEncryption`, `AzureKmsEncryption`, and `AzureKmsKey` types
  - Add new access transparency service with access protection and request types
  - Add `ClientComponent` type for ARO-HCP
- Add development workflow improvements
- Prepare infrastructure for Azure node pool enhancements

## 0.1.469

- Update model to add arohcp ManagedIdentitiesRequirements types and endpoints

## 0.1.468

- update golangci to v2
- add macgregor, cristianoveiga to owners/approvers
- leadership flag enhancement: precheck handler added
- move JSON structs (serialization) to ocm-api-model

## 0.1.467

- Update model version v0.0.419
  - Add `LogType` to `NotificationDetailsRequest`

## 0.1.466

- Update model version v0.0.418
  - Add arohcp Version types and endpoints
  - Add missing not-yet-generated changes for node pool encryption

## 0.1.465

- Update model version v0.0.417
  - Add `RhRegionID` field to `Subscription` and `ClusterAuthorizationRequest` type

## 0.1.464

- Update model version to v0.0.416
  - Added annotation to Node Pool State

## 0.1.463

- Updated model version to v0.0.415
  - Added `ClusterCapabilities` type
  - Added `Capabilities` field to `Cluster` type

## 0.1.462

- Updated model version to v0.0.414
  - Added node pool endpoint resources for aro_hcp
  - Added node pool state model in clusters_mgmt
  - Added node pool status resources for aro_hcp
- Updated meta model version to v0.0.67
  - Added async methods

## 0.1.461

- Updated model version to v0.0.413
  - Add the new inflight check endpoints
- Updated meta model version to v0.0.66

## 0.1.460

- Updated model version to v0.0.412
  - Added `ClusterCondition` field to `VersionGate` type

## 0.1.459

- Update model version to v0.0.411
  - Update json values for SDN to OVN Cluster Migration configuration properties

## 0.1.458

- Update metamodel version to v0.0.65
  - Enhance 'ref' mechanism
- Update model versiont to v0.0.410
  - Add ClusterStatus to aro_hcp

## 0.1.457

- Fix leadership/flag not respecting initial context supplied

## 0.1.456

- Update model version to v0.0.409
  - Update cluster migration state to match API object

## 0.1.455

- Update model version to v0.0.408
  - Update cluster migration type constant

## 0.1.454

- Update model version to v0.0.407
  - Add ClusterMigration resources

## 0.1.453

- Update model version to v0.0.406
  - Expose the GCP Shielded VM secure boot setting at the machine pool level

## 0.1.452

- Update model version v0.0.405
- Update metamodel version v0.0.64
  - Add ARO HCP v1alpha1 root resource

## 0.1.451

- Update model version v0.0.404
  - Add WifConfig patch endpoint

## 0.1.450

- Update model version v0.0.403
  - Add `NodesOutboundConnectivity` in azure_type.model to the `Azure` model
  - Add `AzureNodesOutboundConnectivity` model

## 0.1.449

Update model version v0.0.402

- Add WifConfig status endpoint
- Add `WifConfigStatus` model

## 0.1.448

- Update model version v0.0.401
  - Include support for HCP shared vpc fields in `Cluster` and `DNS` models.

## 0.1.447

- Update model version v0.0.400
  - Update ocm-sdk-go examples for Addons
  - Update `Parameters` attribute in addon_installation_type.model to the List of `AddonInstallationParameter` type.
  - Update `Parameters` attribute in addon_type.model to the  List of `AddonParameter` type.

## 0.1.446

- Update model version v0.0.399
  - Improve Clusters Azure API type documentation
  - Add Azure Operators Authentication related types
  - Add `AzureOperatorsAuthentication` type to `Azure` type

## 0.1.445

- Update model version v0.0.398
  - Add updates for GCP Private Service Connect feature
  
## 0.1.444

- Update model version v0.0.397
  - Add `FlapDetection` field to `status_board.statuses` model

## 0.1.443

- Update model version v0.0.396
  - Add `FlapDetection` field to `status_board.status_updates` model

## 0.1.442

- Update model version v0.0.395
  - Add `ASWShard` struct
  - Add `WifEnabled` field to `Version` model

## 0.1.441

- Update model version v0.0.394
  - Add `WifTemplates` field to `WifConfig` model

## 0.1.440

- Update model version v0.0.393
  - Include missing field to hypershift model
  - Include new fields to manifest model
  - Remove v2alpha1 resources

## 0.1.439

- Update model version v0.0.392
  - Add `vm` WIF access type
  - Add `support` field to WifConfig resource

## 0.1.438

- Update model version v0.0.391
  - Add `RegistryConfig` attribute to `Cluster` model
  - Add `RegistryAllowlist` resource and endpoints

## 0.1.437

- Update model version v0.0.390
  - Add `RolePrefix` field to `WifGcp` model

## 0.1.436

- Update model version v0.0.389
  - Add state struct to node pool
  - Add limited support reason override

## 0.1.435

- Add 'OidcThumbprint' type model to v1 and v2alpha1
- Add 'OidcThumbprintInput' type model to v1 and v2alpha1
- Add 'OidcThumbprint' resource model to v1 and v2alpha1

## 0.1.434

- Update model version v0.0.386
  - Add `RootVolume` attribute to `AWSNodePool` model
- Update model version v0.0.387
  - Add `ProjectNumber` field to `WifConfig` model

## 0.1.433

- Update model version v0.0.384
  - Add clusters_mgmt API model for v2alpha1
- Update model version v0.0.385
  - Update WIF endpoint path
  - Remove WIF templates endpoints

## 0.1.432

- Update model version v0.0.383
  - Add `Kind` and `Id` field to GCP `Authentication` structure

## 0.1.431

- Update model version v0.0.382
  - Add `Authentication` field to GCP model

## 0.1.430

- Added WIF endpoints and resources

## 0.1.429

- Update model version v0.0.380
  - Add `Architecture` attribute to `MachineType` model
  - Add `ReleaseImages` attribute to `Version` model

## 0.1.428

- Update model version v0.0.379
  - Add `Ec2MetadataHttpTokens` to `AWSNodePool` resource

## 0.1.427

- Update model version v0.0.378
  - Add `MultiArchEnabled` attribute to `Cluster` model

## 0.1.426

- Update for Windows support including systemcertpool

## 0.1.425

- Update Windows certificates

## 0.1.424

- Update model version v0.0.377
  - Change type of the `OSDiskSizeGibibytes` attribute in the `AzureNodePool` resource from String to Integer.

## 0.1.423

- Update model version v0.0.376
  - Add `AzureNodePool` to `NodePool` type

## 0.1.422

- Update model version v0.0.375
  - Add `ManagementUpgrade` parameters to the `NodePool` model to support additional upgrade configuration.
  - Support server name inference for regionalized OCM redirects
  - Add `AdditionalAllowedPrincipals` to `AWS` type to support additional allowed principal ARNs to be added to the hosted control plane's VPC Endpoint Service.

## 0.1.421

- Update model version v0.0.374
  - Add `CreationTimestamp` and `LastUpdateTimestamp` to `RolePolicyBinding` type
  - Add `access_transparency` service

## 0.1.420

- Update model version v0.0.373
  - Add `subnet_resource_id` to `Azure` resource
  - Add `network_security_group_resource_id to`Azure` resource

## 0.1.419

- Update model version v0.0.372
  - Exposed the `/api/clusters_mgmt/v1/clusters/{id}/kubelet_configs` endpoint
  - Added `name` field to `KubeletConfig`
  - Added `kubelet_configs` field to `NodePool`

## 0.1.418

- Update model version v0.0.371
  - Add Tags to the AWSMachinePool model to support custom AWS tags for day 2 creation of machine pools

## 0.1.417

- Add RolePolicyBindings to the AWS resource model to support STS Arbitrary Policies feature.

## 0.1.416

- Update windows api.openshift certs

## 0.1.415

- Update model version v0.0.369
  - Update metamodel version to 0.0.60
  - [OCM-6294] add /load_balancer_quota_values endpoint
  - [OCM-7027] feat: document pagination and ordering support for break glass
  - [OCM-7144] Add /storage_quota_values endpoint
  - Add Azure resource to Cluster resource.
  - Fix spacing in description of Azure's ManagedResourceGroupName

## 0.1.414

- Update metamodel version v0.0.60

## 0.1.413

- Update model version v0.0.366
  - Fix Default Capabilities.

## 0.1.412

- Update model version v0.0.365
  - [OCM-6763] Add default capability resource to SDK.

## 0.1.411

- Upgrade pgx version and other dependencies

## 0.1.410

- Update model version v0.0.364
  - Add `BreakGlassCredentials` to the `Cluster` resource model.

## 0.1.409

- Update model version v0.0.363
  - Add `NodeDrainGracePeriod` to the `NodePool` model.
- Update model version v0.0.362
  - Changed `UserName` attribute for TokenClaimMappings to `Username`.

## 0.1.408

- Update model version v0.0.361
  - Add `Scope` attribute to `ReservedResource`.
  - Add `Scope` attribute to `ClusterAuthorizationRequest`.

## 0.1.407

- Update model version v0.0.360
  - Add `ComponentRoutes` attribute to `Ingress`

## 0.1.406

- Migrate Keychain `securestore` usage to non-CGO libraries

## 0.1.405

- Update model version v0.0.359
  - Add `ExternalAuthConfig` resource to `Cluster` model.
  - Add `DomainPrefix` to `Cluster` model.

## 0.1.404

- Update model version v0.0.357
  - Add `ExternalAuth` to `ExternalAuthConfig` model

## 0.1.403

- Update model version v0.0.356
  - Reverting change to remove provision shard from cluster

## 0.1.402

- Update model version v0.0.355
  - Removed undefined api calls from the model
  - Add support to `securestore` that allows the caller to define a keyring target
  - Additional `securestore` Error Standardization
  - Add error handling to `securestore` for denied Keychain access due to permissions

## 0.1.401

- Update model version v0.0.353
  - Added support for `PackageImage` for `clusters_mgmt`

## 0.1.400

- Update model version v0.0.352
  - Remove `StatusBoard` `fullname` search parameter.
  - Deprecate `Notify` resource.

## 0.1.399

- Additional error handling for `securestore`

## 0.1.398

- Add regions support from ocm shards
- Don't error when response is 204 and no content-type

## 0.1.397

- Add `NewUnauthenticatedConnectionBuilder` function to allow creating a
  `Connection` without client side authentication

## 0.1.396

- Remove redundant fields from /notify_details
- Add `ExternalAuthConfig` to `Cluster` model.

## 0.1.395

- Add `SubnetOutposts` and `AvailabilityZoneTypes` to `aws_node_pool_type` and `aws_machine_pool_type` resources.

## 0.1.394

- Added Device Code flow to `authentication`
- Update model version v0.0.347
  - Add `HostedControlPlaneDefault` boolean to `Version` Type model.

## 0.1.393

- Add authentication using OAuth2 and PCKE
- Add secure token storage

## 0.1.392

- Update model version v0.0.346
  - Modify notify_details response

## 0.1.391

- Update model version v0.0.345
  - Add `validate_credentials` resource to `AwsInquiries`
- Update model version v0.0.344
  - Add the /notify_details endpoint to the SDK

## 0.1.390

- Update model version v0.0.343
  - Add `Platform`to `subnet_network_verification_type` resource

## 0.1.389

- Update model version v0.0.342
  - Add `Search` and `Order` methods to List `/api/clusters_mgmt/v1/clusters/{id}/node_pools`

## 0.1.388

- Update model version v0.0.341
  - Add DELETE /api/addons_mgmt/v1/clusters/{id}/addons endpoint

## 0.1.387

- Update model version v0.0.340
  - Add get `Platform` to `network_verification_type` resource

## 0.1.386

- Update model version to v0.0.339
  - Add `MachineTypes` to `GCPInquiriesClient` endpoints

## 0.1.385

- Update model version to v0.0.338
  - Add `ProductTechnologyPreviews` and `ProductMinimalVersions` endpoints

## 0.1.384

- Updated client for `KubeletConfig` to align `post` and `update` function signatures

## 0.1.383

- Update model version to v0.0.336
  - Added `security` field to Cluster Service GCP field

## 0.1.382

- Update model version to v0.0.335
  - Add `doc_references` field in `LogEntry`
  - Add tags to subnet network verification resource

- Update model version to v0.0.334
  - Add Search method to status_board status_updates model

## 0.1.381

- Update model version to v0.0.333
  - Add `/api/clusters_mgmt/v1/clusters/{id}/kubelet_config` endpoint
  - Add `KubeletConfig` struct
  - Update `Cluster` struct to be able to optionally embed the `KubeletConfig` struct

## 0.1.380

- Update model version v0.0.332
  - Add `AdditionalInfraSecurityGroupIds` to `AWS` type
  - Add `AdditionalControlPlaneSecurityGroupIds` to `AWS` type
- Update model version v0.0.331
  - Add `Search` method to `status_board` `products_resource`, `applications_resource`, and `services_resource models`

## 0.1.379

- Require Go 1.21

## 0.1.378

- Update model version v0.0.330
  - Add `Update` method to `HypershiftConfig` resource

## 0.1.377

- Update model version v0.0.329
  - Add get `ClusterId` to `network_verification_type` resource

## 0.1.376

- Update model version v0.0.328
  - Add get `VPC` to `Cluster` resource

## 0.1.375

- Update model version v0.0.327
  - Add `BestEffort` to method `Delete` in `Cluster`

## 0.1.374

- Update model version v0.0.326
  - Add `BackplaneURL` to `Environment` type

## 0.1.373

- Update model version to v0.0.325
  - Add `OrganizationId` to `FeatureReviewRequest` type

## 0.1.372

- Update model version to v0.0.324
  - Add `CreatedAt` to `LogEntry` type
  - Add `CreatedBy` to `LogEntry` type

## 0.1.371

- Update model version to v0.0.323
  - Add `GCPMarketplaceEnabled` to `version` type

## 0.1.370

- Update model version to v0.0.322
  - Add AdditionalComputeSecurityGroupIds to AWS type
  - Add AdditionalSecurityGroupIds to AWS Machine Pool type
  - Add AwsSecurityGroups to VPC type

## 0.1.369

- Update model version to v0.0.321
  - Exposes `/api/clusters_mgmt/v1/aws_inquiries/sts_account_roles` in the SDK

## 0.1.368

- Update model version v0.0.318
  - Add `ImageOverrides` to `Version` type

## 0.1.367

- Windows: Update SSO CA
- Update model version v0.0.315
  - Add DisplayName and Description properties to `BillingModelItem`

## 0.1.366

- Update model version v0.0.314
  - Add new resources and a type for `BillingModelItem`

## 0.1.365

- Update model version v0.0.312
  - Added support for `AddonInstallations` endpoints for `addons_mgmt`.
  - Updated APIs for `AddonStatus`, `AddonStatusCondition`, `AddonSubOperator` and `AddonVersion`.

## 0.1.364

- Update model version v0.0.311
  - Add a new resource to OSL clusters/cluster_log

## 0.1.363

- Modify SelfAccessReview to return IsOCMInternal field

## 0.1.362

- Redact aws access and secret access keys from debug logs

## 0.1.361

- Update model version v0.0.309
  - Modify access review response to include `is_ocm_internal` field.
  - Add the remainder of cluster-autoscaler parameters.

## 0.1.360

- Update model version v0.0.307
  - Move `PrivateHostedZoneID` and `PrivateHostedZoneRoleARN` to `aws_type` resource

## 0.1.359

- Update model version v0.0.306
  - Fix upgrade related constants JSON output to align with existing values
- Update model version v0.0.305
  - Add `PrivateHostedZoneID` and `PrivateHostedZoneRoleARN` to `cluster_type` resource

## 0.1.358

- Update model version v0.0.304
  - Add upgrade related constants also for `NodePoolUpgradePolicy`.
  - Change DNS domain field names.

## 0.1.357

- Update model version v0.0.303
  - Add upgrade related constants.
- Update model version v0.0.302
  - Add property `MarketplaceGCP` to `billing_model_type` in `clusters_mgmt` and `accounts_mgmt`
  - Document `GovCloud`, `KMSLocationID` and `KMSLocationName` fields to `CloudRegion`
  - Document `fetchRegions=true` to `cloud_providers`

## 0.1.356

- Update model version v0.0.301
  - Update name for `ClusterStsSupportRole` resource and type to `StsSupportJumpRole`

## 0.1.355

- Update model version v0.0.300
  - Add `UserDefined` in dns domain resource

## 0.1.354

- Update model version v0.0.299
  - OCM-209 | feat: Add cluster autoscaler API resources
  - OCM-209 | feat: Add autoscaler locator in cluster resource

## 0.1.353

- Prevent connection leak in retry wrapper

## 0.1.352

- Update model version v0.0.297
  - Add managed ingress attributes
  - Fix `fetchLabels` and `fetchAccounts` url parameter names
  - Add `ClusterStsSupportRole` resource and type

## 0.1.351

- Update model version v0.0.296
  - Add json annotation to `DeleteAssociatedResources` parameter in account resource

## 0.1.350

- Update model version v0.0.294
  - Add `DeleteAssociatedResources` locator to account resource
- Update model version v0.0.295
  - Update `ReservedAt` to `ReservedAtTimestamp` in dns domain type
- Update metamodel version 0.0.59:
  - Honor`@http` annotation for query parameters

## 0.1.349

- Update model to version v0.0.293
  - Add label list to OSDFM cluster request payloads
  - Replace references to labels in OSDFM cluster structs with the labels themselves
  - Fix typos in OSDFM cluster Label struct fields
  - Add HashedPassword field to clusters_mgmt to provide encrypted value
  - Add cluster autoscaler structs

## 0.1.348

- Update model version v0.0.291
  - Add Reason to access review responses
  - Enable users to provide both hashed and plain-text passwords
  - API model for network verification

## 0.1.347

- Update model version v0.0.290
  - Rename `MachineTypeRootVolume` to `RootVolume`
  - Put `RootVolume` in `ClusterNodes`
  - add contracts to cloud accounts (#765)

## 0.1.346

- Update model version v0.0.289
  - Add Load balancer type to Ingress model
  - remove unused API endpoints

## 0.1.345

- Update model version v0.0.288
  - Add `DNSDomains` resource to the `root_resource`.
  - Complete OSD FM api for SDK usage.

## 0.1.344

- Update model version v0.0.287
  - Add Htpasswd to Cluster

## 0.1.343

- Update model version v0.0.286
  - Add MachineTypeRootVolume to MachinePool

## 0.1.342

- Update model version v0.0.285
  - Changed DNS Domain from Class to a Struct.
  - Change dns domain type to class and remove ID.

## 0.1.341

- Update model version v0.0.282
  - Changing parameter name from HttpTokensState to Ec2MetadataHttpTokens

## 0.1.340

- Update model version v0.0.281
  - Add `RootVolume` of type `MachineTypeRootVolume` to `MachineType` type.

## 0.1.339

- Update model version v0.0.280
  - Add `HttpTokensState` to `AWS` resource.

## 0.1.338

- Windows: Update API CA

## 0.1.337

- Windows: Update SSO CA

## 0.1.336

- Update model version v0.0.279
  - Add `AuditLog` to `AWS` resource.
  - Add `RoleArn` attribute to the `AuditLog` model.

## 0.1.335

- Update model version v0.0.278
  - Add InflightChecks locator to cluster resource
  - Add BillingAccountID to AWS model

## 0.1.334

- Update model version v0.0.276
  - Add delete method to `Account` resource.
  - Add `tuning_configs` endpoints.
  - Add `tuning_configs` field to Node Pools.

## 0.1.333

- Update model version v0.0.275
  - Add pending delete cluster API.
- Update model version v0.0.274
  - Add `Subnets` property to the CloudProviderData model.

## 0.1.332

- Update model version v0.0.273
  - update metamodel version 0.0.57
  - remove circular dependencies from clusters mgmt

## 0.1.331

- Update model to v0.0.272
  - adding quota version

## 0.1.330

- Update model to v0.0.271
  - Adding `version_inquiry` endpoint to Managed Services.

## 0.1.329

- Update model to v0.0.270
  - adding quota auth to root resource model

## 0.1.328

- Update model to v0.0.269
  - Add `DeleteProtection` resource to `Cluster` resource.
  - adding quota auth models

## 0.1.327

- Update model to v0.0.268
  - Replace `OidcConfigId` for `OidcConfig` in `STS` resource.

## 0.1.326

- Update model to v0.0.267
  - Add `OidcConfigId` to `STS` resource.
  - Remove `OidcPrivateKeySecretArn` from `STS` resource.

## 0.1.325

- Update model to v0.0.266
  - Adjust `Oidc Configs` endpoints.

## 0.1.324

- Update model to v0.0.265
  - Rename `HypershiftEnabled` boolean to `HostedControlPlaneEnabled` in `Version` Type model.

## 0.1.323

- Update model to v0.0.264
  - Add `Hosted Oidc Configs` endpoints.

## 0.1.322

- Update model to v0.0.263
  - Add `HypershiftEnabled` boolean to `Version` Type model.

## 0.1.321

- Update model to v0.0.262
  - Add `Control Plane Upgrade Scheduler` endpoints.

## 0.1.320

- Update to model v0.0.261
  - Add `commonAnnotations` and `commonLabels` to addons
- Update to Addon structs and openapi.json for supporting
  - `commonAnnotations`
  - `commonLabels`

## 0.1.319

- Update to model v0.0.260
  - Add `ManagedPolicies` field to the `STS` type model.

## 0.1.318

- Update to model v0.0.259
  - Add master and infra instance types to cluster nodes
- Update to model v0.0.258
  - Export cluster name for mgmt, mgmt_parent, and svc clusters

## 0.1.317

- Update to model v0.0.257
  - Add `ByoOidc` type to Cluster type model
  - Add addon upgrade policy to clusters_mgmt
  - Add `Labels` and `Taints` to NodePool type

## 0.1.316

- Update to model v0.0.256
  - Add `LogType` field to Cluster Log type model
  - Fix Addon status type and value constants

## 0.1.315

- Update to model v0.0.255
  - Add `Version` field to node pool

## 0.1.314

- Update to model v0.0.254
  - Add `PrivateLinkConfiguration` type with related endpoints

## 0.1.313

- Update to model v0.0.253
  - Update Permission resource attributes
    - Rename ResourceType to Resource

## 0.1.312

- Update to metamodel 0.0.57

## 0.1.311

- Update to model v0.0.252
  - Update `STS` resource attributes
    - Remove `BoundServiceAccountSigningKey`
    - Remove `BoundServiceAccountKeyKmsId`
    - Rename `BoundServiceAccountKeySecretArn` to `OidcPrivateKeySecretArn`

## 0.1.310

- Update to model v0.0.251
  - Update `NodePool` with status attributes
  - Added `current_compute` attribute in `ClusterStatus` for hosted clusters.
  - Added missing variable to `addon environment variable` for addons mgmt

## 0.1.309

- Update to model v0.0.250
  - Add `Addon Inquiries API` to `addons_mgmt`

## 0.1.308

- Update to model v0.0.249
  - Add `BoundServiceAccountKeySecretArn` attribute to the `Sts` model.

## 0.1.307

- Update to model v0.0.248
  - Add `AwsEtcdEncryption` type model and reference from `AWS`.
  - Add `Enabled` attribute to `STS` model.

## 0.1.306

- Update to model v0.0.247
  - Corrected `Metrics` type on `DeletedSubscription`

## 0.1.305

- Update to model v0.0.246
  - Add Search to `Capabilities` resource

## 0.1.304

- Update to model v0.0.245
  - Add `BoundServiceAccountKeyKmsId` attribute to the `Sts` model.

## 0.1.303

- Update to model v0.0.244
  - Add `ARN` attribute to the `AWSSTSPolicy` model.

## 0.1.302

- Update to model v0.0.243
  - Add `BoundServiceAccountSigningKey` attribute to the `Sts` model.
- Update to model v0.0.242
  - Add `AddonNamespace` resource model.
  - Add `CommonLabels` attribute to the `Addon` model.
  - Add `CommonAnnotations` attribute to the `Addon` model.
  - Add `MachineType` locator on `MachineTypes` model.

## 0.1.301

- Update to model v0.0.241
  - Add `DeletedSubscriptions`
  - Add `AddonCluster`
  - Add `AddonStatus`

## 0.1.300

- Update PR check to include go v1.19
- Update goimports to v0.4.0
- Update to model v0.0.240
  - Fix `AddonConfig` on `AddonConfigType` resource model.

## 0.1.299

- Update to model 0.0.239
  - Fixes for `NodePoolAutoScaling` and `AWSNodePool`.

## 0.1.298

- Update to model 0.0.238
  - `NodePool` fixes.

## 0.1.297

- Update to model 0.0.237
  - Add `NodePool` resource types.
  - Add `NodePools` locator to `Cluster` type.

## 0.1.296

- Update to model 0.0.236
  - Add extra fields to label model:
    - Type
    - ManagedBy
    - AccountID
    - SubscriptionID
    - OrganizationID

## 0.1.295

- Update to model 0.0.235
  - Add `capabilities` resource model.

## 0.1.294

- Update to model 0.0.233
  - Add SupportsHypershift property to CloudRegion model.

## 0.1.293

- Update to model 0.0.232
  - Modify `availabilityZone` property in CloudProviderData model from `string` to `[]string`.

## 0.1.292

- Update to model 0.0.231
  - Add `AvailabilityZone` property to CloudProviderData model.
  - Add `Public` property to Subnetwork model.

## 0.1.291

- Update to model 0.0.230
  - Add creation timestamp and modification timestamp to provision shard
  - Add pull secret for addon version
  - Add addon secret props for addon version config
  - Add additional catalog sources for addon version
  - Add addon parameter condition

## 0.1.290

- Update to model 0.0.229
  - Add Addon Management models
  - Add GCP Encryption Keys to cluster model
- Add client for Fleet Management service
- Add client for Addons Management service

## 0.1.289

- Update to model 0.0.228
  - Add hypershift endpoint with its ManagementCluster.
  - Align hypershift case usage.
  - [Hypershift] Expose /manifests
  - Added expiry setting to managed service clusters.
  - Added Manifests to external_configuration.
  - Add marketplace specific enum for clusters mgmt
  - Add Search method to ProvisionShards

## 0.1.288

- Windows: Update CA

## 0.1.287

- Update to model 0.0.223:
  - Add Version property to CloudProviderData model.
  - Add InfraID property to Cluster model.
  - Drop deprecated DisplayName property from ClusterRegistration model.
  - Add ConsoleUrl and DisplayName properties to ClusterRegistration model and correct documentation.

## 0.1.286

- Update to model 0.0.220:
  - Add `ManagedBy` property in RoleBinding type

## 0.1.285

- Update to model 0.0.219:
  - Add billing model to addon installations

## 0.1.284

- Update to model 0.0.218:
  - Change provision shard to include kube client configurations and server URL.

## 0.1.283

- Update to model 0.0.217:
  - Change provision shard to include kube client configurations.
  - Add GCP volume size to flavour API.
  - Add fleet manager related structures and API.

## 0.1.282

- Update to model 0.0.215:
  - Add hypershift config to provision shard API.

## 0.1.281

- Update to model 0.0.214:
  - Add locator `label` to Generic Labels resource.

## 0.1.280

- Update to model 0.0.213:
  - Add update function to provision shard API.

## 0.1.279

- Update to model 0.0.212:
  - Add status to provision shard API.

## 0.1.278

- Update to model 0.0.211:
  - Remove `DisplayName` field from Cluster model.
  - Add API for adding and removing a provision shard.

## 0.1.277

- Update to model 0.0.209:
  - Add `capabilities` field to account type.

## 0.1.276

- Update to model 0.0.208:
  - Add `delete` method to registryCredentials type.

## 0.1.275

- Update to model 0.0.207:
  - Add `Subnets` field to machinePool type.

## 0.1.274

- Update to model 0.0.206:
  - Add `ExcludeSubscriptionStatuses` field to ResourceReview type.
  - Add `dry_run` flag to ClusterDeleteRequest type.

## 0.1.273

- Update to model 0.0.205:
  - Add `BillingMarketplaceAccount` field to ReservedResource type.

## 0.1.272

- Update to model 0.0.204:
  - Remove volume type from flavour
  - Add Network Configuration for Managed Services

## 0.1.271

- Update to model 0.0.203:
  - Add `MarketplaceAWS`, `MarketplaceAzure`, `MarketplaceRHM` billing models.

## 0.1.270

- Update to model 0.0.202:
  - Add `CloudAccount` type.
  - Add `CloudAccounts` field to QuotaCost type.
  - Add `BillingMarketplaceAccount` field to Subscription type.

## 0.1.269

- authentication: Allow client credential grants with basic auth
- Update to model 0.0.201:
  - Adding groups claim to openID IDP

## 0.1.268

- Update to model 0.0.200:
  - Add `hypershift.enabled` field to the cluster type.

## 0.1.267

- Update to model 0.0.199:
  - Fix cred request api model parameters

## 0.1.266

- Update to model 0.0.198:
  - Add cred request to api model
  - Add AWSRegionMachineTypes endpoint to api model
- windows: Update certificates

## 0.1.265

- Update to model 0.0.197:
  - Change inflight check type Details field to Interface

## 0.1.264

- Update to model 0.0.196:
  - Added Machine Pool Security Group Filters for Machine Pools and Cluster Nodes
  - Drop RoleARN from AddOnInstallation

## 0.1.263 Apr 19 2022

- Update to model 0.0.195:
  - Added Import method to the HTPasswd IDP user collection.
  - Added credential request type and updated the addon type to include it

## 0.1.262 Apr 14 2022

- Update to model 0.0.194:
  - Added availability zone fields to managed service cluster struct.

## 0.1.261 Apr 14 2022

- Update to model 0.0.193:
  - Add limitedSupportReasonCount to cluster status struct.
  - Add inflight check API.

## 0.1.260 Apr 11 2022

- Update to model 0.0.191:
  - Fix JSON representation of log severity.

## 0.1.259 Apr 8 2022

- Update to model 0.0.190:
  - Fix JSON names of identity provider types.
  - Add enable minor version upgrades flag to upgrade policy.

## 0.1.258 Apr 5 2022

- Update to model 0.0.189:
  - Added QuotaRules to ocm-sdk-go
  - Added no_proxy field to the proxy project
  - Added errors resource.
  - Added errors support for status-board.

## 0.1.257 Apr 1 2022

- Add `web-rca` examples.
- Update to metamodel 0.0.53:
  - Don't consider `Status` and `Error` built-in request parameters.
- Update to metamodel 0.0.54:
  - Remove generation of experimental server code.
- Update to model 0.0.188:
  - Add Status query param for incidents resource.

## 0.1.256 Mar 31 2022

- Update to model 0.0.187:
  - Add new `web-rca` service.

## 0.1.255 Mar 30 2022

- Update to model 0.0.186:
  - Add ManagementCluster to ProvisionShard

## 0.1.254 Mar 30 2022

- Update to model 0.0.185:
  - Fixes to Cloud Resources endpoints.

## 0.1.253 Mar 29 2022

- Update to model 0.1.184:
  - Adding Cloud Resources endpoints.

## 0.1.252 Mar 17 2022

- Update to model 0.1.183:
  - Added field for parameters to be specified for managed services.

## 0.1.251 Mar 15 2022

- Update to model 0.1.182:
  - Adding `service_mgmt` service.

## 0.1.250 Mar 14 2022

- Update to model 0.0.181:
  - Add aws sts policy
  - Add ReleaseImage to Version

## 0.1.249 Mar 9 2022

- Update to model 0.0.180:
  - Add CloudProvider info to ProvisionShards.

## 0.1.248 Mar 9 2022

- Update to model 0.0.179:
  - Fix cluster logs URL, should be `cluster_logs` instead of `cluster_logs_uuid`.

## 0.1.247 Mar 8 2022

- Update to metamodel 0.0.52:
  - Add support for annotations.
  - Add `@json` and `@http` annotations.
  - Add `@go` annotation.
  - Add original text to names.
  - Add `Impersonate` method to support the `Impersonate-User` header.

## 0.1.246 Mar 7 2022

- Update to model 0.0.178:
  - Add `managed_service` field to add-on type.
  - Add `credentials_secret` field to add-on type.
  - Add `region` field to provision shard.

## 0.1.245 Mar 3 2022

- Update to model 0.0.177:
  - Fix update method of environment endpoint, should be `Update` instead of
    `Patch`.
  - Remove unimplemented `POST /api/service_logs/v1/cluster_logs/clusters/{uuid}/cluster_logs`
    method.

## 0.1.244 Mar 02 2022

- Update to model 0.0.176
  - adding new endpoint for 'environment'

## 0.1.243 Mar 02 2022

- Update to model 0.0.175
  - adding new apis for addon config attribute
  - adding list of requirements to addon parameter options
  - adding name fields to VPCs and Subnetworks
  - rename addon env object

## 0.1.242 Feb 16 2022

- Update to model 0.0.174
  - adding rhit_account_id to Account class

## 0.1.241 Feb 11 2022

- Update to model 0.0.173:
  - addons: Support attributes necessary for STS.
  - Add ProductIds param to Status Resource.
  - Add Role bindings to Subscription.

## 0.1.240 Feb 4 2022

- Update to model 0.0.172:
  - Remove deprecated `SKUs` endpoint.
  - Remove deprecated quota summary resource and type.
  - Add QuotaVersion to ClusterAuth.
  - Allow adding/removing operator roles.

## 0.1.239 Feb 3 2022

- Update to metamodel 0.0.51:
  - Check for `io.EOF` before trying to parse response body.

## 0.1.238 Jan 28 2022

- Update to model 0.0.170:
  - Add `ServiceInfo` type to status board service.

## 0.1.237 Jan 25 2022

- Update to metamodel 0.0.50:
  - Fix format of date query parameters so that it is RFC3339.

## 0.1.236 Jan 25 2022

- Update to model v0.0.169
  - Version gate type: Add warning message field

## 0.1.235 Jan 11 2022

- Install metamodel with `go install`
- Update to model v0.0.168
  - Fix description of various API attributes
  - OVN: Add network type selection
  - adding field to hold validation error message

## 0.1.234 Jan 4 2022

- Update to version 4 of JWT library.

  Note that this is a backwards compatibility breaking change because the
  `jwt.Token` type is used as parameter in the `authentication.ContextWithToken`
  and `authentication.TokenFromContext` methods. If you are using those methods
  then you will have to change your code to import `github.com/jwt-go/jwt/v4`
  instead of `github.com/jwt-go/jwt`.

- Update to Ginkgo 2.

  This change affects only the tests, but if you are using _Ginkgo_ in your
  project and you are still using version 1 then you may find issues when
  running the `ginkgo` command because both versions of the library will be
  included in your tests binaries and both will try to use the `flag` package to
  create conflicting command line flags. If that is the case the best approach
  is to update your project go use version 2 as well.

## 0.1.233 Dec 28 2021

- Update to model 0.0.167:
  - Change field name in version gate agreement link to version gate.

## 0.1.232 Dec 26 2021

- Update to model 0.0.166:
  - Change version gate agreement URL.

## 0.1.231 Dec 22 2021

- Update to model 0.0.165:
  - Move version gates to be top level resource.
  - Add version raw id prefix to version gates.
- Update version of JSON iterator.
- Update SQL drivers.
- Don't use github.com/ghodss/yaml
- Require Go 1.16

## 0.1.230 Dec 21 2021

- Update to model 0.0.163:
  - Add support for deleting version gate.

## 0.1.229 Dec 21 2021

- Update to model 0.0.162:
  - Add support for adding version gate.

## 0.1.228 Dec 20 2021

- Update to metamodel 0.0.46:
  - Remove unused imports.
  - Check result of `Flush` method.
  - Cancel poll context.
  - Avoid some ineffectual assignments.
  - Explicitly use `jsoniter` package selector.

- Update to model 0.0.161:
  - Support for version gate agreements.

## 0.1.227 Dec 20 2021

- Update to model 0.0.160:
  - Change version gates URL.

## 0.1.226 Dec 20 2021

- Update to model 0.0.159:
  - Support for version gates.

## 0.1.225 Dec 19 2021

- MatchJQ should require at least one result
- Update to model 0.0.158:
  - Adding subnetworks to vpc inquiry
  - Add statuses path to service model, add some comments.
  - [SDB-2509] Update OSL API schema to be compatible with ocm-sdk-go

## 0.1.224 Dec 10 2021

- Support pull-secret access token as a valid token.

- Update to model 0.0.156:
  - Add `updates` method to status board product resource.
  - Fix status get method of status board.

## 0.1.223 Dec 9 2021

- Update to model 0.0.155:
  - Add `status_board` service.

## 0.1.222 Dec 3 2021

This version doesn't contain changes to the functionality, only to the
development and build workflows:

- Rename `master` branch to `main`.

  To adapt your local repository to the new branch name run the following
  commands:

  ```shell
  git branch -m master main
  git fetch origin
  git branch -u origin/main main
  git remote set-head origin -a
  ```

- Automatically add changes from `CHANGES.md` to release descriptions.

## 0.1.221 Dec 1 2021

- Modify `func (c *Connection) Close()` to return nil in case the connection is already closed.

## 0.1.220 Nov 25 2021

- Added utilities to test with `jq` expressions and JSON patches.

## 0.1.219 Nov 22 2021

- Update to model 0.0.153:
  - Enable FIPS mode

## 0.1.218 Nov 22 2021

- Update to metamodel 0.0.44:
  - Check for loops in locator paths.
  - Add `Empty` method to builders.

## 0.1.217 Nov 18 2021

- Update to model 0.0.152
  - Update type `resource` to `clusterResources`
  - Revert "Add Name field to LDAP identity provider"
  - Remove addon install mode `singleNamespace`
  - Add addon install mode `ownNamespace`
  - Add channel to addon version class

## 0.1.216 Nov 15 2021

- Update to metamodel 0.0.43:
  - Add `status` attribute to errors.

## 0.1.215 Nov 7 2021

- Update to model 0.0.151:
  - Add Name field to LDAP identity provider

## 0.1.214 Oct 27 2021

- Update to model 0.0.150:
  - Fix addon installation version (addon_version vs version)
  - Remove no_proxy attribute from SDK
  - Add body to the external tracking event

## 0.1.213 Oct 26 2021

- Update to model 0.0.42
  - Accept iterator as parameter in `helpers.NewIterator`
- Retry when `REFUSED_STREAM`

## 0.1.212 Oct 12 2021

- Fix typo in HTTP internal server error message
- Copy data package from CLI
- Add support for digging inside maps
- Replace expired CA certs for windows
- Update model to 0.0.149
  - Add Addon Versions

## 0.1.211 Oct 05 2021

- Update to model 0.0.148:
  - Revert archived clusters endpoint.

## 0.1.210 Sep 27 2021

- Update to model 0.0.147:
  - Add missing connection to clusters collection in the `service_logs` service.

## 0.1.209 Sep 21 2021

- Avoid hard-coded private keys
- Bump API model to v0.0.146
  - Add Status to AddOnRequirement

## 0.1.208 Sep 13 2021

- Add Archived cluster endpoint
- Add cluster waiting state

## 0.1.207 Sep 13 2021

- Add cluster-wide proxy

## 0.1.206 Sep 9 2021

- Update model to v0.0.143:
  - Add Add() method to Limited Support Reasons resource

## 0.1.205 Sep 9 2021

- Update model to v0.0.142:
  - Add Limited Support Reason API

## 0.1.204 Aug 25 2021

- Change level of token retry messages to `DEBUG`.

## 0.1.203 Aug 23 2021

- Retry requests with body
- Update to metamodel 0.0.39
  - Add Details to Errors

## 0.1.202 Aug 18 2021

- Update model to v0.0.141:
  - Add support for EndOfLifeTimestamp in version

## 0.1.201 Aug 17 2021

- Retry when server sends go away before settings.
- Reject impersonation.

## 0.1.200 Aug 11 2021

- Update model to v0.0.140:
  - Add check_optional_terms to TermsReview and SelfTermsReview
  - Add reduceClusterList to ResourceReview and SelfResourceReview

## 0.1.199 Aug 10 2021

Changes in this release are mainly intended to simplify packaging of the SDK in
Fedora, see [issue #421](https://github.com/openshift-online/ocm-sdk-go/issues/421)
for details.

- Use `golang-jwt/jwt` instead of `dgrijalva/jwt-go`.

  The `dgrijalva/jwt-go` library is no longer maintained and `golang-jwt/jwt` is
  a community maintained fork. See <https://github.com/dgrijalva/jwt-go/issues/462>
  for detailts.

  Parts of the public interface of the SDK use this library, so this is a
  backwards compatibility breaking change. Projects using the SDK will need to
  switch to the new library, specially if they are using the
  `context.ContextWithToken` or `context.TokenFromContext` functions. The change
  should only require changing the import paths, as the fork is fully compatible
  with the original library.

  A simple way to do the required changes is the following command:

  ```shell
  find . -name '*.go' | xargs sed -i 's|dgrijalva/jwt-go|golang-jwt/jwt|'
  ```

  This also addresses
  [CVE-2020-2610](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-26160[CVE-2020-26160),
  but that vulnerability doesn't currently affect the SDK because the
  authentication handler doesn't use the `aud` claim.

- Use [microcosm-cc/bluemonday](https://github.com/microcosm-cc/bluemonday)
  instead of [grokify/html-strip-tags-go](https://github.com/grokify/html-strip-tags-go)
  for HTML sanitizing.

- Use [json-iterator/go](https://github.com/json-iterator/go) instead of
  [c0b/go-ordered-json](https://gitlab.com/c0b/go-ordered-json) to ensure ordered
  JSON in debug output.

## 0.1.198 Aug 03 2021

- Update model to v0.0.139:
  - Add kms key arn to aws ccs cluster encryption

## 0.1.197 Jul 22 2021

- Update model to v0.0.138:
  - Add addon inquiries endpoints get and list

## 0.1.196 Jul 20 2021

- Update model to v0.0.137:
  - Add Cluster Resource Type
  - Remove Cluster Deployment GET

## 0.1.195 Jul 19 2021

- Update model to v0.0.136:
  - Added Spot Market Options to MachinePool

## 0.1.194 Jul 6 2021

- Update to version 0.0.38 of the metamodel in order to fix a conflict between
  `github.com/golang/glog` and `github.com/istio/glog` that prevents building
  packages that use the SDK.

## 0.1.193 Jul 6 2021

- Update the version number in the `version.go` file.

## 0.1.192 Jul 6 2021

- Fix initialization of metrics registerer.
- Update model to v0.0.135:
  - Add user name to service log `LogEntry`.

## 0.1.191 Jul 5 2021

- Add leadership flag.

## 0.1.190 Jun 23 2021

- Don't use refresh token if have client credentials.

## 0.1.189 Jun 23 2021

- Don't require refresh token for client credentials grant.

## 0.1.188 Jun 23 2021

- Update model to v0.0.134
  - Added InternalOnly flag to SubscriptionNotify

## 0.1.187 Jun 16 2021

- Update model to v0.0.133
  - Added capabilities support to Organization

## 0.1.186 Jun 11 2021

- Update model to v0.0.132
  - Added Disable workload monitoring
  - add authorizations feature review and self feature review models.
  - add locators for feature review and self feature review
  - Add ccs_only flag for cloud regions

## 0.1.185 Jun 10 2021

- Add `All` and `Any` functions to the logging package.

## 0.1.184 Jun 1 2021

- Update ocm-api-model to v0.0.129
  - Add cloud provider inquiries to api-model
  - sts: Add support role ARN

## 0.1.183 May 31 2021

- Update model to v0.0.128
  - Remove the Dashboards resource.
  - Add CCSOnly, GenericName fields to machine type.
  - Add AcceleratedComputing value to MachineTypeCategory enum.

## 0.1.182 May 25 2021

- Update model to v0.0.126
  - Add ClusterConfigurationMode type under ClusterStatus
  - sts: Change custom roles to instance roles

## 0.1.181 May 25 2021

- Add support for cookie authentication.

## 0.1.180 May 25 2021

- Remove `AccessToken` authentication.

## 0.1.179 May 20 2021

- Allow building with Go 1.13.

## 0.1.178 May 18 2021

- Update model to v0.0.125
  - Added Custom IAM roles For STS

## 0.1.177 May 13 2021

- Update of parameters to JobQueue#Push

## 0.1.176 May 12 2021

- Added Arguments to JobQueue#Pop

## 0.1.175 May 6 2021

- Update `Logger` interface to include a `Fatal` log level.
  - Fatal level will call `os.Exit(1)` after writing the message.
  - Fatal level is always active.

IMPORTANT: This version breaks backwards compatibility in the `Logger`
interface, as all implementations now require a `Fatal` method to be implemented.

## 0.1.174 Apr 14 2021

- Miscellaneous fixes to JobQueue service.
- Bump ocm-api-model to v0.0.121

## 0.1.173 May 03 2021

- Bump ocm-api-model to v0.0.119
  - STS: Support attributes to allow STS clusters

## 0.1.172 May 03 2021

- Bump ocm-api-model to v0.0.118

## 0.1.171 Apr 14 2021

- Add JobQueue service.

## 0.1.170 Apr 14 2021

- Add `Tolerance` parameter to authentication handler.

## 0.1.169 Apr 13 2021

- Bump ocm-api-model to v0.0.115
  - Add event_code and site_code to TermsReviewRequest type
  - Add new SelfTermsReviewRequest type

## 0.1.168 Apr 6 2021

- Bump ocm-api-model to v0.0.114
  - related-resources: Add resource type and cloud provider
  - event: Track ad-hoc authenticated events

## 0.1.167 Apr 6 2021

- Move token logic to separate transport wrapper.
- Use defaults from authentication package.
- Update to model 0.0.113:
  - Add RelatedResources struct to QuotaCost.

## 0.1.166 Mar 30 2021

- Move client selection logic to separate type
- Update ocm-api-model to v0.0.112
  - Add Options to AddOnParameter type.
  - aws: Support PrivateLink for fully-private clusters

## 0.1.165 Mar 22 2021

- Fix wrong TLS server name (issue
  [356](https://github.com/openshift-online/ocm-sdk-go/issues/356)).

## 0.1.164 Mar 17 2021

- Change default user agent to `OCM-SDK`.
- Update to model 0.0.111:
  - Add subscription metrics.
  - Add `deprovision` and `force` parameters to delete cluster method.
  - Ensure all subscription fields are available.

## 0.1.163 Mar 5 2021

- Enable compression.
- Remove logger from metrics transport wrapper.
- Reorder execution of tests.
- Add metrics handler wrapper.
- Add `h2c` support.
- Enable HTTP/2.

## 0.1.162 Feb 22 2021

- Update to model 0.0.110:
  - organization: Add quota_cost endpoint resources

## 0.1.161 Feb 22 2021

- Update to model 0.0.109:
  - Remove deprecated 'upgrade_channel_group' field.
- Run tests in parallel.
- Add documentation of upgrade policy states.

## 0.1.160 Feb 17 2021

- Improve testing of metrics.
- Update to model 0.0.108:
  - Add `billing_model` attribute to the `ReservedResource` type.
  - Add `cluster_billing_model` attribute to the `Subscriptioin` type.

## 0.1.159 Feb 15 2021

- Add metrics package
- Update API model to v0.0.107:
  - add addon sub operator type

## 0.1.158 Feb 12 2021

- Downgrade from Go 1.15 to Go 1.14. This has been requested by users of the SDK
  that can't upgrade to Go 1.15 because it isn't available in RHEL 8 and because
  of the issues that Go 1.15 introduces related to the obsolete `CN` attribute
  of X.509 certificates. The only negative effect of this downgrade is that
  timeouts or deadlines set for requests sent using TLS over Unix sockets
  will be ignored.

## 0.1.157 Feb 8 2021

- Accept Empty Reader as non-nil req body
- Missplaced return after warning from send
- Add trusted CA certificates for Windows
- Update metamodel to v0.0.36:
  - Use Go 1.15
  - Add `documentedSupport` and `namedSupport`
  - Add `typedSupport`
  - Make reporter streams configurable
  - Add presence bitmap
- Update model to v0.0.106:
  - Add billing_model field to cluster type
  - subscriptions: Add label locator

## 0.1.156 Feb 4 2021

- Update to model 0.0.105:
  - Add cluster hibernation support
- Declare go 1.15 in go.mod
- connection: Skip loading SystemCertPool on Windows

## 0.1.155 Jan 27 2021

- Update to model 0.0.104:
  - Add addon requirement type.

## 0.1.154 Jan 26 2021

- Update to model 0.0.103:
  - Remove `cluster_admin_enabled` attribute from cluster type.
  - Add missing subscription, cluster authorization and plan attributes.

## 0.1.153 Jan 21 2021

- Add support for customizing the error responses of the authentication handler.
- Add support for connecting to the server using Unix sockets.

## 0.1.152 Dec 17 2020

- Update model to v0.0.102
  - add default value to add-on parameter type
  - Add upgrade channel group for a cluster

## 0.1.151 Dec 2 2020

- Move logging code to `logging` package
- Rename `Metrics` to `MetricsSubsystem`
- Add method to read metrics subsystem
- Load metrics subsystem configuration from file
- Load string if it doesn't look like a file
- Load `os.Stdin` in dump configuration example
- Reject URLs without scheme or host name
- Add redirection tests
- update model to 0.0.101

## 0.1.150 Nov 25 2020

- update model to 0.0.100

## 0.1.149 Nov 24 2020

- Fix issue in the method that returns the URL of a connection: it was returning
  an empty string when no alternative URLs were configured.

## 0.1.148 Nov 23 2020

- Rename `!shell` configuration tag to `!script`.
- Add `!yaml` configuration tag.

## 0.1.147 Nov 19 2020

- Add support for alternative URLs.
- Add support for loading trusted CA files.
- Add suppott for loading configuration from YAML file.

## 0.1.146 Nov 17 2020

- Add `EvaluateTemplate` function for tests
- Don't crash in debug mode deserializing an empty response
- Update model to v0.0.99
  - Add deletion add-on installation endpoint
  - Add Update method to addon installation resource
- Update metamodel tp v0.0.35
  - Update to version 4.8 of Antlr
  - Wrap errors

## 0.1.145 Nov 10 2020

- Update model to v0.0.98

## 0.1.144 Nov 2 2020

- Update model to v0.0.96
  - Add Enabled to AddOnParameter type.

## 0.1.143 Oct 27 2020

- Update api-model to v0.0.95
  - Add SubnetIDs field to AWS type.

## 0.1.142 Oct 26 2020

- Allow disabling keep alive connections in the SDK connection transport
- Update api-model to v0.0.94
  - version: Rename field from MOA to ROSA
  - [AMS] Add IncludeRedHatAssociates to SubscriptionNotify

## 0.1.141 Oct 21 2020

- Update api-model to v0.0.92
  - Add RawID field to Version type

## 0.1.140 Oct 14 2020

- Update api-model to v0.0.91
  - Remove redudant fields
  - flavours: Remove infra and compute nodes
  - Add AddOnParameter modal type Update AddOn to include list of AddOnParameters
  - Add AddOnInstallationParameter modal type Update AddOnInstallation to include list of AddOnInstallationParameters

## 0.1.139 Oct 11 2020

- Update api-model to v0.0.90
  - Add machine pools locator
  - Add compute node labels
- Interpret HTML entities in logged summary of error response
- Use new limits for content summary

## 0.1.138 Oct 5 2020

- Update api-model to v0.0.88
  - Add missing machine pools resource

## 0.1.137 Oct 5 2020

- Update api-model to v0.0.87
  - Add missing machine pool resource

## 0.1.136 Oct 5 2020

- Update metamodel to v0.0.34
  - Support numeric initialisms
- Update api-model to v0.0.86
  - Added New Error Message implementation
  - idp: Add HTPasswd provider
  - Uptdating SDK with GCP credentials

## 0.1.135 Oct 5 2020

- Update API model to v0.0.85
  - Add upgrade policy state

## 0.1.133 Sep 30 2020

- increase the limit/size of content summary
- Update metamodel to v0.0.33
  - json: Support NoContent on POST responses

## 0.1.132 Sep 24 2020

- Update model to v0.0.83
  - add external resources to add on type model
  - SDA-2952 - Add "hidden" option to AddOn

## 0.1.131 Sep 23 2020

- Support http proxy

## 0.1.130 Sep 21 2020

Update model to v0.0.82

- Added Install Error Details From Provisioner

## 0.1.129 Sep 21 2020

Request a token valid for longer than 1 min

## 0.1.128 Sep 14 2020

Go mod tidy

## 0.1.127 Sep 14 2020

Update to model v0.0.81

- Add key to label_type
- Remove ID from upgrade label

Also included as part of model v0.0.80

- Add upgrade policy type and resource
- Add terms review and self terms review
- Add dashboards summary

## 0.1.126 Sep 7 2020

Update to model v0.0.79

- Add 'available_upgrades' list to version type
- Add CCS type and Attribute to Cluster type

## 0.1.125 Sep 4 2020

Update to model v0.0.78

- Added New DNS_READY
- version: Add moa_enabled flag

## 0.1.124 Aug 28 2020

- Set token expiry function public
- Allow auth header of type AccessToken

## 0.1.123 Aug 23 2020

- Remove get tokens on first attempt log entry
- Update to metamodel v0.0.32
- Update to model v0.0.77
  - Add ChannelGroup attribute to version
  - Add avaialble AWS regions method

## 0.1.122 Aug 18 2020

- Better logging and metrics when retrying SSO
- Assume expiration is 0 when missing in the token

## 0.1.121 Aug 18 2020

- BROKEN: DO NOT USE

## 0.1.120 Aug 13 2020

- Update to model v0.0.76
  - Add missing link to provision shard

## 0.1.119 Aug 10 2020

- Add support for retry getting access token in case of http 5xx

## 0.1.118 Aug 7 2020

- Update to model v0.0.75
  - Added support_case resource
  - Added token_authorization to root_resource

## 0.1.117 Aug 5 2020

- Update to model v0.0.73
  - [CS] Add hive_config to the provision shard
  - [CS] Improving cluster logs endpoint
  - [AMS] Added token authorization endpoint

## 0.1.116 Aug 3 2020

- Added support for http PUT method
- Update to model v0.0.73
  - Add capability_review endpoint
  - Add support_cases endpoint

## 0.1.115 Jul 30 2020

- Update to metamodel v0.0.31
  - Adding List type to checkUpdate validator

- Update to model v0.0.72
  - Fix comment
  - Expose if a region supports multi AZ
  - Add Update Identity Provider
  - removing 'deprovision' suffix from logs endpoint
  - add post method to subscription resource
  - Add labels field to external configuration type
  - Implement Batch Patch Ingresses API endpoint

## 0.1.114 Jul 21 2020

- Update to model v0.0.71
  - Add API for getting cluster's provision shard
  - Add API for getting provision shards

## 0.1.113 Jul 14 2020

- Update to model v0.0.70
  - Add API for custerdeployment labels
  - add organization_id to cluster_registration
  - label: Fix erroneous file extensions
  - MachineType: Expose instance size enum

## 0.1.112 Jul 5 2020

- Update to model v0.0.69
  - Added top level sku_rules endpoint to AMS
  - Change the feature toggle API to be /feature_toggles/id/query using POST with org id as context

## 0.1.111 Jul 1 2020

- Update to model v0.0.67
  - [AMS] Added SkuCount to ResourceQuota type

## 0.1.110 Jun 30 2020

- Update to model v0.0.66
  - Change feature toggle query to be POST with payload containing organization ID

## 0.1.109 Jun 29 2020

- Update to model v0.0.65
  - Added Uninstall Log
  - Added syncset API
  - Update to metamodel v0.0.30

## 0.1.108 Jun 21 2020

- Update to model v0.0.64
  - Added Notify to root_resource in AMS

## 0.1.107 Jun 18 2020

- Update to model v0.0.63
  - cluster: Remove support for expiration_timestamp
  - Added top-level Notify endpoint to AMS

## 0.1.106 Jun 9 2020

- Update to metamodel v0.0.29:
  - pr_check: Lock in dependency versions for test pipeline
  - Fix setter for Poll request params

- Update to model v0.0.62:
  - Add subscription notify endpoint

- Update to model v0.0.61:
  - accounts_mgmt: Add 'fields' parameter to all list-requests
  - accounts_mgmt: Support for Labels resources

- Update to model v0.0.60:
  - Add parameters 'offset' and 'tail' to log resource

## 0.1.105 May 21 2020

- Update to model 0.0.59:
  - Add feature_toggle endpoint and api model

## 0.1.104 May 15 2020

- Update to model v0.0.58
  - AddOns: Add docs_link attribute
  - Update to metamodel v0.0.28

## 0.1.102 May 15 2020

- Update to model v0.0.57:
  - AddOnInstallations: Remove DELETE operation
  - Added Label to Account

- Update to metamodel v0.0.28:
  - OpenAPI: Fix expected response

## 0.1.101 May 5 2020

- Update to model 0.0.56
  - Add Labels to Organization

## 0.1.100 Apr 23 2020

- Update to model 0.0.55
  - Add enabled field to region
  - Adding metrics.nodes to api model
  - Adding cluster ingresses endpoint
  - ClusterNodes: Add ComputeMachineType
  - Network: Added HostPrefix

## 0.1.99 Apr 7 2020

- Update to model 0.0.54
  - Add HealthState field to Cluster type
  - Refactor alerts and operator conditions to contain only 'CriticalAlerts' and 'OperatorsConditionFailing'
  - Adding computeNodesSockets to cluster metrics
  - Fix pull secret deletion path
  - Remove unsupported cluster state
  - Add machine type category

- Update to metamodel 0.0.27
  - Update file header year to 2020

## 0.1.98 Apr 6 2020

- Update to model 0.0.53
  - Add pull secret deletion
  - Products: Add product attribute to cluster object
  - Products: Support for top-level cluster types
  - Add ClusterOperatorsConditions type
  - Add ClusterAlertsFiring type and field in ClusterMetrics

## 0.1.97 Mar 26 2020

- Update to model 0.0.52
  - Add Subscription Model changes.

## 0.1.96 Mar 24 2020

- Update to model 0.0.50
  - Add Ingress type
  - Add sockets to cluster_metrics_type

## 0.1.95 Mar 22 2020

- Update to model 0.0.48:
  - Fix `OpenID` attributes.
  - Add Cluster API listening method.

## 0.1.94 Mar 19 2020

- Update to model 0.0.47:
  - Add ClusterAdminEnabled flag.
  - Add PullSecrets endpoint.
  - Fix `LDAPIdentityProvider` attribute name.

## 0.1.93 Mar 18 2020

- Update to model 0.0.46:
  - Add missing fields for add-on installation
  - Add operator name to add-ons

## 0.1.92 Mar 11 2020

- Update to model 0.0.45:
  - Add Organizations field to GitHub IDP

## 0.1.91 Mar 5 2020

- Update to model 0.0.42:
  - Add `client_secret` attribute to _GitHub_ identity provider.

## 0.1.90 Mar 2 2020

- Request new tokens when the _OpenID_ server returns error code `invalid_grant`
during the refresh token grant.

- Check that responses from the _OpenID_ server contain `application/json` in
the `Content-Type` header, and improve the error messages generated in that
case so that they contain a summary of the content.

- Honor cookies sent by the _OpenID_ and API servers.

## 0.1.89 Feb 26 2020

- Update to metamodel 0.0.26.

  The more relevant change in the new version of the metamodel is the new
  `operation_id` attribute added to error objects and error messages. An error
  object like this:

  ```json
  {
    "kind": "Error",
    "id": "401",
    "href": "/api/clusters_mgmt/v1/errors/401",
    "code": "CLUSTERS-MGMT-401",
    "reason": "My reason",
    "operation_id": "456"
  }
  ```

  Will result in the following error string (in one single line):

  ```
  identifier is '401', code is 'CLUSTERS-MGMT-401' and
  operation identifier is '456': My reason
  ```

  This addresses issue [150](https://github.com/openshift-online/ocm-sdk-go/issues/150).

## 0.1.88 Feb 20 2020

- Remove _service_ and _version_ parameters from the builder of the
  authentication handler. This is a backwards compatibility breaking change
  that requires changes in the code that creates the authentication handler. For
  example, if the current code is like this:

  ```go
  handler, err := authentication.NewHandler().
          Logger(logger).
          Service("clusters_mgmt").
          Version("v1").
          Public("...").
          KeysFile("...").
          KeysURL("...").
          ACLFile("...").
          Next(next).
          Build()
  if err != nil {
          ...
  }
  ```

  It will need to be changed to this:

   ```go
  handler, err := authentication.NewHandler().
          Logger(logger).
          Public("...").
          KeysFile("...").
          KeysURL("...").
          ACLFile("...").
          Next(next).
          Build()
  if err != nil {
          ...
  }
  ```

  Note that the only change required is removing the calls to the `Service` and
  `Version` methods of the builder. The handler will now extract those values
  from the request URL.

  This is specially important for programs that use the same authentication
  handler for multiple services.

- Update to metamodel 0.0.25:
  - Run the `gofmt` command only once for all generated files instead of running
   it once per each generated file.
  - Avoid generating code with constructs that would then be simplified by the
   `-s` flag of the `gofmt` command.

## 0.1.87 Feb 14 2020

- Preserver order of attributes of JSON documents sent to the log when debug
  mode is enabled.
- Update to metamodel 0.0.24:
  - Add `Content-Type` to responses sent by the generated server code.
  - Don't require developer to explicitly remove the `/api` when using the
   server code.
  - Remove redundant quotes from error responses sent by the generated
   server code.

## 0.1.86 Feb 13 2020

- Update to model 0.0.41:
  - Add `target_namespace` and `install_mode` attributes to `AddOn` type.
  - Add `state` attribute to `AWSInfrastructureAccessRole` type.

## 0.1.85 Feb 12 2020

- Update to metamodel 0.0.23:
  - Fix missing _OpenAPI_ paths due to incorrect use of `append`.

## 0.1.84 Feb 5 2020

- Add method to update flavour.

## 0.1.83 Feb 3 2020

- Check content type of HTTP responses and return an error if it isn't JSON.
- Update to model 0.0.39:
  - Add types and resources for cluster operator metrics.
  - Add `deleting` and `removed` states to AWS infrastructure access role grant
   status.

## 0.1.82 Jan 23 2020

- Update to model 0.0.38:
  - Add `search` and `order` parameters to the method that lists registry
   credentials.
  - Add `labels` parameter to the method that lists subscriptions.
  - Add types and resources for management of AWS infrastructure access roles.

## 0.1.81 Jan 16 2020

- Add ability to intercept request and response using a transport middleware
   of type `http.RoundTripper`.

## 0.1.80 Jan 13 2020

- Add body details in case of error from token provider.

## 0.1.79 Jan 9 2020

- Update to metamodel 0.0.22:
  - Fix generation of _OpenAPI_ paths so that all the characters are lower case.

## 0.1.78 Jan 8 2020

- Fix URL prefix for the logs service.
- Update to metamodel 0.0.21:
  - Use JSON iterator instead of the default JSON Go package.

## 0.1.77 Jan 8 2020

- Don't require Go 1.13.
- Update to model 0.0.37:
  - Add new `service_logs` service.
  - Add types and resources for machine types.

## 0.1.76 Jan 3 2020

- Update to model 0.0.36:
  - Add types and resources for AWS infrastructure access roles.
  - Add GCP flavour and change AWS flavour to contain also the instance type.

## 0.1.75 Jan 1 2020

- Update to model 0.0.35:
  - Add `CurrentAccess` support.

## 0.1.74 Dec 31 2019

- Update to model 0.0.33:
  - Add the `CreatedAt` and `UpdatedAt` attributes to the `Subscription` type.

## 0.1.73 Dec 24 2019

- Update to model 0.0.32:
  - Replace `AddOns` with `AddOnInstallations`.

## 0.1.72 Dec 19 2019

- Update to model 0.0.31:
  - Add `ban_code` attribute to `Account` type.

## 0.1.71 Dec 19 2019

- Authentication handler sends 401 instead of 511.
- Authentication handler sends the `WWW-Authenticate` response header.
- Authentication handler doesn't send authentication failures to the log.

## 0.1.70 Dec 18 2019

- Update to metamodel 0.0.20:
  - Fix conversion of errors to JSON so that the `kind` attribute is generated
   correctly.

- Add authentication handler.

## 0.1.69 Dec 17 2019

- Update to model 0.0.30:
  - Add support for `ClusterUUID` field.

## 0.1.68 Dec 12 2019

- Update to metamodel 0.0.19:
  - Don't fail on wrong kind.

## 0.1.67 Dec 12 2019

- Don't check kinds of add-ons installations.

## 0.1.66 Dec 12 2019

- Update to model 0.0.29:
  - Allow subscription identifier on role binding.

## 0.1.65 Dec 10 2019

- Update to model 0.0.28:
  - Add `AddOnInstallation` type.

## 0.1.64 Dec 4 2019

- Update to model 0.0.27:
  - Add `resource_name` and `resource_cost` attributes to the add-on type.

## 0.1.63 Dec 2 2019

- Update to model 0.0.26:
  - Remove obsolete `aws` and `version` fields from the `Flavour` type.
  - Add instance type fields to the `Flavour` type.
  - Add `AWSVolume` and `AWSFlavour` types.
  - Add attributes required for _BYOC_.
  - Fix direction of `Body` parameters of updates.

## 0.1.62 Nov 28 2019

- Update to model 0.0.25:
  - Allow patching role binding.

## 0.1.61 Nov 25 2019

- Update to metamodel 0.0.18:
  - Add stage URL and `securitySchemes` to the generated _OpenAPI_
   specifications.

## 0.1.60 Nov 23 2019

- Update to model 0.0.24:
  - Fix directions of paging parameters.
  - Fix direction of `Body` parameter of `Update`.
  - Add default values to paging parameters.
  - Update to metamodel 0.0.17.

- Update to metamodel 0.0.17:
  - Add semantic checks.
  - Add support for default values.
  - Check default values of paging parameters.

## 0.1.59 Nov 20 2019

- Update to model 0.0.23:
  - Add infra nodes to `FlavourNodes`.
  - Refactor flavour nodes.

## 0.1.58 Nov 19 2019

- Update to metamodel 0.0.16:
  - Add simple conversion from AsciiDoc to Markdown.

## 0.1.57 Nov 19 2019

- Update to metamodel 0.0.15:
  - Add support for the version metadata resource.

## 0.1.56 Nov 19 2019

- Update to model 0.0.22:
  - Add `socket_total_by_node_roles_os` metric query.

## 0.1.55 Nov 17 2019

- Update to model 0.0.21:
  - Added add-on resources and types.
  - Added subscription reserved resources collection.

## 0.1.54 Nov 17 2019

- Drop support for _developers.redhat.com_.

- Update to metamodel 0.0.14:
  - Add `Poll` method to clients that have a `Get` method.

## 0.1.53 Nov 14 2019

- Update to model 0.0.20:
  - Query resource quota from root and delete by identifier.

- Update to metamodel 0.0.13:
  - Fix imports of `helpers` and `errors` packages.

## 0.1.52 Nov 8 2019

- Update to model 0.0.19:
  - Added identifiers to role binding type.

## 0.1.51 Nov 7 2019

- Update to model 0.0.18:
  - Added support to search role bindings and resource quota.

## 0.1.50 Nov 4 2019

- Update to metamodel 0.0.12:
  - Add _OpenAPI_ specification generator.

## 0.1.49 Oct 28 2019

- Update to model 0.0.17:
  - Added `Disconnected`, `DisplayName` and `ExternalClusterID` attributes to the
   cluster authorization request type.

## 0.1.48 Oct 27 2019

- Update to model 0.0.16:
  - Added `ResourceReview` resource to the authorizations service.

- Update to metamodel 0.0.11:
  - Improve parsing of initialisms.
  - Fix the method not allowed code.
  - Send not found when server returns `nil` target.
  - Generate service and version servers.
  - Don't generate files with execution permission.

## 0.1.47 Oct 25 2019

- Update to metamodel 0.0.10:
  - Make HTTP adapters stateless.

## 0.1.46 Oct 24 2019

- Update to model 0.0.15:
  - Added `search` parameter to the accounts `List` method.

## 0.1.45 Oct 24 2019

- Update to model 0.0.14:
  - Added `SKU` type.
  - Improved organizations.
  - Improved roles.

## 0.1.44 Oct 15 2019

- Upate to model 0.0.13:
  - Added `AccessTokenAuth` type.
  - Added `auths` attribute to `AccessToken` type.

- Update to metamodel 0.0.9:
  - Generate shorter adapter names.
  - Use constants from the `http` package.
  - Shorter _read_ and _write_ names.
  - Rename `SetStatusCode` to `Status`.
  - Improve naming of variables.
  - Set default status.
  - Move errors and helpers generators to separate files.

## 0.1.43 Oct 10 2019

- Update to model 0.0.12:
  - Add `access_review` resource.

## 0.1.41 Oct 10 2019

- Update to model 0.0.11:
  - Add `export_control_review` resource.

## 0.1.40 Oct 7 2019

- Update to model 0.0.10:
  - Add `cpu_total_by_node_roles_os` metric query.

## 0.1.39 Oct 7 2019

- Update to model 0.0.9:
  - Add `type` attribute to the `ResourceQuota` type.
  - Add `config_managed` attribute to the `RoleBinding` type.

## 0.1.38 Sep 17 2019

- Update to model 0.0.8:
  - Update methods don't return body.

## 0.1.37 Sep 16 2019

- Update to model 0.0.7:
  - Add `search` parameter to the `List` method of the subscriptions resource.

## 0.1.36 Sep 16 2019

- Update to model 0.0.6:
  - Remove the `creator` attribute of the `Cluster` type.

- Update to metamodel 0.0.7:
  - Add `Copy` method to builders.

## 0.1.35 Sep 12 2019

- Update to model 0.0.5:
  - Add `order` parameter to the methods to list accounts and subscriptions.

## 0.1.34 Sep 11 2019

- Use access token that is about to expire if there is no other mechanism to
  obtain a new one.

- Update to model 0.0.3:
  - Add `order` parameter to the collections that suport it.
  - Add cloud providers collection.

## 0.1.33 Sep 10 2019

- Update to model 0.0.2:
  - Add `DisplayName` attribute to `Subscription` type.

- Update to metamodel 0.0.5:
  - Fix generation of field names for query parameters.
  - Remove `query` and `path` fields from request objects.
  - Remove unused imports.

## 0.1.32 Sep 03 2019

- Makefile generates code using the ocm-api-metamodel v0.0.4.

- Generated servers parse request query parameters.

## 0.1.31 Aug 28 2019

- Generated servers enforce no trailing slashes as well send `Content-Type` header.

## 0.1.30 Aug 27 2019

- Renamed package to `github.com/openshift-online/ocm-sdk-go`.

## 0.1.29 Aug 26 2019

- Generated servers can handle routes with and without trailing slashes.

- Clone metamodel for code generation

- Clone model for code generation

- Rename main package

## 0.1.28 Aug 22 2019

- Add Context parameter to Server methods.

## 0.1.27 Aug 22 2019

- Add generated servers.

- Changes ClusterRegistration response type from long to string .

## 0.1.26 Aug 13 2019

- Add support for the `compute_nodes_cpu` and `compute_nodes_memory` metrics.

## 0.1.25 Aug 11 2019

- Add support for quota summary.

- Fix the data type of the cluster registration expiration date.

## 0.1.24 Jun 28 2019

- Automatically select the deprecated _OpenID_ server when authenticating with
  user name and password.

## 0.1.23 Jun 27 2019

- Don't show cluster admin credentials in the debug log.

## 0.1.22 Jun 27 2019

- Don't send warnings about toke issuer when no tokens are used.

- Fix the names of the methods used to set the V values of the `glog` logger.

## 0.1.21 Jun 26 2019

- Added methods to get connection attributes like token URL, client identifier,
  etc.

## 0.1.20 Jun 26 2019

- Switch from `developers.redhat.com` to `sso.redhat.com`.

## 0.1.19 Jun 25 2019

- Added `GetMethod` and `GetPath` methods to HTTP requests.

- Added `Header` method to HTTP responses.

## 0.1.18 Jun 21 2019

- Added support for the `expiration_timestamp` attribute of the `Cluster` type.

## 0.1.17 Jun 20 2019

- Added support for the `name` attribute of the `Dashboard` type.

- Added to lists a new `Get` method to get elements by index.

## 0.1.16 Jun 19 2019

- Added to response types getter methods that return the value of the parameter
  and a boolean flag that indicates if there is actually a value.

## 0.1.15 Jun 19 2019

- Add support for the `versions` collection.

## 0.1.14 Jun 4 2019

- Redact sensitive fields in debug logs.

- Don't crash when there is no response.

## 0.1.13 May 22 2019

- Added support for building objects with attributes that are lists of structs.

## 0.1.12 May 20 2019

- Added support for deleting subscriptions.

- Added Prometheus metrics.

## 0.1.11 May 15 2019

- Increase token slack to one minute.

## 0.1.10 May 8 2019

- Improved support for contexts, adding the `BuildContext`, `TokensContext` and
  `SendContext` methods.

IMPORTANT: This version breaks backwards compatibility in the `Logger`
interface, as all the methods require now a first `ctx` parameter.

## 0.1.9 May 3 2019

- Added cluster credentials resource.

## 0.1.8 May 2 2019

- Moved basic cluster metrics to the `metrics` attribute.

- Added `Empty` method to lists and struct typess.

## 0.1.7 May 1 2019

- Always close connections used to request access tokens.

## 0.1.6 Apr 23 2019

- Add typed interface.

## 0.1.5 Apr 17 2019

- Changed package path to `github.com/openshift-online/uhc-sdk-go`.

## 0.1.4 Apr 3 2019

- Don't panic when no refresh token is provided.

## 0.1.3 Mar 27 2019

- Don't close body in round tripper.

## 0.1.2 Mar 23 2019

- Add support for offline access tokens.

## 0.1.1 Jan 25 2019

- Change the `glog` logger so that it uses `--v=0` for errors, warnings and
  information messages and `--v=1` for debug messages.

## 0.1.0 Jan 24 2019

- Renamed the project from `api-client` to `uhc-sdk`.

- Moved the command line tool to a new `uhc-cli` project.

## 0.0.13 Jan 24 2019

- Add `context` and `timeout` parameters to all requests.

- Scrub password from debug log.

## 0.0.12 Dec 19 2018

- Add `TrustedCAs` parameter to the connection builder.

## 0.0.11 Dec 17 2018

- Check that `T` is passed to the testing logger.

## 0.0.10 Nov 27 2018

- Implement terminal check correctly for _macOS_.

## 0.0.9 Nov 22 2018

- Don't include the testing logger in the binary.

- Added support for printing refresh tokens.

- Added support for setting the _OpenID_ scopes.

- Added a new `StdLogger` that sends log messages to the standard output and
  error streams.
